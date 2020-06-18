package customerrors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackedErrors(t *testing.T) {
	errMsgs := []error{errors.New("Error 1"), errors.New("Error 2")}

	tests := map[string]struct {
		inputErr    []error
		expectedMsg string
	}{
		"error passed in": {
			inputErr:    errMsgs,
			expectedMsg: "Error 1, Error 2",
		},
		"error not passed in": {
			inputErr:    nil,
			expectedMsg: "",
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			result := StackErrors(test.inputErr)
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
