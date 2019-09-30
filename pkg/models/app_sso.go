package models

type AppSso struct {
	ClientId     string             `json:"client_id"`
	ClientSecret string             `json:"client_secret"`
	MetadataUrl  string             `json:"metadata_url"`
	AcsUrl       string             `json:"acs_url"`
	SlsUrl       string             `json:"sls_url"`
	Issuer       string             `json:"issuer"`
	Certificate  *AppSsoCertificate `json:"certificate"`
}
