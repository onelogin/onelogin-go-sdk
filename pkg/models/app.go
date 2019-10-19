package models

import "time"

// App is the contract for apps api v2.
type App struct {
	ID                 int32                    `json:"id,omitempty"`
	Name               string                   `json:"name,omitempty"`
	Visible            bool                     `json:"visible"`
	Description        string                   `json:"description,omitempty"`
	Notes              string                   `json:"notes,omitempty"`
	IconURL            string                   `json:"icon_url,omitempty"`
	AuthMethod         int32                    `json:"auth_method,omitempty"`
	PolicyID           int32                    `json:"policy_id,omitempty"`
	AllowAssumedSignin bool                     `json:"allow_assumed_signin"`
	TabID              int32                    `json:"tab_id,omitempty"`
	ConnectorID        int32                    `json:"connector_id,omitempty"`
	CreatedAt          time.Time                `json:"created_at,omitempty"`
	UpdatedAt          time.Time                `json:"updated_at,omitempty"`
	Provisioning       *AppProvisioning         `json:"provisioning,omitempty"`
	Sso                *AppSso                  `json:"sso,omitempty"`
	Configuration      *AppConfiguration        `json:"configuration,omitempty"`
	Parameters         map[string]AppParameters `json:"parameters,omitempty"`
	RoleIDs            []int                    `json:"role_ids,omitempty"`
}
