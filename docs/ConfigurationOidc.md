# ConfigurationOidc

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginUrl** | **string** | The OpenId Connect Client Id. Note that client_secret is only returned after Creating an App. | 
**RedirectUri** | **string** | Comma or newline separated list of valid redirect uris for the OpenId Connect Authorization Code flow. | 
**AccessTokenExpirationMinutes** | **int32** | Number of minutes the refresh token will be valid for. | 
**RefreshTokenExpirationMinutes** | **int32** | Number of minutes the refresh token will be valid for. | 
**TokenEndpointAuthMethod** | **int32** | - 0: Basic - 1: POST - 2: None / PKCE | 
**OidcApplicationType** | **int32** | - 0 : Web - 1 : Native / Mobile | 

## Methods

### NewConfigurationOidc

`func NewConfigurationOidc(loginUrl string, redirectUri string, accessTokenExpirationMinutes int32, refreshTokenExpirationMinutes int32, tokenEndpointAuthMethod int32, oidcApplicationType int32, ) *ConfigurationOidc`

NewConfigurationOidc instantiates a new ConfigurationOidc object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigurationOidcWithDefaults

`func NewConfigurationOidcWithDefaults() *ConfigurationOidc`

NewConfigurationOidcWithDefaults instantiates a new ConfigurationOidc object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginUrl

`func (o *ConfigurationOidc) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *ConfigurationOidc) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *ConfigurationOidc) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.


### GetRedirectUri

`func (o *ConfigurationOidc) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *ConfigurationOidc) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *ConfigurationOidc) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.


### GetAccessTokenExpirationMinutes

`func (o *ConfigurationOidc) GetAccessTokenExpirationMinutes() int32`

GetAccessTokenExpirationMinutes returns the AccessTokenExpirationMinutes field if non-nil, zero value otherwise.

### GetAccessTokenExpirationMinutesOk

`func (o *ConfigurationOidc) GetAccessTokenExpirationMinutesOk() (*int32, bool)`

GetAccessTokenExpirationMinutesOk returns a tuple with the AccessTokenExpirationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokenExpirationMinutes

`func (o *ConfigurationOidc) SetAccessTokenExpirationMinutes(v int32)`

SetAccessTokenExpirationMinutes sets AccessTokenExpirationMinutes field to given value.


### GetRefreshTokenExpirationMinutes

`func (o *ConfigurationOidc) GetRefreshTokenExpirationMinutes() int32`

GetRefreshTokenExpirationMinutes returns the RefreshTokenExpirationMinutes field if non-nil, zero value otherwise.

### GetRefreshTokenExpirationMinutesOk

`func (o *ConfigurationOidc) GetRefreshTokenExpirationMinutesOk() (*int32, bool)`

GetRefreshTokenExpirationMinutesOk returns a tuple with the RefreshTokenExpirationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenExpirationMinutes

`func (o *ConfigurationOidc) SetRefreshTokenExpirationMinutes(v int32)`

SetRefreshTokenExpirationMinutes sets RefreshTokenExpirationMinutes field to given value.


### GetTokenEndpointAuthMethod

`func (o *ConfigurationOidc) GetTokenEndpointAuthMethod() int32`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *ConfigurationOidc) GetTokenEndpointAuthMethodOk() (*int32, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *ConfigurationOidc) SetTokenEndpointAuthMethod(v int32)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.


### GetOidcApplicationType

`func (o *ConfigurationOidc) GetOidcApplicationType() int32`

GetOidcApplicationType returns the OidcApplicationType field if non-nil, zero value otherwise.

### GetOidcApplicationTypeOk

`func (o *ConfigurationOidc) GetOidcApplicationTypeOk() (*int32, bool)`

GetOidcApplicationTypeOk returns a tuple with the OidcApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcApplicationType

`func (o *ConfigurationOidc) SetOidcApplicationType(v int32)`

SetOidcApplicationType sets OidcApplicationType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


