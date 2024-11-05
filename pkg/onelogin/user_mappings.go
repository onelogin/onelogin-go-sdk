package onelogin

import (
	mod "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
	utl "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/utilities"
)

const (
	MappingsPath string = "api/2/mappings"
)

func (sdk *OneloginSDK) ListMappings() (interface{}, error) {
	p, err := utl.BuildAPIPath(MappingsPath)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) CreateMapping(mapping mod.UserMapping) (interface{}, error) {
	p, err := utl.BuildAPIPath(MappingsPath)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Post(&p, mapping)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) DeleteMapping(mappingID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(MappingsPath, mappingID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Delete(&p)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetMapping(mappingID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(MappingsPath, mappingID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) ListActions() (interface{}, error) {
	p, err := utl.BuildAPIPath(MappingsPath, "actions")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) UpdateMapping(mappingID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(MappingsPath, mappingID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) BulkSortMappings(mappingIDs []int) (interface{}, error) {
	p, err := utl.BuildAPIPath(MappingsPath, "bulk_sort")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, mappingIDs)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) ListActionValues(actionValue string) (interface{}, error) {
	p, err := utl.BuildAPIPath(MappingsPath, "actions", actionValue, "values")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) ListConditionValues(conditionValue string) (interface{}, error) {
	p, err := utl.BuildAPIPath(MappingsPath, "conditions", conditionValue, "values")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) ListConditionOperators(conditionValue string) (interface{}, error) {
	p, err := utl.BuildAPIPath(MappingsPath, "conditions", conditionValue, "operators")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) DryrunMapping(mappingID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(MappingsPath, mappingID, "dryrun")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Post(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

// https://<subdomain>/api/2/mappings/conditions
func (sdk *OneloginSDK) ListConditions() (interface{}, error) {
	p, err := utl.BuildAPIPath(MappingsPath, "conditions")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}
