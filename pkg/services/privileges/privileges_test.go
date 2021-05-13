package privileges

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/internal/test"
	"github.com/onelogin/onelogin-go-sdk/pkg/oltypes"
	"github.com/stretchr/testify/assert"
)

func TestQuery(t *testing.T) {
	tests := map[string]struct {
		queryPayload     *PrivilegeQuery
		expectedResponse []Privilege
		expectedError    []error
		repository       *test.MockRepository
	}{
		"it gets one privilege with assigned resources": {
			queryPayload:     &PrivilegeQuery{Limit: "1"},
			expectedResponse: []Privilege{Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("privilege")}},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal([]Privilege{Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("privilege")}})
				},
			},
		},
		"it returns the remote default limit of privileges and assigned resources if no query given": {
			queryPayload: &PrivilegeQuery{},
			expectedResponse: []Privilege{
				Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("privilege")},
				Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("name2")},
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal([]Privilege{
						Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("privilege")},
						Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("name2")},
					})
				},
			},
		},
		"it returns an error if the call to /privileges fails": {
			queryPayload:     &PrivilegeQuery{},
			expectedError:    []error{errors.New("error")},
			expectedResponse: nil,
			repository:       &test.MockRepository{},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			actual, err := svc.QueryWithAssignments(test.queryPayload)
			assert.Equal(t, test.expectedResponse, actual)
			if len(test.expectedError) > 0 {
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestQueryWithAssignments(t *testing.T) {
	tests := map[string]struct {
		queryPayload     *PrivilegeQuery
		expectedResponse []Privilege
		expectedError    []error
		repository       *test.MockRepository
	}{
		"it gets one privilege with assigned resources": {
			queryPayload:     &PrivilegeQuery{Limit: "1"},
			expectedResponse: []Privilege{Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("privilege"), RoleIDs: []int{1, 2, 3}, UserIDs: []int{5, 6, 7}}},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal([]Privilege{Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("privilege")}})
				},
				ReReadFunc: []func(r interface{}) ([]byte, error){
					func(r interface{}) ([]byte, error) {
						return json.Marshal(map[string]interface{}{"total": 4, "roles": []int{1, 2, 3}})
					},
					func(r interface{}) ([]byte, error) {
						return json.Marshal(map[string]interface{}{"total": 4, "users": []int{5, 6, 7}})
					},
				},
			},
		},
		"it returns the remote default limit of privileges and assigned resources if no query given": {
			queryPayload: &PrivilegeQuery{},
			expectedResponse: []Privilege{
				Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("privilege"), UserIDs: []int{5, 6, 7}, RoleIDs: []int{1, 2, 3}},
				Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("name2"), UserIDs: []int{5, 6, 7}, RoleIDs: []int{1, 2, 3}},
			},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal([]Privilege{
						Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("privilege")},
						Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("name2")},
					})
				},
				ReReadFunc: []func(r interface{}) ([]byte, error){
					func(r interface{}) ([]byte, error) {
						return json.Marshal(map[string]interface{}{"total": 4, "roles": []int{1, 2, 3}})
					},
					func(r interface{}) ([]byte, error) {
						return json.Marshal(map[string]interface{}{"total": 4, "users": []int{5, 6, 7}})
					},
					func(r interface{}) ([]byte, error) {
						return json.Marshal(map[string]interface{}{"total": 4, "roles": []int{1, 2, 3}})
					},
					func(r interface{}) ([]byte, error) {
						return json.Marshal(map[string]interface{}{"total": 4, "users": []int{5, 6, 7}})
					},
				},
			},
		},
		"it returns many privileges and assigned resources, where no assigned resources included if subsequent assigned resource call fails intermettently if no query given": {
			queryPayload: &PrivilegeQuery{},
			expectedResponse: []Privilege{
				Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("privilege"), RoleIDs: []int{1, 2, 3}},
				Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("name2"), UserIDs: []int{5, 6, 7}, RoleIDs: []int{1, 2, 3}},
			},
			expectedError: []error{errors.New("unable to read [users]"), nil},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal([]Privilege{
						Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("privilege")},
						Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("name2")},
					})
				},
				ReReadFunc: []func(r interface{}) ([]byte, error){
					func(r interface{}) ([]byte, error) {
						return json.Marshal(map[string]interface{}{"total": 4, "roles": []int{1, 2, 3}})
					},
					func(r interface{}) ([]byte, error) {
						return []byte{}, errors.New("fail")
					},
					func(r interface{}) ([]byte, error) {
						return json.Marshal(map[string]interface{}{"total": 4, "roles": []int{1, 2, 3}})
					},
					func(r interface{}) ([]byte, error) {
						return json.Marshal(map[string]interface{}{"total": 4, "users": []int{5, 6, 7}})
					},
				},
			},
		},
		"it returns an error if the call to /privileges fails": {
			queryPayload:     &PrivilegeQuery{},
			expectedError:    []error{errors.New("error")},
			expectedResponse: nil,
			repository:       &test.MockRepository{},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			actual, err := svc.QueryWithAssignments(test.queryPayload)
			assert.Equal(t, test.expectedResponse, actual)
			if len(test.expectedError) > 0 {
				fmt.Println("actual", actual, "exp", test.expectedResponse, "err", err, "should err", test.expectedError)
				assert.Equal(t, test.expectedError, err)
			}
		})
	}
}

