package mocks

import (
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) GetUsers(query *models.UserQuery) ([]models.User, error) {
	args := m.Called(query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockClient) GetApps(query *models.AppQuery) ([]models.App, error) {
	args := m.Called(query)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.App), args.Error(1)
}

func (m *MockClient) CreateUser(user models.User) (models.User, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return models.User{}, args.Error(1)
	}
	return args.Get(0).(models.User), args.Error(1)
}
