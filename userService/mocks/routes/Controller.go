// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// DeleteUser provides a mock function with given fields: c
func (_m *Controller) DeleteUser(c *gin.Context) {
	_m.Called(c)
}

// GetUserById provides a mock function with given fields: c
func (_m *Controller) GetUserById(c *gin.Context) {
	_m.Called(c)
}

// GetUsers provides a mock function with given fields: c
func (_m *Controller) GetUsers(c *gin.Context) {
	_m.Called(c)
}

// UpdateUser provides a mock function with given fields: c
func (_m *Controller) UpdateUser(c *gin.Context) {
	_m.Called(c)
}
