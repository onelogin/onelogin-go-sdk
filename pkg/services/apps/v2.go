package apps

import (
	"encoding/json"
	"fmt"

	"github.com/onelogin/onelogin-go-sdk/internal/customerrors"
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
func New(repo services.Repository, host string) *V2Service {
	return &V2Service{
		Endpoint:     fmt.Sprintf("%s/api/2/apps", host),
		Repository:   repo,
		ErrorContext: errAppsV2Context,
	}
}

// Query retrieves all the apps from the repository that meet the query criteria passed in the
// request payload. If an empty payload is given, it will retrieve all apps
func (svc *V2Service) Query(query *AppsQuery) ([]App, error) {
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
	json.Unmarshal(resp, &apps)
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
	json.Unmarshal(resp, &app)
	return &app, nil
}

// Create creates a new app, and if successful, it returns
// the http response and the pointer to the app.
func (svc *V2Service) Create(app *App) (*App, error) {
	var newApp App
	resp, err := svc.Repository.Create(olhttp.OLHTTPRequest{
		URL:        svc.Endpoint,
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    app,
	})
	if err != nil {
		return &newApp, err
	}
	json.Unmarshal(resp, &newApp)
	return &newApp, nil
}

// Update updates an existing app, and if successful, it returns
// the http response and the pointer to the updated app.
func (svc *V2Service) Update(id int32, app *App) (*App, error) {
	var updatedApp App
	resp, err := svc.Repository.Update(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    app,
	})
	if err != nil {
		return &updatedApp, err
	}
	json.Unmarshal(resp, &updatedApp)
	pruneParamErr := svc.pruneParameters(&app.Parameters, &updatedApp)

	if pruneParamErr != nil {
		recoveredAppState, recoverErr := svc.GetOne(*updatedApp.ID)
		if recoverErr != nil {
			return nil, err
		}
		return recoveredAppState, pruneParamErr
	}
	// re-read the app so we return one with all the parameters changes made via each individual parameters call
	return svc.GetOne(*updatedApp.ID)
}

// Destroy deletes the app for the id, and if successful, it returns nil
func (svc *V2Service) Destroy(id int32) error {
	if _, err := svc.Repository.Destroy(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	}); err != nil {
		return err
	}
	return nil
}

// Given a list of requested rules, go to the API, and pluck (delete) all the rules that are not on the
// request list. At this point the app holds all existing rules in the API.
// Rules not on the request list are assumed to be removed by the caller.
func (svc *V2Service) pruneParameters(requestedParams *map[string]AppParameters, app *App) error {
	var delErrors []error
	keepMap := make(map[int32]bool, len(*requestedParams))
	for _, param := range *requestedParams {
		keepMap[*param.ID] = true
	}
	// no need to call down app parameters specifically like we do for rules. parameters returned as part of app update
	for _, delCandidate := range app.Parameters {
		if !keepMap[*delCandidate.ID] {
			if _, err := svc.Repository.Destroy(olhttp.OLHTTPRequest{
				URL:        fmt.Sprintf("%s/%d/parameters/%d", svc.Endpoint, *app.ID, *delCandidate.ID),
				Headers:    map[string]string{"Content-Type": "application/json"},
				AuthMethod: "bearer",
			}); err != nil {
				delErrors = append(delErrors, err)
			}
		}
	}
	return customerrors.StackErrors(delErrors)
}
