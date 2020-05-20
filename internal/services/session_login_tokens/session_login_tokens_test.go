package sessionlogintokens

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/onelogin/onelogin-go-sdk/internal/services"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AuthenticatorMock struct {
	mock.Mock
}

func (auth *AuthenticatorMock) Authorize() (string, error) {
	args := auth.Called()
	return args.Get(0).(string), args.Error(1)
}

func (auth *AuthenticatorMock) ReAuthorize() (string, error) {
	args := auth.Called()
	return args.Get(0).(string), args.Error(1)
}

type TestIns struct {
	mockRequest *SessionLoginTokenRequest
}

type TestMocks struct {
	authMockErr           error
	authMockStatusCode    int
	appMockPayload        interface{}
	appMockStatusCodeResp int
	tokenExpiredError     error
	tokenRefreshError     error
	mockResponse          *SessionLoginToken
	mockStatusCodeResp    int
}

type ExpectedOuts struct {
	expectedErrorMsg     string
	expectedSessionToken string
	expectedStateToken   string
	expectedStatusCode   int
}

func MockServer(mocks *TestMocks) *httptest.Server {
	// set up the mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if mocks.tokenExpiredError != nil {
			w.WriteHeader(http.StatusUnauthorized)
			mocks.tokenExpiredError = nil
		} else {
			w.WriteHeader(mocks.mockStatusCodeResp)
			if mocks.mockResponse != nil {
				body, _ := json.Marshal(mocks.mockResponse)
				w.Write(body)
			}
		}
	}))
	return ts
}

func MockAuthenticator(mocks TestMocks) *AuthenticatorMock {
	mockedObj := &AuthenticatorMock{}
	mockedObj.On("Authorize", mock.Anything).Return("accessToken", mocks.authMockErr)
	if mocks.tokenExpiredError != nil {
		mockedObj.On("ReAuthorize", mock.Anything).Return("accessToken", mocks.tokenRefreshError)
	}
	return mockedObj
}

func TestCreate(t *testing.T) {
	tests := map[string]struct {
		Ins   TestIns
		Mocks TestMocks
		Outs  ExpectedOuts
	}{
		"If request is valid and user credentials are valid, it returns a session token": {
			Ins: TestIns{
				mockRequest: &SessionLoginTokenRequest{
					UsernameOrEmail: oltypes.String("test"),
					Password:        oltypes.String("test"),
					Subdomain:       oltypes.String("test"),
				},
			},
			Mocks: TestMocks{
				authMockStatusCode: http.StatusOK,
				mockStatusCodeResp: http.StatusCreated,
				mockResponse: &SessionLoginToken{
					ReturnToURL:  oltypes.String("http://test.com"),
					SessionToken: oltypes.String("test"),
					StateToken:   oltypes.String("test"),
				},
			},
			Outs: ExpectedOuts{
				expectedSessionToken: "test",
				expectedStateToken:   "test",
				expectedStatusCode:   http.StatusCreated,
			},
		},
		"An unauthorized request is given a 401 response": {
			Ins: TestIns{
				mockRequest: &SessionLoginTokenRequest{
					UsernameOrEmail: oltypes.String("bad_guy"),
					Password:        oltypes.String("bad_guy"),
					Subdomain:       oltypes.String("bad_guy"),
				},
			},
			Mocks: TestMocks{
				authMockErr:        errors.New("test"),
				mockStatusCodeResp: http.StatusBadRequest,
				mockResponse: &SessionLoginToken{
					ReturnToURL:  oltypes.String("http://test.com"),
					SessionToken: oltypes.String("test"),
					StateToken:   oltypes.String("test"),
				},
			},
			Outs: ExpectedOuts{
				expectedErrorMsg: "error: context: [Session Login Tokens v2 service], error_message: [test]",
			},
		},
		"Invalid user credentials returns a 401 response": {
			Ins: TestIns{
				mockRequest: &SessionLoginTokenRequest{
					UsernameOrEmail: oltypes.String("bad"),
					Subdomain:       oltypes.String("bad"),
				},
			},
			Mocks: TestMocks{
				authMockStatusCode: http.StatusOK,
				mockStatusCodeResp: http.StatusBadRequest,
			},
			Outs: ExpectedOuts{
				expectedErrorMsg: "request error: context: Session Login Tokens v2 service, status_code: [400], error_message: Bad Request",
			},
		},
		"expired access token re-auths and succeeds": {
			Mocks: TestMocks{
				tokenExpiredError: errors.New("expired"),
				mockResponse: &SessionLoginToken{
					ReturnToURL:  oltypes.String("http://test.com"),
					SessionToken: oltypes.String("test"),
					StateToken:   oltypes.String("test"),
				},
				mockStatusCodeResp: http.StatusOK,
			},
			Outs: ExpectedOuts{
				expectedSessionToken: "test",
				expectedStateToken:   "test",
				expectedStatusCode:   http.StatusOK,
			},
		},
		"expired access token re-auths and fails": {
			Mocks: TestMocks{
				tokenExpiredError:  errors.New("expired"),
				tokenRefreshError:  errors.New("failed to renew access token"),
				mockStatusCodeResp: http.StatusUnauthorized,
			},
			Outs: ExpectedOuts{
				expectedErrorMsg:   "error: context: [Session Login Tokens v2 service], error_message: [failed to renew access token]",
				expectedStatusCode: http.StatusUnauthorized,
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			client := &http.Client{Timeout: time.Second * 5}
			mockedAuthenticator := MockAuthenticator(test.Mocks)
			srv := MockServer(&test.Mocks)
			defer srv.Close()
			sessionLoginToken := New(&services.APIServiceConfig{
				BaseURL: srv.URL,
				Client:  client,
				Auth:    mockedAuthenticator,
			})

			resultResp, resultPayload, resultErr := sessionLoginToken.Create(test.Ins.mockRequest)
			if test.Outs.expectedErrorMsg != "" {
				assert.NotNil(t, resultErr)
				assert.Equal(t, test.Outs.expectedErrorMsg, resultErr.Error())
			} else {
				assert.Nil(t, resultErr)
				assert.Equal(t, test.Outs.expectedStatusCode, resultResp.StatusCode)
				fmt.Println("RES", resultPayload, resultErr)
				resultSessionToken, _ := oltypes.GetStringVal(resultPayload.SessionToken)
				resultStateToken, _ := oltypes.GetStringVal(resultPayload.StateToken)
				assert.Equal(t, test.Outs.expectedSessionToken, resultSessionToken)
				assert.Equal(t, test.Outs.expectedStateToken, resultStateToken)
			}
		})
	}
}
