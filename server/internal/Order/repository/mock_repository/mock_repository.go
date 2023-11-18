// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"
	dto "server/internal/domain/dto"

	gomock "github.com/golang/mock/gomock"
)

// MockOrderRepositoryI is a mock of OrderRepositoryI interface.
type MockOrderRepositoryI struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepositoryIMockRecorder
}

// MockOrderRepositoryIMockRecorder is the mock recorder for MockOrderRepositoryI.
type MockOrderRepositoryIMockRecorder struct {
	mock *MockOrderRepositoryI
}

// NewMockOrderRepositoryI creates a new mock instance.
func NewMockOrderRepositoryI(ctrl *gomock.Controller) *MockOrderRepositoryI {
	mock := &MockOrderRepositoryI{ctrl: ctrl}
	mock.recorder = &MockOrderRepositoryIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepositoryI) EXPECT() *MockOrderRepositoryIMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockOrderRepositoryI) CreateOrder(order *dto.DBReqCreateOrder) (*dto.RespCreateOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", order)
	ret0, _ := ret[0].(*dto.RespCreateOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderRepositoryIMockRecorder) CreateOrder(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderRepositoryI)(nil).CreateOrder), order)
}

// GetOrder mocks base method.
func (m *MockOrderRepositoryI) GetOrder(reqOrder *dto.ReqGetOneOrder) (*dto.RespGetOneOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", reqOrder)
	ret0, _ := ret[0].(*dto.RespGetOneOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockOrderRepositoryIMockRecorder) GetOrder(reqOrder interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockOrderRepositoryI)(nil).GetOrder), reqOrder)
}

// GetOrders mocks base method.
func (m *MockOrderRepositoryI) GetOrders(userId uint) ([]*dto.RespGetOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", userId)
	ret0, _ := ret[0].([]*dto.RespGetOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockOrderRepositoryIMockRecorder) GetOrders(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockOrderRepositoryI)(nil).GetOrders), userId)
}

// UpdateOrder mocks base method.
func (m *MockOrderRepositoryI) UpdateOrder(order *dto.ReqUpdateOrder) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", order)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockOrderRepositoryIMockRecorder) UpdateOrder(order interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockOrderRepositoryI)(nil).UpdateOrder), order)
}
