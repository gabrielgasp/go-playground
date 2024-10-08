// Code generated by mockery v2.43.2. DO NOT EDIT.

package io

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// MockWriterTo is an autogenerated mock type for the WriterTo type
type MockWriterTo struct {
	mock.Mock
}

type MockWriterTo_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWriterTo) EXPECT() *MockWriterTo_Expecter {
	return &MockWriterTo_Expecter{mock: &_m.Mock}
}

// WriteTo provides a mock function with given fields: w
func (_m *MockWriterTo) WriteTo(w io.Writer) (int64, error) {
	ret := _m.Called(w)

	if len(ret) == 0 {
		panic("no return value specified for WriteTo")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(io.Writer) (int64, error)); ok {
		return rf(w)
	}
	if rf, ok := ret.Get(0).(func(io.Writer) int64); ok {
		r0 = rf(w)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(io.Writer) error); ok {
		r1 = rf(w)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWriterTo_WriteTo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WriteTo'
type MockWriterTo_WriteTo_Call struct {
	*mock.Call
}

// WriteTo is a helper method to define mock.On call
//   - w io.Writer
func (_e *MockWriterTo_Expecter) WriteTo(w interface{}) *MockWriterTo_WriteTo_Call {
	return &MockWriterTo_WriteTo_Call{Call: _e.mock.On("WriteTo", w)}
}

func (_c *MockWriterTo_WriteTo_Call) Run(run func(w io.Writer)) *MockWriterTo_WriteTo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(io.Writer))
	})
	return _c
}

func (_c *MockWriterTo_WriteTo_Call) Return(n int64, err error) *MockWriterTo_WriteTo_Call {
	_c.Call.Return(n, err)
	return _c
}

func (_c *MockWriterTo_WriteTo_Call) RunAndReturn(run func(io.Writer) (int64, error)) *MockWriterTo_WriteTo_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWriterTo creates a new instance of MockWriterTo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWriterTo(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWriterTo {
	mock := &MockWriterTo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
