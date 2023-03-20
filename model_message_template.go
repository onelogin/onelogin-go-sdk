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

// checks if the MessageTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessageTemplate{}

// MessageTemplate struct for MessageTemplate
type MessageTemplate struct {
	Id *int32 `json:"id,omitempty"`
	AccountId *int32 `json:"account_id,omitempty"`
	// Template type that describes the source (sms, voice, email) and purpose (registration, invite, etc)
	Type string `json:"type"`
	// The 2 character language locale for the template. e.g. en = English, es = Spanish
	Locale string `json:"locale"`
	Template MessageTemplateTemplate `json:"template"`
	TemplateClass *string `json:"template_class,omitempty"`
	// Last time template was updated
	UpdatedAt *string `json:"updated_at,omitempty"`
	// brand id number
	BrandId *int32 `json:"brand_id,omitempty"`
}

// NewMessageTemplate instantiates a new MessageTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageTemplate(type_ string, locale string, template MessageTemplateTemplate) *MessageTemplate {
	this := MessageTemplate{}
	this.Type = type_
	this.Locale = locale
	this.Template = template
	return &this
}

// NewMessageTemplateWithDefaults instantiates a new MessageTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageTemplateWithDefaults() *MessageTemplate {
	this := MessageTemplate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MessageTemplate) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageTemplate) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MessageTemplate) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *MessageTemplate) SetId(v int32) {
	o.Id = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *MessageTemplate) GetAccountId() int32 {
	if o == nil || isNil(o.AccountId) {
		var ret int32
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageTemplate) GetAccountIdOk() (*int32, bool) {
	if o == nil || isNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *MessageTemplate) HasAccountId() bool {
	if o != nil && !isNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int32 and assigns it to the AccountId field.
func (o *MessageTemplate) SetAccountId(v int32) {
	o.AccountId = &v
}

// GetType returns the Type field value
func (o *MessageTemplate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MessageTemplate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MessageTemplate) SetType(v string) {
	o.Type = v
}

// GetLocale returns the Locale field value
func (o *MessageTemplate) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *MessageTemplate) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *MessageTemplate) SetLocale(v string) {
	o.Locale = v
}

// GetTemplate returns the Template field value
func (o *MessageTemplate) GetTemplate() MessageTemplateTemplate {
	if o == nil {
		var ret MessageTemplateTemplate
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *MessageTemplate) GetTemplateOk() (*MessageTemplateTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *MessageTemplate) SetTemplate(v MessageTemplateTemplate) {
	o.Template = v
}

// GetTemplateClass returns the TemplateClass field value if set, zero value otherwise.
func (o *MessageTemplate) GetTemplateClass() string {
	if o == nil || isNil(o.TemplateClass) {
		var ret string
		return ret
	}
	return *o.TemplateClass
}

// GetTemplateClassOk returns a tuple with the TemplateClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageTemplate) GetTemplateClassOk() (*string, bool) {
	if o == nil || isNil(o.TemplateClass) {
		return nil, false
	}
	return o.TemplateClass, true
}

// HasTemplateClass returns a boolean if a field has been set.
func (o *MessageTemplate) HasTemplateClass() bool {
	if o != nil && !isNil(o.TemplateClass) {
		return true
	}

	return false
}

// SetTemplateClass gets a reference to the given string and assigns it to the TemplateClass field.
func (o *MessageTemplate) SetTemplateClass(v string) {
	o.TemplateClass = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MessageTemplate) GetUpdatedAt() string {
	if o == nil || isNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageTemplate) GetUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MessageTemplate) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *MessageTemplate) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetBrandId returns the BrandId field value if set, zero value otherwise.
func (o *MessageTemplate) GetBrandId() int32 {
	if o == nil || isNil(o.BrandId) {
		var ret int32
		return ret
	}
	return *o.BrandId
}

// GetBrandIdOk returns a tuple with the BrandId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessageTemplate) GetBrandIdOk() (*int32, bool) {
	if o == nil || isNil(o.BrandId) {
		return nil, false
	}
	return o.BrandId, true
}

// HasBrandId returns a boolean if a field has been set.
func (o *MessageTemplate) HasBrandId() bool {
	if o != nil && !isNil(o.BrandId) {
		return true
	}

	return false
}

// SetBrandId gets a reference to the given int32 and assigns it to the BrandId field.
func (o *MessageTemplate) SetBrandId(v int32) {
	o.BrandId = &v
}

func (o MessageTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessageTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	// skip: account_id is readOnly
	toSerialize["type"] = o.Type
	toSerialize["locale"] = o.Locale
	toSerialize["template"] = o.Template
	// skip: template_class is readOnly
	// skip: updated_at is readOnly
	// skip: brand_id is readOnly
	return toSerialize, nil
}

type NullableMessageTemplate struct {
	value *MessageTemplate
	isSet bool
}

func (v NullableMessageTemplate) Get() *MessageTemplate {
	return v.value
}

func (v *NullableMessageTemplate) Set(val *MessageTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageTemplate(val *MessageTemplate) *NullableMessageTemplate {
	return &NullableMessageTemplate{value: val, isSet: true}
}

func (v NullableMessageTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

