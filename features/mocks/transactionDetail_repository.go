package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockITransactionDetailRepository is a mock of ITransactionDetailRepository interface.
type MockITransactionDetailRepository struct {
	ctrl     *gomock.Controller
	recorder *MockITransactionDetailRepositoryMockRecorder
}

// MockITransactionDetailRepositoryMockRecorder is the mock recorder for MockITransactionDetailRepository.
type MockITransactionDetailRepositoryMockRecorder struct {
	mock *MockITransactionDetailRepository
}

// NewMockITransactionDetailRepository creates a new mock instance.
func NewMockITransactionDetailRepository(ctrl *gomock.Controller) *MockITransactionDetailRepository {
	mock := &MockITransactionDetailRepository{ctrl: ctrl}
	mock.recorder = &MockITransactionDetailRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITransactionDetailRepository) EXPECT() *MockITransactionDetailRepositoryMockRecorder {
	return m.recorder
}

// CreateTransactionDetail mocks base method.
func (m *MockITransactionDetailRepository) CreateTransactionDetail(arg0 *request.TransactionDetail) (response.TransactionDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransactionDetail", arg0)
	ret0, _ := ret[0].(response.TransactionDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransactionDetail indicates an expected call of CreateTransactionDetail.
func (mr *MockITransactionDetailRepositoryMockRecorder) CreateTransactionDetail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransactionDetail", reflect.TypeOf((*MockITransactionDetailRepository)(nil).CreateTransactionDetail), arg0)
}

// DeleteTransactionDetail mocks base method.
func (m *MockITransactionDetailRepository) DeleteTransactionDetail(arg0 string) (error, response.TransactionDetail) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTransactionDetail", arg0)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].(response.TransactionDetail)
	return ret0, ret1
}

// DeleteTransactionDetail indicates an expected call of DeleteTransactionDetail.
func (mr *MockITransactionDetailRepositoryMockRecorder) DeleteTransactionDetail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTransactionDetail", reflect.TypeOf((*MockITransactionDetailRepository)(nil).DeleteTransactionDetail), arg0)
}

// GetAllTransactionDetail mocks base method.
func (m *MockITransactionDetailRepository) GetAllTransactionDetail() (error, []response.TransactionDetail) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTransactionDetail")
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].([]response.TransactionDetail)
	return ret0, ret1
}

// GetAllTransactionDetail indicates an expected call of GetAllTransactionDetail.
func (mr *MockITransactionDetailRepositoryMockRecorder) GetAllTransactionDetail() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTransactionDetail", reflect.TypeOf((*MockITransactionDetailRepository)(nil).GetAllTransactionDetail))
}

// GetTransactionDetail mocks base method.
func (m *MockITransactionDetailRepository) GetTransactionDetail(arg0 string) (error, response.TransactionDetail) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionDetail", arg0)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].(response.TransactionDetail)
	return ret0, ret1
}

// GetTransactionDetail indicates an expected call of GetTransactionDetail.
func (mr *MockITransactionDetailRepositoryMockRecorder) GetTransactionDetail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionDetail", reflect.TypeOf((*MockITransactionDetailRepository)(nil).GetTransactionDetail), arg0)
}

// UpdateTransactionDetail mocks base method.
func (m *MockITransactionDetailRepository) UpdateTransactionDetail(arg0 string, arg1 request.TransactionDetail) (error, response.TransactionDetail) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransactionDetail", arg0, arg1)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].(response.TransactionDetail)
	return ret0, ret1
}

// UpdateTransactionDetail indicates an expected call of UpdateTransactionDetail.
func (mr *MockITransactionDetailRepositoryMockRecorder) UpdateTransactionDetail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransactionDetail", reflect.TypeOf((*MockITransactionDetailRepository)(nil).UpdateTransactionDetail), arg0, arg1)
}