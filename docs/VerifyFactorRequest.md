# VerifyFactorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Otp** | Pointer to **string** | OTP code provided by the device or SMS message sent to user. | [optional] 
**DeviceId** | Pointer to **int32** | ID of the specified device which has been registerd for the given user. | [optional] 

## Methods

### NewVerifyFactorRequest

`func NewVerifyFactorRequest() *VerifyFactorRequest`

NewVerifyFactorRequest instantiates a new VerifyFactorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyFactorRequestWithDefaults

`func NewVerifyFactorRequestWithDefaults() *VerifyFactorRequest`

NewVerifyFactorRequestWithDefaults instantiates a new VerifyFactorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtp

`func (o *VerifyFactorRequest) GetOtp() string`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *VerifyFactorRequest) GetOtpOk() (*string, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *VerifyFactorRequest) SetOtp(v string)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *VerifyFactorRequest) HasOtp() bool`

HasOtp returns a boolean if a field has been set.

### GetDeviceId

`func (o *VerifyFactorRequest) GetDeviceId() int32`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *VerifyFactorRequest) GetDeviceIdOk() (*int32, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *VerifyFactorRequest) SetDeviceId(v int32)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *VerifyFactorRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


