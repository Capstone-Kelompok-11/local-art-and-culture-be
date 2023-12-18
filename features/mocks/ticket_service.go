package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockITicketService is a mock of ITicketService interface.
type MockITicketService struct {
	ctrl     *gomock.Controller
	recorder *MockITicketServiceMockRecorder
}

// MockITicketServiceMockRecorder is the mock recorder for MockITicketService.
type MockITicketServiceMockRecorder struct {
	mock *MockITicketService
}

// NewMockITicketService creates a new mock instance.
func NewMockITicketService(ctrl *gomock.Controller) *MockITicketService {
	mock := &MockITicketService{ctrl: ctrl}
	mock.recorder = &MockITicketServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITicketService) EXPECT() *MockITicketServiceMockRecorder {
	return m.recorder
}

// CreateTicket mocks base method.
func (m *MockITicketService) CreateTicket(arg0 *request.Ticket) (response.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTicket", arg0)
	ret0, _ := ret[0].(response.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTicket indicates an expected call of CreateTicket.
func (mr *MockITicketServiceMockRecorder) CreateTicket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTicket", reflect.TypeOf((*MockITicketService)(nil).CreateTicket), arg0)        
}

// DeleteTicket mocks base method.
func (m *MockITicketService) DeleteTicket(arg0 string) (response.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTicket", arg0)
	ret0, _ := ret[0].(response.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTicket indicates an expected call of DeleteTicket.
func (mr *MockITicketServiceMockRecorder) DeleteTicket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTicket", reflect.TypeOf((*MockITicketService)(nil).DeleteTicket), arg0)        
}

// GetAllTicket mocks base method.
func (m *MockITicketService) GetAllTicket(arg0 string) ([]response.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTicket", arg0)
	ret0, _ := ret[0].([]response.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTicket indicates an expected call of GetAllTicket.
func (mr *MockITicketServiceMockRecorder) GetAllTicket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTicket", reflect.TypeOf((*MockITicketService)(nil).GetAllTicket), arg0)        
}

// GetTicket mocks base method.
func (m *MockITicketService) GetTicket(arg0 string) (response.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTicket", arg0)
	ret0, _ := ret[0].(response.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTicket indicates an expected call of GetTicket.
func (mr *MockITicketServiceMockRecorder) GetTicket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicket", reflect.TypeOf((*MockITicketService)(nil).GetTicket), arg0)
}

// UpdateTicket mocks base method.
func (m *MockITicketService) UpdateTicket(arg0 string, arg1 request.Ticket) (response.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTicket", arg0, arg1)
	ret0, _ := ret[0].(response.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTicket indicates an expected call of UpdateTicket.
func (mr *MockITicketServiceMockRecorder) UpdateTicket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTicket", reflect.TypeOf((*MockITicketService)(nil).UpdateTicket), arg0, arg1)  
}