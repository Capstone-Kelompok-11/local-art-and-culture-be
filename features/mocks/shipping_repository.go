package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIShippingRepository is a mock of IShippingRepository interface.
type MockIShippingRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIShippingRepositoryMockRecorder
}

// MockIShippingRepositoryMockRecorder is the mock recorder for MockIShippingRepository.
type MockIShippingRepositoryMockRecorder struct {
	mock *MockIShippingRepository
}

// NewMockIShippingRepository creates a new mock instance.
func NewMockIShippingRepository(ctrl *gomock.Controller) *MockIShippingRepository {
	mock := &MockIShippingRepository{ctrl: ctrl}
	mock.recorder = &MockIShippingRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIShippingRepository) EXPECT() *MockIShippingRepositoryMockRecorder {
	return m.recorder
}

// CreateShippingMethod mocks base method.
func (m *MockIShippingRepository) CreateShippingMethod(arg0 *request.Shipping) (response.Shipping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateShippingMethod", arg0)
	ret0, _ := ret[0].(response.Shipping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateShippingMethod indicates an expected call of CreateShippingMethod.
func (mr *MockIShippingRepositoryMockRecorder) CreateShippingMethod(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateShippingMethod", reflect.TypeOf((*MockIShippingRepository)(nil).CreateShippingMethod), arg0)      
}

// DeleteShipping mocks base method.
func (m *MockIShippingRepository) DeleteShipping(arg0 string) (response.Shipping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteShipping", arg0)
	ret0, _ := ret[0].(response.Shipping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteShipping indicates an expected call of DeleteShipping.
func (mr *MockIShippingRepositoryMockRecorder) DeleteShipping(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteShipping", reflect.TypeOf((*MockIShippingRepository)(nil).DeleteShipping), arg0)
}

// GetAllShipping mocks base method.
func (m *MockIShippingRepository) GetAllShipping() ([]response.Shipping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllShipping")
	ret0, _ := ret[0].([]response.Shipping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllShipping indicates an expected call of GetAllShipping.
func (mr *MockIShippingRepositoryMockRecorder) GetAllShipping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllShipping", reflect.TypeOf((*MockIShippingRepository)(nil).GetAllShipping))
}

// GetShipping mocks base method.
func (m *MockIShippingRepository) GetShipping(arg0 string) (response.Shipping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShipping", arg0)
	ret0, _ := ret[0].(response.Shipping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShipping indicates an expected call of GetShipping.
func (mr *MockIShippingRepositoryMockRecorder) GetShipping(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShipping", reflect.TypeOf((*MockIShippingRepository)(nil).GetShipping), arg0)
}

// UpdateShipping mocks base method.
func (m *MockIShippingRepository) UpdateShipping(arg0 string, arg1 request.Shipping) (response.Shipping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateShipping", arg0, arg1)
	ret0, _ := ret[0].(response.Shipping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateShipping indicates an expected call of UpdateShipping.
func (mr *MockIShippingRepositoryMockRecorder) UpdateShipping(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateShipping", reflect.TypeOf((*MockIShippingRepository)(nil).UpdateShipping), arg0, arg1)
}