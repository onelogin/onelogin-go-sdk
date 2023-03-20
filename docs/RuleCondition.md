# RuleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the operator | [optional] 
**Value** | Pointer to **string** | The condition operator value to use when creating or updating rules. | [optional] 

## Methods

### NewRuleCondition

`func NewRuleCondition() *RuleCondition`

NewRuleCondition instantiates a new RuleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleConditionWithDefaults

`func NewRuleConditionWithDefaults() *RuleCondition`

NewRuleConditionWithDefaults instantiates a new RuleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RuleCondition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleCondition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleCondition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleCondition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *RuleCondition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RuleCondition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RuleCondition) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RuleCondition) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


