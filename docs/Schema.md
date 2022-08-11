# Schema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Apps unique ID in OneLogin. | [optional] 
**ConnectorId** | Pointer to **int32** | ID of the apps underlying connector. | [optional] 
**Name** | Pointer to **string** | App name. | [optional] 
**Description** | Pointer to **string** | Freeform description of the app. | [optional] 
**Notes** | Pointer to **string** | Freeform notes about the app. | [optional] 
**PolicyId** | Pointer to **int32** | The security policy assigned to the app. | [optional] 
**BrandId** | Pointer to **int32** | The custom login page branding to use for this app. Applies to app initiated logins via OIDC or SAML. | [optional] 
**IconUrl** | Pointer to **string** | A link to the apps icon url. | [optional] 
**Visible** | Pointer to **bool** | Indicates if the app is visible in the OneLogin portal. | [optional] 
**AuthMethod** | Pointer to **int32** | An ID indicating the type of app. | [optional] 
**TabId** | Pointer to **int32** | ID of the OneLogin portal tab that the app is assigned to. | [optional] 
**CreatedAt** | Pointer to **string** | The date the app was created. | [optional] 
**UpdatedAt** | Pointer to **string** | The date the app was last updated. | [optional] 
**RoleIds** | Pointer to **[]int32** | List of Role IDs that are assigned to the app. On App Create or Update the entire array is replaced with the values provided. | [optional] 
**AllowAssumedSignin** | Pointer to **bool** | Indicates whether or not administrators can access the app as a user that they have assumed control over. | [optional] 
**Provisioning** | Pointer to [**SchemaProvisioning**](SchemaProvisioning.md) |  | [optional] 
**Sso** | Pointer to **map[string]interface{}** |  | [optional] 
**Configuration** | Pointer to **map[string]interface{}** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**EnforcementPoint** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSchema

`func NewSchema() *Schema`

NewSchema instantiates a new Schema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaWithDefaults

`func NewSchemaWithDefaults() *Schema`

NewSchemaWithDefaults instantiates a new Schema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Schema) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Schema) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Schema) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Schema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConnectorId

`func (o *Schema) GetConnectorId() int32`

GetConnectorId returns the ConnectorId field if non-nil, zero value otherwise.

### GetConnectorIdOk

`func (o *Schema) GetConnectorIdOk() (*int32, bool)`

GetConnectorIdOk returns a tuple with the ConnectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorId

`func (o *Schema) SetConnectorId(v int32)`

SetConnectorId sets ConnectorId field to given value.

### HasConnectorId

`func (o *Schema) HasConnectorId() bool`

HasConnectorId returns a boolean if a field has been set.

### GetName

`func (o *Schema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Schema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Schema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Schema) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Schema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Schema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Schema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Schema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotes

`func (o *Schema) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *Schema) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *Schema) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *Schema) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetPolicyId

`func (o *Schema) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *Schema) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *Schema) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *Schema) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetBrandId

`func (o *Schema) GetBrandId() int32`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *Schema) GetBrandIdOk() (*int32, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *Schema) SetBrandId(v int32)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *Schema) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### GetIconUrl

`func (o *Schema) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *Schema) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *Schema) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *Schema) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetVisible

`func (o *Schema) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *Schema) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *Schema) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *Schema) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetAuthMethod

`func (o *Schema) GetAuthMethod() int32`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *Schema) GetAuthMethodOk() (*int32, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *Schema) SetAuthMethod(v int32)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *Schema) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetTabId

`func (o *Schema) GetTabId() int32`

GetTabId returns the TabId field if non-nil, zero value otherwise.

### GetTabIdOk

`func (o *Schema) GetTabIdOk() (*int32, bool)`

GetTabIdOk returns a tuple with the TabId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTabId

`func (o *Schema) SetTabId(v int32)`

SetTabId sets TabId field to given value.

### HasTabId

`func (o *Schema) HasTabId() bool`

HasTabId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Schema) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Schema) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Schema) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Schema) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Schema) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Schema) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Schema) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Schema) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetRoleIds

`func (o *Schema) GetRoleIds() []int32`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *Schema) GetRoleIdsOk() (*[]int32, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *Schema) SetRoleIds(v []int32)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *Schema) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetAllowAssumedSignin

`func (o *Schema) GetAllowAssumedSignin() bool`

GetAllowAssumedSignin returns the AllowAssumedSignin field if non-nil, zero value otherwise.

### GetAllowAssumedSigninOk

`func (o *Schema) GetAllowAssumedSigninOk() (*bool, bool)`

GetAllowAssumedSigninOk returns a tuple with the AllowAssumedSignin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAssumedSignin

`func (o *Schema) SetAllowAssumedSignin(v bool)`

SetAllowAssumedSignin sets AllowAssumedSignin field to given value.

### HasAllowAssumedSignin

`func (o *Schema) HasAllowAssumedSignin() bool`

HasAllowAssumedSignin returns a boolean if a field has been set.

### GetProvisioning

`func (o *Schema) GetProvisioning() SchemaProvisioning`

GetProvisioning returns the Provisioning field if non-nil, zero value otherwise.

### GetProvisioningOk

`func (o *Schema) GetProvisioningOk() (*SchemaProvisioning, bool)`

GetProvisioningOk returns a tuple with the Provisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioning

`func (o *Schema) SetProvisioning(v SchemaProvisioning)`

SetProvisioning sets Provisioning field to given value.

### HasProvisioning

`func (o *Schema) HasProvisioning() bool`

HasProvisioning returns a boolean if a field has been set.

### GetSso

`func (o *Schema) GetSso() map[string]interface{}`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *Schema) GetSsoOk() (*map[string]interface{}, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *Schema) SetSso(v map[string]interface{})`

SetSso sets Sso field to given value.

### HasSso

`func (o *Schema) HasSso() bool`

HasSso returns a boolean if a field has been set.

### GetConfiguration

`func (o *Schema) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *Schema) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *Schema) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *Schema) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetParameters

`func (o *Schema) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *Schema) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *Schema) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *Schema) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetEnforcementPoint

`func (o *Schema) GetEnforcementPoint() map[string]interface{}`

GetEnforcementPoint returns the EnforcementPoint field if non-nil, zero value otherwise.

### GetEnforcementPointOk

`func (o *Schema) GetEnforcementPointOk() (*map[string]interface{}, bool)`

GetEnforcementPointOk returns a tuple with the EnforcementPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforcementPoint

`func (o *Schema) SetEnforcementPoint(v map[string]interface{})`

SetEnforcementPoint sets EnforcementPoint field to given value.

### HasEnforcementPoint

`func (o *Schema) HasEnforcementPoint() bool`

HasEnforcementPoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


