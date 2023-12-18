package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIFilesService is a mock of IFilesService interface.
type MockIFilesService struct {
	ctrl     *gomock.Controller
	recorder *MockIFilesServiceMockRecorder
}

// MockIFilesServiceMockRecorder is the mock recorder for MockIFilesService.
type MockIFilesServiceMockRecorder struct {
	mock *MockIFilesService
}

// NewMockIFilesService creates a new mock instance.
func NewMockIFilesService(ctrl *gomock.Controller) *MockIFilesService {
	mock := &MockIFilesService{ctrl: ctrl}
	mock.recorder = &MockIFilesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFilesService) EXPECT() *MockIFilesServiceMockRecorder {
	return m.recorder
}

// CreateFiles mocks base method.
func (m *MockIFilesService) CreateFiles(arg0 *request.Files) (response.Files, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFiles", arg0)
	ret0, _ := ret[0].(response.Files)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFiles indicates an expected call of CreateFiles.
func (mr *MockIFilesServiceMockRecorder) CreateFiles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFiles", reflect.TypeOf((*MockIFilesService)(nil).CreateFiles), arg0)
}

// DeleteFiles mocks base method.
func (m *MockIFilesService) DeleteFiles(arg0 string) (response.Files, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFiles", arg0)
	ret0, _ := ret[0].(response.Files)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFiles indicates an expected call of DeleteFiles.
func (mr *MockIFilesServiceMockRecorder) DeleteFiles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFiles", reflect.TypeOf((*MockIFilesService)(nil).DeleteFiles), arg0)
}

// GetAllFiles mocks base method.
func (m *MockIFilesService) GetAllFiles(arg0 string) ([]response.Files, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFiles", arg0)
	ret0, _ := ret[0].([]response.Files)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFiles indicates an expected call of GetAllFiles.
func (mr *MockIFilesServiceMockRecorder) GetAllFiles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFiles", reflect.TypeOf((*MockIFilesService)(nil).GetAllFiles), arg0)
}

// GetFiles mocks base method.
func (m *MockIFilesService) GetFiles(arg0 string) (response.Files, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFiles", arg0)
	ret0, _ := ret[0].(response.Files)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFiles indicates an expected call of GetFiles.
func (mr *MockIFilesServiceMockRecorder) GetFiles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFiles", reflect.TypeOf((*MockIFilesService)(nil).GetFiles), arg0)
}

// UpdateFiles mocks base method.
func (m *MockIFilesService) UpdateFiles(arg0 string, arg1 request.Files) (response.Files, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFiles", arg0, arg1)
	ret0, _ := ret[0].(response.Files)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFiles indicates an expected call of UpdateFiles.
func (mr *MockIFilesServiceMockRecorder) UpdateFiles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFiles", reflect.TypeOf((*MockIFilesService)(nil).UpdateFiles), arg0, arg1)     
}