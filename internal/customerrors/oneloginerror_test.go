package customerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneLoginErrorWrapper(t *testing.T) {
	errMsg := "Test Error"
	context := "Test"

	tests := map[string]struct {
		context     string
		inputErr    error
		expectedMsg string
	}{
		"error passed in": {
			context:     context,
			inputErr:    errors.New(errMsg),
			expectedMsg: fmt.Sprintf("error: context: [%s], error_message: [%s]", context, errMsg),
		},
		"error not passed in": {
			context:     context,
			inputErr:    nil,
			expectedMsg: "",
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			result := OneloginErrorWrapper(test.context, test.inputErr)

			if test.inputErr != nil {
				assert.NotNil(t, result)
				assert.Equal(t, test.expectedMsg, result.Error())
			} else if test.inputErr == nil {
				assert.Nil(t, result)
			} else {
				assert.Fail(t, "Test should not reach here")
			}
		})
	}

}
