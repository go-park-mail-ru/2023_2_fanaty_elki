// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"
	entity "server/internal/domain/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockProductRepositoryI is a mock of ProductRepositoryI interface.
type MockProductRepositoryI struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepositoryIMockRecorder
}

// MockProductRepositoryIMockRecorder is the mock recorder for MockProductRepositoryI.
type MockProductRepositoryIMockRecorder struct {
	mock *MockProductRepositoryI
}

// NewMockProductRepositoryI creates a new mock instance.
func NewMockProductRepositoryI(ctrl *gomock.Controller) *MockProductRepositoryI {
	mock := &MockProductRepositoryI{ctrl: ctrl}
	mock.recorder = &MockProductRepositoryIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepositoryI) EXPECT() *MockProductRepositoryIMockRecorder {
	return m.recorder
}

// GetProductByID mocks base method.
func (m *MockProductRepositoryI) GetProductByID(id uint) (*entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductByID", id)
	ret0, _ := ret[0].(*entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductByID indicates an expected call of GetProductByID.
func (mr *MockProductRepositoryIMockRecorder) GetProductByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductByID", reflect.TypeOf((*MockProductRepositoryI)(nil).GetProductByID), id)
}

// GetProductsByMenuTypeID mocks base method.
func (m *MockProductRepositoryI) GetProductsByMenuTypeID(id uint) ([]*entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsByMenuTypeID", id)
	ret0, _ := ret[0].([]*entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductsByMenuTypeID indicates an expected call of GetProductsByMenuTypeID.
func (mr *MockProductRepositoryIMockRecorder) GetProductsByMenuTypeID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsByMenuTypeID", reflect.TypeOf((*MockProductRepositoryI)(nil).GetProductsByMenuTypeID), id)
}

// GetRestaurantIDByProduct mocks base method.
func (m *MockProductRepositoryI) GetRestaurantIDByProduct(id uint) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurantIDByProduct", id)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurantIDByProduct indicates an expected call of GetRestaurantIDByProduct.
func (mr *MockProductRepositoryIMockRecorder) GetRestaurantIDByProduct(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurantIDByProduct", reflect.TypeOf((*MockProductRepositoryI)(nil).GetRestaurantIDByProduct), id)
}

// SearchProducts mocks base method.
func (m *MockProductRepositoryI) SearchProducts(word string) ([]*entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchProducts", word)
	ret0, _ := ret[0].([]*entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchProducts indicates an expected call of SearchProducts.
func (mr *MockProductRepositoryIMockRecorder) SearchProducts(word interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchProducts", reflect.TypeOf((*MockProductRepositoryI)(nil).SearchProducts), word)
}
