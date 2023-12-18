package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIWishlistRepository is a mock of IWishlistRepository interface.
type MockIWishlistRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIWishlistRepositoryMockRecorder
}

// MockIWishlistRepositoryMockRecorder is the mock recorder for MockIWishlistRepository.
type MockIWishlistRepositoryMockRecorder struct {
	mock *MockIWishlistRepository
}

// NewMockIWishlistRepository creates a new mock instance.
func NewMockIWishlistRepository(ctrl *gomock.Controller) *MockIWishlistRepository {
	mock := &MockIWishlistRepository{ctrl: ctrl}
	mock.recorder = &MockIWishlistRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIWishlistRepository) EXPECT() *MockIWishlistRepositoryMockRecorder {
	return m.recorder
}

// CreateWishlist mocks base method.
func (m *MockIWishlistRepository) CreateWishlist(arg0 *request.Wishlist) (response.Wishlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWishlist", arg0)
	ret0, _ := ret[0].(response.Wishlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWishlist indicates an expected call of CreateWishlist.
func (mr *MockIWishlistRepositoryMockRecorder) CreateWishlist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWishlist", reflect.TypeOf((*MockIWishlistRepository)(nil).CreateWishlist), arg0)
}

// DeleteWishlist mocks base method.
func (m *MockIWishlistRepository) DeleteWishlist(arg0 string) (response.Wishlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteWishlist", arg0)
	ret0, _ := ret[0].(response.Wishlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteWishlist indicates an expected call of DeleteWishlist.
func (mr *MockIWishlistRepositoryMockRecorder) DeleteWishlist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteWishlist", reflect.TypeOf((*MockIWishlistRepository)(nil).DeleteWishlist), arg0)
}

// GetAllWishlist mocks base method.
func (m *MockIWishlistRepository) GetAllWishlist(arg0 string) ([]response.Wishlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllWishlist", arg0)
	ret0, _ := ret[0].([]response.Wishlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllWishlist indicates an expected call of GetAllWishlist.
func (mr *MockIWishlistRepositoryMockRecorder) GetAllWishlist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllWishlist", reflect.TypeOf((*MockIWishlistRepository)(nil).GetAllWishlist), arg0)
}

// GetWishlist mocks base method.
func (m *MockIWishlistRepository) GetWishlist(arg0 string) (response.Wishlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWishlist", arg0)
	ret0, _ := ret[0].(response.Wishlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWishlist indicates an expected call of GetWishlist.
func (mr *MockIWishlistRepositoryMockRecorder) GetWishlist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWishlist", reflect.TypeOf((*MockIWishlistRepository)(nil).GetWishlist), arg0)     
}

// UpdateWishlist mocks base method.
func (m *MockIWishlistRepository) UpdateWishlist(arg0 string, arg1 request.Wishlist) (response.Wishlist, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateWishlist", arg0, arg1)
	ret0, _ := ret[0].(response.Wishlist)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateWishlist indicates an expected call of UpdateWishlist.
func (mr *MockIWishlistRepositoryMockRecorder) UpdateWishlist(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateWishlist", reflect.TypeOf((*MockIWishlistRepository)(nil).UpdateWishlist), arg0, arg1)
}