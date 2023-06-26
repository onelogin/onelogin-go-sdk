package models

type GenerateSAMLTokenRequest struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
	AppID           string `json:"app_id"`
	Subdomain       string `json:"subdomain"`
	IPAddress       string `json:"ip_address,omitempty"`
}

type VerifyMFATokenRequest struct {
	AppID       string `json:"app_id"`
	DeviceID    string `json:"device_id"`
	StateToken  string `json:"state_token"`
	OTPToken    string `json:"otp_token,omitempty"`
	DoNotNotify bool   `json:"do_not_notify,omitempty"`
}

type EnrollFactorRequest struct {
	FactorID      int    `json:"factor_id"`
	DisplayName   string `json:"display_name"`
	ExpiresIn     string `json:"expires_in,omitempty"`
	Verified      bool   `json:"verified,omitempty"`
	RedirectTo    string `json:"redirect_to,omitempty"`
	CustomMessage string `json:"custom_message,omitempty"`
}
type ActivateFactorRequest struct {
	DeviceID      string `json:"device_id"`
	ExpiresIn     string `json:"expires_in,omitempty"`
	RedirectTo    string `json:"redirect_to,omitempty"`
	CustomMessage string `json:"custom_message,omitempty"`
}

type GenerateMFATokenRequest struct {
	ExpiresIn string `json:"expires_in,omitempty"`
	Reusable  bool   `json:"reusable,omitempty"`
}