func TestGetOne(t *testing.T) {
	tests := map[string]struct {
		id               string
		expectedResponse *Privilege
		expectedError    error
		repository       *test.MockRepository
	}{
		"it gets one privilege by id": {
			id:               "asdf",
			expectedResponse: &Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("first"), RoleIDs: []int{1, 2, 3}, UserIDs: []int{5, 6, 7}},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("first")})
				},
				ReReadFunc: []func(r interface{}) ([]byte, error){
					func(r interface{}) ([]byte, error) {
						return json.Marshal(map[string]interface{}{"total": 4, "roles": []int{1, 2, 3}})
					},
					func(r interface{}) ([]byte, error) {
						return json.Marshal(map[string]interface{}{"total": 4, "users": []int{5, 6, 7}})
					},
				},
			},
		},
		"it returns a privilege by without attached resources if attached resource request fails": {
			id:               "asdf",
			expectedResponse: &Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("first")},
			repository: &test.MockRepository{
				ReadFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("first")})
				},
				ReReadFunc: []func(r interface{}) ([]byte, error){
					func(r interface{}) ([]byte, error) {
						return json.Marshal(map[string]interface{}{"total": 4, "garbage": 123})
					},
					func(r interface{}) ([]byte, error) {
						return []byte{}, errors.New("fail")
					},
				},
			},
		},
		"it returns an error if the call to /privileges fails": {
			expectedError:    errors.New("error"),
			expectedResponse: nil,
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
		updatePayload  *Privilege
		expectedResult *Privilege
		expectedError  error
		repository     *test.MockRepository
	}{
		"it updates a privilege": {
			updatePayload:  &Privilege{Name: oltypes.String("update"), ID: oltypes.String("asdf")},
			expectedResult: &Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("update")},
			repository: &test.MockRepository{
				UpdateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("update")})
				},
			},
		},
		"it returns an error if no id on resource": {
			updatePayload:  &Privilege{Name: oltypes.String("update")},
			expectedResult: &Privilege{},
			expectedError:  errors.New("No ID Given"),
			repository:     &test.MockRepository{},
		},
		"it returns an error if something went wrong": {
			updatePayload:  &Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("update")},
			expectedResult: &Privilege{},
			expectedError:  errors.New("error"),
			repository:     &test.MockRepository{},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			err := svc.Update(test.updatePayload)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Equal(t, test.expectedResult, test.updatePayload)
			}
		})
	}
}

func TestCreate(t *testing.T) {
	tests := map[string]struct {
		createPayload  *Privilege
		expectedResult *Privilege
		expectedError  error
		repository     *test.MockRepository
	}{
		"it creates a privilege": {
			createPayload:  &Privilege{Name: oltypes.String("create")},
			expectedResult: &Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("create")},
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("create")})
				},
				SubCreateFunc: []func(r interface{}) ([]byte, error){
					func(r interface{}) ([]byte, error) {
						return json.Marshal(map[string]bool{"success": true})
					},
				},
			},
		},
		"it creates a privilege and adds users and roles": {
			createPayload:  &Privilege{Name: oltypes.String("create"), RoleIDs: []int{1, 2, 3}, UserIDs: []int{5, 6, 7}},
			expectedResult: &Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("create"), RoleIDs: []int{1, 2, 3}, UserIDs: []int{5, 6, 7}},
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("create")})
				},
				SubCreateFunc: []func(r interface{}) ([]byte, error){
					func(r interface{}) ([]byte, error) {
						return json.Marshal(map[string]bool{"success": true})
					},
					func(r interface{}) ([]byte, error) {
						return json.Marshal(map[string]bool{"success": true})
					},
				},
			},
		},
		"it returns an error if attaching resources fails": {
			createPayload: &Privilege{Name: oltypes.String("create"), RoleIDs: []int{1, 2, 3}, UserIDs: []int{5, 6, 7}},
			expectedError: errors.New("unable to assign [roles users]"),
			repository: &test.MockRepository{
				CreateFunc: func(r interface{}) ([]byte, error) {
					return json.Marshal(Privilege{ID: oltypes.String("asdf"), Name: oltypes.String("create")})
				},
				SubCreateFunc: []func(r interface{}) ([]byte, error){
					func(r interface{}) ([]byte, error) {
						return json.Marshal(map[string]bool{"garbage": true})
					},
					func(r interface{}) ([]byte, error) {
						return json.Marshal(map[string]bool{"success": false})
					},
				},
				DestroyFunc: func(r interface{}) ([]byte, error) {
					return nil, nil
				},
			},
		},
		"it returns an error if something went wrong": {
			createPayload:  &Privilege{Name: oltypes.String("create")},
			expectedResult: &Privilege{},
			expectedError:  errors.New("error"),
			repository:     &test.MockRepository{},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			svc := New(test.repository, "test.com")
			err := svc.Create(test.createPayload)
			if test.expectedError != nil {
				assert.Equal(t, test.expectedError, err)
			} else {
				assert.Equal(t, test.expectedResult, test.createPayload)
			}
		})
	}
}

func TestDestroy(t *testing.T) {
	tests := map[string]struct {
		id               string
		repository       *test.MockRepository
		expectedResponse *Privilege
		expectedError    error
	}{
		"it destroys one privilege": {
			id: "test",
			repository: &test.MockRepository{
				DestroyFunc: func(r interface{}) ([]byte, error) {
					return nil, nil
				},
			},
			expectedResponse: &Privilege{},
		},
		"it returns an error if there is a problem finding the app": {
			id:               "test",
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
