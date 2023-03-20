# GenerateOTPRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresIn** | Pointer to **int32** | Set the duration of the token in seconds. Token expiration defaults to 259200 seconds &#x3D; 72 hours. 72 hours is the max value. | [optional] 
**Reusable** | Pointer to **bool** | Defines if the token is reusable multiple times within the expiry window. Value defaults to false. If set to true, token can be used multiple times, until it expires. | [optional] [default to false]

## Methods

### NewGenerateOTPRequest

`func NewGenerateOTPRequest() *GenerateOTPRequest`

NewGenerateOTPRequest instantiates a new GenerateOTPRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateOTPRequestWithDefaults

`func NewGenerateOTPRequestWithDefaults() *GenerateOTPRequest`

NewGenerateOTPRequestWithDefaults instantiates a new GenerateOTPRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresIn

`func (o *GenerateOTPRequest) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *GenerateOTPRequest) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *GenerateOTPRequest) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *GenerateOTPRequest) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetReusable

`func (o *GenerateOTPRequest) GetReusable() bool`

GetReusable returns the Reusable field if non-nil, zero value otherwise.

### GetReusableOk

`func (o *GenerateOTPRequest) GetReusableOk() (*bool, bool)`

GetReusableOk returns a tuple with the Reusable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusable

`func (o *GenerateOTPRequest) SetReusable(v bool)`

SetReusable sets Reusable field to given value.

### HasReusable

`func (o *GenerateOTPRequest) HasReusable() bool`

HasReusable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


