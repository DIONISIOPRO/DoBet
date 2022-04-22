// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/namuethopro/dobet-auth/domain"
	mock "github.com/stretchr/testify/mock"
)

// AuthService is an autogenerated mock type for the AuthService type
type AuthService struct {
	mock.Mock
}

// Login provides a mock function with given fields: _a0
func (_m *AuthService) Login(_a0 domain.LoginDetails) (string, string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(domain.LoginDetails) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(domain.LoginDetails) string); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(domain.LoginDetails) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Logout provides a mock function with given fields: token
func (_m *AuthService) Logout(token string) error {
	ret := _m.Called(token)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RefreshToken provides a mock function with given fields: token
func (_m *AuthService) RefreshToken(token string) (string, string, error) {
	ret := _m.Called(token)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string) string); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(token)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
