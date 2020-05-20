package sessionlogintokens

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/onelogin/onelogin-go-sdk/internal/services"
	"github.com/onelogin/onelogin-go-sdk/internal/services/olhttp"

	"github.com/onelogin/onelogin-go-sdk/internal/customerrors"
)

// #nosec
const errSessionLoginTokenV1Context = "Session Login Tokens v2 service"

type V1Service struct {
	Endpoint, ErrorContext string
	Config                 services.APIServiceConfig
	Auth                   services.Authenticator
}

// NewSessionLoginTokenV1 creates the new apps service v2.
func New(cfg *services.APIServiceConfig) V1Service {
	return V1Service{
		Endpoint:     fmt.Sprintf("%s/api/1/login/auth", cfg.BaseURL),
		Config:       *cfg,
		ErrorContext: errSessionLoginTokenV1Context,
	}
}

// CreateSessionLoginToken takes a SessionLoginToken request that represents an end-user's credentials
// and returns a Session Token that represents an authenticated session
func (svc *V1Service) Create(request *SessionLoginTokenRequest) (*http.Response, *SessionLoginToken, error) {
	accessToken, err := svc.Config.Auth.Authorize()
	if err != nil {
		return nil, nil, customerrors.OneloginErrorWrapper(svc.ErrorContext, err)
	}

	headers := map[string]string{
		"Content-type":  "application/json",
		"Authorization": "Bearer " + accessToken,
	}

	req, err := olhttp.NewOneloginRequest(svc.Endpoint, http.MethodGet, headers, nil)
	if err = customerrors.OneloginErrorWrapper(svc.ErrorContext, err); err != nil {
		return nil, nil, err
	}

	resp, err := svc.Config.Client.Do(req)

	if resp.StatusCode == 401 {

		_, err := svc.Config.Auth.ReAuthorize()
		if err != nil {
			return nil, nil, customerrors.OneloginErrorWrapper(svc.ErrorContext, err)
		}
		return svc.Create(request)
	}

	if err = customerrors.ReqErrorWrapper(resp, svc.ErrorContext, err); err != nil {
		return resp, nil, err
	}

	var bodyResp SessionLoginToken

	err = json.NewDecoder(resp.Body).Decode(&bodyResp)

	if err = customerrors.OneloginErrorWrapper(svc.ErrorContext, err); err != nil {
		return resp, nil, err
	}

	return resp, &bodyResp, nil
}
