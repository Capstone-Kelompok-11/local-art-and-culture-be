package mocks

import (
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIMidtransService is a mock of IMidtransService interface.
type MockIMidtransService struct {
	ctrl     *gomock.Controller
	recorder *MockIMidtransServiceMockRecorder
}

// MockIMidtransServiceMockRecorder is the mock recorder for MockIMidtransService.
type MockIMidtransServiceMockRecorder struct {
	mock *MockIMidtransService
}

// NewMockIMidtransService creates a new mock instance.
func NewMockIMidtransService(ctrl *gomock.Controller) *MockIMidtransService {
	mock := &MockIMidtransService{ctrl: ctrl}
	mock.recorder = &MockIMidtransServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIMidtransService) EXPECT() *MockIMidtransServiceMockRecorder {
	return m.recorder
}

// Verification mocks base method.
func (m *MockIMidtransService) Verification(arg0, arg1 string) (error, response.Transaction) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Verification", arg0, arg1)
	ret0, _ := ret[0].(error)
	ret1, _ := ret[1].(response.Transaction)
	return ret0, ret1
}

// Verification indicates an expected call of Verification.
func (mr *MockIMidtransServiceMockRecorder) Verification(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verification", reflect.TypeOf((*MockIMidtransService)(nil).Verification), arg0, arg1)
}