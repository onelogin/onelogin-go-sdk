package apprules

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
		queryPayload     *AppRuleQuery
		mockLegalValues  *MockLegalValuesService
		expectedResponse []AppRule
		expectedError    error
		repository       *test.MockRepository
	}{
		"It gets the rules for an app": {
			queryPayload:    &AppRuleQuery{AppID: "1"},
			mockLegalValues: &MockLegalValuesService{},
			expectedResponse: []AppRule{AppRule{
				ID:         oltypes.Int32(1),
				Name:       oltypes.String("name"),
				Conditions: []AppRuleConditions{{Source: oltypes.String("test"), Operator: oltypes.String(">"), Value: oltypes.String("90")}},
				Actions:    []AppRuleActions{{Action: oltypes.String("test"), Value: []string{"test"}, Expression: oltypes.String(".*")}}},
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal([]AppRule{AppRule{
						ID:         oltypes.Int32(1),
						Name:       oltypes.String("name"),
						Conditions: []AppRuleConditions{{Source: oltypes.String("test"), Operator: oltypes.String(">"), Value: oltypes.String("90")}},
						Actions:    []AppRuleActions{{Action: oltypes.String("test"), Value: []string{"test"}, Expression: oltypes.String(".*")}}},
					})
				},
			},
		},
		"It returns an error if app does not exist": {
			queryPayload:    &AppRuleQuery{AppID: "-1"},
			mockLegalValues: &MockLegalValuesService{},
			expectedError:   errors.New("error"),
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return nil, errors.New("error")
				},
			},
		},
		"It returns an empty array if no rules exist": {
			queryPayload:     &AppRuleQuery{AppID: "1"},
			mockLegalValues:  &MockLegalValuesService{},
			expectedResponse: []AppRule{AppRule{}},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal([]AppRule{AppRule{}})
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
		appID, appRuleID int32
		expectedResponse *AppRule
		mockLegalValues  *MockLegalValuesService
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets one app rule": {
			appID:           1,
			appRuleID:       1,
			mockLegalValues: &MockLegalValuesService{},
			expectedResponse: &AppRule{
				ID:         oltypes.Int32(1),
				Name:       oltypes.String("name"),
				Conditions: []AppRuleConditions{{Source: oltypes.String("test"), Operator: oltypes.String(">"), Value: oltypes.String("90")}},
				Actions:    []AppRuleActions{{Action: oltypes.String("test"), Value: []string{"test"}, Expression: oltypes.String(".*")}},
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(AppRule{
						ID:         oltypes.Int32(1),
						Name:       oltypes.String("name"),
						Conditions: []AppRuleConditions{{Source: oltypes.String("test"), Operator: oltypes.String(">"), Value: oltypes.String("90")}},
						Actions:    []AppRuleActions{{Action: oltypes.String("test"), Value: []string{"test"}, Expression: oltypes.String(".*")}},
					})
				},
			},
		},
		"it returns an error if the app does not exist": {
			appID:           -1,
			appRuleID:       1,
			mockLegalValues: &MockLegalValuesService{},
			expectedError:   errors.New("error"),
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return nil, errors.New("error")
				},
			},
		},
		"it returns an empty response if app rule does not exist": {
			appID:            1,
			appRuleID:        -1,
			mockLegalValues:  &MockLegalValuesService{},
			expectedResponse: &AppRule{},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal([]AppRule{AppRule{}})
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, test.mockLegalValues, "test.com")
			actual, err := svc.GetOne(test.appID, test.appRuleID)
			assert.Equal(t, test.expectedResponse, actual)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := map[string]struct {
		appRuleID        int32
		mockLegalValues  *MockLegalValuesService
		updatePayload    AppRule
		expectedResponse *AppRule
		expectedError    error
		repository       *test.MockRepository
	}{
		"it updates the app rule": {
			updatePayload: AppRule{
				ID:         oltypes.Int32(1),
				AppID:      oltypes.Int32(1),
				Name:       oltypes.String("update me"),
				Conditions: []AppRuleConditions{{Source: oltypes.String("ri"), Operator: oltypes.String("set_status"), Value: oltypes.String("12345")}},
				Actions:    []AppRuleActions{{Action: oltypes.String("set_status"), Value: []string{"12345"}, Expression: oltypes.String(".*")}},
			},
			expectedResponse: &AppRule{
				ID:         oltypes.Int32(1),
				AppID:      oltypes.Int32(1),
				Name:       oltypes.String("update me"),
				Conditions: []AppRuleConditions{{Source: oltypes.String("ri"), Operator: oltypes.String("set_status"), Value: oltypes.String("12345")}},
				Actions:    []AppRuleActions{{Action: oltypes.String("ri"), Value: []string{"12345"}, Expression: oltypes.String(".*")}},
			},
			mockLegalValues: &MockLegalValuesService{
				DoFunc: func(addr string, o interface{}) error {
					json.Unmarshal([]byte(`[{"value":"ri"}, {"value":"has_role"}, {"value": "12345"}, {"value": "set_status"}]`), o)
					return nil
				},
			},
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(AppRule{
						ID:         oltypes.Int32(1),
						AppID:      oltypes.Int32(1),
						Name:       oltypes.String("update me"),
						Conditions: []AppRuleConditions{{Source: oltypes.String("ri"), Operator: oltypes.String("set_status"), Value: oltypes.String("12345")}},
						Actions:    []AppRuleActions{{Action: oltypes.String("ri"), Value: []string{"12345"}, Expression: oltypes.String(".*")}},
					})
				},
			},
		},
		"it updates one app rule allowing freeform inputs when no valid values returned": {
			updatePayload: AppRule{
				ID:         oltypes.Int32(1),
				AppID:      oltypes.Int32(1),
				Name:       oltypes.String("update me"),
				Conditions: []AppRuleConditions{{Source: oltypes.String("test"), Operator: oltypes.String(">"), Value: oltypes.String("90")}},
				Actions:    []AppRuleActions{{Action: oltypes.String("test"), Value: []string{"test"}, Expression: oltypes.String(".*")}},
			},
			expectedResponse: &AppRule{
				ID:         oltypes.Int32(1),
				AppID:      oltypes.Int32(1),
				Name:       oltypes.String("update me"),
				Conditions: []AppRuleConditions{{Source: oltypes.String("test"), Operator: oltypes.String(">"), Value: oltypes.String("90")}},
				Actions:    []AppRuleActions{{Action: oltypes.String("test"), Value: []string{"test"}, Expression: oltypes.String(".*")}},
			},
			mockLegalValues: &MockLegalValuesService{
				DoFunc: func(addr string, o interface{}) error {
					json.Unmarshal([]byte(`[]`), o)
					return nil
				},
			},
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(AppRule{
						ID:         oltypes.Int32(1),
						AppID:      oltypes.Int32(1),
						Name:       oltypes.String("update me"),
						Conditions: []AppRuleConditions{{Source: oltypes.String("test"), Operator: oltypes.String(">"), Value: oltypes.String("90")}},
						Actions:    []AppRuleActions{{Action: oltypes.String("test"), Value: []string{"test"}, Expression: oltypes.String(".*")}},
					})
				},
			},
		},
		"it returns an error if an invalid condition or action value given": {
			updatePayload: AppRule{
				ID:         oltypes.Int32(1),
				AppID:      oltypes.Int32(1),
				Name:       oltypes.String("update me"),
				Conditions: []AppRuleConditions{{Source: oltypes.String("test"), Operator: oltypes.String(">"), Value: oltypes.String("90")}},
				Actions:    []AppRuleActions{{Action: oltypes.String("test"), Value: []string{"test"}, Expression: oltypes.String(".*")}},
			},
			expectedError: errors.New("update me.conditions.source must be one of [ri has_role 12345], got: test, update me.conditions.value must be one of [ri has_role 12345], got: 90, update me.conditions.operator must be one of [ri has_role 12345], got: >, update me.actions.action must be one of [ri has_role 12345], got: test, update me.actions.values must be one of [ri has_role 12345], got: test"),
			mockLegalValues: &MockLegalValuesService{
				DoFunc: func(addr string, o interface{}) error {
					json.Unmarshal([]byte(`[{"value":"ri"}, {"value":"has_role"}, {"value": "12345"}]`), o)
					return nil
				},
			},
			repository: &test.MockRepository{},
		},
		"it returns an error if the app does not exist": {
			updatePayload: AppRule{
				ID:         oltypes.Int32(1),
				AppID:      oltypes.Int32(1),
				Name:       oltypes.String("update me"),
				Conditions: []AppRuleConditions{{Source: oltypes.String("test"), Operator: oltypes.String(">"), Value: oltypes.String("90")}},
				Actions:    []AppRuleActions{{Action: oltypes.String("test"), Value: []string{"test"}, Expression: oltypes.String(".*")}},
			},
			mockLegalValues:  &MockLegalValuesService{},
			expectedError:    errors.New("error"),
			expectedResponse: &AppRule{},
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return nil, errors.New("error")
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, test.mockLegalValues, "test.com")
			actual, err := svc.Update(1, &test.updatePayload)
			assert.Equal(t, test.expectedResponse, actual)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	tests := map[string]struct {
		appID            int32
		mockLegalValues  *MockLegalValuesService
		createPayload    AppRule
		expectedResponse *AppRule
		expectedError    error
		repository       *test.MockRepository
	}{

		"it creates the app rule": {
			createPayload: AppRule{
				AppID:      oltypes.Int32(-1),
				Name:       oltypes.String("new app rule"),
				Conditions: []AppRuleConditions{{Source: oltypes.String("test"), Operator: oltypes.String(">"), Value: oltypes.String("90")}},
				Actions:    []AppRuleActions{{Action: oltypes.String("test"), Value: []string{"test"}, Expression: oltypes.String(".*")}},
			},
			expectedResponse: &AppRule{
				ID:         oltypes.Int32(1),
				AppID:      oltypes.Int32(1),
				Name:       oltypes.String("new app rule"),
				Conditions: []AppRuleConditions{{Source: oltypes.String("test"), Operator: oltypes.String(">"), Value: oltypes.String("90")}},
				Actions:    []AppRuleActions{{Action: oltypes.String("test"), Value: []string{"test"}, Expression: oltypes.String(".*")}},
			},
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(AppRule{
						ID:         oltypes.Int32(1),
						AppID:      oltypes.Int32(1),
						Name:       oltypes.String("new app rule"),
						Conditions: []AppRuleConditions{{Source: oltypes.String("test"), Operator: oltypes.String(">"), Value: oltypes.String("90")}},
						Actions:    []AppRuleActions{{Action: oltypes.String("test"), Value: []string{"test"}, Expression: oltypes.String(".*")}},
					})
				},
			},
		},
		"it returns an error if the app does not exist": {
			createPayload: AppRule{
				ID:         oltypes.Int32(1),
				AppID:      oltypes.Int32(-1),
				Name:       oltypes.String("new app rule"),
				Conditions: []AppRuleConditions{{Source: oltypes.String("test"), Operator: oltypes.String(">"), Value: oltypes.String("90")}},
				Actions:    []AppRuleActions{{Action: oltypes.String("test"), Value: []string{"test"}, Expression: oltypes.String(".*")}},
			},
			expectedError:    errors.New("error"),
			expectedResponse: &AppRule{},
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					return nil, errors.New("error")
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, test.mockLegalValues, "test.com")
			actual, err := svc.Create(&test.createPayload)
			assert.Equal(t, test.expectedResponse, actual)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestDestroy(t *testing.T) {
	tests := map[string]struct {
		appID, appRuleID int32
		mockLegalValues  *MockLegalValuesService
		updatePayload    AppRule
		expectedResponse AppRule
		expectedError    error
		repository       *test.MockRepository
	}{
		"it destroys the app rule": {
			appID:           1,
			appRuleID:       1,
			mockLegalValues: &MockLegalValuesService{},
			repository: &test.MockRepository{
				DestroyFunc: func(r interface{}) ([]byte, error) {
					return nil, nil
				},
			},
		},
		"it returns an error if the app does not exist": {
			appID:           -1,
			appRuleID:       1,
			mockLegalValues: &MockLegalValuesService{},
			expectedError:   errors.New("error"),
			repository: &test.MockRepository{
				DestroyFunc: func(r interface{}) ([]byte, error) {
					return nil, errors.New("error")
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, test.mockLegalValues, "test.com")
			err := svc.Destroy(test.appID, test.appRuleID)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}
