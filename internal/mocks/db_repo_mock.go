// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozonmp/bss-workplace-api/internal/repo (interfaces: WorkplaceRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
	model "github.com/ozonmp/bss-workplace-api/internal/model"
)

// MockWorkplaceRepo is a mock of WorkplaceRepo interface.
type MockWorkplaceRepo struct {
	ctrl     *gomock.Controller
	recorder *MockWorkplaceRepoMockRecorder
}

// MockWorkplaceRepoMockRecorder is the mock recorder for MockWorkplaceRepo.
type MockWorkplaceRepoMockRecorder struct {
	mock *MockWorkplaceRepo
}

// NewMockWorkplaceRepo creates a new mock instance.
func NewMockWorkplaceRepo(ctrl *gomock.Controller) *MockWorkplaceRepo {
	mock := &MockWorkplaceRepo{ctrl: ctrl}
	mock.recorder = &MockWorkplaceRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorkplaceRepo) EXPECT() *MockWorkplaceRepoMockRecorder {
	return m.recorder
}

// CreateWorkplace mocks base method.
func (m *MockWorkplaceRepo) CreateWorkplace(arg0 context.Context, arg1 string, arg2 uint32, arg3 *sqlx.Tx) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWorkplace", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWorkplace indicates an expected call of CreateWorkplace.
func (mr *MockWorkplaceRepoMockRecorder) CreateWorkplace(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWorkplace", reflect.TypeOf((*MockWorkplaceRepo)(nil).CreateWorkplace), arg0, arg1, arg2, arg3)
}

// DescribeWorkplace mocks base method.
func (m *MockWorkplaceRepo) DescribeWorkplace(arg0 context.Context, arg1 uint64) (*model.Workplace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeWorkplace", arg0, arg1)
	ret0, _ := ret[0].(*model.Workplace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeWorkplace indicates an expected call of DescribeWorkplace.
func (mr *MockWorkplaceRepoMockRecorder) DescribeWorkplace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeWorkplace", reflect.TypeOf((*MockWorkplaceRepo)(nil).DescribeWorkplace), arg0, arg1)
}

// ListWorkplaces mocks base method.
func (m *MockWorkplaceRepo) ListWorkplaces(arg0 context.Context, arg1, arg2 uint64) ([]model.Workplace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListWorkplaces", arg0, arg1, arg2)
	ret0, _ := ret[0].([]model.Workplace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListWorkplaces indicates an expected call of ListWorkplaces.
func (mr *MockWorkplaceRepoMockRecorder) ListWorkplaces(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListWorkplaces", reflect.TypeOf((*MockWorkplaceRepo)(nil).ListWorkplaces), arg0, arg1, arg2)
}

// RemoveWorkplace mocks base method.
func (m *MockWorkplaceRepo) RemoveWorkplace(arg0 context.Context, arg1 uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveWorkplace", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveWorkplace indicates an expected call of RemoveWorkplace.
func (mr *MockWorkplaceRepoMockRecorder) RemoveWorkplace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveWorkplace", reflect.TypeOf((*MockWorkplaceRepo)(nil).RemoveWorkplace), arg0, arg1)
}
