package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"

	"github.com/onelogin/onelogin-go-sdk/internal/customerrors"
	"github.com/onelogin/onelogin-go-sdk/pkg/models"
)

const (
	errAppsV2Context = "apps v2 service"
)

// AppsV2 is the apps service v2.
type AppsV2 struct {
	BaseURL      string
	client       *http.Client
	Auth         Authenticator
	ErrorContext string
}

// AppsV2Config is the config for apps service v2.
type AppsV2Config struct {
	BaseURL string
	Client  *http.Client
	Auth    Authenticator
}

// NewAppsV2 creates the new apps service v2.
func NewAppsV2(cfg *AppsV2Config) *AppsV2 {
	return &AppsV2{
		BaseURL:      fmt.Sprintf("%s/api/2/apps", cfg.BaseURL),
		client:       cfg.Client,
		Auth:         cfg.Auth,
		ErrorContext: errAppsV2Context,
	}
}

func (apps *AppsV2) GetAppRules(appId int32) (*http.Response, []models.AppRule, error) {
	respAuth, authResp, err := Authorize(apps.Auth)

	if authResErr := customerrors.ReqErrorWrapper(respAuth, apps.ErrorContext, err); authResErr != nil {
		return nil, nil, authResErr
	}

	accessToken, isValid := oltypes.GetStringVal(authResp.AccessToken)

	if !isValid {
		return nil, nil, customerrors.OneloginErrorWrapper(apps.ErrorContext, ErrValueMissing)
	}

	headers := map[string]string{
		"Content-type":  "application/json",
		"Authorization": "Bearer " + accessToken,
	}

	req, err := setUpRequest(apps.BaseURL, http.MethodGet, headers, nil)

	if err = customerrors.OneloginErrorWrapper(apps.ErrorContext, err); err != nil {
		return nil, nil, err
	}
	resp, err := apps.client.Do(req)
	if err != nil {
		return nil, []models.AppRule{}, nil
	}
	return resp, []models.AppRule{}, nil
}

func (apps *AppsV2) GetApps(query *models.AppsQuery) (*http.Response, []models.App, error) {
	respAuth, authResp, err := Authorize(apps.Auth)

	if authResErr := customerrors.ReqErrorWrapper(respAuth, apps.ErrorContext, err); authResErr != nil {
		return nil, nil, authResErr
	}

	accessToken, isValid := oltypes.GetStringVal(authResp.AccessToken)

	if !isValid {
		return nil, nil, customerrors.OneloginErrorWrapper(apps.ErrorContext, ErrValueMissing)
	}

	headers := map[string]string{
		"Content-type":  "application/json",
		"Authorization": "Bearer " + accessToken,
	}

	req, err := setUpRequest(apps.BaseURL, http.MethodGet, headers, nil)

	if err = customerrors.OneloginErrorWrapper(apps.ErrorContext, err); err != nil {
		return nil, nil, err
	}

	params := req.URL.Query()

	if query.Limit != "" {
		params.Add("limit", query.Limit)
	}

	if query.Page != "" {
		params.Add("page", query.Page)
	}
	if query.Name != "" {
		params.Add("name", query.Name)
	}
	if query.ConnectorID != "" {
		params.Add("connector_id", query.ConnectorID)
	}
	if query.AuthMethod != "" {
		params.Add("auth_method", query.AuthMethod)
	}
	if query.Cursor != "" {
		params.Add("cursor", query.Cursor)
	}
	req.URL.RawQuery = params.Encode()

	resp, err := apps.client.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	var bodyResp []models.App

	if err = customerrors.OneloginErrorWrapper(apps.ErrorContext, json.NewDecoder(resp.Body).Decode(&bodyResp)); err != nil {
		return resp, nil, err
	}

	return resp, bodyResp, nil
}

// GetAppByID retrieves the app by id, and if successful, it returns
// the http response and the pointer to the app.
func (apps *AppsV2) GetAppByID(id int32) (*http.Response, *models.App, error) {
	respAuth, authResp, err := Authorize(apps.Auth)

	if authResErr := customerrors.ReqErrorWrapper(respAuth, apps.ErrorContext, err); authResErr != nil {
		return nil, nil, authResErr
	}

	accessToken, isValid := oltypes.GetStringVal(authResp.AccessToken)

	if !isValid {
		return nil, nil, customerrors.OneloginErrorWrapper(apps.ErrorContext, ErrValueMissing)
	}

	url := fmt.Sprintf("%s/%d", apps.BaseURL, id)

	headers := map[string]string{
		"Content-type":  "application/json",
		"Authorization": "Bearer " + accessToken,
	}

	req, err := setUpRequest(url, http.MethodGet, headers, nil)

	if err = customerrors.OneloginErrorWrapper(apps.ErrorContext, err); err != nil {
		return nil, nil, err
	}

	resp, err := apps.client.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	if resError := customerrors.ReqErrorWrapper(resp, apps.ErrorContext, err); resError != nil {
		return resp, nil, resError
	}

	var bodyResp models.App

	if err = customerrors.OneloginErrorWrapper(apps.ErrorContext, json.NewDecoder(resp.Body).Decode(&bodyResp)); err != nil {
		return resp, nil, err
	}

	return resp, &bodyResp, nil
}

