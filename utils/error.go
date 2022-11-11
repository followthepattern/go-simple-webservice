package utils

import (
	"net/http"
)

type Error struct {
	code    int
	message string
}

func New(code int, message string) error {
	return &Error{code: code, message: message}
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) Code() int {
	return e.code
}

func GetErrorCode(err error) int {
	switch err.(type) {
	case *Error:
		return err.(*Error).Code()
	}
	return http.StatusBadRequest
}
