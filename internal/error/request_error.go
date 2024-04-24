package error

import (
	"fmt"
)

type RequestError struct {
	Message string
}

func (e RequestError) Error() string {
	return fmt.Sprintf("Request error: %s", e.Message)
}

func NewRequestError(message string) *RequestError {
	return &RequestError{
		Message: message,
	}
}
