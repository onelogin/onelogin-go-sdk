package onelogin

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
	mod "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
	utl "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/utilities"
)

const (
	UserPathV1 string = "api/1/users"
	UserPathV2 string = "api/2/users"
)

// Users V2
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
func (sdk *OneloginSDK) GetUsersModels(query mod.Queryable) ([]models.User, error) {
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

	tmp,err := utl.CheckHTTPResponse(resp)
	if err != nil {
		return nil, err
	}

	var users []models.User
	tmpBytes,err := json.Marshal(tmp)
	if err != nil {
		return nil, err
	}
	err= json.Unmarshal(tmpBytes,&users)
	if err != nil {
		return nil, err
	}
	return users,nil
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

// Users V1
func (sdk *OneloginSDK) UpdatePasswordSecure(id int) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV1, "set_password_using_salt", id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) UpdatePasswordInsecure(id int) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV1, "set_password_clear", id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) LockUserAccount(id int) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV1, id, "lock_user")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

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

func (sdk *OneloginSDK) AssignRolesToUser(userID int, roles []int) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV1, userID, "add_roles")
	if err != nil {
		return nil, err
	}
	payload := map[string][]int{"role_id_array": roles}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) SetUserState(userID, state int) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV1, userID, "set_state")
	if err != nil {
		return nil, err
	}
	payload := map[string]int{"state": state}
	resp, err := sdk.Client.Put(&p, payload)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) RemoveUserRole(userID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV1, userID, "remove_roles")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetCustomAttributes() (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV1, "custom_attributes")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) SetCustomAttributes(userID int, attr interface{}) (interface{}, error) {
	p, err := utl.BuildAPIPath(UserPathV1, userID, "set_custom_attributes")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, attr)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}
