package smarthooks

import (
	"errors"
	"github.com/onelogin/onelogin-go-sdk/pkg/utils"
	"time"
)

// SmartHookQuery represents available query parameters
type SmartHookQuery struct {
	Limit  string
	Page   string
	Cursor string
	Type   string
}

// SmartHook represents a OneLogin SmartHook with associated resource data
type SmartHook struct {
	ID              *string           `json:"id,omitempty"`
	Type            *string           `json:"type,omitempty"`
	Status          *string           `json:"status,omitempty"`
	Disabled        *bool             `json:"disabled,omitempty"`
	Retries         *int32            `json:"retries,omitempty"`
	Timeout         *int32            `json:"timeout,omitempty"`
	RiskEnabled     *bool             `json:"risk_enabled,omitempty"`
	LocationEnabled *bool             `json:"location_enabled,omitempty"`
	Packages        map[string]string `json:"packages,omitempty"`
	CreatedAt       *time.Time        `json:"created_at,omitempty"`
	UpdatedAt       *time.Time        `json:"updated_at,omitempty"`
	Function        *string           `json:"function,omitempty"`
	Options         *SmartHookOptions `json:"options,omitempty"`
	EnvVars         []EnvVar          `json:"env_vars,omitempty"`
}

// SmartHookOptions represents the options to be associated with a SmartHook
type SmartHookOptions struct {
	RiskEnabled          *bool `json:"risk_enabled,omitempty"`
	MFADeviceInfoEnabled *bool `json:"mfa_device_info_enabled,omitempty"`
	LocationEnabled      *bool `json:"location_enabled,omitempty"`
}

// EnvVar represents an Environment Variable to be associated with a SmartHook
type EnvVar struct {
	ID        *string    `json:"id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// EncodeFunction mutates the reciever to base64 encode whatever value is on the Function field
func (hook *SmartHook) EncodeFunction() error {
	if hook.Function == nil {
		return errors.New("No Function Definition Given")
	}
	str := utils.EncodeString(*hook.Function)
	hook.Function = &str
	return nil
}

// DecodeFunction mutates the reciever to base64 decode whatever value is on the Function field
func (hook *SmartHook) DecodeFunction() error {
	if hook.Function == nil {
		return errors.New("No Function Definition Given")
	}
	str := utils.DecodeString(*hook.Function)
	hook.Function = &str
	return nil
}
