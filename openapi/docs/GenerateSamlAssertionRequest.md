# GenerateSamlAssertionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsernameOrEmail** | **string** | Set this to the username or email of the OneLogin user accessing the app for which you want to generate a SAML token. | 
**Password** | **string** | Password of the OneLogin user accessing the app for which you want to generate a SAML token. | 
**AppId** | **string** | App ID of the app for which you want to generate a SAML token. This is the app ID in OneLogin. | 
**Subdomain** | **string** | Set to the subdomain of the OneLogin user accessing the app for which you want to generate a SAML token. | 
**IpAddress** | Pointer to **string** | Whitelisted IP address, if MFA is required and you need to honor IP address whitelisting defined in MFA policies. | [optional] 

## Methods

### NewGenerateSamlAssertionRequest

`func NewGenerateSamlAssertionRequest(usernameOrEmail string, password string, appId string, subdomain string, ) *GenerateSamlAssertionRequest`

NewGenerateSamlAssertionRequest instantiates a new GenerateSamlAssertionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateSamlAssertionRequestWithDefaults

`func NewGenerateSamlAssertionRequestWithDefaults() *GenerateSamlAssertionRequest`

NewGenerateSamlAssertionRequestWithDefaults instantiates a new GenerateSamlAssertionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsernameOrEmail

`func (o *GenerateSamlAssertionRequest) GetUsernameOrEmail() string`

GetUsernameOrEmail returns the UsernameOrEmail field if non-nil, zero value otherwise.

### GetUsernameOrEmailOk

`func (o *GenerateSamlAssertionRequest) GetUsernameOrEmailOk() (*string, bool)`

GetUsernameOrEmailOk returns a tuple with the UsernameOrEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsernameOrEmail

`func (o *GenerateSamlAssertionRequest) SetUsernameOrEmail(v string)`

SetUsernameOrEmail sets UsernameOrEmail field to given value.


### GetPassword

`func (o *GenerateSamlAssertionRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GenerateSamlAssertionRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GenerateSamlAssertionRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetAppId

`func (o *GenerateSamlAssertionRequest) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *GenerateSamlAssertionRequest) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *GenerateSamlAssertionRequest) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetSubdomain

`func (o *GenerateSamlAssertionRequest) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *GenerateSamlAssertionRequest) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *GenerateSamlAssertionRequest) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.


### GetIpAddress

`func (o *GenerateSamlAssertionRequest) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *GenerateSamlAssertionRequest) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *GenerateSamlAssertionRequest) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *GenerateSamlAssertionRequest) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


