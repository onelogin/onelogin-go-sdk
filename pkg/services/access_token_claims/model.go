package accesstokenclaims

type AccessTokenClaimsQuery struct {
	AuthServerID string
}

type AccessTokenClaim struct {
	ID                       *int32
	AuthServerID             *int32
	Label                    *string
	UserAttributeMappings    *string
	UserAttributeMacros      *string
	AttributeTransformations *string
	SkipIfBlank              *bool
	Values                   []string
	DefaultValues            *string
	ProvisionedEntitlements  *bool
}
