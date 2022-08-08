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

// ListAppUsers200ResponseInner struct for ListAppUsers200ResponseInner
type ListAppUsers200ResponseInner struct {
	Firstname *string `json:"firstname,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Email *string `json:"email,omitempty"`
	Lastname *string `json:"lastname,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewListAppUsers200ResponseInner instantiates a new ListAppUsers200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAppUsers200ResponseInner() *ListAppUsers200ResponseInner {
	this := ListAppUsers200ResponseInner{}
	return &this
}

// NewListAppUsers200ResponseInnerWithDefaults instantiates a new ListAppUsers200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAppUsers200ResponseInnerWithDefaults() *ListAppUsers200ResponseInner {
	this := ListAppUsers200ResponseInner{}
	return &this
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *ListAppUsers200ResponseInner) GetFirstname() string {
	if o == nil || o.Firstname == nil {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAppUsers200ResponseInner) GetFirstnameOk() (*string, bool) {
	if o == nil || o.Firstname == nil {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *ListAppUsers200ResponseInner) HasFirstname() bool {
	if o != nil && o.Firstname != nil {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *ListAppUsers200ResponseInner) SetFirstname(v string) {
	o.Firstname = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListAppUsers200ResponseInner) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAppUsers200ResponseInner) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListAppUsers200ResponseInner) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ListAppUsers200ResponseInner) SetId(v int32) {
	o.Id = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ListAppUsers200ResponseInner) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAppUsers200ResponseInner) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ListAppUsers200ResponseInner) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ListAppUsers200ResponseInner) SetEmail(v string) {
	o.Email = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *ListAppUsers200ResponseInner) GetLastname() string {
	if o == nil || o.Lastname == nil {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAppUsers200ResponseInner) GetLastnameOk() (*string, bool) {
	if o == nil || o.Lastname == nil {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *ListAppUsers200ResponseInner) HasLastname() bool {
	if o != nil && o.Lastname != nil {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *ListAppUsers200ResponseInner) SetLastname(v string) {
	o.Lastname = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ListAppUsers200ResponseInner) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAppUsers200ResponseInner) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ListAppUsers200ResponseInner) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ListAppUsers200ResponseInner) SetUsername(v string) {
	o.Username = &v
}

func (o ListAppUsers200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Firstname != nil {
		toSerialize["firstname"] = o.Firstname
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Lastname != nil {
		toSerialize["lastname"] = o.Lastname
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableListAppUsers200ResponseInner struct {
	value *ListAppUsers200ResponseInner
	isSet bool
}

func (v NullableListAppUsers200ResponseInner) Get() *ListAppUsers200ResponseInner {
	return v.value
}

func (v *NullableListAppUsers200ResponseInner) Set(val *ListAppUsers200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListAppUsers200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListAppUsers200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAppUsers200ResponseInner(val *ListAppUsers200ResponseInner) *NullableListAppUsers200ResponseInner {
	return &NullableListAppUsers200ResponseInner{value: val, isSet: true}
}

func (v NullableListAppUsers200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAppUsers200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


