package clientcredentials

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/onelogin/onelogin-go-sdk/internal/customerrors"
	"github.com/onelogin/onelogin-go-sdk/internal/services"
	"github.com/onelogin/onelogin-go-sdk/internal/services/olhttp"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
)

const (
	// ClientCredentialsText is the key used for the auth request.
	ClientCredentialsText = "client_credentials"

	// errAuthContext is the context.
	errAuthContext = "auth v2 service"
)

type AuthService struct {
 	Endpoint, Token, ErrorContext string
	Config services.AuthServiceConfig
	Client *http.Client
}
// New uses the cfg to generate the new auth service, and returns
// the created auth service for version 2.
func New(cfg *services.AuthServiceConfig) *AuthService {
	return &AuthService{
		Endpoint: fmt.Sprintf("%s/auth/oauth2/v2/token", cfg.BaseURL),
		Config: *cfg,
		ErrorContext: errAuthContext,
	}
}

// Create authorizes the credentials for the ClientId and ClientSecret, and returns, if successfull,
// the http response and the auth response.
func (auth *AuthService) Create() (*http.Response, *ClientCredential, error) {
	payload := AuthBody{
		GrantType: ClientCredentialsText,
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	url := fmt.Sprintf("%s/auth/oauth2/v2/token", auth.Config.BaseURL)

	req, err := olhttp.NewOneloginRequest(url, http.MethodPost, headers, payload)

	if oneloginErr := customerrors.OneloginErrorWrapper(auth.ErrorContext, err); oneloginErr != nil {
		return nil, nil, oneloginErr
	}

	req.SetBasicAuth(auth.Config.ClientID, auth.Config.ClientSecret)

	resp, err := auth.Client.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	if respErr := customerrors.ReqErrorWrapper(resp, auth.ErrorContext, err); respErr != nil {
		return resp, nil, respErr
	}

	var output ClientCredential

	if oneloginErr := customerrors.OneloginErrorWrapper(auth.ErrorContext, json.NewDecoder(resp.Body).Decode(&output)); oneloginErr != nil {
		return resp, nil, oneloginErr
	}

	return resp, &output, nil
}

// Authorize calls the Authorize function of the Authenticator interface. Adds a fresh access token if none exists
func (svc *AuthService) Authorize() (string, error) {
	if svc.Token == "" {
		if err := setToken(svc); err != nil {
			return "", err
		}
	}
	return svc.Token, nil
}

// ReAuthorize calls the Authorize function of the Authenticator interface. Overwrites the access token with a new token
func (svc *AuthService) ReAuthorize() (string, error) {
	if err := setToken(svc); err != nil {
		return "", err
	}
	return svc.Token, nil
}

func setToken(svc *AuthService) error {
	_, cred, err := svc.Create()
	if err != nil {
		return  err
	}
	svc.Token, _ = oltypes.GetStringVal(cred.AccessToken)
	return nil
}
