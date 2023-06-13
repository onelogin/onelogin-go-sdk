package authentication

import (
	"errors"
)

type MockAuthenticator struct {
	GetTokenFunc         func() (string, error)
	NewAuthenticatorFunc func() Authenticator
	GenerateTokenFunc    func() error
	RevokeTokenFunc      func(token, domain *string) error
}

func (a *MockAuthenticator) GetToken() (string, error) {
	if a.GetTokenFunc != nil {
		return a.GetTokenFunc()
	}
	return "", errors.New("GetTokenFunc is not implemented")
}

func (a *MockAuthenticator) NewAuthenticatorF() Authenticator {
	if a.NewAuthenticatorFunc != nil {
		return a.NewAuthenticatorFunc()
	}
	return *NewAuthenticator()
}

func (a *MockAuthenticator) GenerateToken() error {
	if a.GenerateTokenFunc != nil {
		return a.GenerateTokenFunc()
	}
	return errors.New("GenerateTokenFunc is not implemented")
}

func (a *MockAuthenticator) RevokeToken(token, domain *string) error {
	if a.RevokeTokenFunc != nil {
		return a.RevokeTokenFunc(token, domain)
	}
	return errors.New("RevokeTokenFunc is not implemented")
}
