package event

import (
	mocks "github/namuethopro/dobet-user/mocks/event"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestListenningToqueues(t *testing.T) {
	mockEventProcessor := new(mocks.EventProcessor)
	mockEventSubscriber := new(mocks.EventSubscreber)
	mockEventSubscriber.On("SubscribeToQueue", mock.Anything).Return(nil, nil)
	listenner := NewRabbitMQEventListenner(mockEventProcessor, mockEventSubscriber, nil)
	listenner.ListenningToqueues()
	mockEventSubscriber.AssertExpectations(t)
}
