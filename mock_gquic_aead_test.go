// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/walmtwang/quic-go (interfaces: GQUICAEAD)

// Package quic is a generated GoMock package.
package quic

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	protocol "github.com/walmtwang/quic-go/internal/protocol"
)

// MockGQUICAEAD is a mock of GQUICAEAD interface
type MockGQUICAEAD struct {
	ctrl     *gomock.Controller
	recorder *MockGQUICAEADMockRecorder
}

// MockGQUICAEADMockRecorder is the mock recorder for MockGQUICAEAD
type MockGQUICAEADMockRecorder struct {
	mock *MockGQUICAEAD
}

// NewMockGQUICAEAD creates a new mock instance
func NewMockGQUICAEAD(ctrl *gomock.Controller) *MockGQUICAEAD {
	mock := &MockGQUICAEAD{ctrl: ctrl}
	mock.recorder = &MockGQUICAEADMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGQUICAEAD) EXPECT() *MockGQUICAEADMockRecorder {
	return m.recorder
}

// Open mocks base method
func (m *MockGQUICAEAD) Open(arg0, arg1 []byte, arg2 protocol.PacketNumber, arg3 []byte) ([]byte, protocol.EncryptionLevel, error) {
	ret := m.ctrl.Call(m, "Open", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(protocol.EncryptionLevel)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Open indicates an expected call of Open
func (mr *MockGQUICAEADMockRecorder) Open(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockGQUICAEAD)(nil).Open), arg0, arg1, arg2, arg3)
}
