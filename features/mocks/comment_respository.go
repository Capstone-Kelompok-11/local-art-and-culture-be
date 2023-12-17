package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockICommentRepository is a mock of ICommentRepository interface.
type MockICommentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICommentRepositoryMockRecorder
}

// MockICommentRepositoryMockRecorder is the mock recorder for MockICommentRepository.
type MockICommentRepositoryMockRecorder struct {
	mock *MockICommentRepository
}

// NewMockICommentRepository creates a new mock instance.
func NewMockICommentRepository(ctrl *gomock.Controller) *MockICommentRepository {
	mock := &MockICommentRepository{ctrl: ctrl}
	mock.recorder = &MockICommentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICommentRepository) EXPECT() *MockICommentRepositoryMockRecorder {
	return m.recorder
}

// CreateComment mocks base method.
func (m *MockICommentRepository) CreateComment(arg0 *request.Comment) (response.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", arg0)
	ret0, _ := ret[0].(response.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockICommentRepositoryMockRecorder) CreateComment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockICommentRepository)(nil).CreateComment), arg0)  
}

// DeleteComment mocks base method.
func (m *MockICommentRepository) DeleteComment(arg0 string) (response.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", arg0)
	ret0, _ := ret[0].(response.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockICommentRepositoryMockRecorder) DeleteComment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockICommentRepository)(nil).DeleteComment), arg0)  
}

// GetAllComment mocks base method.
func (m *MockICommentRepository) GetAllComment() ([]response.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllComment")
	ret0, _ := ret[0].([]response.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllComment indicates an expected call of GetAllComment.
func (mr *MockICommentRepositoryMockRecorder) GetAllComment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllComment", reflect.TypeOf((*MockICommentRepository)(nil).GetAllComment))        
}

// GetComment mocks base method.
func (m *MockICommentRepository) GetComment(arg0 string) (response.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComment", arg0)
	ret0, _ := ret[0].(response.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComment indicates an expected call of GetComment.
func (mr *MockICommentRepositoryMockRecorder) GetComment(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComment", reflect.TypeOf((*MockICommentRepository)(nil).GetComment), arg0)        
}

// UpdateComment mocks base method.
func (m *MockICommentRepository) UpdateComment(arg0 string, arg1 request.Comment) (response.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateComment", arg0, arg1)
	ret0, _ := ret[0].(response.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateComment indicates an expected call of UpdateComment.
func (mr *MockICommentRepositoryMockRecorder) UpdateComment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComment", reflect.TypeOf((*MockICommentRepository)(nil).UpdateComment), arg0, arg1)
}