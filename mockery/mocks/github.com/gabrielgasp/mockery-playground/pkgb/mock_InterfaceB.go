// Code generated by mockery v2.43.2. DO NOT EDIT.

package pkgb

import mock "github.com/stretchr/testify/mock"

// MockInterfaceB is an autogenerated mock type for the InterfaceB type
type MockInterfaceB struct {
	mock.Mock
}

type MockInterfaceB_Expecter struct {
	mock *mock.Mock
}

func (_m *MockInterfaceB) EXPECT() *MockInterfaceB_Expecter {
	return &MockInterfaceB_Expecter{mock: &_m.Mock}
}

// MethodB provides a mock function with given fields:
func (_m *MockInterfaceB) MethodB() {
	_m.Called()
}

// MockInterfaceB_MethodB_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MethodB'
type MockInterfaceB_MethodB_Call struct {
	*mock.Call
}

// MethodB is a helper method to define mock.On call
func (_e *MockInterfaceB_Expecter) MethodB() *MockInterfaceB_MethodB_Call {
	return &MockInterfaceB_MethodB_Call{Call: _e.mock.On("MethodB")}
}

func (_c *MockInterfaceB_MethodB_Call) Run(run func()) *MockInterfaceB_MethodB_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockInterfaceB_MethodB_Call) Return() *MockInterfaceB_MethodB_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockInterfaceB_MethodB_Call) RunAndReturn(run func()) *MockInterfaceB_MethodB_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockInterfaceB creates a new instance of MockInterfaceB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockInterfaceB(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockInterfaceB {
	mock := &MockInterfaceB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
