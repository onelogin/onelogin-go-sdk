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
	Function        *string           `json:"function,omitempty"`
}

// InflatedSmartHook represents a OneLogin SmartHook with associated resource data
type InflatedSmartHook struct {
	ID              *string           `json:"id,omitempty"`
	Type            *string           `json:"type,omitempty"`
	Status          *string           `json:"status,omitempty"`
	Disabled        *bool             `json:"disabled,omitempty"`
	Retries         *int32            `json:"retries,omitempty"`
	Timeout         *int32            `json:"timeout,omitempty"`
	RiskEnabled     *bool             `json:"risk_enabled,omitempty"`
	LocationEnabled *bool             `json:"location_enabled,omitempty"`
	EnvVars         []EnvVar          `json:"env_vars,omitempty"`
	Packages        map[string]string `json:"packages,omitempty"`
	CreatedAt       *time.Time        `json:"created_at,omitempty"`
	UpdatedAt       *time.Time        `json:"updated_at,omitempty"`
	Function        *string           `json:"function,omitempty"`
}

// EnvVar represents an Environment Variable to be associated with a SmartHook
type EnvVar struct {
	ID        *string    `json:"id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type HasEncodable interface {
	GetEncodableField() *string
	SetEncodableField(s *string)
}

func (hook *SmartHook) GetEncodableField() *string {
	return hook.Function
}

func (hook *SmartHook) SetEncodableField(s *string) {
	hook.Function = s
}

func (hook *InflatedSmartHook) GetEncodableField() *string {
	return hook.Function
}

func (hook *InflatedSmartHook) SetEncodableField(s *string) {
	hook.Function = s
}
