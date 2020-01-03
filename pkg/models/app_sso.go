package models

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
