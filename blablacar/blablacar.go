package blablacar

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://public-api.blablacar.com/"
	userAgent      = "go-blablacar"
)

type Client struct {
	client  *http.Client
	BaseURL *url.URL

	UserAgent string

	token string
}

func NewClient(token string) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)

	return &Client{
		client:    &http.Client{},
		BaseURL:   baseURL,
		UserAgent: userAgent,
		token:     token,
	}
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.BaseURL)
	}

	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	q := u.Query()
	q.Add("key", c.token)

	u.RawQuery = q.Encode()

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (c *Client) BareDo(ctx context.Context, req *http.Request) (*http.Response, error) {
	if ctx == nil {
		return nil, errors.New("ctx must be not nil")
	}

	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, err
	}

	err = CheckResponse(resp)
	if err != nil {
		defer resp.Body.Close()
	}

	return resp, err
}

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.BareDo(ctx, req)
	if err != nil {
		return resp, err
	}

	defer resp.Body.Close()

	switch v := v.(type) {
	case nil:
	case io.Writer:
		_, err = io.Copy(v, resp.Body)
	default:
		decErr := json.NewDecoder(resp.Body).Decode(v)

		if decErr == io.EOF {
			decErr = nil // ignore EOF errors caused by empty response body
		}

		if decErr != nil {
			err = decErr
		}
	}

	return resp, err
}

type messageError struct {
	Message string `json:"message"`
}

type tripError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type malformedError struct {
	Code  string `json:"code"`
	Type  string `json:"type"`
	Field string `json:"field"`
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	var err error

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	// This code written for error handling
	if r.StatusCode == http.StatusUnauthorized {
		msgErr := messageError{}

		err = json.Unmarshal(data, &msgErr)
		if err != nil {
			return err
		}

		err = &Error{
			Code:    ErrInvalidToken,
			Message: msgErr.Message,
		}
	}

	if r.StatusCode == http.StatusMethodNotAllowed {
		msgErr := messageError{}

		err = json.Unmarshal(data, &msgErr)
		if err != nil {
			return err
		}

		err = &Error{
			Code:    ErrMethodNotAllowed,
			Message: msgErr.Message,
		}
	}

	if r.StatusCode == http.StatusNotFound {
		tripError := tripError{}

		err = json.Unmarshal(data, &tripError)
		if err != nil {
			return err
		}

		if tripError.Error == "" {
			msgErr := messageError{}

			err = json.Unmarshal(data, &msgErr)
			if err != nil {
				return err
			}

			err = &Error{
				Code:    ErrRouteNotFound,
				Message: msgErr.Message,
			}
		} else {
			err = &Error{
				Code:    ErrTripNotFound,
				Message: tripError.ErrorDescription,
			}
		}
	}

	if r.StatusCode == http.StatusBadRequest {
		var malformedErrors []malformedError

		err = json.Unmarshal(data, &malformedErrors)
		if err != nil {
			return err
		}

		err = &Error{
			Code:    ErrMalformedRequest,
			Message: fmt.Sprintf("%s (%s)", malformedErrors[0].Type, malformedErrors[0].Field),
		}
	}

	return err
}
