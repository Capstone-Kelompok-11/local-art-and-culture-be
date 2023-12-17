package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIGuestRepository is a mock of IGuestRepository interface.
type MockIGuestRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIGuestRepositoryMockRecorder
}

// MockIGuestRepositoryMockRecorder is the mock recorder for MockIGuestRepository.
type MockIGuestRepositoryMockRecorder struct {
	mock *MockIGuestRepository
}

// NewMockIGuestRepository creates a new mock instance.
func NewMockIGuestRepository(ctrl *gomock.Controller) *MockIGuestRepository {
	mock := &MockIGuestRepository{ctrl: ctrl}
	mock.recorder = &MockIGuestRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGuestRepository) EXPECT() *MockIGuestRepositoryMockRecorder {
	return m.recorder
}

// CreateGuest mocks base method.
func (m *MockIGuestRepository) CreateGuest(arg0 *request.Guest) (response.Guest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateGuest", arg0)
	ret0, _ := ret[0].(response.Guest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateGuest indicates an expected call of CreateGuest.
func (mr *MockIGuestRepositoryMockRecorder) CreateGuest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateGuest", reflect.TypeOf((*MockIGuestRepository)(nil).CreateGuest), arg0)        
}

// DeleteGuest mocks base method.
func (m *MockIGuestRepository) DeleteGuest(arg0 string) (response.Guest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteGuest", arg0)
	ret0, _ := ret[0].(response.Guest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteGuest indicates an expected call of DeleteGuest.
func (mr *MockIGuestRepositoryMockRecorder) DeleteGuest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteGuest", reflect.TypeOf((*MockIGuestRepository)(nil).DeleteGuest), arg0)        
}

// GetAllGuest mocks base method.
func (m *MockIGuestRepository) GetAllGuest(arg0 string) ([]response.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllGuest", arg0)
	ret0, _ := ret[0].([]response.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllGuest indicates an expected call of GetAllGuest.
func (mr *MockIGuestRepositoryMockRecorder) GetAllGuest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllGuest", reflect.TypeOf((*MockIGuestRepository)(nil).GetAllGuest), arg0)        
}

// GetGuest mocks base method.
func (m *MockIGuestRepository) GetGuest(arg0 string) (response.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGuest", arg0)
	ret0, _ := ret[0].(response.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGuest indicates an expected call of GetGuest.
func (mr *MockIGuestRepositoryMockRecorder) GetGuest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGuest", reflect.TypeOf((*MockIGuestRepository)(nil).GetGuest), arg0)
}

// UpdateGuest mocks base method.
func (m *MockIGuestRepository) UpdateGuest(arg0 string, arg1 request.Guest) (response.Guest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGuest", arg0, arg1)
	ret0, _ := ret[0].(response.Guest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateGuest indicates an expected call of UpdateGuest.
func (mr *MockIGuestRepositoryMockRecorder) UpdateGuest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGuest", reflect.TypeOf((*MockIGuestRepository)(nil).UpdateGuest), arg0, arg1)  
}