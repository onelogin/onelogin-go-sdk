package apps

import "time"

type AppsQuery struct {
	Limit       string
	Page        string
	Name        string
	ConnectorID string
	AuthMethod  string
	Cursor      string
}

// App is the contract for apps api v2.
type App struct {
	ID                 *int32                   `json:"id,omitempty"`
	Name               *string                  `json:"name,omitempty"`
	Visible            *bool                    `json:"visible,omitempty"`
	Description        *string                  `json:"description,,omitempty"`
	Notes              *string                  `json:"notes,omitempty"`
	IconURL            *string                  `json:"icon_url,omitempty"`
	AuthMethod         *int32                   `json:"auth_method,omitempty"`
	PolicyID           *int32                   `json:"policy_id,omitempty"`
	AllowAssumedSignin *bool                    `json:"allow_assumed_signin,omitempty"`
	TabID              *int32                   `json:"tab_id,omitempty"`
	ConnectorID        *int32                   `json:"connector_id,omitempty"`
	CreatedAt          time.Time                `json:"created_at,omitempty"`
	UpdatedAt          time.Time                `json:"updated_at,omitempty"`
	Provisioning       *AppProvisioning         `json:"provisioning"`
	Sso                *AppSso                  `json:"sso"`
	Configuration      *AppConfiguration        `json:"configuration"`
	Parameters         map[string]AppParameters `json:"parameters"`
	RoleIDs            []int                    `json:"role_ids"`
}

// AppProvisioning is the contract for provisioning.
type AppProvisioning struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// AppSso is the contract for apps sso.
type AppSso struct {
	ClientID     *string            `json:"client_id,omitempty"`
	ClientSecret *string            `json:"client_secret,omitempty"`
	MetadataURL  *string            `json:"metadata_url,omitempty"`
	AcsURL       *string            `json:"acs_url,omitempty"`
	SlsURL       *string            `json:"sls_url,omitempty"`
	Issuer       *string            `json:"issuer,omitempty"`
	Certificate  *AppSsoCertificate `json:"certificate"`
}

// AppSsoCertificate is the contract for sso certificate.
type AppSsoCertificate struct {
	ID    *int32  `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// AppConfiguration is the contract for configuration.
type AppConfiguration struct {
	RedirectURI                   *string `json:"redirect_uri,omitempty"`
	RefreshTokenExpirationMinutes *int32  `json:"refresh_token_expiration_minutes,omitempty"`
	LoginURL                      *string `json:"login_url,omitempty"`
	OidcApplicationType           *int32  `json:"oidc_application_type,omitempty"`
	TokenEndpointAuthMethod       *int32  `json:"token_endpoint_auth_method,omitempty"`
	AccessTokenExpirationMinutes  *int32  `json:"access_token_expiration_minutes,omitempty"`
	ProviderArn                   *string `json:"provider_arn,omitempty"`
	SignatureAlgorithm            *string `json:"signature_algorithm,omitempty"`
}

// AppParameters is the contract for parameters.
type AppParameters struct {
	ID                        *int32  `json:"id,omitempty"`
	Label                     *string `json:"label,omitempty"`
	UserAttributeMappings     *string `json:"user_attribute_mappings,omitempty"`
	UserAttributeMacros       *string `json:"user_attribute_macros,omitempty"`
	AttributesTransformations *string `json:"attributes_transformations,omitempty"`
	SkipIfBlank               *bool   `json:"skip_if_blank,omitempty"`
	Values                    *string `json:"values,omitempty,omitempty"`
	DefaultValues             *string `json:"default_values,omitempty"`
	ParamKeyName              *string `json:"param_key_name,omitempty"`
	ProvisionedEntitlements   *bool   `json:"provisioned_entitlements,omitempty"`
	SafeEntitlementsEnabled   *bool   `json:"safe_entitlements_enabled,omitempty"`
	IncludeInSamlAssertion    *bool   `json:"include_in_saml_assertion,omitempty"`
}
