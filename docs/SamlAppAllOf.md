# SamlAppAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | [**ConfigurationSaml**](ConfigurationSaml.md) |  | 
**Sso** | Pointer to [**SsoSaml**](SsoSaml.md) |  | [optional] 
**Parameters** | [**SamlAppAllOfParameters**](SamlAppAllOfParameters.md) |  | 

## Methods

### NewSamlAppAllOf

`func NewSamlAppAllOf(configuration ConfigurationSaml, parameters SamlAppAllOfParameters, ) *SamlAppAllOf`

NewSamlAppAllOf instantiates a new SamlAppAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlAppAllOfWithDefaults

`func NewSamlAppAllOfWithDefaults() *SamlAppAllOf`

NewSamlAppAllOfWithDefaults instantiates a new SamlAppAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *SamlAppAllOf) GetConfiguration() ConfigurationSaml`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *SamlAppAllOf) GetConfigurationOk() (*ConfigurationSaml, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *SamlAppAllOf) SetConfiguration(v ConfigurationSaml)`

SetConfiguration sets Configuration field to given value.


### GetSso

`func (o *SamlAppAllOf) GetSso() SsoSaml`

GetSso returns the Sso field if non-nil, zero value otherwise.

### GetSsoOk

`func (o *SamlAppAllOf) GetSsoOk() (*SsoSaml, bool)`

GetSsoOk returns a tuple with the Sso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSso

`func (o *SamlAppAllOf) SetSso(v SsoSaml)`

SetSso sets Sso field to given value.

### HasSso

`func (o *SamlAppAllOf) HasSso() bool`

HasSso returns a boolean if a field has been set.

### GetParameters

`func (o *SamlAppAllOf) GetParameters() SamlAppAllOfParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *SamlAppAllOf) GetParametersOk() (*SamlAppAllOfParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *SamlAppAllOf) SetParameters(v SamlAppAllOfParameters)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


