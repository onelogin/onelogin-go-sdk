package customerrors

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReqErrorWrapper(t *testing.T) {
	context := "test"
	testErrMsg := "Test error"

	tests := map[string]struct {
		responseHttp   *http.Response
		err            error
		context        string
		expectedErrMsg string
	}{
		"http response not nil and status code >= 400 and no error passed in": {
			responseHttp:   &http.Response{StatusCode: http.StatusBadRequest},
			err:            nil,
			context:        context,
			expectedErrMsg: fmt.Sprintf("request error: context: %s, status_code: [%d], error_message: %s", context, http.StatusBadRequest, http.StatusText(http.StatusBadRequest)),
		},
		"err but no http response": {
			responseHttp:   nil,
			err:            errors.New(testErrMsg),
			context:        context,
			expectedErrMsg: fmt.Sprintf("request error: context: %s, error_message: %s", context, testErrMsg),
		},
		"no http response and no error passed in": {
			responseHttp:   nil,
			err:            nil,
			context:        context,
			expectedErrMsg: "",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			resultErr := ReqErrorWrapper(test.responseHttp, test.context, test.err)
			if test.responseHttp == nil && test.err == nil {
				assert.Nil(t, resultErr)
			} else {
				assert.NotNil(t, resultErr)
				assert.Equal(t, test.expectedErrMsg, resultErr.Error())
			}
		})
	}
}
