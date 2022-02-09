// Code generated by mockery v2.7.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IQueue is an autogenerated mock type for the IQueue type
type IQueue struct {
	mock.Mock
}

// Dequeue provides a mock function with given fields:
func (_m *IQueue) Dequeue() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// Enqueue provides a mock function with given fields: _a0
func (_m *IQueue) Enqueue(_a0 interface{}) {
	_m.Called(_a0)
}

// Length provides a mock function with given fields:
func (_m *IQueue) Length() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}