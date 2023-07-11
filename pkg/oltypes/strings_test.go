package oltypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetStringVal(t *testing.T) {
	validStrValue := "test"

	tests := map[string]struct {
		input                    *string
		expectedValidationOutput bool
		expectedOutput           string
	}{
		"should return value when value passed in": {
			input:                    String(validStrValue),
			expectedValidationOutput: true,
			expectedOutput:           validStrValue,
		},
		"should return false when nil passed in as value": {
			input:                    nil,
			expectedValidationOutput: false,
			expectedOutput:           EmptyString,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			resultVal, isValid := GetStringVal(test.input)
			assert.Equal(t, test.expectedValidationOutput, isValid)
			assert.Equal(t, test.expectedOutput, resultVal)
		})
	}

}
