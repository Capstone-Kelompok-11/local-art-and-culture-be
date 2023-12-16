package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIShippingService is a mock of IShippingService interface.
type MockIShippingService struct {
	ctrl     *gomock.Controller
	recorder *MockIShippingServiceMockRecorder
}

// MockIShippingServiceMockRecorder is the mock recorder for MockIShippingService.
type MockIShippingServiceMockRecorder struct {
	mock *MockIShippingService
}

// NewMockIShippingService creates a new mock instance.
func NewMockIShippingService(ctrl *gomock.Controller) *MockIShippingService {
	mock := &MockIShippingService{ctrl: ctrl}
	mock.recorder = &MockIShippingServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIShippingService) EXPECT() *MockIShippingServiceMockRecorder {
	return m.recorder
}

// CreateShippingMethod mocks base method.
func (m *MockIShippingService) CreateShippingMethod(arg0 *request.Shipping) (response.Shipping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateShippingMethod", arg0)
	ret0, _ := ret[0].(response.Shipping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateShippingMethod indicates an expected call of CreateShippingMethod.
func (mr *MockIShippingServiceMockRecorder) CreateShippingMethod(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateShippingMethod", reflect.TypeOf((*MockIShippingService)(nil).CreateShippingMethod), arg0)
}

// DeleteShipping mocks base method.
func (m *MockIShippingService) DeleteShipping(arg0 string) (response.Shipping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteShipping", arg0)
	ret0, _ := ret[0].(response.Shipping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteShipping indicates an expected call of DeleteShipping.
func (mr *MockIShippingServiceMockRecorder) DeleteShipping(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteShipping", reflect.TypeOf((*MockIShippingService)(nil).DeleteShipping), arg0)
}

// GetAllShipping mocks base method.
func (m *MockIShippingService) GetAllShipping() ([]response.Shipping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllShipping")
	ret0, _ := ret[0].([]response.Shipping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllShipping indicates an expected call of GetAllShipping.
func (mr *MockIShippingServiceMockRecorder) GetAllShipping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllShipping", reflect.TypeOf((*MockIShippingService)(nil).GetAllShipping))
}

// GetShipping mocks base method.
func (m *MockIShippingService) GetShipping(arg0 string) (response.Shipping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShipping", arg0)
	ret0, _ := ret[0].(response.Shipping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShipping indicates an expected call of GetShipping.
func (mr *MockIShippingServiceMockRecorder) GetShipping(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShipping", reflect.TypeOf((*MockIShippingService)(nil).GetShipping), arg0)
}

// UpdateShipping mocks base method.
func (m *MockIShippingService) UpdateShipping(arg0 string, arg1 request.Shipping) (response.Shipping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateShipping", arg0, arg1)
	ret0, _ := ret[0].(response.Shipping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateShipping indicates an expected call of UpdateShipping.
func (mr *MockIShippingServiceMockRecorder) UpdateShipping(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateShipping", reflect.TypeOf((*MockIShippingService)(nil).UpdateShipping), arg0, arg1)
}