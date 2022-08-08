# AuthServerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audiences** | Pointer to **[]string** |  | [optional] 
**RefreshTokenExpirationMinutes** | Pointer to **int32** |  | [optional] 
**ResourceIdentifier** | Pointer to **string** |  | [optional] 
**AccessTokenExpirationMinutes** | Pointer to **int32** |  | [optional] 

## Methods

### NewAuthServerConfiguration

`func NewAuthServerConfiguration() *AuthServerConfiguration`

NewAuthServerConfiguration instantiates a new AuthServerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthServerConfigurationWithDefaults

`func NewAuthServerConfigurationWithDefaults() *AuthServerConfiguration`

NewAuthServerConfigurationWithDefaults instantiates a new AuthServerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudiences

`func (o *AuthServerConfiguration) GetAudiences() []string`

GetAudiences returns the Audiences field if non-nil, zero value otherwise.

### GetAudiencesOk

`func (o *AuthServerConfiguration) GetAudiencesOk() (*[]string, bool)`

GetAudiencesOk returns a tuple with the Audiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudiences

`func (o *AuthServerConfiguration) SetAudiences(v []string)`

SetAudiences sets Audiences field to given value.

### HasAudiences

`func (o *AuthServerConfiguration) HasAudiences() bool`

HasAudiences returns a boolean if a field has been set.

### GetRefreshTokenExpirationMinutes

`func (o *AuthServerConfiguration) GetRefreshTokenExpirationMinutes() int32`

GetRefreshTokenExpirationMinutes returns the RefreshTokenExpirationMinutes field if non-nil, zero value otherwise.

### GetRefreshTokenExpirationMinutesOk

`func (o *AuthServerConfiguration) GetRefreshTokenExpirationMinutesOk() (*int32, bool)`

GetRefreshTokenExpirationMinutesOk returns a tuple with the RefreshTokenExpirationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenExpirationMinutes

`func (o *AuthServerConfiguration) SetRefreshTokenExpirationMinutes(v int32)`

SetRefreshTokenExpirationMinutes sets RefreshTokenExpirationMinutes field to given value.

### HasRefreshTokenExpirationMinutes

`func (o *AuthServerConfiguration) HasRefreshTokenExpirationMinutes() bool`

HasRefreshTokenExpirationMinutes returns a boolean if a field has been set.

### GetResourceIdentifier

`func (o *AuthServerConfiguration) GetResourceIdentifier() string`

GetResourceIdentifier returns the ResourceIdentifier field if non-nil, zero value otherwise.

### GetResourceIdentifierOk

`func (o *AuthServerConfiguration) GetResourceIdentifierOk() (*string, bool)`

GetResourceIdentifierOk returns a tuple with the ResourceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIdentifier

`func (o *AuthServerConfiguration) SetResourceIdentifier(v string)`

SetResourceIdentifier sets ResourceIdentifier field to given value.

### HasResourceIdentifier

`func (o *AuthServerConfiguration) HasResourceIdentifier() bool`

HasResourceIdentifier returns a boolean if a field has been set.

### GetAccessTokenExpirationMinutes

`func (o *AuthServerConfiguration) GetAccessTokenExpirationMinutes() int32`

GetAccessTokenExpirationMinutes returns the AccessTokenExpirationMinutes field if non-nil, zero value otherwise.

### GetAccessTokenExpirationMinutesOk

`func (o *AuthServerConfiguration) GetAccessTokenExpirationMinutesOk() (*int32, bool)`

GetAccessTokenExpirationMinutesOk returns a tuple with the AccessTokenExpirationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenExpirationMinutes

`func (o *AuthServerConfiguration) SetAccessTokenExpirationMinutes(v int32)`

SetAccessTokenExpirationMinutes sets AccessTokenExpirationMinutes field to given value.

### HasAccessTokenExpirationMinutes

`func (o *AuthServerConfiguration) HasAccessTokenExpirationMinutes() bool`

HasAccessTokenExpirationMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


