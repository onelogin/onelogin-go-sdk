package apprules

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
		queryPayload     *AppRuleQuery
		expectedResponse []AppRule
		expectedError    error
		repository       *test.MockRepository
	}{
		"It gets the rules for an app": {
			queryPayload: &AppRuleQuery{AppID: "1"},
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
			queryPayload:  &AppRuleQuery{AppID: "-1"},
			expectedError: errors.New("error"),
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return nil, errors.New("error")
				},
			},
		},
		"It returns an empty array if no rules exist": {
			queryPayload:     &AppRuleQuery{AppID: "1"},
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
		appID, appRuleID int32
		expectedResponse *AppRule
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets one app rule": {
			appID:     1,
			appRuleID: 1,
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
			appID:         -1,
			appRuleID:     1,
			expectedError: errors.New("error"),
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return nil, errors.New("error")
				},
			},
		},
		"it returns an empty response if app rule does not exist": {
			appID:            1,
			appRuleID:        -1,
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
			svc := New(test.repository, "test.com")
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
		"it returns an error if the app does not exist": {
			updatePayload: AppRule{
				ID:         oltypes.Int32(1),
				AppID:      oltypes.Int32(1),
				Name:       oltypes.String("update me"),
				Conditions: []AppRuleConditions{{Source: oltypes.String("test"), Operator: oltypes.String(">"), Value: oltypes.String("90")}},
				Actions:    []AppRuleActions{{Action: oltypes.String("test"), Value: []string{"test"}, Expression: oltypes.String(".*")}},
			},
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
			svc := New(test.repository, "test.com")
			actual, err := svc.Update(&test.updatePayload)
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
			svc := New(test.repository, "test.com")
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
		updatePayload    AppRule
		expectedResponse AppRule
		expectedError    error
		repository       *test.MockRepository
	}{
		"it destroys the app rule": {
			appID:     1,
			appRuleID: 1,
			repository: &test.MockRepository{
				DestroyFunc: func(r interface{}) ([]byte, error) {
					return nil, nil
				},
			},
		},
		"it returns an error if the app does not exist": {
			appID:         -1,
			appRuleID:     1,
			expectedError: errors.New("error"),
			repository: &test.MockRepository{
				DestroyFunc: func(r interface{}) ([]byte, error) {
					return nil, errors.New("error")
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			err := svc.Destroy(test.appID, test.appRuleID)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}
