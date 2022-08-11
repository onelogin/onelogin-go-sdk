# CreateAuthorizationServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Configuration** | Pointer to [**AuthServerConfiguration**](AuthServerConfiguration.md) |  | [optional] 

## Methods

### NewCreateAuthorizationServerRequest

`func NewCreateAuthorizationServerRequest() *CreateAuthorizationServerRequest`

NewCreateAuthorizationServerRequest instantiates a new CreateAuthorizationServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAuthorizationServerRequestWithDefaults

`func NewCreateAuthorizationServerRequestWithDefaults() *CreateAuthorizationServerRequest`

NewCreateAuthorizationServerRequestWithDefaults instantiates a new CreateAuthorizationServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAuthorizationServerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAuthorizationServerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAuthorizationServerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateAuthorizationServerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateAuthorizationServerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAuthorizationServerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAuthorizationServerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAuthorizationServerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetConfiguration

`func (o *CreateAuthorizationServerRequest) GetConfiguration() AuthServerConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CreateAuthorizationServerRequest) GetConfigurationOk() (*AuthServerConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CreateAuthorizationServerRequest) SetConfiguration(v AuthServerConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *CreateAuthorizationServerRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


