package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUserService is a mock of IUserService interface.
type MockIUserService struct {
	ctrl     *gomock.Controller
	recorder *MockIUserServiceMockRecorder
}

// MockIUserServiceMockRecorder is the mock recorder for MockIUserService.
type MockIUserServiceMockRecorder struct {
	mock *MockIUserService
}

// NewMockIUserService creates a new mock instance.
func NewMockIUserService(ctrl *gomock.Controller) *MockIUserService {
	mock := &MockIUserService{ctrl: ctrl}
	mock.recorder = &MockIUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserService) EXPECT() *MockIUserServiceMockRecorder {
	return m.recorder
}

// CalculatePaginationValues mocks base method.
func (m *MockIUserService) CalculatePaginationValues(arg0, arg1, arg2 int) (int, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculatePaginationValues", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// CalculatePaginationValues indicates an expected call of CalculatePaginationValues.
func (mr *MockIUserServiceMockRecorder) CalculatePaginationValues(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculatePaginationValues", reflect.TypeOf((*MockIUserService)(nil).CalculatePaginationValues), arg0, arg1, arg2)
}

// CountUsersByRole mocks base method.
func (m *MockIUserService) CountUsersByRole(arg0 uint) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountUsersByRole", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountUsersByRole indicates an expected call of CountUsersByRole.
func (mr *MockIUserServiceMockRecorder) CountUsersByRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountUsersByRole", reflect.TypeOf((*MockIUserService)(nil).CountUsersByRole), arg0)  
}

// DeleteUser mocks base method.
func (m *MockIUserService) DeleteUser(arg0 string) (response.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(response.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockIUserServiceMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIUserService)(nil).DeleteUser), arg0)
}

// GetAllUser mocks base method.
func (m *MockIUserService) GetAllUser(arg0 string, arg1, arg2 int) ([]response.User, map[string]int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUser", arg0, arg1, arg2)
	ret0, _ := ret[0].([]response.User)
	ret1, _ := ret[1].(map[string]int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllUser indicates an expected call of GetAllUser.
func (mr *MockIUserServiceMockRecorder) GetAllUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUser", reflect.TypeOf((*MockIUserService)(nil).GetAllUser), arg0, arg1, arg2)  
}

// GetNextPage mocks base method.
func (m *MockIUserService) GetNextPage(arg0, arg1 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextPage", arg0, arg1)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetNextPage indicates an expected call of GetNextPage.
func (mr *MockIUserServiceMockRecorder) GetNextPage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextPage", reflect.TypeOf((*MockIUserService)(nil).GetNextPage), arg0, arg1)      
}

// GetPrevPage mocks base method.
func (m *MockIUserService) GetPrevPage(arg0 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrevPage", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetPrevPage indicates an expected call of GetPrevPage.
func (mr *MockIUserServiceMockRecorder) GetPrevPage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrevPage", reflect.TypeOf((*MockIUserService)(nil).GetPrevPage), arg0)
}

// GetUser mocks base method.
func (m *MockIUserService) GetUser(arg0 string) (response.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(response.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockIUserServiceMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIUserService)(nil).GetUser), arg0)
}

// LoginUser mocks base method.
func (m *MockIUserService) LoginUser(arg0 *request.User) (response.Creators, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", arg0)
	ret0, _ := ret[0].(response.Creators)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginUser indicates an expected call of LoginUser.
func (mr *MockIUserServiceMockRecorder) LoginUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockIUserService)(nil).LoginUser), arg0)
}

// RegisterUser mocks base method.
func (m *MockIUserService) RegisterUser(arg0 *request.User) (response.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", arg0)
	ret0, _ := ret[0].(response.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockIUserServiceMockRecorder) RegisterUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockIUserService)(nil).RegisterUser), arg0)
}

// UpdateUser mocks base method.
func (m *MockIUserService) UpdateUser(arg0 string, arg1 request.User) (response.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(response.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockIUserServiceMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockIUserService)(nil).UpdateUser), arg0, arg1)        
}