// CreateApp creates a new app, and if successful, it returns
// the http response and the pointer to the app.
func (apps *AppsV2) CreateApp(app *models.App) (*http.Response, *models.App, error) {

	respAuth, authResp, err := Authorize(apps.Auth)

	if respErr := customerrors.ReqErrorWrapper(respAuth, apps.ErrorContext, err); respErr != nil {
		return nil, nil, respErr
	}

	accessToken, isValid := oltypes.GetStringVal(authResp.AccessToken)

	if !isValid {
		return nil, nil, customerrors.OneloginErrorWrapper(apps.ErrorContext, ErrValueMissing)
	}

	headers := map[string]string{
		"Content-type":  "application/json",
		"Authorization": "Bearer " + accessToken,
	}

	req, err := setUpRequest(apps.BaseURL, http.MethodPost, headers, app)

	if err = customerrors.OneloginErrorWrapper(apps.ErrorContext, err); err != nil {
		return nil, nil, err
	}

	resp, err := apps.client.Do(req)

	if err = customerrors.ReqErrorWrapper(resp, apps.ErrorContext, err); err != nil {
		return resp, nil, err
	}

	var bodyResp models.App

	err = json.NewDecoder(resp.Body).Decode(&bodyResp)

	if err = customerrors.OneloginErrorWrapper(apps.ErrorContext, err); err != nil {
		return resp, nil, err
	}

	return resp, &bodyResp, nil
}

// UpdateAppByID updates an existing app, and if successful, it returns
// the http response and the pointer to the updated app.
func (apps *AppsV2) UpdateAppByID(id int32, app *models.App) (*http.Response, *models.App, error) {
	respAuth, authResp, err := Authorize(apps.Auth)

	if respErr := customerrors.ReqErrorWrapper(respAuth, apps.ErrorContext, err); respErr != nil {
		return nil, nil, respErr
	}

	accessToken, isValid := oltypes.GetStringVal(authResp.AccessToken)

	if !isValid {
		return nil, nil, customerrors.OneloginErrorWrapper(apps.ErrorContext, ErrValueMissing)
	}

	url := fmt.Sprintf("%s/%d", apps.BaseURL, id)

	headers := map[string]string{
		"Content-type":  "application/json",
		"Authorization": "Bearer " + accessToken,
	}

	req, err := setUpRequest(url, http.MethodPut, headers, app)

	if err = customerrors.OneloginErrorWrapper(apps.ErrorContext, err); err != nil {
		return nil, nil, err
	}

	resp, err := apps.client.Do(req)

	if resp != nil {
		defer resp.Body.Close()
	}

	if err = customerrors.ReqErrorWrapper(resp, apps.ErrorContext, err); err != nil {
		return nil, nil, err
	}

	var bodyResp models.App

	err = json.NewDecoder(resp.Body).Decode(&bodyResp)

	if err = customerrors.OneloginErrorWrapper(apps.ErrorContext, err); err != nil {
		return resp, nil, err
	}

	return resp, &bodyResp, nil
}

// DeleteApp deletes the app for the id, and if successful, it returns
// the http response.
func (apps *AppsV2) DeleteApp(id int32) (*http.Response, error) {

	respAuth, authResp, err := Authorize(apps.Auth)

	if err = customerrors.ReqErrorWrapper(respAuth, apps.ErrorContext, err); err != nil {
		return nil, err
	}

	accessToken, isValid := oltypes.GetStringVal(authResp.AccessToken)

	if !isValid {
		return nil, customerrors.OneloginErrorWrapper(apps.ErrorContext, ErrValueMissing)
	}

	url := fmt.Sprintf("%s/%d", apps.BaseURL, id)

	headers := map[string]string{
		"Content-type":  "application/json",
		"Authorization": "Bearer " + accessToken,
	}

	req, err := setUpRequest(url, http.MethodDelete, headers, nil)

	resp, err := apps.client.Do(req)

	if err = customerrors.ReqErrorWrapper(resp, apps.ErrorContext, err); err != nil {
		return nil, err
	}

	return resp, nil
}
