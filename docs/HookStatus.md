# HookStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | responses status nam | [optional] 
**Message** | Pointer to **string** | your operation was successful | [optional] 

## Methods

### NewHookStatus

`func NewHookStatus() *HookStatus`

NewHookStatus instantiates a new HookStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHookStatusWithDefaults

`func NewHookStatusWithDefaults() *HookStatus`

NewHookStatusWithDefaults instantiates a new HookStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *HookStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HookStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HookStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *HookStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMessage

`func (o *HookStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HookStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HookStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HookStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


