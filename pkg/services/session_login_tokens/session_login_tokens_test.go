package sessionlogintokens

import (
	"errors"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/internal/test"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	tests := map[string]struct {
		createPayload    *SessionLoginTokenRequest
		expectedResponse *SessionLoginToken
		repository       *test.MockRepository
		expectedError    error
	}{
		"it creates one sessionLoginToken": {
			createPayload:    &SessionLoginTokenRequest{UsernameOrEmail: oltypes.String("name")},
			expectedResponse: &SessionLoginToken{SessionToken: oltypes.String("name")},
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}, o interface{}) error {
					*o.(*SessionLoginToken) = SessionLoginToken{StateToken: oltypes.String("state")}
					return nil
				},
			},
		},
		"it returns an error if there is a bad request": {
			createPayload:    &SessionLoginTokenRequest{UsernameOrEmail: oltypes.String("???")},
			expectedResponse: nil,
			expectedError:    errors.New("bad request"),
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}, o interface{}) error {
					return errors.New("bad request")
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			actual, err := svc.Create(test.createPayload)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.NotNil(t, actual)
			}
		})
	}
}
