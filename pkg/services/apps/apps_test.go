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

func TestQuery(t *testing.T) {
	tests := map[string]struct {
		queryPayload     *AppsQuery
		expectedResponse []App
		expectedError    error
	}{
		"it gets one app": {
			queryPayload: &AppsQuery{Limit: "1"},
			expectedResponse: []App{
				App{ID: oltypes.Int32(1), Name: oltypes.String("name"), Rules: []AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}}},
			},
		},
		"it returns the remote default limit of apps if no query given": {
			queryPayload: &AppsQuery{},
			expectedResponse: []App{
				App{ID: oltypes.Int32(1), Name: oltypes.String("name"), Rules: []AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}}},
				App{ID: oltypes.Int32(2), Name: oltypes.String("name2"), Rules: []AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}}},
			},
		},
		"it returns an empty response if no apps meet the criteria": {
			queryPayload:     &AppsQuery{Name: "???"},
			expectedResponse: []App{},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(MockRepository{
				DoFunc: func(r interface{}) ([]byte, error) {
					availableApps := map[string]App{
						"name":  App{ID: oltypes.Int32(1), Name: oltypes.String("name")},
						"name2": App{ID: oltypes.Int32(2), Name: oltypes.String("name2")},
					}
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1/rules" || req.URL == "test.com/api/2/apps/2/rules" {
						out, _ := json.Marshal([]AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}})
						return out, nil
					}
					var out []byte
					if test.queryPayload.Limit == "1" {
						out, _ = json.Marshal([]App{availableApps["name"]})
						return out, nil
					} else {
						apps := []App{}
						for n, a := range availableApps {
							if test.queryPayload.Name == n || test.queryPayload.Name == "" {
								apps = append(apps, a)
							}
						}
						out, _ = json.Marshal(apps)
						return out, nil
					}
				},
			}, "test.com",
			)
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
	}{
		"it gets one app": {
			id: int32(1),
			expectedResponse: &App{
				ID:    oltypes.Int32(1),
				Name:  oltypes.String("name"),
				Rules: []AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}},
			},
		},
		"it returns an error if there is a problem finding the app": {
			id:               int32(2),
			expectedResponse: nil,
			expectedError:    errors.New("not found"),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(MockRepository{
				DoFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1" {
						out, _ := json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("name")})
						return out, nil
					}
					if req.URL == "test.com/api/2/apps/1/rules" {
						out, _ := json.Marshal([]AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}})
						return out, nil
					}
					return nil, errors.New("not found")
				},
			}, "test.com")
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
						ID:   oltypes.Int32(2),
						Name: oltypes.String("new_rule"),
					},
				},
			},
		},
		"it returns an error if there is a problem finding the app": {
			id:               int32(2),
			expectedResponse: nil,
			expectedError:    errors.New("not found"),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(MockRepository{
				DoFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1" {
						a := App{ID: test.expectedResponse.ID, Name: test.expectedResponse.Name}
						out, _ := json.Marshal(a)
						return out, nil
					}
					if req.URL == "test.com/api/2/apps/1/rules" {
						resp := map[string]int32{"id": *test.expectedResponse.Rules[1].ID}
						out, _ := json.Marshal(resp)
						return out, nil
					}
					if req.URL == "test.com/api/2/apps/1/rules/1" {
						resp := map[string]int32{"id": *test.expectedResponse.Rules[0].ID}
						out, _ := json.Marshal(resp)
						return out, nil
					}
					return nil, errors.New("not found")
				},
			}, "test.com")
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
		expectedResponse *App
		expectedError    error
	}{
		"it destroys one app": {
			id:               int32(1),
			expectedResponse: &App{},
		},
		"it returns an error if there is a problem finding the app": {
			id:               int32(2),
			expectedResponse: nil,
			expectedError:    errors.New("not found"),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(MockRepository{
				DoFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					if req.URL == "test.com/api/2/apps/1" {
						return nil, nil
					}
					return nil, errors.New("not found")
				},
			}, "test.com")
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
	}{
		"it creates one app": {
			createPayload:    &App{Name: oltypes.String("name"), Rules: []AppRule{AppRule{Name: oltypes.String("rule")}}},
			expectedResponse: &App{ID: oltypes.Int32(1), Name: oltypes.String("name"), Rules: []AppRule{AppRule{ID: oltypes.Int32(1), Name: oltypes.String("rule")}}},
		},
		"it returns an error if there is a bad request": {
			createPayload:    &App{ID: oltypes.Int32(1), Name: oltypes.String("not allowed value")},
			expectedResponse: nil,
			expectedError:    errors.New("bad request"),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(MockRepository{
				DoFunc: func(r interface{}) ([]byte, error) {
					if test.expectedError == nil {
						req := r.(olhttp.OLHTTPRequest)
						if req.URL == "test.com/api/2/apps" {
							app := req.Payload.(*App)
							app.ID = oltypes.Int32(int32(1))
							resp := App{Name: app.Name, ID: app.ID}
							out, _ := json.Marshal(resp)
							return out, nil
						} else {
							resp := map[string]int{"id": 1}
							out, _ := json.Marshal(resp)
							return out, nil
						}
					}
					return nil, errors.New("bad request")
				},
			}, "test.com")
			actual, err := svc.Create(test.createPayload)
			assert.Equal(t, test.expectedResponse, actual)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}
