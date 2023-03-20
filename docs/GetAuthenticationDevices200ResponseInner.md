# GetAuthenticationDevices200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** | MFA device identifier. | [optional] 
**UserDisplayName** | Pointer to **string** | Authentication factor display name assigned by users when they register the device. | [optional] 
**TypeDisplayName** | Pointer to **string** | Authentication factor display name as it appears to users upon initial registration, as defined by admins at Settings &gt; Authentication Factors. | [optional] 
**AuthFactorName** | Pointer to **string** | Authentication factor name, as it appears to administrators in OneLogin. | [optional] 
**Default** | Pointer to **bool** | true &#x3D; is user’s default MFA device for OneLogin. | [optional] [default to false]

## Methods

### NewGetAuthenticationDevices200ResponseInner

`func NewGetAuthenticationDevices200ResponseInner() *GetAuthenticationDevices200ResponseInner`

NewGetAuthenticationDevices200ResponseInner instantiates a new GetAuthenticationDevices200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAuthenticationDevices200ResponseInnerWithDefaults

`func NewGetAuthenticationDevices200ResponseInnerWithDefaults() *GetAuthenticationDevices200ResponseInner`

NewGetAuthenticationDevices200ResponseInnerWithDefaults instantiates a new GetAuthenticationDevices200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *GetAuthenticationDevices200ResponseInner) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *GetAuthenticationDevices200ResponseInner) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *GetAuthenticationDevices200ResponseInner) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *GetAuthenticationDevices200ResponseInner) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetUserDisplayName

`func (o *GetAuthenticationDevices200ResponseInner) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *GetAuthenticationDevices200ResponseInner) GetUserDisplayNameOk() (*string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayName

`func (o *GetAuthenticationDevices200ResponseInner) SetUserDisplayName(v string)`

SetUserDisplayName sets UserDisplayName field to given value.

### HasUserDisplayName

`func (o *GetAuthenticationDevices200ResponseInner) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### GetTypeDisplayName

`func (o *GetAuthenticationDevices200ResponseInner) GetTypeDisplayName() string`

GetTypeDisplayName returns the TypeDisplayName field if non-nil, zero value otherwise.

### GetTypeDisplayNameOk

`func (o *GetAuthenticationDevices200ResponseInner) GetTypeDisplayNameOk() (*string, bool)`

GetTypeDisplayNameOk returns a tuple with the TypeDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDisplayName

`func (o *GetAuthenticationDevices200ResponseInner) SetTypeDisplayName(v string)`

SetTypeDisplayName sets TypeDisplayName field to given value.

### HasTypeDisplayName

`func (o *GetAuthenticationDevices200ResponseInner) HasTypeDisplayName() bool`

HasTypeDisplayName returns a boolean if a field has been set.

### GetAuthFactorName

`func (o *GetAuthenticationDevices200ResponseInner) GetAuthFactorName() string`

GetAuthFactorName returns the AuthFactorName field if non-nil, zero value otherwise.

### GetAuthFactorNameOk

`func (o *GetAuthenticationDevices200ResponseInner) GetAuthFactorNameOk() (*string, bool)`

GetAuthFactorNameOk returns a tuple with the AuthFactorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthFactorName

`func (o *GetAuthenticationDevices200ResponseInner) SetAuthFactorName(v string)`

SetAuthFactorName sets AuthFactorName field to given value.

### HasAuthFactorName

`func (o *GetAuthenticationDevices200ResponseInner) HasAuthFactorName() bool`

HasAuthFactorName returns a boolean if a field has been set.

### GetDefault

`func (o *GetAuthenticationDevices200ResponseInner) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *GetAuthenticationDevices200ResponseInner) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *GetAuthenticationDevices200ResponseInner) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *GetAuthenticationDevices200ResponseInner) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


