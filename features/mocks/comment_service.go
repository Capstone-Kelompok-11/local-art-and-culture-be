package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockICommentService is a mock of ICommentService interface.
type MockICommentService struct {
	ctrl     *gomock.Controller
	recorder *MockICommentServiceMockRecorder
}

// MockICommentServiceMockRecorder is the mock recorder for MockICommentService.
type MockICommentServiceMockRecorder struct {
	mock *MockICommentService
}

// NewMockICommentService creates a new mock instance.
func NewMockICommentService(ctrl *gomock.Controller) *MockICommentService {
	mock := &MockICommentService{ctrl: ctrl}
	mock.recorder = &MockICommentServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICommentService) EXPECT() *MockICommentServiceMockRecorder {
	return m.recorder
}

// CreateComment mocks base method.
func (m *MockICommentService) CreateComment(arg0 *request.Comment) (response.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", arg0)
	ret0, _ := ret[0].(response.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockICommentServiceMockRecorder) CreateComment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockICommentService)(nil).CreateComment), arg0)     
}

// DeleteComment mocks base method.
func (m *MockICommentService) DeleteComment(arg0 string) (response.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", arg0)
	ret0, _ := ret[0].(response.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockICommentServiceMockRecorder) DeleteComment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockICommentService)(nil).DeleteComment), arg0)     
}

// GetAllComment mocks base method.
func (m *MockICommentService) GetAllComment() ([]response.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllComment")
	ret0, _ := ret[0].([]response.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllComment indicates an expected call of GetAllComment.
func (mr *MockICommentServiceMockRecorder) GetAllComment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllComment", reflect.TypeOf((*MockICommentService)(nil).GetAllComment))
}

// GetComment mocks base method.
func (m *MockICommentService) GetComment(arg0 string) (response.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComment", arg0)
	ret0, _ := ret[0].(response.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComment indicates an expected call of GetComment.
func (mr *MockICommentServiceMockRecorder) GetComment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComment", reflect.TypeOf((*MockICommentService)(nil).GetComment), arg0)
}

// UpdateComment mocks base method.
func (m *MockICommentService) UpdateComment(arg0 string, arg1 request.Comment) (response.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateComment", arg0, arg1)
	ret0, _ := ret[0].(response.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateComment indicates an expected call of UpdateComment.
func (mr *MockICommentServiceMockRecorder) UpdateComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComment", reflect.TypeOf((*MockICommentService)(nil).UpdateComment), arg0, arg1)
}