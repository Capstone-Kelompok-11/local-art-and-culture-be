package services

import (
	"errors"
	//"lokasani/entity/request"
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

	mockLikeRepository.EXPECT().GetAllLike("LikeID").Return(expectedLikes, nil).Times(1)

	likeService := NewLikeService(mockLikeRepository)

	resultLikes, err := likeService.GetAllLike("LikeID")

	assert.NoError(t, err)
	assert.Equal(t, expectedLikes, resultLikes)
}

func TestGetAllLike_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLikeRepository := mocks.NewMockILikeRepository(ctrl)

	expectedError := errors.New("some error")

	mockLikeRepository.EXPECT().GetAllLike("LikeID").Return(nil, expectedError).Times(1)

	likeService := NewLikeService(mockLikeRepository)

	resultLikes, err := likeService.GetAllLike("LikeID")

	assert.Error(t, err)
	assert.Nil(t, resultLikes)
	assert.NotEqual(t, err, expectedError.Error())
}

// func TestUpdateLike_Success(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockLikeRepository := mocks.NewMockILikeRepository(ctrl)

// 	expectedLike := response.Like{
// 		Id:     1,
// 		UserId: 1,
// 	}

// 	mockLikeRepository.EXPECT().UpdateLike(gomock.Any()).Return(expectedLike, nil).Times(1)

// 	likeService := NewLikeService(mockLikeRepository)

// 	requestLike := request.Like{
// 		UserId: 1,
// 	}

// 	_, err := likeService.UpdateLike(requestLike)

// 	assert.NotNil(t, err)
// 	//assert.Equal(t, expectedLike, resultLike)
// }