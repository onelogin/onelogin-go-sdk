# SetUserStateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **int32** | Set to the state value. Valid values include:   - 0 : Unapproved   - 1 : Approved   - 2 : Rejected   - 3 : Unlicensed | 

## Methods

### NewSetUserStateRequest

`func NewSetUserStateRequest(state int32, ) *SetUserStateRequest`

NewSetUserStateRequest instantiates a new SetUserStateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetUserStateRequestWithDefaults

`func NewSetUserStateRequestWithDefaults() *SetUserStateRequest`

NewSetUserStateRequestWithDefaults instantiates a new SetUserStateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *SetUserStateRequest) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SetUserStateRequest) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SetUserStateRequest) SetState(v int32)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


