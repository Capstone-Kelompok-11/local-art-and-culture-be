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

func TestCreateWishlist_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWishlistRepository := mocks.NewMockIWishlistRepository(ctrl)

	mockRequestWishlist := &request.Wishlist{
		Quantity: 1,
		ProductId: 1,
		TicketId: 1,
	}

	expectedResponseWishlist := response.Wishlist{
		Id:          1,
		Quantity: 1,
		ProductId: 1,
		TicketId: 1,
	}
	mockWishlistRepository.EXPECT().CreateWishlist(mockRequestWishlist).Return(expectedResponseWishlist, nil)

	WishlistService := NewWishlistService(mockWishlistRepository)
	result, err := WishlistService.CreateWishlist(mockRequestWishlist)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if result.Id != expectedResponseWishlist.Id || result.Quantity != expectedResponseWishlist.Quantity || result.ProductId != expectedResponseWishlist.ProductId || result.TicketId != expectedResponseWishlist.TicketId {
		t.Errorf("Expected response %v, but got %v", expectedResponseWishlist, result)
	}
}

func TestCreateWishlist_Failure_QtyIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockWishlistRepository := mocks.NewMockIWishlistRepository(ctrl)

    mockRequestWishlist := &request.Wishlist{
        Quantity: 0,
		ProductId: 1,
		TicketId: 1,
    }

    expectedError := errors.New("quantity is empty")

    WishlistService := NewWishlistService(mockWishlistRepository)

    result, err := WishlistService.CreateWishlist(mockRequestWishlist)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.Quantity != 0 || result.ProductId != 0 || result.TicketId != 0 {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestCreateWishlist_Failure_DatabaseError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockWishlistRepository := mocks.NewMockIWishlistRepository(ctrl)

    mockRequestWishlist := &request.Wishlist{
		Quantity: 1,
		ProductId: 1,
		TicketId: 1,
    }

    expectedError := errors.New("failed to create new Wishlist")

    mockWishlistRepository.EXPECT().CreateWishlist(mockRequestWishlist).Return(response.Wishlist{}, expectedError)

    WishlistService := NewWishlistService(mockWishlistRepository)

    result, _ := WishlistService.CreateWishlist(mockRequestWishlist)

	if result.Id != 0 || result.Quantity != 0 || result.ProductId != 0 || result.TicketId != 0 {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeleteWishlist_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWishlistRepository := mocks.NewMockIWishlistRepository(ctrl)

	expectedWishlist := response.Wishlist{
		Id:     1,
		Quantity: 1,
		ProductId: 1,
		TicketId: 1,
	}

	mockWishlistRepository.EXPECT().DeleteWishlist("WishlistID").Return(expectedWishlist, nil).Times(1)

	WishlistService := NewWishlistService(mockWishlistRepository)

	resultWishlist, err := WishlistService.DeleteWishlist("WishlistID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultWishlist, expectedWishlist) {
		t.Errorf("Expected Wishlist %+v, but got %+v", expectedWishlist, resultWishlist)
	}
}

func TestDeleteWishlist_Failure_ErrDelete(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockWishlistRepository := mocks.NewMockIWishlistRepository(ctrl)

    expectedError := errors.New("failed to delete data from database")

    mockWishlistRepository.EXPECT().DeleteWishlist(gomock.Any()).Return(response.Wishlist{}, errors.New("failed to delete data from database"))

    WishlistService := NewWishlistService(mockWishlistRepository)

    result, err := WishlistService.DeleteWishlist("WishlistID")

    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error %v, but got %v", expectedError, err)
    }

    if result.Id != 0 || result.Quantity != 0 || result.ProductId != 0 || result.TicketId != 0 {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeleteWishlist_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWishlistRepository := mocks.NewMockIWishlistRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	WishlistService := NewWishlistService(mockWishlistRepository)

	_, err := WishlistService.DeleteWishlist("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestGetAllWishlist_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWishlistRepository := mocks.NewMockIWishlistRepository(ctrl)

	expectedWishlists := []response.Wishlist{
		{Id: 1, Quantity: 1, ProductId: 1, TicketId: 1},
		{Id: 2, Quantity: 2, ProductId: 2, TicketId: 2},
	}

	var expectedError error

	mockWishlistRepository.EXPECT().GetAllWishlist("filter").Return(expectedWishlists, expectedError).Times(1)

	WishlistService := NewWishlistService(mockWishlistRepository)

	resultWishlists, err := WishlistService.GetAllWishlist("filter")

	if err != expectedError {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultWishlists, expectedWishlists) {
		t.Errorf("Expected Wishlists %+v, but got %+v", expectedWishlists, resultWishlists)
	}
}

func TestGetAllWishlist_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWishlistRepository := mocks.NewMockIWishlistRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockWishlistRepository.EXPECT().GetAllWishlist("filter").Return(nil, expectedError).Times(1)

	WishlistService := NewWishlistService(mockWishlistRepository)

	resultWishlists, err := WishlistService.GetAllWishlist("filter")

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if resultWishlists != nil {
		t.Errorf("Expected nil Wishlists, but got %+v", resultWishlists)
	}
}

func TestGetWishlist_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWishlistRepository := mocks.NewMockIWishlistRepository(ctrl)

	expectedWishlist := response.Wishlist{
		Id:     1,
		Quantity: 1,
		ProductId: 1,
		TicketId: 1,
	}

	mockWishlistRepository.EXPECT().GetWishlist("WishlistID").Return(expectedWishlist, nil).Times(1)

	WishlistService := NewWishlistService(mockWishlistRepository)

	resultWishlist, err := WishlistService.GetWishlist("WishlistID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultWishlist, expectedWishlist) {
		t.Errorf("Expected Wishlist %+v, but got %+v", expectedWishlist, resultWishlist)
	}
}

