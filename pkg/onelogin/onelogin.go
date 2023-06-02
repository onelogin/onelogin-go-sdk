package onelogin

import (
	"github.com/onelogin/onelogin-go-sdk/internal/api"
	"github.com/onelogin/onelogin-go-sdk/internal/authentication"
	olError "github.com/onelogin/onelogin-go-sdk/internal/error"
)

// OneloginSDK represents the Onelogin SDK.
type OneloginSDK struct {
	Auth      *authentication.Authenticator
	ApiClient *api.Client
}

// NewOneloginSDK creates a new instance of the Onelogin SDK.
func NewOneloginSDK() *OneloginSDK {
	return &OneloginSDK{
		Auth:      authentication.NewAuthenticator(),
		ApiClient: api.NewClient(),
	}
}

// GetToken performs the authentication process using the env credentials.
func (sdk *OneloginSDK) GetToken() (string, error) {
	// Call the authenticator to perform the authentication process
	token, err := sdk.Auth.GetToken()
	if err != nil {
		return "", olError.NewAuthenticationError("OneLoginSDK token retrieval error")
	}
	if token == "" {
		return sdk.Auth.GenerateToken()
	}
	return token, nil
}

// ...
// Implement additional methods for various SDK functionalities
// ...
