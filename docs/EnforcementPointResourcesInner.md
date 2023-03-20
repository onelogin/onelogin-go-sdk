# EnforcementPointResourcesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**IsPathRegex** | Pointer to **NullableBool** |  | [optional] 
**RequireAuth** | Pointer to **bool** |  | [optional] 
**Permission** | Pointer to **string** |  | [optional] 
**Conditions** | Pointer to **string** | required if permission &#x3D;&#x3D; \&quot;conditions\&quot; | [optional] 

## Methods

### NewEnforcementPointResourcesInner

`func NewEnforcementPointResourcesInner() *EnforcementPointResourcesInner`

NewEnforcementPointResourcesInner instantiates a new EnforcementPointResourcesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnforcementPointResourcesInnerWithDefaults

`func NewEnforcementPointResourcesInnerWithDefaults() *EnforcementPointResourcesInner`

NewEnforcementPointResourcesInnerWithDefaults instantiates a new EnforcementPointResourcesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *EnforcementPointResourcesInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *EnforcementPointResourcesInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *EnforcementPointResourcesInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *EnforcementPointResourcesInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetIsPathRegex

`func (o *EnforcementPointResourcesInner) GetIsPathRegex() bool`

GetIsPathRegex returns the IsPathRegex field if non-nil, zero value otherwise.

### GetIsPathRegexOk

`func (o *EnforcementPointResourcesInner) GetIsPathRegexOk() (*bool, bool)`

GetIsPathRegexOk returns a tuple with the IsPathRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPathRegex

`func (o *EnforcementPointResourcesInner) SetIsPathRegex(v bool)`

SetIsPathRegex sets IsPathRegex field to given value.

### HasIsPathRegex

`func (o *EnforcementPointResourcesInner) HasIsPathRegex() bool`

HasIsPathRegex returns a boolean if a field has been set.

### SetIsPathRegexNil

`func (o *EnforcementPointResourcesInner) SetIsPathRegexNil(b bool)`

 SetIsPathRegexNil sets the value for IsPathRegex to be an explicit nil

### UnsetIsPathRegex
`func (o *EnforcementPointResourcesInner) UnsetIsPathRegex()`

UnsetIsPathRegex ensures that no value is present for IsPathRegex, not even an explicit nil
### GetRequireAuth

`func (o *EnforcementPointResourcesInner) GetRequireAuth() bool`

GetRequireAuth returns the RequireAuth field if non-nil, zero value otherwise.

### GetRequireAuthOk

`func (o *EnforcementPointResourcesInner) GetRequireAuthOk() (*bool, bool)`

GetRequireAuthOk returns a tuple with the RequireAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireAuth

`func (o *EnforcementPointResourcesInner) SetRequireAuth(v bool)`

SetRequireAuth sets RequireAuth field to given value.

### HasRequireAuth

`func (o *EnforcementPointResourcesInner) HasRequireAuth() bool`

HasRequireAuth returns a boolean if a field has been set.

### GetPermission

`func (o *EnforcementPointResourcesInner) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *EnforcementPointResourcesInner) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *EnforcementPointResourcesInner) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *EnforcementPointResourcesInner) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### GetConditions

`func (o *EnforcementPointResourcesInner) GetConditions() string`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *EnforcementPointResourcesInner) GetConditionsOk() (*string, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *EnforcementPointResourcesInner) SetConditions(v string)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *EnforcementPointResourcesInner) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