func TestGetWishlist_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWishlistRepository := mocks.NewMockIWishlistRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockWishlistRepository.EXPECT().GetWishlist("WishlistID").Return(response.Wishlist{}, expectedError).Times(1)

	WishlistService := NewWishlistService(mockWishlistRepository)

	resultWishlist, err := WishlistService.GetWishlist("WishlistID")

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if resultWishlist.Id != 0 || resultWishlist.Quantity != 0 || resultWishlist.ProductId != 0 || resultWishlist.TicketId != 0  {
		t.Errorf("Expected empty result, but got %+v", resultWishlist)
	}
}

func TestGetWishlist_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWishlistRepository := mocks.NewMockIWishlistRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	WishlistService := NewWishlistService(mockWishlistRepository)

	_, err := WishlistService.GetWishlist("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestUpdateWishlist_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWishlistRepository := mocks.NewMockIWishlistRepository(ctrl)

	expectedUpdatedWishlist := response.Wishlist{
		Id:          1,
		Quantity: 1,
		ProductId: 1,
		TicketId: 1,
	}

	var expectedError error

	mockWishlistRepository.EXPECT().UpdateWishlist("WishlistID", gomock.Any()).Return(expectedUpdatedWishlist, expectedError).Times(1)

	WishlistService := NewWishlistService(mockWishlistRepository)

	requestWishlist := request.Wishlist{
		Quantity: 1,
		ProductId: 1,
		TicketId: 1,
	}

	resultWishlist, err := WishlistService.UpdateWishlist("WishlistID", requestWishlist)

	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if !reflect.DeepEqual(resultWishlist, expectedUpdatedWishlist) {
		t.Errorf("Expected %+v, but got %+v", expectedUpdatedWishlist, resultWishlist)
	}
}

func TestUpdateWishlist_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWishlistRepository := mocks.NewMockIWishlistRepository(ctrl)

	expectedError := errors.New("failed to update data")

	mockWishlistRepository.EXPECT().UpdateWishlist("WishlistID", gomock.Any()).Return(response.Wishlist{}, expectedError).Times(1)

	WishlistService := NewWishlistService(mockWishlistRepository)

	requestWishlist := request.Wishlist{
		Quantity: 1,
		ProductId: 1,
		TicketId: 1,
	}

	resultWishlist, err := WishlistService.UpdateWishlist("WishlistID", requestWishlist)

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if resultWishlist.Id != 0 || resultWishlist.Quantity != 0 || resultWishlist.ProductId != 0 || resultWishlist.TicketId != 0  {
		t.Errorf("Expected empty result, but got %+v", resultWishlist)
	}
}

func TestUpdateWishlist_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockWishlistRepository := mocks.NewMockIWishlistRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	WishlistService := NewWishlistService(mockWishlistRepository)

	_, err := WishlistService.UpdateWishlist("", request.Wishlist{})

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}