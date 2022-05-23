// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/dionisiopro/dobet-auth/domain"
	mock "github.com/stretchr/testify/mock"
)

// AuthEventPublisher is an autogenerated mock type for the AuthEventPublisher type
type AuthEventPublisher struct {
	mock.Mock
}

// Publish provides a mock function with given fields: name, event
func (_m *AuthEventPublisher) Publish(name string, event domain.Event) error {
	ret := _m.Called(name, event)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, domain.Event) error); ok {
		r0 = rf(name, event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
