// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozonmp/bss-workplace-api/internal/repo (interfaces: WorkplaceEventRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
	model "github.com/ozonmp/bss-workplace-api/internal/model"
)

// MockWorkplaceEventRepo is a mock of WorkplaceEventRepo interface.
type MockWorkplaceEventRepo struct {
	ctrl     *gomock.Controller
	recorder *MockWorkplaceEventRepoMockRecorder
}

// MockWorkplaceEventRepoMockRecorder is the mock recorder for MockWorkplaceEventRepo.
type MockWorkplaceEventRepoMockRecorder struct {
	mock *MockWorkplaceEventRepo
}

// NewMockWorkplaceEventRepo creates a new mock instance.
func NewMockWorkplaceEventRepo(ctrl *gomock.Controller) *MockWorkplaceEventRepo {
	mock := &MockWorkplaceEventRepo{ctrl: ctrl}
	mock.recorder = &MockWorkplaceEventRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkplaceEventRepo) EXPECT() *MockWorkplaceEventRepoMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockWorkplaceEventRepo) Add(arg0 context.Context, arg1 model.WorkplaceEvent, arg2 *sqlx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockWorkplaceEventRepoMockRecorder) Add(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockWorkplaceEventRepo)(nil).Add), arg0, arg1, arg2)
}

// Lock mocks base method.
func (m *MockWorkplaceEventRepo) Lock(arg0 context.Context, arg1 uint64, arg2 *sqlx.Tx) ([]model.WorkplaceEvent, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lock", arg0, arg1, arg2)
	ret0, _ := ret[0].([]model.WorkplaceEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Lock indicates an expected call of Lock.
func (mr *MockWorkplaceEventRepoMockRecorder) Lock(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockWorkplaceEventRepo)(nil).Lock), arg0, arg1, arg2)
}

// Remove mocks base method.
func (m *MockWorkplaceEventRepo) Remove(arg0 context.Context, arg1 []uint64, arg2 *sqlx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockWorkplaceEventRepoMockRecorder) Remove(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockWorkplaceEventRepo)(nil).Remove), arg0, arg1, arg2)
}

// Unlock mocks base method.
func (m *MockWorkplaceEventRepo) Unlock(arg0 context.Context, arg1 []uint64, arg2 *sqlx.Tx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlock", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unlock indicates an expected call of Unlock.
func (mr *MockWorkplaceEventRepoMockRecorder) Unlock(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockWorkplaceEventRepo)(nil).Unlock), arg0, arg1, arg2)
}
