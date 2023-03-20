# AuthClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The attribute name for the claim when its returned in an Access Token | 
**UserAttributeMappings** | Pointer to **string** | A user attribute to map values from For custom attributes prefix the name of the attribute with &#x60;custom_attribute_&#x60;. e.g. To get the value for custom attribute &#x60;employee_id&#x60; use &#x60;custom_attribute_employee_id&#x60;. | [optional] 
**UserAttributeMacros** | Pointer to **string** | When &#x60;user_attribute_mappings&#x60; is set to &#x60;_macro_&#x60; this macro will be used to assign the parameter value. | [optional] 

## Methods

### NewAuthClaim

`func NewAuthClaim(name string, ) *AuthClaim`

NewAuthClaim instantiates a new AuthClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthClaimWithDefaults

`func NewAuthClaimWithDefaults() *AuthClaim`

NewAuthClaimWithDefaults instantiates a new AuthClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthClaim) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthClaim) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthClaim) SetName(v string)`

SetName sets Name field to given value.


### GetUserAttributeMappings

`func (o *AuthClaim) GetUserAttributeMappings() string`

GetUserAttributeMappings returns the UserAttributeMappings field if non-nil, zero value otherwise.

### GetUserAttributeMappingsOk

`func (o *AuthClaim) GetUserAttributeMappingsOk() (*string, bool)`

GetUserAttributeMappingsOk returns a tuple with the UserAttributeMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeMappings

`func (o *AuthClaim) SetUserAttributeMappings(v string)`

SetUserAttributeMappings sets UserAttributeMappings field to given value.

### HasUserAttributeMappings

`func (o *AuthClaim) HasUserAttributeMappings() bool`

HasUserAttributeMappings returns a boolean if a field has been set.

### GetUserAttributeMacros

`func (o *AuthClaim) GetUserAttributeMacros() string`

GetUserAttributeMacros returns the UserAttributeMacros field if non-nil, zero value otherwise.

### GetUserAttributeMacrosOk

`func (o *AuthClaim) GetUserAttributeMacrosOk() (*string, bool)`

GetUserAttributeMacrosOk returns a tuple with the UserAttributeMacros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAttributeMacros

`func (o *AuthClaim) SetUserAttributeMacros(v string)`

SetUserAttributeMacros sets UserAttributeMacros field to given value.

### HasUserAttributeMacros

`func (o *AuthClaim) HasUserAttributeMacros() bool`

HasUserAttributeMacros returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


