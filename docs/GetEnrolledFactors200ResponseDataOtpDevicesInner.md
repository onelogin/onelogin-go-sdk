# GetEnrolledFactors200ResponseDataOtpDevicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | True &#x3D; enabled (used successfully for authentication at least once). False &#x3D; pending (registered but never used). | [optional] 
**Default** | Pointer to **bool** | True &#x3D; is user’s default MFA device for OneLogin. | [optional] 
**StateToken** | Pointer to **string** | A short lived token that is required to Verify the Factor. This token expires in 120 seconds. | [optional] 
**AuthFactorName** | Pointer to **string** | \&quot;Official\&quot; authentication factor name, as it appears to administrators in OneLogin. | [optional] 
**PhoneNumber** | Pointer to **string** | For OTP codes sent via SMS, the phone number receiving the SMS message. | [optional] 
**TypeDisplayName** | Pointer to **string** | Authentication factor display name as it appears to users upon initial registration, as defined by admins at Settings &gt; Authentication Factors. | [optional] 
**NeedsTrigger** | Pointer to **bool** | true: You MUST Activate this Factor to trigger an SMS or Push notification before Verifying the OTP code. false: No Activation required. You can Verify the OTP immediately. MFA factors that provide both push notifications (user accepts notification) and pull code submission (user initiates code submission from device or enters it manually) should appear twice; once with needs_trigger: true and once with needs_trigger: false. | [optional] 
**UserDisplayName** | Pointer to **string** | Authentication factor display name assigned by users when they enroll the device. | [optional] 
**Id** | Pointer to **int32** | MFA device identifier. | [optional] 

## Methods

### NewGetEnrolledFactors200ResponseDataOtpDevicesInner

`func NewGetEnrolledFactors200ResponseDataOtpDevicesInner() *GetEnrolledFactors200ResponseDataOtpDevicesInner`

NewGetEnrolledFactors200ResponseDataOtpDevicesInner instantiates a new GetEnrolledFactors200ResponseDataOtpDevicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEnrolledFactors200ResponseDataOtpDevicesInnerWithDefaults

`func NewGetEnrolledFactors200ResponseDataOtpDevicesInnerWithDefaults() *GetEnrolledFactors200ResponseDataOtpDevicesInner`

NewGetEnrolledFactors200ResponseDataOtpDevicesInnerWithDefaults instantiates a new GetEnrolledFactors200ResponseDataOtpDevicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDefault

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetStateToken

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) GetStateToken() string`

GetStateToken returns the StateToken field if non-nil, zero value otherwise.

### GetStateTokenOk

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) GetStateTokenOk() (*string, bool)`

GetStateTokenOk returns a tuple with the StateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateToken

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) SetStateToken(v string)`

SetStateToken sets StateToken field to given value.

### HasStateToken

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) HasStateToken() bool`

HasStateToken returns a boolean if a field has been set.

### GetAuthFactorName

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) GetAuthFactorName() string`

GetAuthFactorName returns the AuthFactorName field if non-nil, zero value otherwise.

### GetAuthFactorNameOk

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) GetAuthFactorNameOk() (*string, bool)`

GetAuthFactorNameOk returns a tuple with the AuthFactorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthFactorName

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) SetAuthFactorName(v string)`

SetAuthFactorName sets AuthFactorName field to given value.

### HasAuthFactorName

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) HasAuthFactorName() bool`

HasAuthFactorName returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetTypeDisplayName

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) GetTypeDisplayName() string`

GetTypeDisplayName returns the TypeDisplayName field if non-nil, zero value otherwise.

### GetTypeDisplayNameOk

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) GetTypeDisplayNameOk() (*string, bool)`

GetTypeDisplayNameOk returns a tuple with the TypeDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDisplayName

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) SetTypeDisplayName(v string)`

SetTypeDisplayName sets TypeDisplayName field to given value.

### HasTypeDisplayName

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) HasTypeDisplayName() bool`

HasTypeDisplayName returns a boolean if a field has been set.

### GetNeedsTrigger

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) GetNeedsTrigger() bool`

GetNeedsTrigger returns the NeedsTrigger field if non-nil, zero value otherwise.

### GetNeedsTriggerOk

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) GetNeedsTriggerOk() (*bool, bool)`

GetNeedsTriggerOk returns a tuple with the NeedsTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsTrigger

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) SetNeedsTrigger(v bool)`

SetNeedsTrigger sets NeedsTrigger field to given value.

### HasNeedsTrigger

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) HasNeedsTrigger() bool`

HasNeedsTrigger returns a boolean if a field has been set.

### GetUserDisplayName

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) GetUserDisplayName() string`

GetUserDisplayName returns the UserDisplayName field if non-nil, zero value otherwise.

### GetUserDisplayNameOk

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) GetUserDisplayNameOk() (*string, bool)`

GetUserDisplayNameOk returns a tuple with the UserDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDisplayName

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) SetUserDisplayName(v string)`

SetUserDisplayName sets UserDisplayName field to given value.

### HasUserDisplayName

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) HasUserDisplayName() bool`

HasUserDisplayName returns a boolean if a field has been set.

### GetId

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetEnrolledFactors200ResponseDataOtpDevicesInner) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


