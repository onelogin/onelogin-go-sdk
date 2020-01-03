package oltypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBoolVal(t *testing.T) {
	tests := map[string]struct {
		input                    *bool
		expectedValidationOutput bool
		expectedOutput           bool
	}{
		"should return true validation and output when value passed in": {
			input:                    Bool(true),
			expectedValidationOutput: true,
			expectedOutput:           true,
		},
		"should return false validation and output when nil passed in as value": {
			input:                    nil,
			expectedValidationOutput: false,
			expectedOutput:           false,
		},
		"should return true validation and false output when false passed in as value": {
			input:                    Bool(false),
			expectedValidationOutput: true,
			expectedOutput:           false,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			resultVal, isValid := GetBoolVal(test.input)
			assert.Equal(t, test.expectedValidationOutput, isValid)
			assert.Equal(t, test.expectedOutput, resultVal)
		})
	}

}
