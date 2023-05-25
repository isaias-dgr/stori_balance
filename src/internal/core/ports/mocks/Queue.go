// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Queue is an autogenerated mock type for the Queue type
type Queue struct {
	mock.Mock
}

// Send provides a mock function with given fields: message
func (_m *Queue) Send(message *string) error {
	ret := _m.Called(message)

	var r0 error
	if rf, ok := ret.Get(0).(func(*string) error); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewQueue interface {
	mock.TestingT
	Cleanup(func())
}

// NewQueue creates a new instance of Queue. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewQueue(t mockConstructorTestingTNewQueue) *Queue {
	mock := &Queue{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}