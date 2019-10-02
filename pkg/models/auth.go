package models

type AuthBody struct {
	GrantType string `json:"grant_type"`
}

type AuthResp struct {
	AccessToken  string `json:"access_token"`
	CreatedAt    string `json:"created_at"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	AccountId    int    `json:"account_id"`
}
