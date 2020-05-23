package apps

import (
	"encoding/json"
	"fmt"

	"github.com/onelogin/onelogin-go-sdk/pkg/services"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/olhttp"
)

const errAppsV2Context = "apps v2 service"

// V2Service holds the information needed to interface with a repository
type V2Service struct {
	Endpoint, ErrorContext string
	Repository             services.Repository
}

// New creates the new svc service v2.
func New(repo services.Repository, host string) V2Service {
	return V2Service{
		Endpoint:     fmt.Sprintf("%s/api/2/apps", host),
		Repository:   repo,
		ErrorContext: errAppsV2Context,
	}
}

// Query retrieves all the apps from the repository that meet the query criteria passed in the
// request payload. If an empty payload is given, it will retrieve all apps
func (svc V2Service) Query(query *AppsQuery) ([]App, error) {
	resp, err := svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        svc.Endpoint,
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    query,
	})
	if err != nil {
		return nil, err
	}

	var apps []App
	if err = json.Unmarshal(resp, &apps); err != nil {
		return nil, err
	}

	for _, app := range apps {
		resp, err := svc.Repository.Read(olhttp.OLHTTPRequest{
			URL:        fmt.Sprintf("%s/%d/rules", svc.Endpoint, app.ID),
			Headers:    map[string]string{"Content-Type": "application/json"},
			AuthMethod: "bearer",
			Payload:    query,
		})
		if err != nil {
			return nil, err
		}
		var rules []AppRule
		if err = json.Unmarshal(resp, &rules); err != nil {
			return nil, err
		}
		app.Rules = rules
	}

	return apps, nil
}

// GetOne retrieves the app by id, and if successful, it returns
// the http response and the pointer to the app.
func (svc *V2Service) GetOne(id int32) (*App, error) {
	resp, err := svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	})
	if err != nil {
		return nil, err
	}

	var app App
	if err = json.Unmarshal(resp, &app); err != nil {
		return nil, err
	}
	return &app, nil
}

// Create creates a new app, and if successful, it returns
// the http response and the pointer to the app.
func (svc *V2Service) Create(app *App) (*App, error) {
	resp, err := svc.Repository.Create(olhttp.OLHTTPRequest{
		URL:        svc.Endpoint,
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    app,
	})
	if err != nil {
		return nil, err
	}
	var newApp App
	if err = json.Unmarshal(resp, &newApp); err != nil {
		return nil, err
	}
	return &newApp, nil
}

// Update updates an existing app, and if successful, it returns
// the http response and the pointer to the updated app.
func (svc *V2Service) Update(id int32, app *App) (*App, error) {
	resp, err := svc.Repository.Update(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    app,
	})
	if err != nil {
		return nil, err
	}

	var updatedApp App
	if err = json.Unmarshal(resp, &updatedApp); err != nil {
		return nil, err
	}
	return &updatedApp, nil
}

// Destroy deletes the app for the id, and if successful, it returns nil
func (svc *V2Service) Destroy(id int32) error {
	_, err := svc.Repository.Destroy(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	})
	if err != nil {
		return err
	}

	return nil
}
