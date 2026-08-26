package onelogin

import (
	"context"

	mod "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
	utl "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/utilities"
)

const (
	PoliciesPath string = "api/2/policies"
)

// Security policies are deliberately passed as generic bodies rather than a
// models.Policy struct. A policy has some seventy writable attributes split
// across two kinds, and modelling them here would put a second copy of that
// list in this repo -- one that has to stay in step with every consumer's own
// schema, and that needs an SDK release before a caller can set a field the
// API already accepts. A map keeps this package to what it is good at: paths,
// auth, and turning a response into something a caller can read.
//
// These all report errors through CheckHTTPResponseWithErrorBody rather than
// CheckHTTPResponse. On an endpoint this wide, "request failed with status:
// 422" does not say which of the seventy attributes was refused, and the
// answer is sitting in the response body.

// GetPolicies lists the policies on the account. queryParams may be nil.
func (sdk *OneloginSDK) GetPolicies(queryParams mod.Queryable) (interface{}, error) {
	return sdk.GetPoliciesWithContext(context.Background(), queryParams)
}

func (sdk *OneloginSDK) GetPoliciesWithContext(ctx context.Context, queryParams mod.Queryable) (interface{}, error) {
	p, err := utl.BuildAPIPath(PoliciesPath)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.GetWithContext(ctx, &p, queryParams)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponseWithErrorBody(resp)
}

// GetPolicyByID returns a single policy with every attribute applicable to its
// kind. Fields belonging to the other kind are absent rather than empty, so a
// missing key does not mean the setting is off.
func (sdk *OneloginSDK) GetPolicyByID(policyID int) (interface{}, error) {
	return sdk.GetPolicyByIDWithContext(context.Background(), policyID)
}

func (sdk *OneloginSDK) GetPolicyByIDWithContext(ctx context.Context, policyID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(PoliciesPath, policyID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.GetWithContext(ctx, &p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponseWithErrorBody(resp)
}

// CreatePolicy creates a policy. The body must carry a name and a kind, either
// "user" or "app"; kind cannot be changed afterwards. Attributes belonging to
// the other kind are refused with a 422 naming the field.
func (sdk *OneloginSDK) CreatePolicy(policy map[string]interface{}) (interface{}, error) {
	return sdk.CreatePolicyWithContext(context.Background(), policy)
}

func (sdk *OneloginSDK) CreatePolicyWithContext(ctx context.Context, policy map[string]interface{}) (interface{}, error) {
	p, err := utl.BuildAPIPath(PoliciesPath)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.PostWithContext(ctx, &p, policy)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponseWithErrorBody(resp)
}

// UpdatePolicy applies a partial update. Only the attributes present in the
// body are touched, with one exception: the authentication factor id lists
// replace their whole set rather than merging into it. kind is immutable and
// sending it is an error.
func (sdk *OneloginSDK) UpdatePolicy(policyID int, policy map[string]interface{}) (interface{}, error) {
	return sdk.UpdatePolicyWithContext(context.Background(), policyID, policy)
}

func (sdk *OneloginSDK) UpdatePolicyWithContext(ctx context.Context, policyID int, policy map[string]interface{}) (interface{}, error) {
	p, err := utl.BuildAPIPath(PoliciesPath, policyID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.PutWithContext(ctx, &p, policy)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponseWithErrorBody(resp)
}

// DeletePolicy removes a policy. Users, groups and apps assigned to it are not
// deleted; their assignment is cleared and they fall back to the account
// default. The account default itself cannot be deleted.
func (sdk *OneloginSDK) DeletePolicy(policyID int) (interface{}, error) {
	return sdk.DeletePolicyWithContext(context.Background(), policyID)
}

func (sdk *OneloginSDK) DeletePolicyWithContext(ctx context.Context, policyID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(PoliciesPath, policyID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.DeleteWithContext(ctx, &p)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponseWithErrorBody(resp)
}

// SetDefaultPolicy makes a policy the account default. User policies only.
func (sdk *OneloginSDK) SetDefaultPolicy(policyID int) (interface{}, error) {
	return sdk.SetDefaultPolicyWithContext(context.Background(), policyID)
}

func (sdk *OneloginSDK) SetDefaultPolicyWithContext(ctx context.Context, policyID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(PoliciesPath, policyID, "set_default")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.PutWithContext(ctx, &p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponseWithErrorBody(resp)
}
