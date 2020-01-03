package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/onelogin/onelogin-go-sdk/pkg/models"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TypeAuthV2Mock struct {
	mock.Mock
}

// mocked authorize function which returns mocked response
func (auth *TypeAuthV2Mock) Authorize() (*http.Response, *models.AuthResp, error) {
	args := auth.Called()
	return args.Get(0).(*http.Response), args.Get(1).(*models.AuthResp), args.Error(2)
}

func TestGetAppByID(t *testing.T) {
	tests := map[string]struct {
		authMockAuthToken          string
		authMockErr                error
		authMockStatusCode         int
		appMockPayload             *models.App
		appMockStatusCodeResp      int
		isAppErrExpected           bool
		expectedAppErrorMsg        string
		expectedAppRespID          int32
		expectedAppRespConnectorID int32
		expectedAppStatusCode      int
	}{
		"auth returns access token and successful response": {
			authMockAuthToken:  "testAuthToken",
			authMockErr:        nil,
			authMockStatusCode: http.StatusOK,
			appMockPayload: &models.App{
				ID:          oltypes.Int32(1234),
				ConnectorID: oltypes.Int32(1111),
			},
			appMockStatusCodeResp:      http.StatusOK,
			isAppErrExpected:           false,
			expectedAppErrorMsg:        "",
			expectedAppRespID:          1234,
			expectedAppRespConnectorID: 1111,
			expectedAppStatusCode:      http.StatusOK,
		},
		"auth returns 401 status code": {
			authMockAuthToken:  "testAuthToken",
			authMockErr:        errors.New("test"),
			authMockStatusCode: http.StatusUnauthorized,
			appMockPayload: &models.App{
				ID:          oltypes.Int32(1234),
				ConnectorID: oltypes.Int32(1111),
			},
			appMockStatusCodeResp: http.StatusBadRequest,
			isAppErrExpected:      true,
			expectedAppErrorMsg:   "request error: context: apps v2 service, status_code: [401], error_message: test",
			expectedAppStatusCode: http.StatusCreated,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			authOutput, _ := json.Marshal(&models.AuthResp{
				AccessToken: oltypes.String(test.authMockAuthToken),
			})

			authResp := &http.Response{
				StatusCode: test.authMockStatusCode,
				Body:       ioutil.NopCloser(bytes.NewReader(authOutput)),
			}

			authErr := test.authMockErr

			payloadRes := &models.AuthResp{
				AccessToken: oltypes.String(test.authMockAuthToken),
			}

			// set up the mock server
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.appMockStatusCodeResp)
				if test.appMockPayload != nil {
					body, _ := json.Marshal(test.appMockPayload)
					w.Write(body)
				}
			}))

			defer ts.Close()

			client := &http.Client{
				Timeout: time.Second * 5,
			}

			mockedObj := &TypeAuthV2Mock{}

			mockedObj.On("Authorize", mock.Anything).Return(authResp, payloadRes, authErr)

			apps := NewAppsV2(&AppsV2Config{
				BaseURL: ts.URL,
				Client:  client,
				Auth:    mockedObj,
			})

			idToSearch, _ := oltypes.GetInt32Val(test.appMockPayload.ID)

			resultResp, resultPayload, resultErr := apps.GetAppByID(idToSearch)

			if test.isAppErrExpected {
				assert.NotNil(t, resultErr)
				assert.Equal(t, test.expectedAppErrorMsg, resultErr.Error())
			} else {
				assert.Nil(t, resultErr)

				resultPayloadID, _ := oltypes.GetInt32Val(resultPayload.ID)
				resultPayloadConnectorID, _ := oltypes.GetInt32Val(resultPayload.ConnectorID)

				assert.Equal(t, test.expectedAppStatusCode, resultResp.StatusCode)
				assert.Equal(t, test.expectedAppRespID, resultPayloadID)
				assert.Equal(t, test.expectedAppRespConnectorID, resultPayloadConnectorID)
				assert.Nil(t, resultErr)
			}

			// make sure the mocked function was called
			mockedObj.AssertExpectations(t)
		})
	}
}

