// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// AuthEventQueueCreator is an autogenerated mock type for the AuthEventQueueCreator type
type AuthEventQueueCreator struct {
	mock.Mock
}

// CreateQueues provides a mock function with given fields: _a0
func (_m *AuthEventQueueCreator) CreateQueues(_a0 []string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}