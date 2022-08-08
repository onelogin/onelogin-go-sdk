# GetRiskScoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | The IP address of the User&#39;s request. | 
**UserAgent** | **string** | The user agent of the User&#39;s request. | 
**User** | [**RiskUser**](RiskUser.md) |  | 
**Source** | Pointer to [**Source**](Source.md) |  | [optional] 
**Session** | Pointer to [**Session**](Session.md) |  | [optional] 
**Device** | Pointer to [**RiskDevice**](RiskDevice.md) |  | [optional] 
**Fp** | Pointer to **string** | Set to the value of the __tdli_fp cookie. | [optional] 

## Methods

### NewGetRiskScoreRequest

`func NewGetRiskScoreRequest(ip string, userAgent string, user RiskUser, ) *GetRiskScoreRequest`

NewGetRiskScoreRequest instantiates a new GetRiskScoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRiskScoreRequestWithDefaults

`func NewGetRiskScoreRequestWithDefaults() *GetRiskScoreRequest`

NewGetRiskScoreRequestWithDefaults instantiates a new GetRiskScoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *GetRiskScoreRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetRiskScoreRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetRiskScoreRequest) SetIp(v string)`

SetIp sets Ip field to given value.


### GetUserAgent

`func (o *GetRiskScoreRequest) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *GetRiskScoreRequest) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *GetRiskScoreRequest) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.


### GetUser

`func (o *GetRiskScoreRequest) GetUser() RiskUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetRiskScoreRequest) GetUserOk() (*RiskUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetRiskScoreRequest) SetUser(v RiskUser)`

SetUser sets User field to given value.


### GetSource

`func (o *GetRiskScoreRequest) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetRiskScoreRequest) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetRiskScoreRequest) SetSource(v Source)`

SetSource sets Source field to given value.

### HasSource

`func (o *GetRiskScoreRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSession

`func (o *GetRiskScoreRequest) GetSession() Session`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *GetRiskScoreRequest) GetSessionOk() (*Session, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *GetRiskScoreRequest) SetSession(v Session)`

SetSession sets Session field to given value.

### HasSession

`func (o *GetRiskScoreRequest) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetDevice

`func (o *GetRiskScoreRequest) GetDevice() RiskDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *GetRiskScoreRequest) GetDeviceOk() (*RiskDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *GetRiskScoreRequest) SetDevice(v RiskDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *GetRiskScoreRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetFp

`func (o *GetRiskScoreRequest) GetFp() string`

GetFp returns the Fp field if non-nil, zero value otherwise.

### GetFpOk

`func (o *GetRiskScoreRequest) GetFpOk() (*string, bool)`

GetFpOk returns a tuple with the Fp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFp

`func (o *GetRiskScoreRequest) SetFp(v string)`

SetFp sets Fp field to given value.

### HasFp

`func (o *GetRiskScoreRequest) HasFp() bool`

HasFp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


