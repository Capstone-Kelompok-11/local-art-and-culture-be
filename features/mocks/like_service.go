package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockILikeService is a mock of ILikeService interface.
type MockILikeService struct {
	ctrl     *gomock.Controller
	recorder *MockILikeServiceMockRecorder
}

// MockILikeServiceMockRecorder is the mock recorder for MockILikeService.
type MockILikeServiceMockRecorder struct {
	mock *MockILikeService
}

// NewMockILikeService creates a new mock instance.
func NewMockILikeService(ctrl *gomock.Controller) *MockILikeService {
	mock := &MockILikeService{ctrl: ctrl}
	mock.recorder = &MockILikeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockILikeService) EXPECT() *MockILikeServiceMockRecorder {
	return m.recorder
}

// GetAllLike mocks base method.
func (m *MockILikeService) GetAllLike(arg0 string) ([]response.Like, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllLike", arg0)
	ret0, _ := ret[0].([]response.Like)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllLike indicates an expected call of GetAllLike.
func (mr *MockILikeServiceMockRecorder) GetAllLike(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllLike", reflect.TypeOf((*MockILikeService)(nil).GetAllLike), arg0)
}

// UpdateLike mocks base method.
func (m *MockILikeService) UpdateLike(arg0 request.Like) (response.Like, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLike", arg0)
	ret0, _ := ret[0].(response.Like)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateLike indicates an expected call of UpdateLike.
func (mr *MockILikeServiceMockRecorder) UpdateLike(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLike", reflect.TypeOf((*MockILikeService)(nil).UpdateLike), arg0)
}