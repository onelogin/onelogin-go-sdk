# GetAssignedUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** |  | [optional] 
**Users** | Pointer to **[]int32** |  | [optional] 
**BeforeCursor** | Pointer to **NullableInt32** |  | [optional] 
**PreviousLink** | Pointer to **NullableString** |  | [optional] 
**AfterCursor** | Pointer to **NullableInt32** |  | [optional] 
**NextLink** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetAssignedUser200Response

`func NewGetAssignedUser200Response() *GetAssignedUser200Response`

NewGetAssignedUser200Response instantiates a new GetAssignedUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAssignedUser200ResponseWithDefaults

`func NewGetAssignedUser200ResponseWithDefaults() *GetAssignedUser200Response`

NewGetAssignedUser200ResponseWithDefaults instantiates a new GetAssignedUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetAssignedUser200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetAssignedUser200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetAssignedUser200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetAssignedUser200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsers

`func (o *GetAssignedUser200Response) GetUsers() []int32`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GetAssignedUser200Response) GetUsersOk() (*[]int32, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GetAssignedUser200Response) SetUsers(v []int32)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *GetAssignedUser200Response) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetBeforeCursor

`func (o *GetAssignedUser200Response) GetBeforeCursor() int32`

GetBeforeCursor returns the BeforeCursor field if non-nil, zero value otherwise.

### GetBeforeCursorOk

`func (o *GetAssignedUser200Response) GetBeforeCursorOk() (*int32, bool)`

GetBeforeCursorOk returns a tuple with the BeforeCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeCursor

`func (o *GetAssignedUser200Response) SetBeforeCursor(v int32)`

SetBeforeCursor sets BeforeCursor field to given value.

### HasBeforeCursor

`func (o *GetAssignedUser200Response) HasBeforeCursor() bool`

HasBeforeCursor returns a boolean if a field has been set.

### SetBeforeCursorNil

`func (o *GetAssignedUser200Response) SetBeforeCursorNil(b bool)`

 SetBeforeCursorNil sets the value for BeforeCursor to be an explicit nil

### UnsetBeforeCursor
`func (o *GetAssignedUser200Response) UnsetBeforeCursor()`

UnsetBeforeCursor ensures that no value is present for BeforeCursor, not even an explicit nil
### GetPreviousLink

`func (o *GetAssignedUser200Response) GetPreviousLink() string`

GetPreviousLink returns the PreviousLink field if non-nil, zero value otherwise.

### GetPreviousLinkOk

`func (o *GetAssignedUser200Response) GetPreviousLinkOk() (*string, bool)`

GetPreviousLinkOk returns a tuple with the PreviousLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousLink

`func (o *GetAssignedUser200Response) SetPreviousLink(v string)`

SetPreviousLink sets PreviousLink field to given value.

### HasPreviousLink

`func (o *GetAssignedUser200Response) HasPreviousLink() bool`

HasPreviousLink returns a boolean if a field has been set.

### SetPreviousLinkNil

`func (o *GetAssignedUser200Response) SetPreviousLinkNil(b bool)`

 SetPreviousLinkNil sets the value for PreviousLink to be an explicit nil

### UnsetPreviousLink
`func (o *GetAssignedUser200Response) UnsetPreviousLink()`

UnsetPreviousLink ensures that no value is present for PreviousLink, not even an explicit nil
### GetAfterCursor

`func (o *GetAssignedUser200Response) GetAfterCursor() int32`

GetAfterCursor returns the AfterCursor field if non-nil, zero value otherwise.

### GetAfterCursorOk

`func (o *GetAssignedUser200Response) GetAfterCursorOk() (*int32, bool)`

GetAfterCursorOk returns a tuple with the AfterCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterCursor

`func (o *GetAssignedUser200Response) SetAfterCursor(v int32)`

SetAfterCursor sets AfterCursor field to given value.

### HasAfterCursor

`func (o *GetAssignedUser200Response) HasAfterCursor() bool`

HasAfterCursor returns a boolean if a field has been set.

### SetAfterCursorNil

`func (o *GetAssignedUser200Response) SetAfterCursorNil(b bool)`

 SetAfterCursorNil sets the value for AfterCursor to be an explicit nil

### UnsetAfterCursor
`func (o *GetAssignedUser200Response) UnsetAfterCursor()`

UnsetAfterCursor ensures that no value is present for AfterCursor, not even an explicit nil
### GetNextLink

`func (o *GetAssignedUser200Response) GetNextLink() string`

GetNextLink returns the NextLink field if non-nil, zero value otherwise.

### GetNextLinkOk

`func (o *GetAssignedUser200Response) GetNextLinkOk() (*string, bool)`

GetNextLinkOk returns a tuple with the NextLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextLink

`func (o *GetAssignedUser200Response) SetNextLink(v string)`

SetNextLink sets NextLink field to given value.

### HasNextLink

`func (o *GetAssignedUser200Response) HasNextLink() bool`

HasNextLink returns a boolean if a field has been set.

### SetNextLinkNil

`func (o *GetAssignedUser200Response) SetNextLinkNil(b bool)`

 SetNextLinkNil sets the value for NextLink to be an explicit nil

### UnsetNextLink
`func (o *GetAssignedUser200Response) UnsetNextLink()`

UnsetNextLink ensures that no value is present for NextLink, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


