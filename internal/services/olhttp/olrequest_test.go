package olhttp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestIns struct {
	url     string
	method  string
	headers map[string]string
	payload interface{}
}

type ExpectedOuts struct {
	expectedURL     string
	expectedHeaders http.Header
	expectedError   error
	expectedMethod  string
	expectedPayload interface{}
}

func TestNewOneloginRequest(t *testing.T) {
	tests := map[string]struct {
		Ins  TestIns
		Outs ExpectedOuts
	}{
		"returns a request without payload": {
			Ins: TestIns{
				url:    "test",
				method: http.MethodGet,
				headers: map[string]string{
					"Content-Type":  "application/json",
					"Authorization": "Bearer test",
				},
			},
			Outs: ExpectedOuts{
				expectedURL:    "test",
				expectedMethod: http.MethodGet,
				expectedHeaders: http.Header{
					"Content-Type":  []string{"application/json"},
					"Authorization": []string{"Bearer test"},
				},
			},
		},
		"returns a request with payload": {
			Ins: TestIns{
				url:     "test",
				method:  http.MethodPost,
				payload: map[string]string{"Name": "test"},
				headers: map[string]string{
					"Content-Type":  "application/json",
					"Authorization": "Bearer test",
				},
			},
			Outs: ExpectedOuts{
				expectedURL:     "test",
				expectedMethod:  http.MethodPost,
				expectedPayload: map[string]string{"Name": "test"},
				expectedHeaders: http.Header{
					"Content-Type":  []string{"application/json"},
					"Authorization": []string{"Bearer test"},
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual, err := NewOneloginRequest(test.Ins.url, test.Ins.method, test.Ins.headers, test.Ins.payload)
			if actual.Body != nil {
				var actualBody map[string]string
				json.NewDecoder(actual.Body).Decode(&actualBody)
				defer actual.Body.Close()
				fmt.Println("sdf", actualBody)
				assert.Equal(t, test.Outs.expectedPayload, actualBody)
			}

			assert.Equal(t, test.Outs.expectedURL, actual.URL.String())
			assert.Equal(t, test.Outs.expectedHeaders, actual.Header)
			assert.Equal(t, test.Outs.expectedMethod, actual.Method)

			assert.Nil(t, err)
		})
	}
}
