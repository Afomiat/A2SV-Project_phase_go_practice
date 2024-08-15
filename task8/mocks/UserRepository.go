// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "clean_architecture_Testing/domain"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: user
func (_m *UserRepository) CreateUser(user domain.User) (domain.User, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(domain.User) (domain.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(domain.User) domain.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(domain.User) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUserID provides a mock function with given fields: username
func (_m *UserRepository) DeleteUserID(username string) (domain.User, error) {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUserID")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (domain.User, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) domain.User); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllUsers provides a mock function with given fields:
func (_m *UserRepository) GetAllUsers() ([]domain.User, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllUsers")
	}

	var r0 []domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]domain.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []domain.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMyProfile provides a mock function with given fields: username
func (_m *UserRepository) GetMyProfile(username string) (domain.User, error) {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for GetMyProfile")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (domain.User, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) domain.User); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: id
func (_m *UserRepository) GetUserByID(id string) (domain.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (domain.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) domain.User); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginUser provides a mock function with given fields: username, password
func (_m *UserRepository) LoginUser(username string, password string) (domain.User, error) {
	ret := _m.Called(username, password)

	if len(ret) == 0 {
		panic("no return value specified for LoginUser")
	}

	var r0 domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (domain.User, error)); ok {
		return rf(username, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) domain.User); ok {
		r0 = rf(username, password)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}