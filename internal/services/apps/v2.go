package apps

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/onelogin/onelogin-go-sdk/internal/services"
	"github.com/onelogin/onelogin-go-sdk/internal/services/olhttp"

	"github.com/onelogin/onelogin-go-sdk/internal/customerrors"
)

const errAppsV2Context = "apps v2 service"

type AppV2Service struct {
	Endpoint, ErrorContext string
	Config                 services.APIServiceConfig
	Auth                   services.Authenticator
}

// New creates the new svc service v2.
func New(cfg *services.APIServiceConfig) AppV2Service {
	return AppV2Service{
		Endpoint:     fmt.Sprintf("%s/api/2/svc", cfg.BaseURL),
		Config:       *cfg,
		ErrorContext: errAppsV2Context,
	}
}

func (svc *AppV2Service) Query(query *AppsQuery) (*http.Response, []App, error) {
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

	params := req.URL.Query()
	if query != nil {
		if query.Limit != nil {
			params.Add("limit", *query.Limit)
		}
		if query.Page != nil {
			params.Add("page", *query.Page)
		}
		if query.Name != nil {
			params.Add("name", *query.Name)
		}
		if query.ConnectorID != nil {
			params.Add("connector_id", *query.ConnectorID)
		}
		if query.AuthMethod != nil {
			params.Add("auth_method", *query.AuthMethod)
		}
		if query.Cursor != nil {
			params.Add("cursor", *query.Cursor)
		}
		req.URL.RawQuery = params.Encode()
	}
	resp, err := svc.Config.Client.Do(req)

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode == 401 {
		_, err := svc.Config.Auth.ReAuthorize()
		if err != nil {
			return nil, nil, customerrors.OneloginErrorWrapper(svc.ErrorContext, err)
		}
		return svc.Query(query)
	}

	var bodyResp []App
	if err = customerrors.OneloginErrorWrapper(svc.ErrorContext, json.NewDecoder(resp.Body).Decode(&bodyResp)); err != nil {
		return resp, nil, err
	}

	return resp, bodyResp, nil
}

// GetAppByID retrieves the app by id, and if successful, it returns
// the http response and the pointer to the app.
func (svc *AppV2Service) GetOne(id int32) (*http.Response, *App, error) {
	accessToken, err := svc.Config.Auth.Authorize()
	if err != nil {
		return nil, nil, customerrors.OneloginErrorWrapper(svc.ErrorContext, err)
	}

	url := fmt.Sprintf("%s/%d", svc.Endpoint, id)
	headers := map[string]string{
		"Content-type":  "application/json",
		"Authorization": "Bearer " + accessToken,
	}

	req, err := olhttp.NewOneloginRequest(url, http.MethodGet, headers, nil)

	if err = customerrors.OneloginErrorWrapper(svc.ErrorContext, err); err != nil {
		return nil, nil, err
	}

	resp, err := svc.Config.Client.Do(req)

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode == 401 {
		_, err := svc.Config.Auth.ReAuthorize()
		if err != nil {
			return nil, nil, customerrors.OneloginErrorWrapper(svc.ErrorContext, err)
		}
		return svc.GetOne(id)
	}

	if resError := customerrors.ReqErrorWrapper(resp, svc.ErrorContext, err); resError != nil {
		return resp, nil, resError
	}

	var bodyResp App
	if err = customerrors.OneloginErrorWrapper(svc.ErrorContext, json.NewDecoder(resp.Body).Decode(&bodyResp)); err != nil {
		return resp, nil, err
	}

	return resp, &bodyResp, nil
}

// CreateApp creates a new app, and if successful, it returns
// the http response and the pointer to the app.
func (svc *AppV2Service) Create(app *App) (*http.Response, *App, error) {
	accessToken, err := svc.Config.Auth.Authorize()
	if err != nil {
		return nil, nil, customerrors.OneloginErrorWrapper(svc.ErrorContext, err)
	}

	headers := map[string]string{
		"Content-type":  "application/json",
		"Authorization": "Bearer " + accessToken,
	}

	req, err := olhttp.NewOneloginRequest(svc.Endpoint, http.MethodPost, headers, app)

	if err = customerrors.OneloginErrorWrapper(svc.ErrorContext, err); err != nil {
		return nil, nil, err
	}

	resp, err := svc.Config.Client.Do(req)

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode == 401 {
		_, err := svc.Config.Auth.ReAuthorize()
		if err != nil {
			return nil, nil, customerrors.OneloginErrorWrapper(svc.ErrorContext, err)
		}
		return svc.Create(app)
	}

	if err = customerrors.ReqErrorWrapper(resp, svc.ErrorContext, err); err != nil {
		return resp, nil, err
	}

	var bodyResp App

	if err = customerrors.OneloginErrorWrapper(svc.ErrorContext, json.NewDecoder(resp.Body).Decode(&bodyResp)); err != nil {
		return resp, nil, err
	}

	return resp, &bodyResp, nil
}

// UpdateAppByID updates an existing app, and if successful, it returns
// the http response and the pointer to the updated app.
func (svc *AppV2Service) Update(id int32, app *App) (*http.Response, *App, error) {
	accessToken, err := svc.Config.Auth.Authorize()
	if err != nil {
		return nil, nil, customerrors.OneloginErrorWrapper(svc.ErrorContext, err)
	}

	url := fmt.Sprintf("%s/%d", svc.Endpoint, id)

	headers := map[string]string{
		"Content-type":  "application/json",
		"Authorization": "Bearer " + accessToken,
	}

	req, err := olhttp.NewOneloginRequest(url, http.MethodPut, headers, app)

	if err = customerrors.OneloginErrorWrapper(svc.ErrorContext, err); err != nil {
		return nil, nil, err
	}

	resp, err := svc.Config.Client.Do(req)

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode == 401 {
		_, err := svc.Config.Auth.ReAuthorize()
		if err != nil {
			return nil, nil, customerrors.OneloginErrorWrapper(svc.ErrorContext, err)
		}
		return svc.Update(id, app)
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	if err = customerrors.ReqErrorWrapper(resp, svc.ErrorContext, err); err != nil {
		return nil, nil, err
	}

	var bodyResp App

	err = json.NewDecoder(resp.Body).Decode(&bodyResp)

	if err = customerrors.OneloginErrorWrapper(svc.ErrorContext, err); err != nil {
		return resp, nil, err
	}

	return resp, &bodyResp, nil
}

// DeleteApp deletes the app for the id, and if successful, it returns
// the http response.
func (svc *AppV2Service) Destroy(id int32) (*http.Response, error) {
	accessToken, err := svc.Config.Auth.Authorize()
	if err != nil {
		return nil, customerrors.OneloginErrorWrapper(svc.ErrorContext, err)
	}

	url := fmt.Sprintf("%s/%d", svc.Endpoint, id)

	headers := map[string]string{
		"Content-type":  "application/json",
		"Authorization": "Bearer " + accessToken,
	}

	req, err := olhttp.NewOneloginRequest(url, http.MethodDelete, headers, nil)

	resp, err := svc.Config.Client.Do(req)

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode == 401 {
		_, err := svc.Config.Auth.ReAuthorize()
		if err != nil {
			return nil, customerrors.OneloginErrorWrapper(svc.ErrorContext, err)
		}
		return svc.Destroy(id)
	}

	if err = customerrors.ReqErrorWrapper(resp, svc.ErrorContext, err); err != nil {
		return nil, err
	}

	return resp, nil
}
