package customerrors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneloginErrorWrapperWithErrorPassedIn(t *testing.T) {
	context := "Test"
	errMsg := "Test Error"
	expectedErrMsg := fmt.Sprintf("error: context: [%s], error_message: [%s]", context, errMsg)

	err := errors.New(errMsg)

	outputErr := OneloginErrorWrapper(context, err)
	assert.NotNil(t, outputErr)
	assert.Equal(t, expectedErrMsg, outputErr.Error())
}

func TestOneloginErrorWrapperWithNilErrorPassedIn(t *testing.T) {
	context := "Test"

	var err error = nil

	outputErr := OneloginErrorWrapper(context, err)
	assert.Equal(t, nil, outputErr)
}
