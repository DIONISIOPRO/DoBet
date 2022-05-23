package event

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/dionisiopro/dobet-user/domain"
	mocks "github.com/dionisiopro/dobet-user/mocks/event"

	"github.com/stretchr/testify/assert"
)

func TestAddBalance(t *testing.T) {
	t.Run("sucess", func(t *testing.T) {
		event := domain.AddMoneyEvent{
			UserId: "id",
			Amount: 10,
		}
		data, err := json.Marshal(event)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		mockrepo := new(mocks.IncomingEventProcessorRepository)
		assert.NotNil(t, mockrepo)
		mockrepo.On("AddMoney", event.UserId, event.Amount).Return(nil).Once()
		handler := NewIncomingEventProcessor(mockrepo)
		assert.NotNil(t, handler)
		err = handler.AddBalance(data)
		assert.NoError(t, err)
		mockrepo.AssertExpectations(t)
	})

	t.Run("fail on epo", func(t *testing.T) {
		event := domain.AddMoneyEvent{
			UserId: "id",
			Amount: 10,
		}
		data, err := json.Marshal(event)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		mockrepo := new(mocks.IncomingEventProcessorRepository)
		assert.NotNil(t, mockrepo)
		mockrepo.On("AddMoney", event.UserId, event.Amount).Return(errors.New("")).Once()
		handler := NewIncomingEventProcessor(mockrepo)
		assert.NotNil(t, handler)
		err = handler.AddBalance(data)
		assert.NotNil(t, err)
		mockrepo.AssertExpectations(t)
	})
}

func TestSubtractBalance(t *testing.T) {
	t.Run("sucess", func(t *testing.T) {
		event := domain.SubtractMoneyEvent{
			UserId: "id",
			Amount: 10,
		}
		data, err := json.Marshal(event)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		mockrepo := new(mocks.IncomingEventProcessorRepository)
		assert.NotNil(t, mockrepo)
		mockrepo.On("SubtractMoney", event.UserId, event.Amount).Return(nil).Once()
		handler := NewIncomingEventProcessor(mockrepo)
		assert.NotNil(t, handler)
		err = handler.SubtractBalance(data)
		assert.NoError(t, err)
		mockrepo.AssertExpectations(t)
	})

	t.Run("fail on epo", func(t *testing.T) {
		event := domain.SubtractMoneyEvent{
			UserId: "id",
			Amount: 10,
		}
		data, err := json.Marshal(event)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		mockrepo := new(mocks.IncomingEventProcessorRepository)
		assert.NotNil(t, mockrepo)
		mockrepo.On("SubtractMoney", event.UserId, event.Amount).Return(errors.New("")).Once()
		handler := NewIncomingEventProcessor(mockrepo)
		assert.NotNil(t, handler)
		err = handler.SubtractBalance(data)
		assert.NotNil(t, err)
		mockrepo.AssertExpectations(t)
	})
}
