// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Coderlane/go-minecraft-ping/mcclient (interfaces: MinecraftClient)

// Package mcclient is a generated GoMock package.
package mcclient

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMinecraftClient is a mock of MinecraftClient interface
type MockMinecraftClient struct {
	ctrl     *gomock.Controller
	recorder *MockMinecraftClientMockRecorder
}

// MockMinecraftClientMockRecorder is the mock recorder for MockMinecraftClient
type MockMinecraftClientMockRecorder struct {
	mock *MockMinecraftClient
}

// NewMockMinecraftClient creates a new mock instance
func NewMockMinecraftClient(ctrl *gomock.Controller) *MockMinecraftClient {
	mock := &MockMinecraftClient{ctrl: ctrl}
	mock.recorder = &MockMinecraftClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMinecraftClient) EXPECT() *MockMinecraftClientMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockMinecraftClient) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockMinecraftClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMinecraftClient)(nil).Close))
}

// Handshake mocks base method
func (m *MockMinecraftClient) Handshake(arg0 ClientState) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Handshake", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Handshake indicates an expected call of Handshake
func (mr *MockMinecraftClientMockRecorder) Handshake(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handshake", reflect.TypeOf((*MockMinecraftClient)(nil).Handshake), arg0)
}

// Status mocks base method
func (m *MockMinecraftClient) Status() (*StatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(*StatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status
func (mr *MockMinecraftClientMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockMinecraftClient)(nil).Status))
}