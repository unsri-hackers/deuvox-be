// Code generated by MockGen. DO NOT EDIT.
// Source: ./init.go

// Package mock_auth is a generated GoMock package.
package mock_auth

import (
	context "context"
	model "deuvox/internal/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockauthUC is a mock of authUC interface
type MockauthUC struct {
	ctrl     *gomock.Controller
	recorder *MockauthUCMockRecorder
}

// MockauthUCMockRecorder is the mock recorder for MockauthUC
type MockauthUCMockRecorder struct {
	mock *MockauthUC
}

// NewMockauthUC creates a new mock instance
func NewMockauthUC(ctrl *gomock.Controller) *MockauthUC {
	mock := &MockauthUC{ctrl: ctrl}
	mock.recorder = &MockauthUCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockauthUC) EXPECT() *MockauthUCMockRecorder {
	return m.recorder
}

// Login mocks base method
func (m *MockauthUC) Login(body model.LoginRequest) (model.LoginResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", body)
	ret0, _ := ret[0].(model.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockauthUCMockRecorder) Login(body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockauthUC)(nil).Login), body)
}

// Register mocks base method
func (m *MockauthUC) Register(ctx context.Context, body model.RegisterRequest) (model.RegisterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, body)
	ret0, _ := ret[0].(model.RegisterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockauthUCMockRecorder) Register(ctx, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockauthUC)(nil).Register), ctx, body)
}

// Token mocks base method
func (m *MockauthUC) Token(ctx context.Context, token string) (model.RegisterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token", ctx, token)
	ret0, _ := ret[0].(model.RegisterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Token indicates an expected call of Token
func (mr *MockauthUCMockRecorder) Token(ctx, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockauthUC)(nil).Token), ctx, token)
}
