/*
OneLogin API

OpenAPI Specification for OneLogin

API version: 3.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onelogin

import (
	"encoding/json"
)

// checks if the SamlApp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamlApp{}

// SamlApp struct for SamlApp
type SamlApp struct {
	// Apps unique ID in OneLogin.
	Id *int32 `json:"id,omitempty"`
	// The name of the app.
	Name string `json:"name"`
	// Indicates if the app is visible in the OneLogin portal.
	Visible bool `json:"visible"`
	// Freeform description of the app.
	Description string `json:"description"`
	// Freeform notes about the app.
	Notes *string `json:"notes,omitempty"`
	// A link to the apps icon url
	IconUrl *string `json:"icon_url,omitempty"`
	AuthMethod *AuthMethod `json:"auth_method,omitempty"`
	// The security policy assigned to the app.
	PolicyId int32 `json:"policy_id"`
	// Indicates whether or not administrators can access the app as a user that they have assumed control over.
	AllowAssumedSignin *bool `json:"allow_assumed_signin,omitempty"`
	// ID of the OneLogin portal tab that the app is assigned to.
	TabId *int32 `json:"tab_id,omitempty"`
	// ID of the connector to base the app from.
	ConnectorId int32 `json:"connector_id"`
	// the date the app was created
	CreatedAt *string `json:"created_at,omitempty"`
	// the date the app was last updated
	UpdatedAt *string `json:"updated_at,omitempty"`
	// List of Role IDs that are assigned to the app. On App Create or Update the entire array is replaced with the values provided.
	RoleIds []int32 `json:"role_ids,omitempty"`
	Provisioning *GenericAppProvisioning `json:"provisioning,omitempty"`
	Parameters SamlAppAllOfParameters `json:"parameters"`
	EnforcementPoint *EnforcementPoint `json:"enforcement_point,omitempty"`
	Configuration ConfigurationSaml `json:"configuration"`
	Sso *SsoSaml `json:"sso,omitempty"`
}

// NewSamlApp instantiates a new SamlApp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlApp(name string, visible bool, description string, policyId int32, connectorId int32, parameters SamlAppAllOfParameters, configuration ConfigurationSaml) *SamlApp {
	this := SamlApp{}
	this.Name = name
	this.Visible = visible
	this.Description = description
	this.PolicyId = policyId
	this.ConnectorId = connectorId
	this.Parameters = parameters
	this.Configuration = configuration
	return &this
}

// NewSamlAppWithDefaults instantiates a new SamlApp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlAppWithDefaults() *SamlApp {
	this := SamlApp{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SamlApp) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApp) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SamlApp) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SamlApp) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *SamlApp) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SamlApp) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SamlApp) SetName(v string) {
	o.Name = v
}

// GetVisible returns the Visible field value
func (o *SamlApp) GetVisible() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value
// and a boolean to check if the value has been set.
func (o *SamlApp) GetVisibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visible, true
}

// SetVisible sets field value
func (o *SamlApp) SetVisible(v bool) {
	o.Visible = v
}

// GetDescription returns the Description field value
func (o *SamlApp) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SamlApp) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SamlApp) SetDescription(v string) {
	o.Description = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *SamlApp) GetNotes() string {
	if o == nil || isNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApp) GetNotesOk() (*string, bool) {
	if o == nil || isNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *SamlApp) HasNotes() bool {
	if o != nil && !isNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *SamlApp) SetNotes(v string) {
	o.Notes = &v
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise.
func (o *SamlApp) GetIconUrl() string {
	if o == nil || isNil(o.IconUrl) {
		var ret string
		return ret
	}
	return *o.IconUrl
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApp) GetIconUrlOk() (*string, bool) {
	if o == nil || isNil(o.IconUrl) {
		return nil, false
	}
	return o.IconUrl, true
}

// HasIconUrl returns a boolean if a field has been set.
func (o *SamlApp) HasIconUrl() bool {
	if o != nil && !isNil(o.IconUrl) {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given string and assigns it to the IconUrl field.
func (o *SamlApp) SetIconUrl(v string) {
	o.IconUrl = &v
}

// GetAuthMethod returns the AuthMethod field value if set, zero value otherwise.
func (o *SamlApp) GetAuthMethod() AuthMethod {
	if o == nil || isNil(o.AuthMethod) {
		var ret AuthMethod
		return ret
	}
	return *o.AuthMethod
}

// GetAuthMethodOk returns a tuple with the AuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApp) GetAuthMethodOk() (*AuthMethod, bool) {
	if o == nil || isNil(o.AuthMethod) {
		return nil, false
	}
	return o.AuthMethod, true
}

// HasAuthMethod returns a boolean if a field has been set.
func (o *SamlApp) HasAuthMethod() bool {
	if o != nil && !isNil(o.AuthMethod) {
		return true
	}

	return false
}

// SetAuthMethod gets a reference to the given AuthMethod and assigns it to the AuthMethod field.
func (o *SamlApp) SetAuthMethod(v AuthMethod) {
	o.AuthMethod = &v
}

// GetPolicyId returns the PolicyId field value
func (o *SamlApp) GetPolicyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value
// and a boolean to check if the value has been set.
func (o *SamlApp) GetPolicyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PolicyId, true
}

// SetPolicyId sets field value
func (o *SamlApp) SetPolicyId(v int32) {
	o.PolicyId = v
}

// GetAllowAssumedSignin returns the AllowAssumedSignin field value if set, zero value otherwise.
func (o *SamlApp) GetAllowAssumedSignin() bool {
	if o == nil || isNil(o.AllowAssumedSignin) {
		var ret bool
		return ret
	}
	return *o.AllowAssumedSignin
}

// GetAllowAssumedSigninOk returns a tuple with the AllowAssumedSignin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApp) GetAllowAssumedSigninOk() (*bool, bool) {
	if o == nil || isNil(o.AllowAssumedSignin) {
		return nil, false
	}
	return o.AllowAssumedSignin, true
}

// HasAllowAssumedSignin returns a boolean if a field has been set.
func (o *SamlApp) HasAllowAssumedSignin() bool {
	if o != nil && !isNil(o.AllowAssumedSignin) {
		return true
	}

	return false
}

// SetAllowAssumedSignin gets a reference to the given bool and assigns it to the AllowAssumedSignin field.
func (o *SamlApp) SetAllowAssumedSignin(v bool) {
	o.AllowAssumedSignin = &v
}

// GetTabId returns the TabId field value if set, zero value otherwise.
func (o *SamlApp) GetTabId() int32 {
	if o == nil || isNil(o.TabId) {
		var ret int32
		return ret
	}
	return *o.TabId
}

// GetTabIdOk returns a tuple with the TabId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApp) GetTabIdOk() (*int32, bool) {
	if o == nil || isNil(o.TabId) {
		return nil, false
	}
	return o.TabId, true
}

// HasTabId returns a boolean if a field has been set.
func (o *SamlApp) HasTabId() bool {
	if o != nil && !isNil(o.TabId) {
		return true
	}

	return false
}

// SetTabId gets a reference to the given int32 and assigns it to the TabId field.
func (o *SamlApp) SetTabId(v int32) {
	o.TabId = &v
}

// GetConnectorId returns the ConnectorId field value
func (o *SamlApp) GetConnectorId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConnectorId
}

// GetConnectorIdOk returns a tuple with the ConnectorId field value
// and a boolean to check if the value has been set.
func (o *SamlApp) GetConnectorIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorId, true
}

// SetConnectorId sets field value
func (o *SamlApp) SetConnectorId(v int32) {
	o.ConnectorId = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SamlApp) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApp) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SamlApp) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *SamlApp) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SamlApp) GetUpdatedAt() string {
	if o == nil || isNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApp) GetUpdatedAtOk() (*string, bool) {
	if o == nil || isNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SamlApp) HasUpdatedAt() bool {
	if o != nil && !isNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *SamlApp) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetRoleIds returns the RoleIds field value if set, zero value otherwise.
func (o *SamlApp) GetRoleIds() []int32 {
	if o == nil || isNil(o.RoleIds) {
		var ret []int32
		return ret
	}
	return o.RoleIds
}

// GetRoleIdsOk returns a tuple with the RoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApp) GetRoleIdsOk() ([]int32, bool) {
	if o == nil || isNil(o.RoleIds) {
		return nil, false
	}
	return o.RoleIds, true
}

// HasRoleIds returns a boolean if a field has been set.
func (o *SamlApp) HasRoleIds() bool {
	if o != nil && !isNil(o.RoleIds) {
		return true
	}

	return false
}

// SetRoleIds gets a reference to the given []int32 and assigns it to the RoleIds field.
func (o *SamlApp) SetRoleIds(v []int32) {
	o.RoleIds = v
}

// GetProvisioning returns the Provisioning field value if set, zero value otherwise.
func (o *SamlApp) GetProvisioning() GenericAppProvisioning {
	if o == nil || isNil(o.Provisioning) {
		var ret GenericAppProvisioning
		return ret
	}
	return *o.Provisioning
}

// GetProvisioningOk returns a tuple with the Provisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApp) GetProvisioningOk() (*GenericAppProvisioning, bool) {
	if o == nil || isNil(o.Provisioning) {
		return nil, false
	}
	return o.Provisioning, true
}

// HasProvisioning returns a boolean if a field has been set.
func (o *SamlApp) HasProvisioning() bool {
	if o != nil && !isNil(o.Provisioning) {
		return true
	}

	return false
}

// SetProvisioning gets a reference to the given GenericAppProvisioning and assigns it to the Provisioning field.
func (o *SamlApp) SetProvisioning(v GenericAppProvisioning) {
	o.Provisioning = &v
}

// GetParameters returns the Parameters field value
func (o *SamlApp) GetParameters() SamlAppAllOfParameters {
	if o == nil {
		var ret SamlAppAllOfParameters
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *SamlApp) GetParametersOk() (*SamlAppAllOfParameters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value
func (o *SamlApp) SetParameters(v SamlAppAllOfParameters) {
	o.Parameters = v
}

// GetEnforcementPoint returns the EnforcementPoint field value if set, zero value otherwise.
func (o *SamlApp) GetEnforcementPoint() EnforcementPoint {
	if o == nil || isNil(o.EnforcementPoint) {
		var ret EnforcementPoint
		return ret
	}
	return *o.EnforcementPoint
}

// GetEnforcementPointOk returns a tuple with the EnforcementPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApp) GetEnforcementPointOk() (*EnforcementPoint, bool) {
	if o == nil || isNil(o.EnforcementPoint) {
		return nil, false
	}
	return o.EnforcementPoint, true
}

// HasEnforcementPoint returns a boolean if a field has been set.
func (o *SamlApp) HasEnforcementPoint() bool {
	if o != nil && !isNil(o.EnforcementPoint) {
		return true
	}

	return false
}

// SetEnforcementPoint gets a reference to the given EnforcementPoint and assigns it to the EnforcementPoint field.
func (o *SamlApp) SetEnforcementPoint(v EnforcementPoint) {
	o.EnforcementPoint = &v
}

// GetConfiguration returns the Configuration field value
func (o *SamlApp) GetConfiguration() ConfigurationSaml {
	if o == nil {
		var ret ConfigurationSaml
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *SamlApp) GetConfigurationOk() (*ConfigurationSaml, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// SetConfiguration sets field value
func (o *SamlApp) SetConfiguration(v ConfigurationSaml) {
	o.Configuration = v
}

// GetSso returns the Sso field value if set, zero value otherwise.
func (o *SamlApp) GetSso() SsoSaml {
	if o == nil || isNil(o.Sso) {
		var ret SsoSaml
		return ret
	}
	return *o.Sso
}

// GetSsoOk returns a tuple with the Sso field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlApp) GetSsoOk() (*SsoSaml, bool) {
	if o == nil || isNil(o.Sso) {
		return nil, false
	}
	return o.Sso, true
}

// HasSso returns a boolean if a field has been set.
func (o *SamlApp) HasSso() bool {
	if o != nil && !isNil(o.Sso) {
		return true
	}

	return false
}

// SetSso gets a reference to the given SsoSaml and assigns it to the Sso field.
func (o *SamlApp) SetSso(v SsoSaml) {
	o.Sso = &v
}

func (o SamlApp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamlApp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: id is readOnly
	toSerialize["name"] = o.Name
	toSerialize["visible"] = o.Visible
	toSerialize["description"] = o.Description
	if !isNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !isNil(o.IconUrl) {
		toSerialize["icon_url"] = o.IconUrl
	}
	if !isNil(o.AuthMethod) {
		toSerialize["auth_method"] = o.AuthMethod
	}
	toSerialize["policy_id"] = o.PolicyId
	if !isNil(o.AllowAssumedSignin) {
		toSerialize["allow_assumed_signin"] = o.AllowAssumedSignin
	}
	if !isNil(o.TabId) {
		toSerialize["tab_id"] = o.TabId
	}
	toSerialize["connector_id"] = o.ConnectorId
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !isNil(o.RoleIds) {
		toSerialize["role_ids"] = o.RoleIds
	}
	if !isNil(o.Provisioning) {
		toSerialize["provisioning"] = o.Provisioning
	}
	toSerialize["parameters"] = o.Parameters
	if !isNil(o.EnforcementPoint) {
		toSerialize["enforcement_point"] = o.EnforcementPoint
	}
	toSerialize["configuration"] = o.Configuration
	if !isNil(o.Sso) {
		toSerialize["sso"] = o.Sso
	}
	return toSerialize, nil
}

type NullableSamlApp struct {
	value *SamlApp
	isSet bool
}

func (v NullableSamlApp) Get() *SamlApp {
	return v.value
}

func (v *NullableSamlApp) Set(val *SamlApp) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlApp) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlApp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlApp(val *SamlApp) *NullableSamlApp {
	return &NullableSamlApp{value: val, isSet: true}
}

func (v NullableSamlApp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlApp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

