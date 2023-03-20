# UpdateClientAppRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scopes** | Pointer to **[]int32** | An array of Scope IDs the scopes the app can request | [optional] 

## Methods

### NewUpdateClientAppRequest

`func NewUpdateClientAppRequest() *UpdateClientAppRequest`

NewUpdateClientAppRequest instantiates a new UpdateClientAppRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClientAppRequestWithDefaults

`func NewUpdateClientAppRequestWithDefaults() *UpdateClientAppRequest`

NewUpdateClientAppRequestWithDefaults instantiates a new UpdateClientAppRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopes

`func (o *UpdateClientAppRequest) GetScopes() []int32`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *UpdateClientAppRequest) GetScopesOk() (*[]int32, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *UpdateClientAppRequest) SetScopes(v []int32)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *UpdateClientAppRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


