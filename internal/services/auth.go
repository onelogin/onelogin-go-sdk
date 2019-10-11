package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/onelogin/onelogin-go-sdk/internal/customerrors"
	"github.com/onelogin/onelogin-go-sdk/pkg/models"
)

const (
	ClientCredentialsText = "client_credentials"
)

type Authenticator interface {
	Authorize() (*http.Response, *models.AuthResp, error)
}

type AuthV2 struct {
	ClientId     string
	ClientSecret string
	client       *http.Client
	baseUrl      string
	ErrorContext string
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
		ErrorContext: "auth v2 service",
	}
}

// Authorize authorizes the credentials for the ClientId and ClientSecret, and returns, if successfull,
// the http response and the auth response.
func (auth *AuthV2) Authorize() (*http.Response, *models.AuthResp, error) {
	reqBody, err := json.Marshal(models.AuthBody{
		GrantType: ClientCredentialsText,
	})

	oneloginErr := customerrors.OneloginErrorWrapper(auth.ErrorContext, err)

	if oneloginErr != nil {
		return nil, nil, oneloginErr
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/auth/oauth2/v2/token", auth.baseUrl), bytes.NewBuffer(reqBody))

	oneloginErr = customerrors.OneloginErrorWrapper(auth.ErrorContext, err)

	if oneloginErr != nil {
		return nil, nil, oneloginErr
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(auth.ClientId, auth.ClientSecret)

	resp, err := auth.client.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	respErr := customerrors.ReqErrorWrapper(resp, auth.ErrorContext, err)

	if respErr != nil {
		return resp, nil, respErr
	}

	var output models.AuthResp

	oneloginErr = customerrors.OneloginErrorWrapper(auth.ErrorContext, json.NewDecoder(resp.Body).Decode(&output))

	if oneloginErr != nil {
		return resp, nil, oneloginErr
	}

	return resp, &output, nil
}

func Authorize(auth Authenticator) (*http.Response, *models.AuthResp, error) {
	return auth.Authorize()
}
