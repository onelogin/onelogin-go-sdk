# Privilege

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Privilege** | [**PrivilegePrivilege**](PrivilegePrivilege.md) |  | 

## Methods

### NewPrivilege

`func NewPrivilege(name string, privilege PrivilegePrivilege, ) *Privilege`

NewPrivilege instantiates a new Privilege object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegeWithDefaults

`func NewPrivilegeWithDefaults() *Privilege`

NewPrivilegeWithDefaults instantiates a new Privilege object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Privilege) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Privilege) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Privilege) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Privilege) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Privilege) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Privilege) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Privilege) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Privilege) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Privilege) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Privilege) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Privilege) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPrivilege

`func (o *Privilege) GetPrivilege() PrivilegePrivilege`

GetPrivilege returns the Privilege field if non-nil, zero value otherwise.

### GetPrivilegeOk

`func (o *Privilege) GetPrivilegeOk() (*PrivilegePrivilege, bool)`

GetPrivilegeOk returns a tuple with the Privilege field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilege

`func (o *Privilege) SetPrivilege(v PrivilegePrivilege)`

SetPrivilege sets Privilege field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


