# Envvar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for the Hook Environment Variable | [optional] 
**Name** | Pointer to **string** | The name of the environment variable. | [optional] 
**CreatedAt** | Pointer to **string** | The ISO8601 formatted date that the environment variable was created. | [optional] 
**UpdatedAt** | Pointer to **string** | The ISO8601 formatted date that the environment variable was last updated. | [optional] 

## Methods

### NewEnvvar

`func NewEnvvar() *Envvar`

NewEnvvar instantiates a new Envvar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvvarWithDefaults

`func NewEnvvarWithDefaults() *Envvar`

NewEnvvarWithDefaults instantiates a new Envvar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Envvar) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Envvar) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Envvar) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Envvar) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Envvar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Envvar) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Envvar) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Envvar) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Envvar) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Envvar) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Envvar) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Envvar) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Envvar) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Envvar) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Envvar) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Envvar) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


