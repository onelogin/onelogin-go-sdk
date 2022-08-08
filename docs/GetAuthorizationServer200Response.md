# GetAuthorizationServer200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**AuthServerConfiguration**](AuthServerConfiguration.md) |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewGetAuthorizationServer200Response

`func NewGetAuthorizationServer200Response() *GetAuthorizationServer200Response`

NewGetAuthorizationServer200Response instantiates a new GetAuthorizationServer200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAuthorizationServer200ResponseWithDefaults

`func NewGetAuthorizationServer200ResponseWithDefaults() *GetAuthorizationServer200Response`

NewGetAuthorizationServer200ResponseWithDefaults instantiates a new GetAuthorizationServer200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *GetAuthorizationServer200Response) GetConfiguration() AuthServerConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *GetAuthorizationServer200Response) GetConfigurationOk() (*AuthServerConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *GetAuthorizationServer200Response) SetConfiguration(v AuthServerConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *GetAuthorizationServer200Response) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetId

`func (o *GetAuthorizationServer200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAuthorizationServer200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAuthorizationServer200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetAuthorizationServer200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDescription

`func (o *GetAuthorizationServer200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetAuthorizationServer200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetAuthorizationServer200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetAuthorizationServer200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *GetAuthorizationServer200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAuthorizationServer200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAuthorizationServer200Response) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAuthorizationServer200Response) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


