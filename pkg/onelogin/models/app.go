package models

import (
	"encoding/json"
	"fmt"
)

type App struct {
	ID          *int32  `json:"id,omitempty"`
	ConnectorID *int32  `json:"connector_id"`
	Name        *string `json:"name"`
	Description *string `json:"description,omitempty"`
	Notes       *string `json:"notes,omitempty"`
	// PolicyID is the app policy enforced when users sign in to this app, and
	// BrandID the brand its login page uses. Both name a record the app is
	// assigned to, and both are three-state on the wire:
	//
	//	omitted   leave the current assignment alone
	//	a number  assign that policy or brand
	//	null      unassign, falling back to the account default
	//
	// A nil pointer covers the first and a non-nil pointer the second. Nothing
	// covers the third: omitempty drops only the nil pointer, so a pointer can
	// never produce a null. Setting one to 0 is not the way round it -- the app
	// endpoint answers 0 with 422, naming the record it went looking for:
	//
	//	policy_id  "The associated Policy with ID 0 could not be found"
	//	brand_id   "The associated AccountBrand with ID 0 could not be found"
	//
	// which is the same thing it says about an ID that does not exist. Both
	// were confirmed against the API, along with null clearing each field.
	//
	// ClearPolicyID and ClearBrandID below are what send the null.
	PolicyID           *int                  `json:"policy_id,omitempty"`
	BrandID            *int                  `json:"brand_id,omitempty"`
	IconURL            *string               `json:"icon_url,omitempty"`
	Visible            *bool                 `json:"visible,omitempty"`
	AuthMethod         *int                  `json:"auth_method,omitempty"`
	TabID              *int                  `json:"tab_id,omitempty"`
	CreatedAt          *string               `json:"created_at,omitempty"`
	UpdatedAt          *string               `json:"updated_at,omitempty"`
	RoleIDs            *[]int                `json:"role_ids,omitempty"`
	AllowAssumedSignin *bool                 `json:"allow_assumed_signin,omitempty"`
	Provisioning       *Provisioning         `json:"provisioning,omitempty"`
	SSO                interface{}           `json:"sso,omitempty"`
	Configuration      interface{}           `json:"configuration,omitempty"`
	Parameters         *map[string]Parameter `json:"parameters,omitempty"`
	EnforcementPoint   *EnforcementPoint     `json:"enforcement_point,omitempty"`

	// ClearPolicyID and ClearBrandID send the corresponding field as JSON
	// null, which is how the app endpoint is asked to unassign a policy or a
	// brand. They are not fields of the API resource, hence `json:"-"`; they
	// are instructions to MarshalJSON.
	//
	// Setting one alongside a non-nil pointer for the same field is a
	// contradiction -- assign this, and also unassign it -- and MarshalJSON
	// reports it rather than picking a winner, because either choice would
	// silently do half of what the caller asked.
	ClearPolicyID bool `json:"-"`
	ClearBrandID  bool `json:"-"`
}

// MarshalJSON encodes the app, turning ClearPolicyID and ClearBrandID into the
// nulls that unassign a policy or a brand.
//
// With neither flag set it returns exactly what the struct tags alone would
// produce, so callers that predate the flags are unaffected.
//
// When a flag is set the encoding is rebuilt from a map, which orders keys
// alphabetically rather than by field. The values are carried across as
// json.RawMessage and so are not re-encoded, and object key order carries no
// meaning in JSON, so only the byte layout differs.
func (a App) MarshalJSON() ([]byte, error) {
	// A distinct type with the same fields and none of the methods. Marshaling
	// App directly here would call this method again.
	type plain App

	raw, err := json.Marshal(plain(a))
	if err != nil {
		return nil, err
	}

	if !a.ClearPolicyID && !a.ClearBrandID {
		return raw, nil
	}

	if a.ClearPolicyID && a.PolicyID != nil {
		return nil, fmt.Errorf("policy_id cannot be assigned (%d) and cleared in the same request: PolicyID and ClearPolicyID are both set", *a.PolicyID)
	}
	if a.ClearBrandID && a.BrandID != nil {
		return nil, fmt.Errorf("brand_id cannot be assigned (%d) and cleared in the same request: BrandID and ClearBrandID are both set", *a.BrandID)
	}

	var fields map[string]json.RawMessage
	if err := json.Unmarshal(raw, &fields); err != nil {
		return nil, err
	}

	null := json.RawMessage("null")
	if a.ClearPolicyID {
		fields["policy_id"] = null
	}
	if a.ClearBrandID {
		fields["brand_id"] = null
	}

	return json.Marshal(fields)
}

type Provisioning struct {
	Enabled bool `json:"enabled"`
}

type SSO interface {
	ValidateSSO() error
}

type SSOOpenId struct {
	ClientID string `json:"client_id"`
}

type SSOSAML struct {
	MetadataURL string      `json:"metadata_url"`
	AcsURL      string      `json:"acs_url"`
	SlsURL      string      `json:"sls_url"`
	Issuer      string      `json:"issuer"`
	Certificate Certificate `json:"certificate"`
}

