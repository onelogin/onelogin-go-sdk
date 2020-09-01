package scopes

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
		queryPayload     *ScopesQuery
		expectedResponse []Scope
		expectedError    error
		repository       *test.MockRepository
	}{
		"It queries all scopes assigend to an auth server": {
			queryPayload: &ScopesQuery{AuthServerID: "1"},
			expectedResponse: []Scope{
				Scope{
					ID:           oltypes.Int32(int32(1)),
					AuthServerID: oltypes.Int32(int32(1)),
					Value:        oltypes.String("value"),
					Description:  oltypes.String("description"),
				},
				Scope{
					ID:           oltypes.Int32(int32(1)),
					AuthServerID: oltypes.Int32(int32(1)),
					Value:        oltypes.String("value"),
					Description:  oltypes.String("description"),
				},
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal([]Scope{
						Scope{
							ID:           oltypes.Int32(int32(1)),
							AuthServerID: oltypes.Int32(int32(1)),
							Value:        oltypes.String("value"),
							Description:  oltypes.String("description"),
						},
						Scope{
							ID:           oltypes.Int32(int32(1)),
							AuthServerID: oltypes.Int32(int32(1)),
							Value:        oltypes.String("value"),
							Description:  oltypes.String("description"),
						},
					})
				},
			},
		},
		"it reports an error": {
			queryPayload: &ScopesQuery{AuthServerID: "1"},
			expectedResponse: []Scope{
				Scope{
					ID:           oltypes.Int32(int32(1)),
					AuthServerID: oltypes.Int32(int32(1)),
					Value:        oltypes.String("value"),
					Description:  oltypes.String("description"),
				},
				Scope{
					ID:           oltypes.Int32(int32(1)),
					AuthServerID: oltypes.Int32(int32(1)),
					Value:        oltypes.String("value"),
					Description:  oltypes.String("description"),
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
			scopes, err := svc.Query(test.queryPayload)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Equal(t, test.expectedResponse, scopes)
			}
		})
	}
}

func TestGetOne(t *testing.T) {
	tests := map[string]struct {
		authServerID, scopeID int32
		expectedError         error
		expectedResponse      Scope
		repository            *test.MockRepository
	}{
		"it gets a scope by id": {
			authServerID: int32(1),
			scopeID:      int32(1),
			expectedResponse: Scope{
				ID:           oltypes.Int32(int32(1)),
				AuthServerID: oltypes.Int32(int32(1)),
				Value:        oltypes.String("value"),
				Description:  oltypes.String("description"),
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(
						Scope{
							ID:           oltypes.Int32(int32(1)),
							AuthServerID: oltypes.Int32(int32(1)),
							Value:        oltypes.String("value"),
							Description:  oltypes.String("description"),
						},
					)
				},
			},
		},
		"it reports an error": {
			scopeID:       int32(1),
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
			scope, err := svc.GetOne(test.authServerID, test.scopeID)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Equal(t, test.expectedResponse, *scope)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	tests := map[string]struct {
		payload          *Scope
		expectedError    error
		expectedResponse *Scope
		repository       *test.MockRepository
	}{
		"it creates a scope and associates it with an auth server": {
			payload: &Scope{
				AuthServerID: oltypes.Int32(int32(1)),
				Value:        oltypes.String("value"),
				Description:  oltypes.String("description"),
			},
			expectedResponse: &Scope{
				ID:           oltypes.Int32(int32(1)),
				AuthServerID: oltypes.Int32(int32(1)),
				Value:        oltypes.String("value"),
				Description:  oltypes.String("description"),
			},
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(map[string]int32{"id": int32(1)})
				},
			},
		},
		"it returns an error": {
			payload: &Scope{
				ID:           oltypes.Int32(int32(1)),
				AuthServerID: oltypes.Int32(int32(1)),
				Value:        oltypes.String("value"),
				Description:  oltypes.String("description"),
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
		payload          *Scope
		expectedError    error
		expectedResponse *Scope
		repository       *test.MockRepository
	}{
		"it updates a scope": {
			payload: &Scope{
				ID:           oltypes.Int32(int32(1)),
				AuthServerID: oltypes.Int32(int32(1)),
				Value:        oltypes.String("value"),
				Description:  oltypes.String("description"),
			},
			expectedResponse: &Scope{
				ID:           oltypes.Int32(int32(1)),
				AuthServerID: oltypes.Int32(int32(1)),
				Value:        oltypes.String("value"),
				Description:  oltypes.String("description"),
			},
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(map[string]int32{"id": int32(1)})
				},
			},
		},
		"it reports an error": {
			payload: &Scope{
				ID:           oltypes.Int32(int32(1)),
				AuthServerID: oltypes.Int32(int32(1)),
				Value:        oltypes.String("value"),
				Description:  oltypes.String("description"),
			},
			expectedError: errors.New("error"),
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return nil, errors.New("error")
				},
			},
		},
		"it reports an error if no parent resource id given": {
			payload: &Scope{
				ID:          oltypes.Int32(int32(1)),
				Value:       oltypes.String("value"),
				Description: oltypes.String("description"),
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
		authServerID, scopeID int32
		expectedError         error
		repository            *test.MockRepository
	}{
		"it deletes a scope": {
			authServerID: int32(1),
			scopeID:      int32(1),
			repository: &test.MockRepository{
				DestroyFunc: func(r interface{}) ([]byte, error) {
					return nil, nil
				},
			},
		},
		"it reports an error": {
			scopeID: int32(1),
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
			err := svc.Destroy(test.authServerID, test.scopeID)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}
