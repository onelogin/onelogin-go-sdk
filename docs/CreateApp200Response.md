# CreateApp200Response

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
**Parameters** | [**AppParameters**](AppParameters.md) |  | 
**EnforcementPoint** | Pointer to [**EnforcementPoint**](EnforcementPoint.md) |  | [optional] 
**Configuration** | [**ConfigurationSaml**](ConfigurationSaml.md) |  | 
**Sso** | Pointer to [**SsoSaml**](SsoSaml.md) |  | [optional] 

## Methods

### NewCreateApp200Response

`func NewCreateApp200Response(name string, visible bool, description string, policyId int32, connectorId int32, parameters AppParameters, configuration ConfigurationSaml, ) *CreateApp200Response`

NewCreateApp200Response instantiates a new CreateApp200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApp200ResponseWithDefaults

`func NewCreateApp200ResponseWithDefaults() *CreateApp200Response`

NewCreateApp200ResponseWithDefaults instantiates a new CreateApp200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateApp200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateApp200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateApp200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CreateApp200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CreateApp200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateApp200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateApp200Response) SetName(v string)`

SetName sets Name field to given value.


### GetVisible

`func (o *CreateApp200Response) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *CreateApp200Response) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *CreateApp200Response) SetVisible(v bool)`

SetVisible sets Visible field to given value.


### GetDescription

`func (o *CreateApp200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateApp200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateApp200Response) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetNotes

`func (o *CreateApp200Response) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateApp200Response) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateApp200Response) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateApp200Response) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetIconUrl

`func (o *CreateApp200Response) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *CreateApp200Response) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *CreateApp200Response) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *CreateApp200Response) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetAuthMethod

`func (o *CreateApp200Response) GetAuthMethod() AuthMethod`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *CreateApp200Response) GetAuthMethodOk() (*AuthMethod, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *CreateApp200Response) SetAuthMethod(v AuthMethod)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *CreateApp200Response) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetPolicyId

`func (o *CreateApp200Response) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *CreateApp200Response) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *CreateApp200Response) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.


### GetAllowAssumedSignin

`func (o *CreateApp200Response) GetAllowAssumedSignin() bool`

GetAllowAssumedSignin returns the AllowAssumedSignin field if non-nil, zero value otherwise.

### GetAllowAssumedSigninOk

`func (o *CreateApp200Response) GetAllowAssumedSigninOk() (*bool, bool)`

GetAllowAssumedSigninOk returns a tuple with the AllowAssumedSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAssumedSignin

`func (o *CreateApp200Response) SetAllowAssumedSignin(v bool)`

SetAllowAssumedSignin sets AllowAssumedSignin field to given value.

### HasAllowAssumedSignin

`func (o *CreateApp200Response) HasAllowAssumedSignin() bool`

HasAllowAssumedSignin returns a boolean if a field has been set.

### GetTabId

`func (o *CreateApp200Response) GetTabId() int32`

GetTabId returns the TabId field if non-nil, zero value otherwise.

### GetTabIdOk

`func (o *CreateApp200Response) GetTabIdOk() (*int32, bool)`

GetTabIdOk returns a tuple with the TabId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabId

`func (o *CreateApp200Response) SetTabId(v int32)`

SetTabId sets TabId field to given value.

### HasTabId

`func (o *CreateApp200Response) HasTabId() bool`

HasTabId returns a boolean if a field has been set.

### GetConnectorId

`func (o *CreateApp200Response) GetConnectorId() int32`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *CreateApp200Response) GetConnectorIdOk() (*int32, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *CreateApp200Response) SetConnectorId(v int32)`

SetConnectorId sets ConnectorId field to given value.


### GetCreatedAt

`func (o *CreateApp200Response) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateApp200Response) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateApp200Response) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateApp200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CreateApp200Response) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateApp200Response) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateApp200Response) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CreateApp200Response) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRoleIds

`func (o *CreateApp200Response) GetRoleIds() []int32`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *CreateApp200Response) GetRoleIdsOk() (*[]int32, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *CreateApp200Response) SetRoleIds(v []int32)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *CreateApp200Response) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetProvisioning

`func (o *CreateApp200Response) GetProvisioning() GenericAppProvisioning`

GetProvisioning returns the Provisioning field if non-nil, zero value otherwise.

### GetProvisioningOk

`func (o *CreateApp200Response) GetProvisioningOk() (*GenericAppProvisioning, bool)`

GetProvisioningOk returns a tuple with the Provisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioning

`func (o *CreateApp200Response) SetProvisioning(v GenericAppProvisioning)`

SetProvisioning sets Provisioning field to given value.

### HasProvisioning

`func (o *CreateApp200Response) HasProvisioning() bool`

HasProvisioning returns a boolean if a field has been set.

### GetParameters

`func (o *CreateApp200Response) GetParameters() AppParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CreateApp200Response) GetParametersOk() (*AppParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CreateApp200Response) SetParameters(v AppParameters)`

SetParameters sets Parameters field to given value.


### GetEnforcementPoint

`func (o *CreateApp200Response) GetEnforcementPoint() EnforcementPoint`

GetEnforcementPoint returns the EnforcementPoint field if non-nil, zero value otherwise.

### GetEnforcementPointOk

`func (o *CreateApp200Response) GetEnforcementPointOk() (*EnforcementPoint, bool)`

GetEnforcementPointOk returns a tuple with the EnforcementPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementPoint

`func (o *CreateApp200Response) SetEnforcementPoint(v EnforcementPoint)`

SetEnforcementPoint sets EnforcementPoint field to given value.

### HasEnforcementPoint

`func (o *CreateApp200Response) HasEnforcementPoint() bool`

HasEnforcementPoint returns a boolean if a field has been set.

### GetConfiguration

`func (o *CreateApp200Response) GetConfiguration() ConfigurationSaml`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CreateApp200Response) GetConfigurationOk() (*ConfigurationSaml, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CreateApp200Response) SetConfiguration(v ConfigurationSaml)`

SetConfiguration sets Configuration field to given value.


### GetSso

`func (o *CreateApp200Response) GetSso() SsoSaml`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *CreateApp200Response) GetSsoOk() (*SsoSaml, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *CreateApp200Response) SetSso(v SsoSaml)`

SetSso sets Sso field to given value.

### HasSso

`func (o *CreateApp200Response) HasSso() bool`

HasSso returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


