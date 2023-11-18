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

// DeleteAdmin mocks base method.
func (m *MockIAdminService) DeleteAdmin(arg0 string) (response.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAdmin", arg0)
	ret0, _ := ret[0].(response.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAdmin indicates an expected call of DeleteAdmin.
func (mr *MockIAdminServiceMockRecorder) DeleteAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAdmin", reflect.TypeOf((*MockIAdminService)(nil).DeleteAdmin), arg0)
}

// GetAdmin mocks base method.
func (m *MockIAdminService) GetAdmin(arg0 string) (response.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdmin", arg0)
	ret0, _ := ret[0].(response.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdmin indicates an expected call of GetAdmin.
func (mr *MockIAdminServiceMockRecorder) GetAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdmin", reflect.TypeOf((*MockIAdminService)(nil).GetAdmin), arg0)
}

// GetAllAdmin mocks base method.
func (m *MockIAdminService) GetAllAdmin() ([]response.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAdmin")
	ret0, _ := ret[0].([]response.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAdmin indicates an expected call of GetAllAdmin.
func (mr *MockIAdminServiceMockRecorder) GetAllAdmin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAdmin", reflect.TypeOf((*MockIAdminService)(nil).GetAllAdmin))
}

// LoginAdmin mocks base method.
func (m *MockIAdminService) LoginAdmin(arg0 *request.Admin) (response.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginAdmin", arg0)
	ret0, _ := ret[0].(response.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginAdmin indicates an expected call of LoginAdmin.
func (mr *MockIAdminServiceMockRecorder) LoginAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginAdmin", reflect.TypeOf((*MockIAdminService)(nil).LoginAdmin), arg0)
}

// RegisterAdmin mocks base method.
func (m *MockIAdminService) RegisterAdmin(arg0 *request.Admin) (response.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterAdmin", arg0)
	ret0, _ := ret[0].(response.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterAdmin indicates an expected call of RegisterAdmin.
func (mr *MockIAdminServiceMockRecorder) RegisterAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAdmin", reflect.TypeOf((*MockIAdminService)(nil).RegisterAdmin), arg0)
}

// UpdateAdmin mocks base method.
func (m *MockIAdminService) UpdateAdmin(arg0 string, arg1 request.Admin) (response.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdmin", arg0, arg1)
	ret0, _ := ret[0].(response.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAdmin indicates an expected call of UpdateAdmin.
func (mr *MockIAdminServiceMockRecorder) UpdateAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdmin", reflect.TypeOf((*MockIAdminService)(nil).UpdateAdmin), arg0, arg1)
}