func TestDeleteAppByID(t *testing.T) {
	tests := map[string]struct {
		authMockAuthToken     string
		authMockErr           error
		authMockStatusCode    int
		appIDtoDelete         int32
		appMockStatusCodeResp int
		isAppErrExpected      bool
		expectedAppErrorMsg   string
		expectedAppStatusCode int
	}{
		"auth returns access token and successful deletion": {
			authMockAuthToken:     "testAuthToken",
			authMockErr:           nil,
			authMockStatusCode:    http.StatusOK,
			appIDtoDelete:         1234,
			appMockStatusCodeResp: http.StatusOK,
			isAppErrExpected:      false,
			expectedAppErrorMsg:   "",
			expectedAppStatusCode: http.StatusOK,
		},
		"auth returns access token and unsuccessful deletion w/ status code 404": {
			authMockAuthToken:     "testAuthToken",
			authMockErr:           nil,
			authMockStatusCode:    http.StatusOK,
			appIDtoDelete:         1234,
			appMockStatusCodeResp: http.StatusNotFound,
			isAppErrExpected:      true,
			expectedAppErrorMsg:   "request error: context: apps v2 service, status_code: [404], error_message: Not Found",
		},
		"auth returns 401 status code": {
			authMockAuthToken:     "testAuthToken",
			authMockErr:           errors.New("test"),
			authMockStatusCode:    http.StatusUnauthorized,
			appIDtoDelete:         1234,
			appMockStatusCodeResp: http.StatusNotFound,
			isAppErrExpected:      true,
			expectedAppErrorMsg:   "request error: context: apps v2 service, status_code: [401], error_message: test",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// set up mock auth response
			authOutput, _ := json.Marshal(&models.AuthResp{
				AccessToken: oltypes.String(test.authMockAuthToken),
			})

			authResp := &http.Response{
				StatusCode: test.authMockStatusCode,
				Body:       ioutil.NopCloser(bytes.NewReader(authOutput)),
			}

			authErr := test.authMockErr

			payloadRes := &models.AuthResp{
				AccessToken: oltypes.String(test.authMockAuthToken),
			}

			// set up the mock server
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.appMockStatusCodeResp)
			}))

			defer ts.Close()

			client := &http.Client{
				Timeout: time.Second * 5,
			}

			mockedObj := &TypeAuthV2Mock{}

			mockedObj.On("Authorize", mock.Anything).Return(authResp, payloadRes, authErr)

			apps := NewAppsV2(&AppsV2Config{
				BaseURL: ts.URL,
				Client:  client,
				Auth:    mockedObj,
			})

			resultResp, resultErr := apps.DeleteApp(test.appIDtoDelete)

			if test.isAppErrExpected {
				assert.NotNil(t, resultErr)
				assert.Equal(t, test.expectedAppErrorMsg, resultErr.Error())
			} else {
				assert.Nil(t, resultErr)
				assert.Equal(t, test.expectedAppStatusCode, resultResp.StatusCode)
				assert.Nil(t, resultErr)
			}

			// make sure the mocked function was called
			mockedObj.AssertExpectations(t)
		})
	}
}

