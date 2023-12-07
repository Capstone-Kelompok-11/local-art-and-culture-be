package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIEventService is a mock of IEventService interface.
type MockIEventService struct {
	ctrl     *gomock.Controller
	recorder *MockIEventServiceMockRecorder
}

// MockIEventServiceMockRecorder is the mock recorder for MockIEventService.
type MockIEventServiceMockRecorder struct {
	mock *MockIEventService
}

// NewMockIEventService creates a new mock instance.
func NewMockIEventService(ctrl *gomock.Controller) *MockIEventService {
	mock := &MockIEventService{ctrl: ctrl}
	mock.recorder = &MockIEventServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIEventService) EXPECT() *MockIEventServiceMockRecorder {
	return m.recorder
}

// CalculatePaginationValues mocks base method.
func (m *MockIEventService) CalculatePaginationValues(arg0, arg1, arg2 int) (int, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculatePaginationValues", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// CalculatePaginationValues indicates an expected call of CalculatePaginationValues.
func (mr *MockIEventServiceMockRecorder) CalculatePaginationValues(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculatePaginationValues", reflect.TypeOf((*MockIEventService)(nil).CalculatePaginationValues), arg0, arg1, arg2)
}

// CreateEvent mocks base method.
func (m *MockIEventService) CreateEvent(arg0 *request.Event) (response.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEvent", arg0)
	ret0, _ := ret[0].(response.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEvent indicates an expected call of CreateEvent.
func (mr *MockIEventServiceMockRecorder) CreateEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEvent", reflect.TypeOf((*MockIEventService)(nil).CreateEvent), arg0)
}

// DeleteEvent mocks base method.
func (m *MockIEventService) DeleteEvent(arg0 string) (response.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteEvent", arg0)
	ret0, _ := ret[0].(response.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteEvent indicates an expected call of DeleteEvent.
func (mr *MockIEventServiceMockRecorder) DeleteEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteEvent", reflect.TypeOf((*MockIEventService)(nil).DeleteEvent), arg0)
}

// GetAllEvent mocks base method.
func (m *MockIEventService) GetAllEvent(arg0, arg1, arg2 string, arg3, arg4 int) ([]response.Event, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEvent", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]response.Event)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllEvent indicates an expected call of GetAllEvent.
func (mr *MockIEventServiceMockRecorder) GetAllEvent(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEvent", reflect.TypeOf((*MockIEventService)(nil).GetAllEvent), arg0, arg1, arg2, arg3, arg4)
}

// GetEvent mocks base method.
func (m *MockIEventService) GetEvent(arg0 string) (response.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEvent", arg0)
	ret0, _ := ret[0].(response.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvent indicates an expected call of GetEvent.
func (mr *MockIEventServiceMockRecorder) GetEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvent", reflect.TypeOf((*MockIEventService)(nil).GetEvent), arg0)
}

// GetNextPage mocks base method.
func (m *MockIEventService) GetNextPage(arg0, arg1 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextPage", arg0, arg1)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetNextPage indicates an expected call of GetNextPage.
func (mr *MockIEventServiceMockRecorder) GetNextPage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextPage", reflect.TypeOf((*MockIEventService)(nil).GetNextPage), arg0, arg1)
}

// GetPrevPage mocks base method.
func (m *MockIEventService) GetPrevPage(arg0 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrevPage", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetPrevPage indicates an expected call of GetPrevPage.
func (mr *MockIEventServiceMockRecorder) GetPrevPage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrevPage", reflect.TypeOf((*MockIEventService)(nil).GetPrevPage), arg0)
}

// UpdateEvent mocks base method.
func (m *MockIEventService) UpdateEvent(arg0 string, arg1 request.Event) (response.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEvent", arg0, arg1)
	ret0, _ := ret[0].(response.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEvent indicates an expected call of UpdateEvent.
func (mr *MockIEventServiceMockRecorder) UpdateEvent(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEvent", reflect.TypeOf((*MockIEventService)(nil).UpdateEvent), arg0, arg1)
}