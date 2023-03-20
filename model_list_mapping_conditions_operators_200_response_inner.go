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

// checks if the ListMappingConditionsOperators200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMappingConditionsOperators200ResponseInner{}

// ListMappingConditionsOperators200ResponseInner struct for ListMappingConditionsOperators200ResponseInner
type ListMappingConditionsOperators200ResponseInner struct {
	// Name or description of operator
	Name *string `json:"name,omitempty"`
	// The condition operator value to use when creating or updating User Mappings.
	Value *string `json:"value,omitempty"`
}

// NewListMappingConditionsOperators200ResponseInner instantiates a new ListMappingConditionsOperators200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMappingConditionsOperators200ResponseInner() *ListMappingConditionsOperators200ResponseInner {
	this := ListMappingConditionsOperators200ResponseInner{}
	return &this
}

// NewListMappingConditionsOperators200ResponseInnerWithDefaults instantiates a new ListMappingConditionsOperators200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMappingConditionsOperators200ResponseInnerWithDefaults() *ListMappingConditionsOperators200ResponseInner {
	this := ListMappingConditionsOperators200ResponseInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListMappingConditionsOperators200ResponseInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMappingConditionsOperators200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListMappingConditionsOperators200ResponseInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListMappingConditionsOperators200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ListMappingConditionsOperators200ResponseInner) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMappingConditionsOperators200ResponseInner) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ListMappingConditionsOperators200ResponseInner) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ListMappingConditionsOperators200ResponseInner) SetValue(v string) {
	o.Value = &v
}

func (o ListMappingConditionsOperators200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMappingConditionsOperators200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableListMappingConditionsOperators200ResponseInner struct {
	value *ListMappingConditionsOperators200ResponseInner
	isSet bool
}

func (v NullableListMappingConditionsOperators200ResponseInner) Get() *ListMappingConditionsOperators200ResponseInner {
	return v.value
}

func (v *NullableListMappingConditionsOperators200ResponseInner) Set(val *ListMappingConditionsOperators200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListMappingConditionsOperators200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListMappingConditionsOperators200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMappingConditionsOperators200ResponseInner(val *ListMappingConditionsOperators200ResponseInner) *NullableListMappingConditionsOperators200ResponseInner {
	return &NullableListMappingConditionsOperators200ResponseInner{value: val, isSet: true}
}

func (v NullableListMappingConditionsOperators200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMappingConditionsOperators200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


