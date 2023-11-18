package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIRoleService is a mock of IRoleService interface.
type MockIRoleService struct {
	ctrl     *gomock.Controller
	recorder *MockIRoleServiceMockRecorder
}

// MockIRoleServiceMockRecorder is the mock recorder for MockIRoleService.
type MockIRoleServiceMockRecorder struct {
	mock *MockIRoleService
}

// NewMockIRoleService creates a new mock instance.
func NewMockIRoleService(ctrl *gomock.Controller) *MockIRoleService {
	mock := &MockIRoleService{ctrl: ctrl}
	mock.recorder = &MockIRoleServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRoleService) EXPECT() *MockIRoleServiceMockRecorder {
	return m.recorder
}

// CreateRole mocks base method.
func (m *MockIRoleService) CreateRole(arg0 *request.Role) (response.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRole", arg0)
	ret0, _ := ret[0].(response.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRole indicates an expected call of CreateRole.
func (mr *MockIRoleServiceMockRecorder) CreateRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRole", reflect.TypeOf((*MockIRoleService)(nil).CreateRole), arg0)
}

// DeleteRole mocks base method.
func (m *MockIRoleService) DeleteRole(arg0 string) (response.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRole", arg0)
	ret0, _ := ret[0].(response.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRole indicates an expected call of DeleteRole.
func (mr *MockIRoleServiceMockRecorder) DeleteRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRole", reflect.TypeOf((*MockIRoleService)(nil).DeleteRole), arg0)
}

// GetAllRole mocks base method.
func (m *MockIRoleService) GetAllRole() ([]response.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRole")
	ret0, _ := ret[0].([]response.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRole indicates an expected call of GetAllRole.
func (mr *MockIRoleServiceMockRecorder) GetAllRole() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRole", reflect.TypeOf((*MockIRoleService)(nil).GetAllRole))
}

// GetRole mocks base method.
func (m *MockIRoleService) GetRole(arg0 string) (response.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole", arg0)
	ret0, _ := ret[0].(response.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRole indicates an expected call of GetRole.
func (mr *MockIRoleServiceMockRecorder) GetRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockIRoleService)(nil).GetRole), arg0)
}

// UpdateRole mocks base method.
func (m *MockIRoleService) UpdateRole(arg0 string, arg1 request.Role) (response.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRole", arg0, arg1)
	ret0, _ := ret[0].(response.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRole indicates an expected call of UpdateRole.
func (mr *MockIRoleServiceMockRecorder) UpdateRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRole", reflect.TypeOf((*MockIRoleService)(nil).UpdateRole), arg0, arg1)
}