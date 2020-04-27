package models

// SessionLoginTokenRequest is the contract for users api v1.
type SessionLoginTokenRequest struct {
	UsernameOrEmail *string `json:"username_or_email,omitempty"`
	Password        *string `json:"password,omitempty"`
	Subdomain       *string `json:"subdomain,omitempty"`
}
