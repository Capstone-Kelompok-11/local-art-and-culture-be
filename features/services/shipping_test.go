package services

import (
	"errors"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/mocks"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateShipping_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShippingRepository := mocks.NewMockIShippingRepository(ctrl)

	expectedShipping := response.Shipping{
		Id:     1,
		Name: "test",
	}

	mockShippingRepository.EXPECT().CreateShippingMethod(gomock.Any()).Return(expectedShipping, nil).Times(1)

	ShippingService := NewShippingService(mockShippingRepository)

	requestShipping := &request.Shipping{
		Name: "test",
	}

	resultShipping, err := ShippingService.CreateShippingMethod(requestShipping)

	assert.NoError(t, err)
	assert.Equal(t, expectedShipping, resultShipping)
}

func TestCreateShipping_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShippingRepository := mocks.NewMockIShippingRepository(ctrl)

	expectedError := errors.New("failed to create Shipping")

	mockShippingRepository.EXPECT().CreateShippingMethod(gomock.Any()).Return(response.Shipping{}, expectedError).Times(1)

	ShippingService := NewShippingService(mockShippingRepository)

	requestShipping := &request.Shipping{
		Name: "test",
	}

	resultShipping, err := ShippingService.CreateShippingMethod(requestShipping)

	assert.Error(t, err)
	assert.Equal(t, response.Shipping{}, resultShipping)
}

func TestCreateShipping_Failure_NameEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

	mockShippingRepository := mocks.NewMockIShippingRepository(ctrl)

    mockRequestShipping := &request.Shipping{
		Name: "",
    }

    expectedError := errors.New("role is empty")

    ShippingService := NewShippingService(mockShippingRepository)

    result, err := ShippingService.CreateShippingMethod(mockRequestShipping)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.Name != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeleteShipping_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShippingRepository := mocks.NewMockIShippingRepository(ctrl)

	expectedShipping := response.Shipping{
		Id:     1,
		Name: "test",
	}

	mockShippingRepository.EXPECT().DeleteShipping("ShippingID").Return(expectedShipping, nil).Times(1)

	ShippingService := NewShippingService(mockShippingRepository)

	resultShipping, err := ShippingService.DeleteShipping("ShippingID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultShipping, expectedShipping) {
		t.Errorf("Expected Shipping %+v, but got %+v", expectedShipping, resultShipping)
	}
}

func TestDeleteShipping_Failure_ErrDelete(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockShippingRepository := mocks.NewMockIShippingRepository(ctrl)

    expectedError := errors.New("can't delete this shipping method")

    mockShippingRepository.EXPECT().DeleteShipping(gomock.Any()).Return(response.Shipping{}, errors.New("failed to delete data from database"))

    ShippingService := NewShippingService(mockShippingRepository)

    result, err := ShippingService.DeleteShipping("ShippingID")

    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error %v, but got %v", expectedError, err)
    }

    if result.Id != 0 || result.Name != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeleteShipping_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShippingRepository := mocks.NewMockIShippingRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	ShippingService := NewShippingService(mockShippingRepository)

	_, err := ShippingService.DeleteShipping("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestGetAllShipping_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShippingRepository := mocks.NewMockIShippingRepository(ctrl)

	expectedShippings := []response.Shipping{
		{Id: 1, Name: "test 1"},
		{Id: 2, Name: "test 2"},
	}

	var expectedError error

	mockShippingRepository.EXPECT().GetAllShipping().Return(expectedShippings, expectedError).Times(1)

	ShippingService := NewShippingService(mockShippingRepository)

	resultShippings, err := ShippingService.GetAllShipping()

	if err != expectedError {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultShippings, expectedShippings) {
		t.Errorf("Expected Shippings %+v, but got %+v", expectedShippings, resultShippings)
	}
}

func TestGetAllShipping_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShippingRepository := mocks.NewMockIShippingRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockShippingRepository.EXPECT().GetAllShipping().Return(nil, expectedError).Times(1)

	ShippingService := NewShippingService(mockShippingRepository)

	resultShippings, err := ShippingService.GetAllShipping()

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if resultShippings != nil {
		t.Errorf("Expected nil Shippings, but got %+v", resultShippings)
	}
}

func TestGetShipping_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShippingRepository := mocks.NewMockIShippingRepository(ctrl)

	expectedShipping := response.Shipping{
		Id:     1,
		Name: "test",
	}

	mockShippingRepository.EXPECT().GetShipping("ShippingID").Return(expectedShipping, nil).Times(1)

	ShippingService := NewShippingService(mockShippingRepository)

	resultShipping, err := ShippingService.GetShipping("ShippingID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultShipping, expectedShipping) {
		t.Errorf("Expected Shipping %+v, but got %+v", expectedShipping, resultShipping)
	}
}

func TestGetShipping_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShippingRepository := mocks.NewMockIShippingRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockShippingRepository.EXPECT().GetShipping("ShippingID").Return(response.Shipping{}, expectedError).Times(1)

	ShippingService := NewShippingService(mockShippingRepository)

	resultShipping, err := ShippingService.GetShipping("ShippingID")

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if resultShipping.Id != 0 || resultShipping.Name != "" {
		t.Errorf("Expected empty result, but got %+v", resultShipping)
	}
}

func TestGetShipping_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShippingRepository := mocks.NewMockIShippingRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	ShippingService := NewShippingService(mockShippingRepository)

	_, err := ShippingService.GetShipping("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestUpdateShipping_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShippingRepository := mocks.NewMockIShippingRepository(ctrl)

	expectedUpdatedShipping := response.Shipping{
		Id:          1,
		Name:      "Update",
	}

	var expectedError error

	mockShippingRepository.EXPECT().UpdateShipping("ShippingID", gomock.Any()).Return(expectedUpdatedShipping, expectedError).Times(1)

	ShippingService := NewShippingService(mockShippingRepository)

	requestShipping := request.Shipping{
		Name:      "Update",
	}

	resultShipping, err := ShippingService.UpdateShipping("ShippingID", requestShipping)

	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if !reflect.DeepEqual(resultShipping, expectedUpdatedShipping) {
		t.Errorf("Expected %+v, but got %+v", expectedUpdatedShipping, resultShipping)
	}
}

func TestUpdateShipping_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShippingRepository := mocks.NewMockIShippingRepository(ctrl)

	expectedError := errors.New("failed to update data")

	mockShippingRepository.EXPECT().UpdateShipping("ShippingID", gomock.Any()).Return(response.Shipping{}, expectedError).Times(1)

	ShippingService := NewShippingService(mockShippingRepository)

	requestShipping := request.Shipping{
		Name:      "Update",
	}

	resultShipping, err := ShippingService.UpdateShipping("ShippingID", requestShipping)

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if resultShipping.Id != 0 || resultShipping.Name != ""{
		t.Errorf("Expected empty result, but got %+v", resultShipping)
	}
}

func TestUpdateShipping_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockShippingRepository := mocks.NewMockIShippingRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	ShippingService := NewShippingService(mockShippingRepository)

	_, err := ShippingService.UpdateShipping("", request.Shipping{})

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}