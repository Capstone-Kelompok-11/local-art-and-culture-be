package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIProductService is a mock of IProductService interface.
type MockIProductService struct {
	ctrl     *gomock.Controller
	recorder *MockIProductServiceMockRecorder
}

// MockIProductServiceMockRecorder is the mock recorder for MockIProductService.
type MockIProductServiceMockRecorder struct {
	mock *MockIProductService
}

// NewMockIProductService creates a new mock instance.
func NewMockIProductService(ctrl *gomock.Controller) *MockIProductService {
	mock := &MockIProductService{ctrl: ctrl}
	mock.recorder = &MockIProductServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProductService) EXPECT() *MockIProductServiceMockRecorder {
	return m.recorder
}

// CalculatePaginationValues mocks base method.
func (m *MockIProductService) CalculatePaginationValues(arg0, arg1, arg2 int) (int, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculatePaginationValues", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// CalculatePaginationValues indicates an expected call of CalculatePaginationValues.
func (mr *MockIProductServiceMockRecorder) CalculatePaginationValues(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculatePaginationValues", reflect.TypeOf((*MockIProductService)(nil).CalculatePaginationValues), arg0, arg1, arg2)
}

// CreateProduct mocks base method.
func (m *MockIProductService) CreateProduct(arg0 *request.Product) ([]response.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0)
	ret0, _ := ret[0].([]response.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockIProductServiceMockRecorder) CreateProduct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockIProductService)(nil).CreateProduct), arg0)     
}

// DeleteProduct mocks base method.
func (m *MockIProductService) DeleteProduct(arg0 string) (response.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", arg0)
	ret0, _ := ret[0].(response.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockIProductServiceMockRecorder) DeleteProduct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockIProductService)(nil).DeleteProduct), arg0)     
}

// GetAllProduct mocks base method.
func (m *MockIProductService) GetAllProduct(arg0 string, arg1, arg2 int) ([]response.Product, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProduct", arg0, arg1, arg2)
	ret0, _ := ret[0].([]response.Product)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllProduct indicates an expected call of GetAllProduct.
func (mr *MockIProductServiceMockRecorder) GetAllProduct(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProduct", reflect.TypeOf((*MockIProductService)(nil).GetAllProduct), arg0, arg1, arg2)
}

// GetNextPage mocks base method.
func (m *MockIProductService) GetNextPage(arg0, arg1 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextPage", arg0, arg1)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetNextPage indicates an expected call of GetNextPage.
func (mr *MockIProductServiceMockRecorder) GetNextPage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextPage", reflect.TypeOf((*MockIProductService)(nil).GetNextPage), arg0, arg1)   
}

// GetPrevPage mocks base method.
func (m *MockIProductService) GetPrevPage(arg0 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrevPage", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetPrevPage indicates an expected call of GetPrevPage.
func (mr *MockIProductServiceMockRecorder) GetPrevPage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrevPage", reflect.TypeOf((*MockIProductService)(nil).GetPrevPage), arg0)
}

// GetProduct mocks base method.
func (m *MockIProductService) GetProduct(arg0 string) (response.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", arg0)
	ret0, _ := ret[0].(response.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockIProductServiceMockRecorder) GetProduct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockIProductService)(nil).GetProduct), arg0)
}

// GetTrendingProduct mocks base method.
func (m *MockIProductService) GetTrendingProduct(arg0 string, arg1, arg2 int) ([]response.Products, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrendingProduct", arg0, arg1, arg2)
	ret0, _ := ret[0].([]response.Products)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTrendingProduct indicates an expected call of GetTrendingProduct.
func (mr *MockIProductServiceMockRecorder) GetTrendingProduct(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrendingProduct", reflect.TypeOf((*MockIProductService)(nil).GetTrendingProduct), arg0, arg1, arg2)
}

// UpdateProduct mocks base method.
func (m *MockIProductService) UpdateProduct(arg0 string, arg1 request.Product) (response.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", arg0, arg1)
	ret0, _ := ret[0].(response.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockIProductServiceMockRecorder) UpdateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockIProductService)(nil).UpdateProduct), arg0, arg1)
}

// getCategoryId mocks base method.
func (m *MockIProductService) getCategoryId(arg0 string) uint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getCategoryId", arg0)
	ret0, _ := ret[0].(uint)
	return ret0
}

// getCategoryId indicates an expected call of getCategoryId.
func (mr *MockIProductServiceMockRecorder) getCategoryId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getCategoryId", reflect.TypeOf((*MockIProductService)(nil).getCategoryId), arg0)     
}