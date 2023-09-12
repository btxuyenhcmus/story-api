// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "readtruyen-api/domain"
	mock "github.com/stretchr/testify/mock"
)

// LoginUsecase is an autogenerated mock type for the LoginUsecase type
type LoginUsecase struct {
	mock.Mock
}

// CreateAccessToken provides a mock function with given fields: user, secret, expiry
func (_m *LoginUsecase) CreateAccessToken(user *domain.User, secret string, expiry int) (string, error) {
	ret := _m.Called(user, secret, expiry)

	var r0 string
	if rf, ok := ret.Get(0).(func(*domain.User, string, int) string); ok {
		r0 = rf(user, secret, expiry)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.User, string, int) error); ok {
		r1 = rf(user, secret, expiry)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRefreshToken provides a mock function with given fields: user, secret, expiry
func (_m *LoginUsecase) CreateRefreshToken(user *domain.User, secret string, expiry int) (string, error) {
	ret := _m.Called(user, secret, expiry)

	var r0 string
	if rf, ok := ret.Get(0).(func(*domain.User, string, int) string); ok {
		r0 = rf(user, secret, expiry)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*domain.User, string, int) error); ok {
		r1 = rf(user, secret, expiry)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: c, email
func (_m *LoginUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ret := _m.Called(c, email)

	var r0 domain.User
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.User); ok {
		r0 = rf(c, email)
	} else {
		r0 = ret.Get(0).(domain.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewLoginUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewLoginUsecase creates a new instance of LoginUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLoginUsecase(t mockConstructorTestingTNewLoginUsecase) *LoginUsecase {
	mock := &LoginUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}