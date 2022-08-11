# GetUserApps200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The App ID | [optional] 
**IconUrl** | Pointer to **string** | A url for the icon that represents the app in the OneLogin portal | [optional] 
**Extension** | Pointer to **bool** | Boolean that indicates if the OneLogin browser extension is required to launch this app. | [optional] 
**LoginId** | Pointer to **int32** | Unqiue identifier for this user and app combination. | [optional] 
**Name** | Pointer to **string** | The name of the app. | [optional] 
**ProvisioningStatus** | Pointer to **string** |  | [optional] 
**ProvisioningState** | Pointer to **string** | If provisioning is enabled this indicates the state of provisioning for the given user. | [optional] 
**ProvisioningEnabled** | Pointer to **bool** | Indicates if provisioning is enabled for this app. | [optional] 

## Methods

### NewGetUserApps200ResponseInner

`func NewGetUserApps200ResponseInner() *GetUserApps200ResponseInner`

NewGetUserApps200ResponseInner instantiates a new GetUserApps200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserApps200ResponseInnerWithDefaults

`func NewGetUserApps200ResponseInnerWithDefaults() *GetUserApps200ResponseInner`

NewGetUserApps200ResponseInnerWithDefaults instantiates a new GetUserApps200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetUserApps200ResponseInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetUserApps200ResponseInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetUserApps200ResponseInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetUserApps200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIconUrl

`func (o *GetUserApps200ResponseInner) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *GetUserApps200ResponseInner) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *GetUserApps200ResponseInner) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *GetUserApps200ResponseInner) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetExtension

`func (o *GetUserApps200ResponseInner) GetExtension() bool`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *GetUserApps200ResponseInner) GetExtensionOk() (*bool, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *GetUserApps200ResponseInner) SetExtension(v bool)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *GetUserApps200ResponseInner) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetLoginId

`func (o *GetUserApps200ResponseInner) GetLoginId() int32`

GetLoginId returns the LoginId field if non-nil, zero value otherwise.

### GetLoginIdOk

`func (o *GetUserApps200ResponseInner) GetLoginIdOk() (*int32, bool)`

GetLoginIdOk returns a tuple with the LoginId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginId

`func (o *GetUserApps200ResponseInner) SetLoginId(v int32)`

SetLoginId sets LoginId field to given value.

### HasLoginId

`func (o *GetUserApps200ResponseInner) HasLoginId() bool`

HasLoginId returns a boolean if a field has been set.

### GetName

`func (o *GetUserApps200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetUserApps200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetUserApps200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetUserApps200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvisioningStatus

`func (o *GetUserApps200ResponseInner) GetProvisioningStatus() string`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *GetUserApps200ResponseInner) GetProvisioningStatusOk() (*string, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *GetUserApps200ResponseInner) SetProvisioningStatus(v string)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *GetUserApps200ResponseInner) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### GetProvisioningState

`func (o *GetUserApps200ResponseInner) GetProvisioningState() string`

GetProvisioningState returns the ProvisioningState field if non-nil, zero value otherwise.

### GetProvisioningStateOk

`func (o *GetUserApps200ResponseInner) GetProvisioningStateOk() (*string, bool)`

GetProvisioningStateOk returns a tuple with the ProvisioningState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningState

`func (o *GetUserApps200ResponseInner) SetProvisioningState(v string)`

SetProvisioningState sets ProvisioningState field to given value.

### HasProvisioningState

`func (o *GetUserApps200ResponseInner) HasProvisioningState() bool`

HasProvisioningState returns a boolean if a field has been set.

### GetProvisioningEnabled

`func (o *GetUserApps200ResponseInner) GetProvisioningEnabled() bool`

GetProvisioningEnabled returns the ProvisioningEnabled field if non-nil, zero value otherwise.

### GetProvisioningEnabledOk

`func (o *GetUserApps200ResponseInner) GetProvisioningEnabledOk() (*bool, bool)`

GetProvisioningEnabledOk returns a tuple with the ProvisioningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningEnabled

`func (o *GetUserApps200ResponseInner) SetProvisioningEnabled(v bool)`

SetProvisioningEnabled sets ProvisioningEnabled field to given value.

### HasProvisioningEnabled

`func (o *GetUserApps200ResponseInner) HasProvisioningEnabled() bool`

HasProvisioningEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


