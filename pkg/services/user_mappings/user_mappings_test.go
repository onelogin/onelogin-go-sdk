package usermappings

import (
	"encoding/json"
	"errors"
	"github.com/onelogin/onelogin-go-sdk/internal/test"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockLegalValuesService struct {
	DoFunc func(addr string, o interface{}) error
}

func (svc MockLegalValuesService) Query(addr string, o interface{}) error {
	if svc.DoFunc != nil {
		return svc.DoFunc(addr, o)
	}
	return errors.New("legal val error")
}

func TestQuery(t *testing.T) {
	tests := map[string]struct {
		queryPayload     *UserMappingsQuery
		mockLegalValues  *MockLegalValuesService
		expectedResponse []UserMapping
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets one mapping": {
			queryPayload:    &UserMappingsQuery{Limit: "1"},
			mockLegalValues: &MockLegalValuesService{},
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
			svc := New(test.repository, test.mockLegalValues, "test.com")
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
			svc := New(test.repository, MockLegalValuesService{}, "test.com")
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
		mockLegalValues  *MockLegalValuesService
		expectedError    error
		repository       *test.MockRepository
	}{
		"it updates one user mapping": {
			id: int32(1),
			updatePayload: &UserMapping{
				ID:         oltypes.Int32(1),
				Name:       oltypes.String("updated"),
				Conditions: []UserMappingConditions{{Operator: oltypes.String("ri"), Source: oltypes.String("has_role"), Value: oltypes.String("12345")}},
				Actions:    []UserMappingActions{{Value: []string{"12345"}, Action: oltypes.String("set_status")}},
			},
			expectedResponse: &UserMapping{
				ID:         oltypes.Int32(1),
				Name:       oltypes.String("updated"),
				Conditions: []UserMappingConditions{{Operator: oltypes.String("ri"), Source: oltypes.String("has_role"), Value: oltypes.String("12345")}},
				Actions:    []UserMappingActions{{Value: []string{"12345"}, Action: oltypes.String("set_status")}},
			},
			mockLegalValues: &MockLegalValuesService{
				DoFunc: func(addr string, o interface{}) error {
					json.Unmarshal([]byte(`[{"value":"ri"}, {"value":"has_role"}, {"value": "12345"}, {"value": "set_status"}]`), o)
					return nil
				},
			},
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(map[string]int32{"id": 1})
				},
			},
		},
		"it updates one user mapping allowing freeform inputs when no valid values returned": {
			id: int32(1),
			updatePayload: &UserMapping{
				ID:   oltypes.Int32(1),
				Name: oltypes.String("updated"),
				Conditions: []UserMappingConditions{{
					Operator: oltypes.String("ri"),
					Source:   oltypes.String("has_role"),
					Value:    oltypes.String("12345"),
				}},
				Actions: []UserMappingActions{{
					Value:  []string{"1"},
					Action: oltypes.String("set_status"),
				}},
			},
			expectedResponse: &UserMapping{
				ID:   oltypes.Int32(1),
				Name: oltypes.String("updated"),
				Conditions: []UserMappingConditions{{
					Operator: oltypes.String("ri"),
					Source:   oltypes.String("has_role"),
					Value:    oltypes.String("12345"),
				}},
				Actions: []UserMappingActions{{
					Value:  []string{"1"},
					Action: oltypes.String("set_status"),
				}},
			},
			mockLegalValues: &MockLegalValuesService{
				DoFunc: func(addr string, o interface{}) error {
					json.Unmarshal([]byte(`[]`), o)
					return nil
				},
			},
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(map[string]int32{"id": 1})
				},
			},
		},
		"it returns an error if an invalid condition or action value given": {
			id: int32(1),
			updatePayload: &UserMapping{
				ID:   oltypes.Int32(1),
				Name: oltypes.String("updated"),
				Conditions: []UserMappingConditions{{
					Operator: oltypes.String("asdf"),
					Source:   oltypes.String("asdf"),
					Value:    oltypes.String("asdf"),
				}},
				Actions: []UserMappingActions{{
					Value:  []string{"2"},
					Action: oltypes.String("asdf"),
				}},
			},
			expectedError: errors.New("updated.actions.values must be one of [ri has_role 12345], got: 2"),
			mockLegalValues: &MockLegalValuesService{
				DoFunc: func(addr string, o interface{}) error {
					json.Unmarshal([]byte(`[{"value":"ri"}, {"value":"has_role"}, {"value": "12345"}]`), o)
					return nil
				},
			},
			repository: &test.MockRepository{},
		},
		"it returns an error if there is a problem finding the mapping": {
			id: int32(2),
			updatePayload: &UserMapping{
				ID:   oltypes.Int32(1),
				Name: oltypes.String("updated"),
				Conditions: []UserMappingConditions{{
					Operator: oltypes.String("ri"),
					Source:   oltypes.String("has_role"),
					Value:    oltypes.String("12345"),
				}},
				Actions: []UserMappingActions{{
					Value:  []string{"1"},
					Action: oltypes.String("set_status"),
				}},
			},
			expectedError: errors.New("error"),
			mockLegalValues: &MockLegalValuesService{
				DoFunc: func(addr string, o interface{}) error {
					json.Unmarshal([]byte(`[{"value":"ri"}, "value":"has_role", "value": "12345"]`), o)
					return nil
				},
			},
			repository: &test.MockRepository{},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, test.mockLegalValues, "test.com")
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
		mockLegalValues  *MockLegalValuesService
		expectedError    error
		repository       *test.MockRepository
	}{
		"it creates one mapping": {
			createPayload: &UserMapping{
				Name: oltypes.String("rule"),
				Conditions: []UserMappingConditions{{
					Operator: oltypes.String("ri"),
					Source:   oltypes.String("has_role"),
					Value:    oltypes.String("12345"),
				}},
				Actions: []UserMappingActions{{
					Value:  []string{"1"},
					Action: oltypes.String("set_status"),
				}},
			},
			expectedResponse: &UserMapping{
				ID:   oltypes.Int32(1),
				Name: oltypes.String("rule"),
				Conditions: []UserMappingConditions{{
					Operator: oltypes.String("ri"),
					Source:   oltypes.String("has_role"),
					Value:    oltypes.String("12345"),
				}},
				Actions: []UserMappingActions{{
					Value:  []string{"1"},
					Action: oltypes.String("set_status"),
				}},
			},
			mockLegalValues: &MockLegalValuesService{
				DoFunc: func(addr string, o interface{}) error {
					json.Unmarshal([]byte(`[{"value":"ri"}, "value":"has_role", "value": "12345"]`), o)
					return nil
				},
			},
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(map[string]int32{"id": 1})
				},
			},
		},
		"it returns an error if an invalid condition or action value given": {
			createPayload: &UserMapping{
				Name: oltypes.String("updated"),
				Conditions: []UserMappingConditions{{
					Operator: oltypes.String("asdf"),
					Source:   oltypes.String("asdf"),
					Value:    oltypes.String("asdf"),
				}},
				Actions: []UserMappingActions{{
					Value:  []string{"2"},
					Action: oltypes.String("asdf"),
				}},
			},
			expectedError: errors.New("updated.actions.values must be one of [ri has_role 12345], got: 2"),
			mockLegalValues: &MockLegalValuesService{
				DoFunc: func(addr string, o interface{}) error {
					json.Unmarshal([]byte(`[{"value":"ri"}, {"value":"has_role"}, {"value": "12345"}]`), o)
					return nil
				},
			},
			repository: &test.MockRepository{},
		},
		"it returns an error if there is a bad request": {
			createPayload: &UserMapping{ID: oltypes.Int32(1), Name: oltypes.String("not allowed value")},
			expectedError: errors.New("error"),
			repository:    &test.MockRepository{},
			mockLegalValues: &MockLegalValuesService{
				DoFunc: func(addr string, o interface{}) error {
					json.Unmarshal([]byte(`[{"value":"ri"}, "value":"has_role", "value": "12345"]`), o)
					return nil
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, test.mockLegalValues, "test.com")
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
		mockLegalValues  *MockLegalValuesService
		expectedResponse *UserMapping
		expectedError    error
	}{
		"it destroys one user mapping": {
			id: int32(1),
			repository: &test.MockRepository{
				DestroyFunc: func(r interface{}) ([]byte, error) {
					return nil, nil
				},
			},
			expectedResponse: &UserMapping{},
		},
		"it returns an error if there is a problem finding the user mapping": {
			id:               int32(2),
			repository:       &test.MockRepository{},
			expectedResponse: nil,
			expectedError:    errors.New("error"),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, MockLegalValuesService{}, "test.com")
			err := svc.Destroy(test.id)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestValidateMappingValues(t *testing.T) {
	tests := map[string]struct {
		inputMapping    *UserMapping
		mockLegalValues *MockLegalValuesService
		expectedError   error
	}{
		"it returns nil for valid mapping values": {
			inputMapping: &UserMapping{
				ID:         oltypes.Int32(1),
				Name:       oltypes.String("updated"),
				Conditions: []UserMappingConditions{{Operator: oltypes.String("ri"), Source: oltypes.String("has_role"), Value: oltypes.String("12345")}},
				Actions:    []UserMappingActions{{Value: []string{"12345"}, Action: oltypes.String("set_status")}},
			},
			mockLegalValues: &MockLegalValuesService{
				DoFunc: func(addr string, o interface{}) error {
					json.Unmarshal([]byte(`[{"value":"ri"}, {"value":"has_role"}, {"value": "12345"}, {"value": "set_status"}]`), o)
					return nil
				},
			},
		},
		"it updates one user mapping allowing freeform inputs when no valid values returned": {
			inputMapping: &UserMapping{
				ID:   oltypes.Int32(1),
				Name: oltypes.String("updated"),
				Conditions: []UserMappingConditions{{
					Operator: oltypes.String("ri"),
					Source:   oltypes.String("has_role"),
					Value:    oltypes.String("12345"),
				}},
				Actions: []UserMappingActions{{
					Value:  []string{"1"},
					Action: oltypes.String("set_status"),
				}},
			},
			mockLegalValues: &MockLegalValuesService{
				DoFunc: func(addr string, o interface{}) error {
					json.Unmarshal([]byte(`[]`), o)
					return nil
				},
			},
		},
		"it returns an error if an invalid condition or action value given": {
			inputMapping: &UserMapping{
				ID:   oltypes.Int32(1),
				Name: oltypes.String("updated"),
				Conditions: []UserMappingConditions{{
					Operator: oltypes.String("asdf"),
					Source:   oltypes.String("asdf"),
					Value:    oltypes.String("asdf"),
				}},
				Actions: []UserMappingActions{{
					Value:  []string{"2"},
					Action: oltypes.String("asdf"),
				}},
			},
			expectedError: errors.New("updated.actions.values must be one of [ri has_role 12345], got: 2"),
			mockLegalValues: &MockLegalValuesService{
				DoFunc: func(addr string, o interface{}) error {
					json.Unmarshal([]byte(`[{"value":"ri"}, {"value":"has_role"}, {"value": "12345"}]`), o)
					return nil
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actualErr := validateMappingValues(test.inputMapping, test.mockLegalValues)
			assert.Equal(t, test.expectedError, actualErr)
		})
	}
}
