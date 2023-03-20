# SamlFactor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **string** | App ID of the app for which you want to generate a SAML token. This is the app ID in OneLogin. | 
**DeviceId** | **string** | Provide the MFA device_id you are submitting for verification. The device_id is supplied by the Generate SAML Assertion API. | 
**StateToken** | **string** | Provide the state_token associated with the MFA device_id you are submitting for verification. The state_token is supplied by the Generate SAML Assertion API. | 
**OtpToken** | Pointer to **string** | Provide the OTP value for the MFA factor you are submitting for verification. For some MFA factors; such as OneLogin OTP SMS, which requires the user to request an OTP; the otp_token value is not required, and if not included, returns a 200 OK - Pending result. You’ll make a subsequent Verify Factor API call to provide the otp_token value once it has been provided to the user. | [optional] 
**DoNotNotify** | Pointer to **bool** | When verifying MFA via Protect Push, set this to true to stop additional push notifications being sent to the OneLogin Protect device. | [optional] 

## Methods

### NewSamlFactor

`func NewSamlFactor(appId string, deviceId string, stateToken string, ) *SamlFactor`

NewSamlFactor instantiates a new SamlFactor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlFactorWithDefaults

`func NewSamlFactorWithDefaults() *SamlFactor`

NewSamlFactorWithDefaults instantiates a new SamlFactor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *SamlFactor) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *SamlFactor) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *SamlFactor) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetDeviceId

`func (o *SamlFactor) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *SamlFactor) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *SamlFactor) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetStateToken

`func (o *SamlFactor) GetStateToken() string`

GetStateToken returns the StateToken field if non-nil, zero value otherwise.

### GetStateTokenOk

`func (o *SamlFactor) GetStateTokenOk() (*string, bool)`

GetStateTokenOk returns a tuple with the StateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateToken

`func (o *SamlFactor) SetStateToken(v string)`

SetStateToken sets StateToken field to given value.


### GetOtpToken

`func (o *SamlFactor) GetOtpToken() string`

GetOtpToken returns the OtpToken field if non-nil, zero value otherwise.

### GetOtpTokenOk

`func (o *SamlFactor) GetOtpTokenOk() (*string, bool)`

GetOtpTokenOk returns a tuple with the OtpToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpToken

`func (o *SamlFactor) SetOtpToken(v string)`

SetOtpToken sets OtpToken field to given value.

### HasOtpToken

`func (o *SamlFactor) HasOtpToken() bool`

HasOtpToken returns a boolean if a field has been set.

### GetDoNotNotify

`func (o *SamlFactor) GetDoNotNotify() bool`

GetDoNotNotify returns the DoNotNotify field if non-nil, zero value otherwise.

### GetDoNotNotifyOk

`func (o *SamlFactor) GetDoNotNotifyOk() (*bool, bool)`

GetDoNotNotifyOk returns a tuple with the DoNotNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotNotify

`func (o *SamlFactor) SetDoNotNotify(v bool)`

SetDoNotNotify sets DoNotNotify field to given value.

### HasDoNotNotify

`func (o *SamlFactor) HasDoNotNotify() bool`

HasDoNotNotify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


