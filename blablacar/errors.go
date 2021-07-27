package blablacar

// ErrorType is a type of error.
type ErrorType int

const (
	ErrInvalidToken     ErrorType = 1
	ErrRouteNotFound    ErrorType = 2
	ErrMethodNotAllowed ErrorType = 3
	ErrTripNotFound     ErrorType = 4
	ErrMalformedRequest ErrorType = 5
)

// Error struct.
type Error struct {
	Code    ErrorType
	Message string
}

// Error returns message of error.
func (e *Error) Error() string {
	return "go-blablacar: " + e.Message
}
