package sessionlogintokens

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/onelogin/onelogin-go-sdk/internal/services/client_credentials"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TypeAuthV2Mock struct {
	mock.Mock
}

// mocked authorize function which returns mocked response
func (auth *TypeAuthV2Mock) Authorize() (*http.Response, *clientcredentials.ClientCredential, error) {
	args := auth.Called()
	return args.Get(0).(*http.Response), args.Get(1).(*clientcredentials.ClientCredential), args.Error(2)
}

func TestCreateSessionLoginToken(t *testing.T) {
	tests := map[string]struct {
		authMockAuthToken    string
		authMockErr          error
		authMockStatusCode   int
		mockRequest          *SessionLoginTokenRequest
		mockResponse         *SessionLoginToken
		mockStatusCodeResp   int
		isErrExpected        bool
		expectedErrorMsg     string
		expectedSessionToken string
		expectedStateToken   string
		expectedStatusCode   int
	}{
		"auth returns access token and successful response": {
			authMockAuthToken:  "testAuthToken",
			authMockErr:        nil,
			authMockStatusCode: http.StatusOK,
			mockRequest: &SessionLoginTokenRequest{
				UsernameOrEmail: oltypes.String("test"),
				Password:        oltypes.String("test"),
				Subdomain:       oltypes.String("test"),
			},
			mockResponse: &SessionLoginToken{
				SessionToken: oltypes.String("test"),
				ReturnToURL:  oltypes.String("test"),
				StateToken:   oltypes.String("test"),
			},
			mockStatusCodeResp:   http.StatusCreated,
			isErrExpected:        false,
			expectedErrorMsg:     "",
			expectedSessionToken: "test",
			expectedStateToken:   "test",
			expectedStatusCode:   http.StatusCreated,
		},
		"auth returns 401 status code": {
			authMockAuthToken:  "testAuthToken",
			authMockErr:        errors.New("test"),
			authMockStatusCode: http.StatusUnauthorized,
			mockRequest: &SessionLoginTokenRequest{
				UsernameOrEmail: oltypes.String("bad_guy"),
				Password:        oltypes.String("bad_guy"),
				Subdomain:       oltypes.String("bad_guy"),
			},
			mockStatusCodeResp: http.StatusBadRequest,
			isErrExpected:      true,
			expectedErrorMsg:   "request error: context: apps v2 service, status_code: [401], error_message: test",
		},
		"app api call returns a 400": {
			authMockAuthToken:  "testAuthToken",
			authMockErr:        nil,
			authMockStatusCode: http.StatusOK,
			mockRequest: &SessionLoginTokenRequest{
				UsernameOrEmail: oltypes.String("test"),
				Subdomain:       oltypes.String("test"),
			},
			mockStatusCodeResp: http.StatusBadRequest,
			isErrExpected:      true,
			expectedErrorMsg:   "request error: context: apps v2 service, status_code: [400], error_message: Bad Request",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// set up mock auth response
			authOutput, _ := json.Marshal(&clientcredentials.ClientCredential{
				AccessToken: oltypes.String(test.authMockAuthToken),
			})

			clientCredential := &http.Response{
				StatusCode: test.authMockStatusCode,
				Body:       ioutil.NopCloser(bytes.NewReader(authOutput)),
			}

			authErr := test.authMockErr

			payloadRes := &clientcredentials.ClientCredential{
				AccessToken: oltypes.String(test.authMockAuthToken),
			}

			// set up the mock server
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.mockStatusCodeResp)
				if test.mockResponse != nil {
					body, _ := json.Marshal(test.mockResponse)
					w.Write(body)
				}
			}))

			defer ts.Close()

			client := &http.Client{
				Timeout: time.Second * 5,
			}

			mockedObj := &TypeAuthV2Mock{}

			mockedObj.On("Authorize", mock.Anything).Return(clientCredential, payloadRes, authErr)

			sessionLoginToken := NewSessionLoginTokenV1(&SessionLoginTokenV1Config{
				BaseURL: ts.URL,
				Client:  client,
				Auth:    mockedObj,
			})

			resultResp, resultPayload, resultErr := sessionLoginToken.CreateSessionLoginToken(test.mockRequest)

			if test.isErrExpected {
				assert.NotNil(t, resultErr)
				assert.Equal(t, test.expectedErrorMsg, resultErr.Error())
			} else {
				assert.Nil(t, resultErr)
				assert.Equal(t, test.expectedStatusCode, resultResp.StatusCode)

				resultSessionToken, _ := oltypes.GetStringVal(resultPayload.SessionToken)
				resultStateToken, _ := oltypes.GetStringVal(resultPayload.StateToken)

				assert.Equal(t, test.expectedSessionToken, resultSessionToken)
				assert.Equal(t, test.expectedStateToken, resultStateToken)
				assert.Nil(t, resultErr)
			}

			// make sure the mocked function was called
			mockedObj.AssertExpectations(t)
		})
	}
}
