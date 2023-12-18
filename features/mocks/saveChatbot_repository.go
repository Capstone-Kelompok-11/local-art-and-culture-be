package mocks

import (
	models "lokasani/entity/models"
	response "lokasani/entity/response"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockISaveRepository is a mock of ISaveRepository interface.
type MockISaveRepository struct {
	ctrl     *gomock.Controller
	recorder *MockISaveRepositoryMockRecorder
}

// MockISaveRepositoryMockRecorder is the mock recorder for MockISaveRepository.
type MockISaveRepositoryMockRecorder struct {
	mock *MockISaveRepository
}

// NewMockISaveRepository creates a new mock instance.
func NewMockISaveRepository(ctrl *gomock.Controller) *MockISaveRepository {
	mock := &MockISaveRepository{ctrl: ctrl}
	mock.recorder = &MockISaveRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISaveRepository) EXPECT() *MockISaveRepositoryMockRecorder {
	return m.recorder
}

// GetAllChatbot mocks base method.
func (m *MockISaveRepository) GetAllChatbot(arg0 uint) ([]*response.SaveChatbot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllChatbot", arg0)
	ret0, _ := ret[0].([]*response.SaveChatbot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllChatbot indicates an expected call of GetAllChatbot.
func (mr *MockISaveRepositoryMockRecorder) GetAllChatbot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllChatbot", reflect.TypeOf((*MockISaveRepository)(nil).GetAllChatbot), arg0)     
}

// SaveChatbot mocks base method.
func (m *MockISaveRepository) SaveChatbot(arg0 models.SaveChatbot) (response.SaveChatbot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveChatbot", arg0)
	ret0, _ := ret[0].(response.SaveChatbot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveChatbot indicates an expected call of SaveChatbot.
func (mr *MockISaveRepositoryMockRecorder) SaveChatbot(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveChatbot", reflect.TypeOf((*MockISaveRepository)(nil).SaveChatbot), arg0)
}