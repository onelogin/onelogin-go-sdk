# TrackRiskEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verb** | **string** | Verbs are used to distinguish between different types of events. | 
**Ip** | **string** | The IP address of the User&#39;s request. | 
**UserAgent** | **string** | The user agent of the User&#39;s request. | 
**User** | [**RiskUser**](RiskUser.md) |  | 
**Source** | Pointer to [**Source**](Source.md) |  | [optional] 
**Session** | Pointer to [**Session**](Session.md) |  | [optional] 
**Device** | Pointer to [**RiskDevice**](RiskDevice.md) |  | [optional] 
**Fp** | Pointer to **string** | Set to the value of the __tdli_fp cookie. | [optional] 
**Published** | Pointer to **string** | Date and time of the event in IS08601 format. Useful for preloading old events. Defaults to date time this API request is received. | [optional] 

## Methods

### NewTrackRiskEventRequest

`func NewTrackRiskEventRequest(verb string, ip string, userAgent string, user RiskUser, ) *TrackRiskEventRequest`

NewTrackRiskEventRequest instantiates a new TrackRiskEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackRiskEventRequestWithDefaults

`func NewTrackRiskEventRequestWithDefaults() *TrackRiskEventRequest`

NewTrackRiskEventRequestWithDefaults instantiates a new TrackRiskEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerb

`func (o *TrackRiskEventRequest) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *TrackRiskEventRequest) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *TrackRiskEventRequest) SetVerb(v string)`

SetVerb sets Verb field to given value.


### GetIp

`func (o *TrackRiskEventRequest) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *TrackRiskEventRequest) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *TrackRiskEventRequest) SetIp(v string)`

SetIp sets Ip field to given value.


### GetUserAgent

`func (o *TrackRiskEventRequest) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *TrackRiskEventRequest) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *TrackRiskEventRequest) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.


### GetUser

`func (o *TrackRiskEventRequest) GetUser() RiskUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TrackRiskEventRequest) GetUserOk() (*RiskUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TrackRiskEventRequest) SetUser(v RiskUser)`

SetUser sets User field to given value.


### GetSource

`func (o *TrackRiskEventRequest) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TrackRiskEventRequest) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TrackRiskEventRequest) SetSource(v Source)`

SetSource sets Source field to given value.

### HasSource

`func (o *TrackRiskEventRequest) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSession

`func (o *TrackRiskEventRequest) GetSession() Session`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *TrackRiskEventRequest) GetSessionOk() (*Session, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *TrackRiskEventRequest) SetSession(v Session)`

SetSession sets Session field to given value.

### HasSession

`func (o *TrackRiskEventRequest) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetDevice

`func (o *TrackRiskEventRequest) GetDevice() RiskDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *TrackRiskEventRequest) GetDeviceOk() (*RiskDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *TrackRiskEventRequest) SetDevice(v RiskDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *TrackRiskEventRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetFp

`func (o *TrackRiskEventRequest) GetFp() string`

GetFp returns the Fp field if non-nil, zero value otherwise.

### GetFpOk

`func (o *TrackRiskEventRequest) GetFpOk() (*string, bool)`

GetFpOk returns a tuple with the Fp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFp

`func (o *TrackRiskEventRequest) SetFp(v string)`

SetFp sets Fp field to given value.

### HasFp

`func (o *TrackRiskEventRequest) HasFp() bool`

HasFp returns a boolean if a field has been set.

### GetPublished

`func (o *TrackRiskEventRequest) GetPublished() string`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *TrackRiskEventRequest) GetPublishedOk() (*string, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *TrackRiskEventRequest) SetPublished(v string)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *TrackRiskEventRequest) HasPublished() bool`

HasPublished returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


