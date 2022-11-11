// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/favonia/cloudflare-ddns/internal/setter (interfaces: Setter)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	netip "net/netip"
	reflect "reflect"

	api "github.com/favonia/cloudflare-ddns/internal/api"
	domain "github.com/favonia/cloudflare-ddns/internal/domain"
	ipnet "github.com/favonia/cloudflare-ddns/internal/ipnet"
	pp "github.com/favonia/cloudflare-ddns/internal/pp"
	gomock "github.com/golang/mock/gomock"
)

// MockSetter is a mock of Setter interface.
type MockSetter struct {
	ctrl     *gomock.Controller
	recorder *MockSetterMockRecorder
}

// MockSetterMockRecorder is the mock recorder for MockSetter.
type MockSetterMockRecorder struct {
	mock *MockSetter
}

// NewMockSetter creates a new mock instance.
func NewMockSetter(ctrl *gomock.Controller) *MockSetter {
	mock := &MockSetter{ctrl: ctrl}
	mock.recorder = &MockSetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSetter) EXPECT() *MockSetterMockRecorder {
	return m.recorder
}

// Set mocks base method.
func (m *MockSetter) Set(arg0 context.Context, arg1 pp.PP, arg2 domain.Domain, arg3 ipnet.Type, arg4 netip.Addr, arg5 api.TTL, arg6 bool) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockSetterMockRecorder) Set(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockSetter)(nil).Set), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}
