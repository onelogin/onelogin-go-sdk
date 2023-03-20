# Brand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Enabled** | **bool** | Indicates if the brand is enabled or not | [default to false]
**CustomSupportEnabled** | **bool** | Indicates if the custom support is enabled. If enabled, the login page includes the ability to submit a support request. | 
**CustomColor** | **string** | Primary brand color | 
**CustomAccentColor** | **string** | Secondary brand color | 
**CustomMaskingColor** | **string** | Color for the masking layer above the background image of the branded login screen. | 
**CustomMaskingOpacity** | **int32** | Opacity for the custom_masking_color. | 
**MfaEnrollmentMessage** | **string** | Text that replaces the default text displayed on the initial screen of the MFA Registration. | 
**EnableCustomLabelForLoginScreen** | **bool** | Indicates if the custom Username/Email field label is enabled or not | 
**CustomLabelTextForLoginScreen** | **string** | Custom label for the Username/Email field on the login screen. See example here. | 
**LoginInstruction** | **string** | Text for the login instruction screen, styled in Markdown. | 
**LoginInstructionTitle** | **string** | Link text to show login instruction screen. | 
**HideOneloginFooter** | **bool** | Indicates if the OneLogin footer will appear at the bottom of the login page. | 
**Background** | [**BrandBackground**](BrandBackground.md) |  | 
**Logo** | [**BrandLogo**](BrandLogo.md) |  | 

## Methods

### NewBrand

`func NewBrand(id int32, enabled bool, customSupportEnabled bool, customColor string, customAccentColor string, customMaskingColor string, customMaskingOpacity int32, mfaEnrollmentMessage string, enableCustomLabelForLoginScreen bool, customLabelTextForLoginScreen string, loginInstruction string, loginInstructionTitle string, hideOneloginFooter bool, background BrandBackground, logo BrandLogo, ) *Brand`

NewBrand instantiates a new Brand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandWithDefaults

`func NewBrandWithDefaults() *Brand`

NewBrandWithDefaults instantiates a new Brand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Brand) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Brand) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Brand) SetId(v int32)`

SetId sets Id field to given value.


### GetEnabled

`func (o *Brand) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Brand) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Brand) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCustomSupportEnabled

`func (o *Brand) GetCustomSupportEnabled() bool`

GetCustomSupportEnabled returns the CustomSupportEnabled field if non-nil, zero value otherwise.

### GetCustomSupportEnabledOk

`func (o *Brand) GetCustomSupportEnabledOk() (*bool, bool)`

GetCustomSupportEnabledOk returns a tuple with the CustomSupportEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSupportEnabled

`func (o *Brand) SetCustomSupportEnabled(v bool)`

SetCustomSupportEnabled sets CustomSupportEnabled field to given value.


### GetCustomColor

`func (o *Brand) GetCustomColor() string`

GetCustomColor returns the CustomColor field if non-nil, zero value otherwise.

### GetCustomColorOk

`func (o *Brand) GetCustomColorOk() (*string, bool)`

GetCustomColorOk returns a tuple with the CustomColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomColor

`func (o *Brand) SetCustomColor(v string)`

SetCustomColor sets CustomColor field to given value.


### GetCustomAccentColor

`func (o *Brand) GetCustomAccentColor() string`

GetCustomAccentColor returns the CustomAccentColor field if non-nil, zero value otherwise.

### GetCustomAccentColorOk

`func (o *Brand) GetCustomAccentColorOk() (*string, bool)`

GetCustomAccentColorOk returns a tuple with the CustomAccentColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAccentColor

`func (o *Brand) SetCustomAccentColor(v string)`

SetCustomAccentColor sets CustomAccentColor field to given value.


### GetCustomMaskingColor

`func (o *Brand) GetCustomMaskingColor() string`

GetCustomMaskingColor returns the CustomMaskingColor field if non-nil, zero value otherwise.

### GetCustomMaskingColorOk

`func (o *Brand) GetCustomMaskingColorOk() (*string, bool)`

GetCustomMaskingColorOk returns a tuple with the CustomMaskingColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaskingColor

`func (o *Brand) SetCustomMaskingColor(v string)`

SetCustomMaskingColor sets CustomMaskingColor field to given value.


### GetCustomMaskingOpacity

`func (o *Brand) GetCustomMaskingOpacity() int32`

GetCustomMaskingOpacity returns the CustomMaskingOpacity field if non-nil, zero value otherwise.

### GetCustomMaskingOpacityOk

