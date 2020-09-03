package clientapps

import "github.com/onelogin/onelogin-go-sdk/pkg/services/auth_servers/scopes"

type ClientAppsQuery struct {
	AuthServerID string
}

type ClientApp struct {
	ID           *int32         `json:"app_id,omitempty"`
	AuthServerID *int32         `json:"auth_server_id,omitempty"`
	APIAuthID    *int32         `json:"api_auth_id,omitempty"`
	Name         *string        `json:"name,omitempty"`
	Scopes       []scopes.Scope `json:"scopes,omitempty"`
	ScopeIDs     []int32        `json:"scope_ids,omitempty"`
}
