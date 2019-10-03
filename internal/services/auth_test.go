package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/onelogin/onelogin-go-sdk/pkg/models"

	"github.com/stretchr/testify/assert"
)

func TestAuthenticateWhenAccessTokenIsReturned(t *testing.T) {
	expectedAccessToken := "test"
	expectedStatusCode := http.StatusOK

	byteObj, _ := json.Marshal(models.AuthResp{
		AccessToken: expectedAccessToken,
	})

	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(expectedStatusCode)
		w.Write(byteObj)
	}))

	defer ts.Close()

	cfg := &AuthConfigV2{
		Client:       httpClient,
		ClientId:     "",
		ClientSecret: "",
		BaseUrl:      ts.URL,
	}

	auth := NewAuthV2(cfg)

	resp, body, err := auth.Authenticate()

	assert.NotNil(t, body)
	assert.Nil(t, err)
	assert.Equal(t, expectedStatusCode, resp.StatusCode)
	assert.Equal(t, expectedAccessToken, body.AccessToken)
}

func TestAuthenticateWhenGreaterThan200IsReturned(t *testing.T) {
	expectedStatusCode := http.StatusUnauthorized
	expectedErr := errors.New(fmt.Sprintf("Auth responded with status code of [ %d ]", http.StatusUnauthorized))

	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(expectedStatusCode)
	}))

	defer ts.Close()

	cfg := &AuthConfigV2{
		Client:       httpClient,
		ClientId:     "",
		ClientSecret: "",
		BaseUrl:      ts.URL,
	}

	auth := NewAuthV2(cfg)

	resp, body, err := auth.Authenticate()

	assert.Nil(t, body)
	assert.NotNil(t, err)
	assert.Equal(t, expectedErr.Error(), err.Error())
	assert.Equal(t, expectedStatusCode, resp.StatusCode)
}
