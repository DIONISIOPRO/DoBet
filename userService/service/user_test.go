package service

import (
	"errors"
	"github/namuethopro/dobet-user/domain"
	mocks "github/namuethopro/dobet-user/mocks/service"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var validUser = domain.User{
	User_id:         "",
	Phone_number:    "123456789",
	First_name:      "Dio",
	Last_name:       "paulo",
	Password:        "12334",
	Account_balance: 0,
	Created_at:      time.Now(),
	Updated_at:      time.Now(),
	IsAdmin:         false,
	RefreshTokens:   []string{},
}
var users = []domain.User{
	validUser,
}

func TestCreateUser(t *testing.T) {
	t.Run("fail Invalid user", func(t *testing.T) {
		var userRepoMock = new(mocks.UserRepository)
		invaliduser1 := validUser
		invaliduser1.Phone_number = "25468"
		userService := newUserService(nil, nil, nil, nil, nil)
		name, err := userService.CreateUser(invaliduser1)
		assert.NotNil(t, err)
		assert.Equal(t, "", name)
		assert.Equal(t, "the lenght of number should be 9", err.Error())
		invaliduser1.Phone_number = "d4d455555"
		name, err = userService.CreateUser(invaliduser1)
		assert.NotNil(t, err)
		assert.Equal(t, "", name)
		assert.Equal(t, "your number is not valid", err.Error())
		userRepoMock.AssertExpectations(t)
	})

	t.Run("sucess", func(t *testing.T) {
		var userRepoMock = new(mocks.UserRepository)
		var userEvManagerMock = new(mocks.EventPublisher)
		userRepoMock.On("CreateUser", validUser).Return("name", nil).Once()
		userEvManagerMock.On("Publish", mock.Anything, mock.Anything).Return(nil).Once()
		userService := newUserService(userRepoMock, userEvManagerMock, nil, nil, nil)
		name, err := userService.CreateUser(validUser)
		assert.NoError(t, err)
		assert.Equal(t, "name", name)
		userRepoMock.AssertExpectations(t)
		userEvManagerMock.AssertExpectations(t)
	})

	t.Run("fail in Repo", func(t *testing.T) {
		var userRepoMock = new(mocks.UserRepository)
		errRepo := errors.New("repo err")
		userRepoMock.On("CreateUser", validUser).Return("name", errRepo).Once()
		userService := newUserService(userRepoMock, nil, nil, nil, nil)
		name, err := userService.CreateUser(validUser)
		assert.NotNil(t, err)
		assert.Equal(t, "", name)
		userRepoMock.AssertExpectations(t)
	})

	t.Run("fail in eventmanager", func(t *testing.T) {
		var userRepoMock = new(mocks.UserRepository)
		var userEvManagerMock = new(mocks.EventPublisher)
		errevent := errors.New("event err")
		userRepoMock.On("CreateUser", validUser).Return("name", nil).Once()
		userEvManagerMock.On("Publish", mock.Anything, mock.Anything).Return(errevent).Once()
		userService := newUserService(userRepoMock, userEvManagerMock, nil, nil, nil)
		name, err := userService.CreateUser(validUser)
		assert.NotNil(t, err)
		assert.Equal(t, "name", name)
		userRepoMock.AssertExpectations(t)
		userEvManagerMock.AssertExpectations(t)
	})
}

func TestGetUsers(t *testing.T) {
	t.Run("sucess", func(t *testing.T) {
		var userRepoMock = new(mocks.UserRepository)
		page, perpage := int64(1), int64(9)
		startIndex := (page - 1) * perpage
		userRepoMock.On("GetUsers", startIndex, perpage).Return(users, nil).Once()
		userService := newUserService(userRepoMock, nil, nil, nil, nil)
		allUsers, err := userService.GetUsers(int64(page), int64(perpage))
		assert.NoError(t, err)
		assert.Equal(t, allUsers, users)
		userRepoMock.AssertExpectations(t)
	})

	t.Run("sucess whit page < 1", func(t *testing.T) {
		var userRepoMock = new(mocks.UserRepository)
		page, perpage := int64(0), int64(5)
		userRepoMock.On("GetUsers", int64(0), perpage).Return(users, nil).Once()
		userService := newUserService(userRepoMock, nil, nil, nil, nil)
		allUsers, err := userService.GetUsers(int64(page), int64(perpage))
		assert.NoError(t, err)
		assert.Equal(t, allUsers, users)
		userRepoMock.AssertExpectations(t)
	})

	t.Run("sucess whit perpage < 1", func(t *testing.T) {
		var userRepoMock = new(mocks.UserRepository)
		page, perpage := int64(0), int64(0)
		userRepoMock.On("GetUsers", int64(0), int64(9)).Return(users, nil).Once()
		userService := newUserService(userRepoMock, nil, nil, nil, nil)
		allUsers, err := userService.GetUsers(int64(page), int64(perpage))
		assert.NoError(t, err)
		assert.Equal(t, allUsers, users)
		userRepoMock.AssertExpectations(t)
	})

	t.Run("fail from repo", func(t *testing.T) {
		var userRepoMock = new(mocks.UserRepository)
		page, perpage := int64(1), int64(9)
		startIndex := (page - 1) * perpage
		userRepoMock.On("GetUsers", startIndex, perpage).Return(nil, errors.New("")).Once()
		userService := newUserService(userRepoMock, nil, nil, nil, nil)
		_, err := userService.GetUsers(int64(page), int64(perpage))
		assert.NotNil(t, err)
		userRepoMock.AssertExpectations(t)
	})

}

func TestGetUserById(t *testing.T) {
	t.Run("sucess", func(t *testing.T) {
		var userRepoMock = new(mocks.UserRepository)
		userRepoMock.On("GetUserById", "id").Return(validUser, nil).Once()
		userService := newUserService(userRepoMock, nil, nil, nil, nil)
		user, err := userService.GetUserById("id")
		assert.NoError(t, err)
		assert.Equal(t, user, validUser)
		userRepoMock.AssertExpectations(t)
	})

	t.Run("fail from repo", func(t *testing.T) {
		var userRepoMock = new(mocks.UserRepository)
		userRepoMock.On("GetUserById", "id").Return(domain.User{}, errors.New("")).Once()
		userService := newUserService(userRepoMock, nil, nil, nil, nil)
		_, err := userService.GetUserById("id")
		assert.NotNil(t, err)
		userRepoMock.AssertExpectations(t)
	})

	t.Run("fail from id empty", func(t *testing.T) {
		var userRepoMock = new(mocks.UserRepository)
		var userService = newUserService(nil, nil, nil, nil, nil)
		_, err := userService.GetUserById("")
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "user id is empty")
		userRepoMock.AssertExpectations(t)
	})
}

