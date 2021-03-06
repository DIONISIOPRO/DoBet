// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	amqp "github.com/streadway/amqp"

	mock "github.com/stretchr/testify/mock"
)

// EventSubscreber is an autogenerated mock type for the EventSubscreber type
type EventSubscreber struct {
	mock.Mock
}

// SubscribeToQueue provides a mock function with given fields: name
func (_m *EventSubscreber) SubscribeToQueue(name string) (<-chan amqp.Delivery, error) {
	ret := _m.Called(name)

	var r0 <-chan amqp.Delivery
	if rf, ok := ret.Get(0).(func(string) <-chan amqp.Delivery); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan amqp.Delivery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
