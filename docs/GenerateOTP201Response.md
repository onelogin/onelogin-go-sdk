# GenerateOTP201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MfaToken** | Pointer to **string** | Token can function as a temporary MFA token. It can be used to authenticate for any app when valid. | [optional] 
**Reusable** | Pointer to **bool** | true indcates the token can be used multiple times, until it expires. false indicates the token is invalid after a single use or once it expires. Defaults to false. | [optional] [default to false]
**ExpiresAt** | Pointer to **string** | Defines the expiration time and date for the token. Format is UTC time. | [optional] 
**DeviceId** | Pointer to **string** | A unique identifier for the temp otp device that has been created for this token. | [optional] 

## Methods

### NewGenerateOTP201Response

`func NewGenerateOTP201Response() *GenerateOTP201Response`

NewGenerateOTP201Response instantiates a new GenerateOTP201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateOTP201ResponseWithDefaults

`func NewGenerateOTP201ResponseWithDefaults() *GenerateOTP201Response`

NewGenerateOTP201ResponseWithDefaults instantiates a new GenerateOTP201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMfaToken

`func (o *GenerateOTP201Response) GetMfaToken() string`

GetMfaToken returns the MfaToken field if non-nil, zero value otherwise.

### GetMfaTokenOk

`func (o *GenerateOTP201Response) GetMfaTokenOk() (*string, bool)`

GetMfaTokenOk returns a tuple with the MfaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaToken

`func (o *GenerateOTP201Response) SetMfaToken(v string)`

SetMfaToken sets MfaToken field to given value.

### HasMfaToken

`func (o *GenerateOTP201Response) HasMfaToken() bool`

HasMfaToken returns a boolean if a field has been set.

### GetReusable

`func (o *GenerateOTP201Response) GetReusable() bool`

GetReusable returns the Reusable field if non-nil, zero value otherwise.

### GetReusableOk

`func (o *GenerateOTP201Response) GetReusableOk() (*bool, bool)`

GetReusableOk returns a tuple with the Reusable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusable

`func (o *GenerateOTP201Response) SetReusable(v bool)`

SetReusable sets Reusable field to given value.

### HasReusable

`func (o *GenerateOTP201Response) HasReusable() bool`

HasReusable returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GenerateOTP201Response) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GenerateOTP201Response) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GenerateOTP201Response) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GenerateOTP201Response) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetDeviceId

`func (o *GenerateOTP201Response) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *GenerateOTP201Response) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *GenerateOTP201Response) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *GenerateOTP201Response) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


