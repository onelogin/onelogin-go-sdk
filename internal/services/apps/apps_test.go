package apps

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
	queryRequestMock *AppsQuery
	requestPayload   *App
	AppID            int32
}

type TestMocks struct {
	authMockErr           error
	appMockPayload        interface{}
	appMockStatusCodeResp int
	tokenExpiredError     error
	tokenRefreshError     error
}

type ExpectedOuts struct {
	expectedAppRespDescription string
	expectedAppErrorMsg        string
	expectedAppStatusCode      int
	expectedAppCount           int
	expectedApps               []App
	expectedAppRespID          int
	expectedAppRespConnectorID int
}

func MockServer(mocks *TestMocks) *httptest.Server {
	// set up the mock server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if mocks.tokenExpiredError != nil {
			w.WriteHeader(http.StatusUnauthorized)
			mocks.tokenExpiredError = nil
		} else {
			w.WriteHeader(mocks.appMockStatusCodeResp)
			if mocks.appMockPayload != nil {
				body, _ := json.Marshal(mocks.appMockPayload)
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

func TestQuery(t *testing.T) {
	tests := map[string]struct {
		Ins   TestIns
		Mocks TestMocks
		Outs  ExpectedOuts
	}{
		"returns one app if limit 1 given": {
			Ins: TestIns{
				queryRequestMock: &AppsQuery{
					Limit: oltypes.String("1"),
				},
			},
			Mocks: TestMocks{
				appMockPayload: []App{
					App{
						ID:          oltypes.Int32(1234),
						ConnectorID: oltypes.Int32(1111),
					},
				},
				appMockStatusCodeResp: http.StatusOK,
			},
			Outs: ExpectedOuts{
				expectedAppStatusCode: http.StatusOK,
				expectedAppCount:      1,
			},
		},
		"returns all apps if empty query given": {
			Ins: TestIns{
				queryRequestMock: &AppsQuery{},
			},
			Mocks: TestMocks{
				appMockPayload: []App{
					App{
						ID:          oltypes.Int32(1234),
						ConnectorID: oltypes.Int32(1111),
					},
					App{
						ID:          oltypes.Int32(5678),
						ConnectorID: oltypes.Int32(2222),
					},
				},
				appMockStatusCodeResp: http.StatusOK,
			},
			Outs: ExpectedOuts{
				expectedAppStatusCode: http.StatusOK,
				expectedAppCount:      2,
			},
		},
		"auth returns 401 status code": {
			Mocks: TestMocks{
				authMockErr:           errors.New("test"),
				appMockStatusCodeResp: http.StatusUnauthorized,
			},
			Outs: ExpectedOuts{
				expectedAppErrorMsg: "error: context: [apps v2 service], error_message: [test]",
			},
		},
		"expired access token re-auths and succeeds": {
			Mocks: TestMocks{
				tokenExpiredError: errors.New("expired"),
				appMockPayload: []App{
					App{
						ID:          oltypes.Int32(1234),
						ConnectorID: oltypes.Int32(1111),
					},
				},
				appMockStatusCodeResp: http.StatusOK,
			},
			Outs: ExpectedOuts{
				expectedAppStatusCode: http.StatusOK,
				expectedAppCount:      1,
			},
		},
		"expired access token re-auths and fails": {
			Mocks: TestMocks{
				tokenExpiredError: errors.New("expired"),
				tokenRefreshError: errors.New("failed to renew access token"),
			},
			Outs: ExpectedOuts{
				expectedAppErrorMsg:   "error: context: [apps v2 service], error_message: [failed to renew access token]",
				expectedAppStatusCode: http.StatusUnauthorized,
				expectedAppCount:      1,
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			client := &http.Client{Timeout: time.Second * 5}
			mockedAuthenticator := MockAuthenticator(test.Mocks)
			srv := MockServer(&test.Mocks)
			defer srv.Close()
			apps := New(&services.APIServiceConfig{
				BaseURL: srv.URL,
				Client:  client,
				Auth:    mockedAuthenticator,
			})
			resultResp, resultPayload, resultErr := apps.Query(test.Ins.queryRequestMock)

			if test.Outs.expectedAppErrorMsg != "" {
				assert.NotNil(t, resultErr)
				assert.Equal(t, test.Outs.expectedAppErrorMsg, resultErr.Error())
			} else {
				assert.Nil(t, resultErr)
				actualPayload, _ := json.Marshal(resultPayload)
				expectedPayload, _ := json.Marshal(test.Mocks.appMockPayload)
				assert.Equal(t, expectedPayload, actualPayload)
				assert.Equal(t, test.Outs.expectedAppStatusCode, resultResp.StatusCode)
				assert.Equal(t, test.Outs.expectedAppCount, len(resultPayload))
				assert.Nil(t, resultErr)
			}
		})
	}
}

func TestGetOne(t *testing.T) {
	tests := map[string]struct {
		Ins   TestIns
		Mocks TestMocks
		Outs  ExpectedOuts
	}{
		"auth returns access token and successful response": {
			Ins: TestIns{
				AppID: 1234,
			},
			Mocks: TestMocks{
				appMockPayload: &App{
					ID:          oltypes.Int32(1234),
					ConnectorID: oltypes.Int32(1111),
				},
				appMockStatusCodeResp: http.StatusOK,
			},
			Outs: ExpectedOuts{
				expectedAppRespID:          1234,
				expectedAppRespConnectorID: 1111,
				expectedAppStatusCode:      http.StatusOK,
			},
		},
		"auth returns 401 status code": {
			Ins: TestIns{},
			Mocks: TestMocks{
				authMockErr:           errors.New("test"),
				appMockStatusCodeResp: http.StatusUnauthorized,
			},
			Outs: ExpectedOuts{
				expectedAppErrorMsg:   "error: context: [apps v2 service], error_message: [test]",
				expectedAppStatusCode: http.StatusCreated,
			},
		},
		"expired access token re-auths and succeeds": {
			Mocks: TestMocks{
				tokenExpiredError: errors.New("expired"),
				appMockPayload: App{
					ID:          oltypes.Int32(1234),
					ConnectorID: oltypes.Int32(1111),
				},
				appMockStatusCodeResp: http.StatusOK,
			},
			Outs: ExpectedOuts{
				expectedAppRespID:          1234,
				expectedAppRespConnectorID: 1111,
				expectedAppStatusCode:      http.StatusOK,
				expectedAppCount:           1,
			},
		},
		"expired access token re-auths and fails": {
			Mocks: TestMocks{
				tokenExpiredError:     errors.New("expired"),
				tokenRefreshError:     errors.New("failed to renew access token"),
				appMockStatusCodeResp: http.StatusUnauthorized,
			},
			Outs: ExpectedOuts{
				expectedAppErrorMsg:   "error: context: [apps v2 service], error_message: [failed to renew access token]",
				expectedAppStatusCode: http.StatusUnauthorized,
				expectedAppCount:      1,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			client := &http.Client{Timeout: time.Second * 5}
			mockedAuthenticator := MockAuthenticator(test.Mocks)
			srv := MockServer(&test.Mocks)
			defer srv.Close()
			apps := New(&services.APIServiceConfig{
				BaseURL: srv.URL,
				Client:  client,
				Auth:    mockedAuthenticator,
			})

			resultResp, resultPayload, resultErr := apps.GetOne(test.Ins.AppID)

			if test.Outs.expectedAppErrorMsg != "" {
				assert.NotNil(t, resultErr)
				assert.Equal(t, test.Outs.expectedAppErrorMsg, resultErr.Error())
			} else {
				actualPayload, _ := json.Marshal(resultPayload)
				expectedPayload, _ := json.Marshal(test.Mocks.appMockPayload)
				assert.Equal(t, expectedPayload, actualPayload)
				assert.Equal(t, test.Outs.expectedAppStatusCode, resultResp.StatusCode)
				assert.Nil(t, resultErr)
			}
		})
	}
}

func TestDestroy(t *testing.T) {
	tests := map[string]struct {
		Ins   TestIns
		Mocks TestMocks
		Outs  ExpectedOuts
	}{
		"auth returns access token and successful deletion": {
			Ins: TestIns{
				AppID: 1234,
			},
			Mocks: TestMocks{
				appMockStatusCodeResp: http.StatusOK,
			},
			Outs: ExpectedOuts{
				expectedAppStatusCode: http.StatusOK,
			},
		},
		"auth returns access token and unsuccessful deletion w/ status code 404": {
			Ins: TestIns{
				AppID: 1234,
			},
			Mocks: TestMocks{
				appMockStatusCodeResp: http.StatusNotFound,
			},
			Outs: ExpectedOuts{
				expectedAppErrorMsg: "request error: context: apps v2 service, status_code: [404], error_message: Not Found",
			},
		},
		"auth returns 401 status code": {
			Ins: TestIns{
				AppID: 1234,
			},
			Mocks: TestMocks{
				authMockErr:           errors.New("test"),
				appMockStatusCodeResp: http.StatusUnauthorized,
			},
			Outs: ExpectedOuts{
				expectedAppErrorMsg: "error: context: [apps v2 service], error_message: [test]",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			client := &http.Client{Timeout: time.Second * 5}
			mockedAuthenticator := MockAuthenticator(test.Mocks)
			srv := MockServer(&test.Mocks)
			defer srv.Close()
			apps := New(&services.APIServiceConfig{
				BaseURL: srv.URL,
				Client:  client,
				Auth:    mockedAuthenticator,
			})

			resultResp, resultErr := apps.Destroy(test.Ins.AppID)

			if test.Outs.expectedAppErrorMsg != "" {
				assert.NotNil(t, resultErr)
				assert.Equal(t, test.Outs.expectedAppErrorMsg, resultErr.Error())
			} else {
				assert.Nil(t, resultErr)
				assert.Equal(t, test.Outs.expectedAppStatusCode, resultResp.StatusCode)
			}
		})
	}
}

func TestCreateApp(t *testing.T) {
	tests := map[string]struct {
		Ins   TestIns
		Mocks TestMocks
		Outs  ExpectedOuts
	}{
		"auth returns access token and successful response": {
			Ins: TestIns{
				requestPayload: &App{
					ID:          oltypes.Int32(1234),
					ConnectorID: oltypes.Int32(1111),
				},
			},
			Mocks: TestMocks{
				appMockPayload: &App{
					ID:          oltypes.Int32(1234),
					ConnectorID: oltypes.Int32(1111),
				},
				appMockStatusCodeResp: http.StatusCreated,
			},
			Outs: ExpectedOuts{
				expectedAppRespID:          1234,
				expectedAppRespConnectorID: 1111,
				expectedAppStatusCode:      http.StatusCreated,
			},
		},
		"auth returns 401 status code": {
			Ins: TestIns{
				requestPayload: &App{
					ID:          oltypes.Int32(1234),
					ConnectorID: oltypes.Int32(1111),
				},
			},
			Mocks: TestMocks{
				authMockErr:           errors.New("test"),
				appMockStatusCodeResp: http.StatusUnauthorized,
			},
			Outs: ExpectedOuts{
				expectedAppErrorMsg: "error: context: [apps v2 service], error_message: [test]",
			},
		},
		"bad app api call returns a 400": {
			Ins: TestIns{
				requestPayload: &App{
					ID:          oltypes.Int32(1234),
					ConnectorID: oltypes.Int32(1111),
				},
			},
			Mocks: TestMocks{
				appMockStatusCodeResp: http.StatusBadRequest,
			},
			Outs: ExpectedOuts{
				expectedAppErrorMsg: "request error: context: apps v2 service, status_code: [400], error_message: Bad Request",
			},
		},
		"expired access token re-auths and succeeds": {
			Mocks: TestMocks{
				tokenExpiredError: errors.New("expired"),
				appMockPayload: App{
					ID:          oltypes.Int32(1234),
					ConnectorID: oltypes.Int32(1111),
				},
				appMockStatusCodeResp: http.StatusOK,
			},
			Outs: ExpectedOuts{
				expectedAppRespID:          1234,
				expectedAppRespConnectorID: 1111,
				expectedAppStatusCode:      http.StatusOK,
				expectedAppCount:           1,
			},
		},
		"expired access token re-auths and fails": {
			Mocks: TestMocks{
				tokenExpiredError:     errors.New("expired"),
				tokenRefreshError:     errors.New("failed to renew access token"),
				appMockStatusCodeResp: http.StatusUnauthorized,
			},
			Outs: ExpectedOuts{
				expectedAppErrorMsg:   "error: context: [apps v2 service], error_message: [failed to renew access token]",
				expectedAppStatusCode: http.StatusUnauthorized,
				expectedAppCount:      1,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			client := &http.Client{Timeout: time.Second * 5}
			mockedAuthenticator := MockAuthenticator(test.Mocks)
			srv := MockServer(&test.Mocks)
			defer srv.Close()
			apps := New(&services.APIServiceConfig{
				BaseURL: srv.URL,
				Client:  client,
				Auth:    mockedAuthenticator,
			})

			resultResp, resultPayload, resultErr := apps.Create(test.Ins.requestPayload)

			if test.Outs.expectedAppErrorMsg != "" {
				assert.NotNil(t, resultErr)
				assert.Equal(t, test.Outs.expectedAppErrorMsg, resultErr.Error())
			} else {
				if resultErr != nil {
					fmt.Println(resultErr)
				}
				assert.Nil(t, resultErr)
				actualPayload, _ := json.Marshal(resultPayload)
				expectedPayload, _ := json.Marshal(test.Mocks.appMockPayload)
				assert.Equal(t, expectedPayload, actualPayload)
				assert.Equal(t, test.Outs.expectedAppStatusCode, resultResp.StatusCode)
				assert.Nil(t, resultErr)
			}
		})
	}
}

func TestUpdateApp(t *testing.T) {
	tests := map[string]struct {
		Ins   TestIns
		Mocks TestMocks
		Outs  ExpectedOuts
	}{
		"auth returns access token and successful response": {
			Ins: TestIns{
				AppID: 1234,
				requestPayload: &App{
					ID:          oltypes.Int32(1234),
					ConnectorID: oltypes.Int32(1111),
					Description: oltypes.String("UPDATE"),
				},
			},
			Mocks: TestMocks{
				appMockPayload: &App{
					ID:          oltypes.Int32(1234),
					ConnectorID: oltypes.Int32(1111),
					Description: oltypes.String("UPDATE"),
				},
				appMockStatusCodeResp: http.StatusOK,
			},
			Outs: ExpectedOuts{
				expectedAppRespID:          1234,
				expectedAppRespConnectorID: 1111,
				expectedAppRespDescription: "UPDATE",
				expectedAppStatusCode:      http.StatusOK,
			},
		},
		"auth returns 401 status code": {
			Ins: TestIns{
				AppID: 1234,
				requestPayload: &App{
					ID:          oltypes.Int32(1234),
					ConnectorID: oltypes.Int32(1111),
				},
			},
			Mocks: TestMocks{
				authMockErr:           errors.New("test"),
				appMockStatusCodeResp: http.StatusBadRequest,
			},
			Outs: ExpectedOuts{
				expectedAppErrorMsg: "error: context: [apps v2 service], error_message: [test]",
			},
		},
		"app api call returns a 404": {
			Ins: TestIns{
				AppID: 1234,
				requestPayload: &App{
					ID:          oltypes.Int32(1234),
					ConnectorID: oltypes.Int32(1111),
				},
			},
			Mocks: TestMocks{
				appMockStatusCodeResp: http.StatusNotFound,
			},
			Outs: ExpectedOuts{
				expectedAppErrorMsg: "request error: context: apps v2 service, status_code: [404], error_message: Not Found",
			},
		},
		"expired access token re-auths and succeeds": {
			Mocks: TestMocks{
				tokenExpiredError: errors.New("expired"),
				appMockPayload: App{
					ID:          oltypes.Int32(1234),
					ConnectorID: oltypes.Int32(1111),
				},
				appMockStatusCodeResp: http.StatusOK,
			},
			Outs: ExpectedOuts{
				expectedAppRespID:          1234,
				expectedAppRespConnectorID: 1111,
				expectedAppStatusCode:      http.StatusOK,
				expectedAppCount:           1,
			},
		},
		"expired access token re-auths and fails": {
			Mocks: TestMocks{
				tokenExpiredError:     errors.New("expired"),
				tokenRefreshError:     errors.New("failed to renew access token"),
				appMockStatusCodeResp: http.StatusUnauthorized,
			},
			Outs: ExpectedOuts{
				expectedAppErrorMsg:   "error: context: [apps v2 service], error_message: [failed to renew access token]",
				expectedAppStatusCode: http.StatusUnauthorized,
				expectedAppCount:      1,
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			client := &http.Client{Timeout: time.Second * 5}
			mockedAuthenticator := MockAuthenticator(test.Mocks)
			srv := MockServer(&test.Mocks)
			defer srv.Close()
			apps := New(&services.APIServiceConfig{
				BaseURL: srv.URL,
				Client:  client,
				Auth:    mockedAuthenticator,
			})

			resultResp, resultPayload, resultErr := apps.Update(test.Ins.AppID, test.Ins.requestPayload)

			if test.Outs.expectedAppErrorMsg != "" {
				assert.NotNil(t, resultErr)
				assert.Equal(t, test.Outs.expectedAppErrorMsg, resultErr.Error())
			} else {
				assert.Nil(t, resultErr)
				actualPayload, _ := json.Marshal(resultPayload)
				expectedPayload, _ := json.Marshal(test.Mocks.appMockPayload)
				assert.Equal(t, expectedPayload, actualPayload)
				assert.Equal(t, test.Outs.expectedAppStatusCode, resultResp.StatusCode)
				assert.Nil(t, resultErr)
			}
		})
	}
}
