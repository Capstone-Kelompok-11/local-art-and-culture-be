package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIRoleRepository is a mock of IRoleRepository interface.
type MockIRoleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRoleRepositoryMockRecorder
}

// MockIRoleRepositoryMockRecorder is the mock recorder for MockIRoleRepository.
type MockIRoleRepositoryMockRecorder struct {
	mock *MockIRoleRepository
}

// NewMockIRoleRepository creates a new mock instance.
func NewMockIRoleRepository(ctrl *gomock.Controller) *MockIRoleRepository {
	mock := &MockIRoleRepository{ctrl: ctrl}
	mock.recorder = &MockIRoleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRoleRepository) EXPECT() *MockIRoleRepositoryMockRecorder {
	return m.recorder
}

// CreateRole mocks base method.
func (m *MockIRoleRepository) CreateRole(arg0 *request.Role) (response.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRole", arg0)
	ret0, _ := ret[0].(response.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRole indicates an expected call of CreateRole.
func (mr *MockIRoleRepositoryMockRecorder) CreateRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRole", reflect.TypeOf((*MockIRoleRepository)(nil).CreateRole), arg0)
}

// DeleteRole mocks base method.
func (m *MockIRoleRepository) DeleteRole(arg0 string) (response.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRole", arg0)
	ret0, _ := ret[0].(response.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRole indicates an expected call of DeleteRole.
func (mr *MockIRoleRepositoryMockRecorder) DeleteRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRole", reflect.TypeOf((*MockIRoleRepository)(nil).DeleteRole), arg0)
}

// GetAllRole mocks base method.
func (m *MockIRoleRepository) GetAllRole() ([]response.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRole")
	ret0, _ := ret[0].([]response.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRole indicates an expected call of GetAllRole.
func (mr *MockIRoleRepositoryMockRecorder) GetAllRole() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRole", reflect.TypeOf((*MockIRoleRepository)(nil).GetAllRole))
}

// GetRole mocks base method.
func (m *MockIRoleRepository) GetRole(arg0 string) (response.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRole", arg0)
	ret0, _ := ret[0].(response.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRole indicates an expected call of GetRole.
func (mr *MockIRoleRepositoryMockRecorder) GetRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRole", reflect.TypeOf((*MockIRoleRepository)(nil).GetRole), arg0)
}

// UpdateRole mocks base method.
func (m *MockIRoleRepository) UpdateRole(arg0 string, arg1 request.Role) (response.Role, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRole", arg0, arg1)
	ret0, _ := ret[0].(response.Role)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRole indicates an expected call of UpdateRole.
func (mr *MockIRoleRepositoryMockRecorder) UpdateRole(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRole", reflect.TypeOf((*MockIRoleRepository)(nil).UpdateRole), arg0, arg1)
}