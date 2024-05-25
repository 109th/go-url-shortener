// Code generated by mockery v2.43.1. DO NOT EDIT.

package mockery

import mock "github.com/stretchr/testify/mock"

// MockServer is an autogenerated mock type for the Server type
type MockServer struct {
	mock.Mock
}

type MockServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockServer) EXPECT() *MockServer_Expecter {
	return &MockServer_Expecter{mock: &_m.Mock}
}

// GetURL provides a mock function with given fields: id
func (_m *MockServer) GetURL(id string) (string, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetURL")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockServer_GetURL_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetURL'
type MockServer_GetURL_Call struct {
	*mock.Call
}

// GetURL is a helper method to define mock.On call
//   - id string
func (_e *MockServer_Expecter) GetURL(id interface{}) *MockServer_GetURL_Call {
	return &MockServer_GetURL_Call{Call: _e.mock.On("GetURL", id)}
}

func (_c *MockServer_GetURL_Call) Run(run func(id string)) *MockServer_GetURL_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockServer_GetURL_Call) Return(_a0 string, _a1 error) *MockServer_GetURL_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServer_GetURL_Call) RunAndReturn(run func(string) (string, error)) *MockServer_GetURL_Call {
	_c.Call.Return(run)
	return _c
}

// SaveURL provides a mock function with given fields: url
func (_m *MockServer) SaveURL(url string) (string, error) {
	ret := _m.Called(url)

	if len(ret) == 0 {
		panic("no return value specified for SaveURL")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(url)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(url)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockServer_SaveURL_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveURL'
type MockServer_SaveURL_Call struct {
	*mock.Call
}

// SaveURL is a helper method to define mock.On call
//   - url string
func (_e *MockServer_Expecter) SaveURL(url interface{}) *MockServer_SaveURL_Call {
	return &MockServer_SaveURL_Call{Call: _e.mock.On("SaveURL", url)}
}

func (_c *MockServer_SaveURL_Call) Run(run func(url string)) *MockServer_SaveURL_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockServer_SaveURL_Call) Return(_a0 string, _a1 error) *MockServer_SaveURL_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockServer_SaveURL_Call) RunAndReturn(run func(string) (string, error)) *MockServer_SaveURL_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockServer creates a new instance of MockServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockServer {
	mock := &MockServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}