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
	DoFunc func(r interface{}) ([]byte, error)
}

func (mr MockRepository) Read(r interface{}) ([]byte, error) {
	req := r.(olhttp.OLHTTPRequest)
	if req.URL == "test.com/api/2/apps/1" {
		out, _ := json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("name")})
		return out, nil
	}
	if req.URL == "test.com/api/2/apps/1/rules" {
		out, _ := json.Marshal([]AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}})
		return out, nil
	}
	if mr.DoFunc != nil {
		return mr.DoFunc(r)
	}
	return nil, nil
}

func (mr MockRepository) Create(r interface{}) ([]byte, error) {
	if mr.DoFunc != nil {
		return mr.DoFunc(r)
	}
	req := r.(olhttp.OLHTTPRequest)
	if req.URL == "test.com/api/2/apps/1/rules" {
		resp := map[string]int32{"id": 1}
		out, _ := json.Marshal(resp)
		return out, nil
	}
	if req.URL == "test.com/api/2/apps" {
		app := req.Payload.(*App)
		app.ID = oltypes.Int32(int32(1))
		resp := App{Name: app.Name, ID: app.ID}
		out, _ := json.Marshal(resp)
		return out, nil
	}
	return nil, nil
}

func (mr MockRepository) Update(r interface{}) ([]byte, error) {
	if mr.DoFunc != nil {
		return mr.DoFunc(r)
	}
	req := r.(olhttp.OLHTTPRequest)
	if req.URL == "test.com/api/2/apps/1" {
		app := req.Payload.(*App)
		a := App{ID: app.ID, Name: oltypes.String("updated")}
		out, _ := json.Marshal(a)
		return out, nil
	}
	if req.URL == "test.com/api/2/apps/1/rules/1" {
		resp := map[string]int32{"id": 1}
		out, _ := json.Marshal(resp)
		return out, nil
	}
	return nil, nil
}

func (mr MockRepository) Destroy(r interface{}) ([]byte, error) {
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
		repository       MockRepository
	}{
		"it gets one app": {
			queryPayload: &AppsQuery{Limit: "1"},
			expectedResponse: []App{
				App{ID: oltypes.Int32(1), Name: oltypes.String("name"), Rules: []AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}}},
			},
			repository: MockRepository{
				DoFunc: func(r interface{}) ([]byte, error) {
					var out []byte
					out, _ = json.Marshal([]App{App{ID: oltypes.Int32(1), Name: oltypes.String("name")}})
					return out, nil
				},
			},
		},
		"it returns the remote default limit of apps if no query given": {
			queryPayload: &AppsQuery{},
			expectedResponse: []App{
				App{ID: oltypes.Int32(1), Name: oltypes.String("name"), Rules: []AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}}},
				App{ID: oltypes.Int32(1), Name: oltypes.String("name2"), Rules: []AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}}},
			},
			repository: MockRepository{
				DoFunc: func(r interface{}) ([]byte, error) {
					availableApps := []App{
						App{ID: oltypes.Int32(1), Name: oltypes.String("name")},
						App{ID: oltypes.Int32(1), Name: oltypes.String("name2")},
					}
					var out []byte
					out, _ = json.Marshal(availableApps)
					return out, nil
				},
			},
		},
		"it returns an empty response if no apps meet the criteria": {
			queryPayload:     &AppsQuery{Name: "???"},
			expectedResponse: []App{},
			repository: MockRepository{
				DoFunc: func(r interface{}) ([]byte, error) {
					var out []byte
					apps := []App{}
					out, _ = json.Marshal(apps)
					return out, nil
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
		repository       MockRepository
	}{
		"it gets one app": {
			id: int32(1),
			expectedResponse: &App{
				ID:    oltypes.Int32(1),
				Name:  oltypes.String("name"),
				Rules: []AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}},
			},
			repository: MockRepository{},
		},
		"it returns an error if there is a problem finding the app": {
			id:               int32(2),
			expectedResponse: nil,
			expectedError:    errors.New("not found"),
			repository: MockRepository{
				DoFunc: func(r interface{}) ([]byte, error) { return nil, errors.New("not found") },
			},
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
		repository       MockRepository
	}{
		"it updates one app": {
			id: int32(1),
			updatePayload: &App{
				ID:   oltypes.Int32(1),
				Name: oltypes.String("original"),
				Rules: []AppRule{
					AppRule{
						ID:   oltypes.Int32(1),
						Name: oltypes.String("updated_rule"),
					},
					AppRule{
						Name: oltypes.String("new_rule"),
					},
				},
			},
			expectedResponse: &App{
				ID:   oltypes.Int32(1),
				Name: oltypes.String("updated"),
				Rules: []AppRule{
					AppRule{
						ID:   oltypes.Int32(1),
						Name: oltypes.String("updated_rule"),
					},
					AppRule{
						ID:   oltypes.Int32(1),
						Name: oltypes.String("new_rule"),
					},
				},
			},
			repository: MockRepository{},
		},
		"it returns an error if there is a problem finding the app": {
			id:               int32(2),
			expectedResponse: &App{},
			expectedError:    errors.New("not found"),
			repository: MockRepository{
				DoFunc: func(r interface{}) ([]byte, error) { return nil, errors.New("not found") },
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
		repository       MockRepository
		expectedResponse *App
		expectedError    error
	}{
		"it destroys one app": {
			id:               int32(1),
			repository:       MockRepository{},
			expectedResponse: &App{},
		},
		"it returns an error if there is a problem finding the app": {
			id: int32(2),
			repository: MockRepository{
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
		repository       MockRepository
	}{
		"it creates one app": {
			createPayload:    &App{Name: oltypes.String("name"), Rules: []AppRule{AppRule{Name: oltypes.String("rule")}}},
			expectedResponse: &App{ID: oltypes.Int32(1), Name: oltypes.String("name"), Rules: []AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}}},
			repository:       MockRepository{},
		},
		"it returns an error if there is a bad request": {
			createPayload:    &App{ID: oltypes.Int32(1), Name: oltypes.String("not allowed value")},
			expectedResponse: &App{},
			expectedError:    errors.New("bad request"),
			repository: MockRepository{
				DoFunc: func(r interface{}) ([]byte, error) {
					return nil, errors.New("bad request")
				},
			},
		},
		"it returns the app with error if there is a bad rules request": {
			createPayload:    &App{ID: oltypes.Int32(1), Name: oltypes.String("name"), Rules: []AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("not allowed value")}}},
			expectedResponse: &App{ID: oltypes.Int32(1), Name: oltypes.String("name"), Rules: []AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("not allowed value")}}},
			expectedError:    errors.New("bad request"),
			repository: MockRepository{
				DoFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1/rules/1" {
						return nil, errors.New("bad request")
					}
					if req.URL == "test.com/api/2/apps" {
						app := req.Payload.(*App)
						app.ID = oltypes.Int32(int32(1))
						resp := App{Name: app.Name, ID: app.ID}
						out, _ := json.Marshal(resp)
						return out, nil
					}
					if req.URL == "test.com/api/2/apps/1" {
						app := req.Payload.(*App)
						app.ID = oltypes.Int32(int32(1))
						resp := App{Name: app.Name, ID: app.ID}
						out, _ := json.Marshal(resp)
						return out, nil
					}
					return nil, nil
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
