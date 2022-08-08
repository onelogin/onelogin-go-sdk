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

// Connector struct for Connector
type Connector struct {
	// The connectors unique ID in OneLogin.
	Id *int32 `json:"id,omitempty"`
	// The name of the connector.
	Name *string `json:"name,omitempty"`
	AuthMethod *AuthMethod `json:"auth_method,omitempty"`
	// Indicates if apps created using this connector will be allowed to create custom parameters.
	AllowsNewParameters *bool `json:"allows_new_parameters,omitempty"`
	// A link to the apps icon url.
	IconUrl *string `json:"icon_url,omitempty"`
}

// NewConnector instantiates a new Connector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnector() *Connector {
	this := Connector{}
	return &this
}

// NewConnectorWithDefaults instantiates a new Connector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorWithDefaults() *Connector {
	this := Connector{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Connector) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connector) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Connector) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Connector) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Connector) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connector) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Connector) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Connector) SetName(v string) {
	o.Name = &v
}

// GetAuthMethod returns the AuthMethod field value if set, zero value otherwise.
func (o *Connector) GetAuthMethod() AuthMethod {
	if o == nil || o.AuthMethod == nil {
		var ret AuthMethod
		return ret
	}
	return *o.AuthMethod
}

// GetAuthMethodOk returns a tuple with the AuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connector) GetAuthMethodOk() (*AuthMethod, bool) {
	if o == nil || o.AuthMethod == nil {
		return nil, false
	}
	return o.AuthMethod, true
}

// HasAuthMethod returns a boolean if a field has been set.
func (o *Connector) HasAuthMethod() bool {
	if o != nil && o.AuthMethod != nil {
		return true
	}

	return false
}

// SetAuthMethod gets a reference to the given AuthMethod and assigns it to the AuthMethod field.
func (o *Connector) SetAuthMethod(v AuthMethod) {
	o.AuthMethod = &v
}

// GetAllowsNewParameters returns the AllowsNewParameters field value if set, zero value otherwise.
func (o *Connector) GetAllowsNewParameters() bool {
	if o == nil || o.AllowsNewParameters == nil {
		var ret bool
		return ret
	}
	return *o.AllowsNewParameters
}

// GetAllowsNewParametersOk returns a tuple with the AllowsNewParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connector) GetAllowsNewParametersOk() (*bool, bool) {
	if o == nil || o.AllowsNewParameters == nil {
		return nil, false
	}
	return o.AllowsNewParameters, true
}

// HasAllowsNewParameters returns a boolean if a field has been set.
func (o *Connector) HasAllowsNewParameters() bool {
	if o != nil && o.AllowsNewParameters != nil {
		return true
	}

	return false
}

// SetAllowsNewParameters gets a reference to the given bool and assigns it to the AllowsNewParameters field.
func (o *Connector) SetAllowsNewParameters(v bool) {
	o.AllowsNewParameters = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *Connector) GetIconUrl() string {
	if o == nil || o.IconUrl == nil {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connector) GetIconUrlOk() (*string, bool) {
	if o == nil || o.IconUrl == nil {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *Connector) HasIconUrl() bool {
	if o != nil && o.IconUrl != nil {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *Connector) SetIconUrl(v string) {
	o.IconUrl = &v
}

func (o Connector) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.AuthMethod != nil {
		toSerialize["auth_method"] = o.AuthMethod
	}
	if o.AllowsNewParameters != nil {
		toSerialize["allows_new_parameters"] = o.AllowsNewParameters
	}
	if o.IconUrl != nil {
		toSerialize["icon_url"] = o.IconUrl
	}
	return json.Marshal(toSerialize)
}

type NullableConnector struct {
	value *Connector
	isSet bool
}

func (v NullableConnector) Get() *Connector {
	return v.value
}

func (v *NullableConnector) Set(val *Connector) {
	v.value = val
	v.isSet = true
}

func (v NullableConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnector(val *Connector) *NullableConnector {
	return &NullableConnector{value: val, isSet: true}
}

func (v NullableConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


