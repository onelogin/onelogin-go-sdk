# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **int32** | an ID for the device type that must be submitted with the Verify Factor API call. | [optional] 
**DeviceType** | Pointer to **string** | Lists an available MFA device type, such as OneLogin OTP SMS or Google Authenticator. | [optional] 

## Methods

### NewDevice

`func NewDevice() *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *Device) GetDeviceId() int32`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *Device) GetDeviceIdOk() (*int32, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *Device) SetDeviceId(v int32)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *Device) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceType

`func (o *Device) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *Device) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *Device) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *Device) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


