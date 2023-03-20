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

// checks if the AuthId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthId{}

// AuthId struct for AuthId
type AuthId struct {
	// Unique ID for the Scope
	Id *int32 `json:"id,omitempty"`
}

// NewAuthId instantiates a new AuthId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthId() *AuthId {
	this := AuthId{}
	return &this
}

// NewAuthIdWithDefaults instantiates a new AuthId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthIdWithDefaults() *AuthId {
	this := AuthId{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthId) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthId) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthId) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AuthId) SetId(v int32) {
	o.Id = &v
}

func (o AuthId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableAuthId struct {
	value *AuthId
	isSet bool
}

func (v NullableAuthId) Get() *AuthId {
	return v.value
}

func (v *NullableAuthId) Set(val *AuthId) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthId) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthId(val *AuthId) *NullableAuthId {
	return &NullableAuthId{value: val, isSet: true}
}

func (v NullableAuthId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

