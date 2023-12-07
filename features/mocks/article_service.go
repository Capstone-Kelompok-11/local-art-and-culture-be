package mocks

import (
	request "lokasani/entity/request"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIArticleService is a mock of IArticleService interface.
type MockIArticleService struct {
	ctrl     *gomock.Controller
	recorder *MockIArticleServiceMockRecorder
}

// MockIArticleServiceMockRecorder is the mock recorder for MockIArticleService.
type MockIArticleServiceMockRecorder struct {
	mock *MockIArticleService
}

// NewMockIArticleService creates a new mock instance.
func NewMockIArticleService(ctrl *gomock.Controller) *MockIArticleService {
	mock := &MockIArticleService{ctrl: ctrl}
	mock.recorder = &MockIArticleServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIArticleService) EXPECT() *MockIArticleServiceMockRecorder {
	return m.recorder
}

// CalculatePaginationValues mocks base method.
func (m *MockIArticleService) CalculatePaginationValues(arg0, arg1, arg2 int) (int, int) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculatePaginationValues", arg0, arg1, arg2)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// CalculatePaginationValues indicates an expected call of CalculatePaginationValues.
func (mr *MockIArticleServiceMockRecorder) CalculatePaginationValues(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculatePaginationValues", reflect.TypeOf((*MockIArticleService)(nil).CalculatePaginationValues), arg0, arg1, arg2)
}

// CreateArticle mocks base method.
func (m *MockIArticleService) CreateArticle(arg0 *request.Article) (response.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateArticle", arg0)
	ret0, _ := ret[0].(response.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateArticle indicates an expected call of CreateArticle.
func (mr *MockIArticleServiceMockRecorder) CreateArticle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateArticle", reflect.TypeOf((*MockIArticleService)(nil).CreateArticle), arg0)
}

// DeleteArticle mocks base method.
func (m *MockIArticleService) DeleteArticle(arg0 string) (response.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteArticle", arg0)
	ret0, _ := ret[0].(response.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteArticle indicates an expected call of DeleteArticle.
func (mr *MockIArticleServiceMockRecorder) DeleteArticle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteArticle", reflect.TypeOf((*MockIArticleService)(nil).DeleteArticle), arg0)
}

// GetAllArticle mocks base method.
func (m *MockIArticleService) GetAllArticle(arg0 string, arg1, arg2 int) ([]response.Article, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllArticle", arg0, arg1, arg2)
	ret0, _ := ret[0].([]response.Article)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllArticle indicates an expected call of GetAllArticle.
func (mr *MockIArticleServiceMockRecorder) GetAllArticle(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllArticle", reflect.TypeOf((*MockIArticleService)(nil).GetAllArticle), arg0, arg1, arg2)
}

// GetArticle mocks base method.
func (m *MockIArticleService) GetArticle(arg0 string) (response.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArticle", arg0)
	ret0, _ := ret[0].(response.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetArticle indicates an expected call of GetArticle.
func (mr *MockIArticleServiceMockRecorder) GetArticle(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArticle", reflect.TypeOf((*MockIArticleService)(nil).GetArticle), arg0)
}

// GetNextPage mocks base method.
func (m *MockIArticleService) GetNextPage(arg0, arg1 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNextPage", arg0, arg1)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetNextPage indicates an expected call of GetNextPage.
func (mr *MockIArticleServiceMockRecorder) GetNextPage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextPage", reflect.TypeOf((*MockIArticleService)(nil).GetNextPage), arg0, arg1)
}

// GetPrevPage mocks base method.
func (m *MockIArticleService) GetPrevPage(arg0 int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPrevPage", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// GetPrevPage indicates an expected call of GetPrevPage.
func (mr *MockIArticleServiceMockRecorder) GetPrevPage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPrevPage", reflect.TypeOf((*MockIArticleService)(nil).GetPrevPage), arg0)
}

// GetTrendingArticle mocks base method.
func (m *MockIArticleService) GetTrendingArticle(arg0 string, arg1, arg2 int) ([]response.Article, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrendingArticle", arg0, arg1, arg2)
	ret0, _ := ret[0].([]response.Article)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTrendingArticle indicates an expected call of GetTrendingArticle.
func (mr *MockIArticleServiceMockRecorder) GetTrendingArticle(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrendingArticle", reflect.TypeOf((*MockIArticleService)(nil).GetTrendingArticle), arg0, arg1, arg2)      
}

// UpdateArticle mocks base method.
func (m *MockIArticleService) UpdateArticle(arg0 string, arg1 request.Article) (response.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateArticle", arg0, arg1)
	ret0, _ := ret[0].(response.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateArticle indicates an expected call of UpdateArticle.
func (mr *MockIArticleServiceMockRecorder) UpdateArticle(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateArticle", reflect.TypeOf((*MockIArticleService)(nil).UpdateArticle), arg0, arg1)
}