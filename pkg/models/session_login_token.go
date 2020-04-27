package models

import "time"

// SessionLoginToken is the contract for users api v1.
type SessionLoginToken struct {
	ExpiresAt    time.Time `json:"expires_at,omitempty"`
	ReturnToURL  *string   `json:"return_to_url,omitempty"`
	SessionToken *string   `json:"session_token,omitempty"`
	StateToken   *string   `json:"state_token,omitempty"`
}
