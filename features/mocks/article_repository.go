package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIArticleRepository is a mock of IArticleRepository interface.
type MockIArticleRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIArticleRepositoryMockRecorder
}

// MockIArticleRepositoryMockRecorder is the mock recorder for MockIArticleRepository.
type MockIArticleRepositoryMockRecorder struct {
	mock *MockIArticleRepository
}

// NewMockIArticleRepository creates a new mock instance.
func NewMockIArticleRepository(ctrl *gomock.Controller) *MockIArticleRepository {
	mock := &MockIArticleRepository{ctrl: ctrl}
	mock.recorder = &MockIArticleRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIArticleRepository) EXPECT() *MockIArticleRepositoryMockRecorder {
	return m.recorder
}

// CreateArticle mocks base method.
func (m *MockIArticleRepository) CreateArticle(arg0 *request.Article) (response.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArticle", arg0)
	ret0, _ := ret[0].(response.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateArticle indicates an expected call of CreateArticle.
func (mr *MockIArticleRepositoryMockRecorder) CreateArticle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArticle", reflect.TypeOf((*MockIArticleRepository)(nil).CreateArticle), arg0)
}

// DeleteArticle mocks base method.
func (m *MockIArticleRepository) DeleteArticle(arg0 string) (response.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteArticle", arg0)
	ret0, _ := ret[0].(response.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteArticle indicates an expected call of DeleteArticle.
func (mr *MockIArticleRepositoryMockRecorder) DeleteArticle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteArticle", reflect.TypeOf((*MockIArticleRepository)(nil).DeleteArticle), arg0)
}

// GetAllArticle mocks base method.
func (m *MockIArticleRepository) GetAllArticle(arg0 string, arg1, arg2 int) ([]response.Article, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllArticle", arg0, arg1, arg2)
	ret0, _ := ret[0].([]response.Article)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllArticle indicates an expected call of GetAllArticle.
func (mr *MockIArticleRepositoryMockRecorder) GetAllArticle(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllArticle", reflect.TypeOf((*MockIArticleRepository)(nil).GetAllArticle), arg0, arg1, arg2)
}

// GetArticle mocks base method.
func (m *MockIArticleRepository) GetArticle(arg0 string) (response.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticle", arg0)
	ret0, _ := ret[0].(response.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticle indicates an expected call of GetArticle.
func (mr *MockIArticleRepositoryMockRecorder) GetArticle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticle", reflect.TypeOf((*MockIArticleRepository)(nil).GetArticle), arg0)
}

// GetTotalLikes mocks base method.
func (m *MockIArticleRepository) GetTotalLikes(arg0 uint, arg1 string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalLikes", arg0, arg1)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalLikes indicates an expected call of GetTotalLikes.
func (mr *MockIArticleRepositoryMockRecorder) GetTotalLikes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalLikes", reflect.TypeOf((*MockIArticleRepository)(nil).GetTotalLikes), arg0, arg1)
}

// GetTotalLikesLastTwoWeeks mocks base method.
func (m *MockIArticleRepository) GetTotalLikesLastTwoWeeks(arg0 uint, arg1 string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTotalLikesLastTwoWeeks", arg0, arg1)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTotalLikesLastTwoWeeks indicates an expected call of GetTotalLikesLastTwoWeeks.
func (mr *MockIArticleRepositoryMockRecorder) GetTotalLikesLastTwoWeeks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTotalLikesLastTwoWeeks", reflect.TypeOf((*MockIArticleRepository)(nil).GetTotalLikesLastTwoWeeks), arg0, arg1)
}

// GetTrendingArticle mocks base method.
func (m *MockIArticleRepository) GetTrendingArticle(arg0 string, arg1, arg2 int) ([]response.Article, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrendingArticle", arg0, arg1, arg2)
	ret0, _ := ret[0].([]response.Article)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTrendingArticle indicates an expected call of GetTrendingArticle.
func (mr *MockIArticleRepositoryMockRecorder) GetTrendingArticle(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrendingArticle", reflect.TypeOf((*MockIArticleRepository)(nil).GetTrendingArticle), arg0, arg1, arg2)   
}

// UpdateArticle mocks base method.
func (m *MockIArticleRepository) UpdateArticle(arg0 string, arg1 request.Article) (response.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateArticle", arg0, arg1)
	ret0, _ := ret[0].(response.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateArticle indicates an expected call of UpdateArticle.
func (mr *MockIArticleRepositoryMockRecorder) UpdateArticle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateArticle", reflect.TypeOf((*MockIArticleRepository)(nil).UpdateArticle), arg0, arg1)
}