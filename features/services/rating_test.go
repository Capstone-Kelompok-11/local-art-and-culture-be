package services

import (
	"errors"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/mocks"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCreateRating_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRatingRepository := mocks.NewMockIRatingRepository(ctrl)

	mockRequestRating := &request.Rating{
		Rating: 10,
		Ulasan: "test",
		UserId: 1,
		ProductId: 1,
	}

	expectedResponseRating := response.Rating{
		Id:          1,
		Rating: 10,
		Ulasan: "test",
		UserId: 1,
		ProductId: 1,
		Users: response.Users{},
		Product: response.Products{},
	}
	mockRatingRepository.EXPECT().CreateRating(mockRequestRating).Return(expectedResponseRating, nil)

	RatingService := NewRatingService(mockRatingRepository)
	result, err := RatingService.CreateRating(mockRequestRating)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if result.Id != expectedResponseRating.Id || result.Rating!= expectedResponseRating.Rating|| result.Ulasan != expectedResponseRating.Ulasan || result.UserId != expectedResponseRating.UserId || result.ProductId != expectedResponseRating.ProductId {
		t.Errorf("Expected response %v, but got %v", expectedResponseRating, result)
	}
}

func TestCreateRating_Failure_RatingIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRatingRepository := mocks.NewMockIRatingRepository(ctrl)

    mockRequestRating := &request.Rating{
        Rating: 0,
		Ulasan: "test",
		UserId: 1,
		ProductId: 1,
    }

    expectedError := errors.New("rating is empty")

    RatingService := NewRatingService(mockRatingRepository)

    result, err := RatingService.CreateRating(mockRequestRating)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.Rating != 0 || result.UserId != 0 || result.ProductId != 0 || result.Ulasan != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestCreateRating_Failure_UlasanIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRatingRepository := mocks.NewMockIRatingRepository(ctrl)

    mockRequestRating := &request.Rating{
        Rating: 10,
		Ulasan: "",
		UserId: 1,
		ProductId: 1,
    }

    expectedError := errors.New("feedback is empty")

    RatingService := NewRatingService(mockRatingRepository)

    result, err := RatingService.CreateRating(mockRequestRating)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.Rating != 0 || result.UserId != 0 || result.ProductId != 0 || result.Ulasan != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestCreateRating_Failure_DatabaseError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRatingRepository := mocks.NewMockIRatingRepository(ctrl)

    mockRequestRating := &request.Rating{
		Rating: 10,
		Ulasan: "test",
		UserId: 1,
		ProductId: 1,
    }

    expectedError := errors.New("failed to create new Rating")

    mockRatingRepository.EXPECT().CreateRating(mockRequestRating).Return(response.Rating{}, expectedError)

    RatingService := NewRatingService(mockRatingRepository)

    result, _ := RatingService.CreateRating(mockRequestRating)

    if result.Id != 0 || result.Rating != 0 || result.UserId != 0 || result.ProductId != 0 || result.Ulasan != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestGetAllRating_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRatingRepository := mocks.NewMockIRatingRepository(ctrl)

	expectedRatings := []response.Rating{
		{Id: 1, Rating: 10, Ulasan: "test1", UserId: 1, ProductId: 1},
		{Id: 2, Rating: 20, Ulasan: "test2", UserId: 2, ProductId: 2},
	}

	var expectedError error

	mockRatingRepository.EXPECT().GetAllRating("filter").Return(expectedRatings, expectedError).Times(1)

	RatingService := NewRatingService(mockRatingRepository)

	resultRatings, err := RatingService.GetAllRating("filter")

	if err != expectedError {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultRatings, expectedRatings) {
		t.Errorf("Expected Ratings %+v, but got %+v", expectedRatings, resultRatings)
	}
}

func TestGetAllRating_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRatingRepository := mocks.NewMockIRatingRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockRatingRepository.EXPECT().GetAllRating("filter").Return(nil, expectedError).Times(1)

	RatingService := NewRatingService(mockRatingRepository)

	resultRatings, err := RatingService.GetAllRating("filter")

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if resultRatings != nil {
		t.Errorf("Expected nil Ratings, but got %+v", resultRatings)
	}
}

