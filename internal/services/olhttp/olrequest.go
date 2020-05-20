package olhttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/onelogin/onelogin-go-sdk/internal/customerrors"
)

const olRequestContext = "ol request"

var errInvalidRequestInput = errors.New("Invalid input for request creation")

// allowed http methods
var allowedMethods = map[string]string{
	http.MethodGet:    http.MethodGet,
	http.MethodPost:   http.MethodPost,
	http.MethodPut:    http.MethodPut,
	http.MethodDelete: http.MethodDelete,
}

// NewOneloginRequest sets up the http request with url, method, headers, and the payload, if one is needed. If successful, it returns
// the http request.
func NewOneloginRequest(url string, method string, headers map[string]string, payload interface{}) (*http.Request, error) {
	methodToUse := strings.ToUpper(method)

	// make sure that the method is valid and allowed
	if _, isThere := allowedMethods[methodToUse]; !isThere {
		return nil, customerrors.OneloginErrorWrapper(olRequestContext, errInvalidRequestInput)
	}

	var req *http.Request
	var reqErr error

	// add payload if put or post
	if (methodToUse == http.MethodPost || methodToUse == http.MethodPut) && payload != nil {
		bodyToSend, marshErr := json.Marshal(payload)
		// make sure marshaling was successful
		if err := customerrors.OneloginErrorWrapper(olRequestContext, marshErr); err != nil {
			return nil, err
		}
		req, reqErr = http.NewRequest(methodToUse, url, bytes.NewBuffer(bodyToSend))
	} else {
		req, reqErr = http.NewRequest(methodToUse, url, nil)
	}

	if err := customerrors.OneloginErrorWrapper(olRequestContext, reqErr); err != nil {
		return nil, err
	}

	// set headers
	for key, val := range headers {
		req.Header.Set(key, val)
	}

	return req, nil
}
