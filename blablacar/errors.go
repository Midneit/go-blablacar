package blablacar

type ErrorType int

const (
	ErrInvalidToken     ErrorType = 1
	ErrRouteNotFound    ErrorType = 2
	ErrMethodNotAllowed ErrorType = 3
	ErrTripNotFound     ErrorType = 4
	ErrMalformedRequest ErrorType = 5
)

type Error struct {
	Code    ErrorType
	Message string
}

func (e *Error) Error() string {
	return "go-blablacar: " + e.Message
}
