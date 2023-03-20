# AddRolesToUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoleIdArray** | **[]int32** | Set to an array of one or more role IDs. The IDs must be positive integers. | 

## Methods

### NewAddRolesToUserRequest

`func NewAddRolesToUserRequest(roleIdArray []int32, ) *AddRolesToUserRequest`

NewAddRolesToUserRequest instantiates a new AddRolesToUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRolesToUserRequestWithDefaults

`func NewAddRolesToUserRequestWithDefaults() *AddRolesToUserRequest`

NewAddRolesToUserRequestWithDefaults instantiates a new AddRolesToUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoleIdArray

`func (o *AddRolesToUserRequest) GetRoleIdArray() []int32`

GetRoleIdArray returns the RoleIdArray field if non-nil, zero value otherwise.

### GetRoleIdArrayOk

`func (o *AddRolesToUserRequest) GetRoleIdArrayOk() (*[]int32, bool)`

GetRoleIdArrayOk returns a tuple with the RoleIdArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIdArray

`func (o *AddRolesToUserRequest) SetRoleIdArray(v []int32)`

SetRoleIdArray sets RoleIdArray field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


