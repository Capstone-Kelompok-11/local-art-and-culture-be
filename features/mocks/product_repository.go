package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIProductRepository is a mock of IProductRepository interface.
type MockIProductRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIProductRepositoryMockRecorder
}

// MockIProductRepositoryMockRecorder is the mock recorder for MockIProductRepository.
type MockIProductRepositoryMockRecorder struct {
	mock *MockIProductRepository
}

// NewMockIProductRepository creates a new mock instance.
func NewMockIProductRepository(ctrl *gomock.Controller) *MockIProductRepository {
	mock := &MockIProductRepository{ctrl: ctrl}
	mock.recorder = &MockIProductRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProductRepository) EXPECT() *MockIProductRepositoryMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockIProductRepository) CreateProduct(arg0 *request.Product) (response.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0)
	ret0, _ := ret[0].(response.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockIProductRepositoryMockRecorder) CreateProduct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockIProductRepository)(nil).CreateProduct), arg0)  
}

// DeleteProduct mocks base method.
func (m *MockIProductRepository) DeleteProduct(arg0 string) (response.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", arg0)
	ret0, _ := ret[0].(response.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockIProductRepositoryMockRecorder) DeleteProduct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockIProductRepository)(nil).DeleteProduct), arg0)  
}

// GetAllProduct mocks base method.
func (m *MockIProductRepository) GetAllProduct(arg0 string, arg1, arg2 int) ([]response.Product, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllProduct", arg0, arg1, arg2)
	ret0, _ := ret[0].([]response.Product)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllProduct indicates an expected call of GetAllProduct.
func (mr *MockIProductRepositoryMockRecorder) GetAllProduct(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllProduct", reflect.TypeOf((*MockIProductRepository)(nil).GetAllProduct), arg0, arg1, arg2)
}

// GetProduct mocks base method.
func (m *MockIProductRepository) GetProduct(arg0 string) (response.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", arg0)
	ret0, _ := ret[0].(response.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockIProductRepositoryMockRecorder) GetProduct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockIProductRepository)(nil).GetProduct), arg0)        
}

// GetTotalLikes mocks base method.
func (m *MockIProductRepository) GetTotalLikes(arg0 uint, arg1 string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalLikes", arg0, arg1)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalLikes indicates an expected call of GetTotalLikes.
func (mr *MockIProductRepositoryMockRecorder) GetTotalLikes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalLikes", reflect.TypeOf((*MockIProductRepository)(nil).GetTotalLikes), arg0, arg1)
}

// GetTotalLikesLastTwoWeeks mocks base method.
func (m *MockIProductRepository) GetTotalLikesLastTwoWeeks(arg0 uint, arg1 string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalLikesLastTwoWeeks", arg0, arg1)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalLikesLastTwoWeeks indicates an expected call of GetTotalLikesLastTwoWeeks.
func (mr *MockIProductRepositoryMockRecorder) GetTotalLikesLastTwoWeeks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalLikesLastTwoWeeks", reflect.TypeOf((*MockIProductRepository)(nil).GetTotalLikesLastTwoWeeks), arg0, arg1)
}

// GetTrendingProduct mocks base method.
func (m *MockIProductRepository) GetTrendingProduct(arg0 string, arg1, arg2 int) ([]response.Products, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrendingProduct", arg0, arg1, arg2)
	ret0, _ := ret[0].([]response.Products)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTrendingProduct indicates an expected call of GetTrendingProduct.
func (mr *MockIProductRepositoryMockRecorder) GetTrendingProduct(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrendingProduct", reflect.TypeOf((*MockIProductRepository)(nil).GetTrendingProduct), arg0, arg1, arg2)
}

// UpdateProduct mocks base method.
func (m *MockIProductRepository) UpdateProduct(arg0 string, arg1 request.Product) (response.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", arg0, arg1)
	ret0, _ := ret[0].(response.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockIProductRepositoryMockRecorder) UpdateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockIProductRepository)(nil).UpdateProduct), arg0, arg1)
}