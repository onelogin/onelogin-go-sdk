package test

import (
	"errors"
)

type MockRepository struct {
	ReadFunc      func(r interface{}) ([][]byte, error)
	ReReadFunc    []func(r interface{}) ([][]byte, error)
	CreateFunc    func(r interface{}) ([]byte, error)
	SubCreateFunc []func(r interface{}) ([]byte, error)
	UpdateFunc    func(r interface{}) ([]byte, error)
	DestroyFunc   func(r interface{}) ([]byte, error)
	ReReads       int
	SubCreates    int
}

// Read will invoke a repository's ReadFunc. The ReadFunc is meant to describe the
// expected return value from the repository, and gives the opportunity to define the
// response based on request parameters such as the request URL. A ReReadFunc is allowed
// because RESTful reads are idempotent, however, you may want to mock a subsequent read,
// following an error recovery scenario where some updates were saved, but others were abandoned.
func (mr *MockRepository) Read(r interface{}) ([][]byte, error) {
	if mr.ReReads > 0 && len(mr.ReReadFunc) > 0 {
		mr.ReReads++
		return mr.ReReadFunc[mr.ReReads-2](r)
	}
	if mr.ReadFunc != nil {
		mr.ReReads++
		return mr.ReadFunc(r)
	}
	return nil, errors.New("error")
}

// Create invokes the repository's CreateFunc giving the opportunity to specify the return
// value for a repository's Create function
func (mr *MockRepository) Create(r interface{}) ([]byte, error) {
	if mr.SubCreates > 0 && len(mr.SubCreateFunc) > 0 {
		mr.SubCreates++
		return mr.SubCreateFunc[mr.SubCreates-2](r)
	}
	if mr.CreateFunc != nil {
		mr.SubCreates++
		return mr.CreateFunc(r)
	}
	return nil, errors.New("error")
}

// Update invokes the repository's UpdateFunc giving the opportunity to specify the return
// value for a repository's Update function
func (mr *MockRepository) Update(r interface{}) ([]byte, error) {
	if mr.UpdateFunc != nil {
		return mr.UpdateFunc(r)
	}
	return nil, errors.New("error")
}

// Destroy invokes the repository's DestroyFunc giving the opportunity to specify the return
// value for a repository's Destroy function
func (mr *MockRepository) Destroy(r interface{}) ([]byte, error) {
	if mr.DestroyFunc != nil {
		return mr.DestroyFunc(r)
	}
	return nil, errors.New("error")
}
