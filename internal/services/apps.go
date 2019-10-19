package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/onelogin/onelogin-go-sdk/internal/customerrors"
	"github.com/onelogin/onelogin-go-sdk/pkg/models"
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
		BaseURL:      cfg.BaseURL,
		client:       cfg.Client,
		Auth:         cfg.Auth,
		ErrorContext: "apps v2 service",
	}
}

// GetAppByID retrieves the app by id, and if successful, it returns
// the http response and the pointer to the app.
func (apps *AppsV2) GetAppByID(id int32) (*http.Response, *models.App, error) {
	url := fmt.Sprintf("%s/api/2/apps/%d", apps.BaseURL, id)
	req, err := http.NewRequest("GET", url, nil)

	if err = customerrors.OneloginErrorWrapper(apps.ErrorContext, err); err != nil {
		return nil, nil, err
	}

	respAuth, authResp, err := Authorize(apps.Auth)

	if authResErr := customerrors.ReqErrorWrapper(respAuth, apps.ErrorContext, err); authResErr != nil {
		return nil, nil, authResErr
	}

	req.Header.Set("Content-type", "application/json")
	bearer := "Bearer " + authResp.AccessToken
	req.Header.Add("Authorization", bearer)

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

	appToCreate, err := json.Marshal(app)

	if err = customerrors.OneloginErrorWrapper(apps.ErrorContext, err); err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/2/apps", apps.BaseURL), bytes.NewBuffer(appToCreate))

	if err = customerrors.OneloginErrorWrapper(apps.ErrorContext, err); err != nil {
		return nil, nil, err
	}

	respAuth, authResp, err := Authorize(apps.Auth)

	if err = customerrors.ReqErrorWrapper(respAuth, apps.ErrorContext, err); err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-type", "application/json")
	bearer := "Bearer " + authResp.AccessToken
	req.Header.Add("Authorization", bearer)

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

	appToCreate, err := json.Marshal(app)

	if err = customerrors.OneloginErrorWrapper(apps.ErrorContext, err); err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/api/2/apps/%d", apps.BaseURL, id), bytes.NewBuffer(appToCreate))

	if err = customerrors.OneloginErrorWrapper(apps.ErrorContext, err); err != nil {
		return nil, nil, err
	}

	respAuth, authResp, err := Authorize(apps.Auth)

	if err = customerrors.ReqErrorWrapper(respAuth, apps.ErrorContext, err); err != nil {
		return nil, nil, err
	}

	req.Header.Set("Content-type", "application/json")
	bearer := "Bearer " + authResp.AccessToken
	req.Header.Add("Authorization", bearer)

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
	url := fmt.Sprintf("%s/api/2/apps/%d", apps.BaseURL, id)
	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return nil, err
	}

	respAuth, authResp, err := Authorize(apps.Auth)

	if err = customerrors.ReqErrorWrapper(respAuth, apps.ErrorContext, err); err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")
	bearer := "Bearer " + authResp.AccessToken
	req.Header.Add("Authorization", bearer)

	resp, err := apps.client.Do(req)

	if err = customerrors.ReqErrorWrapper(resp, apps.ErrorContext, err); err != nil {
		return nil, err
	}

	return resp, nil
}
