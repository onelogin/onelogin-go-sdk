package saml

// SAMLAssertion test
type SAMLAssertion struct {
	Data        *string    `json:"data,omitempty"`
	Message     *string    `json:"message,omitempty"`
	StateToken  *string    `json:"state_token,omitempty"`
	User        *User      `json:"user,omitempty"`
	Devices     []*Devices `json:"devices,omitempty"`
	CallbackURL *string    `json:"callback_url,omitempty"`
}

// User user
type User struct {
	ID        *string
	Email     *string
	Username  *string
	Firstname *string
	Lastname  *string
}

// Devices devices
type Devices struct {
	ID   *int    `json:"device_id,omitempty"`
	Type *string `json:"device_type,omitempty"`
}

// SAMLAssertionRequest test
type SAMLAssertionRequest struct {
	UsernameOrEmail *string `json:"username_or_email,omitempty"`
	Password        *string `json:"password,omitempty"`
	AppID           *string `json:"app_id,omitempty"`
	Subdomain       *string `json:"subdomain,omitempty"`
	IPAddress       *string `json:"ip_address,omitempty"`
}

// VerifyFactor test
type VerifyFactor struct {
	Data    *string `json:"data,omitempty"`
	Message *string `json:"message,omitempty"`
}

// VerifyFactorRequest test
type VerifyFactorRequest struct {
	AppID       *string `json:"app_id,omitempty"`
	DeviceID    *string `json:"device_id,omitempty"`
	StateToken  *string `json:"state_token,omitempty"`
	OTPToken    *string `json:"otp_token,omitempty"`
	DoNotVerify *bool   `json:"do_not_verify,omitempty"`
}
