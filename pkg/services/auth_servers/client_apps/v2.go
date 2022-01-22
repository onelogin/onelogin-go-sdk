package clientapps

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/onelogin/onelogin-go-sdk/pkg/services"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/olhttp"
)

const errAppsV2Context = "access token claims v2 service"

// V2Service holds the information needed to interface with a repository
type V2Service struct {
	Endpoint, ErrorContext string
	Repository             services.Repository
}

// New creates the new svc service v2.
func New(repo services.Repository, host string) *V2Service {
	return &V2Service{
		Endpoint:     fmt.Sprintf("%s/api/2/api_authorizations", host),
		Repository:   repo,
		ErrorContext: errAppsV2Context,
	}
}

// Query retrieves all the access token claims from the repository that meet the query criteria passed in the
// request payload. If an empty payload is given, it will retrieve all access token claims.
func (svc *V2Service) Query(query *ClientAppsQuery) ([]ClientApp, error) {
	resp, err := svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%s/clients", svc.Endpoint, query.AuthServerID),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	})
	if err != nil {
		return nil, err
	}

	var clients []ClientApp
	for _, bytes := range resp {
		var unmarshalled []ClientApp
		json.Unmarshal(bytes, &unmarshalled)
		clients = append(clients, unmarshalled...)
	}
	return clients, nil
}

// Create creates a new access token claim in place and returns an error if something went wrong
func (svc *V2Service) Create(clientApp *ClientApp) error {
	if clientApp.AuthServerID == nil {
		return errors.New("AuthServerID required on the payload")
	}
	resp, err := svc.Repository.Create(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d/clients", svc.Endpoint, *clientApp.AuthServerID),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    clientApp,
	})
	if err != nil {
		return err
	}

	respObj := map[string]int32{}
	json.Unmarshal(resp, &respObj)
	clientApp.ID = oltypes.Int32(respObj["app_id"])
	clientApp.APIAuthID = oltypes.Int32(respObj["api_auth_id"])

	return nil
}

// Update updates an existing access token claim in place or returns an error if something went wrong
func (svc *V2Service) Update(clientApp *ClientApp) error {
	if clientApp.ID == nil || clientApp.AuthServerID == nil {
		return errors.New("Both ID and AuthServerID are required on the payload")
	}
	resp, err := svc.UpdateRaw(*clientApp.AuthServerID, *clientApp.ID, clientApp)
	if err != nil {
		return err
	}

	respObj := map[string]int32{}
	json.Unmarshal(resp, &respObj)
	clientApp.ID = oltypes.Int32(respObj["app_id"])
	clientApp.APIAuthID = oltypes.Int32(respObj["api_auth_id"])

	return nil
}

// UpdateRaw updates an existing access token claim and returns the raw response or an error if something went wrong
func (svc *V2Service) UpdateRaw(authServerId int32, appId int32, clientApp interface{}) ([]byte, error) {
	return svc.Repository.Update(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d/clients/%d", svc.Endpoint, authServerId, appId),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    clientApp,
	})
}

// Destroy takes the access token claim id and access token claim id and removes the access token claim from the API.
// Returns an error if something went wrong.
func (svc *V2Service) Destroy(clientId int32, id int32) error {
	if _, err := svc.Repository.Destroy(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d/clients/%d", svc.Endpoint, clientId, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	}); err != nil {
		return err
	}
	return nil
}
