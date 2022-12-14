// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// BookController is an autogenerated mock type for the BookController type
type BookController struct {
	mock.Mock
}

// CreateOne provides a mock function with given fields: c
func (_m *BookController) CreateOne(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteOneByID provides a mock function with given fields: c
func (_m *BookController) DeleteOneByID(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMany provides a mock function with given fields: c
func (_m *BookController) GetMany(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetOneByID provides a mock function with given fields: c
func (_m *BookController) GetOneByID(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateOneByID provides a mock function with given fields: c
func (_m *BookController) UpdateOneByID(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewBookController interface {
	mock.TestingT
	Cleanup(func())
}

// NewBookController creates a new instance of BookController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBookController(t mockConstructorTestingTNewBookController) *BookController {
	mock := &BookController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
