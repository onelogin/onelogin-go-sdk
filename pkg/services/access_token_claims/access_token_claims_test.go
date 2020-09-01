package accesstokenclaims

import (
	"encoding/json"
	"errors"
	"github.com/onelogin/onelogin-go-sdk/internal/test"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuery(t *testing.T) {
	tests := map[string]struct {
		queryPayload     *AccessTokenClaimsQuery
		expectedResponse []AccessTokenClaim
		expectedError    error
		repository       *test.MockRepository
	}{
		"It queries all claims assigend to an auth server": {
			queryPayload: &AccessTokenClaimsQuery{AuthServerID: "1"},
			expectedResponse: []AccessTokenClaim{
				AccessTokenClaim{
					ID:                       oltypes.Int32(int32(1)),
					AuthServerID:             oltypes.Int32(int32(1)),
					Label:                    oltypes.String("label"),
					UserAttributeMappings:    oltypes.String("mapping"),
					UserAttributeMacros:      oltypes.String("macro"),
					AttributeTransformations: oltypes.String("a"),
					SkipIfBlank:              oltypes.Bool(true),
					Values:                   []string{"values"},
					DefaultValues:            oltypes.String("default values"),
					ProvisionedEntitlements:  oltypes.Bool(true),
				},
				AccessTokenClaim{
					ID:                       oltypes.Int32(int32(2)),
					AuthServerID:             oltypes.Int32(int32(1)),
					Label:                    oltypes.String("label"),
					UserAttributeMappings:    oltypes.String("mapping"),
					UserAttributeMacros:      oltypes.String("macro"),
					AttributeTransformations: oltypes.String("a"),
					SkipIfBlank:              oltypes.Bool(true),
					Values:                   []string{"values"},
					DefaultValues:            oltypes.String("default values"),
					ProvisionedEntitlements:  oltypes.Bool(true),
				},
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal([]AccessTokenClaim{
						AccessTokenClaim{
							ID:                       oltypes.Int32(int32(1)),
							AuthServerID:             oltypes.Int32(int32(1)),
							Label:                    oltypes.String("label"),
							UserAttributeMappings:    oltypes.String("mapping"),
							UserAttributeMacros:      oltypes.String("macro"),
							AttributeTransformations: oltypes.String("a"),
							SkipIfBlank:              oltypes.Bool(true),
							Values:                   []string{"values"},
							DefaultValues:            oltypes.String("default values"),
							ProvisionedEntitlements:  oltypes.Bool(true),
						},
						AccessTokenClaim{
							ID:                       oltypes.Int32(int32(2)),
							AuthServerID:             oltypes.Int32(int32(1)),
							Label:                    oltypes.String("label"),
							UserAttributeMappings:    oltypes.String("mapping"),
							UserAttributeMacros:      oltypes.String("macro"),
							AttributeTransformations: oltypes.String("a"),
							SkipIfBlank:              oltypes.Bool(true),
							Values:                   []string{"values"},
							DefaultValues:            oltypes.String("default values"),
							ProvisionedEntitlements:  oltypes.Bool(true),
						},
					})
				},
			},
		},
		"it reports an error": {
			queryPayload: &AccessTokenClaimsQuery{AuthServerID: "1"},
			expectedResponse: []AccessTokenClaim{
				AccessTokenClaim{
					ID:                       oltypes.Int32(int32(1)),
					AuthServerID:             oltypes.Int32(int32(1)),
					Label:                    oltypes.String("label"),
					UserAttributeMappings:    oltypes.String("mapping"),
					UserAttributeMacros:      oltypes.String("macro"),
					AttributeTransformations: oltypes.String("a"),
					SkipIfBlank:              oltypes.Bool(true),
					Values:                   []string{"values"},
					DefaultValues:            oltypes.String("default values"),
					ProvisionedEntitlements:  oltypes.Bool(true),
				},
				AccessTokenClaim{
					ID:                       oltypes.Int32(int32(2)),
					AuthServerID:             oltypes.Int32(int32(1)),
					Label:                    oltypes.String("label"),
					UserAttributeMappings:    oltypes.String("mapping"),
					UserAttributeMacros:      oltypes.String("macro"),
					AttributeTransformations: oltypes.String("a"),
					SkipIfBlank:              oltypes.Bool(true),
					Values:                   []string{"values"},
					DefaultValues:            oltypes.String("default values"),
					ProvisionedEntitlements:  oltypes.Bool(true),
				},
			},
			expectedError: errors.New("error"),
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return nil, errors.New("error")
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			claims, err := svc.Query(test.queryPayload)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Equal(t, test.expectedResponse, claims)
			}
		})
	}
}

