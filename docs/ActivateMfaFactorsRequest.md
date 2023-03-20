# ActivateMfaFactorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StateTokenExpiresIn** | Pointer to **int32** | Optional. Sets the window of time in seconds that the factor must be verified within. Defaults to 120 seconds (2 minutes). Max 900 seconds (15 minutes). | [optional] 
**NumericSmsOtp** | Pointer to **bool** | Optional. Defaults to false. Only applies to SMS factor. When set to &#x60;true&#x60; a 6 digit numeric code will be sent to the user instead of the standard code which is alphanumeric. | [optional] 
**SmsMessage** | Pointer to **string** | Optional. Only applies to SMS factor. A message template that will be sent via SMS. Max length of the message after template items are inserted is 160 characters. The following template variables can be included in the message. - {{otp_code}} - The security code. - {{expiration}} - The number of minutes until the one time code expires. | [optional] 

## Methods

### NewActivateMfaFactorsRequest

`func NewActivateMfaFactorsRequest() *ActivateMfaFactorsRequest`

NewActivateMfaFactorsRequest instantiates a new ActivateMfaFactorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivateMfaFactorsRequestWithDefaults

`func NewActivateMfaFactorsRequestWithDefaults() *ActivateMfaFactorsRequest`

NewActivateMfaFactorsRequestWithDefaults instantiates a new ActivateMfaFactorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStateTokenExpiresIn

`func (o *ActivateMfaFactorsRequest) GetStateTokenExpiresIn() int32`

GetStateTokenExpiresIn returns the StateTokenExpiresIn field if non-nil, zero value otherwise.

### GetStateTokenExpiresInOk

`func (o *ActivateMfaFactorsRequest) GetStateTokenExpiresInOk() (*int32, bool)`

GetStateTokenExpiresInOk returns a tuple with the StateTokenExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTokenExpiresIn

`func (o *ActivateMfaFactorsRequest) SetStateTokenExpiresIn(v int32)`

SetStateTokenExpiresIn sets StateTokenExpiresIn field to given value.

### HasStateTokenExpiresIn

`func (o *ActivateMfaFactorsRequest) HasStateTokenExpiresIn() bool`

HasStateTokenExpiresIn returns a boolean if a field has been set.

### GetNumericSmsOtp

`func (o *ActivateMfaFactorsRequest) GetNumericSmsOtp() bool`

GetNumericSmsOtp returns the NumericSmsOtp field if non-nil, zero value otherwise.

### GetNumericSmsOtpOk

`func (o *ActivateMfaFactorsRequest) GetNumericSmsOtpOk() (*bool, bool)`

GetNumericSmsOtpOk returns a tuple with the NumericSmsOtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericSmsOtp

`func (o *ActivateMfaFactorsRequest) SetNumericSmsOtp(v bool)`

SetNumericSmsOtp sets NumericSmsOtp field to given value.

### HasNumericSmsOtp

`func (o *ActivateMfaFactorsRequest) HasNumericSmsOtp() bool`

HasNumericSmsOtp returns a boolean if a field has been set.

### GetSmsMessage

`func (o *ActivateMfaFactorsRequest) GetSmsMessage() string`

GetSmsMessage returns the SmsMessage field if non-nil, zero value otherwise.

### GetSmsMessageOk

`func (o *ActivateMfaFactorsRequest) GetSmsMessageOk() (*string, bool)`

GetSmsMessageOk returns a tuple with the SmsMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsMessage

`func (o *ActivateMfaFactorsRequest) SetSmsMessage(v string)`

SetSmsMessage sets SmsMessage field to given value.

### HasSmsMessage

`func (o *ActivateMfaFactorsRequest) HasSmsMessage() bool`

HasSmsMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


