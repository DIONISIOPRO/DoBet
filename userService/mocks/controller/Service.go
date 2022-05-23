// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/dionisiopro/dobet-user/domain"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// DeleteUser provides a mock function with given fields: userid
func (_m *Service) DeleteUser(userid string) error {
	ret := _m.Called(userid)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserById provides a mock function with given fields: userId
func (_m *Service) GetUserById(userId string) (domain.User, error) {
	ret := _m.Called(userId)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(string) domain.User); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByPhone provides a mock function with given fields: phone
func (_m *Service) GetUserByPhone(phone string) (domain.User, error) {
	ret := _m.Called(phone)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(string) domain.User); ok {
		r0 = rf(phone)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(phone)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: page, perpage
func (_m *Service) GetUsers(page int64, perpage int64) ([]domain.User, error) {
	ret := _m.Called(page, perpage)

	var r0 []domain.User
	if rf, ok := ret.Get(0).(func(int64, int64) []domain.User); ok {
		r0 = rf(page, perpage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64) error); ok {
		r1 = rf(page, perpage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: userid, user
func (_m *Service) UpdateUser(userid string, user domain.User) error {
	ret := _m.Called(userid, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, domain.User) error); ok {
		r0 = rf(userid, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
