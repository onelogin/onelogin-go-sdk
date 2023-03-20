# ActionObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The action to apply | [optional] 
**Value** | Pointer to **[]string** | Only applicable to provisioned and set_* actions. Items in the array will be a plain text string or valid value for the selected action. | [optional] 

## Methods

### NewActionObj

`func NewActionObj() *ActionObj`

NewActionObj instantiates a new ActionObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionObjWithDefaults

`func NewActionObjWithDefaults() *ActionObj`

NewActionObjWithDefaults instantiates a new ActionObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ActionObj) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ActionObj) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ActionObj) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ActionObj) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetValue

`func (o *ActionObj) GetValue() []string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ActionObj) GetValueOk() (*[]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ActionObj) SetValue(v []string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ActionObj) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


