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

// checks if the GetRoleById200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRoleById200ResponseDataInner{}

// GetRoleById200ResponseDataInner struct for GetRoleById200ResponseDataInner
type GetRoleById200ResponseDataInner struct {
	// role's unique ID in Onelogin
	Id *int32 `json:"id,omitempty"`
	// Role name
	Name *string `json:"name,omitempty"`
}

// NewGetRoleById200ResponseDataInner instantiates a new GetRoleById200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRoleById200ResponseDataInner() *GetRoleById200ResponseDataInner {
	this := GetRoleById200ResponseDataInner{}
	return &this
}

// NewGetRoleById200ResponseDataInnerWithDefaults instantiates a new GetRoleById200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRoleById200ResponseDataInnerWithDefaults() *GetRoleById200ResponseDataInner {
	this := GetRoleById200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetRoleById200ResponseDataInner) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleById200ResponseDataInner) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetRoleById200ResponseDataInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetRoleById200ResponseDataInner) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetRoleById200ResponseDataInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRoleById200ResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetRoleById200ResponseDataInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetRoleById200ResponseDataInner) SetName(v string) {
	o.Name = &v
}

func (o GetRoleById200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRoleById200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableGetRoleById200ResponseDataInner struct {
	value *GetRoleById200ResponseDataInner
	isSet bool
}

func (v NullableGetRoleById200ResponseDataInner) Get() *GetRoleById200ResponseDataInner {
	return v.value
}

func (v *NullableGetRoleById200ResponseDataInner) Set(val *GetRoleById200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRoleById200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRoleById200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRoleById200ResponseDataInner(val *GetRoleById200ResponseDataInner) *NullableGetRoleById200ResponseDataInner {
	return &NullableGetRoleById200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetRoleById200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRoleById200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


