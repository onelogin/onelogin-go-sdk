# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **string** | A username for the user. | [optional] 
**Email** | Pointer to **string** | A valid email for the user. | [optional] 
**Firstname** | Pointer to **string** | The user&#39;s first name. | [optional] 
**Lastname** | Pointer to **string** | The user&#39;s last name. | [optional] 
**Title** | Pointer to **string** | The user&#39;s job title. | [optional] 
**Department** | Pointer to **string** | The user&#39;s department. | [optional] 
**Company** | Pointer to **string** | The user&#39;s company. | [optional] 
**Comment** | Pointer to **string** | Free text related to the user. | [optional] 
**GroupId** | Pointer to **int32** | The ID of the Group in OneLogin that the user is assigned to. | [optional] 
**RoleIds** | Pointer to **[]int32** | A list of OneLogin Role IDs of the user | [optional] 
**Phone** | Pointer to **string** | The E.164 format phone number for a user. | [optional] 
**State** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**DirectoryId** | Pointer to **int32** | The ID of the OneLogin Directory of the user. | [optional] 
**TrustedIdpId** | Pointer to **int32** | The ID of the OneLogin Trusted IDP of the user. | [optional] 
**ManagerAdId** | Pointer to **string** | The ID of the user&#39;s manager in Active Directory. | [optional] 
**ManagerUserId** | Pointer to **string** | The OneLogin User ID for the user&#39;s manager. | [optional] 
**SamaccountName** | Pointer to **string** | The user&#39;s Active Directory username. | [optional] 
**MemberOf** | Pointer to **string** | The user&#39;s directory membership. | [optional] 
**Userprincipalname** | Pointer to **string** | The principle name of the user. | [optional] 
**DistinguishedName** | Pointer to **string** | The distinguished name of the user. | [optional] 
**ExternalId** | Pointer to **string** | The ID of the user in an external directory. | [optional] 
**ActivatedAt** | Pointer to **string** |  | [optional] 
**LastLogin** | Pointer to **string** |  | [optional] 
**InvitationSentAt** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** |  | [optional] 
**PreferredLocaleCode** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**CustomAttributes** | Pointer to **map[string]interface{}** |  | [optional] 
**InvalidLoginAttempts** | Pointer to **int32** |  | [optional] 
**LockedUntil** | Pointer to **string** |  | [optional] 
**PasswordChangedAt** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** | The password to set for a user. | [optional] 
**PasswordConfirmation** | Pointer to **string** | Required if the password is being set. | [optional] 
**PasswordAlgorithm** | Pointer to **string** | Use this when importing a password that&#39;s already hashed. Prepend the salt value to the cleartext password value before SHA-256-encoding it | [optional] 
**Salt** | Pointer to **string** | The salt value used with the password_algorithm. | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstname

`func (o *User) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *User) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *User) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *User) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *User) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *User) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *User) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *User) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetTitle

`func (o *User) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *User) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *User) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *User) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDepartment

`func (o *User) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *User) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *User) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *User) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetCompany

`func (o *User) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *User) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *User) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *User) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetComment

`func (o *User) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *User) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *User) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *User) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetGroupId

`func (o *User) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *User) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *User) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *User) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetRoleIds

`func (o *User) GetRoleIds() []int32`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *User) GetRoleIdsOk() (*[]int32, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *User) SetRoleIds(v []int32)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *User) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### GetPhone

`func (o *User) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *User) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *User) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *User) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetState

`func (o *User) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *User) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *User) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *User) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStatus

