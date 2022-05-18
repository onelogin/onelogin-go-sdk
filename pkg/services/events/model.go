package events

import "time"

// EventQuery represents available query parameters
type EventQuery struct {
	Limit       string
	Page        string
	Cursor      string
	ClientID    *int32     `json:"client_id,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	DirectoryID *int32     `json:"directory_id,omitempty"`
	EventTypeID *int32     `json:"event_type_id,omitempty"`
	ID          *int32     `json:"id,omitempty"`
	Resolution  *string    `json:"resolution,omitempty"`
	Since       *time.Time `json:"since,omitempty"`
	Until       *time.Time `json:"until,omitempty"`
	UserID      *int32     `json:"user_id,omitempty"`
}

// EventType represents the EventType resource in OneLogin
type EventType struct {
	ID          *int32  `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description,omitempty"`
}

// Event represents the Event resource in OneLogin
type Event struct {
	AccountID            *int32     `json:"account_id,omitempty"`
	ActorSystem          *string    `json:"actor_system,omitempty"`
	ActorUserID          *int32     `json:"actor_user_id,omitempty"`
	ActorUserName        *string    `json:"actor_user_name,omitempty"`
	AppID                *int32     `json:"app_id,omitempty"`
	AppName              *string    `json:"app_name,omitempty"`
	AssumingActingUserID *int32     `json:"assuming_acting_user_id,omitempty"`
	BrowserFingerprint   *string    `json:"browser_fingerprint,omitempty"`
	ClientID             *int32     `json:"client_id,omitempty"`
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	CustomMessage        *string    `json:"custom_message,omitempty"`
	DirectoryID          *int32     `json:"directory_id,omitempty"`
	DirectorySyncRunID   *int32     `json:"directory_sync_run_id,omitempty"`
	ErrorDescription     *string    `json:"error_description,omitempty"`
	EventTypeID          *int32     `json:"event_type_id,omitempty"`
	EventTypeIds         *int32     `json:"event_type_ids,omitempty"`
	GroupID              *int32     `json:"group_id,omitempty"`
	GroupName            *string    `json:"group_name,omitempty"`
	ID                   *int32     `json:"id,omitempty"`
	Ipaddr               *string    `json:"ipaddr,omitempty"`
	Notes                *string    `json:"notes,omitempty"`
	OperationName        *string    `json:"operation_name,omitempty"`
	OtpDeviceID          *int32     `json:"otp_device_id,omitempty"`
	OtpDeviceName        *string    `json:"otp_device_name,omitempty"`
	PolicyID             *int32     `json:"policy_id,omitempty"`
	PolicyName           *string    `json:"policy_name,omitempty"`
	ProxyIP              *string    `json:"proxy_ip,omitempty"`
	Resolution           *string    `json:"resolution,omitempty"`
	ResourceTypeID       *int32     `json:"resource_type_id,omitempty"`
	RiskCookieID         *string    `json:"risk_cookie_id,omitempty"`
	RiskReasons          *string    `json:"risk_reasons,omitempty"`
	RiskScore            *int32     `json:"risk_score,omitempty"`
	RoleID               *int32     `json:"role_id,omitempty"`
	RoleName             *string    `json:"role_name,omitempty"`
	Since                *time.Time `json:"since,omitempty"`
	Until                *time.Time `json:"until,omitempty"`
	UserID               *int32     `json:"user_id,omitempty"`
	UserName             *string    `json:"user_name,omitempty"`
}
