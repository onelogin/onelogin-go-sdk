package onelogin

import (
	"fmt"

	mod "github.com/onelogin/onelogin-go-sdk/internal/models"
)

var (
	RolePath  string = "/api/2/roles"
	RolePathI string = "/api/2/roles/%d"
)

func (sdk *OneloginSDK) CreateRole(role mod.Role) ([]byte, error) {
	return sdk.ApiClient.Post(RolePath, nil, role)
}

func (sdk *OneloginSDK) GetRoles(queryParams map[string]string) ([]byte, error) {
	return sdk.ApiClient.Get(RolePath, queryParams)
}

func (sdk *OneloginSDK) GetRoleByID(id int, queryParams map[string]string) ([]byte, error) {
	return sdk.ApiClient.Get(fmt.Sprintf(RolePathI, id), queryParams)
}

func (sdk *OneloginSDK) UpdateRole(id int, role mod.Role, queryParams map[string]string) ([]byte, error) {
	return sdk.ApiClient.Put(fmt.Sprintf(RolePathI, id), queryParams, role)
}

func (sdk *OneloginSDK) DeleteRole(id int, queryParams map[string]string) ([]byte, error) {
	return sdk.ApiClient.Delete(fmt.Sprintf(RolePathI, id), queryParams)
}
