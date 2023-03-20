# UpdatePasswordInsecureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | Set to the password value using cleartext. Hashes are never stored as cleartext. They are stored securely using cryptographic hash. OneLogin continuously upgrades the strength of the hash. Ensure that the value meets the password strength requirements set for the account. | 
**PasswordConfirmation** | **string** | Ensure that this value matches the password value exactly. | 
**ValidatePolicy** | Pointer to **bool** | Will passwords validate against the User Policy. Defaults to false. | [optional] [default to false]

## Methods

### NewUpdatePasswordInsecureRequest

`func NewUpdatePasswordInsecureRequest(password string, passwordConfirmation string, ) *UpdatePasswordInsecureRequest`

NewUpdatePasswordInsecureRequest instantiates a new UpdatePasswordInsecureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePasswordInsecureRequestWithDefaults

`func NewUpdatePasswordInsecureRequestWithDefaults() *UpdatePasswordInsecureRequest`

NewUpdatePasswordInsecureRequestWithDefaults instantiates a new UpdatePasswordInsecureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *UpdatePasswordInsecureRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdatePasswordInsecureRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdatePasswordInsecureRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPasswordConfirmation

`func (o *UpdatePasswordInsecureRequest) GetPasswordConfirmation() string`

GetPasswordConfirmation returns the PasswordConfirmation field if non-nil, zero value otherwise.

### GetPasswordConfirmationOk

`func (o *UpdatePasswordInsecureRequest) GetPasswordConfirmationOk() (*string, bool)`

GetPasswordConfirmationOk returns a tuple with the PasswordConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordConfirmation

`func (o *UpdatePasswordInsecureRequest) SetPasswordConfirmation(v string)`

SetPasswordConfirmation sets PasswordConfirmation field to given value.


### GetValidatePolicy

`func (o *UpdatePasswordInsecureRequest) GetValidatePolicy() bool`

GetValidatePolicy returns the ValidatePolicy field if non-nil, zero value otherwise.

### GetValidatePolicyOk

`func (o *UpdatePasswordInsecureRequest) GetValidatePolicyOk() (*bool, bool)`

GetValidatePolicyOk returns a tuple with the ValidatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatePolicy

`func (o *UpdatePasswordInsecureRequest) SetValidatePolicy(v bool)`

SetValidatePolicy sets ValidatePolicy field to given value.

### HasValidatePolicy

`func (o *UpdatePasswordInsecureRequest) HasValidatePolicy() bool`

HasValidatePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


