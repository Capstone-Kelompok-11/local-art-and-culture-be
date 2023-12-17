package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockICategoryService is a mock of ICategoryService interface.
type MockICategoryService struct {
	ctrl     *gomock.Controller
	recorder *MockICategoryServiceMockRecorder
}

// MockICategoryServiceMockRecorder is the mock recorder for MockICategoryService.
type MockICategoryServiceMockRecorder struct {
	mock *MockICategoryService
}

// NewMockICategoryService creates a new mock instance.
func NewMockICategoryService(ctrl *gomock.Controller) *MockICategoryService {
	mock := &MockICategoryService{ctrl: ctrl}
	mock.recorder = &MockICategoryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICategoryService) EXPECT() *MockICategoryServiceMockRecorder {
	return m.recorder
}

// CreateCategory mocks base method.
func (m *MockICategoryService) CreateCategory(arg0 *request.Category) (response.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory", arg0)
	ret0, _ := ret[0].(response.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCategory indicates an expected call of CreateCategory.
func (mr *MockICategoryServiceMockRecorder) CreateCategory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockICategoryService)(nil).CreateCategory), arg0)  
}

// DeleteCategory mocks base method.
func (m *MockICategoryService) DeleteCategory(arg0 string) (response.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategory", arg0)
	ret0, _ := ret[0].(response.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCategory indicates an expected call of DeleteCategory.
func (mr *MockICategoryServiceMockRecorder) DeleteCategory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategory", reflect.TypeOf((*MockICategoryService)(nil).DeleteCategory), arg0)  
}

// GetAllCategory mocks base method.
func (m *MockICategoryService) GetAllCategory(arg0 string) ([]response.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCategory", arg0)
	ret0, _ := ret[0].([]response.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCategory indicates an expected call of GetAllCategory.
func (mr *MockICategoryServiceMockRecorder) GetAllCategory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCategory", reflect.TypeOf((*MockICategoryService)(nil).GetAllCategory), arg0)  
}

// GetCategory mocks base method.
func (m *MockICategoryService) GetCategory(arg0 string) (response.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategory", arg0)
	ret0, _ := ret[0].(response.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategory indicates an expected call of GetCategory.
func (mr *MockICategoryServiceMockRecorder) GetCategory(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategory", reflect.TypeOf((*MockICategoryService)(nil).GetCategory), arg0)        
}

// UpdateCategory mocks base method.
func (m *MockICategoryService) UpdateCategory(arg0 string, arg1 request.Category) (response.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCategory", arg0, arg1)
	ret0, _ := ret[0].(response.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCategory indicates an expected call of UpdateCategory.
func (mr *MockICategoryServiceMockRecorder) UpdateCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategory", reflect.TypeOf((*MockICategoryService)(nil).UpdateCategory), arg0, arg1)
}