// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ledgerwatch/erigon/polygon/p2p (interfaces: Service)

// Package p2p is a generated GoMock package.
package p2p

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/ledgerwatch/erigon/core/types"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// DownloadHeaders mocks base method.
func (m *MockService) DownloadHeaders(arg0 context.Context, arg1, arg2 uint64, arg3 PeerId) ([]*types.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadHeaders", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*types.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DownloadHeaders indicates an expected call of DownloadHeaders.
func (mr *MockServiceMockRecorder) DownloadHeaders(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadHeaders", reflect.TypeOf((*MockService)(nil).DownloadHeaders), arg0, arg1, arg2, arg3)
}

// MaxPeers mocks base method.
func (m *MockService) MaxPeers() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxPeers")
	ret0, _ := ret[0].(int)
	return ret0
}

// MaxPeers indicates an expected call of MaxPeers.
func (mr *MockServiceMockRecorder) MaxPeers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxPeers", reflect.TypeOf((*MockService)(nil).MaxPeers))
}

// PeersSyncProgress mocks base method.
func (m *MockService) PeersSyncProgress() PeersSyncProgress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PeersSyncProgress")
	ret0, _ := ret[0].(PeersSyncProgress)
	return ret0
}

// PeersSyncProgress indicates an expected call of PeersSyncProgress.
func (mr *MockServiceMockRecorder) PeersSyncProgress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeersSyncProgress", reflect.TypeOf((*MockService)(nil).PeersSyncProgress))
}

// Penalize mocks base method.
func (m *MockService) Penalize(arg0 context.Context, arg1 PeerId) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Penalize", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Penalize indicates an expected call of Penalize.
func (mr *MockServiceMockRecorder) Penalize(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Penalize", reflect.TypeOf((*MockService)(nil).Penalize), arg0, arg1)
}

// Start mocks base method.
func (m *MockService) Start(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", arg0)
}

// Start indicates an expected call of Start.
func (mr *MockServiceMockRecorder) Start(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockService)(nil).Start), arg0)
}

// Stop mocks base method.
func (m *MockService) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockServiceMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockService)(nil).Stop))
}
