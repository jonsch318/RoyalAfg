// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	protos "github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	mock "github.com/stretchr/testify/mock"
)

// UserServiceServer is an autogenerated mock type for the UserServiceServer type
type UserServiceServer struct {
	mock.Mock
}

// GetUserById provides a mock function with given fields: _a0, _a1
func (_m *UserServiceServer) GetUserById(_a0 context.Context, _a1 *protos.GetUser) (*protos.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *protos.User
	if rf, ok := ret.Get(0).(func(context.Context, *protos.GetUser) *protos.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protos.GetUser) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByUsername provides a mock function with given fields: _a0, _a1
func (_m *UserServiceServer) GetUserByUsername(_a0 context.Context, _a1 *protos.GetUser) (*protos.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *protos.User
	if rf, ok := ret.Get(0).(func(context.Context, *protos.GetUser) *protos.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protos.GetUser) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveUser provides a mock function with given fields: _a0, _a1
func (_m *UserServiceServer) SaveUser(_a0 context.Context, _a1 *protos.User) (*protos.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *protos.User
	if rf, ok := ret.Get(0).(func(context.Context, *protos.User) *protos.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protos.User) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: _a0, _a1
func (_m *UserServiceServer) UpdateUser(_a0 context.Context, _a1 *protos.User) (*protos.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *protos.User
	if rf, ok := ret.Get(0).(func(context.Context, *protos.User) *protos.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *protos.User) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}