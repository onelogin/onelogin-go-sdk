package smarthooks

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
		queryPayload     *SmartHookQuery
		expectedResponse []SmartHook
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets one smarthook": {
			queryPayload:     &SmartHookQuery{Limit: "1"},
			expectedResponse: []SmartHook{SmartHook{ID: oltypes.String("abc"), Function: oltypes.String("ZnVuY3Rpb24gbXlGdW5jKCkgey4uLn0=")}},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal([]SmartHook{SmartHook{ID: oltypes.String("abc"), Function: oltypes.String("ZnVuY3Rpb24gbXlGdW5jKCkgey4uLn0=")}})
					return [][]byte{b}, err
				},
			},
		},
		"it returns the remote default limit of smarthooks if no query given": {
			queryPayload: &SmartHookQuery{},
			expectedResponse: []SmartHook{
				SmartHook{ID: oltypes.String("abc"), Function: oltypes.String("ZnVuY3Rpb24gbXlGdW5jKCkgey4uLn0=")},
				SmartHook{ID: oltypes.String("abc"), Function: oltypes.String("ZnVuY3Rpb24gbXlPdGhlckZ1bmMoKSB7Li4ufQ==")},
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal([]SmartHook{
						SmartHook{ID: oltypes.String("abc"), Function: oltypes.String("ZnVuY3Rpb24gbXlGdW5jKCkgey4uLn0=")},
						SmartHook{ID: oltypes.String("abc"), Function: oltypes.String("ZnVuY3Rpb24gbXlPdGhlckZ1bmMoKSB7Li4ufQ==")},
					})
					return [][]byte{b}, err
				},
			},
		},
		"it returns an error if the call to /smarthooks fails": {
			queryPayload:     &SmartHookQuery{},
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
		expectedResponse *SmartHook
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets one smarthook by id": {
			id:               "1",
			expectedResponse: &SmartHook{ID: oltypes.String("abc"), Function: oltypes.String("ZnVuY3Rpb24gbXlGdW5jKCkgey4uLn0=")},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal(SmartHook{ID: oltypes.String("abc"), Function: oltypes.String("ZnVuY3Rpb24gbXlGdW5jKCkgey4uLn0=")})
					return [][]byte{b}, err
				},
			},
		},
		"it returns an error if the call to /smarthooks fails": {
			expectedError:    errors.New("error"),
			expectedResponse: &SmartHook{},
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
		updatePayload  *SmartHook
		expectedResult *SmartHook
		expectedError  error
		repository     *test.MockRepository
	}{
		"it updates a smarthook": {
			updatePayload:  &SmartHook{Function: oltypes.String("function myFunc() {...}"), ID: oltypes.String("abc")},
			expectedResult: &SmartHook{ID: oltypes.String("abc"), Function: oltypes.String("ZnVuY3Rpb24gbXlGdW5jKCkgey4uLn0=")},
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(SmartHook{ID: oltypes.String("abc"), Function: oltypes.String("ZnVuY3Rpb24gbXlGdW5jKCkgey4uLn0=")})
				},
			},
		},
		"it is forgiving of encoded Function": {
			updatePayload:  &SmartHook{ID: oltypes.String("abc"), Function: oltypes.String("ZnVuY3Rpb24gbXlGdW5jKCkgey4uLn0=")},
			expectedResult: &SmartHook{ID: oltypes.String("abc"), Function: oltypes.String("ZnVuY3Rpb24gbXlGdW5jKCkgey4uLn0=")},
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(SmartHook{ID: oltypes.String("abc"), Function: oltypes.String("ZnVuY3Rpb24gbXlGdW5jKCkgey4uLn0=")})
				},
			},
		},
		"it returns an error if no function given": {
			updatePayload:  &SmartHook{ID: oltypes.String("abc")},
			expectedResult: nil,
			expectedError:  errors.New("No Function Definition Given"),
			repository:     &test.MockRepository{},
		},
		"it returns an error if no id on resource": {
			updatePayload:  &SmartHook{Function: oltypes.String("function myFunc() {...}")},
			expectedResult: nil,
			expectedError:  errors.New("No ID Given"),
			repository:     &test.MockRepository{},
		},
		"it returns an error if something went wrong": {
			updatePayload:  &SmartHook{ID: oltypes.String("abc"), Function: oltypes.String("function myFunc() {...}")},
			expectedResult: nil,
			expectedError:  errors.New("error"),
			repository:     &test.MockRepository{},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			test.updatePayload.EncodeFunction()
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
		createPayload  *SmartHook
		expectedResult *SmartHook
		expectedError  error
		repository     *test.MockRepository
	}{
		"it creates a smarthook by encoding a decoded function": {
			createPayload:  &SmartHook{Function: oltypes.String("function myFunc() {...}")},
			expectedResult: &SmartHook{ID: oltypes.String("abc"), Function: oltypes.String("ZnVuY3Rpb24gbXlGdW5jKCkgey4uLn0=")},
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(SmartHook{ID: oltypes.String("abc"), Function: oltypes.String("ZnVuY3Rpb24gbXlGdW5jKCkgey4uLn0=")})
				},
			},
		},
		"it is forgiving of encoded Function": {
			createPayload:  &SmartHook{Function: oltypes.String("ZnVuY3Rpb24gbXlGdW5jKCkgey4uLn0=")},
			expectedResult: &SmartHook{ID: oltypes.String("abc"), Function: oltypes.String("ZnVuY3Rpb24gbXlGdW5jKCkgey4uLn0=")},
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(SmartHook{ID: oltypes.String("abc"), Function: oltypes.String("ZnVuY3Rpb24gbXlGdW5jKCkgey4uLn0=")})
				},
			},
		},
		"it returns an error if no encoded or decoded function given": {
			createPayload:  &SmartHook{},
			expectedResult: nil,
			expectedError:  errors.New("No Function Definition Given"),
			repository:     &test.MockRepository{},
		},
		"it returns an error if something went wrong": {
			createPayload:  &SmartHook{Function: oltypes.String("function myFunc() {...}")},
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
		expectedResponse *SmartHook
		expectedError    error
	}{
		"it destroys one smarthook": {
			id: "1",
			repository: &test.MockRepository{
				DestroyFunc: func(r interface{}) ([]byte, error) {
					return nil, nil
				},
			},
			expectedResponse: &SmartHook{},
		},
		"it returns an error if there is a problem finding the smarthook": {
			id:               "1",
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

func TestEncodeDecode(t *testing.T) {
	tests := map[string]struct {
		Subject *SmartHook
		Decoded *SmartHook
		Encoded *SmartHook
	}{
		"It Encodes and Decodes the hook function": {
			Subject: &SmartHook{ID: oltypes.String("abc"), Function: oltypes.String("function myFunc() {...}")},
			Decoded: &SmartHook{ID: oltypes.String("abc"), Function: oltypes.String("function myFunc() {...}")},
			Encoded: &SmartHook{ID: oltypes.String("abc"), Function: oltypes.String("ZnVuY3Rpb24gbXlGdW5jKCkgey4uLn0=")},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.Subject.EncodeFunction()
			assert.Equal(t, test.Encoded, test.Subject)
			test.Subject.DecodeFunction()
			assert.Equal(t, test.Decoded, test.Decoded)
		})
	}
}
