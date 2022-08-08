# RiskRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The name of this rule | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | The type parameter specifies the type of rule that will be created. | [optional] 
**Target** | Pointer to **string** | The target parameter that will be used when evaluating the rule against an incoming event. | [optional] 
**Filters** | Pointer to **[]string** | A list of IP addresses or country codes or names to evaluate against each event. | [optional] 
**Source** | Pointer to [**Source**](Source.md) |  | [optional] 

## Methods

### NewRiskRule

`func NewRiskRule() *RiskRule`

NewRiskRule instantiates a new RiskRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRuleWithDefaults

`func NewRiskRuleWithDefaults() *RiskRule`

NewRiskRuleWithDefaults instantiates a new RiskRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RiskRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RiskRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskRule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RiskRule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RiskRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RiskRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RiskRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RiskRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *RiskRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskRule) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RiskRule) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTarget

`func (o *RiskRule) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *RiskRule) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *RiskRule) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *RiskRule) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetFilters

`func (o *RiskRule) GetFilters() []string`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *RiskRule) GetFiltersOk() (*[]string, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *RiskRule) SetFilters(v []string)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *RiskRule) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetSource

`func (o *RiskRule) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *RiskRule) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *RiskRule) SetSource(v Source)`

SetSource sets Source field to given value.

### HasSource

`func (o *RiskRule) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


