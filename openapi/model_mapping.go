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

// Mapping struct for Mapping
type Mapping struct {
	Id *int32 `json:"id,omitempty"`
	// The name of the mapping.
	Name string `json:"name"`
	// Indicates if the mapping is enabled or not.
	Enabled bool `json:"enabled"`
	// Indicates how conditions should be matched.
	Match string `json:"match"`
	// Indicates the order of the mapping. When `null` this will default to last position.
	Position int32 `json:"position"`
	// An array of conditions that the user must meet in order for the mapping to be applied.
	Conditions []Condition `json:"conditions,omitempty"`
	// An array of actions that will be applied to the users that are matched by the conditions.
	Actions []Action `json:"actions"`
}

// NewMapping instantiates a new Mapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMapping(name string, enabled bool, match string, position int32, actions []Action) *Mapping {
	this := Mapping{}
	this.Name = name
	this.Enabled = enabled
	this.Match = match
	this.Position = position
	this.Actions = actions
	return &this
}

// NewMappingWithDefaults instantiates a new Mapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMappingWithDefaults() *Mapping {
	this := Mapping{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Mapping) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mapping) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Mapping) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Mapping) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Mapping) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Mapping) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Mapping) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value
func (o *Mapping) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Mapping) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Mapping) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMatch returns the Match field value
func (o *Mapping) GetMatch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Match
}

// GetMatchOk returns a tuple with the Match field value
// and a boolean to check if the value has been set.
func (o *Mapping) GetMatchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Match, true
}

// SetMatch sets field value
func (o *Mapping) SetMatch(v string) {
	o.Match = v
}

// GetPosition returns the Position field value
func (o *Mapping) GetPosition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Position
}

// GetPositionOk returns a tuple with the Position field value
// and a boolean to check if the value has been set.
func (o *Mapping) GetPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Position, true
}

// SetPosition sets field value
func (o *Mapping) SetPosition(v int32) {
	o.Position = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *Mapping) GetConditions() []Condition {
	if o == nil || o.Conditions == nil {
		var ret []Condition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mapping) GetConditionsOk() ([]Condition, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *Mapping) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []Condition and assigns it to the Conditions field.
func (o *Mapping) SetConditions(v []Condition) {
	o.Conditions = v
}

// GetActions returns the Actions field value
func (o *Mapping) GetActions() []Action {
	if o == nil {
		var ret []Action
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *Mapping) GetActionsOk() ([]Action, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *Mapping) SetActions(v []Action) {
	o.Actions = v
}

func (o Mapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["match"] = o.Match
	}
	if true {
		toSerialize["position"] = o.Position
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if true {
		toSerialize["actions"] = o.Actions
	}
	return json.Marshal(toSerialize)
}

type NullableMapping struct {
	value *Mapping
	isSet bool
}

func (v NullableMapping) Get() *Mapping {
	return v.value
}

func (v *NullableMapping) Set(val *Mapping) {
	v.value = val
	v.isSet = true
}

func (v NullableMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMapping(val *Mapping) *NullableMapping {
	return &NullableMapping{value: val, isSet: true}
}

func (v NullableMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


