// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"
	dto "server/internal/domain/dto"
	entity "server/internal/domain/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockSessionRepositoryI is a mock of SessionRepositoryI interface.
type MockSessionRepositoryI struct {
	ctrl     *gomock.Controller
	recorder *MockSessionRepositoryIMockRecorder
}

// MockSessionRepositoryIMockRecorder is the mock recorder for MockSessionRepositoryI.
type MockSessionRepositoryIMockRecorder struct {
	mock *MockSessionRepositoryI
}

// NewMockSessionRepositoryI creates a new mock instance.
func NewMockSessionRepositoryI(ctrl *gomock.Controller) *MockSessionRepositoryI {
	mock := &MockSessionRepositoryI{ctrl: ctrl}
	mock.recorder = &MockSessionRepositoryIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionRepositoryI) EXPECT() *MockSessionRepositoryIMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockSessionRepositoryI) Check(sessionToken string) (*entity.Cookie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", sessionToken)
	ret0, _ := ret[0].(*entity.Cookie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Check indicates an expected call of Check.
func (mr *MockSessionRepositoryIMockRecorder) Check(sessionToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockSessionRepositoryI)(nil).Check), sessionToken)
}

// Create mocks base method.
func (m *MockSessionRepositoryI) Create(cookie *entity.Cookie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", cookie)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockSessionRepositoryIMockRecorder) Create(cookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSessionRepositoryI)(nil).Create), cookie)
}

// Delete mocks base method.
func (m *MockSessionRepositoryI) Delete(cookie *dto.DBDeleteCookie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", cookie)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockSessionRepositoryIMockRecorder) Delete(cookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockSessionRepositoryI)(nil).Delete), cookie)
}
