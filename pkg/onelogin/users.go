package onelogin

import (
	"encoding/json"
	"errors"

	mod "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
	utl "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/utilities"
)

const (
	UserPathV1 string = "api/1/users"
	UserPathV2 string = "api/2/users"
)

// Users V2
// was ListUsers
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

func (sdk *OneloginSDK) CreateUser(user mod.User) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV2)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Post(&p, user)
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

func (sdk *OneloginSDK) GetCustomAttributes() (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV2, "custom_attributes")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetCustomAttributeByID(id int) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV2, "custom_attributes", id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) CreateCustomAttributes(requestBody interface{}) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV2, "custom_attributes")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Post(&p, requestBody)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) UpdateCustomAttributes(id int, requestBody interface{}) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV2, "custom_attributes", id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, requestBody)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) DeleteCustomAttributes(id int) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV2, "custom_attributes", id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Delete(&p)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetUsersModels(query mod.Queryable) ([]mod.User, error) {
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

	tmp, err := utl.CheckHTTPResponse(resp)
	if err != nil {
		return nil, err
	}

	var users []mod.User
	tmpBytes, err := json.Marshal(tmp)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(tmpBytes, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Users V1

func (sdk *OneloginSDK) GetUserRoles(id int) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV1, id, "roles")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) AddRolesForUser(userID int, requestBody interface{}) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV1, userID, "add_roles")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, requestBody)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) RemoveRolesForUser(userID int, requestBody interface{}) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV1, userID, "remove_roles")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, requestBody)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) UpdatePasswordInsecure(id int, requestBody interface{}) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV1, "set_password_clear_text", id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, requestBody)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) UpdatePasswordSecure(id int, requestBody interface{}) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV1, "set_password_using_salt", id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, requestBody)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) SetCustomAttributes(userID int, requestBody interface{}) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV1, userID, "set_custom_attributes")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, requestBody)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) SetUserState(userID, requestBody interface{}) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV1, userID, "set_state")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, requestBody)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) LogOutUser(userID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV1, userID, "logout")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) LockUserAccount(id int, requestBody interface{}) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV1, id, "lock_user")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, requestBody)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}
