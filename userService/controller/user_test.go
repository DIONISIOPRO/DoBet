package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"github/namuethopro/dobet-user/domain"
	mocks "github/namuethopro/dobet-user/mocks/controller"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockUser = domain.User{
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

func TestGetUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("sucess", func(t *testing.T) {
		mockUsers := []domain.User{}
		mockUsers = append(mockUsers, mockUser)
		service := new(mocks.UserService)
		service.On("GetUsers", mock.AnythingOfType("int64"), mock.AnythingOfType("int64")).Return(mockUsers, nil)
		req, err := http.NewRequest("GET", "/api/v1/user", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(res)
		c.Request = req
		controller := NewUserController(service)
		controller.GetUsers(c)
		code := res.Result().StatusCode
		err = json.Unmarshal(res.Body.Bytes(), &mockUsers)
		assert.NoError(t, err)
		assert.Equal(t, 200, code)
		service.AssertExpectations(t)
	})

	t.Run("error in service", func(t *testing.T) {
		mockUsers := []domain.User{}
		mockUsers = append(mockUsers, mockUser)
		service := new(mocks.UserService)
		service.On("GetUsers", mock.AnythingOfType("int64"), mock.AnythingOfType("int64")).Return(nil, errors.New(""))
		req, err := http.NewRequest("GET", "/api/v1/user", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(res)
		c.Request = req
		controller := NewUserController(service)
		controller.GetUsers(c)
		code := res.Result().StatusCode
		err = json.Unmarshal(res.Body.Bytes(), &mockUsers)
		assert.NotNil(t, err)
		assert.Equal(t, 500, code)
		service.AssertExpectations(t)
	})

}

func TestGetUsersById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("sucess", func(t *testing.T) {
		service := new(mocks.UserService)
		service.On("GetUserById", "5").Return(mockUser, nil)
		req, err := http.NewRequest("GET", "/api/v1/user/5", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(res)
		p := gin.Param{
			Key:   "id",
			Value: "5",
		}
		c.Params = append(c.Params, p)
		c.Request = req
		controller := NewUserController(service)
		controller.GetUserById(c)
		code := res.Result().StatusCode
		err = json.Unmarshal(res.Body.Bytes(), &mockUser)
		assert.NoError(t, err)
		assert.Equal(t, 200, code)
		service.AssertExpectations(t)
	})

	t.Run("error in service", func(t *testing.T) {
		service := new(mocks.UserService)
		service.On("GetUserById", "5").Return(domain.User{}, errors.New(""))
		req, err := http.NewRequest("GET", "/api/v1/user", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(res)
		p := gin.Param{
			Key:   "id",
			Value: "5",
		}
		c.Params = append(c.Params, p)
		c.Request = req
		controller := NewUserController(service)
		controller.GetUserById(c)
		code := res.Result().StatusCode
		assert.Equal(t, 500, code)
		service.AssertExpectations(t)
	})

	t.Run("error: invalid id param", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/api/v1/user", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(res)
		c.Request = req
		controller := NewUserController(nil)
		controller.GetUserById(c)
		code := res.Result().StatusCode
		assert.Equal(t, http.StatusBadRequest, code)
	})

}

func TestGetUsersByPhone(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("sucess", func(t *testing.T) {
		localmockuser := mockUser
		service := new(mocks.UserService)
		service.On("GetUserByPhone", "123456789").Return(mockUser, nil)
		req, err := http.NewRequest("GET", "/api/v1/user/", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(res)
		p := gin.Param{
			Key:   "phone",
			Value: "123456789",
		}
		c.Params = append(c.Params, p)
		c.Request = req
		controller := NewUserController(service)
		controller.GetUserByPhone(c)
		code := res.Result().StatusCode
		err = json.Unmarshal(res.Body.Bytes(), &localmockuser)
		assert.NoError(t, err)
		assert.Equal(t, 200, code)
		service.AssertExpectations(t)
	})

	t.Run("error in service", func(t *testing.T) {
		service := new(mocks.UserService)
		service.On("GetUserByPhone", "123456789").Return(domain.User{}, errors.New(""))
		req, err := http.NewRequest("GET", "/api/v1/user", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(res)
		p := gin.Param{
			Key:   "phone",
			Value: "123456789",
		}
		c.Params = append(c.Params, p)
		c.Request = req
		controller := NewUserController(service)
		controller.GetUserByPhone(c)
		code := res.Result().StatusCode
		assert.Equal(t, 500, code)
		service.AssertExpectations(t)
	})

	t.Run("error: invalid id param", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/api/v1/user", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(res)
		c.Request = req
		controller := NewUserController(nil)
		controller.GetUserByPhone(c)
		code := res.Result().StatusCode
		assert.Equal(t, http.StatusBadRequest, code)
	})

}
func TestDeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	t.Run("sucess", func(t *testing.T) {
		service := new(mocks.UserService)
		service.On("DeleteUser", "5").Return(nil)
		req, err := http.NewRequest("POST", "/api/v1/user/delete", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(res)
		p := gin.Param{
			Key:   "id",
			Value: "5",
		}
		c.Params = append(c.Params, p)
		c.Request = req
		controller := NewUserController(service)
		controller.DeleteUser(c)
		code := res.Result().StatusCode
		assert.Equal(t, 200, code)
		service.AssertExpectations(t)
	})

	t.Run("error in service", func(t *testing.T) {
		service := new(mocks.UserService)
		service.On("DeleteUser", "5").Return(errors.New(""))
		req, err := http.NewRequest("GET", "/api/v1/user", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(res)
		p := gin.Param{
			Key:   "id",
			Value: "5",
		}
		c.Params = append(c.Params, p)
		c.Request = req
		controller := NewUserController(service)
		controller.DeleteUser(c)
		code := res.Result().StatusCode
		assert.Equal(t, 500, code)
		service.AssertExpectations(t)
	})

	t.Run("error: invalid id param", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/api/v1/user", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(res)
		c.Request = req
		controller := NewUserController(nil)
		controller.DeleteUser(c)
		code := res.Result().StatusCode
		assert.Equal(t, http.StatusBadRequest, code)
	})

}

func TestUpdateUser(t *testing.T) {
	localmockuser := mockUser
	gin.SetMode(gin.TestMode)
	t.Run("sucess", func(t *testing.T) {

		service := new(mocks.UserService)
		service.On("UpdateUser", mock.Anything, mock.Anything).Return(nil)
		var buf = &bytes.Buffer{}
		err := json.NewEncoder(buf).Encode(localmockuser)
		assert.NoError(t, err)
		req, err := http.NewRequest("POST", "/api/v1/user/update", buf)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(res)
		p := gin.Param{
			Key:   "id",
			Value: "5",
		}
		c.Params = append(c.Params, p)
		c.Request = req
		controller := NewUserController(service)
		controller.UpdateUser(c)
		code := res.Result().StatusCode
		assert.Equal(t, 200, code)
		service.AssertExpectations(t)
	})

	t.Run("error in service", func(t *testing.T) {
		service := new(mocks.UserService)
		service.On("UpdateUser", "5", mock.Anything).Return(errors.New(""))
		var buf = &bytes.Buffer{}
		err := json.NewEncoder(buf).Encode(localmockuser)
		assert.NoError(t, err)
		req, err := http.NewRequest("POST", "/api/v1/user/update", buf)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(res)
		p := gin.Param{
			Key:   "id",
			Value: "5",
		}
		c.Params = append(c.Params, p)
		c.Request = req
		controller := NewUserController(service)
		controller.UpdateUser(c)
		code := res.Result().StatusCode
		assert.Equal(t, 500, code)
		service.AssertExpectations(t)
	})

	t.Run("error: invalid id param", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/api/v1/user/update", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(res)
		c.Request = req
		controller := NewUserController(nil)
		controller.UpdateUser(c)
		code := res.Result().StatusCode
		assert.Equal(t, http.StatusBadRequest, code)
	})
}
