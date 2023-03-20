# LockAccountUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LockedUntil** | **int32** | Set to the number of minutes for which you want to lock the user account. Set to 0 if you want to lock the user account based on the Lock effective period set in the policy assigned to the user. If no policy is assigned to the user, setting this value to 0 will lock the user’s account until you unlock it Note that this value can not be less time that the Lock Effective Period specified on a user policy. | 

## Methods

### NewLockAccountUserRequest

`func NewLockAccountUserRequest(lockedUntil int32, ) *LockAccountUserRequest`

NewLockAccountUserRequest instantiates a new LockAccountUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLockAccountUserRequestWithDefaults

`func NewLockAccountUserRequestWithDefaults() *LockAccountUserRequest`

NewLockAccountUserRequestWithDefaults instantiates a new LockAccountUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLockedUntil

`func (o *LockAccountUserRequest) GetLockedUntil() int32`

GetLockedUntil returns the LockedUntil field if non-nil, zero value otherwise.

### GetLockedUntilOk

`func (o *LockAccountUserRequest) GetLockedUntilOk() (*int32, bool)`

GetLockedUntilOk returns a tuple with the LockedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedUntil

`func (o *LockAccountUserRequest) SetLockedUntil(v int32)`

SetLockedUntil sets LockedUntil field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


