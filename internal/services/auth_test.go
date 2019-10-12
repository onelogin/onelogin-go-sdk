package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/onelogin/onelogin-go-sdk/pkg/models"

	"github.com/stretchr/testify/assert"
)

func TestAuthorize(t *testing.T) {
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	expectedAccessTokenFor1 := "test"
	expectedStatusCodeFor1 := http.StatusOK
	mockedResponse1, _ := json.Marshal(models.AuthResp{
		AccessToken: expectedAccessTokenFor1,
	})

	expectedStatusCodeFor2 := http.StatusUnauthorized

	tests := map[string]struct {
		mockedResponseBody      []byte
		mockedResponseCode      int
		httpClient              *http.Client
		expectedErroMsg         string
		expectedStatusCode      int
		expectedAccessToken     string
		isErrResExpected        bool
		isHTTPResponseExpected  bool
		isRespPayloadCompNeeded bool
	}{
		"access token is returned": {
			httpClient:              httpClient,
			mockedResponseBody:      mockedResponse1,
			mockedResponseCode:      expectedStatusCodeFor1,
			expectedStatusCode:      expectedStatusCodeFor1,
			expectedAccessToken:     expectedAccessTokenFor1,
			isRespPayloadCompNeeded: true,
			isErrResExpected:        false,
			isHTTPResponseExpected:  true,
		},
		"error is returned when incorrect status code": {
			httpClient:              httpClient,
			mockedResponseCode:      expectedStatusCodeFor2,
			expectedStatusCode:      expectedStatusCodeFor2,
			isErrResExpected:        true,
			isHTTPResponseExpected:  true,
			isRespPayloadCompNeeded: false,
			expectedErroMsg:         "request error: context: auth v2 service, status_code: [401], error_message: Unauthorized",
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			// set up the mock server
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.mockedResponseCode)
				if test.mockedResponseBody != nil {
					w.Write(test.mockedResponseBody)
				}
			}))

			defer ts.Close()

			// end set up the mock server

			cfg := &AuthConfigV2{
				Client:       test.httpClient,
				ClientID:     "",
				ClientSecret: "",
				BaseURL:      ts.URL,
			}

			auth := NewAuthV2(cfg)

			resp, body, err := auth.Authorize()

			// check errors
			if test.isErrResExpected {
				assert.NotNil(t, err)
				assert.Equal(t, err.Error(), test.expectedErroMsg)
			} else {
				assert.Nil(t, err)
			}

			// check http response
			if test.isHTTPResponseExpected {
				assert.NotNil(t, resp)

				if test.isRespPayloadCompNeeded {
					assert.Equal(t, test.expectedAccessToken, body.AccessToken)
				}

				// check the response status code
				assert.Equal(t, test.expectedStatusCode, resp.StatusCode)
			} else {
				assert.Nil(t, resp)
			}
		})
	}
}
