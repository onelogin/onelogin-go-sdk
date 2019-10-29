package models

// AppSsoCertificate is the conract for sso certificate.
type AppSsoCertificate struct {
	ID    int32  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
