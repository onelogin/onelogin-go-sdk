package customerrors

import (
	"errors"
	"fmt"
	"net/http"
)

// RequestError used for http request errors.
type RequestError struct {
	Context    string
	StatusCode int
	Err        error
}

// ReqErrorWrapper creates a new Request error and returns,
// the pointer to the request error.
func ReqErrorWrapper(resp *http.Response, err error, context string) *RequestError {
	code := 0
	errToUse := err

	if resp != nil {
		code = resp.StatusCode
	}

	if errToUse == nil && code >= http.StatusBadRequest {
		errToUse = errors.New(http.StatusText(code))
	}

	if errToUse == nil {
		return nil
	}

	return &RequestError{
		StatusCode: code,
		Err:        errToUse,
		Context:    context,
	}
}

func (err *RequestError) Error() string {
	errMsg := ""
	if err.Err != nil {
		errMsg = err.Err.Error()
	}
	return fmt.Sprintf("request error: context: %s, status_code: [%d], error_message: %s", err.Context, err.StatusCode, errMsg)
}
