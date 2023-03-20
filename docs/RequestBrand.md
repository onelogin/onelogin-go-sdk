# RequestBrand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates if the brand is enabled or not | [optional] [default to false]
**Name** | **string** | Brand’s name for humans. This isn’t related to subdomains. | 
**CustomSupportEnabled** | Pointer to **bool** | Indicates if the custom support is enabled. If enabled, the login page includes the ability to submit a support request. | [optional] 
**CustomColor** | Pointer to **string** | Primary brand color | [optional] 
**CustomAccentColor** | Pointer to **string** | Secondary brand color | [optional] 
**CustomMaskingColor** | Pointer to **string** | Color for the masking layer above the background image of the branded login screen. | [optional] 
**CustomMaskingOpacity** | Pointer to **int32** | Opacity for the custom_masking_color. | [optional] 
**EnableCustomLabelForLoginScreen** | Pointer to **bool** | Indicates if the custom Username/Email field label is enabled or not | [optional] 
**CustomLabelTextForLoginScreen** | Pointer to **string** | Custom label for the Username/Email field on the login screen. See example here. | [optional] 
**LoginInstructionTitle** | Pointer to **string** | Link text to show login instruction screen. | [optional] 
**LoginInstruction** | Pointer to **string** | Text for the login instruction screen, styled in Markdown. | [optional] 
**HideOneloginFooter** | Pointer to **bool** | Indicates if the OneLogin footer will appear at the bottom of the login page. | [optional] 
**MfaEnrollmentMessage** | Pointer to **string** | Text that replaces the default text displayed on the initial screen of the MFA Registration. | [optional] 
**Background** | Pointer to **string** | Base64 encoded image data (JPG/PNG, &lt;5MB) | [optional] 
**Logo** | Pointer to **string** | Base64 encoded image data (PNG, &lt;1MB) | [optional] 

## Methods

### NewRequestBrand

`func NewRequestBrand(name string, ) *RequestBrand`

NewRequestBrand instantiates a new RequestBrand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestBrandWithDefaults

`func NewRequestBrandWithDefaults() *RequestBrand`

