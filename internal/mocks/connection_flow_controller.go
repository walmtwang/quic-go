// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/walmtwang/quic-go/internal/flowcontrol (interfaces: ConnectionFlowController)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/walmtwang/quic-go/internal/protocol"
)

// MockConnectionFlowController is a mock of ConnectionFlowController interface
type MockConnectionFlowController struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionFlowControllerMockRecorder
}

// MockConnectionFlowControllerMockRecorder is the mock recorder for MockConnectionFlowController
type MockConnectionFlowControllerMockRecorder struct {
	mock *MockConnectionFlowController
}

// NewMockConnectionFlowController creates a new mock instance
func NewMockConnectionFlowController(ctrl *gomock.Controller) *MockConnectionFlowController {
	mock := &MockConnectionFlowController{ctrl: ctrl}
	mock.recorder = &MockConnectionFlowControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConnectionFlowController) EXPECT() *MockConnectionFlowControllerMockRecorder {
	return m.recorder
}

// AddBytesRead mocks base method
func (m *MockConnectionFlowController) AddBytesRead(arg0 protocol.ByteCount) {
	m.ctrl.Call(m, "AddBytesRead", arg0)
}

// AddBytesRead indicates an expected call of AddBytesRead
func (mr *MockConnectionFlowControllerMockRecorder) AddBytesRead(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBytesRead", reflect.TypeOf((*MockConnectionFlowController)(nil).AddBytesRead), arg0)
}

// AddBytesSent mocks base method
func (m *MockConnectionFlowController) AddBytesSent(arg0 protocol.ByteCount) {
	m.ctrl.Call(m, "AddBytesSent", arg0)
}

// AddBytesSent indicates an expected call of AddBytesSent
func (mr *MockConnectionFlowControllerMockRecorder) AddBytesSent(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddBytesSent", reflect.TypeOf((*MockConnectionFlowController)(nil).AddBytesSent), arg0)
}

// GetWindowUpdate mocks base method
func (m *MockConnectionFlowController) GetWindowUpdate() protocol.ByteCount {
	ret := m.ctrl.Call(m, "GetWindowUpdate")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// GetWindowUpdate indicates an expected call of GetWindowUpdate
func (mr *MockConnectionFlowControllerMockRecorder) GetWindowUpdate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWindowUpdate", reflect.TypeOf((*MockConnectionFlowController)(nil).GetWindowUpdate))
}

// IsNewlyBlocked mocks base method
func (m *MockConnectionFlowController) IsNewlyBlocked() (bool, protocol.ByteCount) {
	ret := m.ctrl.Call(m, "IsNewlyBlocked")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(protocol.ByteCount)
	return ret0, ret1
}

// IsNewlyBlocked indicates an expected call of IsNewlyBlocked
func (mr *MockConnectionFlowControllerMockRecorder) IsNewlyBlocked() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNewlyBlocked", reflect.TypeOf((*MockConnectionFlowController)(nil).IsNewlyBlocked))
}

// MaybeQueueWindowUpdate mocks base method
func (m *MockConnectionFlowController) MaybeQueueWindowUpdate() {
	m.ctrl.Call(m, "MaybeQueueWindowUpdate")
}

// MaybeQueueWindowUpdate indicates an expected call of MaybeQueueWindowUpdate
func (mr *MockConnectionFlowControllerMockRecorder) MaybeQueueWindowUpdate() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaybeQueueWindowUpdate", reflect.TypeOf((*MockConnectionFlowController)(nil).MaybeQueueWindowUpdate))
}

// SendWindowSize mocks base method
func (m *MockConnectionFlowController) SendWindowSize() protocol.ByteCount {
	ret := m.ctrl.Call(m, "SendWindowSize")
	ret0, _ := ret[0].(protocol.ByteCount)
	return ret0
}

// SendWindowSize indicates an expected call of SendWindowSize
func (mr *MockConnectionFlowControllerMockRecorder) SendWindowSize() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendWindowSize", reflect.TypeOf((*MockConnectionFlowController)(nil).SendWindowSize))
}

// UpdateSendWindow mocks base method
func (m *MockConnectionFlowController) UpdateSendWindow(arg0 protocol.ByteCount) {
	m.ctrl.Call(m, "UpdateSendWindow", arg0)
}

// UpdateSendWindow indicates an expected call of UpdateSendWindow
func (mr *MockConnectionFlowControllerMockRecorder) UpdateSendWindow(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSendWindow", reflect.TypeOf((*MockConnectionFlowController)(nil).UpdateSendWindow), arg0)
}