func TestGetUserByPhone(t *testing.T) {
	t.Run("sucess", func(t *testing.T) {
		var userRepoMock = new(mocks.UserRepository)
		userRepoMock.On("GetUserByPhone", "123456789").Return(validUser, nil).Once()
		userService := newUserService(userRepoMock, nil, nil, nil, nil)
		user, err := userService.GetUserByPhone("123456789")
		assert.NoError(t, err)
		assert.Equal(t, user, validUser)
		userRepoMock.AssertExpectations(t)
	})

	t.Run("fail from repo", func(t *testing.T) {
		var userRepoMock = new(mocks.UserRepository)
		userRepoMock.On("GetUserByPhone", "123456789").Return(domain.User{}, errors.New("")).Once()
		userService := newUserService(userRepoMock, nil, nil, nil, nil)
		_, err := userService.GetUserByPhone("123456789")
		assert.NotNil(t, err)
		userRepoMock.AssertExpectations(t)
	})

	t.Run("fail from id empty", func(t *testing.T) {
		var userRepoMock = new(mocks.UserRepository)
		userService := newUserService(nil, nil, nil, nil, nil)
		_, err := userService.GetUserByPhone("")
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "user id is empty")
		userRepoMock.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("sucess", func(t *testing.T) {
		var userRepoMock = new(mocks.UserRepository)
		var userEvManagerMock = new(mocks.EventPublisher)
		userRepoMock.On("DeleteUser", "id").Return(nil).Once()
		userEvManagerMock.On("Publish", mock.Anything, mock.Anything).Return(nil).Once()
		userService := newUserService(userRepoMock, userEvManagerMock, nil, nil, nil)
		err := userService.DeleteUser("id")
		assert.NoError(t, err)
		userRepoMock.AssertExpectations(t)
		userEvManagerMock.AssertExpectations(t)
	})

	t.Run("fail from repo", func(t *testing.T) {
		var userRepoMock = new(mocks.UserRepository)
		err1 := errors.New("some error")
		userRepoMock.On("DeleteUser", "id").Return(err1).Once()
		userService := newUserService(userRepoMock, nil, nil, nil, nil)
		err := userService.DeleteUser("id")
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "some error")
		userRepoMock.AssertExpectations(t)
	})

	t.Run("fail from id empty", func(t *testing.T) {
		var userRepoMock = new(mocks.UserRepository)
		var userService = newUserService(nil, nil, nil, nil, nil)
		err := userService.DeleteUser("")
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "user id is empty")
		userRepoMock.AssertExpectations(t)
	})

	t.Run("fail from event", func(t *testing.T) {
		var userRepoMock = new(mocks.UserRepository)
		var userEvManagerMock = new(mocks.EventPublisher)
		err1 := errors.New("some error")
		userRepoMock.On("DeleteUser", "id").Return(nil).Once()
		userEvManagerMock.On("Publish", mock.Anything, mock.Anything).Return(err1).Once()
		userServicee := newUserService(userRepoMock, userEvManagerMock, nil, nil, nil)
		err := userServicee.DeleteUser("id")
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "some error")
		userRepoMock.AssertExpectations(t)
		userEvManagerMock.AssertExpectations(t)
	})

}

