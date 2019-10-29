package models

// AppConfiguration is the contract for configuration.
type AppConfiguration struct {
	RedirectURI                   string `json:"redirect_uri,omitempty"`
	RefreshTokenExpirationMinutes int32  `json:"refresh_token_expiration_minutes,omitempty"`
	LoginURL                      string `json:"login_url,omitempty"`
	OidcApplicationType           int32  `json:"oidc_application_type,omitempty"`
	TokenEndpointAuthMethod       int32  `json:"token_endpoint_auth_method,omitempty"`
	AccessTokenExpirationMinutes  int32  `json:"access_token_expiration_minutes,omitempty"`
	ProviderArn                   string `json:"provider_arn,omitempty"`
	SignatureAlgorithm            string `json:"signature_algorithm,omitempty"`
}
