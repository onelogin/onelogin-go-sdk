# AuthServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Auth server unique ID in Onelogin | [optional] [readonly] 
**Name** | **string** | Name of the API. | 
**Description** | **string** | Description of what the API does. | 
**Configuration** | [**AuthServerConfiguration**](AuthServerConfiguration.md) |  | 

## Methods

### NewAuthServer

`func NewAuthServer(name string, description string, configuration AuthServerConfiguration, ) *AuthServer`

NewAuthServer instantiates a new AuthServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthServerWithDefaults

`func NewAuthServerWithDefaults() *AuthServer`

NewAuthServerWithDefaults instantiates a new AuthServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthServer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthServer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthServer) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AuthServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AuthServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthServer) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AuthServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuthServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuthServer) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetConfiguration

`func (o *AuthServer) GetConfiguration() AuthServerConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *AuthServer) GetConfigurationOk() (*AuthServerConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *AuthServer) SetConfiguration(v AuthServerConfiguration)`

SetConfiguration sets Configuration field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