`func (o *User) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *User) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *User) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *User) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDirectoryId

`func (o *User) GetDirectoryId() int32`

GetDirectoryId returns the DirectoryId field if non-nil, zero value otherwise.

### GetDirectoryIdOk

`func (o *User) GetDirectoryIdOk() (*int32, bool)`

GetDirectoryIdOk returns a tuple with the DirectoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectoryId

`func (o *User) SetDirectoryId(v int32)`

SetDirectoryId sets DirectoryId field to given value.

### HasDirectoryId

`func (o *User) HasDirectoryId() bool`

HasDirectoryId returns a boolean if a field has been set.

### GetTrustedIdpId

`func (o *User) GetTrustedIdpId() int32`

GetTrustedIdpId returns the TrustedIdpId field if non-nil, zero value otherwise.

### GetTrustedIdpIdOk

`func (o *User) GetTrustedIdpIdOk() (*int32, bool)`

GetTrustedIdpIdOk returns a tuple with the TrustedIdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedIdpId

`func (o *User) SetTrustedIdpId(v int32)`

SetTrustedIdpId sets TrustedIdpId field to given value.

### HasTrustedIdpId

`func (o *User) HasTrustedIdpId() bool`

HasTrustedIdpId returns a boolean if a field has been set.

### GetManagerAdId

`func (o *User) GetManagerAdId() string`

GetManagerAdId returns the ManagerAdId field if non-nil, zero value otherwise.

### GetManagerAdIdOk

`func (o *User) GetManagerAdIdOk() (*string, bool)`

GetManagerAdIdOk returns a tuple with the ManagerAdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerAdId

`func (o *User) SetManagerAdId(v string)`

SetManagerAdId sets ManagerAdId field to given value.

### HasManagerAdId

`func (o *User) HasManagerAdId() bool`

HasManagerAdId returns a boolean if a field has been set.

### GetManagerUserId

`func (o *User) GetManagerUserId() string`

GetManagerUserId returns the ManagerUserId field if non-nil, zero value otherwise.

### GetManagerUserIdOk

`func (o *User) GetManagerUserIdOk() (*string, bool)`

GetManagerUserIdOk returns a tuple with the ManagerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagerUserId

`func (o *User) SetManagerUserId(v string)`

SetManagerUserId sets ManagerUserId field to given value.

### HasManagerUserId

`func (o *User) HasManagerUserId() bool`

HasManagerUserId returns a boolean if a field has been set.

### GetSamaccountName

`func (o *User) GetSamaccountName() string`

GetSamaccountName returns the SamaccountName field if non-nil, zero value otherwise.

### GetSamaccountNameOk

`func (o *User) GetSamaccountNameOk() (*string, bool)`

GetSamaccountNameOk returns a tuple with the SamaccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamaccountName

`func (o *User) SetSamaccountName(v string)`

SetSamaccountName sets SamaccountName field to given value.

### HasSamaccountName

`func (o *User) HasSamaccountName() bool`

HasSamaccountName returns a boolean if a field has been set.

### GetMemberOf

`func (o *User) GetMemberOf() string`

GetMemberOf returns the MemberOf field if non-nil, zero value otherwise.

### GetMemberOfOk

`func (o *User) GetMemberOfOk() (*string, bool)`

GetMemberOfOk returns a tuple with the MemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberOf

`func (o *User) SetMemberOf(v string)`

SetMemberOf sets MemberOf field to given value.

### HasMemberOf

`func (o *User) HasMemberOf() bool`

HasMemberOf returns a boolean if a field has been set.

### GetUserprincipalname

`func (o *User) GetUserprincipalname() string`

GetUserprincipalname returns the Userprincipalname field if non-nil, zero value otherwise.

### GetUserprincipalnameOk

`func (o *User) GetUserprincipalnameOk() (*string, bool)`

GetUserprincipalnameOk returns a tuple with the Userprincipalname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserprincipalname

`func (o *User) SetUserprincipalname(v string)`

SetUserprincipalname sets Userprincipalname field to given value.

### HasUserprincipalname

`func (o *User) HasUserprincipalname() bool`

HasUserprincipalname returns a boolean if a field has been set.

### GetDistinguishedName

`func (o *User) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *User) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *User) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *User) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### GetExternalId

`func (o *User) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *User) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *User) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *User) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetActivatedAt

`func (o *User) GetActivatedAt() string`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *User) GetActivatedAtOk() (*string, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAt

`func (o *User) SetActivatedAt(v string)`

SetActivatedAt sets ActivatedAt field to given value.

### HasActivatedAt

`func (o *User) HasActivatedAt() bool`

HasActivatedAt returns a boolean if a field has been set.

### GetLastLogin

`func (o *User) GetLastLogin() string`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *User) GetLastLoginOk() (*string, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *User) SetLastLogin(v string)`

SetLastLogin sets LastLogin field to given value.

### HasLastLogin

`func (o *User) HasLastLogin() bool`

HasLastLogin returns a boolean if a field has been set.

### GetInvitationSentAt

`func (o *User) GetInvitationSentAt() string`

GetInvitationSentAt returns the InvitationSentAt field if non-nil, zero value otherwise.

### GetInvitationSentAtOk

`func (o *User) GetInvitationSentAtOk() (*string, bool)`

GetInvitationSentAtOk returns a tuple with the InvitationSentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationSentAt

`func (o *User) SetInvitationSentAt(v string)`

SetInvitationSentAt sets InvitationSentAt field to given value.

### HasInvitationSentAt

`func (o *User) HasInvitationSentAt() bool`

HasInvitationSentAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *User) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *User) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *User) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *User) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPreferredLocaleCode

`func (o *User) GetPreferredLocaleCode() string`

GetPreferredLocaleCode returns the PreferredLocaleCode field if non-nil, zero value otherwise.

### GetPreferredLocaleCodeOk

