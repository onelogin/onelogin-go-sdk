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

// GenerateMfaToken200Response struct for GenerateMfaToken200Response
type GenerateMfaToken200Response struct {
	// Token can function as a temporary MFA token. It can be used to authenticate for any app when valid.
	MfaToken *string `json:"mfa_token,omitempty"`
	// true indcates the token can be used multiple times. false indicates the token is invalid after a single use
	Reusable *bool `json:"reusable,omitempty"`
	// Defines the expiration time and date for the token. Format is UTC time.
	ExpiresAt *string `json:"expires_at,omitempty"`
	// Defines the expiration time and date for the token. Format is UTC time.
	DeviceId *string `json:"device_id,omitempty"`
}

// NewGenerateMfaToken200Response instantiates a new GenerateMfaToken200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateMfaToken200Response() *GenerateMfaToken200Response {
	this := GenerateMfaToken200Response{}
	return &this
}

// NewGenerateMfaToken200ResponseWithDefaults instantiates a new GenerateMfaToken200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateMfaToken200ResponseWithDefaults() *GenerateMfaToken200Response {
	this := GenerateMfaToken200Response{}
	return &this
}

// GetMfaToken returns the MfaToken field value if set, zero value otherwise.
func (o *GenerateMfaToken200Response) GetMfaToken() string {
	if o == nil || o.MfaToken == nil {
		var ret string
		return ret
	}
	return *o.MfaToken
}

// GetMfaTokenOk returns a tuple with the MfaToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateMfaToken200Response) GetMfaTokenOk() (*string, bool) {
	if o == nil || o.MfaToken == nil {
		return nil, false
	}
	return o.MfaToken, true
}

// HasMfaToken returns a boolean if a field has been set.
func (o *GenerateMfaToken200Response) HasMfaToken() bool {
	if o != nil && o.MfaToken != nil {
		return true
	}

	return false
}

// SetMfaToken gets a reference to the given string and assigns it to the MfaToken field.
func (o *GenerateMfaToken200Response) SetMfaToken(v string) {
	o.MfaToken = &v
}

// GetReusable returns the Reusable field value if set, zero value otherwise.
func (o *GenerateMfaToken200Response) GetReusable() bool {
	if o == nil || o.Reusable == nil {
		var ret bool
		return ret
	}
	return *o.Reusable
}

// GetReusableOk returns a tuple with the Reusable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateMfaToken200Response) GetReusableOk() (*bool, bool) {
	if o == nil || o.Reusable == nil {
		return nil, false
	}
	return o.Reusable, true
}

// HasReusable returns a boolean if a field has been set.
func (o *GenerateMfaToken200Response) HasReusable() bool {
	if o != nil && o.Reusable != nil {
		return true
	}

	return false
}

// SetReusable gets a reference to the given bool and assigns it to the Reusable field.
func (o *GenerateMfaToken200Response) SetReusable(v bool) {
	o.Reusable = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *GenerateMfaToken200Response) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateMfaToken200Response) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *GenerateMfaToken200Response) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *GenerateMfaToken200Response) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *GenerateMfaToken200Response) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateMfaToken200Response) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *GenerateMfaToken200Response) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *GenerateMfaToken200Response) SetDeviceId(v string) {
	o.DeviceId = &v
}

func (o GenerateMfaToken200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MfaToken != nil {
		toSerialize["mfa_token"] = o.MfaToken
	}
	if o.Reusable != nil {
		toSerialize["reusable"] = o.Reusable
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.DeviceId != nil {
		toSerialize["device_id"] = o.DeviceId
	}
	return json.Marshal(toSerialize)
}

type NullableGenerateMfaToken200Response struct {
	value *GenerateMfaToken200Response
	isSet bool
}

func (v NullableGenerateMfaToken200Response) Get() *GenerateMfaToken200Response {
	return v.value
}

func (v *NullableGenerateMfaToken200Response) Set(val *GenerateMfaToken200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateMfaToken200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateMfaToken200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateMfaToken200Response(val *GenerateMfaToken200Response) *NullableGenerateMfaToken200Response {
	return &NullableGenerateMfaToken200Response{value: val, isSet: true}
}

func (v NullableGenerateMfaToken200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenerateMfaToken200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

