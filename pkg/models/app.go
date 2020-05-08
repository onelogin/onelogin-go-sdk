package models

import "time"

// App is the contract for apps api v2.
type App struct {
	ID                 *int32                   `json:"id,omitempty"`
	Name               *string                  `json:"name,omitempty"`
	Visible            *bool                    `json:"visible,omitempty"`
	Description        *string                  `json:"description,,omitempty"`
	Notes              *string                  `json:"notes,omitempty"`
	IconURL            *string                  `json:"icon_url,omitempty"`
	AuthMethod         *int32                   `json:"auth_method,omitempty"`
	PolicyID           *int32                   `json:"policy_id,omitempty"`
	AllowAssumedSignin *bool                    `json:"allow_assumed_signin,omitempty"`
	TabID              *int32                   `json:"tab_id,omitempty"`
	ConnectorID        *int32                   `json:"connector_id,omitempty"`
	CreatedAt          time.Time                `json:"created_at,omitempty"`
	UpdatedAt          time.Time                `json:"updated_at,omitempty"`
	Provisioning       *AppProvisioning         `json:"provisioning"`
	Sso                *AppSso                  `json:"sso"`
	Configuration      *AppConfiguration        `json:"configuration"`
	Parameters         map[string]AppParameters `json:"parameters"`
	RoleIDs            []int                    `json:"role_ids"`
}

type AppsQuery struct {
	Limit       string
	Page        string
	Name        string
	ConnectorID string
	AuthMethod  string
	Cursor      string
}
