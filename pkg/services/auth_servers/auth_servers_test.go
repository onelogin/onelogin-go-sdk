package authservers

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
		queryPayload     *AuthServerQuery
		expectedResponse []AuthServer
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets one auth server": {
			queryPayload: &AuthServerQuery{Limit: "1"},
			expectedResponse: []AuthServer{
				AuthServer{
					ID:   oltypes.Int32(1),
					Name: oltypes.String("name"),
					Configuration: &AuthServerConfiguration{
						ResourceIdentifier: oltypes.String("example.com/contacts"),
						Audiences:          []string{"example.com/contacts", "example.com/people"},
					},
				},
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal([]AuthServer{
						AuthServer{
							ID:   oltypes.Int32(1),
							Name: oltypes.String("name"),
							Configuration: &AuthServerConfiguration{
								ResourceIdentifier: oltypes.String("example.com/contacts"),
								Audiences:          []string{"example.com/contacts", "example.com/people"},
							},
						},
					})
				},
			},
		},
		"it returns the remote default limit of auth servers if no query given": {
			queryPayload: &AuthServerQuery{},
			expectedResponse: []AuthServer{
				AuthServer{
					ID:   oltypes.Int32(1),
					Name: oltypes.String("name"),
					Configuration: &AuthServerConfiguration{
						ResourceIdentifier: oltypes.String("example.com/contacts"),
						Audiences:          []string{"example.com/contacts", "example.com/people"},
					},
				},
				AuthServer{
					ID:   oltypes.Int32(1),
					Name: oltypes.String("name2"),
					Configuration: &AuthServerConfiguration{
						ResourceIdentifier: oltypes.String("sample.com/contacts"),
						Audiences:          []string{"sample.com/contacts", "sample.com/people"},
					},
				},
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal([]AuthServer{
						AuthServer{
							ID:   oltypes.Int32(1),
							Name: oltypes.String("name"),
							Configuration: &AuthServerConfiguration{
								ResourceIdentifier: oltypes.String("example.com/contacts"),
								Audiences:          []string{"example.com/contacts", "example.com/people"},
							},
						},
						AuthServer{
							ID:   oltypes.Int32(1),
							Name: oltypes.String("name2"),
							Configuration: &AuthServerConfiguration{
								ResourceIdentifier: oltypes.String("sample.com/contacts"),
								Audiences:          []string{"sample.com/contacts", "sample.com/people"},
							},
						},
					})
				},
			},
		},
		"it returns an error if the call to /api_authorizations fails": {
			queryPayload:     &AuthServerQuery{},
			expectedError:    errors.New("error"),
			expectedResponse: nil,
			repository:       &test.MockRepository{},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			actual, err := svc.Query(test.queryPayload)
			assert.Equal(t, test.expectedResponse, actual)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestGetOne(t *testing.T) {
	tests := map[string]struct {
		id               int32
		expectedResponse *AuthServer
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets one auth server by id": {
			id: int32(1),
			expectedResponse: &AuthServer{
				ID:   oltypes.Int32(1),
				Name: oltypes.String("name"),
				Configuration: &AuthServerConfiguration{
					ResourceIdentifier: oltypes.String("example.com/contacts"),
					Audiences:          []string{"example.com/contacts", "example.com/people"},
				},
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(AuthServer{
						ID:   oltypes.Int32(1),
						Name: oltypes.String("name"),
						Configuration: &AuthServerConfiguration{
							ResourceIdentifier: oltypes.String("example.com/contacts"),
							Audiences:          []string{"example.com/contacts", "example.com/people"},
						},
					})
				},
			},
		},
		"it returns an error if the call to /api_authorizations fails": {
			expectedError:    errors.New("error"),
			expectedResponse: nil,
			repository:       &test.MockRepository{},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			actual, err := svc.GetOne(test.id)
			assert.Equal(t, test.expectedResponse, actual)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := map[string]struct {
		id             int32
		updatePayload  *AuthServer
		expectedResult *AuthServer
		expectedError  error
		repository     *test.MockRepository
	}{
		"it updates a auth server": {
			id:            int32(1),
			updatePayload: &AuthServer{Name: oltypes.String("update")},
			expectedResult: &AuthServer{
				ID:   oltypes.Int32(1),
				Name: oltypes.String("update"),
				Configuration: &AuthServerConfiguration{
					ResourceIdentifier: oltypes.String("example.com/contacts"),
					Audiences:          []string{"example.com/contacts", "example.com/people"},
				},
			},
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(AuthServer{
						ID:   oltypes.Int32(1),
						Name: oltypes.String("update"),
						Configuration: &AuthServerConfiguration{
							ResourceIdentifier: oltypes.String("example.com/contacts"),
							Audiences:          []string{"example.com/contacts", "example.com/people"},
						},
					})
				},
			},
		},
		"it returns an error if something went wrong": {
			id:             int32(2),
			updatePayload:  &AuthServer{ID: oltypes.Int32(1), Name: oltypes.String("update")},
			expectedResult: &AuthServer{},
			expectedError:  errors.New("error"),
			repository:     &test.MockRepository{},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			err := svc.Update(test.id, test.updatePayload)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Equal(t, test.expectedResult, test.updatePayload)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	tests := map[string]struct {
		createPayload  *AuthServer
		expectedResult *AuthServer
		expectedError  error
		repository     *test.MockRepository
	}{
		"it creates a auth server": {
			createPayload: &AuthServer{Name: oltypes.String("name")},
			expectedResult: &AuthServer{
				ID:   oltypes.Int32(1),
				Name: oltypes.String("name"),
				Configuration: &AuthServerConfiguration{
					ResourceIdentifier: oltypes.String("example.com/contacts"),
					Audiences:          []string{"example.com/contacts", "example.com/people"},
				},
			},
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(AuthServer{
						ID:   oltypes.Int32(1),
						Name: oltypes.String("name"),
						Configuration: &AuthServerConfiguration{
							ResourceIdentifier: oltypes.String("example.com/contacts"),
							Audiences:          []string{"example.com/contacts", "example.com/people"},
						},
					})
				},
			},
		},
		"it returns an error if something went wrong": {
			createPayload:  &AuthServer{Name: oltypes.String("name")},
			expectedResult: &AuthServer{},
			expectedError:  errors.New("error"),
			repository:     &test.MockRepository{},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			err := svc.Create(test.createPayload)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Equal(t, test.expectedResult, test.createPayload)
			}
		})
	}
}

func TestDestroy(t *testing.T) {
	tests := map[string]struct {
		id               int32
		repository       *test.MockRepository
		expectedResponse *AuthServer
		expectedError    error
	}{
		"it destroys one auth server": {
			id: int32(1),
			repository: &test.MockRepository{
				DestroyFunc: func(r interface{}) ([]byte, error) {
					return nil, nil
				},
			},
			expectedResponse: &AuthServer{},
		},
		"it returns an error if there is a problem finding the app": {
			id:               int32(2),
			repository:       &test.MockRepository{},
			expectedResponse: nil,
			expectedError:    errors.New("error"),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			err := svc.Destroy(test.id)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
