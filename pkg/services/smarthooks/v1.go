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
func (svc *V1Service) Query(query *SmartHookQuery) ([]InflatedSmartHook, error) {
	resp, err := svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        svc.Endpoint,
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    query,
	})
	if err != nil {
		return nil, err
	}

	var smarthooks []InflatedSmartHook
	json.Unmarshal(resp, &smarthooks)
	decodedSmarthooks := make([]InflatedSmartHook, len(smarthooks))
	for i := range smarthooks {
		decodeFunction(&smarthooks[i])
		decodedSmarthooks[i] = smarthooks[i]
	}
	return decodedSmarthooks, nil
}

// GetOne retrieves the smarthook by id and returns it
func (svc *V1Service) GetOne(id string) (*InflatedSmartHook, error) {
	out := InflatedSmartHook{}
	resp, err := svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%s", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	})
	if err != nil {
		return &out, err
	}
	json.Unmarshal(resp, &out)
	decodeFunction(&out)
	return &out, nil
}

// Create takes a smarthook without an id and attempts to use the parameters to create it
// in the API. Modifies the smarthook in place, or returns an error if one occurs
func (svc *V1Service) Create(smarthook *SmartHook) (*InflatedSmartHook, error) {
	out := InflatedSmartHook{}
	if encErr := encodeFunction(smarthook); encErr != nil {
		return &out, encErr
	}
	resp, err := svc.Repository.Create(olhttp.OLHTTPRequest{
		URL:        svc.Endpoint,
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    smarthook,
	})
	if err != nil {
		return &out, err
	}

	json.Unmarshal(resp, &out)
	return &out, nil
}

// Update takes a smarthook and an id and attempts to use the parameters to update it
// in the API. Modifies the smarthook in place, or returns an error if one occurs
func (svc *V1Service) Update(smarthook *SmartHook) (*InflatedSmartHook, error) {
	out := InflatedSmartHook{}
	if smarthook.ID == nil {
		return &out, errors.New("No ID Given")
	}
	if encErr := encodeFunction(smarthook); encErr != nil {
		return &out, encErr
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
		return &out, err
	}

	json.Unmarshal(resp, &out)
	return &out, nil
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

func encodeFunction(smarthook HasEncodable) error {
	if smarthook.GetEncodableField() == nil {
		return errors.New("No Function Definition Given")
	}
	str := utils.EncodeString(*smarthook.GetEncodableField())
	smarthook.SetEncodableField(&str)
	return nil
}

func decodeFunction(smarthook HasEncodable) error {
	if smarthook.GetEncodableField() == nil {
		return errors.New("No Function Definition Given")
	}
	str := utils.DecodeString(*smarthook.GetEncodableField())
	smarthook.SetEncodableField(&str)
	return nil
}
