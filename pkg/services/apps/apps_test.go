package apps

import (
	"encoding/json"
	"errors"
	"github.com/onelogin/onelogin-go-sdk/internal/test"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/olhttp"
	"github.com/stretchr/testify/assert"
	"testing"
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
				App{ID: oltypes.Int32(1), Name: oltypes.String("name")},
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal([]App{App{ID: oltypes.Int32(1), Name: oltypes.String("name")}})
				},
			},
		},
		"it returns the remote default limit of apps if no query given": {
			queryPayload: &AppsQuery{},
			expectedResponse: []App{
				App{ID: oltypes.Int32(1), Name: oltypes.String("name")},
				App{ID: oltypes.Int32(1), Name: oltypes.String("name2")},
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
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
			expectedResponse: &App{ID: oltypes.Int32(1), Name: oltypes.String("name")},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
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
			},
			expectedResponse: &App{
				ID:   oltypes.Int32(1),
				Name: oltypes.String("updated"),
			},
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(map[string]int32{"id": 2})
				},
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("updated")})
				},
				ReadFunc: func(r interface{}) ([]byte, error) {
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
		"it removes parameters when requested": {
			updatePayload: &App{
				ID:         oltypes.Int32(1),
				Name:       oltypes.String("original"),
				Parameters: map[string]AppParameters{},
			},
			expectedResponse: &App{ID: oltypes.Int32(1), Name: oltypes.String("name")},
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("name"), Parameters: map[string]AppParameters{"test": AppParameters{ID: oltypes.Int32(1)}}})
				},
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("name")})
				},
				ReReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("name")})
				},
				DestroyFunc: func(r interface{}) ([]byte, error) {
					return nil, nil
				},
			},
		},
		"it recovers if delete parameter fails": {
			updatePayload: &App{
				ID:         oltypes.Int32(1),
				Name:       oltypes.String("original"),
				Parameters: map[string]AppParameters{},
			},
			expectedResponse: &App{
				ID:         oltypes.Int32(1),
				Name:       oltypes.String("name"),
				Parameters: map[string]AppParameters{"test": AppParameters{ID: oltypes.Int32(1)}},
			},
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("name"), Parameters: map[string]AppParameters{"test": AppParameters{ID: oltypes.Int32(1)}}})
				},
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(App{ID: oltypes.Int32(1), Name: oltypes.String("name"), Parameters: map[string]AppParameters{"test": AppParameters{ID: oltypes.Int32(1)}}})
				},
				DestroyFunc: func(r interface{}) ([]byte, error) {
					return nil, errors.New("error")
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
			createPayload:    &App{Name: oltypes.String("name")},
			expectedResponse: &App{ID: oltypes.Int32(1), Name: oltypes.String("name")},
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					req := r.(olhttp.OLHTTPRequest)
					app := req.Payload.(*App)
					app.ID = oltypes.Int32(int32(1))
					resp := App{Name: app.Name, ID: app.ID}
					return json.Marshal(resp)
				},
				ReadFunc: func(r interface{}) ([]byte, error) {
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
