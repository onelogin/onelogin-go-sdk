# UpdatePasswordSecureRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | Set to the password value using a SHA-256-encoded value. If you are including your own password_salt value in your request, prepend the salt value to the cleartext password value before SHA-256-encoding it. For example, if your salt value is hello and your cleartext password value is password, the value you need to SHA-256-encode is hellopassword. The resulting encoded value would be b1c788abac15390de987ad17b65ac73c9b475d428a51f245c645a442fddd078b. Note that the alpha characters in this has must all be lower case. | 
**PasswordConfirmation** | **string** | This value must match the password value. | 
**PasswordAlgorithm** | **string** | Set to salt+sha256. | 
**PasswordSalt** | Pointer to **string** | Optional. If your password hash has been salted then you can provide the salt used in this param. This assumes that the salt was prepended to the password before doing the SHA256 hash. The API supports a salt value that is up to 40 characters long. | [optional] 

## Methods

### NewUpdatePasswordSecureRequest

`func NewUpdatePasswordSecureRequest(password string, passwordConfirmation string, passwordAlgorithm string, ) *UpdatePasswordSecureRequest`

NewUpdatePasswordSecureRequest instantiates a new UpdatePasswordSecureRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePasswordSecureRequestWithDefaults

`func NewUpdatePasswordSecureRequestWithDefaults() *UpdatePasswordSecureRequest`

NewUpdatePasswordSecureRequestWithDefaults instantiates a new UpdatePasswordSecureRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *UpdatePasswordSecureRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdatePasswordSecureRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdatePasswordSecureRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPasswordConfirmation

`func (o *UpdatePasswordSecureRequest) GetPasswordConfirmation() string`

GetPasswordConfirmation returns the PasswordConfirmation field if non-nil, zero value otherwise.

### GetPasswordConfirmationOk

`func (o *UpdatePasswordSecureRequest) GetPasswordConfirmationOk() (*string, bool)`

GetPasswordConfirmationOk returns a tuple with the PasswordConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordConfirmation

`func (o *UpdatePasswordSecureRequest) SetPasswordConfirmation(v string)`

SetPasswordConfirmation sets PasswordConfirmation field to given value.


### GetPasswordAlgorithm

`func (o *UpdatePasswordSecureRequest) GetPasswordAlgorithm() string`

GetPasswordAlgorithm returns the PasswordAlgorithm field if non-nil, zero value otherwise.

### GetPasswordAlgorithmOk

`func (o *UpdatePasswordSecureRequest) GetPasswordAlgorithmOk() (*string, bool)`

GetPasswordAlgorithmOk returns a tuple with the PasswordAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordAlgorithm

`func (o *UpdatePasswordSecureRequest) SetPasswordAlgorithm(v string)`

SetPasswordAlgorithm sets PasswordAlgorithm field to given value.


### GetPasswordSalt

`func (o *UpdatePasswordSecureRequest) GetPasswordSalt() string`

GetPasswordSalt returns the PasswordSalt field if non-nil, zero value otherwise.

### GetPasswordSaltOk

`func (o *UpdatePasswordSecureRequest) GetPasswordSaltOk() (*string, bool)`

GetPasswordSaltOk returns a tuple with the PasswordSalt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSalt

`func (o *UpdatePasswordSecureRequest) SetPasswordSalt(v string)`

SetPasswordSalt sets PasswordSalt field to given value.

### HasPasswordSalt

`func (o *UpdatePasswordSecureRequest) HasPasswordSalt() bool`

HasPasswordSalt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


