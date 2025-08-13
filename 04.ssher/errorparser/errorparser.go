package errorparser

import (
	"errors"
	"strconv"
)

type ErrorHandler struct {
}

type baseError struct {
	Code        int
	Description string
}

func newError(code int, description string) *baseError {
	return &baseError{
		Code:        code,
		Description: description,
	}
}

func NewErrorHandler() *ErrorHandler {
	return &ErrorHandler{}
}
func (e *ErrorHandler) Send(code int, description string) error {
	newError := newError(code, description)
	return errors.New("Code: " + strconv.Itoa(newError.Code) + " Message: " + newError.Description)
}
