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

// checks if the ListMessageTemplates200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMessageTemplates200ResponseInner{}

// ListMessageTemplates200ResponseInner struct for ListMessageTemplates200ResponseInner
type ListMessageTemplates200ResponseInner struct {
	// template ID
	Id *int32 `json:"id,omitempty"`
	// indicator if template is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// name of message template
	Name *string `json:"name,omitempty"`
}

// NewListMessageTemplates200ResponseInner instantiates a new ListMessageTemplates200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMessageTemplates200ResponseInner() *ListMessageTemplates200ResponseInner {
	this := ListMessageTemplates200ResponseInner{}
	return &this
}

// NewListMessageTemplates200ResponseInnerWithDefaults instantiates a new ListMessageTemplates200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMessageTemplates200ResponseInnerWithDefaults() *ListMessageTemplates200ResponseInner {
	this := ListMessageTemplates200ResponseInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListMessageTemplates200ResponseInner) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMessageTemplates200ResponseInner) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListMessageTemplates200ResponseInner) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ListMessageTemplates200ResponseInner) SetId(v int32) {
	o.Id = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ListMessageTemplates200ResponseInner) GetEnabled() bool {
	if o == nil || isNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMessageTemplates200ResponseInner) GetEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ListMessageTemplates200ResponseInner) HasEnabled() bool {
	if o != nil && !isNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ListMessageTemplates200ResponseInner) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListMessageTemplates200ResponseInner) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMessageTemplates200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListMessageTemplates200ResponseInner) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListMessageTemplates200ResponseInner) SetName(v string) {
	o.Name = &v
}

func (o ListMessageTemplates200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMessageTemplates200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableListMessageTemplates200ResponseInner struct {
	value *ListMessageTemplates200ResponseInner
	isSet bool
}

func (v NullableListMessageTemplates200ResponseInner) Get() *ListMessageTemplates200ResponseInner {
	return v.value
}

func (v *NullableListMessageTemplates200ResponseInner) Set(val *ListMessageTemplates200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListMessageTemplates200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListMessageTemplates200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMessageTemplates200ResponseInner(val *ListMessageTemplates200ResponseInner) *NullableListMessageTemplates200ResponseInner {
	return &NullableListMessageTemplates200ResponseInner{value: val, isSet: true}
}

func (v NullableListMessageTemplates200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMessageTemplates200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

