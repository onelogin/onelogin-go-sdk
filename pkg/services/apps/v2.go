package apps

import (
	"encoding/json"
	"errors"
	"fmt"

	customerror "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/error"
	mod "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
	utils "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/utilities"
)

const errAppsV2Context = "apps v2 service"

// V2Service holds the information needed to interface with a repository
type V2Service struct {
	Endpoint, ErrorContext string
	Repository             utils.Repository
}

// New creates the new svc service v2.
func New(repo utils.Repository, host string) *V2Service {
	return &V2Service{
		Endpoint:     fmt.Sprintf("%s/api/2/apps", host),
		Repository:   repo,
		ErrorContext: errAppsV2Context,
	}
}

// Query retrieves all the apps from the repository that meet the query criteria passed in the
// request payload. If an empty payload is given, it will retrieve all apps.
func (svc *V2Service) Query(query *mod.AppQuery) ([]mod.App, error) {
	resp, err := svc.Repository.Read(mod.OLHTTPRequest{
		URL:        svc.Endpoint,
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    query,
	})
	if err != nil {
		return nil, err
	}

	var apps []mod.App
	for _, bytes := range resp {
		var unmarshalled []mod.App
		json.Unmarshal(bytes, &unmarshalled)
		if len(apps) == 0 {
			apps = unmarshalled
		} else {
			apps = append(apps, unmarshalled...)
		}
	}

	return apps, nil
}

// GetOne retrieves the app by id, and if successful, it returns
// a pointer to the app.
func (svc *V2Service) GetOne(id int32) (*mod.App, error) {
	resp, err := svc.Repository.Read(mod.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	})
	if err != nil {
		return nil, err
	}

	var app mod.App
	if len(resp) < 1 {
		return nil, errors.New("invalid length of response returned")
	}
	json.Unmarshal(resp[0], &app)

	return &app, nil
}

// GetUsers retrieves the list of users for a given app by id, it returns
// an array of users for the app.
func (svc *V2Service) GetUsers(id int32) ([]mod.UserApp, error) {
	return svc.GetUsersWithQuery(id, nil)
}

// GetUsersWithQuery retrieves the list of users for a given app by id with optional query parameters, 
// it returns an array of users for the app.
func (svc *V2Service) GetUsersWithQuery(id int32, query mod.Queryable) ([]mod.UserApp, error) {
	resp, err := svc.Repository.Read(mod.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d/users", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    query,
	})
	if err != nil {
		return nil, err
	}

	var users []mod.UserApp
	for _, bytes := range resp {
		var unmarshalled []mod.UserApp
		json.Unmarshal(bytes, &unmarshalled)
		if len(users) == 0 {
			users = unmarshalled
		} else {
			users = append(users, unmarshalled...)
		}
	}

	return users, nil
}

// Create creates a new app, and if successful, it returns a pointer to the app.
func (svc *V2Service) Create(app *mod.App) error {
	resp, err := svc.Repository.Create(mod.OLHTTPRequest{
		URL:        svc.Endpoint,
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    app,
	})
	if err != nil {
		return err
	}
	json.Unmarshal(resp, app)
	return nil
}

// Update updates an existing app in place or returns an error if something went wrong

// Update is unique in that the API does not fully support Parameters as first-class
// resources and are managed by nesting them in the App. This means that a partial
// update state could exist if, for example, a parameter failed to delete or be updated
// while other parameter changes succeeded. In order to ensure the client is given an
// accurate representation of what has been persisted to the API, we call out to the GetOne
// to simply return what is currently in the API, rather than updating in place. This is a
// temporary holdover until parameters is dealt with in a consistent fashion to other nested resources like app rules
func (svc *V2Service) Update(app *mod.App) (*mod.App, error) {
	if app.ID == nil {
		return nil, errors.New("no ID given")
	}
	requestedParametersState := make(map[string]mod.Parameter, len(*app.Parameters))
	for k, p := range *app.Parameters {
		requestedParametersState[k] = p
	}
	resp, err := svc.Repository.Update(mod.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d", svc.Endpoint, *app.ID),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    app,
	})
	if err != nil {
		return &mod.App{}, err
	}
	json.Unmarshal(resp, app)

	pruneParamErr := svc.pruneParameters(requestedParametersState, app)

	if pruneParamErr != nil {
		var recoverErr error
		app, recoverErr = svc.GetOne(*app.ID)
		if recoverErr != nil {
			return nil, err
		}
		return app, pruneParamErr
	}
	// re-read the app so we return one with all the parameters changes made via each individual parameters call
	return svc.GetOne(*app.ID)
}

// Destroy deletes the app for the id, and if successful, it returns nil
func (svc *V2Service) Destroy(id int32) error {
	if _, err := svc.Repository.Destroy(mod.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	}); err != nil {
		return err
	}
	return nil
}

// Given a list of requested parameters, go to the API, and pluck (delete) all the parameters that are not on the
// request list. At this point the app holds all existing parameters in the API.
// Rules not on the request list are assumed to be removed by the caller.
func (svc *V2Service) pruneParameters(requestedParams map[string]mod.Parameter, app *mod.App) error {
	var delErrors []error

	// Create a map to keep track of the parameters that should be kept
	keepMap := make(map[int]bool, len(requestedParams))

	// Populate the keepMap with the requested parameters
	for _, param := range requestedParams {
		if param.ID != 0 {
			keepMap[param.ID] = true
		}
	}

	// Iterate through the app's parameters to find those that need to be deleted
	for _, delCandidate := range *app.Parameters {
		// Check if this parameter should be kept
		if !keepMap[delCandidate.ID] {
			// If not, delete it via the API
			_, err := svc.Repository.Destroy(mod.OLHTTPRequest{
				URL:        fmt.Sprintf("%s/%d/parameters/%d", svc.Endpoint, *app.ID, delCandidate.ID),
				Headers:    map[string]string{"Content-Type": "application/json"},
				AuthMethod: "bearer",
			})

			// If there's an error deleting the parameter, accumulate it
			if err != nil {
				delErrors = append(delErrors, err)
			}
		}
	}
	return customerror.StackErrors(delErrors)
}
