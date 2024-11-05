package onelogin

import (
	mod "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
	utl "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/utilities"
)

const (
	APIAuthPath string = "api/2/api_authorizations"
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

// was ListAuthServers
func (sdk *OneloginSDK) GetAuthServers(queryParams mod.Queryable) (interface{}, error) {
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

func (sdk *OneloginSDK) GetAuthServerByID(id int, queryParams mod.Queryable) (interface{}, error) {
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

func (sdk *OneloginSDK) GetAuthClaims(id int, queryParams mod.Queryable) (interface{}, error) {
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

// Scopes related endpoints
func (sdk *OneloginSDK) CreateAuthServerScope(id int, scope mod.Scope) (interface{}, error) {
	p, err := utl.BuildAPIPath(APIAuthPath, id, "scopes")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Post(&p, scope)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) DeleteAuthServerScope(id, scopeID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(APIAuthPath, id, "scopes", scopeID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Delete(&p)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetAuthServerScopes(id int, queryParams mod.Queryable) (interface{}, error) {
	p, err := utl.BuildAPIPath(APIAuthPath, id, "scopes")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, queryParams)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) UpdateAuthServerScope(id, scopeID int, scope mod.Scope) (interface{}, error) {
	p, err := utl.BuildAPIPath(APIAuthPath, id, "scopes", scopeID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, scope)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

// Client App related endpoints

func (sdk *OneloginSDK) CreateClientApp(id int, clientApp mod.ClientApp) (interface{}, error) {
	p, err := utl.BuildAPIPath(APIAuthPath, id, "clients")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Post(&p, clientApp)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetClientApps(id int) (interface{}, error) {
	p, err := utl.BuildAPIPath(APIAuthPath, id, "clients")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) DeleteClientApp(id, clientID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(APIAuthPath, id, "clients", clientID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Delete(&p)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) UpdateClientApp(id, clientID int, clientApp mod.ClientApp) (interface{}, error) {
	p, err := utl.BuildAPIPath(APIAuthPath, id, "clients", clientID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, clientApp)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}
