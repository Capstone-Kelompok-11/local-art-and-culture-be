package mocks


import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockITicketRepository is a mock of ITicketRepository interface.
type MockITicketRepository struct {
	ctrl     *gomock.Controller
	recorder *MockITicketRepositoryMockRecorder
}

// MockITicketRepositoryMockRecorder is the mock recorder for MockITicketRepository.
type MockITicketRepositoryMockRecorder struct {
	mock *MockITicketRepository
}

// NewMockITicketRepository creates a new mock instance.
func NewMockITicketRepository(ctrl *gomock.Controller) *MockITicketRepository {
	mock := &MockITicketRepository{ctrl: ctrl}
	mock.recorder = &MockITicketRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITicketRepository) EXPECT() *MockITicketRepositoryMockRecorder {
	return m.recorder
}

// CreateTicket mocks base method.
func (m *MockITicketRepository) CreateTicket(arg0 *request.Ticket) (response.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTicket", arg0)
	ret0, _ := ret[0].(response.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTicket indicates an expected call of CreateTicket.
func (mr *MockITicketRepositoryMockRecorder) CreateTicket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTicket", reflect.TypeOf((*MockITicketRepository)(nil).CreateTicket), arg0)     
}

// DeleteTicket mocks base method.
func (m *MockITicketRepository) DeleteTicket(arg0 string) (response.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTicket", arg0)
	ret0, _ := ret[0].(response.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTicket indicates an expected call of DeleteTicket.
func (mr *MockITicketRepositoryMockRecorder) DeleteTicket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTicket", reflect.TypeOf((*MockITicketRepository)(nil).DeleteTicket), arg0)     
}

// GetAllTicket mocks base method.
func (m *MockITicketRepository) GetAllTicket(arg0 string) ([]response.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTicket", arg0)
	ret0, _ := ret[0].([]response.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTicket indicates an expected call of GetAllTicket.
func (mr *MockITicketRepositoryMockRecorder) GetAllTicket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTicket", reflect.TypeOf((*MockITicketRepository)(nil).GetAllTicket), arg0)     
}

// GetTicket mocks base method.
func (m *MockITicketRepository) GetTicket(arg0 string) (response.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTicket", arg0)
	ret0, _ := ret[0].(response.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTicket indicates an expected call of GetTicket.
func (mr *MockITicketRepositoryMockRecorder) GetTicket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTicket", reflect.TypeOf((*MockITicketRepository)(nil).GetTicket), arg0)
}

// UpdateTicket mocks base method.
func (m *MockITicketRepository) UpdateTicket(arg0 string, arg1 request.Ticket) (response.Ticket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTicket", arg0, arg1)
	ret0, _ := ret[0].(response.Ticket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTicket indicates an expected call of UpdateTicket.
func (mr *MockITicketRepositoryMockRecorder) UpdateTicket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTicket", reflect.TypeOf((*MockITicketRepository)(nil).UpdateTicket), arg0, arg1)
}