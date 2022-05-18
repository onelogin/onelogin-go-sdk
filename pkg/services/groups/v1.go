package groups

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/onelogin/onelogin-go-sdk/pkg/services"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/olhttp"
)

const errGroupsV1Context = "groups v1 service"

type V1Service struct {
	Endpoint, ErrorContext string
	Repository             services.Repository
}

// New creates the new svc service v1.
func New(repo services.Repository, host string) *V1Service {
	return &V1Service{
		Endpoint:     fmt.Sprintf("%s/api/1/groups", host),
		Repository:   repo,
		ErrorContext: errGroupsV1Context,
	}
}

// Query retrieves all the groups from the repository that meet the query criteria passed in the
// request payload. If an empty payload is given, it will retrieve all roles
func (svc *V1Service) Query(query *GroupQuery) ([]Group, error) {
	resp, err := svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        svc.Endpoint,
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    query,
	})
	if err != nil {
		return nil, err
	}

	var groups []Group
	for _, bytes := range resp {
		var unmarshalled []Group
		json.Unmarshal(bytes, &unmarshalled)
		groups = append(groups, unmarshalled...)
	}
	return groups, nil
}

// GetOne retrieves the role by id and returns it
func (svc *V1Service) GetOne(id int32) (*Group, error) {
	resp, err := svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d", svc.Endpoint, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	})
	if err != nil {
		return nil, err
	}
	var group Group

	if len(resp) < 1 {
		return nil, errors.New("invalid length of response returned")
	}

	json.Unmarshal(resp[0], &group)
	return &group, nil
}
