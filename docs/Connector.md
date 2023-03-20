# Connector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Connectors unique ID in OneLogin. | [optional] 
**Name** | Pointer to **string** | Name of Connector | [optional] 
**IconUrl** | Pointer to **string** | A link to the icon&#39;s url. | [optional] 
**AuthMethod** | Pointer to [**AuthMethod**](AuthMethod.md) |  | [optional] 
**AllowsNewParameters** | Pointer to **bool** | Indicates if apps created using this connector will be allowed to create custom parameters. | [optional] 

## Methods

### NewConnector

`func NewConnector() *Connector`

NewConnector instantiates a new Connector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorWithDefaults

`func NewConnectorWithDefaults() *Connector`

NewConnectorWithDefaults instantiates a new Connector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Connector) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Connector) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Connector) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Connector) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Connector) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Connector) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Connector) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Connector) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIconUrl

`func (o *Connector) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *Connector) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *Connector) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *Connector) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetAuthMethod

`func (o *Connector) GetAuthMethod() AuthMethod`

GetAuthMethod returns the AuthMethod field if non-nil, zero value otherwise.

### GetAuthMethodOk

`func (o *Connector) GetAuthMethodOk() (*AuthMethod, bool)`

GetAuthMethodOk returns a tuple with the AuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMethod

`func (o *Connector) SetAuthMethod(v AuthMethod)`

SetAuthMethod sets AuthMethod field to given value.

### HasAuthMethod

`func (o *Connector) HasAuthMethod() bool`

HasAuthMethod returns a boolean if a field has been set.

### GetAllowsNewParameters

`func (o *Connector) GetAllowsNewParameters() bool`

GetAllowsNewParameters returns the AllowsNewParameters field if non-nil, zero value otherwise.

### GetAllowsNewParametersOk

`func (o *Connector) GetAllowsNewParametersOk() (*bool, bool)`

GetAllowsNewParametersOk returns a tuple with the AllowsNewParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowsNewParameters

`func (o *Connector) SetAllowsNewParameters(v bool)`

SetAllowsNewParameters sets AllowsNewParameters field to given value.

### HasAllowsNewParameters

`func (o *Connector) HasAllowsNewParameters() bool`

HasAllowsNewParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


