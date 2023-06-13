package onelogin

import (
	"github.com/onelogin/onelogin-go-sdk/internal/api"
	olerror "github.com/onelogin/onelogin-go-sdk/internal/error"
)

// OneloginSDK represents the Onelogin SDK.
type OneloginSDK struct {
	Client *api.Client
}

// NewOneloginSDK creates a new instance of the Onelogin SDK.
func NewOneloginSDK() (*OneloginSDK, error) {
	client, err := api.NewClient()
	if err != nil {
		return nil, err
	}
	return &OneloginSDK{Client: client}, nil
}

// GetToken performs the authentication process using the env credentials.
func (sdk *OneloginSDK) GetToken() (string, error) {
	// Call the authenticator to perform the authentication process
	accessTk, err := sdk.Client.Auth.GetToken()
	if err != nil {
		return "", olerror.NewSDKError("Access Token retrieval unsuccessful")
	}
	return accessTk, nil
}

func (sdk *OneloginSDK) GenerateInviteLink() {
	// Implementation for Generate Invite Link endpoint
}

func (sdk *OneloginSDK) SendInviteLink() {
	// Implementation for Send Invite Link endpoint
}

func (sdk *OneloginSDK) ListConnectors() {
	// Implementation for List Connectors endpoint
}

func (sdk *OneloginSDK) CreateDeviceVerification(userID int) {
	// Implementation for Create Device Verification endpoint
}
func (sdk *OneloginSDK) CreateFactorRegistration(userID int)                   {}
func (sdk *OneloginSDK) DeleteEnrolledFactor(userID int, deviceID int)         {}
func (sdk *OneloginSDK) GenerateOTP(userID int)                                {}
func (sdk *OneloginSDK) GetAuthenticationDevices(userID int)                   {}
func (sdk *OneloginSDK) GetUserRegistration(userID int, registrationID int)    {}
func (sdk *OneloginSDK) GetUserVerification(userID int, verificationID int)    {}
func (sdk *OneloginSDK) VerifyUserRegistration(userID int, registrationID int) {}
func (sdk *OneloginSDK) VerifyUserVerification(userID int, verificationID int) {}
func (sdk *OneloginSDK) GenerateMFAToken(userID int)                           {}
func (sdk *OneloginSDK) GetEnrolledFactors(userID int)                         {}
func (sdk *OneloginSDK) GetMFAFactors(userID int)                              {}
func (sdk *OneloginSDK) RemoveMFAFactor(userID int, deviceID int)              {}
func (sdk *OneloginSDK) VerifyMFAFactor(userID int, deviceID int)              {}