func TestUpdateUser(t *testing.T) {
	t.Run("sucess", func(t *testing.T) {
		var userRepoMock = new(mocks.UserRepository)
		var userEvManagerMock = new(mocks.EventPublisher)
		userRepoMock.On("UpdateUser", "id", validUser).Return(nil).Once()
		userEvManagerMock.On("Publish", mock.Anything, mock.Anything).Return(nil).Once()
		userService := newUserService(userRepoMock, userEvManagerMock, nil, nil, nil)
		err := userService.UpdateUser("id", validUser)
		assert.NoError(t, err)
		userRepoMock.AssertExpectations(t)
		userEvManagerMock.AssertExpectations(t)
	})

	t.Run("fail from repo", func(t *testing.T) {
		var userRepoMock = new(mocks.UserRepository)
		err1 := errors.New("some error")
		userRepoMock.On("UpdateUser", "id", validUser).Return(err1).Once()
		userService := newUserService(userRepoMock, nil, nil, nil, nil)
		err := userService.UpdateUser("id", validUser)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "some error")
		userRepoMock.AssertExpectations(t)
	})

	t.Run("fail from id empty", func(t *testing.T) {
		userService := newUserService(nil, nil, nil, nil, nil)
		err := userService.UpdateUser("", validUser)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "user id is empty")
	})

	t.Run("fail from event", func(t *testing.T) {
		var userRepoMock = new(mocks.UserRepository)
		var userEvManagerMock = new(mocks.EventPublisher)
		err1 := errors.New("some error")
		userRepoMock.On("UpdateUser", "id", validUser).Return(nil).Once()
		userEvManagerMock.On("Publish", mock.Anything, mock.Anything).Return(err1).Once()
		userServicee := newUserService(userRepoMock, userEvManagerMock, nil, nil, nil)
		err := userServicee.UpdateUser("id", validUser)
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "some error")
		userRepoMock.AssertExpectations(t)
		userEvManagerMock.AssertExpectations(t)
	})
}
