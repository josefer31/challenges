// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	uuid "github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
)

// IdGenerator is an autogenerated mock type for the IdGenerator type
type IdGenerator struct {
	mock.Mock
}

type IdGenerator_Expecter struct {
	mock *mock.Mock
}

func (_m *IdGenerator) EXPECT() *IdGenerator_Expecter {
	return &IdGenerator_Expecter{mock: &_m.Mock}
}

// Next provides a mock function with given fields:
func (_m *IdGenerator) Next() uuid.UUID {
	ret := _m.Called()

	var r0 uuid.UUID
	if rf, ok := ret.Get(0).(func() uuid.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	return r0
}

// IdGenerator_Next_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Next'
type IdGenerator_Next_Call struct {
	*mock.Call
}

// Next is a helper method to define mock.On call
func (_e *IdGenerator_Expecter) Next() *IdGenerator_Next_Call {
	return &IdGenerator_Next_Call{Call: _e.mock.On("Next")}
}

func (_c *IdGenerator_Next_Call) Run(run func()) *IdGenerator_Next_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IdGenerator_Next_Call) Return(_a0 uuid.UUID) *IdGenerator_Next_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IdGenerator_Next_Call) RunAndReturn(run func() uuid.UUID) *IdGenerator_Next_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewIdGenerator interface {
	mock.TestingT
	Cleanup(func())
}

// NewIdGenerator creates a new instance of IdGenerator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIdGenerator(t mockConstructorTestingTNewIdGenerator) *IdGenerator {
	mock := &IdGenerator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
