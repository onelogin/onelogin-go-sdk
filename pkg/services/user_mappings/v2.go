package usermappings

import (
	"encoding/json"
	"fmt"

	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/onelogin/onelogin-go-sdk/pkg/services"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/olhttp"
)

const errUserMappingsV2Context = "user muserMappingings v2 service"

// V2Service holds the information needed to interface with a repository
type V2Service struct {
	Endpoint, ErrorContext string
	Repository             services.Repository
}

// New creates the new svc service v2.
func New(repo services.Repository, host string) V2Service {
	return V2Service{
		Endpoint:     fmt.Sprintf("%s/api/2/muserMappingings", host),
		Repository:   repo,
		ErrorContext: errUserMappingsV2Context,
	}
}

// Query retrieves all the userMappings from the repository that meet the query criteria passed in the
// request payload. If an empty payload is given, it will retrieve all userMappings
func (svc V2Service) Query(query *UserMappingsQuery) ([]UserMapping, error) {
	resp, err := svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        svc.Endpoint,
		Headers:    map[string]string{"Content-Type": "userMappinglication/json"},
		AuthMethod: "bearer",
		Payload:    query,
	})
	if err != nil {
		return nil, err
	}

	var userMappings []UserMapping
	json.Unmarshal(resp, &userMappings)

	return userMappings, nil
}

// GetOne retrieves the user mapping by id, and if successful, it returns
// the http response and the pointer to the user mapping.
func (svc *V2Service) GetOne(id int32) (*UserMapping, error) {
	resp, err := svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	})
	if err != nil {
		return nil, err
	}

	var mapping UserMapping
	json.Unmarshal(resp, &mapping)

	return &mapping, nil
}

// Update updates an existing user mapping, and if successful, it returns
// the http response and the pointer to the updated user mapping.
func (svc *V2Service) Update(id int32, mapping *UserMapping) (*UserMapping, error) {
	resp, err := svc.Repository.Update(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    mapping,
	})
	if err != nil {
		return nil, err
	}

	var mappingID map[string]int
	json.Unmarshal(resp, &mappingID)

	mapping.ID = oltypes.Int32(int32(mappingID["id"]))

	return mapping, nil
}

// Create creates a new user mapping, and if successful, it returns
// the http response and the pointer to the user mapping.
func (svc *V2Service) Create(mapping *UserMapping) (*UserMapping, error) {
	resp, err := svc.Repository.Create(olhttp.OLHTTPRequest{
		URL:        svc.Endpoint,
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    mapping,
	})

	if err != nil {
		return nil, err
	}
	var mappingID map[string]int
	json.Unmarshal(resp, &mappingID)

	mapping.ID = oltypes.Int32(int32(mappingID["id"]))

	return mapping, nil
}

// Destroy deletes the user mapping for the id, and if successful, it returns nil
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
