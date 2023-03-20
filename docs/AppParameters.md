# AppParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserAttributeMappings** | Pointer to **string** | A user attribute to map values from For custom attributes prefix the name of the attribute with &#x60;custom_attribute_&#x60;. e.g. To get the value for custom attribute &#x60;employee_id&#x60; use &#x60;custom_attribute_employee_id&#x60;. | [optional] 
**UserAttributeMacros** | Pointer to **string** | When &#x60;user_attribute_mappings&#x60; is set to &#x60;_macro_&#x60; this macro will be used to assign the parameter value. | [optional] 
**Label** | Pointer to **string** | The can only be set when creating a new parameter. It can not be updated. | [optional] 
**IncludeInSamlAssertion** | Pointer to **bool** | When true, this parameter will be included in a SAML assertion payload. | [optional] 

## Methods

### NewAppParameters

`func NewAppParameters() *AppParameters`

NewAppParameters instantiates a new AppParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppParametersWithDefaults

`func NewAppParametersWithDefaults() *AppParameters`

NewAppParametersWithDefaults instantiates a new AppParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserAttributeMappings

`func (o *AppParameters) GetUserAttributeMappings() string`

GetUserAttributeMappings returns the UserAttributeMappings field if non-nil, zero value otherwise.

### GetUserAttributeMappingsOk

`func (o *AppParameters) GetUserAttributeMappingsOk() (*string, bool)`

GetUserAttributeMappingsOk returns a tuple with the UserAttributeMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeMappings

`func (o *AppParameters) SetUserAttributeMappings(v string)`

SetUserAttributeMappings sets UserAttributeMappings field to given value.

### HasUserAttributeMappings

`func (o *AppParameters) HasUserAttributeMappings() bool`

HasUserAttributeMappings returns a boolean if a field has been set.

### GetUserAttributeMacros

`func (o *AppParameters) GetUserAttributeMacros() string`

GetUserAttributeMacros returns the UserAttributeMacros field if non-nil, zero value otherwise.

### GetUserAttributeMacrosOk

`func (o *AppParameters) GetUserAttributeMacrosOk() (*string, bool)`

GetUserAttributeMacrosOk returns a tuple with the UserAttributeMacros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeMacros

`func (o *AppParameters) SetUserAttributeMacros(v string)`

SetUserAttributeMacros sets UserAttributeMacros field to given value.

### HasUserAttributeMacros

`func (o *AppParameters) HasUserAttributeMacros() bool`

HasUserAttributeMacros returns a boolean if a field has been set.

### GetLabel

`func (o *AppParameters) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AppParameters) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AppParameters) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AppParameters) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetIncludeInSamlAssertion

`func (o *AppParameters) GetIncludeInSamlAssertion() bool`

GetIncludeInSamlAssertion returns the IncludeInSamlAssertion field if non-nil, zero value otherwise.

### GetIncludeInSamlAssertionOk

`func (o *AppParameters) GetIncludeInSamlAssertionOk() (*bool, bool)`

GetIncludeInSamlAssertionOk returns a tuple with the IncludeInSamlAssertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInSamlAssertion

`func (o *AppParameters) SetIncludeInSamlAssertion(v bool)`

SetIncludeInSamlAssertion sets IncludeInSamlAssertion field to given value.

### HasIncludeInSamlAssertion

`func (o *AppParameters) HasIncludeInSamlAssertion() bool`

HasIncludeInSamlAssertion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


