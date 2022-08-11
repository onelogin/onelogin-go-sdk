# GenerateMfaTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresIn** | Pointer to **string** | Set the duration of the token in seconds. | [optional] 
**Reusable** | Pointer to **bool** | Defines if the token is reusable multiple times within the expiry window. | [optional] 

## Methods

### NewGenerateMfaTokenRequest

`func NewGenerateMfaTokenRequest() *GenerateMfaTokenRequest`

NewGenerateMfaTokenRequest instantiates a new GenerateMfaTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateMfaTokenRequestWithDefaults

`func NewGenerateMfaTokenRequestWithDefaults() *GenerateMfaTokenRequest`

NewGenerateMfaTokenRequestWithDefaults instantiates a new GenerateMfaTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresIn

`func (o *GenerateMfaTokenRequest) GetExpiresIn() string`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *GenerateMfaTokenRequest) GetExpiresInOk() (*string, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *GenerateMfaTokenRequest) SetExpiresIn(v string)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *GenerateMfaTokenRequest) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetReusable

`func (o *GenerateMfaTokenRequest) GetReusable() bool`

GetReusable returns the Reusable field if non-nil, zero value otherwise.

### GetReusableOk

`func (o *GenerateMfaTokenRequest) GetReusableOk() (*bool, bool)`

GetReusableOk returns a tuple with the Reusable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusable

`func (o *GenerateMfaTokenRequest) SetReusable(v bool)`

SetReusable sets Reusable field to given value.

### HasReusable

`func (o *GenerateMfaTokenRequest) HasReusable() bool`

HasReusable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