func TestGetRating_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRatingRepository := mocks.NewMockIRatingRepository(ctrl)

	expectedRating := response.Rating{
		Rating: 10,
		Ulasan: "test",
		UserId: 1,
		ProductId: 1,
	}

	mockRatingRepository.EXPECT().GetRating("RatingID").Return(expectedRating, nil).Times(1)

	RatingService := NewRatingService(mockRatingRepository)

	resultRating, err := RatingService.GetRating("RatingID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultRating, expectedRating) {
		t.Errorf("Expected Rating %+v, but got %+v", expectedRating, resultRating)
	}
}

func TestGetRating_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRatingRepository := mocks.NewMockIRatingRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockRatingRepository.EXPECT().GetRating("RatingID").Return(response.Rating{}, expectedError).Times(1)

	RatingService := NewRatingService(mockRatingRepository)

	resultRating, err := RatingService.GetRating("RatingID")

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if resultRating.Id != 0 || resultRating.Ulasan != "" || resultRating.UserId != 0 || resultRating.ProductId != 0 {
		t.Errorf("Expected empty result, but got %+v", resultRating)
	}
}

func TestGetRating_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRatingRepository := mocks.NewMockIRatingRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	RatingService := NewRatingService(mockRatingRepository)

	_, err := RatingService.GetRating("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestUpdateRating_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRatingRepository := mocks.NewMockIRatingRepository(ctrl)

	expectedUpdatedRating := response.Rating{
		Id:          1,
		Rating: 10,
		Ulasan: "test",
		UserId: 1,
		ProductId: 1,
	}

	var expectedError error

	mockRatingRepository.EXPECT().UpdateRating("RatingID", gomock.Any()).Return(expectedUpdatedRating, expectedError).Times(1)

	RatingService := NewRatingService(mockRatingRepository)

	requestRating := request.Rating{
		Rating: 10,
		Ulasan: "test",
		UserId: 1,
		ProductId: 1,
	}

	resultRating, err := RatingService.UpdateRating("RatingID", requestRating)

	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if !reflect.DeepEqual(resultRating, expectedUpdatedRating) {
		t.Errorf("Expected %+v, but got %+v", expectedUpdatedRating, resultRating)
	}
}

func TestUpdateRating_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRatingRepository := mocks.NewMockIRatingRepository(ctrl)

	expectedError := errors.New("failed to update data")

	mockRatingRepository.EXPECT().UpdateRating("RatingID", gomock.Any()).Return(response.Rating{}, expectedError).Times(1)

	RatingService := NewRatingService(mockRatingRepository)

	requestRating := request.Rating{
		Rating: 10,
		Ulasan: "test",
		UserId: 1,
		ProductId: 1,
	}

	resultRating, err := RatingService.UpdateRating("RatingID", requestRating)

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if resultRating.Id != 0 || resultRating.Ulasan != "" || resultRating.UserId != 0 || resultRating.ProductId != 0 {
		t.Errorf("Expected empty result, but got %+v", resultRating)
	}
}

func TestUpdateRating_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRatingRepository := mocks.NewMockIRatingRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	RatingService := NewRatingService(mockRatingRepository)

	_, err := RatingService.UpdateRating("", request.Rating{})

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestDeleteRating_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRatingRepository := mocks.NewMockIRatingRepository(ctrl)

	expectedRating := response.Rating{
		Id:     1,
		Rating: 10,
		Ulasan: "test",
		UserId: 1,
		ProductId: 1,
	}

	mockRatingRepository.EXPECT().DeleteRating("RatingID").Return(expectedRating, nil).Times(1)

	RatingService := NewRatingService(mockRatingRepository)

	resultRating, err := RatingService.DeleteRating("RatingID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultRating, expectedRating) {
		t.Errorf("Expected Rating %+v, but got %+v", expectedRating, resultRating)
	}
}

func TestDeleteRating_Failure_ErrDelete(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRatingRepository := mocks.NewMockIRatingRepository(ctrl)

    expectedError := errors.New("failed to delete data from database")

    mockRatingRepository.EXPECT().DeleteRating(gomock.Any()).Return(response.Rating{}, errors.New("failed to delete data from database"))

    RatingService := NewRatingService(mockRatingRepository)

    result, err := RatingService.DeleteRating("RatingID")

    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error %v, but got %v", expectedError, err)
    }

    if result.Id != 0 || result.Ulasan != "" || result.UserId != 0 || result.ProductId != 0 || result.Rating != 0  {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeleteRating_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRatingRepository := mocks.NewMockIRatingRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	RatingService := NewRatingService(mockRatingRepository)

	_, err := RatingService.DeleteRating("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}