`func (o *User) GetPreferredLocaleCodeOk() (*string, bool)`

GetPreferredLocaleCodeOk returns a tuple with the PreferredLocaleCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLocaleCode

`func (o *User) SetPreferredLocaleCode(v string)`

SetPreferredLocaleCode sets PreferredLocaleCode field to given value.

### HasPreferredLocaleCode

`func (o *User) HasPreferredLocaleCode() bool`

HasPreferredLocaleCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *User) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *User) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCustomAttributes

`func (o *User) GetCustomAttributes() map[string]interface{}`

GetCustomAttributes returns the CustomAttributes field if non-nil, zero value otherwise.

### GetCustomAttributesOk

`func (o *User) GetCustomAttributesOk() (*map[string]interface{}, bool)`

GetCustomAttributesOk returns a tuple with the CustomAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAttributes

`func (o *User) SetCustomAttributes(v map[string]interface{})`

SetCustomAttributes sets CustomAttributes field to given value.

### HasCustomAttributes

`func (o *User) HasCustomAttributes() bool`

HasCustomAttributes returns a boolean if a field has been set.

### GetInvalidLoginAttempts

`func (o *User) GetInvalidLoginAttempts() int32`

GetInvalidLoginAttempts returns the InvalidLoginAttempts field if non-nil, zero value otherwise.

### GetInvalidLoginAttemptsOk

`func (o *User) GetInvalidLoginAttemptsOk() (*int32, bool)`

GetInvalidLoginAttemptsOk returns a tuple with the InvalidLoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidLoginAttempts

`func (o *User) SetInvalidLoginAttempts(v int32)`

SetInvalidLoginAttempts sets InvalidLoginAttempts field to given value.

### HasInvalidLoginAttempts

`func (o *User) HasInvalidLoginAttempts() bool`

HasInvalidLoginAttempts returns a boolean if a field has been set.

### GetLockedUntil

`func (o *User) GetLockedUntil() string`

GetLockedUntil returns the LockedUntil field if non-nil, zero value otherwise.

### GetLockedUntilOk

`func (o *User) GetLockedUntilOk() (*string, bool)`

GetLockedUntilOk returns a tuple with the LockedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedUntil

`func (o *User) SetLockedUntil(v string)`

SetLockedUntil sets LockedUntil field to given value.

### HasLockedUntil

`func (o *User) HasLockedUntil() bool`

HasLockedUntil returns a boolean if a field has been set.

### GetPasswordChangedAt

`func (o *User) GetPasswordChangedAt() string`

GetPasswordChangedAt returns the PasswordChangedAt field if non-nil, zero value otherwise.

### GetPasswordChangedAtOk

`func (o *User) GetPasswordChangedAtOk() (*string, bool)`

GetPasswordChangedAtOk returns a tuple with the PasswordChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChangedAt

`func (o *User) SetPasswordChangedAt(v string)`

SetPasswordChangedAt sets PasswordChangedAt field to given value.

### HasPasswordChangedAt

`func (o *User) HasPasswordChangedAt() bool`

HasPasswordChangedAt returns a boolean if a field has been set.

### GetPassword

`func (o *User) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *User) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *User) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *User) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordConfirmation

`func (o *User) GetPasswordConfirmation() string`

GetPasswordConfirmation returns the PasswordConfirmation field if non-nil, zero value otherwise.

### GetPasswordConfirmationOk

`func (o *User) GetPasswordConfirmationOk() (*string, bool)`

GetPasswordConfirmationOk returns a tuple with the PasswordConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordConfirmation

`func (o *User) SetPasswordConfirmation(v string)`

SetPasswordConfirmation sets PasswordConfirmation field to given value.

### HasPasswordConfirmation

`func (o *User) HasPasswordConfirmation() bool`

HasPasswordConfirmation returns a boolean if a field has been set.

### GetPasswordAlgorithm

`func (o *User) GetPasswordAlgorithm() string`

GetPasswordAlgorithm returns the PasswordAlgorithm field if non-nil, zero value otherwise.

### GetPasswordAlgorithmOk

`func (o *User) GetPasswordAlgorithmOk() (*string, bool)`

GetPasswordAlgorithmOk returns a tuple with the PasswordAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordAlgorithm

`func (o *User) SetPasswordAlgorithm(v string)`

SetPasswordAlgorithm sets PasswordAlgorithm field to given value.

### HasPasswordAlgorithm

`func (o *User) HasPasswordAlgorithm() bool`

HasPasswordAlgorithm returns a boolean if a field has been set.

### GetSalt

`func (o *User) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *User) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *User) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *User) HasSalt() bool`

HasSalt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


