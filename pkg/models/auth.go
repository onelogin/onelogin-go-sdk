package models

// AuthBody is the request payload for authorization.
type AuthBody struct {
	GrantType string `json:"grant_type"`
}

// AuthResp is the authorization response payload.
type AuthResp struct {
	AccessToken  string `json:"access_token"`
	CreatedAt    string `json:"created_at"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	AccountID    int    `json:"account_id"`
}
