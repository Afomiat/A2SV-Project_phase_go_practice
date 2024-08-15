// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "clean_architecture_Testing/domain"

	mock "github.com/stretchr/testify/mock"
)

// ErrorResponseDomain is an autogenerated mock type for the ErrorResponseDomain type
type ErrorResponseDomain struct {
	mock.Mock
}

// NewErrorResponse provides a mock function with given fields: message
func (_m *ErrorResponseDomain) NewErrorResponse(message string) domain.ErrorResponse {
	ret := _m.Called(message)

	if len(ret) == 0 {
		panic("no return value specified for NewErrorResponse")
	}

	var r0 domain.ErrorResponse
	if rf, ok := ret.Get(0).(func(string) domain.ErrorResponse); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Get(0).(domain.ErrorResponse)
	}

	return r0
}

// NewErrorResponseDomain creates a new instance of ErrorResponseDomain. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewErrorResponseDomain(t interface {
	mock.TestingT
	Cleanup(func())
}) *ErrorResponseDomain {
	mock := &ErrorResponseDomain{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}