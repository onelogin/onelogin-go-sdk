# Schema1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**AddedBy** | Pointer to [**Schema1AddedBy**](Schema1AddedBy.md) |  | [optional] 
**AddedAt** | Pointer to **string** |  | [optional] 
**Assigned** | Pointer to **bool** | Indicates if assigned to role or not. Defaults to true. | [optional] 

## Methods

### NewSchema1

`func NewSchema1() *Schema1`

NewSchema1 instantiates a new Schema1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchema1WithDefaults

`func NewSchema1WithDefaults() *Schema1`

NewSchema1WithDefaults instantiates a new Schema1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Schema1) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Schema1) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Schema1) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Schema1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Schema1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Schema1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Schema1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Schema1) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsername

`func (o *Schema1) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Schema1) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Schema1) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Schema1) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAddedBy

`func (o *Schema1) GetAddedBy() Schema1AddedBy`

GetAddedBy returns the AddedBy field if non-nil, zero value otherwise.

### GetAddedByOk

`func (o *Schema1) GetAddedByOk() (*Schema1AddedBy, bool)`

GetAddedByOk returns a tuple with the AddedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedBy

`func (o *Schema1) SetAddedBy(v Schema1AddedBy)`

SetAddedBy sets AddedBy field to given value.

### HasAddedBy

`func (o *Schema1) HasAddedBy() bool`

HasAddedBy returns a boolean if a field has been set.

### GetAddedAt

`func (o *Schema1) GetAddedAt() string`

GetAddedAt returns the AddedAt field if non-nil, zero value otherwise.

### GetAddedAtOk

`func (o *Schema1) GetAddedAtOk() (*string, bool)`

GetAddedAtOk returns a tuple with the AddedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAt

`func (o *Schema1) SetAddedAt(v string)`

SetAddedAt sets AddedAt field to given value.

### HasAddedAt

`func (o *Schema1) HasAddedAt() bool`

HasAddedAt returns a boolean if a field has been set.

### GetAssigned

`func (o *Schema1) GetAssigned() bool`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *Schema1) GetAssignedOk() (*bool, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *Schema1) SetAssigned(v bool)`

SetAssigned sets Assigned field to given value.

### HasAssigned

`func (o *Schema1) HasAssigned() bool`

HasAssigned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


