package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIRatingRepository is a mock of IRatingRepository interface.
type MockIRatingRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIRatingRepositoryMockRecorder
}

// MockIRatingRepositoryMockRecorder is the mock recorder for MockIRatingRepository.
type MockIRatingRepositoryMockRecorder struct {
	mock *MockIRatingRepository
}

// NewMockIRatingRepository creates a new mock instance.
func NewMockIRatingRepository(ctrl *gomock.Controller) *MockIRatingRepository {
	mock := &MockIRatingRepository{ctrl: ctrl}
	mock.recorder = &MockIRatingRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRatingRepository) EXPECT() *MockIRatingRepositoryMockRecorder {
	return m.recorder
}

// CreateRating mocks base method.
func (m *MockIRatingRepository) CreateRating(arg0 *request.Rating) (response.Rating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRating", arg0)
	ret0, _ := ret[0].(response.Rating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRating indicates an expected call of CreateRating.
func (mr *MockIRatingRepositoryMockRecorder) CreateRating(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRating", reflect.TypeOf((*MockIRatingRepository)(nil).CreateRating), arg0)     
}

// DeleteRating mocks base method.
func (m *MockIRatingRepository) DeleteRating(arg0 string) (response.Rating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRating", arg0)
	ret0, _ := ret[0].(response.Rating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRating indicates an expected call of DeleteRating.
func (mr *MockIRatingRepositoryMockRecorder) DeleteRating(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRating", reflect.TypeOf((*MockIRatingRepository)(nil).DeleteRating), arg0)     
}

// GetAllRating mocks base method.
func (m *MockIRatingRepository) GetAllRating(arg0 string) ([]response.Rating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRating", arg0)
	ret0, _ := ret[0].([]response.Rating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRating indicates an expected call of GetAllRating.
func (mr *MockIRatingRepositoryMockRecorder) GetAllRating(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRating", reflect.TypeOf((*MockIRatingRepository)(nil).GetAllRating), arg0)     
}

// GetRating mocks base method.
func (m *MockIRatingRepository) GetRating(arg0 string) (response.Rating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRating", arg0)
	ret0, _ := ret[0].(response.Rating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRating indicates an expected call of GetRating.
func (mr *MockIRatingRepositoryMockRecorder) GetRating(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRating", reflect.TypeOf((*MockIRatingRepository)(nil).GetRating), arg0)
}

// UpdateRating mocks base method.
func (m *MockIRatingRepository) UpdateRating(arg0 string, arg1 request.Rating) (response.Rating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRating", arg0, arg1)
	ret0, _ := ret[0].(response.Rating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRating indicates an expected call of UpdateRating.
func (mr *MockIRatingRepositoryMockRecorder) UpdateRating(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRating", reflect.TypeOf((*MockIRatingRepository)(nil).UpdateRating), arg0, arg1)
}