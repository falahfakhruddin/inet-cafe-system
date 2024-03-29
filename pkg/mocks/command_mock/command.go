// Code generated by mockery v2.32.0. DO NOT EDIT.

package commandmock

import mock "github.com/stretchr/testify/mock"

// Command is an autogenerated mock type for the Command type
type Command struct {
	mock.Mock
}

// ParseArgs provides a mock function with given fields: _a0
func (_m *Command) ParseArgs(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Run provides a mock function with given fields:
func (_m *Command) Run() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCommand creates a new instance of Command. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCommand(t interface {
	mock.TestingT
	Cleanup(func())
}) *Command {
	mock := &Command{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
