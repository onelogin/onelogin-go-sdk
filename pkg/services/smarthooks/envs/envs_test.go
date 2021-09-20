package smarthookenvs

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/internal/test"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	tests := map[string]struct {
		queryPayload     *SmartHookEnvVarQuery
		expectedResponse []EnvVar
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets one environment variable": {
			queryPayload:     &SmartHookEnvVarQuery{Limit: "abc"},
			expectedResponse: []EnvVar{EnvVar{ID: oltypes.String("abc"), Name: oltypes.String("API_KEY"), Value: oltypes.String("Its a secret to everybody")}},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal([]EnvVar{EnvVar{ID: oltypes.String("abc"), Name: oltypes.String("API_KEY"), Value: oltypes.String("Its a secret to everybody")}})
					return [][]byte{b}, err
				},
			},
		},
		"it returns the remote default limit of smarthooks if no query given": {
			queryPayload: &SmartHookEnvVarQuery{},
			expectedResponse: []EnvVar{
				EnvVar{ID: oltypes.String("abc"), Name: oltypes.String("API_KEY"), Value: oltypes.String("Its a secret to everybody")},
				EnvVar{ID: oltypes.String("abc"), Name: oltypes.String("API_KEY"), Value: oltypes.String("Its a secret to everybody")},
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal([]EnvVar{
						EnvVar{ID: oltypes.String("abc"), Name: oltypes.String("API_KEY"), Value: oltypes.String("Its a secret to everybody")},
						EnvVar{ID: oltypes.String("abc"), Name: oltypes.String("API_KEY"), Value: oltypes.String("Its a secret to everybody")},
					})
					return [][]byte{b}, err
				},
			},
		},
		"it returns an error if the call to /smarthooks fails": {
			queryPayload:     &SmartHookEnvVarQuery{},
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
		id               string
		expectedResponse *EnvVar
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets one smarthook by id": {
			id:               "abc",
			expectedResponse: &EnvVar{ID: oltypes.String("abc"), Name: oltypes.String("API_KEY"), Value: oltypes.String("Its a secret to everybody")},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal(EnvVar{ID: oltypes.String("abc"), Name: oltypes.String("API_KEY"), Value: oltypes.String("Its a secret to everybody")})
					return [][]byte{b}, err
				},
			},
		},
		"it returns an error if the call to /smarthooks fails": {
			expectedError:    errors.New("error"),
			expectedResponse: &EnvVar{},
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
		updatePayload  *EnvVar
		expectedResult *EnvVar
		expectedError  error
		repository     *test.MockRepository
	}{
		"it updates a smarthook": {
			updatePayload:  &EnvVar{ID: oltypes.String("abc"), Value: oltypes.String("SHHHHH!")},
			expectedResult: &EnvVar{ID: oltypes.String("abc"), Name: oltypes.String("API_KEY")},
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(EnvVar{ID: oltypes.String("abc"), Name: oltypes.String("API_KEY")})
				},
			},
		},
		"it attempts a search if name is given": {
			updatePayload:  &EnvVar{Name: oltypes.String("API_KEY"), Value: oltypes.String("secret_to_everybody")},
			expectedResult: &EnvVar{ID: oltypes.String("def"), Name: oltypes.String("API_KEY")},
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					fmt.Println("UP")
					return json.Marshal(EnvVar{ID: oltypes.String("def"), Name: oltypes.String("API_KEY")})
				},
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal([]EnvVar{
						EnvVar{ID: oltypes.String("abc"), Name: oltypes.String("SECRET")},
						EnvVar{ID: oltypes.String("def"), Name: oltypes.String("API_KEY")},
					})
					return [][]byte{b}, err
				},
			},
		},
		"it returns error if name search fails": {
			updatePayload:  &EnvVar{Name: oltypes.String("API_KEY"), Value: oltypes.String("secret_to_everybody")},
			expectedResult: nil,
			expectedError:  errors.New("no ID or Name given"),
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					fmt.Println("UP")
					return json.Marshal(EnvVar{ID: oltypes.String("def"), Name: oltypes.String("API_KEY")})
				},
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal([]EnvVar{})
					return [][]byte{b}, err
				},
			},
		},
		"it returns an error if no value given": {
			updatePayload:  &EnvVar{ID: oltypes.String("abc")},
			expectedResult: nil,
			expectedError:  errors.New("value is required"),
			repository:     &test.MockRepository{},
		},
		"it returns an error if no id on resource": {
			updatePayload:  &EnvVar{Value: oltypes.String("SHHHHH!")},
			expectedResult: nil,
			expectedError:  errors.New("no ID or Name given"),
			repository:     &test.MockRepository{},
		},
		"it returns an error if something went wrong": {
			updatePayload:  &EnvVar{ID: oltypes.String("abc"), Value: oltypes.String("SHHHHH!")},
			expectedResult: nil,
			expectedError:  errors.New("error"),
			repository:     &test.MockRepository{},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			out, err := svc.Update(test.updatePayload)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Equal(t, test.expectedResult, out)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	tests := map[string]struct {
		createPayload  *EnvVar
		expectedResult *EnvVar
		expectedError  error
		repository     *test.MockRepository
	}{
		"it creates an environment variable": {
			createPayload:  &EnvVar{ID: oltypes.String("abc"), Name: oltypes.String("API_KEY"), Value: oltypes.String("Its a secret to everybody")},
			expectedResult: &EnvVar{ID: oltypes.String("abc"), Name: oltypes.String("API_KEY"), Value: oltypes.String("Its a secret to everybody")},
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(EnvVar{ID: oltypes.String("abc"), Name: oltypes.String("API_KEY"), Value: oltypes.String("Its a secret to everybody")})
				},
			},
		},
		"it returns an error if no name or value given": {
			createPayload:  &EnvVar{},
			expectedResult: nil,
			expectedError:  errors.New("Name and Value are both required"),
			repository:     &test.MockRepository{},
		},
		"it returns an error if something went wrong": {
			createPayload:  &EnvVar{ID: oltypes.String("abc"), Name: oltypes.String("API_KEY"), Value: oltypes.String("Its a secret to everybody")},
			expectedResult: nil,
			expectedError:  errors.New("error"),
			repository:     &test.MockRepository{},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			out, err := svc.Create(test.createPayload)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Equal(t, test.expectedResult, out)
			}
		})
	}
}

func TestDestroy(t *testing.T) {
	tests := map[string]struct {
		id               string
		repository       *test.MockRepository
		expectedResponse *EnvVar
		expectedError    error
	}{
		"it destroys one smarthook": {
			id: "abc",
			repository: &test.MockRepository{
				DestroyFunc: func(r interface{}) ([]byte, error) {
					return nil, nil
				},
			},
			expectedResponse: &EnvVar{},
		},
		"it returns an error if there is a problem finding the smarthook": {
			id:               "abc",
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
