package clientcredentials

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/onelogin/onelogin-go-sdk/internal/services"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/stretchr/testify/assert"
)

type TestMocks struct {
	mockedResponseBody interface{}
	mockedResponseCode int
	mockedService      V2Service
}

type ExpectedOuts struct {
	expectedErroMsg        string
	expectedServiceState   V2Service
	expectedStatusCode     int
	expectedAccessToken    ClientCredential
	isHTTPResponseExpected bool
}

func MockServer(mocks *TestMocks) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(mocks.mockedResponseCode)
		if mocks.mockedResponseBody != nil {
			body, _ := json.Marshal(mocks.mockedResponseBody)
			w.Write(body)
		}
	}))
	return ts
}

func TestCreate(t *testing.T) {
	tests := map[string]struct {
		Mocks TestMocks
		Outs  ExpectedOuts
	}{
		"access token is returned": {
			Mocks: TestMocks{
				mockedResponseCode: http.StatusOK,
				mockedResponseBody: ClientCredential{
					AccessToken: oltypes.String("test"),
				},
			},
			Outs: ExpectedOuts{
				expectedStatusCode: http.StatusOK,
				expectedAccessToken: ClientCredential{
					AccessToken: oltypes.String("test"),
				},
				isHTTPResponseExpected: true,
			},
		},
		"error is returned when incorrect status code": {
			Mocks: TestMocks{
				mockedResponseCode: http.StatusUnauthorized,
			},
			Outs: ExpectedOuts{
				expectedStatusCode:     http.StatusUnauthorized,
				isHTTPResponseExpected: true,
				expectedErroMsg:        "request error: context: auth v2 service, status_code: [401], error_message: Unauthorized",
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			client := &http.Client{Timeout: time.Second * 5}
			srv := MockServer(&test.Mocks)
			defer srv.Close()
			cfg := &services.AuthServiceConfig{
				ClientID:     "",
				ClientSecret: "",
				BaseURL:      srv.URL,
				Client:       client,
			}

			auth := New(cfg)

			resp, body, err := auth.Create()
			// check errors
			if test.Outs.expectedErroMsg != "" {
				assert.NotNil(t, err)
				assert.Equal(t, err.Error(), test.Outs.expectedErroMsg)
			} else {
				assert.Nil(t, err)
			}

			// check http response
			if test.Outs.isHTTPResponseExpected {
				assert.NotNil(t, resp)

				actualPayload, _ := json.Marshal(body)
				expectedPayload, _ := json.Marshal(test.Mocks.mockedResponseBody)
				assert.Equal(t, expectedPayload, actualPayload)

				// check the response status code
				assert.Equal(t, test.Outs.expectedStatusCode, resp.StatusCode)
			} else {
				assert.Nil(t, resp)
			}
		})
	}
}

func TestAuthorize(t *testing.T) {
	tests := map[string]struct {
		Mocks TestMocks
		Outs  ExpectedOuts
	}{
		"it memoizes the access token on the service": {
			Mocks: TestMocks{
				mockedResponseCode: http.StatusOK,
				mockedResponseBody: ClientCredential{
					AccessToken: oltypes.String("test"),
				},
			},
			Outs: ExpectedOuts{
				expectedServiceState: V2Service{
					Token: "test",
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			client := &http.Client{Timeout: time.Second * 5}
			srv := MockServer(&test.Mocks)
			defer srv.Close()
			cfg := &services.AuthServiceConfig{
				ClientID:     "",
				ClientSecret: "",
				BaseURL:      srv.URL,
				Client:       client,
			}

			auth := New(cfg)

			auth.Authorize()
			assert.Equal(t, test.Outs.expectedServiceState.Token, auth.Token)
		})
	}
}

func TestReAuthorize(t *testing.T) {
	tests := map[string]struct {
		Mocks TestMocks
		Outs  ExpectedOuts
	}{
		"it resets the access token memoized on the service": {
			Mocks: TestMocks{
				mockedResponseCode: http.StatusOK,
				mockedResponseBody: ClientCredential{
					AccessToken: oltypes.String("after"),
				},
			},
			Outs: ExpectedOuts{
				expectedServiceState: V2Service{
					Token: "after",
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			client := &http.Client{Timeout: time.Second * 5}
			srv := MockServer(&test.Mocks)
			defer srv.Close()
			cfg := &services.AuthServiceConfig{
				ClientID:     "",
				ClientSecret: "",
				BaseURL:      srv.URL,
				Client:       client,
			}

			auth := New(cfg)
			tok, _ := auth.ReAuthorize()
			fmt.Println("FUCK", auth.Token, tok)
			assert.Equal(t, test.Outs.expectedServiceState.Token, auth.Token)
		})
	}
}

func TestSetToken(t *testing.T) {
	tests := map[string]struct {
		Mocks TestMocks
		Outs  ExpectedOuts
	}{
		"it memoizes the access token from the api on the service": {
			Mocks: TestMocks{
				mockedResponseCode: http.StatusOK,
				mockedResponseBody: ClientCredential{
					AccessToken: oltypes.String("test"),
				},
			},
			Outs: ExpectedOuts{
				expectedServiceState: V2Service{
					Token: "test",
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			client := &http.Client{Timeout: time.Second * 5}
			srv := MockServer(&test.Mocks)
			defer srv.Close()
			cfg := &services.AuthServiceConfig{
				ClientID:     "",
				ClientSecret: "",
				BaseURL:      srv.URL,
				Client:       client,
			}

			auth := New(cfg)
			setToken(&auth)

			assert.Equal(t, test.Outs.expectedServiceState.Token, auth.Token)
		})
	}
}
