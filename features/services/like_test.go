package services

import (
	"errors"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetAllLike_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockILikeService(ctrl)

	expectedLikes := []response.Like{{Id: 1, UserId: 1}, {Id: 2, UserId: 2}}
	mockService.EXPECT().GetAllLike(gomock.Any()).Return(expectedLikes, nil)

	likes, err := mockService.GetAllLike("articleID")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(likes) != len(expectedLikes) {
		t.Errorf("Expected %d likes, got %d", len(expectedLikes), len(likes))
	}
}

func TestGetAllLike_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockILikeService(ctrl)

	expectedError := errors.New("some error")
	mockService.EXPECT().GetAllLike(gomock.Any()).Return(nil, expectedError)

	likes, err := mockService.GetAllLike("articleID")

	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}

	if likes != nil {
		t.Errorf("Expected nil likes, got %v", likes)
	}
}

func TestUpdateLike_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockILikeService(ctrl)

	expectedUpdatedLike := response.Like{Id: 1, UserId: 1}
	mockService.EXPECT().UpdateLike(gomock.Any()).Return(expectedUpdatedLike, nil)

	updatedLike, err := mockService.UpdateLike(request.Like{UserId: 1})

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if updatedLike.Id != expectedUpdatedLike.Id {
		t.Errorf("Expected like ID %d, got %d", expectedUpdatedLike.Id, updatedLike.Id)
	}

	if updatedLike.UserId != expectedUpdatedLike.UserId {
		t.Errorf("Expected user ID %d, got %d", expectedUpdatedLike.UserId, updatedLike.UserId)
	}
}

func TestUpdateLike_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockILikeService(ctrl)

	expectedError := errors.New("some error")
	mockService.EXPECT().UpdateLike(gomock.Any()).Return(response.Like{}, expectedError)

	updatedLike, err := mockService.UpdateLike(request.Like{UserId: 1})

	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}

	if updatedLike.Id != 0 {
		t.Errorf("Expected empty like ID, got %d", updatedLike.Id)
	}

	if updatedLike.UserId != 0 {
		t.Errorf("Expected empty user ID, got %d", updatedLike.UserId)
	}
}