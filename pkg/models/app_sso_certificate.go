package models

// AppSsoCertificate is the conract for sso certificate.
type AppSsoCertificate struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
