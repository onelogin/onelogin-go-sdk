package users

import "time"

// UserQuery represents available query parameters
type UserQuery struct {
	Limit          string
	Page           string
	Cursor         string
	CreatedSince   time.Time `json:"created_since,omitempty"`
	CreatedUntil   time.Time `json:"created_until,omitempty"`
	UpdatedSince   time.Time `json:"updated_since,omitempty"`
	UpdatedUntil   time.Time `json:"updated_until,omitempty"`
	LastLoginSince time.Time `json:"last_login_since,omitempty"`
	LastLoginUntil time.Time `json:"last_login_until,omitempty"`
	Firstname      *string   `json:"firstname,omitempty"`
	Lastname       *string   `json:"lastname,omitempty"`
	Username       *string   `json:"username,omitempty"`
	Samaccountname *string   `json:"samaccountname,omitempty"`
	DirectoryID    *string   `json:"directory_id,omitempty"`
	ExternalID     *string   `json:"external_id,omitempty"`
	AppID          *string   `json:"app_id,omitempty"`
	UserIDs        *string   `json:"user_ids,omitempty"`
	Fields         *string   `json:"fields,omitempty"`
}

// User represents a OneLogin User
type User struct {
	Firstname            *string   `json:"firstname,omitempty"`
	Lastname             *string   `json:"lastname,omitempty"`
	Username             *string   `json:"username,omitempty"`
	Email                *string   `json:"email,omitempty"`
	DistinguishedName    *string   `json:"distinguished_name,omitempty"`
	Samaccountname       *string   `json:"samaccountname,omitempty"`
	DirectoryID          *string   `json:"directory_id,omitempty"`
	ExternalID           *string   `json:"external_id,omitempty"`
	MemberOf             *string   `json:"member_of,omitempty"`
	Phone                *string   `json:"phone,omitemepty"`
	CreatedAt            time.Time `json:"created_at,omitempty"`
	UpdatedAt            time.Time `json:"updated_at,omitempty"`
	ActivatedAt          time.Time `json:"Activated_at,omitempty"`
	LastLogin            time.Time `json:"last_login,omitempty"`
	PasswordChangedAt    time.Time `json:"password_changed_at,omitempty"`
	LockedUntil          time.Time `json:"locked_until,omitempty"`
	InvitationSentAt     time.Time `json:"invitation_sent_at,omitempty"`
	State                *int32    `json:"state,omitempty"`
	Status               *int32    `json:"status,omitempty"`
	InvalidLoginAttepmts *int32    `json:"invalid_login_attempts,omitempty"`
	GroupID              *int32    `json:"group_id,omitempty"`
	ID                   *int32    `json:"id,omitempty"`
}
