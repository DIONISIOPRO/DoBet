package repository

import (
	"errors"
	"github.com/dionisiopro/dobet-user/domain"
	mocks "github.com/dionisiopro/dobet-user/mocks/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

var validUser = domain.User{
	User_id:         "id",
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

func TestCreatUser(t *testing.T) {
	t.Run("sucess", func(t *testing.T) {
		mockDriver := new(mocks.MongoDriverUserCollection)
		mockDriver.On("CountDocuments", mock.Anything, mock.Anything).Return(int64(0), nil)
		mockDriver.On("InsertOne", mock.Anything, mock.Anything).Return(nil, nil)
		repo := NewUserRepository(mockDriver)
		id, err := repo.CreateUser(validUser)
		assert.NoError(t, err)
		assert.NotEqual(t, "", id)
		mockDriver.AssertExpectations(t)
	})

	t.Run("user allready exist", func(t *testing.T) {
		mockDriver := new(mocks.MongoDriverUserCollection)
		mockDriver.On("CountDocuments", mock.Anything, mock.Anything).Return(int64(1), nil)
		repo := NewUserRepository(mockDriver)
		id, err := repo.CreateUser(validUser)
		assert.NotNil(t, err)
		assert.Equal(t, "", id)
		mockDriver.AssertExpectations(t)
	})

	t.Run("error while inserting ", func(t *testing.T) {
		mockDriver := new(mocks.MongoDriverUserCollection)
		mockDriver.On("CountDocuments", mock.Anything, mock.Anything).Return(int64(0), nil)
		mockDriver.On("InsertOne", mock.Anything, mock.Anything).Return(nil, errors.New("some error"))
		repo := NewUserRepository(mockDriver)
		id, err := repo.CreateUser(validUser)
		assert.NotNil(t, err)
		assert.Equal(t, "", id)
		assert.Equal(t, "some error", err.Error())
		mockDriver.AssertExpectations(t)
	})

}

func TestUpdateUser(t *testing.T) {
	t.Run("sucess", func(t *testing.T) {
		mockDriver := new(mocks.MongoDriverUserCollection)
		mockDriver.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
		repo := NewUserRepository(mockDriver)
		err := repo.UpdateUser(mock.Anything, validUser)
		assert.NoError(t, err)
		mockDriver.AssertExpectations(t)
	})

	t.Run("error while updating ", func(t *testing.T) {
		mockDriver := new(mocks.MongoDriverUserCollection)
		mockDriver.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("some error"))
		repo := NewUserRepository(mockDriver)
		err := repo.UpdateUser(mock.Anything, validUser)
		assert.NotNil(t, err)
		assert.Equal(t, "some error", err.Error())
		mockDriver.AssertExpectations(t)
	})

}
func TestFind(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		collection := mt.Coll
		repo := NewUserRepository(collection)
		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, prepareUserToSave(validUser))
		second := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, prepareUserToSave(validUser))
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)
		users, err := repo.GetUsers(int64(1), int64(9))
		assert.Nil(t, err)
		assert.IsType(t, []domain.User{}, users)
		assert.Equal(t, 2, len(users))
		assert.Equal(t, validUser.Phone_number, users[0].Phone_number)
	})

	mt.Run("error in mongo", func(mt *mtest.T) {
		collection := mt.Coll
		repo := NewUserRepository(collection)
		mt.AddMockResponses(nil)
		_, err := repo.GetUsers(int64(1), int64(9))
		assert.NotNil(t, err)
	})
}

func TestGetUserById(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		collection := mt.Coll
		dc := prepareUserToSave(validUser)
		repo := NewUserRepository(collection)
		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, dc)
		mt.AddMockResponses(first)
		user, err := repo.GetUserById(dc.Map()["user_id"].(string))
		assert.Nil(t, err)
		assert.IsType(t, domain.User{}, user)
		assert.Equal(t, validUser.Phone_number, user.Phone_number)
	})
	mt.Run("error in mongo", func(mt *mtest.T) {
		collection := mt.Coll
		repo := NewUserRepository(collection)
		mt.AddMockResponses(nil)
		dc := prepareUserToSave(validUser)
		_, err := repo.GetUserById(dc.Map()["user_id"].(string))
		assert.NotNil(t, err)
	})
}

func TestGetUserByPhone(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		collection := mt.Coll
		dc := prepareUserToSave(validUser)
		repo := NewUserRepository(collection)
		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, dc)
		mt.AddMockResponses(first)
		user, err := repo.GetUserByPhone(dc.Map()["phone_number"].(string))
		assert.Nil(t, err)
		assert.IsType(t, domain.User{}, user)
		assert.Equal(t, validUser.Phone_number, user.Phone_number)
	})
	mt.Run("error in mongo", func(mt *mtest.T) {
		collection := mt.Coll
		repo := NewUserRepository(collection)
		mt.AddMockResponses(nil)
		dc := prepareUserToSave(validUser)
		_, err := repo.GetUserByPhone(dc.Map()["phone_number"].(string))
		assert.NotNil(t, err)
	})
}

func TestGetUserBalance(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("success", func(mt *mtest.T) {
		collection := mt.Coll
		dc := prepareUserToSave(validUser)
		repo := NewUserRepository(collection)
		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, dc)
		mt.AddMockResponses(first)
		balance, err := repo.GetUserBalance(dc.Map()["user_id"].(string))
		assert.Nil(t, err)
		assert.Equal(t, validUser.Account_balance, balance)
	})

	mt.Run("error in mongo", func(mt *mtest.T) {
		collection := mt.Coll
		repo := NewUserRepository(collection)
		mt.AddMockResponses(nil)
		dc := prepareUserToSave(validUser)
		balance, err := repo.GetUserBalance(dc.Map()["phone_number"].(string))
		assert.NotNil(t, err)
		assert.Equal(t, float64(-1), balance)
	})

}

func TestDeleteUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		mockDriver := new(mocks.MongoDriverUserCollection)
		mockDriver.On("DeleteOne", mock.Anything, mock.Anything, mock.Anything).Return(nil, nil)
		repo := NewUserRepository(mockDriver)
		err := repo.DeleteUser(validUser.User_id)
		assert.Nil(t, err)
		mockDriver.AssertExpectations(t)
	})

	mt.Run("error deleting", func(mt *mtest.T) {
		mockDriver := new(mocks.MongoDriverUserCollection)
		mockDriver.On("DeleteOne", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("some"))
		repo := NewUserRepository(mockDriver)
		err := repo.DeleteUser(validUser.User_id)
		assert.NotNil(t, err)
		assert.Equal(t, "some", err.Error())
		mockDriver.AssertExpectations(t)
	})

	mt.Run("id emptly", func(mt *mtest.T) {
		repo := NewUserRepository(nil)
		err := repo.DeleteUser("")
		assert.NotNil(t, err)
	})
}

func TestAddMoneyUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		collection := mt.Coll
		dc := prepareUserToSave(validUser)
		repo := NewUserRepository(collection)
		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, dc)
		second := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, dc)
		thirty := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, dc)
		mt.AddMockResponses(first, second, thirty)
		err := repo.AddMoney(validUser.User_id, float64(10))
		assert.Nil(t, err)
	})
}

func TestSubtractMoneyUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		collection := mt.Coll
		dc := prepareUserToSave(validUser)
		repo := NewUserRepository(collection)
		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, dc)
		second := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, dc)
		thirty := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, dc)
		mt.AddMockResponses(first, second, thirty)
		err := repo.SubtractMoney(validUser.User_id, float64(10))
		assert.Nil(t, err)
	})
}
