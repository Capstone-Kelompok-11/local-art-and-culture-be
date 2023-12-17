package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIGuestService is a mock of IGuestService interface.
type MockIGuestService struct {
	ctrl     *gomock.Controller
	recorder *MockIGuestServiceMockRecorder
}

// MockIGuestServiceMockRecorder is the mock recorder for MockIGuestService.
type MockIGuestServiceMockRecorder struct {
	mock *MockIGuestService
}

// NewMockIGuestService creates a new mock instance.
func NewMockIGuestService(ctrl *gomock.Controller) *MockIGuestService {
	mock := &MockIGuestService{ctrl: ctrl}
	mock.recorder = &MockIGuestServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGuestService) EXPECT() *MockIGuestServiceMockRecorder {
	return m.recorder
}

// CreateGuest mocks base method.
func (m *MockIGuestService) CreateGuest(arg0 *request.Guest) (response.Guest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGuest", arg0)
	ret0, _ := ret[0].(response.Guest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGuest indicates an expected call of CreateGuest.
func (mr *MockIGuestServiceMockRecorder) CreateGuest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGuest", reflect.TypeOf((*MockIGuestService)(nil).CreateGuest), arg0)
}

// DeleteGuest mocks base method.
func (m *MockIGuestService) DeleteGuest(arg0 string) (response.Guest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGuest", arg0)
	ret0, _ := ret[0].(response.Guest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteGuest indicates an expected call of DeleteGuest.
func (mr *MockIGuestServiceMockRecorder) DeleteGuest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGuest", reflect.TypeOf((*MockIGuestService)(nil).DeleteGuest), arg0)
}

// GetAllGuest mocks base method.
func (m *MockIGuestService) GetAllGuest(arg0 string) ([]response.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllGuest", arg0)
	ret0, _ := ret[0].([]response.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllGuest indicates an expected call of GetAllGuest.
func (mr *MockIGuestServiceMockRecorder) GetAllGuest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllGuest", reflect.TypeOf((*MockIGuestService)(nil).GetAllGuest), arg0)
}

// GetGuest mocks base method.
func (m *MockIGuestService) GetGuest(arg0 string) (response.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGuest", arg0)
	ret0, _ := ret[0].(response.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGuest indicates an expected call of GetGuest.
func (mr *MockIGuestServiceMockRecorder) GetGuest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGuest", reflect.TypeOf((*MockIGuestService)(nil).GetGuest), arg0)
}

// UpdateGuest mocks base method.
func (m *MockIGuestService) UpdateGuest(arg0 string, arg1 request.Guest) (response.Guest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGuest", arg0, arg1)
	ret0, _ := ret[0].(response.Guest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateGuest indicates an expected call of UpdateGuest.
func (mr *MockIGuestServiceMockRecorder) UpdateGuest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGuest", reflect.TypeOf((*MockIGuestService)(nil).UpdateGuest), arg0, arg1)     
}