`func (o *Brand) GetCustomMaskingOpacityOk() (*int32, bool)`

GetCustomMaskingOpacityOk returns a tuple with the CustomMaskingOpacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaskingOpacity

`func (o *Brand) SetCustomMaskingOpacity(v int32)`

SetCustomMaskingOpacity sets CustomMaskingOpacity field to given value.


### GetMfaEnrollmentMessage

`func (o *Brand) GetMfaEnrollmentMessage() string`

GetMfaEnrollmentMessage returns the MfaEnrollmentMessage field if non-nil, zero value otherwise.

### GetMfaEnrollmentMessageOk

`func (o *Brand) GetMfaEnrollmentMessageOk() (*string, bool)`

GetMfaEnrollmentMessageOk returns a tuple with the MfaEnrollmentMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaEnrollmentMessage

`func (o *Brand) SetMfaEnrollmentMessage(v string)`

SetMfaEnrollmentMessage sets MfaEnrollmentMessage field to given value.


### GetEnableCustomLabelForLoginScreen

`func (o *Brand) GetEnableCustomLabelForLoginScreen() bool`

GetEnableCustomLabelForLoginScreen returns the EnableCustomLabelForLoginScreen field if non-nil, zero value otherwise.

### GetEnableCustomLabelForLoginScreenOk

`func (o *Brand) GetEnableCustomLabelForLoginScreenOk() (*bool, bool)`

GetEnableCustomLabelForLoginScreenOk returns a tuple with the EnableCustomLabelForLoginScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCustomLabelForLoginScreen

`func (o *Brand) SetEnableCustomLabelForLoginScreen(v bool)`

SetEnableCustomLabelForLoginScreen sets EnableCustomLabelForLoginScreen field to given value.


### GetCustomLabelTextForLoginScreen

`func (o *Brand) GetCustomLabelTextForLoginScreen() string`

GetCustomLabelTextForLoginScreen returns the CustomLabelTextForLoginScreen field if non-nil, zero value otherwise.

### GetCustomLabelTextForLoginScreenOk

`func (o *Brand) GetCustomLabelTextForLoginScreenOk() (*string, bool)`

GetCustomLabelTextForLoginScreenOk returns a tuple with the CustomLabelTextForLoginScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLabelTextForLoginScreen

`func (o *Brand) SetCustomLabelTextForLoginScreen(v string)`

SetCustomLabelTextForLoginScreen sets CustomLabelTextForLoginScreen field to given value.


### GetLoginInstruction

`func (o *Brand) GetLoginInstruction() string`

GetLoginInstruction returns the LoginInstruction field if non-nil, zero value otherwise.

### GetLoginInstructionOk

`func (o *Brand) GetLoginInstructionOk() (*string, bool)`

GetLoginInstructionOk returns a tuple with the LoginInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginInstruction

`func (o *Brand) SetLoginInstruction(v string)`

SetLoginInstruction sets LoginInstruction field to given value.


### GetLoginInstructionTitle

`func (o *Brand) GetLoginInstructionTitle() string`

GetLoginInstructionTitle returns the LoginInstructionTitle field if non-nil, zero value otherwise.

### GetLoginInstructionTitleOk

`func (o *Brand) GetLoginInstructionTitleOk() (*string, bool)`

GetLoginInstructionTitleOk returns a tuple with the LoginInstructionTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginInstructionTitle

`func (o *Brand) SetLoginInstructionTitle(v string)`

SetLoginInstructionTitle sets LoginInstructionTitle field to given value.


### GetHideOneloginFooter

`func (o *Brand) GetHideOneloginFooter() bool`

GetHideOneloginFooter returns the HideOneloginFooter field if non-nil, zero value otherwise.

### GetHideOneloginFooterOk

`func (o *Brand) GetHideOneloginFooterOk() (*bool, bool)`

GetHideOneloginFooterOk returns a tuple with the HideOneloginFooter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideOneloginFooter

`func (o *Brand) SetHideOneloginFooter(v bool)`

SetHideOneloginFooter sets HideOneloginFooter field to given value.


### GetBackground

`func (o *Brand) GetBackground() BrandBackground`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *Brand) GetBackgroundOk() (*BrandBackground, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *Brand) SetBackground(v BrandBackground)`

SetBackground sets Background field to given value.


### GetLogo

`func (o *Brand) GetLogo() BrandLogo`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *Brand) GetLogoOk() (*BrandLogo, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *Brand) SetLogo(v BrandLogo)`

SetLogo sets Logo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


