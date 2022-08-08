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

// GetClientApps200ResponseInner struct for GetClientApps200ResponseInner
type GetClientApps200ResponseInner struct {
	Scopes []GetClientApps200ResponseInnerScopesInner `json:"scopes,omitempty"`
	AppId *int32 `json:"app_id,omitempty"`
	Name *string `json:"name,omitempty"`
	ApiAuthId *int32 `json:"api_auth_id,omitempty"`
}

// NewGetClientApps200ResponseInner instantiates a new GetClientApps200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetClientApps200ResponseInner() *GetClientApps200ResponseInner {
	this := GetClientApps200ResponseInner{}
	return &this
}

// NewGetClientApps200ResponseInnerWithDefaults instantiates a new GetClientApps200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetClientApps200ResponseInnerWithDefaults() *GetClientApps200ResponseInner {
	this := GetClientApps200ResponseInner{}
	return &this
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *GetClientApps200ResponseInner) GetScopes() []GetClientApps200ResponseInnerScopesInner {
	if o == nil || o.Scopes == nil {
		var ret []GetClientApps200ResponseInnerScopesInner
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClientApps200ResponseInner) GetScopesOk() ([]GetClientApps200ResponseInnerScopesInner, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *GetClientApps200ResponseInner) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []GetClientApps200ResponseInnerScopesInner and assigns it to the Scopes field.
func (o *GetClientApps200ResponseInner) SetScopes(v []GetClientApps200ResponseInnerScopesInner) {
	o.Scopes = v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *GetClientApps200ResponseInner) GetAppId() int32 {
	if o == nil || o.AppId == nil {
		var ret int32
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClientApps200ResponseInner) GetAppIdOk() (*int32, bool) {
	if o == nil || o.AppId == nil {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *GetClientApps200ResponseInner) HasAppId() bool {
	if o != nil && o.AppId != nil {
		return true
	}

	return false
}

// SetAppId gets a reference to the given int32 and assigns it to the AppId field.
func (o *GetClientApps200ResponseInner) SetAppId(v int32) {
	o.AppId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetClientApps200ResponseInner) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClientApps200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetClientApps200ResponseInner) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetClientApps200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetApiAuthId returns the ApiAuthId field value if set, zero value otherwise.
func (o *GetClientApps200ResponseInner) GetApiAuthId() int32 {
	if o == nil || o.ApiAuthId == nil {
		var ret int32
		return ret
	}
	return *o.ApiAuthId
}

// GetApiAuthIdOk returns a tuple with the ApiAuthId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetClientApps200ResponseInner) GetApiAuthIdOk() (*int32, bool) {
	if o == nil || o.ApiAuthId == nil {
		return nil, false
	}
	return o.ApiAuthId, true
}

// HasApiAuthId returns a boolean if a field has been set.
func (o *GetClientApps200ResponseInner) HasApiAuthId() bool {
	if o != nil && o.ApiAuthId != nil {
		return true
	}

	return false
}

// SetApiAuthId gets a reference to the given int32 and assigns it to the ApiAuthId field.
func (o *GetClientApps200ResponseInner) SetApiAuthId(v int32) {
	o.ApiAuthId = &v
}

func (o GetClientApps200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.AppId != nil {
		toSerialize["app_id"] = o.AppId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ApiAuthId != nil {
		toSerialize["api_auth_id"] = o.ApiAuthId
	}
	return json.Marshal(toSerialize)
}

type NullableGetClientApps200ResponseInner struct {
	value *GetClientApps200ResponseInner
	isSet bool
}

func (v NullableGetClientApps200ResponseInner) Get() *GetClientApps200ResponseInner {
	return v.value
}

func (v *NullableGetClientApps200ResponseInner) Set(val *GetClientApps200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetClientApps200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetClientApps200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetClientApps200ResponseInner(val *GetClientApps200ResponseInner) *NullableGetClientApps200ResponseInner {
	return &NullableGetClientApps200ResponseInner{value: val, isSet: true}
}

func (v NullableGetClientApps200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetClientApps200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

