// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/walmtwang/quic-go (interfaces: PacketHandlerManager)

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/walmtwang/quic-go/internal/protocol"
)

// MockPacketHandlerManager is a mock of PacketHandlerManager interface
type MockPacketHandlerManager struct {
	ctrl     *gomock.Controller
	recorder *MockPacketHandlerManagerMockRecorder
}

// MockPacketHandlerManagerMockRecorder is the mock recorder for MockPacketHandlerManager
type MockPacketHandlerManagerMockRecorder struct {
	mock *MockPacketHandlerManager
}

// NewMockPacketHandlerManager creates a new mock instance
func NewMockPacketHandlerManager(ctrl *gomock.Controller) *MockPacketHandlerManager {
	mock := &MockPacketHandlerManager{ctrl: ctrl}
	mock.recorder = &MockPacketHandlerManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPacketHandlerManager) EXPECT() *MockPacketHandlerManagerMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockPacketHandlerManager) Add(arg0 protocol.ConnectionID, arg1 packetHandler) {
	m.ctrl.Call(m, "Add", arg0, arg1)
}

// Add indicates an expected call of Add
func (mr *MockPacketHandlerManagerMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockPacketHandlerManager)(nil).Add), arg0, arg1)
}

// CloseServer mocks base method
func (m *MockPacketHandlerManager) CloseServer() {
	m.ctrl.Call(m, "CloseServer")
}

// CloseServer indicates an expected call of CloseServer
func (mr *MockPacketHandlerManagerMockRecorder) CloseServer() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseServer", reflect.TypeOf((*MockPacketHandlerManager)(nil).CloseServer))
}

// Remove mocks base method
func (m *MockPacketHandlerManager) Remove(arg0 protocol.ConnectionID) {
	m.ctrl.Call(m, "Remove", arg0)
}

// Remove indicates an expected call of Remove
func (mr *MockPacketHandlerManagerMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockPacketHandlerManager)(nil).Remove), arg0)
}

// SetServer mocks base method
func (m *MockPacketHandlerManager) SetServer(arg0 unknownPacketHandler) {
	m.ctrl.Call(m, "SetServer", arg0)
}

// SetServer indicates an expected call of SetServer
func (mr *MockPacketHandlerManagerMockRecorder) SetServer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetServer", reflect.TypeOf((*MockPacketHandlerManager)(nil).SetServer), arg0)
}
