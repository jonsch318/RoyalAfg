// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	dtos "github.com/jonsch318/royalafg/pkg/dtos"
	mock "github.com/stretchr/testify/mock"

	models "github.com/jonsch318/royalafg/pkg/models"

	mw "github.com/jonsch318/royalafg/pkg/mw"
)

// IAuthentication is an autogenerated mock type for the IAuthentication type
type IAuthentication struct {
	mock.Mock
}

// Login provides a mock function with given fields: username, password
func (_m *IAuthentication) Login(username string, password string) (*models.User, string, error) {
	ret := _m.Called(username, password)

	var r0 *models.User
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (*models.User, string, error)); ok {
		return rf(username, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) *models.User); ok {
		r0 = rf(username, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) string); ok {
		r1 = rf(username, password)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(username, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Logout provides a mock function with given fields: id
func (_m *IAuthentication) Logout(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Register provides a mock function with given fields: dto
func (_m *IAuthentication) Register(dto *dtos.RegisterDto) (*models.User, string, error) {
	ret := _m.Called(dto)

	var r0 *models.User
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(*dtos.RegisterDto) (*models.User, string, error)); ok {
		return rf(dto)
	}
	if rf, ok := ret.Get(0).(func(*dtos.RegisterDto) *models.User); ok {
		r0 = rf(dto)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*dtos.RegisterDto) string); ok {
		r1 = rf(dto)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(*dtos.RegisterDto) error); ok {
		r2 = rf(dto)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// VerifyAuthentication provides a mock function with given fields: user
func (_m *IAuthentication) VerifyAuthentication(user *mw.UserClaims) bool {
	ret := _m.Called(user)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*mw.UserClaims) bool); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// NewIAuthentication creates a new instance of IAuthentication. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIAuthentication(t interface {
	mock.TestingT
	Cleanup(func())
}) *IAuthentication {
	mock := &IAuthentication{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
