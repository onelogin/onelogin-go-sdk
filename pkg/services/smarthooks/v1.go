package smarthooks

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/onelogin/onelogin-go-sdk/pkg/services"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/olhttp"
	"github.com/onelogin/onelogin-go-sdk/pkg/utils"
)

const errSmartHooksV2Context = "smarthooks v1 service"

// V1Service holds the information needed to interface with a repository
type V1Service struct {
	Endpoint, ErrorContext string
	Repository             services.Repository
}

// New creates the new svc service v2.
func New(repo services.Repository, host string) *V1Service {
	// ugly hack, remove after eks migration is done
	if host == "https://oapi.onelogin-shadow01.com" {
		host = "https://oapi-eks.onelogin-shadow01.com"
	}
	return &V1Service{
		Endpoint:     fmt.Sprintf("%s/api/2/hooks", host),
		Repository:   repo,
		ErrorContext: errSmartHooksV2Context,
	}
}

// Query retrieves all the smarthooks from the repository that meet the query criteria passed in the
// request payload. If an empty payload is given, it will retrieve all smarthooks
func (svc *V1Service) Query(query *SmartHookQuery) ([]SmartHook, error) {
	resp, err := svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        svc.Endpoint,
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    query,
	})
	if err != nil {
		return nil, err
	}

	var smarthooks []SmartHook
	json.Unmarshal(resp, &smarthooks)
	decodedSmarthooks := make([]SmartHook, len(smarthooks))
	for i := range smarthooks {
		decodeFunction(&smarthooks[i])
		decodedSmarthooks[i] = smarthooks[i]
	}
	return decodedSmarthooks, nil
}

// GetOne retrieves the smarthook by id and returns it
func (svc *V1Service) GetOne(id string) (*SmartHook, error) {
	resp, err := svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%s", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	})
	if err != nil {
		return nil, err
	}
	var smarthook SmartHook
	json.Unmarshal(resp, &smarthook)
	decodeFunction(&smarthook)
	return &smarthook, nil
}

// Create takes a smarthook without an id and attempts to use the parameters to create it
// in the API. Modifies the smarthook in place, or returns an error if one occurs
func (svc *V1Service) Create(smarthook *SmartHook) error {
	if encErr := encodeFunction(smarthook); encErr != nil {
		return encErr
	}
	resp, err := svc.Repository.Create(olhttp.OLHTTPRequest{
		URL:        svc.Endpoint,
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    smarthook,
	})
	if err != nil {
		return err
	}
	json.Unmarshal(resp, smarthook)
	return nil
}

// Update takes a smarthook and an id and attempts to use the parameters to update it
// in the API. Modifies the smarthook in place, or returns an error if one occurs
func (svc *V1Service) Update(smarthook *SmartHook) error {
	if smarthook.ID == nil {
		return errors.New("No ID Given")
	}
	if encErr := encodeFunction(smarthook); encErr != nil {
		return encErr
	}
	id := *smarthook.ID
	smarthook.ID = nil
	resp, err := svc.Repository.Update(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%s", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    smarthook,
	})
	if err != nil {
		return err
	}
	json.Unmarshal(resp, smarthook)
	return nil
}

// Destroy deletes the smarthook with the given id, and if successful, it returns nil
func (svc *V1Service) Destroy(id string) error {
	if _, err := svc.Repository.Destroy(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%s", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	}); err != nil {
		return err
	}
	return nil
}

func encodeFunction(smarthook *SmartHook) error {
	if smarthook.Function == nil {
		return errors.New("No Function Definition Given")
	}
	str := utils.EncodeString(*smarthook.Function)
	smarthook.Function = &str
	return nil
}

func decodeFunction(smarthook *SmartHook) error {
	if smarthook.Function == nil {
		return errors.New("No Function Definition Given")
	}
	str := utils.DecodeString(*smarthook.Function)
	smarthook.Function = &str
	return nil
}
