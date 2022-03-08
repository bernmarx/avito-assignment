package serviceerrors

type Error struct {
	Msg  string
	Code int
}

func New(_msg string, _Code int) *Error {
	return &Error{_msg, _Code}
}

func (e *Error) Error() string {
	return e.Msg
}
