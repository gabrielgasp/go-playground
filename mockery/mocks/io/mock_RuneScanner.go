// Code generated by mockery v2.43.2. DO NOT EDIT.

package io

import mock "github.com/stretchr/testify/mock"

// MockRuneScanner is an autogenerated mock type for the RuneScanner type
type MockRuneScanner struct {
	mock.Mock
}

type MockRuneScanner_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRuneScanner) EXPECT() *MockRuneScanner_Expecter {
	return &MockRuneScanner_Expecter{mock: &_m.Mock}
}

// ReadRune provides a mock function with given fields:
func (_m *MockRuneScanner) ReadRune() (rune, int, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReadRune")
	}

	var r0 rune
	var r1 int
	var r2 error
	if rf, ok := ret.Get(0).(func() (rune, int, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() rune); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(rune)
	}

	if rf, ok := ret.Get(1).(func() int); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(int)
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockRuneScanner_ReadRune_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadRune'
type MockRuneScanner_ReadRune_Call struct {
	*mock.Call
}

// ReadRune is a helper method to define mock.On call
func (_e *MockRuneScanner_Expecter) ReadRune() *MockRuneScanner_ReadRune_Call {
	return &MockRuneScanner_ReadRune_Call{Call: _e.mock.On("ReadRune")}
}

func (_c *MockRuneScanner_ReadRune_Call) Run(run func()) *MockRuneScanner_ReadRune_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRuneScanner_ReadRune_Call) Return(r rune, size int, err error) *MockRuneScanner_ReadRune_Call {
	_c.Call.Return(r, size, err)
	return _c
}

func (_c *MockRuneScanner_ReadRune_Call) RunAndReturn(run func() (rune, int, error)) *MockRuneScanner_ReadRune_Call {
	_c.Call.Return(run)
	return _c
}

// UnreadRune provides a mock function with given fields:
func (_m *MockRuneScanner) UnreadRune() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for UnreadRune")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRuneScanner_UnreadRune_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UnreadRune'
type MockRuneScanner_UnreadRune_Call struct {
	*mock.Call
}

// UnreadRune is a helper method to define mock.On call
func (_e *MockRuneScanner_Expecter) UnreadRune() *MockRuneScanner_UnreadRune_Call {
	return &MockRuneScanner_UnreadRune_Call{Call: _e.mock.On("UnreadRune")}
}

func (_c *MockRuneScanner_UnreadRune_Call) Run(run func()) *MockRuneScanner_UnreadRune_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRuneScanner_UnreadRune_Call) Return(_a0 error) *MockRuneScanner_UnreadRune_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRuneScanner_UnreadRune_Call) RunAndReturn(run func() error) *MockRuneScanner_UnreadRune_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRuneScanner creates a new instance of MockRuneScanner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRuneScanner(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRuneScanner {
	mock := &MockRuneScanner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}