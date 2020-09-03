package clientapps

import (
	"encoding/json"
	"errors"
	"github.com/onelogin/onelogin-go-sdk/internal/test"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/auth_servers/scopes"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuery(t *testing.T) {
	tests := map[string]struct {
		queryPayload     *ClientAppsQuery
		expectedResponse []ClientApp
		expectedError    error
		repository       *test.MockRepository
	}{
		"It queries all clients assigend to an auth server": {
			queryPayload: &ClientAppsQuery{AuthServerID: "1"},
			expectedResponse: []ClientApp{
				ClientApp{
					ID:           oltypes.Int32(int32(1)),
					AuthServerID: oltypes.Int32(int32(1)),
					APIAuthID:    oltypes.Int32(int32(1)),
					Name:         oltypes.String("name"),
					Scopes: []scopes.Scope{
						scopes.Scope{
							ID:           oltypes.Int32(int32(1)),
							AuthServerID: oltypes.Int32(int32(1)),
							Value:        oltypes.String("value"),
							Description:  oltypes.String("description"),
						},
						scopes.Scope{
							ID:           oltypes.Int32(int32(2)),
							AuthServerID: oltypes.Int32(int32(1)),
							Value:        oltypes.String("value"),
							Description:  oltypes.String("description"),
						},
					},
				},
				ClientApp{
					ID:           oltypes.Int32(int32(2)),
					AuthServerID: oltypes.Int32(int32(1)),
					APIAuthID:    oltypes.Int32(int32(1)),
					Name:         oltypes.String("name"),
					Scopes: []scopes.Scope{
						scopes.Scope{
							ID:           oltypes.Int32(int32(3)),
							AuthServerID: oltypes.Int32(int32(1)),
							Value:        oltypes.String("value"),
							Description:  oltypes.String("description"),
						},
						scopes.Scope{
							ID:           oltypes.Int32(int32(4)),
							AuthServerID: oltypes.Int32(int32(1)),
							Value:        oltypes.String("value"),
							Description:  oltypes.String("description"),
						},
					},
				},
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal([]ClientApp{
						ClientApp{
							ID:           oltypes.Int32(int32(1)),
							AuthServerID: oltypes.Int32(int32(1)),
							APIAuthID:    oltypes.Int32(int32(1)),
							Name:         oltypes.String("name"),
							Scopes: []scopes.Scope{
								scopes.Scope{
									ID:           oltypes.Int32(int32(1)),
									AuthServerID: oltypes.Int32(int32(1)),
									Value:        oltypes.String("value"),
									Description:  oltypes.String("description"),
								},
								scopes.Scope{
									ID:           oltypes.Int32(int32(2)),
									AuthServerID: oltypes.Int32(int32(1)),
									Value:        oltypes.String("value"),
									Description:  oltypes.String("description"),
								},
							},
						},
						ClientApp{
							ID:           oltypes.Int32(int32(2)),
							AuthServerID: oltypes.Int32(int32(1)),
							APIAuthID:    oltypes.Int32(int32(1)),
							Name:         oltypes.String("name"),
							Scopes: []scopes.Scope{
								scopes.Scope{
									ID:           oltypes.Int32(int32(3)),
									AuthServerID: oltypes.Int32(int32(1)),
									Value:        oltypes.String("value"),
									Description:  oltypes.String("description"),
								},
								scopes.Scope{
									ID:           oltypes.Int32(int32(4)),
									AuthServerID: oltypes.Int32(int32(1)),
									Value:        oltypes.String("value"),
									Description:  oltypes.String("description"),
								},
							},
						},
					})
				},
			},
		},
		"it reports an error": {
			queryPayload:  &ClientAppsQuery{AuthServerID: "1"},
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
			clients, err := svc.Query(test.queryPayload)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Equal(t, test.expectedResponse, clients)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	tests := map[string]struct {
		payload          *ClientApp
		expectedError    error
		expectedResponse *ClientApp
		repository       *test.MockRepository
	}{
		"it creates a client and associates it with an auth server": {
			payload: &ClientApp{
				AuthServerID: oltypes.Int32(int32(1)),
				Name:         oltypes.String("name"),
				ScopeIDs:     []int32{int32(1), int32(2)},
			},
			expectedResponse: &ClientApp{
				ID:           oltypes.Int32(int32(1)),
				AuthServerID: oltypes.Int32(int32(1)),
				APIAuthID:    oltypes.Int32(int32(1)),
				Name:         oltypes.String("name"),
				ScopeIDs:     []int32{int32(1), int32(2)},
			},
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(map[string]int32{"app_id": int32(1), "api_auth_id": int32(1)})
				},
			},
		},
		"it errs out if parent ID not given on payload": {
			payload: &ClientApp{
				Name:     oltypes.String("name"),
				ScopeIDs: []int32{int32(1), int32(2)},
			},
			expectedError: errors.New("AuthServerID required on the payload"),
		},
		"it returns an error": {
			payload: &ClientApp{
				AuthServerID: oltypes.Int32(int32(1)),
				APIAuthID:    oltypes.Int32(int32(1)),
				Name:         oltypes.String("name"),
				ScopeIDs:     []int32{int32(3), int32(4)},
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
		payload          *ClientApp
		expectedError    error
		expectedResponse *ClientApp
		repository       *test.MockRepository
	}{
		"it updates a client": {
			payload: &ClientApp{
				ID:           oltypes.Int32(int32(1)),
				AuthServerID: oltypes.Int32(int32(1)),
				Name:         oltypes.String("name"),
				ScopeIDs:     []int32{int32(3), int32(4)},
			},
			expectedResponse: &ClientApp{
				ID:           oltypes.Int32(int32(1)),
				AuthServerID: oltypes.Int32(int32(1)),
				APIAuthID:    oltypes.Int32(int32(1)),
				Name:         oltypes.String("name"),
				ScopeIDs:     []int32{int32(3), int32(4)},
			},
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(map[string]int32{"app_id": int32(1), "api_auth_id": int32(1)})
				},
			},
		},
		"it reports an error": {
			payload: &ClientApp{
				ID:           oltypes.Int32(int32(1)),
				AuthServerID: oltypes.Int32(int32(1)),
				Name:         oltypes.String("name"),
				ScopeIDs:     []int32{int32(3), int32(4)},
			},
			expectedError: errors.New("error"),
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return nil, errors.New("error")
				},
			},
		},
		"it reports an error if no parent resource id given": {
			payload: &ClientApp{
				ID:       oltypes.Int32(int32(1)),
				Name:     oltypes.String("name"),
				ScopeIDs: []int32{int32(3), int32(4)},
			},
			expectedError: errors.New("Both ID and AuthServerID are required on the payload"),
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
		authServerID, clientID int32
		expectedError          error
		repository             *test.MockRepository
	}{
		"it deletes a client": {
			authServerID: int32(1),
			clientID:     int32(1),
			repository: &test.MockRepository{
				DestroyFunc: func(r interface{}) ([]byte, error) {
					return nil, nil
				},
			},
		},
		"it reports an error": {
			clientID: int32(1),
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
			err := svc.Destroy(test.authServerID, test.clientID)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}
