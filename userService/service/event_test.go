package service

import (
	"encoding/json"
	"errors"
	"github/namuethopro/dobet-user/domain"
	mocks "github/namuethopro/dobet-user/mocks/service"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddBalance(t *testing.T) {
	t.Run("sucess", func(t *testing.T) {
		lock := &sync.Mutex{}
		event := domain.AddMoneyEvent{
			UserId: "id",
			Amount: 10,
		}
		data, err := json.Marshal(event)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		mockrepo := new(mocks.IventHandlerUserRepository)
		assert.NotNil(t, mockrepo)
		mockrepo.On("AddMoney", event.UserId, event.Amount).Return(nil).Once()
		handler := NewIncomingEventHandler(lock, mockrepo, nil, nil)
		assert.NotNil(t, handler)
		err = handler.AddBalance(data)
		assert.NoError(t, err)
		mockrepo.AssertExpectations(t)
	})

	t.Run("fail on epo", func(t *testing.T) {
		lock := &sync.Mutex{}
		event := domain.AddMoneyEvent{
			UserId: "id",
			Amount: 10,
		}
		data, err := json.Marshal(event)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		mockrepo := new(mocks.IventHandlerUserRepository)
		assert.NotNil(t, mockrepo)
		mockrepo.On("AddMoney", event.UserId, event.Amount).Return(errors.New("")).Once()
		handler := NewIncomingEventHandler(lock, mockrepo, nil, nil)
		assert.NotNil(t, handler)
		err = handler.AddBalance(data)
		assert.NotNil(t, err)
		mockrepo.AssertExpectations(t)
	})
}

func TestSubtractBalance(t *testing.T) {
	t.Run("sucess", func(t *testing.T) {
		lock := &sync.Mutex{}
		event := domain.SubtractMoneyEvent{
			UserId: "id",
			Amount: 10,
		}
		data, err := json.Marshal(event)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		mockrepo := new(mocks.IventHandlerUserRepository)
		mockReservermoney := new(mocks.ReserveMoneyHandler)
		assert.NotNil(t, mockrepo)
		mockrepo.On("SubtractMoney", event.UserId, event.Amount).Return(nil).Once()
		mockReservermoney.On("UnReserveMoney", mock.Anything, mock.Anything, mock.Anything).Times(1)
		handler := NewIncomingEventHandler(lock, mockrepo, nil, mockReservermoney)
		assert.NotNil(t, handler)
		err = handler.SubtractBalance(data)
		assert.NoError(t, err)
		mockrepo.AssertExpectations(t)
		mockReservermoney.AssertExpectations(t)
	})

	t.Run("fail on epo", func(t *testing.T) {
		lock := &sync.Mutex{}
		event := domain.SubtractMoneyEvent{
			UserId: "id",
			Amount: 10,
		}
		data, err := json.Marshal(event)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		mockrepo := new(mocks.IventHandlerUserRepository)
		mockReservermoney := new(mocks.ReserveMoneyHandler)
		assert.NotNil(t, mockrepo)
		mockrepo.On("SubtractMoney", event.UserId, event.Amount).Return(errors.New("")).Once()
		handler := NewIncomingEventHandler(lock, mockrepo, nil, mockReservermoney)
		assert.NotNil(t, handler)
		err = handler.SubtractBalance(data)
		assert.NotNil(t, err)
		mockrepo.AssertExpectations(t)
		mockReservermoney.AssertExpectations(t)
	})
}

func TestCheckMoney(t *testing.T) {
	t.Run("sucess", func(t *testing.T) {
		lock := &sync.Mutex{}
		event := domain.CheckMoneyEvent{
			UserId: "id",
			Amount: 10,
			Hash:   "",
		}
		data, err := json.Marshal(event)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		mockrepo := new(mocks.IventHandlerUserRepository)
		mockReservermoney := new(mocks.ReserveMoneyHandler)
		mockEvenManager := new(mocks.IventHandlerUserEventManager)
		assert.NotNil(t, mockrepo)
		assert.NotNil(t, mockReservermoney)
		assert.NotNil(t, mockEvenManager)
		mockEvenManager.On("Publish", mock.Anything, mock.Anything).Return(nil).Once()
		mockrepo.On("GetUserBalance", event.UserId).Return(float64(20), nil).Once()
		mockReservermoney.On("ReserveMoney", mock.Anything, mock.Anything, mock.Anything).Once()
		handler := NewIncomingEventHandler(lock, mockrepo, mockEvenManager, mockReservermoney)
		assert.NotNil(t, handler)
		err = handler.CheckMoney(data)
		assert.NoError(t, err)
		mockrepo.AssertExpectations(t)
		mockReservermoney.AssertExpectations(t)
		mockReservermoney.AssertExpectations(t)
	})

	t.Run("fail on epo", func(t *testing.T) {
		error1 := errors.New("")
		lock := &sync.Mutex{}
		event := domain.SubtractMoneyEvent{
			UserId: "id",
			Amount: 10,
		}
		data, err := json.Marshal(event)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		mockrepo := new(mocks.IventHandlerUserRepository)
		assert.NotNil(t, mockrepo)
		mockrepo.On("GetUserBalance", mock.Anything).Return(float64(1), error1).Once()
		handler := NewIncomingEventHandler(lock, mockrepo, nil, nil)
		assert.NotNil(t, handler)
		err = handler.CheckMoney(data)
		assert.NotNil(t, err)
		mockrepo.AssertExpectations(t)
	})
}
