package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIAdminService is a mock of IAdminService interface.
type MockIAdminService struct {
	ctrl     *gomock.Controller
	recorder *MockIAdminServiceMockRecorder
}

// MockIAdminServiceMockRecorder is the mock recorder for MockIAdminService.
type MockIAdminServiceMockRecorder struct {
	mock *MockIAdminService
}

// NewMockIAdminService creates a new mock instance.
func NewMockIAdminService(ctrl *gomock.Controller) *MockIAdminService {
	mock := &MockIAdminService{ctrl: ctrl}
	mock.recorder = &MockIAdminServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAdminService) EXPECT() *MockIAdminServiceMockRecorder {
	return m.recorder
}

// CalculatePaginationValues mocks base method.
func (m *MockIAdminService) CalculatePaginationValues(arg0, arg1, arg2 int) (int, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculatePaginationValues", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// CalculatePaginationValues indicates an expected call of CalculatePaginationValues.
func (mr *MockIAdminServiceMockRecorder) CalculatePaginationValues(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculatePaginationValues", reflect.TypeOf((*MockIAdminService)(nil).CalculatePaginationValues), arg0, arg1, arg2)
}

// DeleteAdmin mocks base method.
func (m *MockIAdminService) DeleteAdmin(arg0 string) (response.SuperAdmin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAdmin", arg0)
	ret0, _ := ret[0].(response.SuperAdmin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAdmin indicates an expected call of DeleteAdmin.
func (mr *MockIAdminServiceMockRecorder) DeleteAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAdmin", reflect.TypeOf((*MockIAdminService)(nil).DeleteAdmin), arg0)
}

// GetAdmin mocks base method.
func (m *MockIAdminService) GetAdmin(arg0 string) (response.SuperAdmin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdmin", arg0)
	ret0, _ := ret[0].(response.SuperAdmin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdmin indicates an expected call of GetAdmin.
func (mr *MockIAdminServiceMockRecorder) GetAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdmin", reflect.TypeOf((*MockIAdminService)(nil).GetAdmin), arg0)
}

// GetAllAdmin mocks base method.
func (m *MockIAdminService) GetAllAdmin(arg0 string, arg1, arg2 int) ([]response.SuperAdmin, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAdmin", arg0, arg1, arg2)
	ret0, _ := ret[0].([]response.SuperAdmin)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllAdmin indicates an expected call of GetAllAdmin.
func (mr *MockIAdminServiceMockRecorder) GetAllAdmin(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAdmin", reflect.TypeOf((*MockIAdminService)(nil).GetAllAdmin), arg0, arg1, arg2)
}

// GetNextPage mocks base method.
func (m *MockIAdminService) GetNextPage(arg0, arg1 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextPage", arg0, arg1)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetNextPage indicates an expected call of GetNextPage.
func (mr *MockIAdminServiceMockRecorder) GetNextPage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextPage", reflect.TypeOf((*MockIAdminService)(nil).GetNextPage), arg0, arg1)     
}

// GetPrevPage mocks base method.
func (m *MockIAdminService) GetPrevPage(arg0 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrevPage", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetPrevPage indicates an expected call of GetPrevPage.
func (mr *MockIAdminServiceMockRecorder) GetPrevPage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrevPage", reflect.TypeOf((*MockIAdminService)(nil).GetPrevPage), arg0)
}

// LoginAdmin mocks base method.
func (m *MockIAdminService) LoginAdmin(arg0 *request.SuperAdmin) (response.SuperAdmin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginAdmin", arg0)
	ret0, _ := ret[0].(response.SuperAdmin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginAdmin indicates an expected call of LoginAdmin.
func (mr *MockIAdminServiceMockRecorder) LoginAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginAdmin", reflect.TypeOf((*MockIAdminService)(nil).LoginAdmin), arg0)
}

// RegisterAdmin mocks base method.
func (m *MockIAdminService) RegisterAdmin(arg0 *request.SuperAdmin) (response.SuperAdmin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterAdmin", arg0)
	ret0, _ := ret[0].(response.SuperAdmin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterAdmin indicates an expected call of RegisterAdmin.
func (mr *MockIAdminServiceMockRecorder) RegisterAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAdmin", reflect.TypeOf((*MockIAdminService)(nil).RegisterAdmin), arg0)       
}

// UpdateAdmin mocks base method.
func (m *MockIAdminService) UpdateAdmin(arg0 string, arg1 request.SuperAdmin) (response.SuperAdmin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdmin", arg0, arg1)
	ret0, _ := ret[0].(response.SuperAdmin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAdmin indicates an expected call of UpdateAdmin.
func (mr *MockIAdminServiceMockRecorder) UpdateAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdmin", reflect.TypeOf((*MockIAdminService)(nil).UpdateAdmin), arg0, arg1)     
}