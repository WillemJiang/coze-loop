// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/coze-dev/coze-loop/backend/infra/limiter (interfaces: IRateLimiter)
//
// Generated by this command:
//
//	mockgen -destination=mocks/rate_limiter.go -package=mocks . IRateLimiter
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	limiter "github.com/coze-dev/coze-loop/backend/infra/limiter"
	gomock "go.uber.org/mock/gomock"
)

// MockIRateLimiter is a mock of IRateLimiter interface.
type MockIRateLimiter struct {
	ctrl     *gomock.Controller
	recorder *MockIRateLimiterMockRecorder
}

// MockIRateLimiterMockRecorder is the mock recorder for MockIRateLimiter.
type MockIRateLimiterMockRecorder struct {
	mock *MockIRateLimiter
}

// NewMockIRateLimiter creates a new mock instance.
func NewMockIRateLimiter(ctrl *gomock.Controller) *MockIRateLimiter {
	mock := &MockIRateLimiter{ctrl: ctrl}
	mock.recorder = &MockIRateLimiterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRateLimiter) EXPECT() *MockIRateLimiterMockRecorder {
	return m.recorder
}

// AllowN mocks base method.
func (m *MockIRateLimiter) AllowN(arg0 context.Context, arg1 string, arg2 int, arg3 ...limiter.LimitOptionFn) (*limiter.Result, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AllowN", varargs...)
	ret0, _ := ret[0].(*limiter.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllowN indicates an expected call of AllowN.
func (mr *MockIRateLimiterMockRecorder) AllowN(arg0, arg1, arg2 any, arg3 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllowN", reflect.TypeOf((*MockIRateLimiter)(nil).AllowN), varargs...)
}
