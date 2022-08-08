/*
OneLogin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// VerifyFactorSamlRequest struct for VerifyFactorSamlRequest
type VerifyFactorSamlRequest struct {
	// App ID of the app for which you want to generate a SAML token. This is the app ID in OneLogin.
	AppId string `json:"app_id"`
	// Provide the MFA device_id you are submitting for verification. The device_id is supplied by the Generate SAML Assertion API.
	DeviceId string `json:"device_id"`
	// state_token associated with the MFA device_id you are submitting. The state_token is supplied by the Generate SAML Assertion API.
	StateToken string `json:"state_token"`
	// Provide the OTP value for the MFA factor you are submitting for verification.
	OtpToken *string `json:"otp_token,omitempty"`
	// When verifying MFA via Protect Push, set this to true to stop additional push notifications being sent to the OneLogin Protect device.
	DoNotNotify *bool `json:"do_not_notify,omitempty"`
}

// NewVerifyFactorSamlRequest instantiates a new VerifyFactorSamlRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVerifyFactorSamlRequest(appId string, deviceId string, stateToken string) *VerifyFactorSamlRequest {
	this := VerifyFactorSamlRequest{}
	this.AppId = appId
	this.DeviceId = deviceId
	this.StateToken = stateToken
	return &this
}

// NewVerifyFactorSamlRequestWithDefaults instantiates a new VerifyFactorSamlRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVerifyFactorSamlRequestWithDefaults() *VerifyFactorSamlRequest {
	this := VerifyFactorSamlRequest{}
	return &this
}

// GetAppId returns the AppId field value
func (o *VerifyFactorSamlRequest) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *VerifyFactorSamlRequest) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *VerifyFactorSamlRequest) SetAppId(v string) {
	o.AppId = v
}

// GetDeviceId returns the DeviceId field value
func (o *VerifyFactorSamlRequest) GetDeviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value
// and a boolean to check if the value has been set.
func (o *VerifyFactorSamlRequest) GetDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeviceId, true
}

// SetDeviceId sets field value
func (o *VerifyFactorSamlRequest) SetDeviceId(v string) {
	o.DeviceId = v
}

// GetStateToken returns the StateToken field value
func (o *VerifyFactorSamlRequest) GetStateToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StateToken
}

// GetStateTokenOk returns a tuple with the StateToken field value
// and a boolean to check if the value has been set.
func (o *VerifyFactorSamlRequest) GetStateTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StateToken, true
}

// SetStateToken sets field value
func (o *VerifyFactorSamlRequest) SetStateToken(v string) {
	o.StateToken = v
}

// GetOtpToken returns the OtpToken field value if set, zero value otherwise.
func (o *VerifyFactorSamlRequest) GetOtpToken() string {
	if o == nil || o.OtpToken == nil {
		var ret string
		return ret
	}
	return *o.OtpToken
}

// GetOtpTokenOk returns a tuple with the OtpToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyFactorSamlRequest) GetOtpTokenOk() (*string, bool) {
	if o == nil || o.OtpToken == nil {
		return nil, false
	}
	return o.OtpToken, true
}

// HasOtpToken returns a boolean if a field has been set.
func (o *VerifyFactorSamlRequest) HasOtpToken() bool {
	if o != nil && o.OtpToken != nil {
		return true
	}

	return false
}

// SetOtpToken gets a reference to the given string and assigns it to the OtpToken field.
func (o *VerifyFactorSamlRequest) SetOtpToken(v string) {
	o.OtpToken = &v
}

// GetDoNotNotify returns the DoNotNotify field value if set, zero value otherwise.
func (o *VerifyFactorSamlRequest) GetDoNotNotify() bool {
	if o == nil || o.DoNotNotify == nil {
		var ret bool
		return ret
	}
	return *o.DoNotNotify
}

// GetDoNotNotifyOk returns a tuple with the DoNotNotify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VerifyFactorSamlRequest) GetDoNotNotifyOk() (*bool, bool) {
	if o == nil || o.DoNotNotify == nil {
		return nil, false
	}
	return o.DoNotNotify, true
}

// HasDoNotNotify returns a boolean if a field has been set.
func (o *VerifyFactorSamlRequest) HasDoNotNotify() bool {
	if o != nil && o.DoNotNotify != nil {
		return true
	}

	return false
}

// SetDoNotNotify gets a reference to the given bool and assigns it to the DoNotNotify field.
func (o *VerifyFactorSamlRequest) SetDoNotNotify(v bool) {
	o.DoNotNotify = &v
}

func (o VerifyFactorSamlRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["app_id"] = o.AppId
	}
	if true {
		toSerialize["device_id"] = o.DeviceId
	}
	if true {
		toSerialize["state_token"] = o.StateToken
	}
	if o.OtpToken != nil {
		toSerialize["otp_token"] = o.OtpToken
	}
	if o.DoNotNotify != nil {
		toSerialize["do_not_notify"] = o.DoNotNotify
	}
	return json.Marshal(toSerialize)
}

type NullableVerifyFactorSamlRequest struct {
	value *VerifyFactorSamlRequest
	isSet bool
}

func (v NullableVerifyFactorSamlRequest) Get() *VerifyFactorSamlRequest {
	return v.value
}

func (v *NullableVerifyFactorSamlRequest) Set(val *VerifyFactorSamlRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVerifyFactorSamlRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVerifyFactorSamlRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVerifyFactorSamlRequest(val *VerifyFactorSamlRequest) *NullableVerifyFactorSamlRequest {
	return &NullableVerifyFactorSamlRequest{value: val, isSet: true}
}

func (v NullableVerifyFactorSamlRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVerifyFactorSamlRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


