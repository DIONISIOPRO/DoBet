// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	domain "github/namuethopro/dobet-user/domain"

	mock "github.com/stretchr/testify/mock"
)

// IncomingEventProcessorConfirmEventManager is an autogenerated mock type for the IncomingEventProcessorConfirmEventManager type
type IncomingEventProcessorConfirmEventManager struct {
	mock.Mock
}

// Publish provides a mock function with given fields: name, _a1
func (_m *IncomingEventProcessorConfirmEventManager) Publish(name string, _a1 domain.Event) error {
	ret := _m.Called(name, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, domain.Event) error); ok {
		r0 = rf(name, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
