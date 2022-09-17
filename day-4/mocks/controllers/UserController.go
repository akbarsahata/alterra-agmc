// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// UserController is an autogenerated mock type for the UserController type
type UserController struct {
	mock.Mock
}

// CreateOne provides a mock function with given fields: c
func (_m *UserController) CreateOne(c echo.Context) error {
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
func (_m *UserController) DeleteOneByID(c echo.Context) error {
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
func (_m *UserController) GetMany(c echo.Context) error {
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
func (_m *UserController) GetOneByID(c echo.Context) error {
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
func (_m *UserController) UpdateOneByID(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserController interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserController creates a new instance of UserController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserController(t mockConstructorTestingTNewUserController) *UserController {
	mock := &UserController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
