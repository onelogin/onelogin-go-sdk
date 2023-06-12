package onelogin

import (
	"errors"

	mod "github.com/onelogin/onelogin-go-sdk/internal/models"
	utl "github.com/onelogin/onelogin-go-sdk/internal/utilities"
)

const (
	UserPathV1 string = "api/1/users"
	UserPathV2 string = "api/2/users"
)

func (sdk *OneloginSDK) CreateUser(user *mod.User) (interface{}, error) {
	p := UserPathV2
	resp, err := sdk.Client.Post(&p, user)
	if err != nil {
		return nil, err
	}

	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetUsers(query mod.Queryable) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV2)
	if err != nil {
		return nil, err
	}

	// Validate query parameters
	validators := query.GetKeyValidators()
	if !utl.ValidateQueryParams(query, validators) {
		return nil, errors.New("invalid query parameters")
	}

	resp, err := sdk.Client.Get(&p, query)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetUserByID(id int, queryParams mod.Queryable) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV2, id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, queryParams)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetUserApps(id int, queryParams mod.Queryable) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV2, id, "apps")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, queryParams)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) UpdateUser(id int, user mod.User) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV2, id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, user)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) DeleteUser(id int) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV2, id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Delete(&p)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) UpdatePasswordSecure(userID string)
