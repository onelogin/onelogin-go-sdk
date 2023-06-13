package onelogin

import (
	"github.com/onelogin/onelogin-go-sdk/internal/models"
	utl "github.com/onelogin/onelogin-go-sdk/internal/utilities"
)

const (
	MFAPath string = "api/2/mfa/users"
)

// https://<subdomain>/api/2/mfa/users/<user_id>/factors
func (sdk *OneloginSDK) GetAuthFactors(userID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(MFAPath, userID, "factors")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

// https://<subdomain>/api/2/mfa/users/<user_id>/registrations
func (sdk *OneloginSDK) EnrollMFAFactor(factor models.EnrollFactorRequest, userID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(MFAPath, userID, "registrations")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Post(&p, factor)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

// https://<subdomain>/api/2/mfa/users/<user_id>/registrations/<registration_id>
func (sdk *OneloginSDK) VerifyMFAEnrollment(userID, registrationID, otp int) (interface{}, error) {
	p, err := utl.BuildAPIPath(MFAPath, userID, "registrations", registrationID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, otp)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}
