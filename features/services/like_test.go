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

func TestUpdateLike_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLikeRepository := mocks.NewMockILikeRepository(ctrl)

	expectedUpdatedLike := response.Like{
		Id:      1,
		UserId:  1,
		SourceId: 1,
		SourceStr: "Like",
		Active:   true,
	}

	mockLikeRepository.EXPECT().UpdateLike(gomock.Any()).Return(expectedUpdatedLike, nil).Times(1)

	likeService := NewLikeService(mockLikeRepository)

	requestLike := request.Like{
		UserId:  1,
		SourceId: 1,
		SourceStr: "Like",
		Active:   true,
	}

	resultLike, err := likeService.UpdateLike(requestLike)

	assert.NoError(t, err)
	assert.Equal(t, expectedUpdatedLike, resultLike)
}

func TestUpdateLike_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLikeRepository := mocks.NewMockILikeRepository(ctrl)

	expectedError := errors.New("update like failed")
	mockLikeRepository.EXPECT().UpdateLike(gomock.Any()).Return(response.Like{}, expectedError).Times(1)

	likeService := NewLikeService(mockLikeRepository)

	requestLike := request.Like{
		UserId:  1,
		SourceId: 1,
		SourceStr: "Like",
		Active:   true,
	}

	resultLike, err := likeService.UpdateLike(requestLike)

	assert.EqualError(t, err, expectedError.Error())
	assert.Equal(t, response.Like{}, resultLike)
}

func TestUpdateLike_Failure_SourceID(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockLikeRepository := mocks.NewMockILikeRepository(ctrl)

    expectedError := errors.New("source is empty`")

    LikeService := NewLikeService(mockLikeRepository)

    _, err := LikeService.UpdateLike(request.Like{})

    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }
}

func TestCreateLike_Failure_SourceStr(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockLikeRepository := mocks.NewMockILikeRepository(ctrl)

    mockRequestLike := request.Like{
        UserId:  1,
		SourceId: 1,
		SourceStr: "",
		Active:   true,
    }

    expectedError := errors.New("source is empty`")

    LikeService := NewLikeService(mockLikeRepository)

    result, err := LikeService.UpdateLike(mockRequestLike)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.UserId != 0 || result.SourceId != 0 || result.SourceStr != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestCreateLike_Failure_UserId(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockLikeRepository := mocks.NewMockILikeRepository(ctrl)

    mockRequestLike := request.Like{
        UserId:  0,
		SourceId: 1,
		SourceStr: "Like",
		Active:   true,
    }

    expectedError := errors.New("can't find any user with this id")

    LikeService := NewLikeService(mockLikeRepository)

    result, err := LikeService.UpdateLike(mockRequestLike)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.UserId != 0 || result.SourceId != 0 || result.SourceStr != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}