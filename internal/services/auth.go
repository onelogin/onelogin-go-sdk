package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/onelogin/onelogin-go-sdk/pkg/models"
)

const (
	CLIENT_CREDENTIALS_TEXT = "client_credentials"
)

type Authenticator interface {
	Authenticate() (*http.Response, *models.AuthResp, error)
}

type AuthV2 struct {
	ClientId     string
	ClientSecret string
	client       *http.Client
	baseUrl      string
}

type AuthConfigV2 struct {
	ClientId     string
	ClientSecret string
	Client       *http.Client
	BaseUrl      string
}

// NewAuthV2 uses the cfg to generate the new auth service, and returns
// the created auth service for version 2.
func NewAuthV2(cfg *AuthConfigV2) *AuthV2 {
	return &AuthV2{
		ClientId:     cfg.ClientId,
		ClientSecret: cfg.ClientSecret,
		client:       cfg.Client,
		baseUrl:      cfg.BaseUrl,
	}
}

// Authenticate authenticates the credentials for the ClientId and ClientSecret, and returns, if successfull,
// the http response and the auth response.
func (auth *AuthV2) Authenticate() (*http.Response, *models.AuthResp, error) {
	reqBody, err := json.Marshal(models.AuthBody{
		GrantType: CLIENT_CREDENTIALS_TEXT,
	})

	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/auth/oauth2/v2/token", auth.baseUrl), bytes.NewBuffer(reqBody))

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(auth.ClientId, auth.ClientSecret)

	if err != nil {
		return nil, nil, err
	}

	resp, err := auth.client.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return resp, nil, err
	}

	if resp.StatusCode > http.StatusOK {
		return resp, nil, errors.New(fmt.Sprintf("Auth responded with status code of [ %d ]", resp.StatusCode))
	}

	if err != nil {
		return resp, nil, err
	}

	var output models.AuthResp

	err = json.NewDecoder(resp.Body).Decode(&output)

	if err != nil {
		return resp, nil, err
	}

	return resp, &output, nil
}

func Authenticate(auth Authenticator) (*http.Response, *models.AuthResp, error) {
	return auth.Authenticate()
}
