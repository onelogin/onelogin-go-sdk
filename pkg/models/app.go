package models

import "time"

type App struct {
	Id                 int32                    `json:"id,omitempty"`
	Name               string                   `json:"name"`
	Visible            bool                     `json:"visible"`
	Description        string                   `json:"description"`
	Notes              string                   `json:"notes"`
	IconUrl            string                   `json:"icon_url"`
	AuthMethod         int32                    `json:"auth_method"`
	PolicyId           int32                    `json:"policy_id"`
	AllowAssumedSignin bool                     `json:"allow_assumed_signin"`
	TabId              int32                    `json:"tab_id"`
	ConnectorId        int32                    `json:"connector_id"`
	CreatedAt          time.Time                `json:"created_at"`
	UpdatedAt          time.Time                `json:"updated_at"`
	Provisioning       *AppProvisioning         `json:"provisioning"`
	Sso                *AppSso                  `json:"sso"`
	Configuration      *AppConfiguration        `json:"configuration"`
	Parameters         map[string]AppParameters `json:"parameters"`
}
