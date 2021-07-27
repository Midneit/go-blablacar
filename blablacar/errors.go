package blablacar

// ErrorType is a type of error.
type ErrorType int

const (
	// ErrInvalidToken error.
	ErrInvalidToken ErrorType = 1
	// ErrRouteNotFound error
	ErrRouteNotFound ErrorType = 2
	// ErrMethodNotAllowed error
	ErrMethodNotAllowed ErrorType = 3
	// ErrTripNotFound error
	ErrTripNotFound ErrorType = 4
	// ErrMalformedRequest error
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
