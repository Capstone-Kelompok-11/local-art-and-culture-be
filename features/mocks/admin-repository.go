package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIAdminRepository is a mock of IAdminRepository interface.
type MockIAdminRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIAdminRepositoryMockRecorder
}

// MockIAdminRepositoryMockRecorder is the mock recorder for MockIAdminRepository.
type MockIAdminRepositoryMockRecorder struct {
	mock *MockIAdminRepository
}

// NewMockIAdminRepository creates a new mock instance.
func NewMockIAdminRepository(ctrl *gomock.Controller) *MockIAdminRepository {
	mock := &MockIAdminRepository{ctrl: ctrl}
	mock.recorder = &MockIAdminRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIAdminRepository) EXPECT() *MockIAdminRepositoryMockRecorder {
	return m.recorder
}

// DeleteAdmin mocks base method.
func (m *MockIAdminRepository) DeleteAdmin(arg0 string) (response.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAdmin", arg0)
	ret0, _ := ret[0].(response.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteAdmin indicates an expected call of DeleteAdmin.
func (mr *MockIAdminRepositoryMockRecorder) DeleteAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAdmin", reflect.TypeOf((*MockIAdminRepository)(nil).DeleteAdmin), arg0)
}

// GetAdmin mocks base method.
func (m *MockIAdminRepository) GetAdmin(arg0 string) (response.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdmin", arg0)
	ret0, _ := ret[0].(response.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdmin indicates an expected call of GetAdmin.
func (mr *MockIAdminRepositoryMockRecorder) GetAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdmin", reflect.TypeOf((*MockIAdminRepository)(nil).GetAdmin), arg0)
}

// GetAllAdmin mocks base method.
func (m *MockIAdminRepository) GetAllAdmin() ([]response.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllAdmin")
	ret0, _ := ret[0].([]response.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllAdmin indicates an expected call of GetAllAdmin.
func (mr *MockIAdminRepositoryMockRecorder) GetAllAdmin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllAdmin", reflect.TypeOf((*MockIAdminRepository)(nil).GetAllAdmin))
}

// LoginAdmin mocks base method.
func (m *MockIAdminRepository) LoginAdmin(arg0 *request.Admin) (response.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginAdmin", arg0)
	ret0, _ := ret[0].(response.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginAdmin indicates an expected call of LoginAdmin.
func (mr *MockIAdminRepositoryMockRecorder) LoginAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginAdmin", reflect.TypeOf((*MockIAdminRepository)(nil).LoginAdmin), arg0)
}

// RegisterAdmin mocks base method.
func (m *MockIAdminRepository) RegisterAdmin(arg0 *request.Admin) (response.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterAdmin", arg0)
	ret0, _ := ret[0].(response.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterAdmin indicates an expected call of RegisterAdmin.
func (mr *MockIAdminRepositoryMockRecorder) RegisterAdmin(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAdmin", reflect.TypeOf((*MockIAdminRepository)(nil).RegisterAdmin), arg0)
}

// UpdateAdmin mocks base method.
func (m *MockIAdminRepository) UpdateAdmin(arg0 string, arg1 request.Admin) (response.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdmin", arg0, arg1)
	ret0, _ := ret[0].(response.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAdmin indicates an expected call of UpdateAdmin.
func (mr *MockIAdminRepositoryMockRecorder) UpdateAdmin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdmin", reflect.TypeOf((*MockIAdminRepository)(nil).UpdateAdmin), arg0, arg1)
}