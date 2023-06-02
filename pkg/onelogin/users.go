package onelogin

import (
	"fmt"

	mod "github.com/onelogin/onelogin-go-sdk/internal/models"
)

const (
	UserPath  string = "/api/2/users"
	UserPathI string = "/api/2/users/%d"
)

func (sdk *OneloginSDK) CreateUser(user mod.User) ([]byte, error) {
	return sdk.ApiClient.Post(UserPath, nil, user)
}

func (sdk *OneloginSDK) GetUsers(queryParams map[string]string) ([]byte, error) {
	return sdk.ApiClient.Get(UserPath, queryParams)
}

func (sdk *OneloginSDK) GetUserByID(id int, queryParams map[string]string) ([]byte, error) {
	return sdk.ApiClient.Get(fmt.Sprintf(UserPathI, id), queryParams)
}

func (sdk *OneloginSDK) UpdateUser(id int, user mod.User, queryParams map[string]string) ([]byte, error) {
	return sdk.ApiClient.Put(fmt.Sprintf(UserPathI, id), queryParams, user)
}