func TestCreateApp(t *testing.T) {
	tests := map[string]struct {
		authMockAuthToken          string
		authMockErr                error
		authMockStatusCode         int
		appMockPayload             *models.App
		appMockStatusCodeResp      int
		isAppErrExpected           bool
		expectedAppErrorMsg        string
		expectedAppRespID          int32
		expectedAppRespConnectorID int32
		expectedAppStatusCode      int
	}{
		"auth returns access token and successful response": {
			authMockAuthToken:  "testAuthToken",
			authMockErr:        nil,
			authMockStatusCode: http.StatusOK,
			appMockPayload: &models.App{
				ID:          oltypes.Int32(1234),
				ConnectorID: oltypes.Int32(1111),
			},
			appMockStatusCodeResp:      http.StatusCreated,
			isAppErrExpected:           false,
			expectedAppErrorMsg:        "",
			expectedAppRespID:          1234,
			expectedAppRespConnectorID: 1111,
			expectedAppStatusCode:      http.StatusCreated,
		},
		"auth returns 401 status code": {
			authMockAuthToken:  "testAuthToken",
			authMockErr:        errors.New("test"),
			authMockStatusCode: http.StatusUnauthorized,
			appMockPayload: &models.App{
				ID:          oltypes.Int32(1234),
				ConnectorID: oltypes.Int32(1111),
			},
			appMockStatusCodeResp: http.StatusBadRequest,
			isAppErrExpected:      true,
			expectedAppErrorMsg:   "request error: context: apps v2 service, status_code: [401], error_message: test",
		},
		"app api call returns a 400": {
			authMockAuthToken:  "testAuthToken",
			authMockErr:        nil,
			authMockStatusCode: http.StatusOK,
			appMockPayload: &models.App{
				ID:          oltypes.Int32(1234),
				ConnectorID: oltypes.Int32(1111),
			},
			appMockStatusCodeResp: http.StatusBadRequest,
			isAppErrExpected:      true,
			expectedAppErrorMsg:   "request error: context: apps v2 service, status_code: [400], error_message: Bad Request",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// set up mock auth response
			authOutput, _ := json.Marshal(&models.AuthResp{
				AccessToken: oltypes.String(test.authMockAuthToken),
			})

			authResp := &http.Response{
				StatusCode: test.authMockStatusCode,
				Body:       ioutil.NopCloser(bytes.NewReader(authOutput)),
			}

			authErr := test.authMockErr

			payloadRes := &models.AuthResp{
				AccessToken: oltypes.String(test.authMockAuthToken),
			}

			// set up the mock server
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.appMockStatusCodeResp)
				if test.appMockPayload != nil {
					body, _ := json.Marshal(test.appMockPayload)
					w.Write(body)
				}
			}))

			defer ts.Close()

			client := &http.Client{
				Timeout: time.Second * 5,
			}

			mockedObj := &TypeAuthV2Mock{}

			mockedObj.On("Authorize", mock.Anything).Return(authResp, payloadRes, authErr)

			apps := NewAppsV2(&AppsV2Config{
				BaseURL: ts.URL,
				Client:  client,
				Auth:    mockedObj,
			})

			resultResp, resultPayload, resultErr := apps.CreateApp(test.appMockPayload)

			if test.isAppErrExpected {
				assert.NotNil(t, resultErr)
				assert.Equal(t, test.expectedAppErrorMsg, resultErr.Error())
			} else {
				assert.Nil(t, resultErr)
				assert.Equal(t, test.expectedAppStatusCode, resultResp.StatusCode)

				resultPayloadID, _ := oltypes.GetInt32Val(resultPayload.ID)
				resultPayloadConnectorID, _ := oltypes.GetInt32Val(resultPayload.ConnectorID)

				assert.Equal(t, test.expectedAppRespID, resultPayloadID)
				assert.Equal(t, test.expectedAppRespConnectorID, resultPayloadConnectorID)
				assert.Nil(t, resultErr)
			}

			// make sure the mocked function was called
			mockedObj.AssertExpectations(t)
		})
	}
}

