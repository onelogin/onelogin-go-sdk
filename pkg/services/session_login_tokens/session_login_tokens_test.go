package sessionlogintokens

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"

	"github.com/stretchr/testify/assert"
)

type MockRepository struct {
	DoFunc func(r interface{}) ([]byte, error)
}

func (mr MockRepository) Read(r interface{}) ([]byte, error) {
	return mr.DoFunc(r)
}

func (mr MockRepository) Create(r interface{}) ([]byte, error) {
	return mr.DoFunc(r)
}

func (mr MockRepository) Update(r interface{}) ([]byte, error) {
	return mr.DoFunc(r)
}

func (mr MockRepository) Destroy(r interface{}) ([]byte, error) {
	return mr.DoFunc(r)
}

func TestCreate(t *testing.T) {
	tests := map[string]struct {
		createPayload    *SessionLoginTokenRequest
		expectedResponse *SessionLoginToken
		expectedError    error
	}{
		"it creates one sessionLoginToken": {
			createPayload:    &SessionLoginTokenRequest{UsernameOrEmail: oltypes.String("name")},
			expectedResponse: &SessionLoginToken{SessionToken: oltypes.String("name")},
		},
		"it returns an error if there is a bad request": {
			createPayload:    &SessionLoginTokenRequest{UsernameOrEmail: oltypes.String("???")},
			expectedResponse: nil,
			expectedError:    errors.New("bad request"),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(MockRepository{
				DoFunc: func(r interface{}) ([]byte, error) {
					if test.expectedError == nil {
						out, _ := json.Marshal(SessionLoginToken{StateToken: oltypes.String("state")})
						return out, nil
					}
					return nil, errors.New("bad request")
				},
			}, "test.com")
			actual, err := svc.Create(test.createPayload)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.NotNil(t, actual)
			}
		})
	}
}
