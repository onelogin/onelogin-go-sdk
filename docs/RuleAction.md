# RuleAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the Action | [optional] 
**Value** | Pointer to **string** | The unique identifier of the action. This should be used when defining actions for a User Mapping. | [optional] 

## Methods

### NewRuleAction

`func NewRuleAction() *RuleAction`

NewRuleAction instantiates a new RuleAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleActionWithDefaults

`func NewRuleActionWithDefaults() *RuleAction`

NewRuleActionWithDefaults instantiates a new RuleAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RuleAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *RuleAction) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RuleAction) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RuleAction) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RuleAction) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


