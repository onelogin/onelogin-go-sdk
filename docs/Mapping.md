# Mapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** | The name of the mapping. | 
**Enabled** | **bool** | Indicates if the mapping is enabled or not. | 
**Match** | **string** | Indicates how conditions should be matched. | 
**Position** | **int32** | Indicates the order of the mapping. When &#x60;null&#x60; this will default to last position. | 
**Conditions** | [**[]Condition**](Condition.md) | An array of conditions that the user must meet in order for the mapping to be applied. | 
**Actions** | [**[]ActionObj**](ActionObj.md) | An array of actions that will be applied to the users that are matched by the conditions. | 

## Methods

### NewMapping

`func NewMapping(name string, enabled bool, match string, position int32, conditions []Condition, actions []ActionObj, ) *Mapping`

NewMapping instantiates a new Mapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMappingWithDefaults

`func NewMappingWithDefaults() *Mapping`

NewMappingWithDefaults instantiates a new Mapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Mapping) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Mapping) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Mapping) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Mapping) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Mapping) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Mapping) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Mapping) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *Mapping) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Mapping) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Mapping) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMatch

`func (o *Mapping) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *Mapping) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *Mapping) SetMatch(v string)`

SetMatch sets Match field to given value.


### GetPosition

`func (o *Mapping) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Mapping) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Mapping) SetPosition(v int32)`

SetPosition sets Position field to given value.


### GetConditions

`func (o *Mapping) GetConditions() []Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *Mapping) GetConditionsOk() (*[]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *Mapping) SetConditions(v []Condition)`

SetConditions sets Conditions field to given value.


### GetActions

`func (o *Mapping) GetActions() []ActionObj`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *Mapping) GetActionsOk() (*[]ActionObj, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *Mapping) SetActions(v []ActionObj)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


