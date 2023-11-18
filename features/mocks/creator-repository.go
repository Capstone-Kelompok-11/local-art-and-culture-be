package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockICreatorRepository is a mock of ICreatorRepository interface.
type MockICreatorRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICreatorRepositoryMockRecorder
}

// MockICreatorRepositoryMockRecorder is the mock recorder for MockICreatorRepository.
type MockICreatorRepositoryMockRecorder struct {
	mock *MockICreatorRepository
}

// NewMockICreatorRepository creates a new mock instance.
func NewMockICreatorRepository(ctrl *gomock.Controller) *MockICreatorRepository {
	mock := &MockICreatorRepository{ctrl: ctrl}
	mock.recorder = &MockICreatorRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICreatorRepository) EXPECT() *MockICreatorRepositoryMockRecorder {
	return m.recorder
}

// CreateCreator mocks base method.
func (m *MockICreatorRepository) CreateCreator(arg0 *request.Creator) (response.Creator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCreator", arg0)
	ret0, _ := ret[0].(response.Creator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCreator indicates an expected call of CreateCreator.
func (mr *MockICreatorRepositoryMockRecorder) CreateCreator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCreator", reflect.TypeOf((*MockICreatorRepository)(nil).CreateCreator), arg0)
}

// DeleteCreator mocks base method.
func (m *MockICreatorRepository) DeleteCreator(arg0 string) (response.Creator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCreator", arg0)
	ret0, _ := ret[0].(response.Creator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteCreator indicates an expected call of DeleteCreator.
func (mr *MockICreatorRepositoryMockRecorder) DeleteCreator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCreator", reflect.TypeOf((*MockICreatorRepository)(nil).DeleteCreator), arg0)
}

// GetAllCreator mocks base method.
func (m *MockICreatorRepository) GetAllCreator() ([]response.UserCreatorResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCreator")
	ret0, _ := ret[0].([]response.UserCreatorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCreator indicates an expected call of GetAllCreator.
func (mr *MockICreatorRepositoryMockRecorder) GetAllCreator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCreator", reflect.TypeOf((*MockICreatorRepository)(nil).GetAllCreator))
}

// GetCreator mocks base method.
func (m *MockICreatorRepository) GetCreator(arg0 string) (response.UserCreatorResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreator", arg0)
	ret0, _ := ret[0].(response.UserCreatorResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCreator indicates an expected call of GetCreator.
func (mr *MockICreatorRepositoryMockRecorder) GetCreator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreator", reflect.TypeOf((*MockICreatorRepository)(nil).GetCreator), arg0)
}

// UpdateCreator mocks base method.
func (m *MockICreatorRepository) UpdateCreator(arg0 string, arg1 request.Creator) (response.Creator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCreator", arg0, arg1)
	ret0, _ := ret[0].(response.Creator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCreator indicates an expected call of UpdateCreator.
func (mr *MockICreatorRepositoryMockRecorder) UpdateCreator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCreator", reflect.TypeOf((*MockICreatorRepository)(nil).UpdateCreator), arg0, arg1)
}