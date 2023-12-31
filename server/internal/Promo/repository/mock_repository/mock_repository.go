// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	reflect "reflect"
	entity "server/internal/domain/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockPromoRepositoryI is a mock of PromoRepositoryI interface.
type MockPromoRepositoryI struct {
	ctrl     *gomock.Controller
	recorder *MockPromoRepositoryIMockRecorder
}

// MockPromoRepositoryIMockRecorder is the mock recorder for MockPromoRepositoryI.
type MockPromoRepositoryIMockRecorder struct {
	mock *MockPromoRepositoryI
}

// NewMockPromoRepositoryI creates a new mock instance.
func NewMockPromoRepositoryI(ctrl *gomock.Controller) *MockPromoRepositoryI {
	mock := &MockPromoRepositoryI{ctrl: ctrl}
	mock.recorder = &MockPromoRepositoryIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPromoRepositoryI) EXPECT() *MockPromoRepositoryIMockRecorder {
	return m.recorder
}

// CheckPromo mocks base method.
func (m *MockPromoRepositoryI) CheckPromo(UserID, PromoID uint) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckPromo", UserID, PromoID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckPromo indicates an expected call of CheckPromo.
func (mr *MockPromoRepositoryIMockRecorder) CheckPromo(UserID, PromoID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckPromo", reflect.TypeOf((*MockPromoRepositoryI)(nil).CheckPromo), UserID, PromoID)
}

// DeletePromo mocks base method.
func (m *MockPromoRepositoryI) DeletePromo(UserID, PromoID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePromo", UserID, PromoID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePromo indicates an expected call of DeletePromo.
func (mr *MockPromoRepositoryIMockRecorder) DeletePromo(UserID, PromoID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePromo", reflect.TypeOf((*MockPromoRepositoryI)(nil).DeletePromo), UserID, PromoID)
}

// DeletePromoFromCart mocks base method.
func (m *MockPromoRepositoryI) DeletePromoFromCart(cartID, PromoID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePromoFromCart", cartID, PromoID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePromoFromCart indicates an expected call of DeletePromoFromCart.
func (mr *MockPromoRepositoryIMockRecorder) DeletePromoFromCart(cartID, PromoID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePromoFromCart", reflect.TypeOf((*MockPromoRepositoryI)(nil).DeletePromoFromCart), cartID, PromoID)
}

// GetPromo mocks base method.
func (m *MockPromoRepositoryI) GetPromo(code string) (*entity.Promo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPromo", code)
	ret0, _ := ret[0].(*entity.Promo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPromo indicates an expected call of GetPromo.
func (mr *MockPromoRepositoryIMockRecorder) GetPromo(code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPromo", reflect.TypeOf((*MockPromoRepositoryI)(nil).GetPromo), code)
}

// GetPromoByID mocks base method.
func (m *MockPromoRepositoryI) GetPromoByID(PromoID uint) (*entity.Promo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPromoByID", PromoID)
	ret0, _ := ret[0].(*entity.Promo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPromoByID indicates an expected call of GetPromoByID.
func (mr *MockPromoRepositoryIMockRecorder) GetPromoByID(PromoID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPromoByID", reflect.TypeOf((*MockPromoRepositoryI)(nil).GetPromoByID), PromoID)
}

// SetPromoToCart mocks base method.
func (m *MockPromoRepositoryI) SetPromoToCart(cartID, PromoID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPromoToCart", cartID, PromoID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPromoToCart indicates an expected call of SetPromoToCart.
func (mr *MockPromoRepositoryIMockRecorder) SetPromoToCart(cartID, PromoID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPromoToCart", reflect.TypeOf((*MockPromoRepositoryI)(nil).SetPromoToCart), cartID, PromoID)
}

// UsePromo mocks base method.
func (m *MockPromoRepositoryI) UsePromo(UserID, PromoID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UsePromo", UserID, PromoID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UsePromo indicates an expected call of UsePromo.
func (mr *MockPromoRepositoryIMockRecorder) UsePromo(UserID, PromoID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UsePromo", reflect.TypeOf((*MockPromoRepositoryI)(nil).UsePromo), UserID, PromoID)
}
