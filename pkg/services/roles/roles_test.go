package roles

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
		queryPayload     *RoleQuery
		expectedResponse []Role
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets one role": {
			queryPayload:     &RoleQuery{Limit: "1"},
			expectedResponse: []Role{Role{ID: oltypes.Int32(123), Name: oltypes.String("my_role")}},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal([]Role{Role{ID: oltypes.Int32(123), Name: oltypes.String("my_role")}})
					return [][]byte{b}, err
				},
			},
		},
		"it returns the remote default limit of roles if no query given": {
			queryPayload: &RoleQuery{},
			expectedResponse: []Role{
				Role{ID: oltypes.Int32(123), Name: oltypes.String("my_role")},
				Role{ID: oltypes.Int32(123), Name: oltypes.String("my_role")},
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal([]Role{
						Role{ID: oltypes.Int32(123), Name: oltypes.String("my_role")},
						Role{ID: oltypes.Int32(123), Name: oltypes.String("my_role")},
					})
					return [][]byte{b}, err
				},
			},
		},
		"it returns an error if the call to /roles fails": {
			queryPayload:     &RoleQuery{},
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
		expectedResponse *Role
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets one role by id": {
			id:               123,
			expectedResponse: &Role{ID: oltypes.Int32(123), Name: oltypes.String("my_role")},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal(Role{ID: oltypes.Int32(123), Name: oltypes.String("my_role")})
					return [][]byte{b}, err
				},
			},
		},
		"it returns an error if the call to /roles fails": {
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
		updatePayload  *Role
		expectedResult *Role
		expectedError  error
		repository     *test.MockRepository
	}{
		"it updates a role": {
			updatePayload:  &Role{Name: oltypes.String("my_role"), ID: oltypes.Int32(123), Apps: []int32{123, 456}},
			expectedResult: &Role{ID: oltypes.Int32(123), Name: oltypes.String("my_role"), Apps: []int32{123, 456}},
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(Role{ID: oltypes.Int32(123), Name: oltypes.String("my_role"), Apps: []int32{123, 456}})
				},
			},
		},
		"it returns an error if no id on resource": {
			updatePayload:  &Role{Name: oltypes.String("my_role")},
			expectedResult: &Role{},
			expectedError:  errors.New("No ID Given"),
			repository:     &test.MockRepository{},
		},
		"it returns an error if something went wrong": {
			updatePayload:  &Role{ID: oltypes.Int32(123), Name: oltypes.String("my_role"), Apps: []int32{123, 456}},
			expectedResult: &Role{},
			expectedError:  errors.New(`{"description":"Invalid id(s): 123"}`),
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					msg, _ := json.Marshal(map[string]string{"description": "Invalid id(s): 123"})
					return []byte{}, errors.New(string(msg))
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			err := svc.Update(test.updatePayload)
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
		createPayload  *Role
		expectedResult *Role
		expectedError  error
		repository     *test.MockRepository
	}{
		"it creates a role": {
			createPayload:  &Role{Name: oltypes.String("my_role")},
			expectedResult: &Role{ID: oltypes.Int32(123), Name: oltypes.String("my_role")},
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(Role{ID: oltypes.Int32(123), Name: oltypes.String("my_role")})
				},
			},
		},
		"it returns an error if something went wrong": {
			createPayload:  &Role{ID: oltypes.Int32(123), Name: oltypes.String("my_role"), Apps: []int32{123, 456}},
			expectedResult: &Role{},
			expectedError:  errors.New(`{"description":"Invalid id(s): 123"}`),
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					msg, _ := json.Marshal(map[string]string{"description": "Invalid id(s): 123"})
					return []byte{}, errors.New(string(msg))
				},
			},
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
		expectedResponse *Role
		expectedError    error
	}{
		"it destroys one role": {
			id: 123,
			repository: &test.MockRepository{
				DestroyFunc: func(r interface{}) ([]byte, error) {
					return nil, nil
				},
			},
			expectedResponse: &Role{},
		},
		"it returns an error if there is a problem finding the role": {
			id:               123,
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
