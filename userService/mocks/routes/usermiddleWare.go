// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// UsermiddleWare is an autogenerated mock type for the UsermiddleWare type
type UsermiddleWare struct {
	mock.Mock
}

// Authenticated provides a mock function with given fields:
func (_m *UsermiddleWare) Authenticated() gin.HandlerFunc {
	ret := _m.Called()

	var r0 gin.HandlerFunc
	if rf, ok := ret.Get(0).(func() gin.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gin.HandlerFunc)
		}
	}

	return r0
}

// IsAdmin provides a mock function with given fields:
func (_m *UsermiddleWare) IsAdmin() gin.HandlerFunc {
	ret := _m.Called()

	var r0 gin.HandlerFunc
	if rf, ok := ret.Get(0).(func() gin.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gin.HandlerFunc)
		}
	}

	return r0
}

// IsOwner provides a mock function with given fields:
func (_m *UsermiddleWare) IsOwner() gin.HandlerFunc {
	ret := _m.Called()

	var r0 gin.HandlerFunc
	if rf, ok := ret.Get(0).(func() gin.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(gin.HandlerFunc)
		}
	}

	return r0
}
