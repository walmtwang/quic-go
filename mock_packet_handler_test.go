// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/walmtwang/quic-go (interfaces: PacketHandler)

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/walmtwang/quic-go/internal/protocol"
)

// MockPacketHandler is a mock of PacketHandler interface
type MockPacketHandler struct {
	ctrl     *gomock.Controller
	recorder *MockPacketHandlerMockRecorder
}

// MockPacketHandlerMockRecorder is the mock recorder for MockPacketHandler
type MockPacketHandlerMockRecorder struct {
	mock *MockPacketHandler
}

// NewMockPacketHandler creates a new mock instance
func NewMockPacketHandler(ctrl *gomock.Controller) *MockPacketHandler {
	mock := &MockPacketHandler{ctrl: ctrl}
	mock.recorder = &MockPacketHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPacketHandler) EXPECT() *MockPacketHandlerMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockPacketHandler) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockPacketHandlerMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPacketHandler)(nil).Close))
}

// GetPerspective mocks base method
func (m *MockPacketHandler) GetPerspective() protocol.Perspective {
	ret := m.ctrl.Call(m, "GetPerspective")
	ret0, _ := ret[0].(protocol.Perspective)
	return ret0
}

// GetPerspective indicates an expected call of GetPerspective
func (mr *MockPacketHandlerMockRecorder) GetPerspective() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPerspective", reflect.TypeOf((*MockPacketHandler)(nil).GetPerspective))
}

// GetVersion mocks base method
func (m *MockPacketHandler) GetVersion() protocol.VersionNumber {
	ret := m.ctrl.Call(m, "GetVersion")
	ret0, _ := ret[0].(protocol.VersionNumber)
	return ret0
}

// GetVersion indicates an expected call of GetVersion
func (mr *MockPacketHandlerMockRecorder) GetVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockPacketHandler)(nil).GetVersion))
}

// destroy mocks base method
func (m *MockPacketHandler) destroy(arg0 error) {
	m.ctrl.Call(m, "destroy", arg0)
}

// destroy indicates an expected call of destroy
func (mr *MockPacketHandlerMockRecorder) destroy(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "destroy", reflect.TypeOf((*MockPacketHandler)(nil).destroy), arg0)
}

// handlePacket mocks base method
func (m *MockPacketHandler) handlePacket(arg0 *receivedPacket) {
	m.ctrl.Call(m, "handlePacket", arg0)
}

// handlePacket indicates an expected call of handlePacket
func (mr *MockPacketHandlerMockRecorder) handlePacket(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handlePacket", reflect.TypeOf((*MockPacketHandler)(nil).handlePacket), arg0)
}
