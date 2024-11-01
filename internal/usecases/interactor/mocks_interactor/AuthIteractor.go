// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks_interactor

import (
	context "context"

	entity "github.com/noffrialdi/auth/internal/usecases/entity"

	mock "github.com/stretchr/testify/mock"

	request "github.com/noffrialdi/auth/internal/handler/api/request"
)

// AuthIteractor is an autogenerated mock type for the AuthIteractor type
type AuthIteractor struct {
	mock.Mock
}

type AuthIteractor_Expecter struct {
	mock *mock.Mock
}

func (_m *AuthIteractor) EXPECT() *AuthIteractor_Expecter {
	return &AuthIteractor_Expecter{mock: &_m.Mock}
}

// Signup provides a mock function with given fields: ctx, req
func (_m *AuthIteractor) Signup(ctx context.Context, req *request.SignUpRequest) (*entity.SignUpResponse, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for Signup")
	}

	var r0 *entity.SignUpResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *request.SignUpRequest) (*entity.SignUpResponse, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *request.SignUpRequest) *entity.SignUpResponse); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.SignUpResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *request.SignUpRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AuthIteractor_Signup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Signup'
type AuthIteractor_Signup_Call struct {
	*mock.Call
}

// Signup is a helper method to define mock.On call
//   - ctx context.Context
//   - req *request.SignUpRequest
func (_e *AuthIteractor_Expecter) Signup(ctx interface{}, req interface{}) *AuthIteractor_Signup_Call {
	return &AuthIteractor_Signup_Call{Call: _e.mock.On("Signup", ctx, req)}
}

func (_c *AuthIteractor_Signup_Call) Run(run func(ctx context.Context, req *request.SignUpRequest)) *AuthIteractor_Signup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*request.SignUpRequest))
	})
	return _c
}

func (_c *AuthIteractor_Signup_Call) Return(respone *entity.SignUpResponse, err error) *AuthIteractor_Signup_Call {
	_c.Call.Return(respone, err)
	return _c
}

func (_c *AuthIteractor_Signup_Call) RunAndReturn(run func(context.Context, *request.SignUpRequest) (*entity.SignUpResponse, error)) *AuthIteractor_Signup_Call {
	_c.Call.Return(run)
	return _c
}

// NewAuthIteractor creates a new instance of AuthIteractor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthIteractor(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthIteractor {
	mock := &AuthIteractor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
