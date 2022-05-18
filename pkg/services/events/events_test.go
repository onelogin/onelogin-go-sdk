package events

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
		queryPayload     *EventQuery
		expectedResponse []Event
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets one event": {
			queryPayload:     &EventQuery{Limit: "1"},
			expectedResponse: []Event{{ID: oltypes.Int32(123), Notes: oltypes.String("my_event")}},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal([]Event{{ID: oltypes.Int32(123), Notes: oltypes.String("my_event")}})
					return [][]byte{b}, err
				},
			},
		},
		"it returns the remote default limit of events if no query given": {
			queryPayload: &EventQuery{},
			expectedResponse: []Event{
				{ID: oltypes.Int32(123), Notes: oltypes.String("my_event")},
				{ID: oltypes.Int32(123), Notes: oltypes.String("my_event")},
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal([]Event{
						{ID: oltypes.Int32(123), Notes: oltypes.String("my_event")},
						{ID: oltypes.Int32(123), Notes: oltypes.String("my_event")},
					})
					return [][]byte{b}, err
				},
			},
		},
		"it returns an error if the call to /events fails": {
			queryPayload:     &EventQuery{},
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
		expectedResponse *Event
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets one event by id": {
			id:               123,
			expectedResponse: &Event{ID: oltypes.Int32(123), Notes: oltypes.String("my_event")},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal(Event{ID: oltypes.Int32(123), Notes: oltypes.String("my_event")})
					return [][]byte{b}, err
				},
			},
		},
		"it returns an error if the call to /events fails": {
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

func TestGetEventTypes(t *testing.T) {
	tests := map[string]struct {
		expectedResponse []EventType
		expectedError    error
		repository       *test.MockRepository
	}{
		"it returns all event types": {
			expectedResponse: []EventType{
				{ID: oltypes.Int32(123), Name: oltypes.String("my_event_type")},
				{ID: oltypes.Int32(123), Name: oltypes.String("my_event_type")},
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([][]byte, error) {
					b, err := json.Marshal([]EventType{
						{ID: oltypes.Int32(123), Name: oltypes.String("my_event_type")},
						{ID: oltypes.Int32(123), Name: oltypes.String("my_event_type")},
					})
					return [][]byte{b}, err
				},
			},
		},
		"it returns an error if the call to /events fails": {
			expectedError:    errors.New("error"),
			expectedResponse: nil,
			repository:       &test.MockRepository{},
		},
	}
	for name, testCfg := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(testCfg.repository, "test.com")
			actual, err := svc.GetTypes()
			assert.Equal(t, testCfg.expectedResponse, actual)
			if testCfg.expectedError != nil {
				assert.Equal(t, testCfg.expectedError, err)
			}
		})
	}
}
