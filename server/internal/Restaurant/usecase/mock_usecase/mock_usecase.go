// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	reflect "reflect"
	dto "server/internal/domain/dto"

	gomock "github.com/golang/mock/gomock"
)

// MockUsecaseI is a mock of UsecaseI interface.
type MockUsecaseI struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseIMockRecorder
}

// MockUsecaseIMockRecorder is the mock recorder for MockUsecaseI.
type MockUsecaseIMockRecorder struct {
	mock *MockUsecaseI
}

// NewMockUsecaseI creates a new mock instance.
func NewMockUsecaseI(ctrl *gomock.Controller) *MockUsecaseI {
	mock := &MockUsecaseI{ctrl: ctrl}
	mock.recorder = &MockUsecaseIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecaseI) EXPECT() *MockUsecaseIMockRecorder {
	return m.recorder
}

// GetCategories mocks base method.
func (m *MockUsecaseI) GetCategories() (*[]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategories")
	ret0, _ := ret[0].(*[]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategories indicates an expected call of GetCategories.
func (mr *MockUsecaseIMockRecorder) GetCategories() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockUsecaseI)(nil).GetCategories))
}

// GetRestaurantById mocks base method.
func (m *MockUsecaseI) GetRestaurantById(id uint) (*dto.RestaurantWithCategories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurantById", id)
	ret0, _ := ret[0].(*dto.RestaurantWithCategories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurantById indicates an expected call of GetRestaurantById.
func (mr *MockUsecaseIMockRecorder) GetRestaurantById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurantById", reflect.TypeOf((*MockUsecaseI)(nil).GetRestaurantById), id)
}

// GetRestaurantProducts mocks base method.
func (m *MockUsecaseI) GetRestaurantProducts(id uint) ([]*dto.MenuTypeWithProducts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurantProducts", id)
	ret0, _ := ret[0].([]*dto.MenuTypeWithProducts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurantProducts indicates an expected call of GetRestaurantProducts.
func (mr *MockUsecaseIMockRecorder) GetRestaurantProducts(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurantProducts", reflect.TypeOf((*MockUsecaseI)(nil).GetRestaurantProducts), id)
}

// GetRestaurants mocks base method.
func (m *MockUsecaseI) GetRestaurants() ([]*dto.RestaurantWithCategories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurants")
	ret0, _ := ret[0].([]*dto.RestaurantWithCategories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurants indicates an expected call of GetRestaurants.
func (mr *MockUsecaseIMockRecorder) GetRestaurants() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurants", reflect.TypeOf((*MockUsecaseI)(nil).GetRestaurants))
}

// GetRestaurantsByCategory mocks base method.
func (m *MockUsecaseI) GetRestaurantsByCategory(name string) ([]*dto.RestaurantWithCategories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRestaurantsByCategory", name)
	ret0, _ := ret[0].([]*dto.RestaurantWithCategories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRestaurantsByCategory indicates an expected call of GetRestaurantsByCategory.
func (mr *MockUsecaseIMockRecorder) GetRestaurantsByCategory(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRestaurantsByCategory", reflect.TypeOf((*MockUsecaseI)(nil).GetRestaurantsByCategory), name)
}

// Search mocks base method.
func (m *MockUsecaseI) Search(word string) ([]*dto.RestaurantWithCategoriesAndProducts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", word)
	ret0, _ := ret[0].([]*dto.RestaurantWithCategoriesAndProducts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockUsecaseIMockRecorder) Search(word interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockUsecaseI)(nil).Search), word)
}
