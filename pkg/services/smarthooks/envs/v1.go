package smarthookenvs

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/onelogin/onelogin-go-sdk/pkg/services"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/olhttp"
)

const errEnvVarsV2Context = "envVar environment variables v1 service"

// SmartHookEnvVarsV1Service holds the information needed to interface with a repository
type SmartHookEnvVarsV1Service struct {
	Endpoint, ErrorContext string
	Repository             services.Repository
}

// New creates the new svc service v2.
func New(repo services.Repository, host string) *SmartHookEnvVarsV1Service {
	return &SmartHookEnvVarsV1Service{
		Endpoint:     fmt.Sprintf("%s/api/2/hooks/envs", host),
		Repository:   repo,
		ErrorContext: errEnvVarsV2Context,
	}
}

// Query retrieves all the envVars from the repository that meet the query criteria passed in the
// request payload. If an empty payload is given, it will retrieve all envVars
func (svc *SmartHookEnvVarsV1Service) Query(query *SmartHookEnvVarQuery) ([]EnvVar, error) {
	resp, err := svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        svc.Endpoint,
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    query,
	})
	if err != nil {
		return nil, err
	}

	var envVars []EnvVar
	json.Unmarshal(resp, &envVars)
	return envVars, nil
}

// GetOne retrieves the envVar by id and returns it
func (svc *SmartHookEnvVarsV1Service) GetOne(id string) (*EnvVar, error) {
	out := EnvVar{}
	resp, err := svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%s", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	})
	if err != nil {
		return &out, err
	}
	json.Unmarshal(resp, &out)
	return &out, nil
}

// Create takes a envVar without an id and attempts to use the parameters to create it
// in the API. Modifies the envVar in place, or returns an error if one occurs
func (svc *SmartHookEnvVarsV1Service) Create(envVar *EnvVar) (*EnvVar, error) {
	out := EnvVar{}
	if envVar.Name == nil || envVar.Value == nil {
		return &out, errors.New("Name and Value are both required")
	}

	resp, err := svc.Repository.Create(olhttp.OLHTTPRequest{
		URL:        svc.Endpoint,
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    envVar,
	})
	if err != nil {
		return &out, err
	}

	json.Unmarshal(resp, &out)
	return &out, nil
}

// Update takes a envVar and an id and attempts to use the parameters to update it
// in the API. Modifies the envVar in place, or returns an error if one occurs
func (svc *SmartHookEnvVarsV1Service) Update(envVar *EnvVar) (*EnvVar, error) {
	out := EnvVar{}
	if envVar.ID == nil {
		return &out, errors.New("No ID Given")
	}
	if envVar.Value == nil {
		return &out, errors.New("Value is required")
	}
	if envVar.Name != nil {
		return &out, errors.New("Name not allowed")
	}

	id := *envVar.ID
	envVar.ID = nil
	resp, err := svc.Repository.Update(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%s", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    envVar,
	})
	if err != nil {
		return &out, err
	}

	json.Unmarshal(resp, &out)
	return &out, nil
}

// Destroy deletes the envVar with the given id, and if successful, it returns nil
func (svc *SmartHookEnvVarsV1Service) Destroy(id string) error {
	if _, err := svc.Repository.Destroy(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%s", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	}); err != nil {
		return err
	}
	return nil
}
