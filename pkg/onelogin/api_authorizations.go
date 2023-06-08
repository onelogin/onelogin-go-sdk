package onelogin

import (
	mod "github.com/onelogin/onelogin-go-sdk/internal/models"
	utl "github.com/onelogin/onelogin-go-sdk/internal/utilities"
)

const (
	APIAuthPath string = "/api/2/api_authorizations"
)

func (sdk *OneloginSDK) CreateAuthServer(authServer *mod.AuthServer) (interface{}, error) {
	p, err := utl.BuildAPIPath(APIAuthPath)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Post(&p, authServer)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetAuthServers(queryParams *mod.Queryable) (interface{}, error) {
	p, err := utl.BuildAPIPath(APIAuthPath)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, queryParams)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetAuthServerByID(id int, queryParams *mod.Queryable) (interface{}, error) {
	p, err := utl.BuildAPIPath(APIAuthPath)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, queryParams)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) UpdateAuthServer(id int, authServer mod.AuthServer) (interface{}, error) {
	p, err := utl.BuildAPIPath(APIAuthPath)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, authServer)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) DeleteAuthServer(id int) (interface{}, error) {
	p, err := utl.BuildAPIPath(APIAuthPath, id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Delete(&p)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

// Claim related endpoints
func (sdk *OneloginSDK) CreateAuthServerClaim(id int, claim mod.AccessTokenClaim) (interface{}, error) {
	p, err := utl.BuildAPIPath(APIAuthPath, id, "claims")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Post(&p, claim)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) DeleteAuthClaim(id, claimID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(APIAuthPath, id, "claims", claimID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Delete(&p)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetAuthClaims(id int, queryParams *mod.Queryable) (interface{}, error) {
	p, err := utl.BuildAPIPath(APIAuthPath, id, "claims")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, queryParams)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) UpdateClaim(id, claimID int, claim mod.AccessTokenClaim) (interface{}, error) {
	p, err := utl.BuildAPIPath(APIAuthPath, id, "claims", claimID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, claim)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}
