// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/dionisiopro/dobet-auth/domain"

	mock "github.com/stretchr/testify/mock"
)

// IncomingEventProcessorRepository is an autogenerated mock type for the IncomingEventProcessorRepository type
type IncomingEventProcessorRepository struct {
	mock.Mock
}

// AddUser provides a mock function with given fields: user
func (_m *IncomingEventProcessorRepository) AddUser(user domain.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveUser provides a mock function with given fields: userId
func (_m *IncomingEventProcessorRepository) RemoveUser(userId string) error {
	ret := _m.Called(userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: userid, user
func (_m *IncomingEventProcessorRepository) UpdateUser(userid string, user domain.User) error {
	ret := _m.Called(userid, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, domain.User) error); ok {
		r0 = rf(userid, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
