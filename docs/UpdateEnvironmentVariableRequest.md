# UpdateEnvironmentVariableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | The secret value that will be encrypted at rest and injected in applicable hook functions at run time. | 

## Methods

### NewUpdateEnvironmentVariableRequest

`func NewUpdateEnvironmentVariableRequest(value string, ) *UpdateEnvironmentVariableRequest`

NewUpdateEnvironmentVariableRequest instantiates a new UpdateEnvironmentVariableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEnvironmentVariableRequestWithDefaults

`func NewUpdateEnvironmentVariableRequestWithDefaults() *UpdateEnvironmentVariableRequest`

NewUpdateEnvironmentVariableRequestWithDefaults instantiates a new UpdateEnvironmentVariableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *UpdateEnvironmentVariableRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateEnvironmentVariableRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateEnvironmentVariableRequest) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


