package onelogin

import (
	mod "github.com/onelogin/onelogin-go-sdk/internal/models"
	utl "github.com/onelogin/onelogin-go-sdk/internal/utilities"
)

var (
	RolePath string = "api/2/roles"
)

func (sdk *OneloginSDK) CreateRole(role *mod.Role) (interface{}, error) {
	p := RolePath
	resp, err := sdk.Client.Post(&p, role)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetRoles(queryParams mod.Queryable) (interface{}, error) {
	p := RolePath
	resp, err := sdk.Client.Get(&p, queryParams)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetRoleByID(id int, queryParams mod.Queryable) (interface{}, error) {
	p, err := utl.BuildAPIPath(RolePath, id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, queryParams)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) UpdateRole(id int, role mod.Role, queryParams map[string]string) (interface{}, error) {
	p, err := utl.BuildAPIPath(RolePath, id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, role)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) DeleteRole(id int, queryParams map[string]string) (interface{}, error) {
	p, err := utl.BuildAPIPath(RolePath, id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Delete(&p)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}
