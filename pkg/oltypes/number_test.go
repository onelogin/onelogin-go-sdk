package oltypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInt32Val(t *testing.T) {
	positiveInt32Value := int32(1)
	zeroInt32Value := int32(0)
	negativeInt32Value := int32(-1)

	tests := map[string]struct {
		input                    *int32
		expectedValidationOutput bool
		expectedOutput           int32
	}{
		"should return value and true validation when positive passed in": {
			input:                    Int32(positiveInt32Value),
			expectedValidationOutput: true,
			expectedOutput:           positiveInt32Value,
		},
		"should return value and true validation when negative passed in": {
			input:                    Int32(negativeInt32Value),
			expectedValidationOutput: true,
			expectedOutput:           negativeInt32Value,
		},
		"should return value and true validation when zero passed in": {
			input:                    Int32(zeroInt32Value),
			expectedValidationOutput: true,
			expectedOutput:           zeroInt32Value,
		},
		"should return zero and false validation when nil passed in as value": {
			input:                    nil,
			expectedValidationOutput: false,
			expectedOutput:           0,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			resultVal, isValid := GetInt32Val(test.input)
			assert.Equal(t, test.expectedValidationOutput, isValid)
			assert.Equal(t, test.expectedOutput, resultVal)
		})
	}

}
func TestGetInt64Val(t *testing.T) {
	positiveInt64Value := int64(1)
	zeroInt64Value := int64(0)
	negativeInt64Value := int64(-1)

	tests := map[string]struct {
		input                    *int64
		expectedValidationOutput bool
		expectedOutput           int64
	}{
		"should return value and true validation when positive value passed in": {
			input:                    Int64(positiveInt64Value),
			expectedValidationOutput: true,
			expectedOutput:           positiveInt64Value,
		},
		"should return value and true validation when zero passed in": {
			input:                    Int64(zeroInt64Value),
			expectedValidationOutput: true,
			expectedOutput:           zeroInt64Value,
		},
		"should return value and true validation when negative value passed in": {
			input:                    Int64(negativeInt64Value),
			expectedValidationOutput: true,
			expectedOutput:           negativeInt64Value,
		},
		"should return zero and false validation when nil passed in as value": {
			input:                    nil,
			expectedValidationOutput: false,
			expectedOutput:           0,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			resultVal, isValid := GetInt64Val(test.input)
			assert.Equal(t, test.expectedValidationOutput, isValid)
			assert.Equal(t, test.expectedOutput, resultVal)
		})
	}

}
