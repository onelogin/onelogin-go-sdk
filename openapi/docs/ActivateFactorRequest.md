# ActivateFactorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **int32** | Required. Specifies the factor to be verified. | [optional] 
**ExpiresIn** | Pointer to **int32** | Optional. Sets the window of time in seconds that the factor must be verified within.  | [optional] 
**RedirectTo** | Pointer to **string** | Optional. Only applies to Email MagicLink factor. | [optional] 
**CustomMessage** | Pointer to **string** | Optional. Only applies to SMS factor. A message template that will be sent via SMS. | [optional] 

## Methods

### NewActivateFactorRequest

`func NewActivateFactorRequest() *ActivateFactorRequest`

NewActivateFactorRequest instantiates a new ActivateFactorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivateFactorRequestWithDefaults

`func NewActivateFactorRequestWithDefaults() *ActivateFactorRequest`

NewActivateFactorRequestWithDefaults instantiates a new ActivateFactorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *ActivateFactorRequest) GetDeviceId() int32`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ActivateFactorRequest) GetDeviceIdOk() (*int32, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ActivateFactorRequest) SetDeviceId(v int32)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *ActivateFactorRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetExpiresIn

`func (o *ActivateFactorRequest) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *ActivateFactorRequest) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *ActivateFactorRequest) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *ActivateFactorRequest) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetRedirectTo

`func (o *ActivateFactorRequest) GetRedirectTo() string`

GetRedirectTo returns the RedirectTo field if non-nil, zero value otherwise.

### GetRedirectToOk

`func (o *ActivateFactorRequest) GetRedirectToOk() (*string, bool)`

GetRedirectToOk returns a tuple with the RedirectTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectTo

`func (o *ActivateFactorRequest) SetRedirectTo(v string)`

SetRedirectTo sets RedirectTo field to given value.

### HasRedirectTo

`func (o *ActivateFactorRequest) HasRedirectTo() bool`

HasRedirectTo returns a boolean if a field has been set.

### GetCustomMessage

`func (o *ActivateFactorRequest) GetCustomMessage() string`

GetCustomMessage returns the CustomMessage field if non-nil, zero value otherwise.

### GetCustomMessageOk

`func (o *ActivateFactorRequest) GetCustomMessageOk() (*string, bool)`

GetCustomMessageOk returns a tuple with the CustomMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMessage

`func (o *ActivateFactorRequest) SetCustomMessage(v string)`

SetCustomMessage sets CustomMessage field to given value.

### HasCustomMessage

`func (o *ActivateFactorRequest) HasCustomMessage() bool`

HasCustomMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


