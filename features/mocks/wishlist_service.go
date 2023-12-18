package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIWishlistService is a mock of IWishlistService interface.
type MockIWishlistService struct {
	ctrl     *gomock.Controller
	recorder *MockIWishlistServiceMockRecorder
}

// MockIWishlistServiceMockRecorder is the mock recorder for MockIWishlistService.
type MockIWishlistServiceMockRecorder struct {
	mock *MockIWishlistService
}

// NewMockIWishlistService creates a new mock instance.
func NewMockIWishlistService(ctrl *gomock.Controller) *MockIWishlistService {
	mock := &MockIWishlistService{ctrl: ctrl}
	mock.recorder = &MockIWishlistServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIWishlistService) EXPECT() *MockIWishlistServiceMockRecorder {
	return m.recorder
}

// CreateWishlist mocks base method.
func (m *MockIWishlistService) CreateWishlist(arg0 *request.Wishlist) (response.Wishlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWishlist", arg0)
	ret0, _ := ret[0].(response.Wishlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWishlist indicates an expected call of CreateWishlist.
func (mr *MockIWishlistServiceMockRecorder) CreateWishlist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWishlist", reflect.TypeOf((*MockIWishlistService)(nil).CreateWishlist), arg0)  
}

// DeleteWishlist mocks base method.
func (m *MockIWishlistService) DeleteWishlist(arg0 string) (response.Wishlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWishlist", arg0)
	ret0, _ := ret[0].(response.Wishlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteWishlist indicates an expected call of DeleteWishlist.
func (mr *MockIWishlistServiceMockRecorder) DeleteWishlist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWishlist", reflect.TypeOf((*MockIWishlistService)(nil).DeleteWishlist), arg0)  
}

// GetAllWishlist mocks base method.
func (m *MockIWishlistService) GetAllWishlist(arg0 string) ([]response.Wishlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllWishlist", arg0)
	ret0, _ := ret[0].([]response.Wishlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllWishlist indicates an expected call of GetAllWishlist.
func (mr *MockIWishlistServiceMockRecorder) GetAllWishlist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllWishlist", reflect.TypeOf((*MockIWishlistService)(nil).GetAllWishlist), arg0)  
}

// GetWishlist mocks base method.
func (m *MockIWishlistService) GetWishlist(arg0 string) (response.Wishlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWishlist", arg0)
	ret0, _ := ret[0].(response.Wishlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWishlist indicates an expected call of GetWishlist.
func (mr *MockIWishlistServiceMockRecorder) GetWishlist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWishlist", reflect.TypeOf((*MockIWishlistService)(nil).GetWishlist), arg0)        
}

// UpdateWishlist mocks base method.
func (m *MockIWishlistService) UpdateWishlist(arg0 string, arg1 request.Wishlist) (response.Wishlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWishlist", arg0, arg1)
	ret0, _ := ret[0].(response.Wishlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWishlist indicates an expected call of UpdateWishlist.
func (mr *MockIWishlistServiceMockRecorder) UpdateWishlist(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWishlist", reflect.TypeOf((*MockIWishlistService)(nil).UpdateWishlist), arg0, arg1)
}