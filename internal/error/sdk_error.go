package error

import "fmt"

type SDKError struct {
	Message string
}

func (e *SDKError) Error() string {
	return fmt.Sprintf("SDK error: %s", e.Message)
}

func NewSDKError(message string) *SDKError {
	return &SDKError{
		Message: message,
	}
}
