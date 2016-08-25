package errors

import "net/http"

// swagger:response ErrorResponse
type ErrorResponse Error

func New(code string) error {
	return &Error{Code: code}
}

// swagger:model Error
type Error struct {
	Code   string
	Header http.Header
}

func (e *Error) Error() string {
	return e.Code
}
