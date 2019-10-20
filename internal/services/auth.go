package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/onelogin/onelogin-go-sdk/internal/customerrors"
	"github.com/onelogin/onelogin-go-sdk/pkg/models"
)

// Auth service constants.
const (
	ClientCredentialsText = "client_credentials"
)

// Authenticator is an interface.
type Authenticator interface {
	Authorize() (*http.Response, *models.AuthResp, error)
}

// AuthV2 is a service for authorization.
type AuthV2 struct {
	ClientID     string
	ClientSecret string
	client       *http.Client
	baseURL      string
	ErrorContext string
}

// AuthConfigV2 is the config for the auth service v2.
type AuthConfigV2 struct {
	ClientID     string
	ClientSecret string
	Client       *http.Client
	BaseURL      string
}

// NewAuthV2 uses the cfg to generate the new auth service, and returns
// the created auth service for version 2.
func NewAuthV2(cfg *AuthConfigV2) *AuthV2 {
	return &AuthV2{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		client:       cfg.Client,
		baseURL:      cfg.BaseURL,
		ErrorContext: "auth v2 service",
	}
}

// Authorize authorizes the credentials for the ClientId and ClientSecret, and returns, if successfull,
// the http response and the auth response.
func (auth *AuthV2) Authorize() (*http.Response, *models.AuthResp, error) {
	reqBody, err := json.Marshal(models.AuthBody{
		GrantType: ClientCredentialsText,
	})

	if oneloginErr := customerrors.OneloginErrorWrapper(auth.ErrorContext, err); oneloginErr != nil {
		return nil, nil, oneloginErr
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/auth/oauth2/v2/token", auth.baseURL), bytes.NewBuffer(reqBody))

	if oneloginErr := customerrors.OneloginErrorWrapper(auth.ErrorContext, err); oneloginErr != nil {
		return nil, nil, oneloginErr
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(auth.ClientID, auth.ClientSecret)

	resp, err := auth.client.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	if respErr := customerrors.ReqErrorWrapper(resp, auth.ErrorContext, err); respErr != nil {
		return resp, nil, respErr
	}

	var output models.AuthResp

	if oneloginErr := customerrors.OneloginErrorWrapper(auth.ErrorContext, json.NewDecoder(resp.Body).Decode(&output)); oneloginErr != nil {
		return resp, nil, oneloginErr
	}

	return resp, &output, nil
}

// Authorize calls the Authorize function of the Authenticator interface.
func Authorize(auth Authenticator) (*http.Response, *models.AuthResp, error) {
	return auth.Authorize()
}
