package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/onelogin/onelogin-go-sdk/internal/customerrors"
	"github.com/onelogin/onelogin-go-sdk/pkg/models"
)

// ClientCredentialsText is the key used for the auth request.
// errAuthContext is the context.
const (
	ClientCredentialsText = "client_credentials"
	errAuthContext        = "auth v2 service"
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
		ErrorContext: errAuthContext,
	}
}

// Authorize authorizes the credentials for the ClientId and ClientSecret, and returns, if successfull,
// the http response and the auth response.
func (auth *AuthV2) Authorize() (*http.Response, *models.AuthResp, error) {
	payload := models.AuthBody{
		GrantType: ClientCredentialsText,
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	url := fmt.Sprintf("%s/auth/oauth2/v2/token", auth.baseURL)

	req, err := setUpRequest(url, http.MethodPost, headers, payload)

	if oneloginErr := customerrors.OneloginErrorWrapper(auth.ErrorContext, err); oneloginErr != nil {
		return nil, nil, oneloginErr
	}

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
