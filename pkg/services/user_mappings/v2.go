package usermappings

import (
	"fmt"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/onelogin/onelogin-go-sdk/pkg/services"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/olhttp"
	"github.com/onelogin/onelogin-go-sdk/pkg/utils"
	"log"

	"sync"
)

const errUserMappingsV2Context = "user mappings v2 service"

// V2Service holds the information needed to interface with a repository
type V2Service struct {
	Endpoint, ErrorContext string
	Repository             services.Repository
}

// New creates the new svc service v2.
func New(repo services.Repository, host string) V2Service {
	return V2Service{
		Endpoint:     fmt.Sprintf("%s/api/2/mappings", host),
		Repository:   repo,
		ErrorContext: errUserMappingsV2Context,
	}
}

// Query retrieves all the userMappings from the repository that meet the query criteria passed in the
// request payload. If an empty payload is given, it will retrieve all userMappings
func (svc *V2Service) Query(query *UserMappingsQuery) ([]UserMapping, error) {
	var (
		userMappings []UserMapping
		err          error
	)
	if err = svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        svc.Endpoint,
		Headers:    map[string]string{"Content-Type": "userMappinglication/json"},
		AuthMethod: "bearer",
		Payload:    query,
	}, &userMappings); err != nil {
		return nil, err
	}
	return userMappings, nil
}

// GetOne retrieves the user mapping by id, and if successful, it returns
// the http response and the pointer to the user mapping.
func (svc *V2Service) GetOne(id int32) (*UserMapping, error) {
	var (
		mapping UserMapping
		err     error
	)
	if err = svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	}, &mapping); err != nil {
		return nil, err
	}

	return &mapping, nil
}

// Update updates an existing user mapping, and if successful, it returns
// the http response and the pointer to the updated user mapping.
func (svc *V2Service) Update(id int32, mapping *UserMapping) (*UserMapping, error) {
	validationErr := svc.ValidateMappingValues(mapping)
	if validationErr != nil {
		return nil, validationErr
	}
	var (
		mappingID map[string]int32
		err       error
	)
	if err = svc.Repository.Update(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    mapping,
	}, &mappingID); err != nil {
		return nil, err
	}
	mapping.ID = oltypes.Int32(int32(mappingID["id"]))
	return mapping, nil
}

// Create creates a new user mapping, and if successful, it returns
// the http response and the pointer to the user mapping.
func (svc *V2Service) Create(mapping *UserMapping) (*UserMapping, error) {
	validationErr := svc.ValidateMappingValues(mapping)
	if validationErr != nil {
		return nil, validationErr
	}
	var (
		mappingID map[string]int32
		err       error
	)
	if err = svc.Repository.Create(olhttp.OLHTTPRequest{
		URL:        svc.Endpoint,
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    mapping,
	}, &mappingID); err != nil {
		return nil, err
	}
	mapping.ID = oltypes.Int32(int32(mappingID["id"]))
	return mapping, nil
}

// Destroy deletes the user mapping for the id, and if successful, it returns nil
func (svc *V2Service) Destroy(id int32) error {
	if err := svc.Repository.Destroy(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	}, nil); err != nil {
		return err
	}
	return nil
}

func (svc *V2Service) ValidateMappingValues(mapping *UserMapping) error {
	legalValRequests := map[string][]string{}
	legalValRequests["mappings/conditions"] = []string{}
	legalValRequests["mappings/actions"] = []string{}
	for _, condition := range mapping.Conditions {
		legalValRequests[fmt.Sprintf("mappings/conditions/%s/values", *condition.Source)] = []string{}
		legalValRequests[fmt.Sprintf("mappings/conditions/%s/operators", *condition.Source)] = []string{}
	}
	for _, action := range mapping.Actions {
		legalValRequests[fmt.Sprintf("mappings/actions/%s/values", *action.Action)] = []string{}
	}

	var wg sync.WaitGroup
	for reqURL := range legalValRequests {
		wg.Add(1)
		go func(reqURL string, legalValRequest map[string][]string) {
			defer wg.Done()
			var (
				legalValResp []map[string]string
				err          error
			)
			if err = svc.Repository.Read(olhttp.OLHTTPRequest{
				AuthMethod: "bearer",
				URL:        fmt.Sprintf("%s/%s", svc.Endpoint, reqURL),
				Headers:    map[string]string{"Content-Type": "application/json"},
			}, &legalValResp); err != nil {
				log.Println("Problem validating mapping", reqURL, err)
			}
			legalVals := make([]string, len(legalValResp))
			for i, legalVal := range legalValResp {
				legalVals[i] = legalVal["value"]
			}
			legalValRequests[reqURL] = legalVals
		}(reqURL, legalValRequests)
	}
	wg.Wait()

	var validationErr error
	for _, condition := range mapping.Conditions {
		if len(legalValRequests["mappings/conditions"]) > 0 {
			err := utils.OneOf(*mapping.Name, *condition.Source, legalValRequests["mappings/conditions"])
			if err != nil {
				log.Println("Illegal value given for condition source")
				validationErr = err
			}
		}
		if len(legalValRequests[fmt.Sprintf("mappings/conditions/%s/values", *condition.Source)]) > 0 {
			err := utils.OneOf(*mapping.Name, *condition.Value, legalValRequests[fmt.Sprintf("mappings/conditions/%s/values", *condition.Source)])
			if err != nil {
				log.Println("Illegal value given for condition value")
				validationErr = err
			}
		}
		if len(legalValRequests[fmt.Sprintf("mappings/conditions/%s/operators", *condition.Source)]) > 0 {
			err := utils.OneOf(*mapping.Name, *condition.Operator, legalValRequests[fmt.Sprintf("mappings/conditions/%s/operators", *condition.Source)])
			if err != nil {
				log.Println("Illegal value given for condition operator")
				validationErr = err
			}
		}
	}

	for _, action := range mapping.Actions {
		if len(legalValRequests["mappings/actions"]) > 0 {
			err := utils.OneOf(*mapping.Name, *action.Action, legalValRequests["mappings/actions"])
			if err != nil {
				log.Println("Illegal value given for action")
				validationErr = err
			}
		}
		for _, val := range action.Value {
			if len(legalValRequests[fmt.Sprintf("mappings/actions/%s/values", *action.Action)]) > 0 {
				err := utils.OneOf(*mapping.Name, val, legalValRequests[fmt.Sprintf("mappings/actions/%s/values", *action.Action)])
				if err != nil {
					log.Println("Illegal value given for action value")
					validationErr = err
				}
			}
		}
	}
	return validationErr
}
