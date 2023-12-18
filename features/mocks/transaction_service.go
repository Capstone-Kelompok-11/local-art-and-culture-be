package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockITransactionService is a mock of ITransactionService interface.
type MockITransactionService struct {
	ctrl     *gomock.Controller
	recorder *MockITransactionServiceMockRecorder
}

// MockITransactionServiceMockRecorder is the mock recorder for MockITransactionService.
type MockITransactionServiceMockRecorder struct {
	mock *MockITransactionService
}

// NewMockITransactionService creates a new mock instance.
func NewMockITransactionService(ctrl *gomock.Controller) *MockITransactionService {
	mock := &MockITransactionService{ctrl: ctrl}
	mock.recorder = &MockITransactionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITransactionService) EXPECT() *MockITransactionServiceMockRecorder {
	return m.recorder
}

// CalculatePaginationValues mocks base method.
func (m *MockITransactionService) CalculatePaginationValues(arg0, arg1, arg2 int) (int, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculatePaginationValues", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// CalculatePaginationValues indicates an expected call of CalculatePaginationValues.
func (mr *MockITransactionServiceMockRecorder) CalculatePaginationValues(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculatePaginationValues", reflect.TypeOf((*MockITransactionService)(nil).CalculatePaginationValues), arg0, arg1, arg2)
}

// ConfirmPayment mocks base method.
func (m *MockITransactionService) ConfirmPayment(arg0, arg1 string) (response.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfirmPayment", arg0, arg1)
	ret0, _ := ret[0].(response.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConfirmPayment indicates an expected call of ConfirmPayment.
func (mr *MockITransactionServiceMockRecorder) ConfirmPayment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmPayment", reflect.TypeOf((*MockITransactionService)(nil).ConfirmPayment), arg0, arg1)
}

// CreateTransaction mocks base method.
func (m *MockITransactionService) CreateTransaction(arg0 *request.Transaction) (response.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", arg0)
	ret0, _ := ret[0].(response.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockITransactionServiceMockRecorder) CreateTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockITransactionService)(nil).CreateTransaction), arg0)
}

// DeleteTransaction mocks base method.
func (m *MockITransactionService) DeleteTransaction(arg0 string) (response.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTransaction", arg0)
	ret0, _ := ret[0].(response.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTransaction indicates an expected call of DeleteTransaction.
func (mr *MockITransactionServiceMockRecorder) DeleteTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTransaction", reflect.TypeOf((*MockITransactionService)(nil).DeleteTransaction), arg0)
}

// GetAllTransaction mocks base method.
func (m *MockITransactionService) GetAllTransaction() ([]response.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTransaction")
	ret0, _ := ret[0].([]response.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTransaction indicates an expected call of GetAllTransaction.
func (mr *MockITransactionServiceMockRecorder) GetAllTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTransaction", reflect.TypeOf((*MockITransactionService)(nil).GetAllTransaction))
}

// GetHistoryTransaction mocks base method.
func (m *MockITransactionService) GetHistoryTransaction(arg0 uint, arg1, arg2 int) ([]*response.Transaction, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistoryTransaction", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*response.Transaction)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetHistoryTransaction indicates an expected call of GetHistoryTransaction.
func (mr *MockITransactionServiceMockRecorder) GetHistoryTransaction(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistoryTransaction", reflect.TypeOf((*MockITransactionService)(nil).GetHistoryTransaction), arg0, arg1, arg2)
}

// GetNextPage mocks base method.
func (m *MockITransactionService) GetNextPage(arg0, arg1 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextPage", arg0, arg1)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetNextPage indicates an expected call of GetNextPage.
func (mr *MockITransactionServiceMockRecorder) GetNextPage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextPage", reflect.TypeOf((*MockITransactionService)(nil).GetNextPage), arg0, arg1)
}

// GetPrevPage mocks base method.
func (m *MockITransactionService) GetPrevPage(arg0 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrevPage", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetPrevPage indicates an expected call of GetPrevPage.
func (mr *MockITransactionServiceMockRecorder) GetPrevPage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrevPage", reflect.TypeOf((*MockITransactionService)(nil).GetPrevPage), arg0)     
}

// GetReportTransaction mocks base method.
func (m *MockITransactionService) GetReportTransaction(arg0 uint, arg1 string) ([]response.TransactionReport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReportTransaction", arg0, arg1)
	ret0, _ := ret[0].([]response.TransactionReport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReportTransaction indicates an expected call of GetReportTransaction.
func (mr *MockITransactionServiceMockRecorder) GetReportTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReportTransaction", reflect.TypeOf((*MockITransactionService)(nil).GetReportTransaction), arg0, arg1)
}

// GetTransaction mocks base method.
func (m *MockITransactionService) GetTransaction(arg0 string) (response.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransaction", arg0)
	ret0, _ := ret[0].(response.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransaction indicates an expected call of GetTransaction.
func (mr *MockITransactionServiceMockRecorder) GetTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockITransactionService)(nil).GetTransaction), arg0)
}

// UpdateTransaction mocks base method.
func (m *MockITransactionService) UpdateTransaction(arg0 string, arg1 request.Transaction) (response.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransaction", arg0, arg1)
	ret0, _ := ret[0].(response.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTransaction indicates an expected call of UpdateTransaction.
func (mr *MockITransactionServiceMockRecorder) UpdateTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransaction", reflect.TypeOf((*MockITransactionService)(nil).UpdateTransaction), arg0, arg1)
}