func TestGetOne(t *testing.T) {
	tests := map[string]struct {
		authServerID, accessTokenClaimID int32
		expectedError                    error
		expectedResponse                 AccessTokenClaim
		repository                       *test.MockRepository
	}{
		"it gets a claim by id": {
			authServerID:       int32(1),
			accessTokenClaimID: int32(1),
			expectedResponse: AccessTokenClaim{
				ID:                       oltypes.Int32(int32(1)),
				AuthServerID:             oltypes.Int32(int32(1)),
				Label:                    oltypes.String("label"),
				UserAttributeMappings:    oltypes.String("mapping"),
				UserAttributeMacros:      oltypes.String("macro"),
				AttributeTransformations: oltypes.String("a"),
				SkipIfBlank:              oltypes.Bool(true),
				Values:                   []string{"values"},
				DefaultValues:            oltypes.String("default values"),
				ProvisionedEntitlements:  oltypes.Bool(true),
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(
						AccessTokenClaim{
							ID:                       oltypes.Int32(int32(1)),
							AuthServerID:             oltypes.Int32(int32(1)),
							Label:                    oltypes.String("label"),
							UserAttributeMappings:    oltypes.String("mapping"),
							UserAttributeMacros:      oltypes.String("macro"),
							AttributeTransformations: oltypes.String("a"),
							SkipIfBlank:              oltypes.Bool(true),
							Values:                   []string{"values"},
							DefaultValues:            oltypes.String("default values"),
							ProvisionedEntitlements:  oltypes.Bool(true),
						},
					)
				},
			},
		},
		"it reports an error": {
			accessTokenClaimID: int32(1),
			expectedError:      errors.New("error"),
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return nil, errors.New("error")
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			claim, err := svc.GetOne(test.authServerID, test.accessTokenClaimID)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Equal(t, test.expectedResponse, *claim)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	tests := map[string]struct {
		payload          *AccessTokenClaim
		expectedError    error
		expectedResponse *AccessTokenClaim
		repository       *test.MockRepository
	}{
		"it creates a claim and associates it with an auth server": {
			payload: &AccessTokenClaim{
				AuthServerID:             oltypes.Int32(int32(1)),
				Label:                    oltypes.String("label"),
				UserAttributeMappings:    oltypes.String("mapping"),
				UserAttributeMacros:      oltypes.String("macro"),
				AttributeTransformations: oltypes.String("a"),
				SkipIfBlank:              oltypes.Bool(true),
				Values:                   []string{"values"},
				DefaultValues:            oltypes.String("default values"),
				ProvisionedEntitlements:  oltypes.Bool(true),
			},
			expectedResponse: &AccessTokenClaim{
				ID:                       oltypes.Int32(int32(1)),
				AuthServerID:             oltypes.Int32(int32(1)),
				Label:                    oltypes.String("label"),
				UserAttributeMappings:    oltypes.String("mapping"),
				UserAttributeMacros:      oltypes.String("macro"),
				AttributeTransformations: oltypes.String("a"),
				SkipIfBlank:              oltypes.Bool(true),
				Values:                   []string{"values"},
				DefaultValues:            oltypes.String("default values"),
				ProvisionedEntitlements:  oltypes.Bool(true),
			},
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(map[string]int32{"id": int32(1)})
				},
			},
		},
		"it returns an error": {
			payload: &AccessTokenClaim{
				AuthServerID:             oltypes.Int32(int32(1)),
				Label:                    oltypes.String("label"),
				UserAttributeMappings:    oltypes.String("mapping"),
				UserAttributeMacros:      oltypes.String("macro"),
				AttributeTransformations: oltypes.String("a"),
				SkipIfBlank:              oltypes.Bool(true),
				Values:                   []string{"values"},
				DefaultValues:            oltypes.String("default values"),
				ProvisionedEntitlements:  oltypes.Bool(true),
			},
			expectedError: errors.New("error"),
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					return nil, errors.New("error")
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			err := svc.Create(test.payload)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Equal(t, test.expectedResponse, test.payload)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := map[string]struct {
		authServerID     int32
		payload          *AccessTokenClaim
		expectedError    error
		expectedResponse *AccessTokenClaim
		repository       *test.MockRepository
	}{
		"it updates a claim": {
			payload: &AccessTokenClaim{
				ID:                       oltypes.Int32(int32(1)),
				AuthServerID:             oltypes.Int32(int32(1)),
				Label:                    oltypes.String("label"),
				UserAttributeMappings:    oltypes.String("mapping"),
				UserAttributeMacros:      oltypes.String("macro"),
				AttributeTransformations: oltypes.String("a"),
				SkipIfBlank:              oltypes.Bool(true),
				Values:                   []string{"values"},
				DefaultValues:            oltypes.String("default values"),
				ProvisionedEntitlements:  oltypes.Bool(true),
			},
			expectedResponse: &AccessTokenClaim{
				ID:                       oltypes.Int32(int32(1)),
				AuthServerID:             oltypes.Int32(int32(1)),
				Label:                    oltypes.String("label"),
				UserAttributeMappings:    oltypes.String("mapping"),
				UserAttributeMacros:      oltypes.String("macro"),
				AttributeTransformations: oltypes.String("a"),
				SkipIfBlank:              oltypes.Bool(true),
				Values:                   []string{"values"},
				DefaultValues:            oltypes.String("default values"),
				ProvisionedEntitlements:  oltypes.Bool(true),
			},
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(map[string]int32{"id": int32(1)})
				},
			},
		},
		"it reports an error": {
			payload: &AccessTokenClaim{
				ID:                       oltypes.Int32(int32(1)),
				AuthServerID:             oltypes.Int32(int32(1)),
				Label:                    oltypes.String("label"),
				UserAttributeMappings:    oltypes.String("mapping"),
				UserAttributeMacros:      oltypes.String("macro"),
				AttributeTransformations: oltypes.String("a"),
				SkipIfBlank:              oltypes.Bool(true),
				Values:                   []string{"values"},
				DefaultValues:            oltypes.String("default values"),
				ProvisionedEntitlements:  oltypes.Bool(true),
			},
			expectedError: errors.New("error"),
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return nil, errors.New("error")
				},
			},
		},
		"it reports an error if no parent resource id given": {
			payload: &AccessTokenClaim{
				ID:                       oltypes.Int32(int32(1)),
				Label:                    oltypes.String("label"),
				UserAttributeMappings:    oltypes.String("mapping"),
				UserAttributeMacros:      oltypes.String("macro"),
				AttributeTransformations: oltypes.String("a"),
				SkipIfBlank:              oltypes.Bool(true),
				Values:                   []string{"values"},
				DefaultValues:            oltypes.String("default values"),
				ProvisionedEntitlements:  oltypes.Bool(true),
			},
			expectedError: errors.New("Both ID and AuthServerID are required on the payload"),
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return nil, errors.New("error")
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			err := svc.Update(test.payload)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Equal(t, test.expectedResponse, test.payload)
			}
		})
	}
}

func TestDestroy(t *testing.T) {
	tests := map[string]struct {
		authServerID, accessTokenClaimID int32
		expectedError                    error
		repository                       *test.MockRepository
	}{
		"it deletes a claim": {
			authServerID:       int32(1),
			accessTokenClaimID: int32(1),
			repository: &test.MockRepository{
				DestroyFunc: func(r interface{}) ([]byte, error) {
					return nil, nil
				},
			},
		},
		"it reports an error": {
			accessTokenClaimID: int32(1),
			repository: &test.MockRepository{
				DestroyFunc: func(r interface{}) ([]byte, error) {
					return nil, errors.New("error")
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			err := svc.Destroy(test.authServerID, test.accessTokenClaimID)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}
