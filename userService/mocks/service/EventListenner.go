// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// EventListenner is an autogenerated mock type for the EventListenner type
type EventListenner struct {
	mock.Mock
}

// ListenningToqueues provides a mock function with given fields: done
func (_m *EventListenner) ListenningToqueues(done <-chan bool) {
	_m.Called(done)
}
