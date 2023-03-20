# CreateFactorRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FactorId** | **int32** | The identifier of the factor to enroll the user with. See Get Available Factors for a list of possible id values. | 
**DisplayName** | **string** | A name for the users device | 
**ExpiresIn** | Pointer to **string** | Defaults to 120. Valid values are: 120-900 seconds. | [optional] 
**Verified** | Pointer to **bool** | Defaults to false. The following factors support verified &#x3D; true as part of the initial registration request: OneLogin SMS, OneLogin Voice, OneLogin Email. When using verified &#x3D; true it is critical that the supported factors have pre-verified values, most likely imported from an existing directory or by the users themselvdes. Factors such as Authenticator and OneLogin Protect do not support verification &#x3D; true as the user interaction is required to verify the factor. | [optional] 
**RedirectTo** | Pointer to **string** | Only applies to Email MagicLink factor. Redirects MagicLink success page to specified URL after 2 seconds. Email must already be configured by the user. | [optional] 
**CustomMessage** | Pointer to **string** | Only applies to SMS factor. A message template that will be sent via SMS. Max length of the message after template items are inserted is 160 characters including the OTP code. SMS must already be configured by the user. The following template variables can be included in the message. - {{otp_code}} - The security code. - {{otp_expiry}} - The number of minutes until the one time code expires. | [optional] 

## Methods

### NewCreateFactorRegistrationRequest

`func NewCreateFactorRegistrationRequest(factorId int32, displayName string, ) *CreateFactorRegistrationRequest`

NewCreateFactorRegistrationRequest instantiates a new CreateFactorRegistrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFactorRegistrationRequestWithDefaults

`func NewCreateFactorRegistrationRequestWithDefaults() *CreateFactorRegistrationRequest`

NewCreateFactorRegistrationRequestWithDefaults instantiates a new CreateFactorRegistrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactorId

`func (o *CreateFactorRegistrationRequest) GetFactorId() int32`

GetFactorId returns the FactorId field if non-nil, zero value otherwise.

### GetFactorIdOk

`func (o *CreateFactorRegistrationRequest) GetFactorIdOk() (*int32, bool)`

GetFactorIdOk returns a tuple with the FactorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorId

`func (o *CreateFactorRegistrationRequest) SetFactorId(v int32)`

SetFactorId sets FactorId field to given value.


### GetDisplayName

`func (o *CreateFactorRegistrationRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateFactorRegistrationRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateFactorRegistrationRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetExpiresIn

`func (o *CreateFactorRegistrationRequest) GetExpiresIn() string`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *CreateFactorRegistrationRequest) GetExpiresInOk() (*string, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *CreateFactorRegistrationRequest) SetExpiresIn(v string)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *CreateFactorRegistrationRequest) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetVerified

`func (o *CreateFactorRegistrationRequest) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *CreateFactorRegistrationRequest) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *CreateFactorRegistrationRequest) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *CreateFactorRegistrationRequest) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetRedirectTo

`func (o *CreateFactorRegistrationRequest) GetRedirectTo() string`

GetRedirectTo returns the RedirectTo field if non-nil, zero value otherwise.

### GetRedirectToOk

`func (o *CreateFactorRegistrationRequest) GetRedirectToOk() (*string, bool)`

GetRedirectToOk returns a tuple with the RedirectTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTo

`func (o *CreateFactorRegistrationRequest) SetRedirectTo(v string)`

SetRedirectTo sets RedirectTo field to given value.

### HasRedirectTo

`func (o *CreateFactorRegistrationRequest) HasRedirectTo() bool`

HasRedirectTo returns a boolean if a field has been set.

### GetCustomMessage

`func (o *CreateFactorRegistrationRequest) GetCustomMessage() string`

GetCustomMessage returns the CustomMessage field if non-nil, zero value otherwise.

### GetCustomMessageOk

`func (o *CreateFactorRegistrationRequest) GetCustomMessageOk() (*string, bool)`

GetCustomMessageOk returns a tuple with the CustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessage

`func (o *CreateFactorRegistrationRequest) SetCustomMessage(v string)`

SetCustomMessage sets CustomMessage field to given value.

### HasCustomMessage

`func (o *CreateFactorRegistrationRequest) HasCustomMessage() bool`

HasCustomMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


