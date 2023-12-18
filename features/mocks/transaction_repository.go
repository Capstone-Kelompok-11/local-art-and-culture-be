package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockITransactionRepository is a mock of ITransactionRepository interface.
type MockITransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockITransactionRepositoryMockRecorder
}

// MockITransactionRepositoryMockRecorder is the mock recorder for MockITransactionRepository.
type MockITransactionRepositoryMockRecorder struct {
	mock *MockITransactionRepository
}

// NewMockITransactionRepository creates a new mock instance.
func NewMockITransactionRepository(ctrl *gomock.Controller) *MockITransactionRepository {
	mock := &MockITransactionRepository{ctrl: ctrl}
	mock.recorder = &MockITransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITransactionRepository) EXPECT() *MockITransactionRepositoryMockRecorder {
	return m.recorder
}

// CreateTransaction mocks base method.
func (m *MockITransactionRepository) CreateTransaction(arg0 *request.Transaction) (response.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", arg0)
	ret0, _ := ret[0].(response.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockITransactionRepositoryMockRecorder) CreateTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockITransactionRepository)(nil).CreateTransaction), arg0)
}

// DeleteTransaction mocks base method.
func (m *MockITransactionRepository) DeleteTransaction(arg0 string) (error, response.Transaction) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTransaction", arg0)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].(response.Transaction)
	return ret0, ret1
}

// DeleteTransaction indicates an expected call of DeleteTransaction.
func (mr *MockITransactionRepositoryMockRecorder) DeleteTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTransaction", reflect.TypeOf((*MockITransactionRepository)(nil).DeleteTransaction), arg0)
}

// GetAllTransaction mocks base method.
func (m *MockITransactionRepository) GetAllTransaction() (error, []response.Transaction) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTransaction")
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].([]response.Transaction)
	return ret0, ret1
}

// GetAllTransaction indicates an expected call of GetAllTransaction.
func (mr *MockITransactionRepositoryMockRecorder) GetAllTransaction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTransaction", reflect.TypeOf((*MockITransactionRepository)(nil).GetAllTransaction))
}

// GetHistoryTransaction mocks base method.
func (m *MockITransactionRepository) GetHistoryTransaction(arg0 uint, arg1, arg2 int) ([]*response.Transaction, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistoryTransaction", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*response.Transaction)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetHistoryTransaction indicates an expected call of GetHistoryTransaction.
func (mr *MockITransactionRepositoryMockRecorder) GetHistoryTransaction(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistoryTransaction", reflect.TypeOf((*MockITransactionRepository)(nil).GetHistoryTransaction), arg0, arg1, arg2)
}

// GetReportTransaction mocks base method.
func (m *MockITransactionRepository) GetReportTransaction(arg0 uint, arg1 string) ([]response.TransactionReport, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReportTransaction", arg0, arg1)
	ret0, _ := ret[0].([]response.TransactionReport)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReportTransaction indicates an expected call of GetReportTransaction.
func (mr *MockITransactionRepositoryMockRecorder) GetReportTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReportTransaction", reflect.TypeOf((*MockITransactionRepository)(nil).GetReportTransaction), arg0, arg1)
}

// GetTransaction mocks base method.
func (m *MockITransactionRepository) GetTransaction(arg0 string) (error, response.Transaction) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransaction", arg0)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].(response.Transaction)
	return ret0, ret1
}

// GetTransaction indicates an expected call of GetTransaction.
func (mr *MockITransactionRepositoryMockRecorder) GetTransaction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransaction", reflect.TypeOf((*MockITransactionRepository)(nil).GetTransaction), arg0)
}

// UpdateTransaction mocks base method.
func (m *MockITransactionRepository) UpdateTransaction(arg0 string, arg1 request.Transaction) (error, response.Transaction) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransaction", arg0, arg1)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].(response.Transaction)
	return ret0, ret1
}

// UpdateTransaction indicates an expected call of UpdateTransaction.
func (mr *MockITransactionRepositoryMockRecorder) UpdateTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransaction", reflect.TypeOf((*MockITransactionRepository)(nil).UpdateTransaction), arg0, arg1)
}