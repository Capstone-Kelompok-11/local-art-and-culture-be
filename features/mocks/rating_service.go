package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIRatingService is a mock of IRatingService interface.
type MockIRatingService struct {
	ctrl     *gomock.Controller
	recorder *MockIRatingServiceMockRecorder
}

// MockIRatingServiceMockRecorder is the mock recorder for MockIRatingService.
type MockIRatingServiceMockRecorder struct {
	mock *MockIRatingService
}

// NewMockIRatingService creates a new mock instance.
func NewMockIRatingService(ctrl *gomock.Controller) *MockIRatingService {
	mock := &MockIRatingService{ctrl: ctrl}
	mock.recorder = &MockIRatingServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRatingService) EXPECT() *MockIRatingServiceMockRecorder {
	return m.recorder
}

// CreateRating mocks base method.
func (m *MockIRatingService) CreateRating(arg0 *request.Rating) (response.Rating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRating", arg0)
	ret0, _ := ret[0].(response.Rating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateRating indicates an expected call of CreateRating.
func (mr *MockIRatingServiceMockRecorder) CreateRating(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRating", reflect.TypeOf((*MockIRatingService)(nil).CreateRating), arg0)        
}

// DeleteRating mocks base method.
func (m *MockIRatingService) DeleteRating(arg0 string) (response.Rating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRating", arg0)
	ret0, _ := ret[0].(response.Rating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteRating indicates an expected call of DeleteRating.
func (mr *MockIRatingServiceMockRecorder) DeleteRating(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRating", reflect.TypeOf((*MockIRatingService)(nil).DeleteRating), arg0)        
}

// GetAllRating mocks base method.
func (m *MockIRatingService) GetAllRating(arg0 string) ([]response.Rating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllRating", arg0)
	ret0, _ := ret[0].([]response.Rating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllRating indicates an expected call of GetAllRating.
func (mr *MockIRatingServiceMockRecorder) GetAllRating(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllRating", reflect.TypeOf((*MockIRatingService)(nil).GetAllRating), arg0)        
}

// GetRating mocks base method.
func (m *MockIRatingService) GetRating(arg0 string) (response.Rating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRating", arg0)
	ret0, _ := ret[0].(response.Rating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRating indicates an expected call of GetRating.
func (mr *MockIRatingServiceMockRecorder) GetRating(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRating", reflect.TypeOf((*MockIRatingService)(nil).GetRating), arg0)
}

// UpdateRating mocks base method.
func (m *MockIRatingService) UpdateRating(arg0 string, arg1 request.Rating) (response.Rating, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRating", arg0, arg1)
	ret0, _ := ret[0].(response.Rating)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateRating indicates an expected call of UpdateRating.
func (mr *MockIRatingServiceMockRecorder) UpdateRating(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRating", reflect.TypeOf((*MockIRatingService)(nil).UpdateRating), arg0, arg1)  
}