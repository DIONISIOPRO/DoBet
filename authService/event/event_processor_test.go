package event

import (
	"encoding/json"
	"errors"
	"github.com/namuethopro/dobet-auth/domain"
	mocks "github.com/namuethopro/dobet-auth/mocks/event"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	t.Run("sucess", func(t *testing.T) {
		user := domain.User{
			User_id: "id",
		}
		event := domain.AddUserEvent{User: user}
		data, err := json.Marshal(event)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		mockrepo := new(mocks.IncomingEventProcessorRepository)
		assert.NotNil(t, mockrepo)
		mockrepo.On("AddUser", user).Return(nil).Once()
		handler := NewIncomingEventProcessor(mockrepo)
		assert.NotNil(t, handler)
		err = handler.AddUser(data)
		assert.NoError(t, err)
		mockrepo.AssertExpectations(t)
	})

	t.Run("fail on epo", func(t *testing.T) {
		user := domain.User{
			User_id: "id",
		}
		event := domain.AddUserEvent{User: user}
		data, err := json.Marshal(event)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		mockrepo := new(mocks.IncomingEventProcessorRepository)
		assert.NotNil(t, mockrepo)
		err1 := errors.New("some error")
		mockrepo.On("AddUser", user).Return(err1).Once()
		handler := NewIncomingEventProcessor(mockrepo)
		assert.NotNil(t, handler)
		err = handler.AddUser(data)
		assert.NotNil(t, err)
		mockrepo.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	t.Run("sucess", func(t *testing.T) {
		user := domain.User{
			User_id: "123",
		}
		event := domain.UpdateUserEvent{User: user, UserId: "123"}
		data, err := json.Marshal(event)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		mockrepo := new(mocks.IncomingEventProcessorRepository)
		assert.NotNil(t, mockrepo)
		mockrepo.On("UpdateUser", user.User_id, user).Return(nil).Once()
		handler := NewIncomingEventProcessor(mockrepo)
		assert.NotNil(t, handler)
		err = handler.UpdateUser(data)
		assert.NoError(t, err)
		mockrepo.AssertExpectations(t)
	})

	t.Run("fail on repo", func(t *testing.T) {
		user := domain.User{
			User_id: "123",
		}
		event := domain.UpdateUserEvent{User: user, UserId: "123"}
		data, err := json.Marshal(event)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		mockrepo := new(mocks.IncomingEventProcessorRepository)
		assert.NotNil(t, mockrepo)
		mockrepo.On("UpdateUser",  user.User_id, user).Return(errors.New("some error")).Once()
		handler := NewIncomingEventProcessor(mockrepo)
		assert.NotNil(t, handler)
		err = handler.UpdateUser(data)
		assert.NotNil(t, err)
		mockrepo.AssertExpectations(t)
	})
}

func TestRemoveUser(t *testing.T) {
	t.Run("sucess", func(t *testing.T) {
		event := domain.DeleteUserEvent{UserId: "id"}
		data, err := json.Marshal(event)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		mockrepo := new(mocks.IncomingEventProcessorRepository)
		assert.NotNil(t, mockrepo)
		mockrepo.On("RemoveUser", event.UserId).Return(nil).Once()
		handler := NewIncomingEventProcessor(mockrepo)
		assert.NotNil(t, handler)
		err = handler.RemoveUser(data)
		assert.NoError(t, err)
		mockrepo.AssertExpectations(t)
	})

	t.Run("fail on epo", func(t *testing.T) {
		event := domain.DeleteUserEvent{UserId: "id"}
		data, err := json.Marshal(event)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		mockrepo := new(mocks.IncomingEventProcessorRepository)
		assert.NotNil(t, mockrepo)
		mockrepo.On("RemoveUser", event.UserId).Return(errors.New("some error")).Once()
		handler := NewIncomingEventProcessor(mockrepo)
		assert.NotNil(t, handler)
		err = handler.RemoveUser(data)
		assert.NotNil(t, err)
		mockrepo.AssertExpectations(t)
	})
}
