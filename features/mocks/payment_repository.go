package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIPaymentRepository is a mock of IPaymentRepository interface.
type MockIPaymentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIPaymentRepositoryMockRecorder
}

// MockIPaymentRepositoryMockRecorder is the mock recorder for MockIPaymentRepository.
type MockIPaymentRepositoryMockRecorder struct {
	mock *MockIPaymentRepository
}

// NewMockIPaymentRepository creates a new mock instance.
func NewMockIPaymentRepository(ctrl *gomock.Controller) *MockIPaymentRepository {
	mock := &MockIPaymentRepository{ctrl: ctrl}
	mock.recorder = &MockIPaymentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPaymentRepository) EXPECT() *MockIPaymentRepositoryMockRecorder {
	return m.recorder
}

// CreatePayment mocks base method.
func (m *MockIPaymentRepository) CreatePayment(arg0 *request.Payment) (response.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePayment", arg0)
	ret0, _ := ret[0].(response.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePayment indicates an expected call of CreatePayment.
func (mr *MockIPaymentRepositoryMockRecorder) CreatePayment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePayment", reflect.TypeOf((*MockIPaymentRepository)(nil).CreatePayment), arg0)
}

// DeletePayment mocks base method.
func (m *MockIPaymentRepository) DeletePayment(arg0 string) (response.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePayment", arg0)
	ret0, _ := ret[0].(response.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletePayment indicates an expected call of DeletePayment.
func (mr *MockIPaymentRepositoryMockRecorder) DeletePayment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePayment", reflect.TypeOf((*MockIPaymentRepository)(nil).DeletePayment), arg0)
}

// GetAllPayment mocks base method.
func (m *MockIPaymentRepository) GetAllPayment() ([]response.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPayment")
	ret0, _ := ret[0].([]response.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPayment indicates an expected call of GetAllPayment.
func (mr *MockIPaymentRepositoryMockRecorder) GetAllPayment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPayment", reflect.TypeOf((*MockIPaymentRepository)(nil).GetAllPayment))
}

// GetPayment mocks base method.
func (m *MockIPaymentRepository) GetPayment(arg0 string) (response.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPayment", arg0)
	ret0, _ := ret[0].(response.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPayment indicates an expected call of GetPayment.
func (mr *MockIPaymentRepositoryMockRecorder) GetPayment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPayment", reflect.TypeOf((*MockIPaymentRepository)(nil).GetPayment), arg0)
}

// UpdatePayment mocks base method.
func (m *MockIPaymentRepository) UpdatePayment(arg0 string, arg1 request.Payment) (response.Payment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePayment", arg0, arg1)
	ret0, _ := ret[0].(response.Payment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePayment indicates an expected call of UpdatePayment.
func (mr *MockIPaymentRepositoryMockRecorder) UpdatePayment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePayment", reflect.TypeOf((*MockIPaymentRepository)(nil).UpdatePayment), arg0, arg1)
}