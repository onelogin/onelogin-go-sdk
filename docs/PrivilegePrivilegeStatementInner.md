# PrivilegePrivilegeStatementInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | **string** | Set to “Allow.” By default, all actions are denied, this Statement allows the listed actions to be executed. | 
**Action** | **[]string** | An array of strings that represent actions within OneLogin. Actions are prefixed with the class of object they are related to and followed by a specific action for the given class. e.g. users:List, where the class is users and the specific action is List. Don’t mix classes within an Action array. To create a privilege that includes multiple different classes, create multiple statements. A wildcard * that includes all actions is supported. Use wildcards to create a Super User privilege. | 
**Scope** | **[]string** | Target the privileged action against specific resources with the scope. The scope pattern is the class of object used by the Action, followed by an ID that represents a resource in OneLogin. e.g. apps/1234, where apps is the class and 1234 is the ID of an app. The wildcard * is supported and indicates that all resources of the class type declared, in the Action, are in scope. The Action and Scope classes must match. However, there is an exception, a scope of roles/{role_id} can be combined with Actions on the user or app class. The exception allows you to target groups of users or apps with specific actions. | 

## Methods

### NewPrivilegePrivilegeStatementInner

`func NewPrivilegePrivilegeStatementInner(effect string, action []string, scope []string, ) *PrivilegePrivilegeStatementInner`

NewPrivilegePrivilegeStatementInner instantiates a new PrivilegePrivilegeStatementInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegePrivilegeStatementInnerWithDefaults

`func NewPrivilegePrivilegeStatementInnerWithDefaults() *PrivilegePrivilegeStatementInner`

NewPrivilegePrivilegeStatementInnerWithDefaults instantiates a new PrivilegePrivilegeStatementInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *PrivilegePrivilegeStatementInner) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *PrivilegePrivilegeStatementInner) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *PrivilegePrivilegeStatementInner) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetAction

`func (o *PrivilegePrivilegeStatementInner) GetAction() []string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PrivilegePrivilegeStatementInner) GetActionOk() (*[]string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PrivilegePrivilegeStatementInner) SetAction(v []string)`

SetAction sets Action field to given value.


### GetScope

`func (o *PrivilegePrivilegeStatementInner) GetScope() []string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *PrivilegePrivilegeStatementInner) GetScopeOk() (*[]string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *PrivilegePrivilegeStatementInner) SetScope(v []string)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


