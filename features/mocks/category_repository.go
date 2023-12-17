package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockICategoryRepository is a mock of ICategoryRepository interface.
type MockICategoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICategoryRepositoryMockRecorder
}

// MockICategoryRepositoryMockRecorder is the mock recorder for MockICategoryRepository.
type MockICategoryRepositoryMockRecorder struct {
	mock *MockICategoryRepository
}

// NewMockICategoryRepository creates a new mock instance.
func NewMockICategoryRepository(ctrl *gomock.Controller) *MockICategoryRepository {
	mock := &MockICategoryRepository{ctrl: ctrl}
	mock.recorder = &MockICategoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICategoryRepository) EXPECT() *MockICategoryRepositoryMockRecorder {
	return m.recorder
}

// CreateCategory mocks base method.
func (m *MockICategoryRepository) CreateCategory(arg0 *request.Category) (response.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory", arg0)
	ret0, _ := ret[0].(response.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCategory indicates an expected call of CreateCategory.
func (mr *MockICategoryRepositoryMockRecorder) CreateCategory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockICategoryRepository)(nil).CreateCategory), arg0)
}

// DeleteCategory mocks base method.
func (m *MockICategoryRepository) DeleteCategory(arg0 string) (response.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategory", arg0)
	ret0, _ := ret[0].(response.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCategory indicates an expected call of DeleteCategory.
func (mr *MockICategoryRepositoryMockRecorder) DeleteCategory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategory", reflect.TypeOf((*MockICategoryRepository)(nil).DeleteCategory), arg0)
}

// GetAllCategory mocks base method.
func (m *MockICategoryRepository) GetAllCategory(arg0 string) ([]response.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCategory", arg0)
	ret0, _ := ret[0].([]response.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCategory indicates an expected call of GetAllCategory.
func (mr *MockICategoryRepositoryMockRecorder) GetAllCategory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCategory", reflect.TypeOf((*MockICategoryRepository)(nil).GetAllCategory), arg0)
}

// GetCategory mocks base method.
func (m *MockICategoryRepository) GetCategory(arg0 string) (response.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategory", arg0)
	ret0, _ := ret[0].(response.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategory indicates an expected call of GetCategory.
func (mr *MockICategoryRepositoryMockRecorder) GetCategory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategory", reflect.TypeOf((*MockICategoryRepository)(nil).GetCategory), arg0)     
}

// GetTypeCategory mocks base method.
func (m *MockICategoryRepository) GetTypeCategory(arg0 string) (response.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTypeCategory", arg0)
	ret0, _ := ret[0].(response.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTypeCategory indicates an expected call of GetTypeCategory.
func (mr *MockICategoryRepositoryMockRecorder) GetTypeCategory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTypeCategory", reflect.TypeOf((*MockICategoryRepository)(nil).GetTypeCategory), arg0)
}

// UpdateCategory mocks base method.
func (m *MockICategoryRepository) UpdateCategory(arg0 string, arg1 request.Category) (response.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCategory", arg0, arg1)
	ret0, _ := ret[0].(response.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCategory indicates an expected call of UpdateCategory.
func (mr *MockICategoryRepositoryMockRecorder) UpdateCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategory", reflect.TypeOf((*MockICategoryRepository)(nil).UpdateCategory), arg0, arg1)
}