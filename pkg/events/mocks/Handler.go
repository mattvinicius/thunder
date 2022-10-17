// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	events "github.com/gothunder/thunder/pkg/events"
	mock "github.com/stretchr/testify/mock"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// Handle provides a mock function with given fields: ctx, topic, payload
func (_m *Handler) Handle(ctx context.Context, topic string, payload []byte) events.HandlerResponse {
	ret := _m.Called(ctx, topic, payload)

	var r0 events.HandlerResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, []byte) events.HandlerResponse); ok {
		r0 = rf(ctx, topic, payload)
	} else {
		r0 = ret.Get(0).(events.HandlerResponse)
	}

	return r0
}

type mockConstructorTestingTNewHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewHandler creates a new instance of Handler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHandler(t mockConstructorTestingTNewHandler) *Handler {
	mock := &Handler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
