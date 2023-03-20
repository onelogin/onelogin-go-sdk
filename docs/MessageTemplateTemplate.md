# MessageTemplateTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | **string** | Custom Email Subject | 
**Html** | **string** | The HTML body of the Custom Email | 
**Plain** | **string** | The Plain text body of the email | 
**Message** | **string** | The body of the SMS message. Max length 160 characters. | 

## Methods

### NewMessageTemplateTemplate

`func NewMessageTemplateTemplate(subject string, html string, plain string, message string, ) *MessageTemplateTemplate`

NewMessageTemplateTemplate instantiates a new MessageTemplateTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageTemplateTemplateWithDefaults

`func NewMessageTemplateTemplateWithDefaults() *MessageTemplateTemplate`

NewMessageTemplateTemplateWithDefaults instantiates a new MessageTemplateTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *MessageTemplateTemplate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *MessageTemplateTemplate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *MessageTemplateTemplate) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetHtml

`func (o *MessageTemplateTemplate) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *MessageTemplateTemplate) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *MessageTemplateTemplate) SetHtml(v string)`

SetHtml sets Html field to given value.


### GetPlain

`func (o *MessageTemplateTemplate) GetPlain() string`

GetPlain returns the Plain field if non-nil, zero value otherwise.

### GetPlainOk

`func (o *MessageTemplateTemplate) GetPlainOk() (*string, bool)`

GetPlainOk returns a tuple with the Plain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlain

`func (o *MessageTemplateTemplate) SetPlain(v string)`

SetPlain sets Plain field to given value.


### GetMessage

`func (o *MessageTemplateTemplate) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MessageTemplateTemplate) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MessageTemplateTemplate) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


