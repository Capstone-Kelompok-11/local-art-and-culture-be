package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockILikeRepository is a mock of ILikeRepository interface.
type MockILikeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockILikeRepositoryMockRecorder
}

// MockILikeRepositoryMockRecorder is the mock recorder for MockILikeRepository.
type MockILikeRepositoryMockRecorder struct {
	mock *MockILikeRepository
}

// NewMockILikeRepository creates a new mock instance.
func NewMockILikeRepository(ctrl *gomock.Controller) *MockILikeRepository {
	mock := &MockILikeRepository{ctrl: ctrl}
	mock.recorder = &MockILikeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockILikeRepository) EXPECT() *MockILikeRepositoryMockRecorder {
	return m.recorder
}

// GetAllLike mocks base method.
func (m *MockILikeRepository) GetAllLike(arg0 string) ([]response.Like, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllLike", arg0)
	ret0, _ := ret[0].([]response.Like)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllLike indicates an expected call of GetAllLike.
func (mr *MockILikeRepositoryMockRecorder) GetAllLike(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllLike", reflect.TypeOf((*MockILikeRepository)(nil).GetAllLike), arg0)
}

// UpdateLike mocks base method.
func (m *MockILikeRepository) UpdateLike(arg0 request.Like) (response.Like, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLike", arg0)
	ret0, _ := ret[0].(response.Like)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateLike indicates an expected call of UpdateLike.
func (mr *MockILikeRepositoryMockRecorder) UpdateLike(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLike", reflect.TypeOf((*MockILikeRepository)(nil).UpdateLike), arg0)
}