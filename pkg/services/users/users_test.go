package users

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/internal/test"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	tests := map[string]struct {
		queryPayload     *UserQuery
		expectedResponse []User
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets one user": {
			queryPayload:     &UserQuery{Limit: "1"},
			expectedResponse: []User{User{ID: oltypes.Int32(1), Firstname: oltypes.String("name")}},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal([]User{User{ID: oltypes.Int32(1), Firstname: oltypes.String("name")}})
					return [][]byte{b}, err
				},
			},
		},
		"it returns the remote default limit of users if no query given": {
			queryPayload: &UserQuery{},
			expectedResponse: []User{
				User{ID: oltypes.Int32(1), Firstname: oltypes.String("name")},
				User{ID: oltypes.Int32(1), Firstname: oltypes.String("name2")},
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal([]User{
						User{ID: oltypes.Int32(1), Firstname: oltypes.String("name")},
						User{ID: oltypes.Int32(1), Firstname: oltypes.String("name2")},
					})
					return [][]byte{b}, err
				},
			},
		},
		"it returns an error if the call to /users fails": {
			queryPayload:     &UserQuery{},
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
		expectedResponse *User
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets one user by id": {
			id:               int32(1),
			expectedResponse: &User{ID: oltypes.Int32(1), Firstname: oltypes.String("first")},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal(User{ID: oltypes.Int32(1), Firstname: oltypes.String("first")})
					return [][]byte{b}, err
				},
			},
		},
		"it returns an error if the call to /users fails": {
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

func TestGetApps(t *testing.T) {
	tests := map[string]struct {
		id               int32
		expectedResponse []UserApp
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets apps for a user": {
			id: 1,
			expectedResponse: []UserApp{{
				ID:                  oltypes.Int32(1),
				IconURL:             oltypes.String("https://example.com/32x32.png"),
				LoginID:             oltypes.Int32(2),
				ProvisioningStatus:  oltypes.String("enabled"),
				ProvisioningState:   oltypes.String("modifying"),
				ProvisioningEnabled: oltypes.Bool(true),
			}},
			expectedError: nil,
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal([]UserApp{{
						ID:                  oltypes.Int32(1),
						IconURL:             oltypes.String("https://example.com/32x32.png"),
						LoginID:             oltypes.Int32(2),
						ProvisioningStatus:  oltypes.String("enabled"),
						ProvisioningState:   oltypes.String("modifying"),
						ProvisioningEnabled: oltypes.Bool(true),
					}})
					return [][]byte{b}, err
				},
			},
		},
		"it returns an error if there is a problem finding the apps": {
			id:               int32(2),
			expectedResponse: nil,
			expectedError:    errors.New("error"),
			repository:       &test.MockRepository{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			actual, err := svc.GetApps(test.id)
			assert.Equal(t, test.expectedResponse, actual)
			assert.Equal(t, test.expectedError, err)
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := map[string]struct {
		updatePayload  *User
		expectedResult *User
		expectedError  error
		repository     *test.MockRepository
	}{
		"it updates a user": {
			updatePayload:  &User{Firstname: oltypes.String("update"), ID: oltypes.Int32(int32(1))},
			expectedResult: &User{ID: oltypes.Int32(1), Firstname: oltypes.String("update")},
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(User{ID: oltypes.Int32(1), Firstname: oltypes.String("update")})
				},
			},
		},
		"it returns an error if no id on resource": {
			updatePayload:  &User{Firstname: oltypes.String("update")},
			expectedResult: &User{},
			expectedError:  errors.New("No ID Given"),
			repository:     &test.MockRepository{},
		},
		"it returns an error if something went wrong": {
			updatePayload:  &User{ID: oltypes.Int32(1), Firstname: oltypes.String("update")},
			expectedResult: &User{},
			expectedError:  errors.New("error"),
			repository:     &test.MockRepository{},
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
		createPayload  *User
		expectedResult *User
		expectedError  error
		repository     *test.MockRepository
	}{
		"it creates a user": {
			createPayload:  &User{Firstname: oltypes.String("create")},
			expectedResult: &User{ID: oltypes.Int32(1), Firstname: oltypes.String("create")},
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(User{ID: oltypes.Int32(1), Firstname: oltypes.String("create")})
				},
			},
		},
		"it returns an error if something went wrong": {
			createPayload:  &User{Firstname: oltypes.String("create")},
			expectedResult: &User{},
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
		expectedResponse *User
		expectedError    error
	}{
		"it destroys one user": {
			id: int32(1),
			repository: &test.MockRepository{
				DestroyFunc: func(r interface{}) ([]byte, error) {
					return nil, nil
				},
			},
			expectedResponse: &User{},
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