func TestUpdateApp(t *testing.T) {
	tests := map[string]struct {
		authMockAuthToken          string
		authMockErr                error
		authMockStatusCode         int
		appMockPayload             *models.App
		appMockStatusCodeResp      int
		isAppErrExpected           bool
		expectedAppErrorMsg        string
		expectedAppRespID          int32
		expectedAppRespConnectorID int32
		expectedAppRespDescription string
		expectedAppStatusCode      int
	}{
		"auth returns access token and successful response": {
			authMockAuthToken:  "testAuthToken",
			authMockErr:        nil,
			authMockStatusCode: http.StatusOK,
			appMockPayload: &models.App{
				ID:          oltypes.Int32(1234),
				ConnectorID: oltypes.Int32(1111),
				Description: oltypes.String("UPDATE"),
			},
			appMockStatusCodeResp:      http.StatusCreated,
			isAppErrExpected:           false,
			expectedAppErrorMsg:        "",
			expectedAppRespID:          1234,
			expectedAppRespConnectorID: 1111,
			expectedAppRespDescription: "UPDATE",
			expectedAppStatusCode:      http.StatusCreated,
		},
		"auth returns 401 status code": {
			authMockAuthToken:  "testAuthToken",
			authMockErr:        errors.New("test"),
			authMockStatusCode: http.StatusUnauthorized,
			appMockPayload: &models.App{
				ID:          oltypes.Int32(1234),
				ConnectorID: oltypes.Int32(1111),
			},
			appMockStatusCodeResp: http.StatusBadRequest,
			isAppErrExpected:      true,
			expectedAppErrorMsg:   "request error: context: apps v2 service, status_code: [401], error_message: test",
		},
		"app api call returns a 404": {
			authMockAuthToken:  "testAuthToken",
			authMockErr:        nil,
			authMockStatusCode: http.StatusOK,
			appMockPayload: &models.App{
				ID:          oltypes.Int32(1234),
				ConnectorID: oltypes.Int32(1111),
			},
			appMockStatusCodeResp: http.StatusBadRequest,
			isAppErrExpected:      true,
			expectedAppErrorMsg:   "request error: context: apps v2 service, status_code: [400], error_message: Bad Request",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// set up mock auth response
			authOutput, _ := json.Marshal(&models.AuthResp{
				AccessToken: oltypes.String(test.authMockAuthToken),
			})

			authResp := &http.Response{
				StatusCode: test.authMockStatusCode,
				Body:       ioutil.NopCloser(bytes.NewReader(authOutput)),
			}

			authErr := test.authMockErr

			payloadRes := &models.AuthResp{
				AccessToken: oltypes.String(test.authMockAuthToken),
			}

			// set up the mock server
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(test.appMockStatusCodeResp)
				if test.appMockPayload != nil {
					body, _ := json.Marshal(test.appMockPayload)
					w.Write(body)
				}
			}))

			defer ts.Close()

			client := &http.Client{
				Timeout: time.Second * 5,
			}

			mockedObj := &TypeAuthV2Mock{}

			mockedObj.On("Authorize", mock.Anything).Return(authResp, payloadRes, authErr)

			apps := NewAppsV2(&AppsV2Config{
				BaseURL: ts.URL,
				Client:  client,
				Auth:    mockedObj,
			})

			idToUpdate, _ := oltypes.GetInt32Val(test.appMockPayload.ID)

			resultResp, resultPayload, resultErr := apps.UpdateAppByID(idToUpdate, test.appMockPayload)

			if test.isAppErrExpected {
				assert.NotNil(t, resultErr)
				assert.Equal(t, test.expectedAppErrorMsg, resultErr.Error())
			} else {
				assert.Nil(t, resultErr)

				respStatusCode := resultResp.StatusCode
				respID, _ := oltypes.GetInt32Val(resultPayload.ID)
				respConnectorID, _ := oltypes.GetInt32Val(resultPayload.ConnectorID)
				respDescription, _ := oltypes.GetStringVal(resultPayload.Description)

				assert.Equal(t, test.expectedAppStatusCode, respStatusCode)
				assert.Equal(t, test.expectedAppRespID, respID)
				assert.Equal(t, test.expectedAppRespConnectorID, respConnectorID)
				assert.Equal(t, test.expectedAppRespDescription, respDescription)
				assert.Nil(t, resultErr)
			}

			// make sure the mocked function was called
			mockedObj.AssertExpectations(t)
		})
	}
}
