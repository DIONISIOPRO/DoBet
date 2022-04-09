package routes

import (
	mocks "github/namuethopro/dobet-user/mocks/routes"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSetUpUserRoutes(t *testing.T) {
	engine := gin.New()
	routeInfo := engine.Routes()
	assert.Equal(t, 0, len(routeInfo))
	mockUserController := new(mocks.UserController)
	mockUserMiddleware := new(mocks.UsermiddleWare)
	router := NewUserRouter(mockUserController, mockUserMiddleware)
	mockUserMiddleware.On("Authenticated").Return(nil)
	mockUserMiddleware.On("IsAdmin").Return(nil)
	mockUserMiddleware.On("IsOwner").Return(nil)
	mockUserController.On("DeleteUser").Return(nil).Once()
	mockUserController.On("GetUserById").Return(nil).Once()
	mockUserController.On("GetUsers").Return(nil).Once()
	mockUserController.On("UpdateUser").Return(nil).Once()
	router.SetupUserRouter(engine)
	routeInfo = engine.Routes()
	assert.Equal(t, 5, len(routeInfo))
	mockUserController.AssertExpectations(t)
	mockUserMiddleware.AssertExpectations(t)
}
