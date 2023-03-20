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

// checks if the AppParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppParameters{}

// AppParameters The parameters section contains parameterized attributes that have defined at the connector level as well as custom attributes that have been defined specifically for this app. Regardless of how they are defined, all parameters have the following attributes. Each parameter is an object with the key for the object being set as the parameters short name.
type AppParameters struct {
	// A user attribute to map values from For custom attributes prefix the name of the attribute with `custom_attribute_`. e.g. To get the value for custom attribute `employee_id` use `custom_attribute_employee_id`.
	UserAttributeMappings *string `json:"user_attribute_mappings,omitempty"`
	// When `user_attribute_mappings` is set to `_macro_` this macro will be used to assign the parameter value.
	UserAttributeMacros *string `json:"user_attribute_macros,omitempty"`
	// The can only be set when creating a new parameter. It can not be updated.
	Label *string `json:"label,omitempty"`
	// When true, this parameter will be included in a SAML assertion payload.
	IncludeInSamlAssertion *bool `json:"include_in_saml_assertion,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AppParameters AppParameters

// NewAppParameters instantiates a new AppParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppParameters() *AppParameters {
	this := AppParameters{}
	return &this
}

// NewAppParametersWithDefaults instantiates a new AppParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppParametersWithDefaults() *AppParameters {
	this := AppParameters{}
	return &this
}

// GetUserAttributeMappings returns the UserAttributeMappings field value if set, zero value otherwise.
func (o *AppParameters) GetUserAttributeMappings() string {
	if o == nil || isNil(o.UserAttributeMappings) {
		var ret string
		return ret
	}
	return *o.UserAttributeMappings
}

// GetUserAttributeMappingsOk returns a tuple with the UserAttributeMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppParameters) GetUserAttributeMappingsOk() (*string, bool) {
	if o == nil || isNil(o.UserAttributeMappings) {
		return nil, false
	}
	return o.UserAttributeMappings, true
}

// HasUserAttributeMappings returns a boolean if a field has been set.
func (o *AppParameters) HasUserAttributeMappings() bool {
	if o != nil && !isNil(o.UserAttributeMappings) {
		return true
	}

	return false
}

// SetUserAttributeMappings gets a reference to the given string and assigns it to the UserAttributeMappings field.
func (o *AppParameters) SetUserAttributeMappings(v string) {
	o.UserAttributeMappings = &v
}

// GetUserAttributeMacros returns the UserAttributeMacros field value if set, zero value otherwise.
func (o *AppParameters) GetUserAttributeMacros() string {
	if o == nil || isNil(o.UserAttributeMacros) {
		var ret string
		return ret
	}
	return *o.UserAttributeMacros
}

// GetUserAttributeMacrosOk returns a tuple with the UserAttributeMacros field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppParameters) GetUserAttributeMacrosOk() (*string, bool) {
	if o == nil || isNil(o.UserAttributeMacros) {
		return nil, false
	}
	return o.UserAttributeMacros, true
}

// HasUserAttributeMacros returns a boolean if a field has been set.
func (o *AppParameters) HasUserAttributeMacros() bool {
	if o != nil && !isNil(o.UserAttributeMacros) {
		return true
	}

	return false
}

// SetUserAttributeMacros gets a reference to the given string and assigns it to the UserAttributeMacros field.
func (o *AppParameters) SetUserAttributeMacros(v string) {
	o.UserAttributeMacros = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *AppParameters) GetLabel() string {
	if o == nil || isNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppParameters) GetLabelOk() (*string, bool) {
	if o == nil || isNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *AppParameters) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *AppParameters) SetLabel(v string) {
	o.Label = &v
}

// GetIncludeInSamlAssertion returns the IncludeInSamlAssertion field value if set, zero value otherwise.
func (o *AppParameters) GetIncludeInSamlAssertion() bool {
	if o == nil || isNil(o.IncludeInSamlAssertion) {
		var ret bool
		return ret
	}
	return *o.IncludeInSamlAssertion
}

// GetIncludeInSamlAssertionOk returns a tuple with the IncludeInSamlAssertion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppParameters) GetIncludeInSamlAssertionOk() (*bool, bool) {
	if o == nil || isNil(o.IncludeInSamlAssertion) {
		return nil, false
	}
	return o.IncludeInSamlAssertion, true
}

// HasIncludeInSamlAssertion returns a boolean if a field has been set.
func (o *AppParameters) HasIncludeInSamlAssertion() bool {
	if o != nil && !isNil(o.IncludeInSamlAssertion) {
		return true
	}

	return false
}

// SetIncludeInSamlAssertion gets a reference to the given bool and assigns it to the IncludeInSamlAssertion field.
func (o *AppParameters) SetIncludeInSamlAssertion(v bool) {
	o.IncludeInSamlAssertion = &v
}

func (o AppParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UserAttributeMappings) {
		toSerialize["user_attribute_mappings"] = o.UserAttributeMappings
	}
	if !isNil(o.UserAttributeMacros) {
		toSerialize["user_attribute_macros"] = o.UserAttributeMacros
	}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.IncludeInSamlAssertion) {
		toSerialize["include_in_saml_assertion"] = o.IncludeInSamlAssertion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AppParameters) UnmarshalJSON(bytes []byte) (err error) {
	varAppParameters := _AppParameters{}

	if err = json.Unmarshal(bytes, &varAppParameters); err == nil {
		*o = AppParameters(varAppParameters)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "user_attribute_mappings")
		delete(additionalProperties, "user_attribute_macros")
		delete(additionalProperties, "label")
		delete(additionalProperties, "include_in_saml_assertion")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAppParameters struct {
	value *AppParameters
	isSet bool
}

func (v NullableAppParameters) Get() *AppParameters {
	return v.value
}

func (v *NullableAppParameters) Set(val *AppParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableAppParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableAppParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppParameters(val *AppParameters) *NullableAppParameters {
	return &NullableAppParameters{value: val, isSet: true}
}

func (v NullableAppParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