NewRequestBrandWithDefaults instantiates a new RequestBrand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *RequestBrand) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RequestBrand) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RequestBrand) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RequestBrand) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *RequestBrand) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestBrand) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestBrand) SetName(v string)`

SetName sets Name field to given value.


### GetCustomSupportEnabled

`func (o *RequestBrand) GetCustomSupportEnabled() bool`

GetCustomSupportEnabled returns the CustomSupportEnabled field if non-nil, zero value otherwise.

### GetCustomSupportEnabledOk

`func (o *RequestBrand) GetCustomSupportEnabledOk() (*bool, bool)`

GetCustomSupportEnabledOk returns a tuple with the CustomSupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSupportEnabled

`func (o *RequestBrand) SetCustomSupportEnabled(v bool)`

SetCustomSupportEnabled sets CustomSupportEnabled field to given value.

### HasCustomSupportEnabled

`func (o *RequestBrand) HasCustomSupportEnabled() bool`

HasCustomSupportEnabled returns a boolean if a field has been set.

### GetCustomColor

`func (o *RequestBrand) GetCustomColor() string`

GetCustomColor returns the CustomColor field if non-nil, zero value otherwise.

### GetCustomColorOk

`func (o *RequestBrand) GetCustomColorOk() (*string, bool)`

GetCustomColorOk returns a tuple with the CustomColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomColor

`func (o *RequestBrand) SetCustomColor(v string)`

SetCustomColor sets CustomColor field to given value.

### HasCustomColor

`func (o *RequestBrand) HasCustomColor() bool`

HasCustomColor returns a boolean if a field has been set.

### GetCustomAccentColor

`func (o *RequestBrand) GetCustomAccentColor() string`

GetCustomAccentColor returns the CustomAccentColor field if non-nil, zero value otherwise.

### GetCustomAccentColorOk

`func (o *RequestBrand) GetCustomAccentColorOk() (*string, bool)`

GetCustomAccentColorOk returns a tuple with the CustomAccentColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAccentColor

`func (o *RequestBrand) SetCustomAccentColor(v string)`

SetCustomAccentColor sets CustomAccentColor field to given value.

### HasCustomAccentColor

`func (o *RequestBrand) HasCustomAccentColor() bool`

HasCustomAccentColor returns a boolean if a field has been set.

### GetCustomMaskingColor

`func (o *RequestBrand) GetCustomMaskingColor() string`

GetCustomMaskingColor returns the CustomMaskingColor field if non-nil, zero value otherwise.

### GetCustomMaskingColorOk

`func (o *RequestBrand) GetCustomMaskingColorOk() (*string, bool)`

GetCustomMaskingColorOk returns a tuple with the CustomMaskingColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaskingColor

`func (o *RequestBrand) SetCustomMaskingColor(v string)`

SetCustomMaskingColor sets CustomMaskingColor field to given value.

### HasCustomMaskingColor

`func (o *RequestBrand) HasCustomMaskingColor() bool`

HasCustomMaskingColor returns a boolean if a field has been set.

### GetCustomMaskingOpacity

`func (o *RequestBrand) GetCustomMaskingOpacity() int32`

GetCustomMaskingOpacity returns the CustomMaskingOpacity field if non-nil, zero value otherwise.

### GetCustomMaskingOpacityOk

`func (o *RequestBrand) GetCustomMaskingOpacityOk() (*int32, bool)`

GetCustomMaskingOpacityOk returns a tuple with the CustomMaskingOpacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaskingOpacity

`func (o *RequestBrand) SetCustomMaskingOpacity(v int32)`

SetCustomMaskingOpacity sets CustomMaskingOpacity field to given value.

### HasCustomMaskingOpacity

`func (o *RequestBrand) HasCustomMaskingOpacity() bool`

HasCustomMaskingOpacity returns a boolean if a field has been set.

### GetEnableCustomLabelForLoginScreen

`func (o *RequestBrand) GetEnableCustomLabelForLoginScreen() bool`

GetEnableCustomLabelForLoginScreen returns the EnableCustomLabelForLoginScreen field if non-nil, zero value otherwise.

### GetEnableCustomLabelForLoginScreenOk

`func (o *RequestBrand) GetEnableCustomLabelForLoginScreenOk() (*bool, bool)`

GetEnableCustomLabelForLoginScreenOk returns a tuple with the EnableCustomLabelForLoginScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCustomLabelForLoginScreen

`func (o *RequestBrand) SetEnableCustomLabelForLoginScreen(v bool)`

SetEnableCustomLabelForLoginScreen sets EnableCustomLabelForLoginScreen field to given value.

### HasEnableCustomLabelForLoginScreen

`func (o *RequestBrand) HasEnableCustomLabelForLoginScreen() bool`

HasEnableCustomLabelForLoginScreen returns a boolean if a field has been set.

### GetCustomLabelTextForLoginScreen

`func (o *RequestBrand) GetCustomLabelTextForLoginScreen() string`

GetCustomLabelTextForLoginScreen returns the CustomLabelTextForLoginScreen field if non-nil, zero value otherwise.

### GetCustomLabelTextForLoginScreenOk

`func (o *RequestBrand) GetCustomLabelTextForLoginScreenOk() (*string, bool)`

GetCustomLabelTextForLoginScreenOk returns a tuple with the CustomLabelTextForLoginScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLabelTextForLoginScreen

`func (o *RequestBrand) SetCustomLabelTextForLoginScreen(v string)`

SetCustomLabelTextForLoginScreen sets CustomLabelTextForLoginScreen field to given value.

### HasCustomLabelTextForLoginScreen

`func (o *RequestBrand) HasCustomLabelTextForLoginScreen() bool`

HasCustomLabelTextForLoginScreen returns a boolean if a field has been set.

### GetLoginInstructionTitle

`func (o *RequestBrand) GetLoginInstructionTitle() string`

GetLoginInstructionTitle returns the LoginInstructionTitle field if non-nil, zero value otherwise.

### GetLoginInstructionTitleOk

`func (o *RequestBrand) GetLoginInstructionTitleOk() (*string, bool)`

GetLoginInstructionTitleOk returns a tuple with the LoginInstructionTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginInstructionTitle

`func (o *RequestBrand) SetLoginInstructionTitle(v string)`

SetLoginInstructionTitle sets LoginInstructionTitle field to given value.

### HasLoginInstructionTitle

`func (o *RequestBrand) HasLoginInstructionTitle() bool`

HasLoginInstructionTitle returns a boolean if a field has been set.

### GetLoginInstruction

`func (o *RequestBrand) GetLoginInstruction() string`

GetLoginInstruction returns the LoginInstruction field if non-nil, zero value otherwise.

### GetLoginInstructionOk

`func (o *RequestBrand) GetLoginInstructionOk() (*string, bool)`

GetLoginInstructionOk returns a tuple with the LoginInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginInstruction

`func (o *RequestBrand) SetLoginInstruction(v string)`

SetLoginInstruction sets LoginInstruction field to given value.

### HasLoginInstruction

`func (o *RequestBrand) HasLoginInstruction() bool`

HasLoginInstruction returns a boolean if a field has been set.

### GetHideOneloginFooter

`func (o *RequestBrand) GetHideOneloginFooter() bool`

GetHideOneloginFooter returns the HideOneloginFooter field if non-nil, zero value otherwise.

### GetHideOneloginFooterOk

`func (o *RequestBrand) GetHideOneloginFooterOk() (*bool, bool)`

GetHideOneloginFooterOk returns a tuple with the HideOneloginFooter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideOneloginFooter

`func (o *RequestBrand) SetHideOneloginFooter(v bool)`

SetHideOneloginFooter sets HideOneloginFooter field to given value.

### HasHideOneloginFooter

`func (o *RequestBrand) HasHideOneloginFooter() bool`

HasHideOneloginFooter returns a boolean if a field has been set.

### GetMfaEnrollmentMessage

`func (o *RequestBrand) GetMfaEnrollmentMessage() string`

GetMfaEnrollmentMessage returns the MfaEnrollmentMessage field if non-nil, zero value otherwise.

### GetMfaEnrollmentMessageOk

`func (o *RequestBrand) GetMfaEnrollmentMessageOk() (*string, bool)`

GetMfaEnrollmentMessageOk returns a tuple with the MfaEnrollmentMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaEnrollmentMessage

`func (o *RequestBrand) SetMfaEnrollmentMessage(v string)`

SetMfaEnrollmentMessage sets MfaEnrollmentMessage field to given value.

### HasMfaEnrollmentMessage

`func (o *RequestBrand) HasMfaEnrollmentMessage() bool`

HasMfaEnrollmentMessage returns a boolean if a field has been set.

### GetBackground

`func (o *RequestBrand) GetBackground() string`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *RequestBrand) GetBackgroundOk() (*string, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *RequestBrand) SetBackground(v string)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *RequestBrand) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### GetLogo

`func (o *RequestBrand) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *RequestBrand) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *RequestBrand) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *RequestBrand) HasLogo() bool`

HasLogo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


