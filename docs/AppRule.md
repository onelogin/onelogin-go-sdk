# AppRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | App Rule ID | [optional] 
**Name** | Pointer to **string** | Rule Name | [optional] 
**Match** | Pointer to **string** | Indicates how conditions should be matched. | [optional] 
**Enabled** | Pointer to **bool** | Indicates if the rule is enabled or not. | [optional] 
**Position** | Pointer to **int32** | Indicates the order of the rule. When &#x60;null&#x60; this will default to last position. | [optional] 
**Conditions** | Pointer to [**[]Condition**](Condition.md) | An array of conditions that the user must meet in order for the rule to be applied. | [optional] 
**Actions** | Pointer to [**[]ActionObj**](ActionObj.md) |  | [optional] 

## Methods

### NewAppRule

`func NewAppRule() *AppRule`

NewAppRule instantiates a new AppRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRuleWithDefaults

`func NewAppRuleWithDefaults() *AppRule`

NewAppRuleWithDefaults instantiates a new AppRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AppRule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppRule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppRule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AppRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AppRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AppRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMatch

`func (o *AppRule) GetMatch() string`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *AppRule) GetMatchOk() (*string, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *AppRule) SetMatch(v string)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *AppRule) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### GetEnabled

`func (o *AppRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AppRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AppRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AppRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPosition

`func (o *AppRule) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *AppRule) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *AppRule) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *AppRule) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetConditions

`func (o *AppRule) GetConditions() []Condition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *AppRule) GetConditionsOk() (*[]Condition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *AppRule) SetConditions(v []Condition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *AppRule) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetActions

`func (o *AppRule) GetActions() []ActionObj`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *AppRule) GetActionsOk() (*[]ActionObj, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *AppRule) SetActions(v []ActionObj)`

SetActions sets Actions field to given value.

### HasActions

`func (o *AppRule) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


