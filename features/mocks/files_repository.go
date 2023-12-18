package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIFilesRepository is a mock of IFilesRepository interface.
type MockIFilesRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIFilesRepositoryMockRecorder
}

// MockIFilesRepositoryMockRecorder is the mock recorder for MockIFilesRepository.
type MockIFilesRepositoryMockRecorder struct {
	mock *MockIFilesRepository
}

// NewMockIFilesRepository creates a new mock instance.
func NewMockIFilesRepository(ctrl *gomock.Controller) *MockIFilesRepository {
	mock := &MockIFilesRepository{ctrl: ctrl}
	mock.recorder = &MockIFilesRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFilesRepository) EXPECT() *MockIFilesRepositoryMockRecorder {
	return m.recorder
}

// CreateFiles mocks base method.
func (m *MockIFilesRepository) CreateFiles(arg0 *request.Files) (response.Files, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateFiles", arg0)
	ret0, _ := ret[0].(response.Files)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateFiles indicates an expected call of CreateFiles.
func (mr *MockIFilesRepositoryMockRecorder) CreateFiles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateFiles", reflect.TypeOf((*MockIFilesRepository)(nil).CreateFiles), arg0)        
}

// DeleteFiles mocks base method.
func (m *MockIFilesRepository) DeleteFiles(arg0 string) (response.Files, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFiles", arg0)
	ret0, _ := ret[0].(response.Files)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFiles indicates an expected call of DeleteFiles.
func (mr *MockIFilesRepositoryMockRecorder) DeleteFiles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFiles", reflect.TypeOf((*MockIFilesRepository)(nil).DeleteFiles), arg0)        
}

// GetAllFiles mocks base method.
func (m *MockIFilesRepository) GetAllFiles(arg0 string) ([]response.Files, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFiles", arg0)
	ret0, _ := ret[0].([]response.Files)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFiles indicates an expected call of GetAllFiles.
func (mr *MockIFilesRepositoryMockRecorder) GetAllFiles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFiles", reflect.TypeOf((*MockIFilesRepository)(nil).GetAllFiles), arg0)        
}

// GetFiles mocks base method.
func (m *MockIFilesRepository) GetFiles(arg0 string) (response.Files, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFiles", arg0)
	ret0, _ := ret[0].(response.Files)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFiles indicates an expected call of GetFiles.
func (mr *MockIFilesRepositoryMockRecorder) GetFiles(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFiles", reflect.TypeOf((*MockIFilesRepository)(nil).GetFiles), arg0)
}

// UpdateFiles mocks base method.
func (m *MockIFilesRepository) UpdateFiles(arg0 string, arg1 request.Files) (response.Files, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFiles", arg0, arg1)
	ret0, _ := ret[0].(response.Files)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFiles indicates an expected call of UpdateFiles.
func (mr *MockIFilesRepositoryMockRecorder) UpdateFiles(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFiles", reflect.TypeOf((*MockIFilesRepository)(nil).UpdateFiles), arg0, arg1)  
}