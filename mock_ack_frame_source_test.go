// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/walmtwang/quic-go (interfaces: AckFrameSource)

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	wire "github.com/walmtwang/quic-go/internal/wire"
)

// MockAckFrameSource is a mock of AckFrameSource interface
type MockAckFrameSource struct {
	ctrl     *gomock.Controller
	recorder *MockAckFrameSourceMockRecorder
}

// MockAckFrameSourceMockRecorder is the mock recorder for MockAckFrameSource
type MockAckFrameSourceMockRecorder struct {
	mock *MockAckFrameSource
}

// NewMockAckFrameSource creates a new mock instance
func NewMockAckFrameSource(ctrl *gomock.Controller) *MockAckFrameSource {
	mock := &MockAckFrameSource{ctrl: ctrl}
	mock.recorder = &MockAckFrameSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAckFrameSource) EXPECT() *MockAckFrameSourceMockRecorder {
	return m.recorder
}

// GetAckFrame mocks base method
func (m *MockAckFrameSource) GetAckFrame() *wire.AckFrame {
	ret := m.ctrl.Call(m, "GetAckFrame")
	ret0, _ := ret[0].(*wire.AckFrame)
	return ret0
}

// GetAckFrame indicates an expected call of GetAckFrame
func (mr *MockAckFrameSourceMockRecorder) GetAckFrame() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAckFrame", reflect.TypeOf((*MockAckFrameSource)(nil).GetAckFrame))
}

// GetStopWaitingFrame mocks base method
func (m *MockAckFrameSource) GetStopWaitingFrame(arg0 bool) *wire.StopWaitingFrame {
	ret := m.ctrl.Call(m, "GetStopWaitingFrame", arg0)
	ret0, _ := ret[0].(*wire.StopWaitingFrame)
	return ret0
}

// GetStopWaitingFrame indicates an expected call of GetStopWaitingFrame
func (mr *MockAckFrameSourceMockRecorder) GetStopWaitingFrame(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStopWaitingFrame", reflect.TypeOf((*MockAckFrameSource)(nil).GetStopWaitingFrame), arg0)
}
