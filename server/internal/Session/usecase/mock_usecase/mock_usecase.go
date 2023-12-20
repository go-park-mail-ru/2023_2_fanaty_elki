// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	reflect "reflect"
	dto "server/internal/domain/dto"
	entity "server/internal/domain/entity"

	gomock "github.com/golang/mock/gomock"
)

// MockSessionUsecaseI is a mock of SessionUsecaseI interface.
type MockSessionUsecaseI struct {
	ctrl     *gomock.Controller
	recorder *MockSessionUsecaseIMockRecorder
}

// MockSessionUsecaseIMockRecorder is the mock recorder for MockSessionUsecaseI.
type MockSessionUsecaseIMockRecorder struct {
	mock *MockSessionUsecaseI
}

// NewMockSessionUsecaseI creates a new mock instance.
func NewMockSessionUsecaseI(ctrl *gomock.Controller) *MockSessionUsecaseI {
	mock := &MockSessionUsecaseI{ctrl: ctrl}
	mock.recorder = &MockSessionUsecaseIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessionUsecaseI) EXPECT() *MockSessionUsecaseIMockRecorder {
	return m.recorder
}

// Check mocks base method.
func (m *MockSessionUsecaseI) Check(SessionToken string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", SessionToken)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Check indicates an expected call of Check.
func (mr *MockSessionUsecaseIMockRecorder) Check(SessionToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockSessionUsecaseI)(nil).Check), SessionToken)
}

// CheckCsrf mocks base method.
func (m *MockSessionUsecaseI) CheckCsrf(sessionToken, csrfToken string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckCsrf", sessionToken, csrfToken)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckCsrf indicates an expected call of CheckCsrf.
func (mr *MockSessionUsecaseIMockRecorder) CheckCsrf(sessionToken, csrfToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckCsrf", reflect.TypeOf((*MockSessionUsecaseI)(nil).CheckCsrf), sessionToken, csrfToken)
}

// CreateCookieAuth mocks base method.
func (m *MockSessionUsecaseI) CreateCookieAuth(cookie *entity.Cookie) (*dto.ReqGetUserProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCookieAuth", cookie)
	ret0, _ := ret[0].(*dto.ReqGetUserProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCookieAuth indicates an expected call of CreateCookieAuth.
func (mr *MockSessionUsecaseIMockRecorder) CreateCookieAuth(cookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCookieAuth", reflect.TypeOf((*MockSessionUsecaseI)(nil).CreateCookieAuth), cookie)
}

// CreateCsrf mocks base method.
func (m *MockSessionUsecaseI) CreateCsrf(sessionToken string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCsrf", sessionToken)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCsrf indicates an expected call of CreateCsrf.
func (mr *MockSessionUsecaseIMockRecorder) CreateCsrf(sessionToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCsrf", reflect.TypeOf((*MockSessionUsecaseI)(nil).CreateCsrf), sessionToken)
}

// GetIDByCookie mocks base method.
func (m *MockSessionUsecaseI) GetIDByCookie(SessionToken string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIDByCookie", SessionToken)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIDByCookie indicates an expected call of GetIDByCookie.
func (mr *MockSessionUsecaseIMockRecorder) GetIDByCookie(SessionToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIDByCookie", reflect.TypeOf((*MockSessionUsecaseI)(nil).GetIDByCookie), SessionToken)
}

// GetUserProfile mocks base method.
func (m *MockSessionUsecaseI) GetUserProfile(sessionToken string) (*dto.ReqGetUserProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProfile", sessionToken)
	ret0, _ := ret[0].(*dto.ReqGetUserProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProfile indicates an expected call of GetUserProfile.
func (mr *MockSessionUsecaseIMockRecorder) GetUserProfile(sessionToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProfile", reflect.TypeOf((*MockSessionUsecaseI)(nil).GetUserProfile), sessionToken)
}

// Login mocks base method.
func (m *MockSessionUsecaseI) Login(user *entity.User) (*entity.Cookie, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", user)
	ret0, _ := ret[0].(*entity.Cookie)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockSessionUsecaseIMockRecorder) Login(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockSessionUsecaseI)(nil).Login), user)
}

// Logout mocks base method.
func (m *MockSessionUsecaseI) Logout(cookie *entity.Cookie) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", cookie)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockSessionUsecaseIMockRecorder) Logout(cookie interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockSessionUsecaseI)(nil).Logout), cookie)
}
