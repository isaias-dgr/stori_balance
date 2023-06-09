// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/isaias-dgr/stori-balance/src/internal/core/domain"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetUser provides a mock function with given fields: id
func (_m *Repository) GetUser(id string) (*domain.User, error) {
	ret := _m.Called(id)

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: account, product, transaction
func (_m *Repository) Save(account string, product string, transaction *domain.Transaction) error {
	ret := _m.Called(account, product, transaction)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *domain.Transaction) error); ok {
		r0 = rf(account, product, transaction)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
