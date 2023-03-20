# Scope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Unique Scope ID value | [optional] 
**Value** | Pointer to **string** | Scope Value | [optional] 
**Description** | Pointer to **string** | Description of the scope | [optional] 

## Methods

### NewScope

`func NewScope() *Scope`

NewScope instantiates a new Scope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeWithDefaults

`func NewScopeWithDefaults() *Scope`

NewScopeWithDefaults instantiates a new Scope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Scope) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Scope) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Scope) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Scope) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *Scope) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Scope) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Scope) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Scope) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDescription

`func (o *Scope) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Scope) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Scope) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Scope) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


