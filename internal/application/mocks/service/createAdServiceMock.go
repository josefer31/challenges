// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	service "polaris/internal/application/service"

	mock "github.com/stretchr/testify/mock"
)

// CreateAdService is an autogenerated mock type for the CreateAdService type
type CreateAdService struct {
	mock.Mock
}

type CreateAdService_Expecter struct {
	mock *mock.Mock
}

func (_m *CreateAdService) EXPECT() *CreateAdService_Expecter {
	return &CreateAdService_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: request
func (_m *CreateAdService) Execute(request service.CreateAdRequest) service.CreateAdResponse {
	ret := _m.Called(request)

	var r0 service.CreateAdResponse
	if rf, ok := ret.Get(0).(func(service.CreateAdRequest) service.CreateAdResponse); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(service.CreateAdResponse)
	}

	return r0
}

// CreateAdService_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type CreateAdService_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - request service.CreateAdRequest
func (_e *CreateAdService_Expecter) Execute(request interface{}) *CreateAdService_Execute_Call {
	return &CreateAdService_Execute_Call{Call: _e.mock.On("Execute", request)}
}

func (_c *CreateAdService_Execute_Call) Run(run func(request service.CreateAdRequest)) *CreateAdService_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(service.CreateAdRequest))
	})
	return _c
}

func (_c *CreateAdService_Execute_Call) Return(_a0 service.CreateAdResponse) *CreateAdService_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CreateAdService_Execute_Call) RunAndReturn(run func(service.CreateAdRequest) service.CreateAdResponse) *CreateAdService_Execute_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewCreateAdService interface {
	mock.TestingT
	Cleanup(func())
}

// NewCreateAdService creates a new instance of CreateAdService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCreateAdService(t mockConstructorTestingTNewCreateAdService) *CreateAdService {
	mock := &CreateAdService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
