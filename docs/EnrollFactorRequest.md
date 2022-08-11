# EnrollFactorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FactorId** | **int32** | The identifier of the factor to enroll the user with. | 
**DisplayName** | **string** | A name for the users device. | 
**ExpiresIn** | Pointer to **string** | Defaults to 120. Valid values are: 120-900 seconds. | [optional] 
**Verified** | Pointer to **bool** | Defaults to false. | [optional] 
**RedirectTo** | Pointer to **string** | Redirects MagicLink success page to specified URL after 2 seconds. | [optional] 
**CustomMessage** | Pointer to **string** | A message template that will be sent via SMS. Max length of the message after template items are inserted is 160 characters including the OTP code. | [optional] 

## Methods

### NewEnrollFactorRequest

`func NewEnrollFactorRequest(factorId int32, displayName string, ) *EnrollFactorRequest`

NewEnrollFactorRequest instantiates a new EnrollFactorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollFactorRequestWithDefaults

`func NewEnrollFactorRequestWithDefaults() *EnrollFactorRequest`

NewEnrollFactorRequestWithDefaults instantiates a new EnrollFactorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactorId

`func (o *EnrollFactorRequest) GetFactorId() int32`

GetFactorId returns the FactorId field if non-nil, zero value otherwise.

### GetFactorIdOk

`func (o *EnrollFactorRequest) GetFactorIdOk() (*int32, bool)`

GetFactorIdOk returns a tuple with the FactorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorId

`func (o *EnrollFactorRequest) SetFactorId(v int32)`

SetFactorId sets FactorId field to given value.


### GetDisplayName

`func (o *EnrollFactorRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EnrollFactorRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EnrollFactorRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetExpiresIn

`func (o *EnrollFactorRequest) GetExpiresIn() string`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *EnrollFactorRequest) GetExpiresInOk() (*string, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *EnrollFactorRequest) SetExpiresIn(v string)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *EnrollFactorRequest) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetVerified

`func (o *EnrollFactorRequest) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *EnrollFactorRequest) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *EnrollFactorRequest) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *EnrollFactorRequest) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetRedirectTo

`func (o *EnrollFactorRequest) GetRedirectTo() string`

GetRedirectTo returns the RedirectTo field if non-nil, zero value otherwise.

### GetRedirectToOk

`func (o *EnrollFactorRequest) GetRedirectToOk() (*string, bool)`

GetRedirectToOk returns a tuple with the RedirectTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTo

`func (o *EnrollFactorRequest) SetRedirectTo(v string)`

SetRedirectTo sets RedirectTo field to given value.

### HasRedirectTo

`func (o *EnrollFactorRequest) HasRedirectTo() bool`

HasRedirectTo returns a boolean if a field has been set.

### GetCustomMessage

`func (o *EnrollFactorRequest) GetCustomMessage() string`

GetCustomMessage returns the CustomMessage field if non-nil, zero value otherwise.

### GetCustomMessageOk

`func (o *EnrollFactorRequest) GetCustomMessageOk() (*string, bool)`

GetCustomMessageOk returns a tuple with the CustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessage

`func (o *EnrollFactorRequest) SetCustomMessage(v string)`

SetCustomMessage sets CustomMessage field to given value.

### HasCustomMessage

`func (o *EnrollFactorRequest) HasCustomMessage() bool`

HasCustomMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


