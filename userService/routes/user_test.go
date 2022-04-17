package routes

import (
	mocks "github/namuethopro/dobet-user/mocks/routes"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSetUpUserRoutes(t *testing.T) {
	engine := gin.New()
	routeInfo := engine.Routes()
	assert.Equal(t, 0, len(routeInfo))
	mockUserController := new(mocks.Controller)
	mockUserMiddleware := new(mocks.MiddleWare)
	router := NewRouter(mockUserController, mockUserMiddleware)
	mockUserMiddleware.On("Authenticated").Return(nil)
	mockUserMiddleware.On("IsAdmin").Return(nil)
	mockUserMiddleware.On("IsOwner").Return(nil)
	mockUserController.On("DeleteUser", mock.Anything).Once()
	mockUserController.On("GetUserById",mock.Anything).Once()
	mockUserController.On("GetUsers", mock.Anything).Once()
	mockUserController.On("UpdateUser", mock.AnythingOfType("*gin.Context")).Once()
	router.SetupUserRouter(engine)
	routeInfo = engine.Routes()
	assert.Equal(t, 6, len(routeInfo))
	mockUserMiddleware.AssertExpectations(t)
}
