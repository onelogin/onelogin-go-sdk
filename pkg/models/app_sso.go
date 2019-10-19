package models

// AppSso is the contact for apps sso.
type AppSso struct {
	ClientID     string             `json:"client_id"`
	ClientSecret string             `json:"client_secret"`
	MetadataURL  string             `json:"metadata_url"`
	AcsURL       string             `json:"acs_url"`
	SlsURL       string             `json:"sls_url"`
	Issuer       string             `json:"issuer"`
	Certificate  *AppSsoCertificate `json:"certificate"`
}
