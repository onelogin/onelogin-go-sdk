package apps

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/internal/test"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/olhttp"
	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	tests := map[string]struct {
		queryPayload     *AppsQuery
		expectedResponse []App
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets one app": {
			queryPayload: &AppsQuery{Limit: "1"},
			expectedResponse: []App{
				App{ID: oltypes.Int32(1), Name: oltypes.String("name"), Rules: []AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}}},
			},
			repository: &test.MockRepository{
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
			repository: &test.MockRepository{
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
			repository: &test.MockRepository{
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
			repository:       &test.MockRepository{},
		},
		"it returns an empty response if no apps meet the criteria": {
			queryPayload:     &AppsQuery{Name: "???"},
			expectedResponse: []App{},
			repository: &test.MockRepository{
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
		repository       *test.MockRepository
	}{
		"it gets one app": {
			id:               int32(1),
			expectedResponse: &App{ID: oltypes.Int32(1), Name: oltypes.String("name"), Rules: []AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}}},
			repository: &test.MockRepository{
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
			repository: &test.MockRepository{
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
		id               int32
		updatePayload    *App
		expectedResponse *App
		expectedError    error
		repository       *test.MockRepository
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
					AppRule{ID: oltypes.Int32(2), Name: oltypes.String("new_rule")},
				},
			},
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(map[string]int32{"id": 2})
				},
				UpdateFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1/rules" {
						return json.Marshal(map[string]int32{"id": 1})
					}
					return json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("updated")})
				},
				ReadFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1/rules" {
						return json.Marshal([]AppRule{
							AppRule{ID: oltypes.Int32(1), Name: oltypes.String("updated_rule")},
							AppRule{ID: oltypes.Int32(2), Name: oltypes.String("new_rule")},
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
			repository:       &test.MockRepository{},
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
			repository: &test.MockRepository{
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
			repository: &test.MockRepository{
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
			repository: &test.MockRepository{
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
				DestroyFunc: func(r interface{}) ([]byte, error) {
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
		repository       *test.MockRepository
		expectedResponse *App
		expectedError    error
	}{
		"it destroys one app": {
			id: int32(1),
			repository: &test.MockRepository{
				DestroyFunc: func(r interface{}) ([]byte, error) {
					return nil, nil
				},
			},
			expectedResponse: &App{},
		},
		"it returns an error if there is a problem finding the app": {
			id:               int32(2),
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

func TestCreate(t *testing.T) {
	tests := map[string]struct {
		createPayload    *App
		expectedResponse *App
		expectedError    error
		repository       *test.MockRepository
	}{
		"it creates one app": {
			createPayload:    &App{Name: oltypes.String("name"), Rules: []AppRule{AppRule{Name: oltypes.String("rule")}}},
			expectedResponse: &App{ID: oltypes.Int32(1), Name: oltypes.String("name"), Rules: []AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}}},
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1/rules" {
						return json.Marshal(map[string]int32{"id": 1})
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
			repository:       &test.MockRepository{},
		},
		"it returns the app with error if there is a bad rules request": {
			createPayload:    &App{Name: oltypes.String("name"), Rules: []AppRule{AppRule{Name: oltypes.String("not allowed value")}}},
			expectedResponse: &App{ID: oltypes.Int32(1), Name: oltypes.String("name")},
			expectedError:    errors.New("error"),
			repository: &test.MockRepository{
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
