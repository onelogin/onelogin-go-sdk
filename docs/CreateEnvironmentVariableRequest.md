# CreateEnvironmentVariableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name for the environment variable that will be used to retrieve the value from a hook function. | 
**Value** | **string** | The secret value that will be encrypted at rest and injected in applicable hook functions at run time. | 

## Methods

### NewCreateEnvironmentVariableRequest

`func NewCreateEnvironmentVariableRequest(name string, value string, ) *CreateEnvironmentVariableRequest`

NewCreateEnvironmentVariableRequest instantiates a new CreateEnvironmentVariableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEnvironmentVariableRequestWithDefaults

`func NewCreateEnvironmentVariableRequestWithDefaults() *CreateEnvironmentVariableRequest`

NewCreateEnvironmentVariableRequestWithDefaults instantiates a new CreateEnvironmentVariableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateEnvironmentVariableRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateEnvironmentVariableRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateEnvironmentVariableRequest) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *CreateEnvironmentVariableRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CreateEnvironmentVariableRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CreateEnvironmentVariableRequest) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


