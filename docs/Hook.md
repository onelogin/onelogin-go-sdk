# Hook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Hook unique ID in OneLogin. | [optional] 
**Type** | **string** | A string describing the type of hook. e.g. &#x60;pre-authentication&#x60; | 
**Disabled** | **bool** | Boolean to enable or disable the hook. Disabled hooks will not run. | [default to true]
**Timeout** | **int32** | The number of seconds to allow the hook function to run before before timing out. Maximum timeout varies based on the type of hook. | [default to 1]
**EnvVars** | **[]string** | Environment Variable objects that will be available via process.env.ENV_VAR_NAME in the hook code. | 
**Runtime** | **string** | The Smart Hooks supported Node.js version to execute this hook with. | 
**Retries** | **int32** | Number of retries if execution fails. | [default to 0]
**Packages** | **map[string]interface{}** | An object containing NPM packages that are bundled with the hook function. | 
**Function** | **string** | A base64 encoded string containing the javascript function code. | 
**ContextVersion** | Pointer to **string** | The semantic version of the content that will be injected into this hook. | [optional] 
**Status** | Pointer to **string** | String describing the state of the hook function. When a hook is ready and disabled is false it will be executed. | [optional] 
**Options** | Pointer to [**HookOptions**](HookOptions.md) |  | [optional] 
**Conditions** | Pointer to [**[]HookConditionsInner**](HookConditionsInner.md) | An array of objects that let you limit the execution of a hook to users in specific roles. | [optional] 
**CreatedAt** | Pointer to **string** | ISO8601 format date that they hook function was created. | [optional] 
**UpdatedAt** | Pointer to **string** | ISO8601 format date that they hook function was last updated. | [optional] 

## Methods

### NewHook

`func NewHook(type_ string, disabled bool, timeout int32, envVars []string, runtime string, retries int32, packages map[string]interface{}, function string, ) *Hook`

NewHook instantiates a new Hook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHookWithDefaults

`func NewHookWithDefaults() *Hook`

NewHookWithDefaults instantiates a new Hook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Hook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Hook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Hook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Hook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *Hook) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Hook) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Hook) SetType(v string)`

SetType sets Type field to given value.


### GetDisabled

`func (o *Hook) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Hook) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Hook) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetTimeout

`func (o *Hook) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *Hook) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *Hook) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### GetEnvVars

`func (o *Hook) GetEnvVars() []string`

GetEnvVars returns the EnvVars field if non-nil, zero value otherwise.

### GetEnvVarsOk

`func (o *Hook) GetEnvVarsOk() (*[]string, bool)`

GetEnvVarsOk returns a tuple with the EnvVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVars

`func (o *Hook) SetEnvVars(v []string)`

SetEnvVars sets EnvVars field to given value.


### GetRuntime

`func (o *Hook) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *Hook) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *Hook) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.


### GetRetries

`func (o *Hook) GetRetries() int32`

GetRetries returns the Retries field if non-nil, zero value otherwise.

### GetRetriesOk

`func (o *Hook) GetRetriesOk() (*int32, bool)`

GetRetriesOk returns a tuple with the Retries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetries

`func (o *Hook) SetRetries(v int32)`

SetRetries sets Retries field to given value.


### GetPackages

`func (o *Hook) GetPackages() map[string]interface{}`

GetPackages returns the Packages field if non-nil, zero value otherwise.

### GetPackagesOk

`func (o *Hook) GetPackagesOk() (*map[string]interface{}, bool)`

GetPackagesOk returns a tuple with the Packages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackages

`func (o *Hook) SetPackages(v map[string]interface{})`

SetPackages sets Packages field to given value.


### GetFunction

`func (o *Hook) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *Hook) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *Hook) SetFunction(v string)`

SetFunction sets Function field to given value.


### GetContextVersion

`func (o *Hook) GetContextVersion() string`

GetContextVersion returns the ContextVersion field if non-nil, zero value otherwise.

### GetContextVersionOk

`func (o *Hook) GetContextVersionOk() (*string, bool)`

GetContextVersionOk returns a tuple with the ContextVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextVersion

`func (o *Hook) SetContextVersion(v string)`

SetContextVersion sets ContextVersion field to given value.

### HasContextVersion

`func (o *Hook) HasContextVersion() bool`

HasContextVersion returns a boolean if a field has been set.

### GetStatus

`func (o *Hook) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Hook) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Hook) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Hook) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOptions

`func (o *Hook) GetOptions() HookOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Hook) GetOptionsOk() (*HookOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Hook) SetOptions(v HookOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Hook) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetConditions

`func (o *Hook) GetConditions() []HookConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *Hook) GetConditionsOk() (*[]HookConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *Hook) SetConditions(v []HookConditionsInner)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *Hook) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Hook) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Hook) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Hook) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Hook) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Hook) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Hook) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Hook) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Hook) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


