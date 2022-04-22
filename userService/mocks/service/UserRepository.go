// Code generated by mockery v2.10.4. DO NOT EDIT.

package mocks

import (
	domain "github.com/namuethopro/dobet-user/domain"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// AddMoney provides a mock function with given fields: userId, amount
func (_m *UserRepository) AddMoney(userId string, amount float64) error {
	ret := _m.Called(userId, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, float64) error); ok {
		r0 = rf(userId, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUser provides a mock function with given fields: userRequest
func (_m *UserRepository) CreateUser(userRequest domain.User) (string, error) {
	ret := _m.Called(userRequest)

	var r0 string
	if rf, ok := ret.Get(0).(func(domain.User) string); ok {
		r0 = rf(userRequest)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.User) error); ok {
		r1 = rf(userRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: userid
func (_m *UserRepository) DeleteUser(userid string) error {
	ret := _m.Called(userid)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUserBalance provides a mock function with given fields: userId
func (_m *UserRepository) GetUserBalance(userId string) (float64, error) {
	ret := _m.Called(userId)

	var r0 float64
	if rf, ok := ret.Get(0).(func(string) float64); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserById provides a mock function with given fields: userId
func (_m *UserRepository) GetUserById(userId string) (domain.User, error) {
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
func (_m *UserRepository) GetUserByPhone(phone string) (domain.User, error) {
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

// GetUsers provides a mock function with given fields: startIndex, perpage
func (_m *UserRepository) GetUsers(startIndex int64, perpage int64) ([]domain.User, error) {
	ret := _m.Called(startIndex, perpage)

	var r0 []domain.User
	if rf, ok := ret.Get(0).(func(int64, int64) []domain.User); ok {
		r0 = rf(startIndex, perpage)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64) error); ok {
		r1 = rf(startIndex, perpage)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubtractMoney provides a mock function with given fields: userId, amount
func (_m *UserRepository) SubtractMoney(userId string, amount float64) error {
	ret := _m.Called(userId, amount)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, float64) error); ok {
		r0 = rf(userId, amount)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: userid, user
func (_m *UserRepository) UpdateUser(userid string, user domain.User) error {
	ret := _m.Called(userid, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, domain.User) error); ok {
		r0 = rf(userid, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
