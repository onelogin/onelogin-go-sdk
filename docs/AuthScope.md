# AuthScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique ID for the Scope | [optional] [readonly] 
**Value** | Pointer to **string** | A value representing the api scope that with be authorized | [optional] 
**Description** | Pointer to **string** | A description of what access the scope enables | [optional] 

## Methods

### NewAuthScope

`func NewAuthScope() *AuthScope`

NewAuthScope instantiates a new AuthScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthScopeWithDefaults

`func NewAuthScopeWithDefaults() *AuthScope`

NewAuthScopeWithDefaults instantiates a new AuthScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthScope) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthScope) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthScope) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuthScope) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *AuthScope) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AuthScope) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AuthScope) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *AuthScope) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDescription

`func (o *AuthScope) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthScope) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthScope) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AuthScope) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


