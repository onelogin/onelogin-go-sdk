# VerifyMfaFactorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StateToken** | Pointer to **string** | The state_token is returned after a successful request to Enroll a Factor or Activate a Factor. The state_token MUST be provided if the needs_trigger attribute from the proceeding calls is set to true. Note that the state_token expires 120 seconds after creation. If the token is expired you will need to Activate the Factor again. | [optional] 
**OtpToken** | Pointer to **string** | OTP code provided by the device or SMS message sent to user. When a device like OneLogin Protect that supports Push has been used you do not need to provide the otp_token and can keep polling this endpoint until the state_token expires. | [optional] 

## Methods

### NewVerifyMfaFactorRequest

`func NewVerifyMfaFactorRequest() *VerifyMfaFactorRequest`

NewVerifyMfaFactorRequest instantiates a new VerifyMfaFactorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyMfaFactorRequestWithDefaults

`func NewVerifyMfaFactorRequestWithDefaults() *VerifyMfaFactorRequest`

NewVerifyMfaFactorRequestWithDefaults instantiates a new VerifyMfaFactorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStateToken

`func (o *VerifyMfaFactorRequest) GetStateToken() string`

GetStateToken returns the StateToken field if non-nil, zero value otherwise.

### GetStateTokenOk

`func (o *VerifyMfaFactorRequest) GetStateTokenOk() (*string, bool)`

GetStateTokenOk returns a tuple with the StateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateToken

`func (o *VerifyMfaFactorRequest) SetStateToken(v string)`

SetStateToken sets StateToken field to given value.

### HasStateToken

`func (o *VerifyMfaFactorRequest) HasStateToken() bool`

HasStateToken returns a boolean if a field has been set.

### GetOtpToken

`func (o *VerifyMfaFactorRequest) GetOtpToken() string`

GetOtpToken returns the OtpToken field if non-nil, zero value otherwise.

### GetOtpTokenOk

`func (o *VerifyMfaFactorRequest) GetOtpTokenOk() (*string, bool)`

GetOtpTokenOk returns a tuple with the OtpToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpToken

`func (o *VerifyMfaFactorRequest) SetOtpToken(v string)`

SetOtpToken sets OtpToken field to given value.

### HasOtpToken

`func (o *VerifyMfaFactorRequest) HasOtpToken() bool`

HasOtpToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


