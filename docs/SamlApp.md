# SamlApp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Apps unique ID in OneLogin. | [optional] [readonly] 
**Name** | **string** | The name of the app. | 
**Visible** | **bool** | Indicates if the app is visible in the OneLogin portal. | 
**Description** | **string** | Freeform description of the app. | 
**Notes** | Pointer to **string** | Freeform notes about the app. | [optional] 
**IconUrl** | Pointer to **string** | A link to the apps icon url | [optional] 
**AuthMethod** | Pointer to [**AuthMethod**](AuthMethod.md) |  | [optional] 
**PolicyId** | **int32** | The security policy assigned to the app. | 
**AllowAssumedSignin** | Pointer to **bool** | Indicates whether or not administrators can access the app as a user that they have assumed control over. | [optional] 
**TabId** | Pointer to **int32** | ID of the OneLogin portal tab that the app is assigned to. | [optional] 
**ConnectorId** | **int32** | ID of the connector to base the app from. | 
**CreatedAt** | Pointer to **string** | the date the app was created | [optional] 
**UpdatedAt** | Pointer to **string** | the date the app was last updated | [optional] 
**RoleIds** | Pointer to **[]int32** | List of Role IDs that are assigned to the app. On App Create or Update the entire array is replaced with the values provided. | [optional] 
**Provisioning** | Pointer to [**GenericAppProvisioning**](GenericAppProvisioning.md) |  | [optional] 
**Parameters** | [**SamlAppAllOfParameters**](SamlAppAllOfParameters.md) |  | 
**EnforcementPoint** | Pointer to [**EnforcementPoint**](EnforcementPoint.md) |  | [optional] 
**Configuration** | [**ConfigurationSaml**](ConfigurationSaml.md) |  | 
**Sso** | Pointer to [**SsoSaml**](SsoSaml.md) |  | [optional] 

## Methods

### NewSamlApp

`func NewSamlApp(name string, visible bool, description string, policyId int32, connectorId int32, parameters SamlAppAllOfParameters, configuration ConfigurationSaml, ) *SamlApp`

NewSamlApp instantiates a new SamlApp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlAppWithDefaults

`func NewSamlAppWithDefaults() *SamlApp`

NewSamlAppWithDefaults instantiates a new SamlApp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SamlApp) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SamlApp) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SamlApp) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SamlApp) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SamlApp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SamlApp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SamlApp) SetName(v string)`

SetName sets Name field to given value.


### GetVisible

`func (o *SamlApp) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *SamlApp) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *SamlApp) SetVisible(v bool)`

SetVisible sets Visible field to given value.


### GetDescription

`func (o *SamlApp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SamlApp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SamlApp) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetNotes

`func (o *SamlApp) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *SamlApp) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *SamlApp) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *SamlApp) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetIconUrl

`func (o *SamlApp) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *SamlApp) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *SamlApp) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *SamlApp) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetAuthMethod

`func (o *SamlApp) GetAuthMethod() AuthMethod`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *SamlApp) GetAuthMethodOk() (*AuthMethod, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *SamlApp) SetAuthMethod(v AuthMethod)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *SamlApp) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetPolicyId

`func (o *SamlApp) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *SamlApp) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *SamlApp) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.


### GetAllowAssumedSignin

`func (o *SamlApp) GetAllowAssumedSignin() bool`

GetAllowAssumedSignin returns the AllowAssumedSignin field if non-nil, zero value otherwise.

### GetAllowAssumedSigninOk

`func (o *SamlApp) GetAllowAssumedSigninOk() (*bool, bool)`

GetAllowAssumedSigninOk returns a tuple with the AllowAssumedSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAssumedSignin

`func (o *SamlApp) SetAllowAssumedSignin(v bool)`

SetAllowAssumedSignin sets AllowAssumedSignin field to given value.

### HasAllowAssumedSignin

`func (o *SamlApp) HasAllowAssumedSignin() bool`

HasAllowAssumedSignin returns a boolean if a field has been set.

### GetTabId

`func (o *SamlApp) GetTabId() int32`

GetTabId returns the TabId field if non-nil, zero value otherwise.

### GetTabIdOk

`func (o *SamlApp) GetTabIdOk() (*int32, bool)`

GetTabIdOk returns a tuple with the TabId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabId

`func (o *SamlApp) SetTabId(v int32)`

SetTabId sets TabId field to given value.

### HasTabId

`func (o *SamlApp) HasTabId() bool`

HasTabId returns a boolean if a field has been set.

### GetConnectorId

`func (o *SamlApp) GetConnectorId() int32`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *SamlApp) GetConnectorIdOk() (*int32, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *SamlApp) SetConnectorId(v int32)`

SetConnectorId sets ConnectorId field to given value.


### GetCreatedAt

`func (o *SamlApp) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SamlApp) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SamlApp) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SamlApp) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SamlApp) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SamlApp) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SamlApp) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SamlApp) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRoleIds

`func (o *SamlApp) GetRoleIds() []int32`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *SamlApp) GetRoleIdsOk() (*[]int32, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *SamlApp) SetRoleIds(v []int32)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *SamlApp) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetProvisioning

`func (o *SamlApp) GetProvisioning() GenericAppProvisioning`

GetProvisioning returns the Provisioning field if non-nil, zero value otherwise.

### GetProvisioningOk

`func (o *SamlApp) GetProvisioningOk() (*GenericAppProvisioning, bool)`

GetProvisioningOk returns a tuple with the Provisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioning

`func (o *SamlApp) SetProvisioning(v GenericAppProvisioning)`

SetProvisioning sets Provisioning field to given value.

### HasProvisioning

`func (o *SamlApp) HasProvisioning() bool`

HasProvisioning returns a boolean if a field has been set.

### GetParameters

`func (o *SamlApp) GetParameters() SamlAppAllOfParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *SamlApp) GetParametersOk() (*SamlAppAllOfParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *SamlApp) SetParameters(v SamlAppAllOfParameters)`

SetParameters sets Parameters field to given value.


### GetEnforcementPoint

`func (o *SamlApp) GetEnforcementPoint() EnforcementPoint`

GetEnforcementPoint returns the EnforcementPoint field if non-nil, zero value otherwise.

### GetEnforcementPointOk

`func (o *SamlApp) GetEnforcementPointOk() (*EnforcementPoint, bool)`

GetEnforcementPointOk returns a tuple with the EnforcementPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementPoint

`func (o *SamlApp) SetEnforcementPoint(v EnforcementPoint)`

SetEnforcementPoint sets EnforcementPoint field to given value.

### HasEnforcementPoint

`func (o *SamlApp) HasEnforcementPoint() bool`

HasEnforcementPoint returns a boolean if a field has been set.

### GetConfiguration

`func (o *SamlApp) GetConfiguration() ConfigurationSaml`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *SamlApp) GetConfigurationOk() (*ConfigurationSaml, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *SamlApp) SetConfiguration(v ConfigurationSaml)`

SetConfiguration sets Configuration field to given value.


### GetSso

`func (o *SamlApp) GetSso() SsoSaml`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *SamlApp) GetSsoOk() (*SsoSaml, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *SamlApp) SetSso(v SsoSaml)`

SetSso sets Sso field to given value.

### HasSso

`func (o *SamlApp) HasSso() bool`

HasSso returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


