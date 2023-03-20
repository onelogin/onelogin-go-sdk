/*
OneLogin API

OpenAPI Specification for OneLogin

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onelogin

import (
	"encoding/json"
)

// checks if the GetAuthenticationDevices200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAuthenticationDevices200ResponseInner{}

// GetAuthenticationDevices200ResponseInner struct for GetAuthenticationDevices200ResponseInner
type GetAuthenticationDevices200ResponseInner struct {
	// MFA device identifier.
	DeviceId *string `json:"device_id,omitempty"`
	// Authentication factor display name assigned by users when they register the device.
	UserDisplayName *string `json:"user_display_name,omitempty"`
	// Authentication factor display name as it appears to users upon initial registration, as defined by admins at Settings > Authentication Factors.
	TypeDisplayName *string `json:"type_display_name,omitempty"`
	// Authentication factor name, as it appears to administrators in OneLogin.
	AuthFactorName *string `json:"auth_factor_name,omitempty"`
	// true = is user’s default MFA device for OneLogin.
	Default *bool `json:"default,omitempty"`
}

// NewGetAuthenticationDevices200ResponseInner instantiates a new GetAuthenticationDevices200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAuthenticationDevices200ResponseInner() *GetAuthenticationDevices200ResponseInner {
	this := GetAuthenticationDevices200ResponseInner{}
	var default_ bool = false
	this.Default = &default_
	return &this
}

// NewGetAuthenticationDevices200ResponseInnerWithDefaults instantiates a new GetAuthenticationDevices200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAuthenticationDevices200ResponseInnerWithDefaults() *GetAuthenticationDevices200ResponseInner {
	this := GetAuthenticationDevices200ResponseInner{}
	var default_ bool = false
	this.Default = &default_
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *GetAuthenticationDevices200ResponseInner) GetDeviceId() string {
	if o == nil || isNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAuthenticationDevices200ResponseInner) GetDeviceIdOk() (*string, bool) {
	if o == nil || isNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *GetAuthenticationDevices200ResponseInner) HasDeviceId() bool {
	if o != nil && !isNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *GetAuthenticationDevices200ResponseInner) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetUserDisplayName returns the UserDisplayName field value if set, zero value otherwise.
func (o *GetAuthenticationDevices200ResponseInner) GetUserDisplayName() string {
	if o == nil || isNil(o.UserDisplayName) {
		var ret string
		return ret
	}
	return *o.UserDisplayName
}

// GetUserDisplayNameOk returns a tuple with the UserDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAuthenticationDevices200ResponseInner) GetUserDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.UserDisplayName) {
		return nil, false
	}
	return o.UserDisplayName, true
}

// HasUserDisplayName returns a boolean if a field has been set.
func (o *GetAuthenticationDevices200ResponseInner) HasUserDisplayName() bool {
	if o != nil && !isNil(o.UserDisplayName) {
		return true
	}

	return false
}

// SetUserDisplayName gets a reference to the given string and assigns it to the UserDisplayName field.
func (o *GetAuthenticationDevices200ResponseInner) SetUserDisplayName(v string) {
	o.UserDisplayName = &v
}

// GetTypeDisplayName returns the TypeDisplayName field value if set, zero value otherwise.
func (o *GetAuthenticationDevices200ResponseInner) GetTypeDisplayName() string {
	if o == nil || isNil(o.TypeDisplayName) {
		var ret string
		return ret
	}
	return *o.TypeDisplayName
}

// GetTypeDisplayNameOk returns a tuple with the TypeDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAuthenticationDevices200ResponseInner) GetTypeDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.TypeDisplayName) {
		return nil, false
	}
	return o.TypeDisplayName, true
}

// HasTypeDisplayName returns a boolean if a field has been set.
func (o *GetAuthenticationDevices200ResponseInner) HasTypeDisplayName() bool {
	if o != nil && !isNil(o.TypeDisplayName) {
		return true
	}

	return false
}

// SetTypeDisplayName gets a reference to the given string and assigns it to the TypeDisplayName field.
func (o *GetAuthenticationDevices200ResponseInner) SetTypeDisplayName(v string) {
	o.TypeDisplayName = &v
}

// GetAuthFactorName returns the AuthFactorName field value if set, zero value otherwise.
func (o *GetAuthenticationDevices200ResponseInner) GetAuthFactorName() string {
	if o == nil || isNil(o.AuthFactorName) {
		var ret string
		return ret
	}
	return *o.AuthFactorName
}

// GetAuthFactorNameOk returns a tuple with the AuthFactorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAuthenticationDevices200ResponseInner) GetAuthFactorNameOk() (*string, bool) {
	if o == nil || isNil(o.AuthFactorName) {
		return nil, false
	}
	return o.AuthFactorName, true
}

// HasAuthFactorName returns a boolean if a field has been set.
func (o *GetAuthenticationDevices200ResponseInner) HasAuthFactorName() bool {
	if o != nil && !isNil(o.AuthFactorName) {
		return true
	}

	return false
}

// SetAuthFactorName gets a reference to the given string and assigns it to the AuthFactorName field.
func (o *GetAuthenticationDevices200ResponseInner) SetAuthFactorName(v string) {
	o.AuthFactorName = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *GetAuthenticationDevices200ResponseInner) GetDefault() bool {
	if o == nil || isNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAuthenticationDevices200ResponseInner) GetDefaultOk() (*bool, bool) {
	if o == nil || isNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *GetAuthenticationDevices200ResponseInner) HasDefault() bool {
	if o != nil && !isNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *GetAuthenticationDevices200ResponseInner) SetDefault(v bool) {
	o.Default = &v
}

func (o GetAuthenticationDevices200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAuthenticationDevices200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DeviceId) {
		toSerialize["device_id"] = o.DeviceId
	}
	if !isNil(o.UserDisplayName) {
		toSerialize["user_display_name"] = o.UserDisplayName
	}
	if !isNil(o.TypeDisplayName) {
		toSerialize["type_display_name"] = o.TypeDisplayName
	}
	if !isNil(o.AuthFactorName) {
		toSerialize["auth_factor_name"] = o.AuthFactorName
	}
	if !isNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	return toSerialize, nil
}

type NullableGetAuthenticationDevices200ResponseInner struct {
	value *GetAuthenticationDevices200ResponseInner
	isSet bool
}

func (v NullableGetAuthenticationDevices200ResponseInner) Get() *GetAuthenticationDevices200ResponseInner {
	return v.value
}

func (v *NullableGetAuthenticationDevices200ResponseInner) Set(val *GetAuthenticationDevices200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAuthenticationDevices200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAuthenticationDevices200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAuthenticationDevices200ResponseInner(val *GetAuthenticationDevices200ResponseInner) *NullableGetAuthenticationDevices200ResponseInner {
	return &NullableGetAuthenticationDevices200ResponseInner{value: val, isSet: true}
}

func (v NullableGetAuthenticationDevices200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAuthenticationDevices200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

