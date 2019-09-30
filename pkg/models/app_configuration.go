package models

type AppConfiguration struct {
	RedirectUri                   string `json:"redirect_uri"`
	RefreshTokenExpirationMinutes int32  `json:"refresh_token_expiration_minutes"`
	LoginUrl                      string `json:"login_url"`
	OidcApplicationType           int32  `json:"oidc_application_type"`
	TokenEndpointAuthMethod       int32  `json:"token_endpoint_auth_method"`
	AccessTokenExpirationMinutes  int32  `json:"access_token_expiration_minutes"`
	ProviderArn                   string `json:"provider_arn"`
	SignatureAlgorithm            string `json:"signature_algorithm"`
}
