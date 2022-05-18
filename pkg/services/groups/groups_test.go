package groups

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
		queryPayload     *GroupQuery
		expectedResponse []Group
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets one group": {
			queryPayload:     &GroupQuery{Limit: "1"},
			expectedResponse: []Group{{ID: oltypes.Int32(123), Name: oltypes.String("my_group")}},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal([]Group{{ID: oltypes.Int32(123), Name: oltypes.String("my_group")}})
					return [][]byte{b}, err
				},
			},
		},
		"it returns the remote default limit of groups if no query given": {
			queryPayload: &GroupQuery{},
			expectedResponse: []Group{
				{ID: oltypes.Int32(123), Name: oltypes.String("my_group")},
				{ID: oltypes.Int32(123), Name: oltypes.String("my_group")},
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal([]Group{
						{ID: oltypes.Int32(123), Name: oltypes.String("my_group")},
						{ID: oltypes.Int32(123), Name: oltypes.String("my_group")},
					})
					return [][]byte{b}, err
				},
			},
		},
		"it returns an error if the call to /groups fails": {
			queryPayload:     &GroupQuery{},
			expectedError:    errors.New("error"),
			expectedResponse: nil,
			repository:       &test.MockRepository{},
		},
	}
	for name, testCfg := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(testCfg.repository, "test.com")
			actual, err := svc.Query(testCfg.queryPayload)
			assert.Equal(t, testCfg.expectedResponse, actual)
			if testCfg.expectedError != nil {
				assert.Equal(t, testCfg.expectedError, err)
			}
		})
	}
}

func TestGetOne(t *testing.T) {
	tests := map[string]struct {
		id               int32
		expectedResponse *Group
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets one group by id": {
			id:               123,
			expectedResponse: &Group{ID: oltypes.Int32(123), Name: oltypes.String("my_group")},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal(Group{ID: oltypes.Int32(123), Name: oltypes.String("my_group")})
					return [][]byte{b}, err
				},
			},
		},
		"it returns an error if the call to /groups fails": {
			expectedError:    errors.New("error"),
			expectedResponse: nil,
			repository:       &test.MockRepository{},
		},
	}
	for name, testCfg := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(testCfg.repository, "test.com")
			actual, err := svc.GetOne(testCfg.id)
			assert.Equal(t, testCfg.expectedResponse, actual)
			if testCfg.expectedError != nil {
				assert.Equal(t, testCfg.expectedError, err)
			}
		})
	}
}
