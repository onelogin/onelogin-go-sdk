package smarthooks

import "time"

// SmartHookQuery represents available query parameters
type SmartHookQuery struct {
	Limit  string
	Page   string
	Cursor string
	Type   string
}

// SmartHook represents a OneLogin SmartHook
type SmartHook struct {
	ID              *string           `json:"id,omitempty"`
	Type            *string           `json:"type,omitempty"`
	Status          *string           `json:"status,omitempty"`
	Disabled        *bool             `json:"disabled,omitempty"`
	Retries         *int32            `json:"retries,omitempty"`
	Timeout         *int32            `json:"timeout,omitempty"`
	RiskEnabled     *bool             `json:"risk_enabled,omitempty"`
	LocationEnabled *bool             `json:"location_enabled,omitempty"`
	EnvVars         []string          `json:"env_vars,omitempty"`
	Packages        map[string]string `json:"packages,omitempty"`
	CreatedAt       *time.Time        `json:"created_at,omitempty"`
	UpdatedAt       *time.Time        `json:"updated_at,omitempty"`
	Function        *string           `json:"function,omitempty"` // this is the field that gets sent to the API as function
}
