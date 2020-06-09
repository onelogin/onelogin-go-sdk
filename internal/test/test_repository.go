package test

import "errors"

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
