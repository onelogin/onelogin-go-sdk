# VerifyUserRegistration200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Registration identifier. | [optional] 
**Status** | Pointer to **string** | pending registration has not been completed successfully. accepted registration has successfully completed. | [optional] 
**DeviceId** | Pointer to **string** | Device id to be used with Verify the Factor. | [optional] 

## Methods

### NewVerifyUserRegistration200Response

`func NewVerifyUserRegistration200Response() *VerifyUserRegistration200Response`

NewVerifyUserRegistration200Response instantiates a new VerifyUserRegistration200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerifyUserRegistration200ResponseWithDefaults

`func NewVerifyUserRegistration200ResponseWithDefaults() *VerifyUserRegistration200Response`

NewVerifyUserRegistration200ResponseWithDefaults instantiates a new VerifyUserRegistration200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VerifyUserRegistration200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VerifyUserRegistration200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VerifyUserRegistration200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VerifyUserRegistration200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *VerifyUserRegistration200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VerifyUserRegistration200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VerifyUserRegistration200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VerifyUserRegistration200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDeviceId

`func (o *VerifyUserRegistration200Response) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *VerifyUserRegistration200Response) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *VerifyUserRegistration200Response) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *VerifyUserRegistration200Response) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


