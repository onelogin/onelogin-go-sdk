package apps

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/olhttp"
	"github.com/stretchr/testify/assert"
)

type MockRepository struct {
	DoFunc     func(r interface{}) ([]byte, error)
	ReadFunc   func(r interface{}) ([]byte, error)
	ReReadFunc func(r interface{}) ([]byte, error)
	CreateFunc func(r interface{}) ([]byte, error)
	UpdateFunc func(r interface{}) ([]byte, error)
	DeleteFunc func(r interface{}) ([]byte, error)
	ReRead     bool
}

func (mr *MockRepository) Read(r interface{}) ([]byte, error) {
	if mr.ReRead && mr.ReReadFunc != nil {
		return mr.ReReadFunc(r)
	}

	if mr.ReadFunc != nil {
		mr.ReRead = true
		return mr.ReadFunc(r)
	}
	return nil, errors.New("error")
}

func (mr *MockRepository) Create(r interface{}) ([]byte, error) {
	if mr.CreateFunc != nil {
		return mr.CreateFunc(r)
	}
	return nil, errors.New("error")
}

func (mr *MockRepository) Update(r interface{}) ([]byte, error) {
	if mr.UpdateFunc != nil {
		return mr.UpdateFunc(r)
	}
	return nil, errors.New("error")
}

func (mr *MockRepository) Destroy(r interface{}) ([]byte, error) {
	if mr.DoFunc != nil {
		return mr.DoFunc(r)
	}
	return nil, nil
}

