package onelogin

import (
	"github.com/onelogin/onelogin-go-sdk/internal/models"
	utl "github.com/onelogin/onelogin-go-sdk/internal/utilities"
)

const (
	SAMLPath string = "api/2/saml_assertion"
)

func (sdk *OneloginSDK) VerifyFactorSAML(request models.VerifyMFATokenRequest) (interface{}, error) {
	p, err := utl.BuildAPIPath(SAMLPath, "verify_factor")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Post(&p, request)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GenerateSAMLAssertion(request models.GenerateSAMLTokenRequest) (interface{}, error) {
	p, err := utl.BuildAPIPath(SAMLPath)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Post(&p, request)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}
