package mocks

import (
	models "lokasani/entity/models"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockISaveService is a mock of ISaveService interface.
type MockISaveService struct {
	ctrl     *gomock.Controller
	recorder *MockISaveServiceMockRecorder
}

// MockISaveServiceMockRecorder is the mock recorder for MockISaveService.
type MockISaveServiceMockRecorder struct {
	mock *MockISaveService
}

// NewMockISaveService creates a new mock instance.
func NewMockISaveService(ctrl *gomock.Controller) *MockISaveService {
	mock := &MockISaveService{ctrl: ctrl}
	mock.recorder = &MockISaveServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISaveService) EXPECT() *MockISaveServiceMockRecorder {
	return m.recorder
}

// GetAllChatbot mocks base method.
func (m *MockISaveService) GetAllChatbot(arg0 uint) ([]*response.SaveChatbot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllChatbot", arg0)
	ret0, _ := ret[0].([]*response.SaveChatbot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllChatbot indicates an expected call of GetAllChatbot.
func (mr *MockISaveServiceMockRecorder) GetAllChatbot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllChatbot", reflect.TypeOf((*MockISaveService)(nil).GetAllChatbot), arg0)        
}

// SaveChatbot mocks base method.
func (m *MockISaveService) SaveChatbot(arg0 models.SaveChatbot) (response.SaveChatbot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveChatbot", arg0)
	ret0, _ := ret[0].(response.SaveChatbot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveChatbot indicates an expected call of SaveChatbot.
func (mr *MockISaveServiceMockRecorder) SaveChatbot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveChatbot", reflect.TypeOf((*MockISaveService)(nil).SaveChatbot), arg0)
}