package errors

// swagger:response ErrorResponse
type ErrorResponse Error

func New(code string) error {
	return &Error{Code: code}
}

// swagger:model Error
type Error struct {
	Code string
}

func (e *Error) Error() string {
	return e.Code
}
