package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplaceSpecialChar(t *testing.T) {
	tests := map[string]struct {
		InputStr    string
		InputRep    string
		ExpectedOut string
	}{
		"It replaces the non alpha-numeric with the specified character": {
			InputStr:    "stuff+test",
			InputRep:    "&",
			ExpectedOut: "stuff&test",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := ReplaceSpecialChar(test.InputStr, test.InputRep)
			assert.Equal(t, test.ExpectedOut, actual)
		})
	}
}

func TestToSnakeCase(t *testing.T) {
	tests := map[string]struct {
		InputStr    string
		ExpectedOut string
	}{
		"It converts the PascalCase to snake_case": {
			InputStr:    "PascalCase",
			ExpectedOut: "pascal_case",
		},
		"It converts the cascalCase to snake_case": {
			InputStr:    "camelCase",
			ExpectedOut: "camel_case",
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := ToSnakeCase(test.InputStr)
			assert.Equal(t, test.ExpectedOut, actual)
		})
	}
}

func TestOneOf(t *testing.T) {
	validOpts := []string{"SHA-1", "SHA-256", "SHA-348", "SHA-512"}
	tests := map[string]struct {
		InputKey       string
		InputValue     string
		ExpectedOutput error
	}{
		"no errors on valid input": {
			InputKey:       "signature_algorithm",
			InputValue:     "SHA-1",
			ExpectedOutput: nil,
		},
		"errors on invalid input": {
			InputKey:       "signature_algorithm",
			InputValue:     "asdf",
			ExpectedOutput: fmt.Errorf("signature_algorithm must be one of %v, got: %s", validOpts, "asdf"),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			errs := OneOf(test.InputKey, test.InputValue, validOpts)
			assert.Equal(t, test.ExpectedOutput, errs)
		})
	}
}
