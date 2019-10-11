package customerrors

import (
	"errors"
	"fmt"
	"net/http"
)

// RequestError used for http request errors.
type RequestError struct {
	context    string
	statusCode int
	err        error
}

// ReqErrorWrapper creates a new Request error and returns,
// the pointer to the request error.
func ReqErrorWrapper(resp *http.Response, err error, context string) error {
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
		statusCode: code,
		err:        errToUse,
		context:    context,
	}
}

func (errReq *RequestError) Error() string {
	errMsg := ""
	if errReq.err != nil {
		errMsg = errReq.err.Error()
	}

	if errReq.statusCode == 0 {
		return fmt.Sprintf("request error: context: %s, error_message: %s", errReq.context, errMsg)
	}

	return fmt.Sprintf("request error: context: %s, status_code: [%d], error_message: %s", errReq.context, errReq.statusCode, errMsg)
}
