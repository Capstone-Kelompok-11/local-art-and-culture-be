package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIPaymentService is a mock of IPaymentService interface.
type MockIPaymentService struct {
	ctrl     *gomock.Controller
	recorder *MockIPaymentServiceMockRecorder
}

// MockIPaymentServiceMockRecorder is the mock recorder for MockIPaymentService.
type MockIPaymentServiceMockRecorder struct {
	mock *MockIPaymentService
}

// NewMockIPaymentService creates a new mock instance.
func NewMockIPaymentService(ctrl *gomock.Controller) *MockIPaymentService {
	mock := &MockIPaymentService{ctrl: ctrl}
	mock.recorder = &MockIPaymentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPaymentService) EXPECT() *MockIPaymentServiceMockRecorder {
	return m.recorder
}

// CreatePayment mocks base method.
func (m *MockIPaymentService) CreatePayment(arg0 *request.Payment) (response.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePayment", arg0)
	ret0, _ := ret[0].(response.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePayment indicates an expected call of CreatePayment.
func (mr *MockIPaymentServiceMockRecorder) CreatePayment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePayment", reflect.TypeOf((*MockIPaymentService)(nil).CreatePayment), arg0)
}

// DeletePayment mocks base method.
func (m *MockIPaymentService) DeletePayment(arg0 string) (response.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePayment", arg0)
	ret0, _ := ret[0].(response.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePayment indicates an expected call of DeletePayment.
func (mr *MockIPaymentServiceMockRecorder) DeletePayment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePayment", reflect.TypeOf((*MockIPaymentService)(nil).DeletePayment), arg0)
}

// GetAllPayment mocks base method.
func (m *MockIPaymentService) GetAllPayment() ([]response.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPayment")
	ret0, _ := ret[0].([]response.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPayment indicates an expected call of GetAllPayment.
func (mr *MockIPaymentServiceMockRecorder) GetAllPayment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPayment", reflect.TypeOf((*MockIPaymentService)(nil).GetAllPayment))
}

// GetPayment mocks base method.
func (m *MockIPaymentService) GetPayment(arg0 string) (response.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPayment", arg0)
	ret0, _ := ret[0].(response.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPayment indicates an expected call of GetPayment.
func (mr *MockIPaymentServiceMockRecorder) GetPayment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPayment", reflect.TypeOf((*MockIPaymentService)(nil).GetPayment), arg0)
}

// UpdatePayment mocks base method.
func (m *MockIPaymentService) UpdatePayment(arg0 string, arg1 request.Payment) (response.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePayment", arg0, arg1)
	ret0, _ := ret[0].(response.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePayment indicates an expected call of UpdatePayment.
func (mr *MockIPaymentServiceMockRecorder) UpdatePayment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePayment", reflect.TypeOf((*MockIPaymentService)(nil).UpdatePayment), arg0, arg1)
}