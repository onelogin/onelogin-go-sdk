# AddClientAppRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | Pointer to **int32** | The ID of the OpenId Connect app to allow access through. | [optional] 
**Scopes** | Pointer to **[]int32** | An array of Scope IDs that represent scopes the app can request | [optional] 

## Methods

### NewAddClientAppRequest

`func NewAddClientAppRequest() *AddClientAppRequest`

NewAddClientAppRequest instantiates a new AddClientAppRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClientAppRequestWithDefaults

`func NewAddClientAppRequestWithDefaults() *AddClientAppRequest`

NewAddClientAppRequestWithDefaults instantiates a new AddClientAppRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *AddClientAppRequest) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AddClientAppRequest) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AddClientAppRequest) SetAppId(v int32)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *AddClientAppRequest) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetScopes

`func (o *AddClientAppRequest) GetScopes() []int32`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *AddClientAppRequest) GetScopesOk() (*[]int32, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *AddClientAppRequest) SetScopes(v []int32)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *AddClientAppRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


