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

// checks if the ListMappingConditions200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMappingConditions200Response{}

// ListMappingConditions200Response struct for ListMappingConditions200Response
type ListMappingConditions200Response struct {
	// Name of Condition
	Name *string `json:"name,omitempty"`
	// The unique identifier of the condition. This should be used when defining conditions for a User Mapping
	Value *string `json:"value,omitempty"`
}

// NewListMappingConditions200Response instantiates a new ListMappingConditions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMappingConditions200Response() *ListMappingConditions200Response {
	this := ListMappingConditions200Response{}
	return &this
}

// NewListMappingConditions200ResponseWithDefaults instantiates a new ListMappingConditions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMappingConditions200ResponseWithDefaults() *ListMappingConditions200Response {
	this := ListMappingConditions200Response{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListMappingConditions200Response) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMappingConditions200Response) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListMappingConditions200Response) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListMappingConditions200Response) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ListMappingConditions200Response) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMappingConditions200Response) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ListMappingConditions200Response) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ListMappingConditions200Response) SetValue(v string) {
	o.Value = &v
}

func (o ListMappingConditions200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMappingConditions200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableListMappingConditions200Response struct {
	value *ListMappingConditions200Response
	isSet bool
}

func (v NullableListMappingConditions200Response) Get() *ListMappingConditions200Response {
	return v.value
}

func (v *NullableListMappingConditions200Response) Set(val *ListMappingConditions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListMappingConditions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListMappingConditions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMappingConditions200Response(val *ListMappingConditions200Response) *NullableListMappingConditions200Response {
	return &NullableListMappingConditions200Response{value: val, isSet: true}
}

func (v NullableListMappingConditions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMappingConditions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


