# ClientAppFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scopes** | Pointer to [**[]Scope**](Scope.md) | List of All Scopes assigned to a client app | [optional] 
**AppId** | Pointer to **int32** | Unique Client App ID | [optional] 
**Name** | Pointer to **string** | Name of client app | [optional] 
**ApiAuthId** | Pointer to **int32** |  | [optional] 

## Methods

### NewClientAppFull

`func NewClientAppFull() *ClientAppFull`

NewClientAppFull instantiates a new ClientAppFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAppFullWithDefaults

`func NewClientAppFullWithDefaults() *ClientAppFull`

NewClientAppFullWithDefaults instantiates a new ClientAppFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopes

`func (o *ClientAppFull) GetScopes() []Scope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ClientAppFull) GetScopesOk() (*[]Scope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ClientAppFull) SetScopes(v []Scope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ClientAppFull) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetAppId

`func (o *ClientAppFull) GetAppId() int32`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ClientAppFull) GetAppIdOk() (*int32, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ClientAppFull) SetAppId(v int32)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ClientAppFull) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetName

`func (o *ClientAppFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientAppFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientAppFull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientAppFull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApiAuthId

`func (o *ClientAppFull) GetApiAuthId() int32`

GetApiAuthId returns the ApiAuthId field if non-nil, zero value otherwise.

### GetApiAuthIdOk

`func (o *ClientAppFull) GetApiAuthIdOk() (*int32, bool)`

GetApiAuthIdOk returns a tuple with the ApiAuthId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAuthId

`func (o *ClientAppFull) SetApiAuthId(v int32)`

SetApiAuthId sets ApiAuthId field to given value.

### HasApiAuthId

`func (o *ClientAppFull) HasApiAuthId() bool`

HasApiAuthId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


