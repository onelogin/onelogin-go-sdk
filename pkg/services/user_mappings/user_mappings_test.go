package usermappings

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/internal/test"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/olhttp"
	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	tests := map[string]struct {
		queryPayload     *UserMappingsQuery
		expectedResponse []UserMapping
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets one mapping": {
			queryPayload: &UserMappingsQuery{Limit: "1"},
			expectedResponse: []UserMapping{
				UserMapping{ID: oltypes.Int32(1), Name: oltypes.String("mapping")},
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal([]UserMapping{UserMapping{ID: oltypes.Int32(1), Name: oltypes.String("mapping")}})
				},
			},
		},
		"it returns the remote default limit of mappings if no query given": {
			queryPayload: &UserMappingsQuery{},
			expectedResponse: []UserMapping{
				UserMapping{ID: oltypes.Int32(1), Name: oltypes.String("name")},
				UserMapping{ID: oltypes.Int32(1), Name: oltypes.String("name2")},
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal([]UserMapping{
						UserMapping{ID: oltypes.Int32(1), Name: oltypes.String("name")},
						UserMapping{ID: oltypes.Int32(1), Name: oltypes.String("name2")},
					})
				},
			},
		},
		"it returns the nothing and error if call to /mappings fails": {
			queryPayload:     &UserMappingsQuery{},
			expectedError:    errors.New("error"),
			expectedResponse: nil,
			repository:       &test.MockRepository{},
		},
		"it returns an empty response if no mappings meet the criteria": {
			queryPayload:     &UserMappingsQuery{HasAction: "???"},
			expectedResponse: []UserMapping{},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal([]UserMapping{})
				},
			},
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
		expectedResponse *UserMapping
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets one mapping": {
			id:               int32(1),
			expectedResponse: &UserMapping{ID: oltypes.Int32(1), Name: oltypes.String("name")},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(UserMapping{ID: oltypes.Int32(1), Name: oltypes.String("name")})
				},
			},
		},
		"it returns an error if there is a problem finding the mapping": {
			id:               int32(2),
			expectedResponse: nil,
			expectedError:    errors.New("error"),
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
		id               int32
		updatePayload    *UserMapping
		expectedResponse *UserMapping
		expectedError    error
		repository       *test.MockRepository
	}{
		"it updates one user mapping": {
			id:               int32(1),
			updatePayload:    &UserMapping{ID: oltypes.Int32(1), Name: oltypes.String("original")},
			expectedResponse: &UserMapping{ID: oltypes.Int32(1), Name: oltypes.String("updated")},
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(UserMapping{ID: oltypes.Int32(1), Name: oltypes.String("updated")})
				},
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(UserMapping{ID: oltypes.Int32(1), Name: oltypes.String("updated")})
				},
			},
		},
		"it returns an error if there is a problem finding the mapping": {
			id:            int32(2),
			expectedError: errors.New("error"),
			repository:    &test.MockRepository{},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			actual, err := svc.Update(test.id, test.updatePayload)
			assert.Equal(t, test.expectedResponse, actual)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	tests := map[string]struct {
		createPayload    *UserMapping
		expectedResponse *UserMapping
		expectedError    error
		repository       *test.MockRepository
	}{
		"it creates one app": {
			createPayload:    &UserMapping{Name: oltypes.String("rule")},
			expectedResponse: &UserMapping{ID: oltypes.Int32(1), Name: oltypes.String("rule")},
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					mapping := req.Payload.(*UserMapping)
					mapping.ID = oltypes.Int32(int32(1))
					resp := UserMapping{Name: mapping.Name, ID: mapping.ID}
					return json.Marshal(resp)
				},
			},
		},
		"it returns an error if there is a bad request": {
			createPayload: &UserMapping{ID: oltypes.Int32(1), Name: oltypes.String("not allowed value")},
			expectedError: errors.New("error"),
			repository:    &test.MockRepository{},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			actual, err := svc.Create(test.createPayload)
			assert.Equal(t, test.expectedResponse, actual)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestDestroy(t *testing.T) {
	tests := map[string]struct {
		id               int32
		repository       *test.MockRepository
		expectedResponse *UserMapping
		expectedError    error
	}{
		"it destroys one user mapping": {
			id:               int32(1),
			repository:       &test.MockRepository{},
			expectedResponse: &UserMapping{},
		},
		"it returns an error if there is a problem finding the user mapping": {
			id: int32(2),
			repository: &test.MockRepository{
				DoFunc: func(r interface{}) ([]byte, error) { return nil, errors.New("not found") },
			},
			expectedResponse: nil,
			expectedError:    errors.New("not found"),
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