type Certificate struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

// ConfigurationOpenId holds the OIDC settings on an app's Configuration tab.
//
// PostLogoutRedirectURI carries one or more URIs as a single newline- or
// comma-separated string, matching how the API stores the connector's
// post_logout_redirect_uri parameter. It is a pointer because the API treats a
// present key as an assignment, giving three distinct requests:
//
//	nil            omit the key; leave the app's existing URIs alone
//	pointer to ""  send ""; clear every URI configured on the app
//	pointer to "…" send the value; replace the app's URIs
type ConfigurationOpenId struct {
	RedirectURI                   string  `json:"redirect_uri"`
	PostLogoutRedirectURI         *string `json:"post_logout_redirect_uri,omitempty"`
	IncludeAmrClaims              *bool   `json:"include_amr_claims,omitempty"`
	LoginURL                      string  `json:"login_url"`
	OidcApplicationType           int     `json:"oidc_application_type"`
	TokenEndpointAuthMethod       int     `json:"token_endpoint_auth_method"`
	AccessTokenExpirationMinutes  int     `json:"access_token_expiration_minutes"`
	RefreshTokenExpirationMinutes int     `json:"refresh_token_expiration_minutes"`
}

type ConfigurationSAML struct {
	ProviderArn        interface{} `json:"provider_arn"`
	SignatureAlgorithm string      `json:"signature_algorithm"`
	CertificateID      int         `json:"certificate_id"`
}

type Parameter struct {
	Values                    interface{} `json:"values,omitempty"`
	UserAttributeMappings     interface{} `json:"user_attribute_mappings,omitempty"`
	ProvisionedEntitlements   bool        `json:"provisioned_entitlements,omitempty"`
	SkipIfBlank               bool        `json:"skip_if_blank,omitempty"`
	ID                        int         `json:"id,omitempty"`
	DefaultValues             interface{} `json:"default_values"`
	AttributesTransformations interface{} `json:"attributes_transformations,omitempty"`
	Label                     string      `json:"label,omitempty"`
	UserAttributeMacros       interface{} `json:"user_attribute_macros,omitempty"`
	IncludeInSamlAssertion    bool        `json:"include_in_saml_assertion,omitempty"`
	// A pointer, unlike the bools above it. With `bool,omitempty` a false is
	// indistinguishable from unset and never reaches the API, so the field
	// could be turned on and never off again.
	SafeEntitlementsEnabled *bool `json:"safe_entitlements_enabled,omitempty"`
}

type EnforcementPoint struct {
	RequireSitewideAuthentication bool        `json:"require_sitewide_authentication"`
	Conditions                    *Conditions `json:"conditions,omitempty"`
	SessionExpiryFixed            Duration    `json:"session_expiry_fixed"`
	SessionExpiryInactivity       Duration    `json:"session_expiry_inactivity"`
	Permissions                   string      `json:"permissions"`
	Token                         string      `json:"token,omitempty"`
	Target                        string      `json:"target"`
	Resources                     []Resource  `json:"resources"`
	ContextRoot                   string      `json:"context_root"`
	UseTargetHostHeader           bool        `json:"use_target_host_header"`
	Vhost                         string      `json:"vhost"`
	LandingPage                   string      `json:"landing_page"`
	CaseSensitive                 bool        `json:"case_sensitive"`
}

type Conditions struct {
	Type  string   `json:"type"`
	Roles []string `json:"roles"`
}

type Duration struct {
	Value int `json:"value"`
	Unit  int `json:"unit"`
}

type Resource struct {
	Path        string  `json:"path"`
	RequireAuth string  `json:"require_authentication"`
	Permissions string  `json:"permissions"`
	Conditions  *string `json:"conditions,omitempty"`
	IsPathRegex *bool   `json:"is_path_regex,omitempty"`
	ResourceID  int     `json:"resource_id,omitempty"`
}

const (
	UnitSeconds = 0
	UnitMinutes = 1
	UnitHours   = 2
)

type AppQuery struct {
	Limit       string  `json:"limit,omitempty"`
	Page        string  `json:"page,omitempty"`
	Cursor      string  `json:"cursor,omitempty"`
	Name        *string `json:"name,omitempty"`
	ConnectorID *int    `json:"connector_id,omitempty"`
	AuthMethod  *int    `json:"auth_method,omitempty"`
}

func (q *AppQuery) GetKeyValidators() map[string]func(interface{}) bool {
	return map[string]func(interface{}) bool{
		"limit":        validateString,
		"page":         validateString,
		"cursor":       validateString,
		"name":         validateString,
		"connector_id": validateInt,
		"auth_method":  validateInt,
	}
}

// AppUserQuery represents query parameters for listing app users with pagination support
type AppUserQuery struct {
	Limit  string `json:"limit,omitempty"`
	Page   string `json:"page,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

func (q *AppUserQuery) GetKeyValidators() map[string]func(interface{}) bool {
	return map[string]func(interface{}) bool{
		"limit":  validateString,
		"page":   validateString,
		"cursor": validateString,
	}
}
