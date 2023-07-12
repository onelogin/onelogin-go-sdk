package onelogin

import (
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
	utl "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/utilities"
)

const (
	MFAPath string = "api/2/mfa/users"
)

// https://<subdomain>/api/2/mfa/users/<user_id>/factors
func (sdk *OneloginSDK) GetAvailableMFAFactors(userID int) (interface{}, error) {
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

// https://<subdomain>/api/2/mfa/users/<user_id>/verifications
func (sdk *OneloginSDK) ActivateMFAFactor(userID int, request models.ActivateFactorRequest) (interface{}, error) {
	p, err := utl.BuildAPIPath(MFAPath, userID, "verifications")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Post(&p, request)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

// https://<subdomain>/api/2/mfa/users/<user_id>/devices/<device_id>
func (sdk *OneloginSDK) RemoveMFAFactor(userID, deviceID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(MFAPath, userID, "devices", deviceID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Delete(&p)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

// https://<subdomain>/api/2/mfa/users/<user_id>/factors
func (sdk *OneloginSDK) GetEnrolledMFAFactors(userID int) (interface{}, error) {
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

// https://<subdomain>/api/2/mfa/users/:user_id/mfa_token
func (sdk *OneloginSDK) GenerateMFAToken(userID int, request models.GenerateMFATokenRequest) (interface{}, error) {
	p, err := utl.BuildAPIPath(MFAPath, userID, "mfa_token")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Post(&p, request)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}
