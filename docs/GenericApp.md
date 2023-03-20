# GenericApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Apps unique ID in OneLogin. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the app. | [optional] 
**Visible** | Pointer to **bool** | Indicates if the app is visible in the OneLogin portal. | [optional] 
**Description** | Pointer to **string** | Freeform description of the app. | [optional] 
**Notes** | Pointer to **string** | Freeform notes about the app. | [optional] 
**IconUrl** | Pointer to **string** | A link to the apps icon url | [optional] 
**AuthMethod** | Pointer to [**AuthMethod**](AuthMethod.md) |  | [optional] 
**PolicyId** | Pointer to **int32** | The security policy assigned to the app. | [optional] 
**AllowAssumedSignin** | Pointer to **bool** | Indicates whether or not administrators can access the app as a user that they have assumed control over. | [optional] 
**TabId** | Pointer to **int32** | ID of the OneLogin portal tab that the app is assigned to. | [optional] 
**ConnectorId** | Pointer to **int32** | ID of the connector to base the app from. | [optional] 
**CreatedAt** | Pointer to **string** | the date the app was created | [optional] 
**UpdatedAt** | Pointer to **string** | the date the app was last updated | [optional] 
**RoleIds** | Pointer to **[]int32** | List of Role IDs that are assigned to the app. On App Create or Update the entire array is replaced with the values provided. | [optional] 
**Provisioning** | Pointer to [**GenericAppProvisioning**](GenericAppProvisioning.md) |  | [optional] 
**Parameters** | Pointer to [**AppParameters**](AppParameters.md) |  | [optional] 
**EnforcementPoint** | Pointer to [**EnforcementPoint**](EnforcementPoint.md) |  | [optional] 

## Methods

### NewGenericApp

`func NewGenericApp() *GenericApp`

NewGenericApp instantiates a new GenericApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericAppWithDefaults

`func NewGenericAppWithDefaults() *GenericApp`

NewGenericAppWithDefaults instantiates a new GenericApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GenericApp) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GenericApp) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GenericApp) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GenericApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GenericApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenericApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenericApp) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GenericApp) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVisible

`func (o *GenericApp) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *GenericApp) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *GenericApp) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *GenericApp) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetDescription

`func (o *GenericApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GenericApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GenericApp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GenericApp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotes

`func (o *GenericApp) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GenericApp) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GenericApp) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *GenericApp) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetIconUrl

`func (o *GenericApp) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *GenericApp) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *GenericApp) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *GenericApp) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetAuthMethod

`func (o *GenericApp) GetAuthMethod() AuthMethod`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *GenericApp) GetAuthMethodOk() (*AuthMethod, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *GenericApp) SetAuthMethod(v AuthMethod)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *GenericApp) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetPolicyId

`func (o *GenericApp) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *GenericApp) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *GenericApp) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *GenericApp) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetAllowAssumedSignin

`func (o *GenericApp) GetAllowAssumedSignin() bool`

GetAllowAssumedSignin returns the AllowAssumedSignin field if non-nil, zero value otherwise.

### GetAllowAssumedSigninOk

`func (o *GenericApp) GetAllowAssumedSigninOk() (*bool, bool)`

GetAllowAssumedSigninOk returns a tuple with the AllowAssumedSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAssumedSignin

`func (o *GenericApp) SetAllowAssumedSignin(v bool)`

SetAllowAssumedSignin sets AllowAssumedSignin field to given value.

### HasAllowAssumedSignin

`func (o *GenericApp) HasAllowAssumedSignin() bool`

HasAllowAssumedSignin returns a boolean if a field has been set.

### GetTabId

`func (o *GenericApp) GetTabId() int32`

GetTabId returns the TabId field if non-nil, zero value otherwise.

### GetTabIdOk

`func (o *GenericApp) GetTabIdOk() (*int32, bool)`

GetTabIdOk returns a tuple with the TabId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabId

`func (o *GenericApp) SetTabId(v int32)`

SetTabId sets TabId field to given value.

### HasTabId

`func (o *GenericApp) HasTabId() bool`

HasTabId returns a boolean if a field has been set.

### GetConnectorId

`func (o *GenericApp) GetConnectorId() int32`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *GenericApp) GetConnectorIdOk() (*int32, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *GenericApp) SetConnectorId(v int32)`

SetConnectorId sets ConnectorId field to given value.

### HasConnectorId

`func (o *GenericApp) HasConnectorId() bool`

HasConnectorId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GenericApp) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GenericApp) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GenericApp) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GenericApp) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GenericApp) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GenericApp) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GenericApp) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GenericApp) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRoleIds

`func (o *GenericApp) GetRoleIds() []int32`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *GenericApp) GetRoleIdsOk() (*[]int32, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *GenericApp) SetRoleIds(v []int32)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *GenericApp) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetProvisioning

`func (o *GenericApp) GetProvisioning() GenericAppProvisioning`

GetProvisioning returns the Provisioning field if non-nil, zero value otherwise.

### GetProvisioningOk

`func (o *GenericApp) GetProvisioningOk() (*GenericAppProvisioning, bool)`

GetProvisioningOk returns a tuple with the Provisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioning

`func (o *GenericApp) SetProvisioning(v GenericAppProvisioning)`

SetProvisioning sets Provisioning field to given value.

### HasProvisioning

`func (o *GenericApp) HasProvisioning() bool`

HasProvisioning returns a boolean if a field has been set.

### GetParameters

`func (o *GenericApp) GetParameters() AppParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *GenericApp) GetParametersOk() (*AppParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *GenericApp) SetParameters(v AppParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *GenericApp) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetEnforcementPoint

`func (o *GenericApp) GetEnforcementPoint() EnforcementPoint`

GetEnforcementPoint returns the EnforcementPoint field if non-nil, zero value otherwise.

### GetEnforcementPointOk

`func (o *GenericApp) GetEnforcementPointOk() (*EnforcementPoint, bool)`

GetEnforcementPointOk returns a tuple with the EnforcementPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementPoint

`func (o *GenericApp) SetEnforcementPoint(v EnforcementPoint)`

SetEnforcementPoint sets EnforcementPoint field to given value.

### HasEnforcementPoint

`func (o *GenericApp) HasEnforcementPoint() bool`

HasEnforcementPoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


