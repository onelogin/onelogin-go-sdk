# OtpDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FactorId** | **int32** | The identifier of the factor to enroll the user with. | [readonly] 
**DisplayName** | **string** | A name for the users device | 
**Number** | **string** | The phone number of the user in E.164 format. | 
**Verified** | Pointer to **bool** | Defaults to false. Some factors like SMS or Voice require that a user recieve a token and then that token is supplied to the Verify endpoint before the device is considered active. You can set verfied to &#x60;true&#x60; which indicates the the users phone number is pre verified and the device can be immediately activated.            | [optional] 

## Methods

### NewOtpDevice

`func NewOtpDevice(factorId int32, displayName string, number string, ) *OtpDevice`

NewOtpDevice instantiates a new OtpDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtpDeviceWithDefaults

`func NewOtpDeviceWithDefaults() *OtpDevice`

NewOtpDeviceWithDefaults instantiates a new OtpDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactorId

`func (o *OtpDevice) GetFactorId() int32`

GetFactorId returns the FactorId field if non-nil, zero value otherwise.

### GetFactorIdOk

`func (o *OtpDevice) GetFactorIdOk() (*int32, bool)`

GetFactorIdOk returns a tuple with the FactorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactorId

`func (o *OtpDevice) SetFactorId(v int32)`

SetFactorId sets FactorId field to given value.


### GetDisplayName

`func (o *OtpDevice) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OtpDevice) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OtpDevice) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetNumber

`func (o *OtpDevice) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *OtpDevice) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *OtpDevice) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetVerified

`func (o *OtpDevice) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *OtpDevice) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *OtpDevice) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *OtpDevice) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


