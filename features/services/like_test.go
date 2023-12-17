package services

import (
	"errors"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllLike_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLikeRepository := mocks.NewMockILikeRepository(ctrl)

	expectedLikes := []response.Like{
		{Id: 1, UserId: 1},
		{Id: 2, UserId: 2},
	}

	mockLikeRepository.EXPECT().GetAllLike("articleID").Return(expectedLikes, nil).Times(1)

	likeService := NewLikeService(mockLikeRepository)

	resultLikes, err := likeService.GetAllLike("articleID")

	assert.NoError(t, err)
	assert.Equal(t, expectedLikes, resultLikes)
}

func TestGetAllLike_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLikeRepository := mocks.NewMockILikeRepository(ctrl)

	expectedError := errors.New("some error")

	mockLikeRepository.EXPECT().GetAllLike("articleID").Return(nil, expectedError).Times(1)

	likeService := NewLikeService(mockLikeRepository)

	resultLikes, err := likeService.GetAllLike("articleID")

	assert.Error(t, err)
	assert.Nil(t, resultLikes)
	assert.NotEqual(t, err, expectedError.Error())
}

func TestUpdateLike_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLikeRepository := mocks.NewMockILikeRepository(ctrl)

	expectedLike := response.Like{
		Id:     1,
		UserId: 1,
	}

	mockLikeRepository.EXPECT().UpdateLike(gomock.Any()).Return(expectedLike, nil).Times(1)

	likeService := NewLikeService(mockLikeRepository)

	requestLike := request.Like{
		UserId: 1,
	}

	resultLike, _ := likeService.UpdateLike(requestLike)

	assert.NotNil(t, expectedLike, resultLike)
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