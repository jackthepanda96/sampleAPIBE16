// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	book "ormapi/app/features/book"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetAll provides a mock function with given fields:
func (_m *Repository) GetAll() (interface{}, error) {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: newBook, user_id
func (_m *Repository) Insert(newBook book.Core, user_id string) (book.Core, error) {
	ret := _m.Called(newBook, user_id)

	var r0 book.Core
	if rf, ok := ret.Get(0).(func(book.Core, string) book.Core); ok {
		r0 = rf(newBook, user_id)
	} else {
		r0 = ret.Get(0).(book.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(book.Core, string) error); ok {
		r1 = rf(newBook, user_id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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
