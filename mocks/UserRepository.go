// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"
	dtorepository "go-template/dto/general/repository"
	entity "go-template/entity"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, u
func (_m *UserRepository) Create(ctx context.Context, u entity.User) (entity.User, error) {
	ret := _m.Called(ctx, u)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(context.Context, entity.User) entity.User); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.User) error); ok {
		r1 = rf(ctx, u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUserForgetPassword provides a mock function with given fields: ctx, urp
func (_m *UserRepository) CreateUserForgetPassword(ctx context.Context, urp entity.UserResetPassword) (entity.UserResetPassword, error) {
	ret := _m.Called(ctx, urp)

	var r0 entity.UserResetPassword
	if rf, ok := ret.Get(0).(func(context.Context, entity.UserResetPassword) entity.UserResetPassword); ok {
		r0 = rf(ctx, urp)
	} else {
		r0 = ret.Get(0).(entity.UserResetPassword)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.UserResetPassword) error); ok {
		r1 = rf(ctx, urp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePreviousResetPassword provides a mock function with given fields: ctx, userId
func (_m *UserRepository) DeletePreviousResetPassword(ctx context.Context, userId int) error {
	ret := _m.Called(ctx, userId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUsedResetPassword provides a mock function with given fields: ctx, urp
func (_m *UserRepository) DeleteUsedResetPassword(ctx context.Context, urp entity.UserResetPassword) (entity.UserResetPassword, error) {
	ret := _m.Called(ctx, urp)

	var r0 entity.UserResetPassword
	if rf, ok := ret.Get(0).(func(context.Context, entity.UserResetPassword) entity.UserResetPassword); ok {
		r0 = rf(ctx, urp)
	} else {
		r0 = ret.Get(0).(entity.UserResetPassword)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.UserResetPassword) error); ok {
		r1 = rf(ctx, urp)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByEmail provides a mock function with given fields: ctx, email
func (_m *UserRepository) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	ret := _m.Called(ctx, email)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindById provides a mock function with given fields: ctx, userId
func (_m *UserRepository) FindById(ctx context.Context, userId int) (entity.User, error) {
	ret := _m.Called(ctx, userId)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(context.Context, int) entity.User); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserWalletAccountByUserId provides a mock function with given fields: ctx, userId
func (_m *UserRepository) FindUserWalletAccountByUserId(ctx context.Context, userId int) (dtorepository.UserWalletDataResponse, error) {
	ret := _m.Called(ctx, userId)

	var r0 dtorepository.UserWalletDataResponse
	if rf, ok := ret.Get(0).(func(context.Context, int) dtorepository.UserWalletDataResponse); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Get(0).(dtorepository.UserWalletDataResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResetPasswordTokenByToken provides a mock function with given fields: ctx, token
func (_m *UserRepository) GetResetPasswordTokenByToken(ctx context.Context, token string) (entity.UserResetPassword, error) {
	ret := _m.Called(ctx, token)

	var r0 entity.UserResetPassword
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.UserResetPassword); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(entity.UserResetPassword)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, u
func (_m *UserRepository) UpdateUser(ctx context.Context, u entity.User) (entity.User, error) {
	ret := _m.Called(ctx, u)

	var r0 entity.User
	if rf, ok := ret.Get(0).(func(context.Context, entity.User) entity.User); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Get(0).(entity.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, entity.User) error); ok {
		r1 = rf(ctx, u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
