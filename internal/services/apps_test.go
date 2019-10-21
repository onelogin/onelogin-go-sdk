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

	"github.com/stretchr/testify/assert"

	"github.com/onelogin/onelogin-go-sdk/pkg/models"
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
		authMockAuthToken      string
		authMockErr            error
		authMockStatusCode     int
		appMockPayload         *models.App
		appMockStatusCodeResp  int
		appMockResErr          error
		isAppErrExpected       bool
		expectedAppErrorMsg    string
		expectedAppRespPayload *models.App
		expectedAppStatusCode  int
	}{
		"auth returns access token": {
			authMockAuthToken:  "testAuthToken",
			authMockErr:        nil,
			authMockStatusCode: http.StatusOK,
			appMockPayload: &models.App{
				ID:          1234,
				ConnectorID: 1111,
			},
			appMockStatusCodeResp: http.StatusCreated,
			appMockResErr:         nil,
			isAppErrExpected:      false,
			expectedAppErrorMsg:   "",
			expectedAppRespPayload: &models.App{
				ID:          1234,
				ConnectorID: 1111,
			},
			expectedAppStatusCode: http.StatusCreated,
		},
		"auth returns http.StatusUnauthorized status code": {
			authMockAuthToken:  "testAuthToken",
			authMockErr:        errors.New("test"),
			authMockStatusCode: http.StatusUnauthorized,
			appMockPayload: &models.App{
				ID:          1234,
				ConnectorID: 1111,
			},
			appMockStatusCodeResp:  http.StatusBadRequest,
			appMockResErr:          nil,
			isAppErrExpected:       true,
			expectedAppErrorMsg:    "request error: context: apps v2 service, status_code: [401], error_message: test",
			expectedAppRespPayload: nil,
			expectedAppStatusCode:  http.StatusCreated,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			authOutput, _ := json.Marshal(&models.AuthResp{
				AccessToken: test.authMockAuthToken,
			})

			authResp := &http.Response{
				StatusCode: test.authMockStatusCode,
				Body:       ioutil.NopCloser(bytes.NewReader(authOutput)),
			}

			authErr := test.authMockErr

			payloadRes := &models.AuthResp{
				AccessToken: test.authMockAuthToken,
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

			resultResp, resultPayload, resultErr := apps.GetAppByID(test.appMockPayload.ID)

			if test.isAppErrExpected {
				assert.NotNil(t, resultErr)
				assert.Equal(t, test.expectedAppErrorMsg, resultErr.Error())
			} else {
				assert.Nil(t, resultErr)
				assert.Equal(t, test.expectedAppStatusCode, resultResp.StatusCode)
				assert.Equal(t, test.expectedAppRespPayload.ID, resultPayload.ID)
				assert.Equal(t, test.expectedAppRespPayload.ConnectorID, resultPayload.ConnectorID)
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
		"auth returns http.StatusUnauthorized status code": {
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
				AccessToken: test.authMockAuthToken,
			})

			authResp := &http.Response{
				StatusCode: test.authMockStatusCode,
				Body:       ioutil.NopCloser(bytes.NewReader(authOutput)),
			}

			authErr := test.authMockErr

			payloadRes := &models.AuthResp{
				AccessToken: test.authMockAuthToken,
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
