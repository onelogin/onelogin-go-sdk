# CreateFactorRegistration201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** | MFA device identifier. | [optional] 
**UserDisplayName** | Pointer to **string** | Authentication factor display name assigned by users when they register the device. | [optional] 
**TypeDisplayName** | Pointer to **string** | Authentication factor display name as it appears to users upon initial registration, as defined by admins at Settings &gt; Authentication Factors. | [optional] 
**AuthFactorName** | Pointer to **string** | Authentication factor name, as it appears to administrators in OneLogin. | [optional] 
**Id** | Pointer to **string** | Verification identifier used in subsequent verification step. | [optional] 
**UserId** | Pointer to **string** | User identifier | [optional] 

## Methods

### NewCreateFactorRegistration201Response

`func NewCreateFactorRegistration201Response() *CreateFactorRegistration201Response`

NewCreateFactorRegistration201Response instantiates a new CreateFactorRegistration201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFactorRegistration201ResponseWithDefaults

`func NewCreateFactorRegistration201ResponseWithDefaults() *CreateFactorRegistration201Response`

NewCreateFactorRegistration201ResponseWithDefaults instantiates a new CreateFactorRegistration201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *CreateFactorRegistration201Response) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *CreateFactorRegistration201Response) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *CreateFactorRegistration201Response) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *CreateFactorRegistration201Response) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetUserDisplayName

`func (o *CreateFactorRegistration201Response) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *CreateFactorRegistration201Response) GetUserDisplayNameOk() (*string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayName

`func (o *CreateFactorRegistration201Response) SetUserDisplayName(v string)`

SetUserDisplayName sets UserDisplayName field to given value.

### HasUserDisplayName

`func (o *CreateFactorRegistration201Response) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### GetTypeDisplayName

`func (o *CreateFactorRegistration201Response) GetTypeDisplayName() string`

GetTypeDisplayName returns the TypeDisplayName field if non-nil, zero value otherwise.

### GetTypeDisplayNameOk

`func (o *CreateFactorRegistration201Response) GetTypeDisplayNameOk() (*string, bool)`

GetTypeDisplayNameOk returns a tuple with the TypeDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDisplayName

`func (o *CreateFactorRegistration201Response) SetTypeDisplayName(v string)`

SetTypeDisplayName sets TypeDisplayName field to given value.

### HasTypeDisplayName

`func (o *CreateFactorRegistration201Response) HasTypeDisplayName() bool`

HasTypeDisplayName returns a boolean if a field has been set.

### GetAuthFactorName

`func (o *CreateFactorRegistration201Response) GetAuthFactorName() string`

GetAuthFactorName returns the AuthFactorName field if non-nil, zero value otherwise.

### GetAuthFactorNameOk

`func (o *CreateFactorRegistration201Response) GetAuthFactorNameOk() (*string, bool)`

GetAuthFactorNameOk returns a tuple with the AuthFactorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthFactorName

`func (o *CreateFactorRegistration201Response) SetAuthFactorName(v string)`

SetAuthFactorName sets AuthFactorName field to given value.

### HasAuthFactorName

`func (o *CreateFactorRegistration201Response) HasAuthFactorName() bool`

HasAuthFactorName returns a boolean if a field has been set.

### GetId

`func (o *CreateFactorRegistration201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateFactorRegistration201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateFactorRegistration201Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateFactorRegistration201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *CreateFactorRegistration201Response) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateFactorRegistration201Response) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateFactorRegistration201Response) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *CreateFactorRegistration201Response) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


