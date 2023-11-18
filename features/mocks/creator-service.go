package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockICreatorService is a mock of ICreatorService interface.
type MockICreatorService struct {
	ctrl     *gomock.Controller
	recorder *MockICreatorServiceMockRecorder
}

// MockICreatorServiceMockRecorder is the mock recorder for MockICreatorService.
type MockICreatorServiceMockRecorder struct {
	mock *MockICreatorService
}

// NewMockICreatorService creates a new mock instance.
func NewMockICreatorService(ctrl *gomock.Controller) *MockICreatorService {
	mock := &MockICreatorService{ctrl: ctrl}
	mock.recorder = &MockICreatorServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICreatorService) EXPECT() *MockICreatorServiceMockRecorder {
	return m.recorder
}

// CreateCreator mocks base method.
func (m *MockICreatorService) CreateCreator(arg0 *request.Creator) (response.Creator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCreator", arg0)
	ret0, _ := ret[0].(response.Creator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCreator indicates an expected call of CreateCreator.
func (mr *MockICreatorServiceMockRecorder) CreateCreator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCreator", reflect.TypeOf((*MockICreatorService)(nil).CreateCreator), arg0)
}

// DeleteCreator mocks base method.
func (m *MockICreatorService) DeleteCreator(arg0 string) (response.Creator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCreator", arg0)
	ret0, _ := ret[0].(response.Creator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCreator indicates an expected call of DeleteCreator.
func (mr *MockICreatorServiceMockRecorder) DeleteCreator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCreator", reflect.TypeOf((*MockICreatorService)(nil).DeleteCreator), arg0)
}

// GetAllCreator mocks base method.
func (m *MockICreatorService) GetAllCreator() ([]response.UserCreatorResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCreator")
	ret0, _ := ret[0].([]response.UserCreatorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCreator indicates an expected call of GetAllCreator.
func (mr *MockICreatorServiceMockRecorder) GetAllCreator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCreator", reflect.TypeOf((*MockICreatorService)(nil).GetAllCreator))
}

// GetCreator mocks base method.
func (m *MockICreatorService) GetCreator(arg0 string) (response.UserCreatorResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreator", arg0)
	ret0, _ := ret[0].(response.UserCreatorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCreator indicates an expected call of GetCreator.
func (mr *MockICreatorServiceMockRecorder) GetCreator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreator", reflect.TypeOf((*MockICreatorService)(nil).GetCreator), arg0)
}

// UpdateCreator mocks base method.
func (m *MockICreatorService) UpdateCreator(arg0 string, arg1 request.Creator) (response.Creator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCreator", arg0, arg1)
	ret0, _ := ret[0].(response.Creator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCreator indicates an expected call of UpdateCreator.
func (mr *MockICreatorServiceMockRecorder) UpdateCreator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCreator", reflect.TypeOf((*MockICreatorService)(nil).UpdateCreator), arg0, arg1)
}