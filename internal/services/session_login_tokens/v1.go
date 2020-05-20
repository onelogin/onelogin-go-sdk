package sessionlogintokens

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/onelogin/onelogin-go-sdk/internal/services"
	"github.com/onelogin/onelogin-go-sdk/internal/services/client_credentials"
	"github.com/onelogin/onelogin-go-sdk/internal/services/olhttp"

	"github.com/onelogin/onelogin-go-sdk/internal/customerrors"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
)

// #nosec
const (
	errSessionLoginTokenV1Context = "apps v2 service"
)

// NewSessionLoginTokenV1 creates the new apps service v2.
func New(cfg *services.APIServiceConfig) *services.APIService {
	return &services.APIService{
		BaseURL:      fmt.Sprintf("%s/api/1/login/auth", cfg.BaseURL),
		Client:       cfg.Client,
		Auth:         cfg.Auth,
		ErrorContext: errSessionLoginTokenV1Context,
	}
}

// CreateSessionLoginToken takes a SessionLoginToken request that represents an end-user's credentials
// and returns a Session Token that represents an authenticated session
func (session_login_tokens *SessionLoginTokenV1) CreateSessionLoginToken(request *SessionLoginTokenRequest) (*http.Response, *SessionLoginToken, error) {
	respAuth, clientCredential, err := clientcredentials.Authorize(session_login_tokens.Auth)
	fmt.Println(*clientCredential.AccessToken)
	if respErr := customerrors.ReqErrorWrapper(respAuth, session_login_tokens.ErrorContext, err); respErr != nil {
		return nil, nil, respErr
	}

	accessToken, isValid := oltypes.GetStringVal(clientCredential.AccessToken)

	if !isValid {
		return nil, nil, customerrors.OneloginErrorWrapper(session_login_tokens.ErrorContext, customerrors.ErrValueMissing)
	}

	headers := map[string]string{
		"Content-type":  "application/json",
		"Authorization": "Bearer " + accessToken,
	}

	req, err := olhttp.NewOneloginRequest(session_login_tokens.BaseURL, http.MethodPost, headers, request)

	if err = customerrors.OneloginErrorWrapper(session_login_tokens.ErrorContext, err); err != nil {
		return nil, nil, err
	}

	resp, err := session_login_tokens.client.Do(req)

	if err = customerrors.ReqErrorWrapper(resp, session_login_tokens.ErrorContext, err); err != nil {
		return resp, nil, err
	}

	var bodyResp SessionLoginToken

	err = json.NewDecoder(resp.Body).Decode(&bodyResp)

	if err = customerrors.OneloginErrorWrapper(session_login_tokens.ErrorContext, err); err != nil {
		return resp, nil, err
	}

	return resp, &bodyResp, nil
}
