# CreateDeviceVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **int32** | Specifies the factor to be verified. | 
**DisplayName** | Pointer to **string** | A name for the users device | [optional] 
**ExpiresIn** | Pointer to **string** | Defaults to 120. Valid values are: 120-900 seconds. | [optional] 
**RedirectTo** | Pointer to **string** | Only applies to Email MagicLink factor. Redirects MagicLink success page to specified URL after 2 seconds. Email must already be configured by the user. | [optional] 
**CustomMessage** | Pointer to **string** | Only applies to SMS factor. A message template that will be sent via SMS. Max length of the message after template items are inserted is 160 characters including the OTP code. SMS must already be configured by the user. The following template variables can be included in the message. - {{otp_code}} - The security code. - {{otp_expiry}} - The number of minutes until the one time code expires. | [optional] 

## Methods

### NewCreateDeviceVerificationRequest

`func NewCreateDeviceVerificationRequest(deviceId int32, ) *CreateDeviceVerificationRequest`

NewCreateDeviceVerificationRequest instantiates a new CreateDeviceVerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDeviceVerificationRequestWithDefaults

`func NewCreateDeviceVerificationRequestWithDefaults() *CreateDeviceVerificationRequest`

NewCreateDeviceVerificationRequestWithDefaults instantiates a new CreateDeviceVerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *CreateDeviceVerificationRequest) GetDeviceId() int32`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *CreateDeviceVerificationRequest) GetDeviceIdOk() (*int32, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *CreateDeviceVerificationRequest) SetDeviceId(v int32)`

SetDeviceId sets DeviceId field to given value.


### GetDisplayName

`func (o *CreateDeviceVerificationRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateDeviceVerificationRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateDeviceVerificationRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CreateDeviceVerificationRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetExpiresIn

`func (o *CreateDeviceVerificationRequest) GetExpiresIn() string`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *CreateDeviceVerificationRequest) GetExpiresInOk() (*string, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *CreateDeviceVerificationRequest) SetExpiresIn(v string)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *CreateDeviceVerificationRequest) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetRedirectTo

`func (o *CreateDeviceVerificationRequest) GetRedirectTo() string`

GetRedirectTo returns the RedirectTo field if non-nil, zero value otherwise.

### GetRedirectToOk

`func (o *CreateDeviceVerificationRequest) GetRedirectToOk() (*string, bool)`

GetRedirectToOk returns a tuple with the RedirectTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTo

`func (o *CreateDeviceVerificationRequest) SetRedirectTo(v string)`

SetRedirectTo sets RedirectTo field to given value.

### HasRedirectTo

`func (o *CreateDeviceVerificationRequest) HasRedirectTo() bool`

HasRedirectTo returns a boolean if a field has been set.

### GetCustomMessage

`func (o *CreateDeviceVerificationRequest) GetCustomMessage() string`

GetCustomMessage returns the CustomMessage field if non-nil, zero value otherwise.

### GetCustomMessageOk

`func (o *CreateDeviceVerificationRequest) GetCustomMessageOk() (*string, bool)`

GetCustomMessageOk returns a tuple with the CustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessage

`func (o *CreateDeviceVerificationRequest) SetCustomMessage(v string)`

SetCustomMessage sets CustomMessage field to given value.

### HasCustomMessage

`func (o *CreateDeviceVerificationRequest) HasCustomMessage() bool`

HasCustomMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


