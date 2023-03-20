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

// checks if the CreateRole201ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRole201ResponseInner{}

// CreateRole201ResponseInner struct for CreateRole201ResponseInner
type CreateRole201ResponseInner struct {
	Id *int32 `json:"id,omitempty"`
}

// NewCreateRole201ResponseInner instantiates a new CreateRole201ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRole201ResponseInner() *CreateRole201ResponseInner {
	this := CreateRole201ResponseInner{}
	return &this
}

// NewCreateRole201ResponseInnerWithDefaults instantiates a new CreateRole201ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRole201ResponseInnerWithDefaults() *CreateRole201ResponseInner {
	this := CreateRole201ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateRole201ResponseInner) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRole201ResponseInner) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateRole201ResponseInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CreateRole201ResponseInner) SetId(v int32) {
	o.Id = &v
}

func (o CreateRole201ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRole201ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableCreateRole201ResponseInner struct {
	value *CreateRole201ResponseInner
	isSet bool
}

func (v NullableCreateRole201ResponseInner) Get() *CreateRole201ResponseInner {
	return v.value
}

func (v *NullableCreateRole201ResponseInner) Set(val *CreateRole201ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRole201ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRole201ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRole201ResponseInner(val *CreateRole201ResponseInner) *NullableCreateRole201ResponseInner {
	return &NullableCreateRole201ResponseInner{value: val, isSet: true}
}

func (v NullableCreateRole201ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRole201ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