func TestQuery(t *testing.T) {
	tests := map[string]struct {
		queryPayload     *AppsQuery
		expectedResponse []App
		expectedError    error
		repository       *MockRepository
	}{
		"it gets one app": {
			queryPayload: &AppsQuery{Limit: "1"},
			expectedResponse: []App{
				App{ID: oltypes.Int32(1), Name: oltypes.String("name"), Rules: []AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}}},
			},
			repository: &MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1/rules" {
						return json.Marshal([]AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}})
					}
					return json.Marshal([]App{App{ID: oltypes.Int32(1), Name: oltypes.String("name")}})
				},
			},
		},
		"it returns the remote default limit of apps if no query given": {
			queryPayload: &AppsQuery{},
			expectedResponse: []App{
				App{ID: oltypes.Int32(1), Name: oltypes.String("name"), Rules: []AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}}},
				App{ID: oltypes.Int32(1), Name: oltypes.String("name2"), Rules: []AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}}},
			},
			repository: &MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1/rules" {
						return json.Marshal([]AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}})
					}
					availableApps := []App{
						App{ID: oltypes.Int32(1), Name: oltypes.String("name")},
						App{ID: oltypes.Int32(1), Name: oltypes.String("name2")},
					}
					return json.Marshal(availableApps)
				},
			},
		},
		"it returns the apps and error if call to /rules fails": {
			queryPayload:  &AppsQuery{},
			expectedError: errors.New("error"),
			expectedResponse: []App{
				App{ID: oltypes.Int32(1), Name: oltypes.String("name")},
				App{ID: oltypes.Int32(1), Name: oltypes.String("name2")},
			},
			repository: &MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1/rules" {
						return nil, errors.New("error")
					}
					availableApps := []App{
						App{ID: oltypes.Int32(1), Name: oltypes.String("name")},
						App{ID: oltypes.Int32(1), Name: oltypes.String("name2")},
					}
					return json.Marshal(availableApps)
				},
			},
		},
		"it returns the nothing and error if call to /apps fails": {
			queryPayload:     &AppsQuery{},
			expectedError:    errors.New("error"),
			expectedResponse: nil,
			repository:       &MockRepository{},
		},
		"it returns an empty response if no apps meet the criteria": {
			queryPayload:     &AppsQuery{Name: "???"},
			expectedResponse: []App{},
			repository: &MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal([]App{})
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
		id               int32
		expectedResponse *App
		expectedError    error
		repository       *MockRepository
	}{
		"it gets one app": {
			id:               int32(1),
			expectedResponse: &App{ID: oltypes.Int32(1), Name: oltypes.String("name"), Rules: []AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}}},
			repository: &MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1/rules" {
						return json.Marshal([]AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}})
					}
					return json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("name")})
				},
			},
		},
		"it returns one app and error if call to /rules fails": {
			expectedError:    errors.New("error"),
			expectedResponse: &App{ID: oltypes.Int32(1), Name: oltypes.String("name")},
			repository: &MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1/rules" {
						return nil, errors.New("error")
					}
					return json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("name")})
				},
			},
		},
		"it returns an error if there is a problem finding the app": {
			id:               int32(2),
			expectedResponse: nil,
			expectedError:    errors.New("error"),
			repository:       &MockRepository{},
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
		id               int32
		updatePayload    *App
		expectedResponse *App
		expectedError    error
		repository       *MockRepository
	}{
		"it updates one app": {
			id: int32(1),
			updatePayload: &App{
				ID:   oltypes.Int32(1),
				Name: oltypes.String("original"),
				Rules: []AppRule{
					AppRule{ID: oltypes.Int32(1), Name: oltypes.String("updated_rule")},
					AppRule{Name: oltypes.String("new_rule")},
				},
			},
			expectedResponse: &App{
				ID:   oltypes.Int32(1),
				Name: oltypes.String("updated"),
				Rules: []AppRule{
					AppRule{ID: oltypes.Int32(1), Name: oltypes.String("updated_rule")},
					AppRule{ID: oltypes.Int32(1), Name: oltypes.String("new_rule")},
				},
			},
			repository: &MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(map[string]int32{"id": 1})
				},
				UpdateFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1/rules" {
						return json.Marshal(AppRule{ID: oltypes.Int32(1), Name: oltypes.String("updated_rule")})
					}
					return json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("updated")})
				},
				ReadFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1/rules" {
						return json.Marshal([]AppRule{
							AppRule{ID: oltypes.Int32(1), Name: oltypes.String("updated_rule")},
							AppRule{ID: oltypes.Int32(1), Name: oltypes.String("new_rule")},
						})
					}
					return json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("updated")})
				},
			},
		},
		"it returns an error if there is a problem finding the app": {
			id:               int32(2),
			expectedResponse: &App{},
			expectedError:    errors.New("error"),
			repository:       &MockRepository{},
		},
		"it returns the app and error if call to /rules fails": {
			updatePayload: &App{
				ID:   oltypes.Int32(1),
				Name: oltypes.String("original"),
				Rules: []AppRule{
					AppRule{ID: oltypes.Int32(1), Name: oltypes.String("updated_rule")},
					AppRule{Name: oltypes.String("new_rule")},
				},
			},
			expectedError:    errors.New("error"),
			expectedResponse: &App{ID: oltypes.Int32(1), Name: oltypes.String("name")},
			repository: &MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					return nil, errors.New("error")
				},
				UpdateFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1/rules" {
						return nil, errors.New("error")
					}
					return json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("name")})
				},
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("name")})
				},
			},
		},
		"it returns the nil and error if call to /rules fails and unable to recover": {
			updatePayload: &App{
				ID:   oltypes.Int32(1),
				Name: oltypes.String("original"),
				Rules: []AppRule{
					AppRule{ID: oltypes.Int32(1), Name: oltypes.String("updated_rule")},
					AppRule{Name: oltypes.String("new_rule")},
				},
			},
			expectedError:    errors.New("error"),
			expectedResponse: nil,
			repository: &MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					return nil, errors.New("error")
				},
				UpdateFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1/rules" {
						return nil, errors.New("error")
					}
					return json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("name")})
				},
			},
		},
		"it removes parameters and rules when requested": {
			updatePayload: &App{
				ID:         oltypes.Int32(1),
				Name:       oltypes.String("original"),
				Parameters: map[string]AppParameters{},
				Rules:      []AppRule{},
			},
			expectedResponse: &App{ID: oltypes.Int32(1), Name: oltypes.String("name"), Rules: []AppRule{}},
			repository: &MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("name"), Parameters: map[string]AppParameters{"test": AppParameters{ID: oltypes.Int32(1)}}})
				},
				ReadFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1/rules" {
						return json.Marshal([]AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}})
					}
					return json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("name")})
				},
				ReReadFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1/rules" {
						return json.Marshal([]AppRule{})
					}
					return json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("name")})
				},
				DeleteFunc: func(r interface{}) ([]byte, error) {
					return nil, nil
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			actual, err := svc.Update(test.id, test.updatePayload)
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
		repository       *MockRepository
		expectedResponse *App
		expectedError    error
	}{
		"it destroys one app": {
			id:               int32(1),
			repository:       &MockRepository{},
			expectedResponse: &App{},
		},
		"it returns an error if there is a problem finding the app": {
			id: int32(2),
			repository: &MockRepository{
				DoFunc: func(r interface{}) ([]byte, error) { return nil, errors.New("not found") },
			},
			expectedResponse: nil,
			expectedError:    errors.New("not found"),
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

func TestCreate(t *testing.T) {
	tests := map[string]struct {
		createPayload    *App
		expectedResponse *App
		expectedError    error
		repository       *MockRepository
	}{
		"it creates one app": {
			createPayload:    &App{Name: oltypes.String("name"), Rules: []AppRule{AppRule{Name: oltypes.String("rule")}}},
			expectedResponse: &App{ID: oltypes.Int32(1), Name: oltypes.String("name"), Rules: []AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}}},
			repository: &MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1/rules" {
						resp := map[string]int32{"id": 1}
						out, _ := json.Marshal(resp)
						return out, nil
					}
					app := req.Payload.(*App)
					app.ID = oltypes.Int32(int32(1))
					resp := App{Name: app.Name, ID: app.ID}
					return json.Marshal(resp)
				},
				ReadFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1/rules" {
						return json.Marshal([]AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}})
					}
					return json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("name")})
				},
			},
		},
		"it returns an error if there is a bad request": {
			createPayload:    &App{ID: oltypes.Int32(1), Name: oltypes.String("not allowed value")},
			expectedResponse: &App{},
			expectedError:    errors.New("error"),
			repository:       &MockRepository{},
		},
		"it returns the app with error if there is a bad rules request": {
			createPayload:    &App{Name: oltypes.String("name"), Rules: []AppRule{AppRule{Name: oltypes.String("not allowed value")}}},
			expectedResponse: &App{ID: oltypes.Int32(1), Name: oltypes.String("name")},
			expectedError:    errors.New("error"),
			repository: &MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1/rules" {
						return nil, errors.New("error")
					}
					return json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("name")})
				},
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("name")})
				},
			},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			actual, err := svc.Create(test.createPayload)
			assert.Equal(t, test.expectedResponse, actual)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}
