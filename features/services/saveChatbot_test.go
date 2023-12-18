package services

import (
	"errors"
	"lokasani/entity/models"
	"lokasani/entity/response"
	"lokasani/features/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllChatbot_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSaveRepository := mocks.NewMockISaveRepository(ctrl)

	expectedResponse := []*response.SaveChatbot{
		{Id: 1, Message: "Chatbot1"},
		{Id: 2, Message: "Chatbot2"},
	}

	mockSaveRepository.EXPECT().GetAllChatbot(gomock.Any()).Return(expectedResponse, nil).Times(1)

	saveService := NewSaveService(mockSaveRepository)

	userID := uint(1)
	result, err := saveService.GetAllChatbot(userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedResponse, result)
}

func TestGetAllChatbot_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSaveRepository := mocks.NewMockISaveRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockSaveRepository.EXPECT().GetAllChatbot(gomock.Any()).Return(nil, expectedError).Times(1)

	saveService := NewSaveService(mockSaveRepository)

	userID := uint(1)
	result, err := saveService.GetAllChatbot(userID)

	assert.EqualError(t, err, expectedError.Error())
	assert.Empty(t, result)
}

func TestUpdateChatbot_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSaveRepository := mocks.NewMockISaveRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	saveService := NewSaveService(mockSaveRepository)

	_, err := saveService.GetAllChatbot(0)

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestSaveChatbot_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSaveRepository := mocks.NewMockISaveRepository(ctrl)

	expectedInput := models.SaveChatbot{
		Message: "berikan rekomendasi event local",
	}
	expectedOutput := response.SaveChatbot{
		Id: 1,
		Message: "success",
		Response : "halo, berikut rekomendasi event local",
	}

	mockSaveRepository.EXPECT().SaveChatbot(expectedInput).Return(expectedOutput, nil).Times(1)

	saveService := NewSaveService(mockSaveRepository)

	result, err := saveService.SaveChatbot(expectedInput)

	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, result)
}