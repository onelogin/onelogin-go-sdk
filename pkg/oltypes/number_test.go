package oltypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInt32Val(t *testing.T) {
	validInt32Value := int32(1)

	tests := map[string]struct {
		input                    *int32
		expectedValidationOutput bool
		expectedOutput           int32
	}{
		"should return value when value passed in": {
			input:                    Int32(validInt32Value),
			expectedValidationOutput: true,
			expectedOutput:           validInt32Value,
		},
		"should return false when nil passed in as value": {
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
	validInt64Value := int64(1)

	tests := map[string]struct {
		input                    *int64
		expectedValidationOutput bool
		expectedOutput           int64
	}{
		"should return value when value passed in": {
			input:                    Int64(validInt64Value),
			expectedValidationOutput: true,
			expectedOutput:           validInt64Value,
		},
		"should return false when nil passed in as value": {
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
