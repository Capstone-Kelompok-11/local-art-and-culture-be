package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockITransactionDetailService is a mock of ITransactionDetailService interface.
type MockITransactionDetailService struct {
	ctrl     *gomock.Controller
	recorder *MockITransactionDetailServiceMockRecorder
}

// MockITransactionDetailServiceMockRecorder is the mock recorder for MockITransactionDetailService.
type MockITransactionDetailServiceMockRecorder struct {
	mock *MockITransactionDetailService
}

// NewMockITransactionDetailService creates a new mock instance.
func NewMockITransactionDetailService(ctrl *gomock.Controller) *MockITransactionDetailService {
	mock := &MockITransactionDetailService{ctrl: ctrl}
	mock.recorder = &MockITransactionDetailServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITransactionDetailService) EXPECT() *MockITransactionDetailServiceMockRecorder {
	return m.recorder
}

// CreateTransactionDetail mocks base method.
func (m *MockITransactionDetailService) CreateTransactionDetail(arg0 *request.TransactionDetail) (response.TransactionDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransactionDetail", arg0)
	ret0, _ := ret[0].(response.TransactionDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransactionDetail indicates an expected call of CreateTransactionDetail.
func (mr *MockITransactionDetailServiceMockRecorder) CreateTransactionDetail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransactionDetail", reflect.TypeOf((*MockITransactionDetailService)(nil).CreateTransactionDetail), arg0)
}

// DeleteTransactionDetail mocks base method.
func (m *MockITransactionDetailService) DeleteTransactionDetail(arg0 string) (response.TransactionDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTransactionDetail", arg0)
	ret0, _ := ret[0].(response.TransactionDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteTransactionDetail indicates an expected call of DeleteTransactionDetail.
func (mr *MockITransactionDetailServiceMockRecorder) DeleteTransactionDetail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTransactionDetail", reflect.TypeOf((*MockITransactionDetailService)(nil).DeleteTransactionDetail), arg0)
}

// GetAllTransactionDetail mocks base method.
func (m *MockITransactionDetailService) GetAllTransactionDetail() ([]response.TransactionDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllTransactionDetail")
	ret0, _ := ret[0].([]response.TransactionDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllTransactionDetail indicates an expected call of GetAllTransactionDetail.
func (mr *MockITransactionDetailServiceMockRecorder) GetAllTransactionDetail() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllTransactionDetail", reflect.TypeOf((*MockITransactionDetailService)(nil).GetAllTransactionDetail))
}

// GetTransactionDetail mocks base method.
func (m *MockITransactionDetailService) GetTransactionDetail(arg0 string) (response.TransactionDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionDetail", arg0)
	ret0, _ := ret[0].(response.TransactionDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionDetail indicates an expected call of GetTransactionDetail.
func (mr *MockITransactionDetailServiceMockRecorder) GetTransactionDetail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionDetail", reflect.TypeOf((*MockITransactionDetailService)(nil).GetTransactionDetail), arg0)
}

// UpdateTransactionDetail mocks base method.
func (m *MockITransactionDetailService) UpdateTransactionDetail(arg0 string, arg1 request.TransactionDetail) (response.TransactionDetail, error) {  
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransactionDetail", arg0, arg1)
	ret0, _ := ret[0].(response.TransactionDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTransactionDetail indicates an expected call of UpdateTransactionDetail.
func (mr *MockITransactionDetailServiceMockRecorder) UpdateTransactionDetail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransactionDetail", reflect.TypeOf((*MockITransactionDetailService)(nil).UpdateTransactionDetail), arg0, arg1)
}