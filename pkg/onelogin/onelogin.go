package onelogin

import (
	"github.com/onelogin/onelogin-go-sdk/internal/api"
	olerror "github.com/onelogin/onelogin-go-sdk/internal/error"
)

// OneloginSDK represents the Onelogin SDK.
type OneloginSDK struct {
	Client *api.Client
}

// NewOneloginSDK creates a new instance of the Onelogin SDK.
func NewOneloginSDK() (*OneloginSDK, error) {
	client, err := api.NewClient()
	if err != nil {
		return nil, err
	}
	return &OneloginSDK{Client: client}, nil
}

// GetToken performs the authentication process using the env credentials.
func (sdk *OneloginSDK) GetToken() (string, error) {
	// Call the authenticator to perform the authentication process
	accessTk, err := sdk.Client.Auth.GetToken()
	if err != nil {
		return "", olerror.NewSDKError("Access Token retrieval unsuccessful")
	}
	return accessTk, nil
}

// MFA-related endpoints
func (sdk *OneloginSDK) ActivateMFAFactor(userID int, deviceID int) {
	// Implementation for Activate MFA Factor endpoint
}

func (sdk *OneloginSDK) EnrollMFAFactor(userID int) {
	// Implementation for Enroll MFA Factor endpoint
}

// Group-related endpoints
func (sdk *OneloginSDK) GetGroupByID(groupID int) {
	// Implementation for Get Group By ID endpoint
}

func (sdk *OneloginSDK) GetGroups() {
	// Implementation for Get Groups endpoint
}

// Invite-related endpoints
func (sdk *OneloginSDK) GenerateInviteLink() {
	// Implementation for Generate Invite Link endpoint
}

func (sdk *OneloginSDK) SendInviteLink() {
	// Implementation for Send Invite Link endpoint
}

// Connector-related endpoints
func (sdk *OneloginSDK) ListConnectors() {
	// Implementation for List Connectors endpoint
}

// App User-related endpoints
func (sdk *OneloginSDK) GetAppUsers(appID int) {
	// Implementation for Get App Users endpoint
}

func (sdk *OneloginSDK) CreateDeviceVerification(userID int) {
	// Implementation for Create Device Verification endpoint
}
func (sdk *OneloginSDK) CreateFactorRegistration(userID int)                   {}
func (sdk *OneloginSDK) DeleteEnrolledFactor(userID int, deviceID int)         {}
func (sdk *OneloginSDK) GenerateOTP(userID int)                                {}
func (sdk *OneloginSDK) GetAuthFactors(userID int)                             {}
func (sdk *OneloginSDK) GetAuthenticationDevices(userID int)                   {}
func (sdk *OneloginSDK) GetUserRegistration(userID int, registrationID int)    {}
func (sdk *OneloginSDK) GetUserVerification(userID int, verificationID int)    {}
func (sdk *OneloginSDK) VerifyUserRegistration(userID int, registrationID int) {}
func (sdk *OneloginSDK) VerifyUserVerification(userID int, verificationID int) {}
func (sdk *OneloginSDK) GenerateMFAToken(userID int)                           {}
func (sdk *OneloginSDK) GetEnrolledFactors(userID int)                         {}
func (sdk *OneloginSDK) GetMFAFactors(userID int)                              {}
func (sdk *OneloginSDK) RemoveMFAFactor(userID int, deviceID int)              {}
func (sdk *OneloginSDK) VerifyMFAFactor(userID int, deviceID int)              {}

func (sdk *OneloginSDK) AddPrivilegeToRole(privilegeID int, roleID int)      {}
func (sdk *OneloginSDK) AssignUsersToPrivilege(privilegeID int)              {}
func (sdk *OneloginSDK) CreatePrivilege()                                    {}
func (sdk *OneloginSDK) DeletePrivilege(privilegeID int)                     {}
func (sdk *OneloginSDK) DeleteRoleFromPrivilege(privilegeID int, roleID int) {}
func (sdk *OneloginSDK) GetAssignedUsers(privilegeID int)                    {}
func (sdk *OneloginSDK) GetPrivilege(privilegeID int)                        {}
func (sdk *OneloginSDK) GetPrivilegeRoles(privilegeID int)                   {}
func (sdk *OneloginSDK) ListPrivileges()                                     {}
func (sdk *OneloginSDK) RemovePrivilegeFromUser(privilegeID int, userID int) {}
func (sdk *OneloginSDK) UpdatePrivilege(privilegeID int)                     {}

func (sdk *OneloginSDK) GenerateSAMLAssertion() {}
func (sdk *OneloginSDK) VerifyFactorSAML()      {}

func (sdk *OneloginSDK) CreateEnvironmentVariable()             {}
func (sdk *OneloginSDK) CreateHook()                            {}
func (sdk *OneloginSDK) DeleteEnvironmentVariable(envVarID int) {}
func (sdk *OneloginSDK) DeleteHook(hookID int)                  {}
func (sdk *OneloginSDK) GetEnvironmentVariable(envVarID int)    {}
func (sdk *OneloginSDK) GetHook(hookID int)                     {}
func (sdk *OneloginSDK) GetHookLogs(hookID int)                 {}
func (sdk *OneloginSDK) ListEnvironmentVariables()              {}
func (sdk *OneloginSDK) ListSmartHooks()                        {}
func (sdk *OneloginSDK) UpdateEnvironmentVariable(envVarID int) {}
func (sdk *OneloginSDK) UpdateSmartHook(hookID int)             {}
func (sdk *OneloginSDK) CreateMapping()                         {}
func (sdk *OneloginSDK) DeleteMapping(mappingID int)            {}
func (sdk *OneloginSDK) GetMapping(mappingID int)               {}
func (sdk *OneloginSDK) ListMappings()                          {}
func (sdk *OneloginSDK) ListActions()                           {}
func (sdk *OneloginSDK) SortMappings()                          {}
func (sdk *OneloginSDK) UpdateMapping(mappingID int)            {}
func (sdk *OneloginSDK) AddRolesToUser(userID int)              {}
func (sdk *OneloginSDK) GetCustomAttributes()                   {}

func (sdk *OneloginSDK) RemoveUserRole(userID int) {}
func (sdk *OneloginSDK) SetUserState(userID int)   {}
