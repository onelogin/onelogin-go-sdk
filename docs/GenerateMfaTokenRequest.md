# GenerateMFAtokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresIn** | Pointer to **int32** | Set the duration of the token in seconds. Token expiration defaults to 259200 seconds &#x3D; 72 hours. 72 hours is the max value. | [optional] 
**Reusable** | Pointer to **bool** | Defines if the token is reusable multiple times within the expiry window. Value defaults to false. If set to true, token can be used multiple times, until it expires. | [optional] [default to false]

## Methods

### NewGenerateMFAtokenRequest

`func NewGenerateMFAtokenRequest() *GenerateMFAtokenRequest`

NewGenerateMFAtokenRequest instantiates a new GenerateMFAtokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateMFAtokenRequestWithDefaults

`func NewGenerateMFAtokenRequestWithDefaults() *GenerateMFAtokenRequest`

NewGenerateMFAtokenRequestWithDefaults instantiates a new GenerateMFAtokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresIn

`func (o *GenerateMFAtokenRequest) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *GenerateMFAtokenRequest) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *GenerateMFAtokenRequest) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *GenerateMFAtokenRequest) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetReusable

`func (o *GenerateMFAtokenRequest) GetReusable() bool`

GetReusable returns the Reusable field if non-nil, zero value otherwise.

### GetReusableOk

`func (o *GenerateMFAtokenRequest) GetReusableOk() (*bool, bool)`

GetReusableOk returns a tuple with the Reusable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReusable

`func (o *GenerateMFAtokenRequest) SetReusable(v bool)`

SetReusable sets Reusable field to given value.

### HasReusable

`func (o *GenerateMFAtokenRequest) HasReusable() bool`

HasReusable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


