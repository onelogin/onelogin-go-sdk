# Action

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The action to apply | [optional] 
**Value** | Pointer to **[]string** | Only applicable to provisioned and set_* actions. Items in the array will be a plain text string or valid value for the selected action. | [optional] 
**Expression** | Pointer to **string** | A regular expression to extract a value. Applies to provisionable, multi-selects, and string actions. | [optional] 
**Scriplet** | Pointer to **string** | A hash containing scriptlet code that returns a value. | [optional] 
**Macro** | Pointer to **string** | A template to construct a value. Applies to default, string, and list actions. | [optional] 

## Methods

### NewAction

`func NewAction() *Action`

NewAction instantiates a new Action object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionWithDefaults

`func NewActionWithDefaults() *Action`

NewActionWithDefaults instantiates a new Action object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *Action) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *Action) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *Action) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *Action) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetValue

`func (o *Action) GetValue() []string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Action) GetValueOk() (*[]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Action) SetValue(v []string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Action) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetExpression

`func (o *Action) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *Action) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *Action) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *Action) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetScriplet

`func (o *Action) GetScriplet() string`

GetScriplet returns the Scriplet field if non-nil, zero value otherwise.

### GetScripletOk

`func (o *Action) GetScripletOk() (*string, bool)`

GetScripletOk returns a tuple with the Scriplet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScriplet

`func (o *Action) SetScriplet(v string)`

SetScriplet sets Scriplet field to given value.

### HasScriplet

`func (o *Action) HasScriplet() bool`

HasScriplet returns a boolean if a field has been set.

### GetMacro

`func (o *Action) GetMacro() string`

GetMacro returns the Macro field if non-nil, zero value otherwise.

### GetMacroOk

`func (o *Action) GetMacroOk() (*string, bool)`

GetMacroOk returns a tuple with the Macro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacro

`func (o *Action) SetMacro(v string)`

SetMacro sets Macro field to given value.

### HasMacro

`func (o *Action) HasMacro() bool`

HasMacro returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


