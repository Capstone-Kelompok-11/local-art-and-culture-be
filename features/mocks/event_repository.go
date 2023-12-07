package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIEventRepository is a mock of IEventRepository interface.
type MockIEventRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIEventRepositoryMockRecorder
}

// MockIEventRepositoryMockRecorder is the mock recorder for MockIEventRepository.
type MockIEventRepositoryMockRecorder struct {
	mock *MockIEventRepository
}

// NewMockIEventRepository creates a new mock instance.
func NewMockIEventRepository(ctrl *gomock.Controller) *MockIEventRepository {
	mock := &MockIEventRepository{ctrl: ctrl}
	mock.recorder = &MockIEventRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIEventRepository) EXPECT() *MockIEventRepositoryMockRecorder {
	return m.recorder
}

// CreateEvent mocks base method.
func (m *MockIEventRepository) CreateEvent(arg0 *request.Event) (response.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", arg0)
	ret0, _ := ret[0].(response.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent.
func (mr *MockIEventRepositoryMockRecorder) CreateEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockIEventRepository)(nil).CreateEvent), arg0)
}

// DeleteEvent mocks base method.
func (m *MockIEventRepository) DeleteEvent(arg0 string) (response.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEvent", arg0)
	ret0, _ := ret[0].(response.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEvent indicates an expected call of DeleteEvent.
func (mr *MockIEventRepositoryMockRecorder) DeleteEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEvent", reflect.TypeOf((*MockIEventRepository)(nil).DeleteEvent), arg0)
}

// GetAllEvent mocks base method.
func (m *MockIEventRepository) GetAllEvent(arg0, arg1, arg2 string, arg3, arg4 int) ([]response.Event, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEvent", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]response.Event)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllEvent indicates an expected call of GetAllEvent.
func (mr *MockIEventRepositoryMockRecorder) GetAllEvent(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEvent", reflect.TypeOf((*MockIEventRepository)(nil).GetAllEvent), arg0, arg1, arg2, arg3, arg4)       
}

// GetEvent mocks base method.
func (m *MockIEventRepository) GetEvent(arg0 string) (response.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvent", arg0)
	ret0, _ := ret[0].(response.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvent indicates an expected call of GetEvent.
func (mr *MockIEventRepositoryMockRecorder) GetEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvent", reflect.TypeOf((*MockIEventRepository)(nil).GetEvent), arg0)
}

// UpdateEvent mocks base method.
func (m *MockIEventRepository) UpdateEvent(arg0 string, arg1 request.Event) (response.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEvent", arg0, arg1)
	ret0, _ := ret[0].(response.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEvent indicates an expected call of UpdateEvent.
func (mr *MockIEventRepositoryMockRecorder) UpdateEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEvent", reflect.TypeOf((*MockIEventRepository)(nil).UpdateEvent), arg0, arg1)
}