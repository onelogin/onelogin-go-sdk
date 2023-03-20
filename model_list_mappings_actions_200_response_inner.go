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

// checks if the ListMappingsActions200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMappingsActions200ResponseInner{}

// ListMappingsActions200ResponseInner struct for ListMappingsActions200ResponseInner
type ListMappingsActions200ResponseInner struct {
	// Name of the action
	Name *string `json:"name,omitempty"`
	// The unique identifier of the action. This should be used when defining actions for a User Mapping.
	Value *string `json:"value,omitempty"`
}

// NewListMappingsActions200ResponseInner instantiates a new ListMappingsActions200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMappingsActions200ResponseInner() *ListMappingsActions200ResponseInner {
	this := ListMappingsActions200ResponseInner{}
	return &this
}

// NewListMappingsActions200ResponseInnerWithDefaults instantiates a new ListMappingsActions200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMappingsActions200ResponseInnerWithDefaults() *ListMappingsActions200ResponseInner {
	this := ListMappingsActions200ResponseInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListMappingsActions200ResponseInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMappingsActions200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListMappingsActions200ResponseInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListMappingsActions200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ListMappingsActions200ResponseInner) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMappingsActions200ResponseInner) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ListMappingsActions200ResponseInner) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ListMappingsActions200ResponseInner) SetValue(v string) {
	o.Value = &v
}

func (o ListMappingsActions200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMappingsActions200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableListMappingsActions200ResponseInner struct {
	value *ListMappingsActions200ResponseInner
	isSet bool
}

func (v NullableListMappingsActions200ResponseInner) Get() *ListMappingsActions200ResponseInner {
	return v.value
}

func (v *NullableListMappingsActions200ResponseInner) Set(val *ListMappingsActions200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListMappingsActions200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListMappingsActions200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMappingsActions200ResponseInner(val *ListMappingsActions200ResponseInner) *NullableListMappingsActions200ResponseInner {
	return &NullableListMappingsActions200ResponseInner{value: val, isSet: true}
}

func (v NullableListMappingsActions200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMappingsActions200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

