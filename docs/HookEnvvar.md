# HookEnvvar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for the Hook Environment Variable | [optional] [readonly] 
**Name** | **string** | The name of the environment variable. | 
**CreatedAt** | Pointer to **string** | The ISO8601 formatted date that the environment variable was created. | [optional] 
**UpdatedAt** | Pointer to **string** | The ISO8601 formatted date that the environment variable was last updated. | [optional] 
**Value** | **string** | The secret value that will be encrypted at rest and injected in applicable hook functions at run time. | 

## Methods

### NewHookEnvvar

`func NewHookEnvvar(name string, value string, ) *HookEnvvar`

NewHookEnvvar instantiates a new HookEnvvar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHookEnvvarWithDefaults

`func NewHookEnvvarWithDefaults() *HookEnvvar`

NewHookEnvvarWithDefaults instantiates a new HookEnvvar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *HookEnvvar) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HookEnvvar) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HookEnvvar) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HookEnvvar) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *HookEnvvar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HookEnvvar) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HookEnvvar) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *HookEnvvar) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HookEnvvar) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HookEnvvar) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HookEnvvar) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *HookEnvvar) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HookEnvvar) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HookEnvvar) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HookEnvvar) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValue

`func (o *HookEnvvar) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *HookEnvvar) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *HookEnvvar) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


