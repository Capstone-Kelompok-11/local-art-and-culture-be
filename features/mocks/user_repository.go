package mocks

import (
	models "lokasani/entity/models"
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUserRepository is a mock of IUserRepository interface.
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository.
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance.
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// CountUsersByRole mocks base method.
func (m *MockIUserRepository) CountUsersByRole(arg0 uint) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountUsersByRole", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountUsersByRole indicates an expected call of CountUsersByRole.
func (mr *MockIUserRepositoryMockRecorder) CountUsersByRole(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountUsersByRole", reflect.TypeOf((*MockIUserRepository)(nil).CountUsersByRole), arg0)
}

// CreateUser mocks base method.
func (m *MockIUserRepository) CreateUser(arg0 *models.Users) (*models.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(*models.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockIUserRepositoryMockRecorder) CreateUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockIUserRepository)(nil).CreateUser), arg0)
}

// DeleteUser mocks base method.
func (m *MockIUserRepository) DeleteUser(arg0 string) (response.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0)
	ret0, _ := ret[0].(response.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockIUserRepositoryMockRecorder) DeleteUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockIUserRepository)(nil).DeleteUser), arg0)
}

// FindByEmail mocks base method.
func (m *MockIUserRepository) FindByEmail(arg0 string) (*models.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", arg0)
	ret0, _ := ret[0].(*models.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockIUserRepositoryMockRecorder) FindByEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockIUserRepository)(nil).FindByEmail), arg0)
}

// GetAllUser mocks base method.
func (m *MockIUserRepository) GetAllUser(arg0 string, arg1, arg2 int) ([]response.User, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUser", arg0, arg1, arg2)
	ret0, _ := ret[0].([]response.User)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllUser indicates an expected call of GetAllUser.
func (mr *MockIUserRepositoryMockRecorder) GetAllUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUser", reflect.TypeOf((*MockIUserRepository)(nil).GetAllUser), arg0, arg1, arg2)
}

// GetUser mocks base method.
func (m *MockIUserRepository) GetUser(arg0 string) (response.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(response.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockIUserRepositoryMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIUserRepository)(nil).GetUser), arg0)
}

// LoginUser mocks base method.
func (m *MockIUserRepository) LoginUser(arg0 *request.User) (response.Creators, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", arg0)
	ret0, _ := ret[0].(response.Creators)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginUser indicates an expected call of LoginUser.
func (mr *MockIUserRepositoryMockRecorder) LoginUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockIUserRepository)(nil).LoginUser), arg0)
}

// RegisterUser mocks base method.
func (m *MockIUserRepository) RegisterUser(arg0 *request.User) (response.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", arg0)
	ret0, _ := ret[0].(response.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockIUserRepositoryMockRecorder) RegisterUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockIUserRepository)(nil).RegisterUser), arg0)       
}

// UpdateUser mocks base method.
func (m *MockIUserRepository) UpdateUser(arg0 string, arg1 request.User) (response.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(response.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockIUserRepositoryMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockIUserRepository)(nil).UpdateUser), arg0, arg1)     
}