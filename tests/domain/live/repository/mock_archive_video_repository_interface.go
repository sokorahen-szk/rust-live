// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/domain/live/repository/archive_video_repository_interface.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/sokorahen-szk/rust-live/internal/domain/live/entity"
	input "github.com/sokorahen-szk/rust-live/internal/domain/live/input"
)

// MockArchiveVideoRepositoryInterface is a mock of ArchiveVideoRepositoryInterface interface.
type MockArchiveVideoRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockArchiveVideoRepositoryInterfaceMockRecorder
}

// MockArchiveVideoRepositoryInterfaceMockRecorder is the mock recorder for MockArchiveVideoRepositoryInterface.
type MockArchiveVideoRepositoryInterfaceMockRecorder struct {
	mock *MockArchiveVideoRepositoryInterface
}

// NewMockArchiveVideoRepositoryInterface creates a new mock instance.
func NewMockArchiveVideoRepositoryInterface(ctrl *gomock.Controller) *MockArchiveVideoRepositoryInterface {
	mock := &MockArchiveVideoRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockArchiveVideoRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArchiveVideoRepositoryInterface) EXPECT() *MockArchiveVideoRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockArchiveVideoRepositoryInterface) Create(arg0 context.Context, arg1 *input.ArchiveVideoInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockArchiveVideoRepositoryInterfaceMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockArchiveVideoRepositoryInterface)(nil).Create), arg0, arg1)
}

// GetByBroadcastId mocks base method.
func (m *MockArchiveVideoRepositoryInterface) GetByBroadcastId(arg0 context.Context, arg1 *entity.VideoBroadcastId) (*entity.ArchiveVideo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByBroadcastId", arg0, arg1)
	ret0, _ := ret[0].(*entity.ArchiveVideo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByBroadcastId indicates an expected call of GetByBroadcastId.
func (mr *MockArchiveVideoRepositoryInterfaceMockRecorder) GetByBroadcastId(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByBroadcastId", reflect.TypeOf((*MockArchiveVideoRepositoryInterface)(nil).GetByBroadcastId), arg0, arg1)
}

// List mocks base method.
func (m *MockArchiveVideoRepositoryInterface) List(arg0 context.Context, arg1 *input.ListArchiveVideoInput) ([]*entity.ArchiveVideo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]*entity.ArchiveVideo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockArchiveVideoRepositoryInterfaceMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockArchiveVideoRepositoryInterface)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockArchiveVideoRepositoryInterface) Update(arg0 context.Context, arg1 *entity.VideoId, arg2 *input.UpdateArchiveVideoInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockArchiveVideoRepositoryInterfaceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockArchiveVideoRepositoryInterface)(nil).Update), arg0, arg1, arg2)
}
