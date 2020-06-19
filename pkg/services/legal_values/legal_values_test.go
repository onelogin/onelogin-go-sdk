package legalvalues

import (
	"encoding/json"
	"errors"
	"github.com/onelogin/onelogin-go-sdk/internal/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLegalValues(t *testing.T) {
	tests := map[string]struct {
		TestClient    *test.MockRepository
		ExpectedOut   []map[string]string
		ExpectedError error
	}{
		"retrieves list of legal values": {
			TestClient: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal([]map[string]string{{"key": "key", "value": "value"}, {"key": "key2", "value": "value2"}})
				},
			},
			ExpectedOut: []map[string]string{{"key": "key", "value": "value"}, {"key": "key2", "value": "value2"}},
		},
		"returns an empty array and error when api call fails": {
			TestClient: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return nil, errors.New("error")
				},
			},
			ExpectedOut:   []map[string]string{},
			ExpectedError: errors.New("error"),
		},
		"returns an empty array and error when api response is not expected shape": {
			TestClient: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(map[string]string{"key": "key"})
				},
			},
			ExpectedOut: []map[string]string{},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.TestClient, "test.com")
			actual := make([]map[string]string, 0)
			err := svc.Query("stuff", &actual)
			assert.Equal(t, test.ExpectedOut, actual)
			if test.ExpectedError != nil {
				assert.Equal(t, test.ExpectedError, err)
			}
		})
	}
}
