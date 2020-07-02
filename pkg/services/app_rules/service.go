package apprules

import (
	"encoding/json"
	"fmt"

	"github.com/onelogin/onelogin-go-sdk/pkg/services"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/olhttp"
)

const errAppsV2Context = "app rules v2 service"

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

func (svc *V2Service) Query(query *AppRuleQuery) ([]AppRule, error) {
	resp, err := svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%s/rules", svc.Endpoint, query.AppID),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	})
	if err != nil {
		return nil, err
	}

	var appRules []AppRule
	json.Unmarshal(resp, &appRules)
	return appRules, nil
}

func (svc *V2Service) GetOne(appId int32, id int32) (*AppRule, error) {
	resp, err := svc.Repository.Read(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d/rules/%d", svc.Endpoint, appId, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	})
	if err != nil {
		return nil, err
	}

	var appRule AppRule
	json.Unmarshal(resp, &appRule)
	return &appRule, nil
}

func (svc *V2Service) Create(appRule *AppRule) (*AppRule, error) {
	var newAppRule AppRule
	resp, err := svc.Repository.Create(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf(""),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    appRule,
	})
	if err != nil {
		return &newAppRule, err
	}
	json.Unmarshal(resp, &newAppRule)
	return &newAppRule, nil
}

func (svc *V2Service) Update(appRule *AppRule) (*AppRule, error) {
	var updatedAppRule AppRule
	resp, err := svc.Repository.Update(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d/rules/%d", svc.Endpoint, *appRule.AppID, *appRule.ID),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
		Payload:    appRule,
	})
	if err != nil {
		return &updatedAppRule, err
	}
	json.Unmarshal(resp, &updatedAppRule)
	return &updatedAppRule, nil
}

func (svc *V2Service) Destroy(appId int32, id int32) error {
	if _, err := svc.Repository.Destroy(olhttp.OLHTTPRequest{
		URL:        fmt.Sprintf("%s/%d/rules/%d", svc.Endpoint, appId, id),
		Headers:    map[string]string{"Content-Type": "application/json"},
		AuthMethod: "bearer",
	}); err != nil {
		return err
	}
	return nil
}
