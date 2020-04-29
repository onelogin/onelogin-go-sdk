package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/onelogin/onelogin-go-sdk/internal/customerrors"
	"github.com/onelogin/onelogin-go-sdk/pkg/models"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
)

// #nosec
const (
	errSessionLoginTokenV1Context = "apps v2 service"
)

// SessionLoginTokenV1 is the apps service v2.
type SessionLoginTokenV1 struct {
	BaseURL      string
	client       *http.Client
	Auth         Authenticator
	ErrorContext string
}

// SessionLoginTokenV1Config is the config for apps service v2.
type SessionLoginTokenV1Config struct {
	BaseURL string
	Client  *http.Client
	Auth    Authenticator
}

// NewSessionLoginTokenV1 creates the new apps service v2.
func NewSessionLoginTokenV1(cfg *SessionLoginTokenV1Config) *SessionLoginTokenV1 {
	return &SessionLoginTokenV1{
		BaseURL:      fmt.Sprintf("%s/api/1/login/auth", cfg.BaseURL),
		client:       cfg.Client,
		Auth:         cfg.Auth,
		ErrorContext: errSessionLoginTokenV1Context,
	}
}

// CreateSessionLoginToken takes a SessionLoginToken request that represents an end-user's credentials
// and returns a Session Token that represents an authenticated session
func (session_login_tokens *SessionLoginTokenV1) CreateSessionLoginToken(request *models.SessionLoginTokenRequest) (*http.Response, *models.SessionLoginToken, error) {
	respAuth, authResp, err := Authorize(session_login_tokens.Auth)
	fmt.Println(*authResp.AccessToken)
	if respErr := customerrors.ReqErrorWrapper(respAuth, session_login_tokens.ErrorContext, err); respErr != nil {
		return nil, nil, respErr
	}

	accessToken, isValid := oltypes.GetStringVal(authResp.AccessToken)

	if !isValid {
		return nil, nil, customerrors.OneloginErrorWrapper(session_login_tokens.ErrorContext, ErrValueMissing)
	}

	headers := map[string]string{
		"Content-type":  "application/json",
		"Authorization": "Bearer " + accessToken,
	}

	req, err := setUpRequest(session_login_tokens.BaseURL, http.MethodPost, headers, request)

	if err = customerrors.OneloginErrorWrapper(session_login_tokens.ErrorContext, err); err != nil {
		return nil, nil, err
	}

	resp, err := session_login_tokens.client.Do(req)

	if err = customerrors.ReqErrorWrapper(resp, session_login_tokens.ErrorContext, err); err != nil {
		return resp, nil, err
	}

	var bodyResp models.SessionLoginToken

	err = json.NewDecoder(resp.Body).Decode(&bodyResp)

	if err = customerrors.OneloginErrorWrapper(session_login_tokens.ErrorContext, err); err != nil {
		return resp, nil, err
	}

	return resp, &bodyResp, nil
}
