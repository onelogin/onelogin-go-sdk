package onelogin

import (
	mod "github.com/onelogin/onelogin-go-sdk/internal/models"
	utl "github.com/onelogin/onelogin-go-sdk/internal/utilities"
)

const (
	UserPath string = "/api/2/users"
)

func (sdk *OneloginSDK) CreateUser(user *mod.User) (interface{}, error) {
	p := UserPath
	resp, err := sdk.Client.Post(&p, user)
	if err != nil {
		return nil, err
	}

	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetUsers(query mod.Queryable) (interface{}, error) {
	p := UserPath

	// Validate query parameters
	//validators := query.GetKeyValidators()
	//if !utl.ValidateQueryParams(query, validators) {
	//	return nil, errors.New("invalid query parameters")
	//}

	resp, err := sdk.Client.Get(&p, &query)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetUserByID(id int, queryParams *mod.Queryable) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPath, id)
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
	p, err := utl.BuildAPIPath(UserPath, id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, user)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}
