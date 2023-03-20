# OidcAppAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | [**ConfigurationOidc**](ConfigurationOidc.md) |  | 
**Sso** | Pointer to [**SsoOidc**](SsoOidc.md) |  | [optional] 

## Methods

### NewOidcAppAllOf

`func NewOidcAppAllOf(configuration ConfigurationOidc, ) *OidcAppAllOf`

NewOidcAppAllOf instantiates a new OidcAppAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOidcAppAllOfWithDefaults

`func NewOidcAppAllOfWithDefaults() *OidcAppAllOf`

NewOidcAppAllOfWithDefaults instantiates a new OidcAppAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *OidcAppAllOf) GetConfiguration() ConfigurationOidc`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *OidcAppAllOf) GetConfigurationOk() (*ConfigurationOidc, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *OidcAppAllOf) SetConfiguration(v ConfigurationOidc)`

SetConfiguration sets Configuration field to given value.


### GetSso

`func (o *OidcAppAllOf) GetSso() SsoOidc`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *OidcAppAllOf) GetSsoOk() (*SsoOidc, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *OidcAppAllOf) SetSso(v SsoOidc)`

SetSso sets Sso field to given value.

### HasSso

`func (o *OidcAppAllOf) HasSso() bool`

HasSso returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


