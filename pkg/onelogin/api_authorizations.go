package onelogin

import (
	"fmt"

	mod "github.com/onelogin/onelogin-go-sdk/internal/models"
)

var (
	APIAuthPath  string = "/api/2/api_authorizations"
	APIAuthPathI string = "/api/2/api_authorizations/%d"
)

func (sdk *OneloginSDK) CreateAuthServer(authServer mod.AuthServer) ([]byte, error) {
	return sdk.ApiClient.Post(APIAuthPath, nil, authServer)
}

func (sdk *OneloginSDK) GetAuthServers(queryParams map[string]string) ([]byte, error) {
	return sdk.ApiClient.Get(APIAuthPath, &queryParams)
}

func (sdk *OneloginSDK) GetAuthServerByID(id int, queryParams map[string]string) ([]byte, error) {
	return sdk.ApiClient.Get(fmt.Sprintf(APIAuthPathI, id), &queryParams)
}

func (sdk *OneloginSDK) UpdateAuthServer(id int, authServer mod.AuthServer, queryParams map[string]string) ([]byte, error) {
	return sdk.ApiClient.Put(fmt.Sprintf(APIAuthPathI, id), &queryParams, authServer)
}

func (sdk *OneloginSDK) DeleteAuthServer(id int, queryParams map[string]string) ([]byte, error) {
	return sdk.ApiClient.Delete(fmt.Sprintf(APIAuthPathI, id), &queryParams)
}

// Claim related endpoints
func (sdk *OneloginSDK) CreateAuthServerClaim(id int, claim mod.AccessTokenClaim) ([]byte, error) {
	claimPath := fmt.Sprintf(APIAuthPathI+"/claims", id)
	return sdk.ApiClient.Post(claimPath, nil, claim)
}

func (sdk *OneloginSDK) DeleteAuthClaim(id, claimID int, queryParams map[string]string) ([]byte, error) {
	claimPath := fmt.Sprintf(APIAuthPathI+"/claims/%d", id, claimID)
	return sdk.ApiClient.Delete(claimPath, &queryParams)
}

func (sdk *OneloginSDK) GetAuthClaims(id int, queryParams map[string]string) ([]byte, error) {
	claimPath := fmt.Sprintf(APIAuthPathI+"/claims", id)
	return sdk.ApiClient.Delete(claimPath, &queryParams)
}

func (sdk *OneloginSDK) UpdateClaim(id, claimID int, claim mod.AccessTokenClaim) ([]byte, error) {
	claimPath := fmt.Sprintf(APIAuthPathI+"/claims/%d", id, claimID)
	return sdk.ApiClient.Put(claimPath, nil, claim)
}
