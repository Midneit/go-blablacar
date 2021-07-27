package blablacar

import (
	"context"
	"strconv"
)

// SearchRequest struct.
type SearchRequest struct {
	FromCoordinate string
	ToCoordinate   string
	Locale         string
	Currency       string
	FromCursor     string
	Count          int
	StartDateLocal string
	EndDateLocal   string
	RequestedSeats int
	RadiusInMeters int
	Sort           string
}

// SearchResponse from API.
type SearchResponse struct {
	Link       string       `json:"link"`
	SearchInfo SearchInfo   `json:"search_info"`
	Trips      []SearchTrip `json:"trips"`
	NextCursor string       `json:"next_cursor"`
}

// Search API method.
//
// https://support.blablacar.com/hc/en-gb/articles/360014199820--Search-V3-API-Documentation
func (c *Client) Search(ctx context.Context, request *SearchRequest) (*SearchResponse, error) {
	u := "api/v3/trips"

	req, err := c.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()

	q.Add("from_coordinate", request.FromCoordinate)
	q.Add("to_coordinate", request.ToCoordinate)

	if request.Locale != "" {
		q.Add("locale", request.Locale)
	}

	if request.Currency != "" {
		q.Add("currency", request.Currency)
	}

	if request.FromCursor != "" {
		q.Add("from_cursor", request.FromCursor)
	}

	if request.Count != 0 {
		q.Add("count", strconv.Itoa(request.Count))
	}

	if request.StartDateLocal != "" {
		q.Add("start_date_local", request.StartDateLocal)
	}

	if request.EndDateLocal != "" {
		q.Add("end_date_local", request.EndDateLocal)
	}

	if request.RequestedSeats != 0 {
		q.Add("requested_seats", strconv.Itoa(request.RequestedSeats))
	}

	if request.RadiusInMeters != 0 {
		q.Add("radius_in_meters", strconv.Itoa(request.RadiusInMeters))
	}

	if request.Sort != "" {
		q.Add("sort", request.Sort)
	}

	req.URL.RawQuery = q.Encode()

	response := SearchResponse{}

	_, err = c.Do(ctx, req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
