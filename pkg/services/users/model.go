package users

import "time"

const (
	USER_STATE_UNAPPROVED int32 = iota
	USER_STATE_APPROVED
	USER_STATE_REJECTED
	USER_STATE_UNLICENSED
)

const (
	USER_STATUS_UNACTIVATED int32 = iota
	USER_STATUS_ACTIVE            // Only users assigned this status can log in to OneLogin.
	USER_STATUS_SUSPENDED
	USER_STATUS_LOCKED
	USER_STATUS_PASSWORD_EXPIRED
	USER_STATUS_AWAITING_PASSWORD_RESET     // The user is required to reset their password.
	unused_user_status_6                    // There is not user status with a value of 6.
	USER_STATUS_PASSWORD_PENDING            // The user has not yet set their password.
	USER_STATUS_SECURITY_QUESTIONS_REQUIRED // The user has not yet set their security questions.
)

// UserQuery represents available query parameters
type UserQuery struct {
	Limit          string
	Page           string
	Cursor         string
	CreatedSince   *time.Time `json:"created_since,omitempty"`
	CreatedUntil   *time.Time `json:"created_until,omitempty"`
	UpdatedSince   *time.Time `json:"updated_since,omitempty"`
	UpdatedUntil   *time.Time `json:"updated_until,omitempty"`
	LastLoginSince *time.Time `json:"last_login_since,omitempty"`
	LastLoginUntil *time.Time `json:"last_login_until,omitempty"`
	Firstname      *string    `json:"firstname,omitempty"`
	Lastname       *string    `json:"lastname,omitempty"`
	Email          *string    `json:"email,omitempty"`
	Username       *string    `json:"username,omitempty"`
	Samaccountname *string    `json:"samaccountname,omitempty"`
	DirectoryID    *string    `json:"directory_id,omitempty"`
	ExternalID     *string    `json:"external_id,omitempty"`
	AppID          *string    `json:"app_id,omitempty"`
	UserIDs        *string    `json:"user_ids,omitempty"`
	Fields         *string    `json:"fields,omitempty"`
}

// User represents a OneLogin User
type User struct {
	Firstname            *string                `json:"firstname,omitempty"`
	Lastname             *string                `json:"lastname,omitempty"`
	Username             *string                `json:"username,omitempty"`
	Email                *string                `json:"email,omitempty"`
	DistinguishedName    *string                `json:"distinguished_name,omitempty"`
	Samaccountname       *string                `json:"samaccountname,omitempty"`
	UserPrincipalName    *string                `json:"userprincipalname,omitempty"`
	MemberOf             *string                `json:"member_of,omitempty"`
	Phone                *string                `json:"phone,omitempty"`
	Password             *string                `json:"password,omitempty"`
	PasswordConfirmation *string                `json:"password_confirmation,omitempty"`
	PasswordAlgorithm    *string                `json:"password_algorithm,omitempty"`
	Salt                 *string                `json:"salt,omitempty"`
	Title                *string                `json:"title,omitempty"`
	Company              *string                `json:"company,omitempty"`
	Department           *string                `json:"department,omitempty"`
	Comment              *string                `json:"comment,omitempty"`
	CreatedAt            *time.Time             `json:"created_at,omitempty"`
	UpdatedAt            *time.Time             `json:"updated_at,omitempty"`
	ActivatedAt          *time.Time             `json:"activated_at,omitempty"`
	LastLogin            *time.Time             `json:"last_login,omitempty"`
	PasswordChangedAt    *time.Time             `json:"password_changed_at,omitempty"`
	LockedUntil          *time.Time             `json:"locked_until,omitempty"`
	InvitationSentAt     *time.Time             `json:"invitation_sent_at,omitempty"`
	State                *int32                 `json:"state,omitempty"`
	Status               *int32                 `json:"status,omitempty"`
	InvalidLoginAttempts *int32                 `json:"invalid_login_attempts,omitempty"`
	GroupID              *int32                 `json:"group_id,omitempty"`
	DirectoryID          *int32                 `json:"directory_id,omitempty"`
	TrustedIDPID         *int32                 `json:"trusted_idp_id,omitempty"`
	ManagerADID          *int32                 `json:"manager_ad_id,omitempty"`
	ManagerUserID        *int32                 `json:"manager_user_id,omitempty"`
	ExternalID           *int32                 `json:"external_id,omitempty"`
	ID                   *int32                 `json:"id,omitempty"`
	CustomAttributes     map[string]interface{} `json:"custom_attributes,omitempty"`
}
