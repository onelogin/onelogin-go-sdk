# VerifyUserRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Otp** | Pointer to **int32** | One-Time-Password (OTP) that was sent to the user based on the chosen factor. OneLogin SMS and OneLogin Email support this OTP code. | [optional] 

## Methods

### NewVerifyUserRegistrationRequest

`func NewVerifyUserRegistrationRequest() *VerifyUserRegistrationRequest`

NewVerifyUserRegistrationRequest instantiates a new VerifyUserRegistrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyUserRegistrationRequestWithDefaults

`func NewVerifyUserRegistrationRequestWithDefaults() *VerifyUserRegistrationRequest`

NewVerifyUserRegistrationRequestWithDefaults instantiates a new VerifyUserRegistrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtp

`func (o *VerifyUserRegistrationRequest) GetOtp() int32`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *VerifyUserRegistrationRequest) GetOtpOk() (*int32, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *VerifyUserRegistrationRequest) SetOtp(v int32)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *VerifyUserRegistrationRequest) HasOtp() bool`

HasOtp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


