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

func (sdk *OneloginSDK) GenerateInviteLink(email string) (interface{}, error) {
	p := "api/1/invites/get_invite_link"
	resp, err := sdk.Client.Post(&p, email)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (sdk *OneloginSDK) ListConnectors() (interface{}, error) {
	p := "api/2/connectors"
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (sdk *OneloginSDK) SendInviteLink(email string) (interface{}, error) {
	p := "api/1/invites/send_invite_link"
	resp, err := sdk.Client.Post(&p, email)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
