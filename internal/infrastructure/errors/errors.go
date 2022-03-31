package errors

// Error is a custom error type
type Error struct {
	Msg  string
	Code int
}

// New creates new Error with custom message and code
func New(_msg string, _Code int) *Error {
	return &Error{_msg, _Code}
}

// Error returns error message
func (e *Error) Error() string {
	return e.Msg
}
