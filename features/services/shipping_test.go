package services

import (
	"errors"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCreateShippingMethod_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIShippingService(ctrl)

	expectedShipping := response.Shipping{Id: 1, Name: "Standard"}
	mockService.EXPECT().CreateShippingMethod(gomock.Any()).Return(expectedShipping, nil)

	createdShipping, err := mockService.CreateShippingMethod(&request.Shipping{Name: "Standard"})

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if createdShipping.Id != expectedShipping.Id {
		t.Errorf("Expected shipping ID %d, got %d", expectedShipping.Id, createdShipping.Id)
	}

	if createdShipping.Name != expectedShipping.Name {
		t.Errorf("Expected shipping method name %s, got %s", expectedShipping.Name, createdShipping.Name)
	}
}

func TestCreateShippingMethod_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIShippingService(ctrl)

	expectedError := errors.New("some error")
	mockService.EXPECT().CreateShippingMethod(gomock.Any()).Return(response.Shipping{}, expectedError)

	createdShipping, err := mockService.CreateShippingMethod(&request.Shipping{Name: "Standard"})

	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}

	if createdShipping.Id != 0 {
		t.Errorf("Expected empty shipping ID, got %d", createdShipping.Id)
	}

	if createdShipping.Name != "" {
		t.Errorf("Expected empty shipping method name, got %s", createdShipping.Name)
	}
}

func TestDeleteShipping_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIShippingService(ctrl)

	expectedDeletedShipping := response.Shipping{Id: 1, Name: "Shipping Service"}
	mockService.EXPECT().DeleteShipping(gomock.Any()).Return(expectedDeletedShipping, nil)

	deletedShipping, err := mockService.DeleteShipping("shipping1")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if deletedShipping.Id != expectedDeletedShipping.Id {
		t.Errorf("Expected shipping ID %d, got %d", expectedDeletedShipping.Id, deletedShipping.Id)
	}

	if deletedShipping.Name != expectedDeletedShipping.Name {
		t.Errorf("Expected shipping name %s, got %s", expectedDeletedShipping.Name, deletedShipping.Name)
	}
}

func TestDeleteShipping_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIShippingService(ctrl)

	expectedError := errors.New("some error")
	mockService.EXPECT().DeleteShipping(gomock.Any()).Return(response.Shipping{}, expectedError)

	deletedShipping, err := mockService.DeleteShipping("shipping1")

	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}

	if deletedShipping.Id != 0 {
		t.Errorf("Expected empty shipping ID, got %d", deletedShipping.Id)
	}

	if deletedShipping.Name != "" {
		t.Errorf("Expected empty shipping name, got %s", deletedShipping.Name)
	}
}

func TestGetAllShipping_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIShippingService(ctrl)

	expectedShippingList := []response.Shipping{
		{Id: 1, Name: "Shipping Service 1"},
		{Id: 2, Name: "Shipping Service 2"},
	}

	mockService.EXPECT().GetAllShipping().Return(expectedShippingList, nil)

	shippingList, err := mockService.GetAllShipping()

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(shippingList) != len(expectedShippingList) {
		t.Errorf("Expected shipping list length %d, got %d", len(expectedShippingList), len(shippingList))
	}

	for i, expectedShipping := range expectedShippingList {
		if shippingList[i].Id != expectedShipping.Id {
			t.Errorf("Expected shipping ID %d, got %d", expectedShipping.Id, shippingList[i].Id)
		}

		if shippingList[i].Name != expectedShipping.Name {
			t.Errorf("Expected shipping name %s, got %s", expectedShipping.Name, shippingList[i].Name)
		}
	}
}

func TestGetAllShipping_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIShippingService(ctrl)

	expectedError := errors.New("some error")
	mockService.EXPECT().GetAllShipping().Return([]response.Shipping{}, expectedError)

	shippingList, err := mockService.GetAllShipping()

	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}

	if len(shippingList) != 0 {
		t.Errorf("Expected empty shipping list, got %v", shippingList)
	}
}

func TestGetShipping_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIShippingService(ctrl)

	expectedShipping := response.Shipping{Id: 1, Name: "Shipping Service"}
	mockService.EXPECT().GetShipping(gomock.Any()).Return(expectedShipping, nil)

	receivedShipping, err := mockService.GetShipping("shipping1")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if receivedShipping.Id != expectedShipping.Id {
		t.Errorf("Expected shipping ID %d, got %d", expectedShipping.Id, receivedShipping.Id)
	}

	if receivedShipping.Name != expectedShipping.Name {
		t.Errorf("Expected shipping name %s, got %s", expectedShipping.Name, receivedShipping.Name)
	}
}

func TestGetShipping_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIShippingService(ctrl)

	expectedError := errors.New("some error")
	mockService.EXPECT().GetShipping(gomock.Any()).Return(response.Shipping{}, expectedError)

	receivedShipping, err := mockService.GetShipping("shipping1")

	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}

	if receivedShipping.Id != 0 {
		t.Errorf("Expected empty shipping ID, got %d", receivedShipping.Id)
	}

	if receivedShipping.Name != "" {
		t.Errorf("Expected empty shipping name, got %s", receivedShipping.Name)
	}
}

func TestUpdateShipping_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIShippingService(ctrl)

	expectedUpdatedShipping := response.Shipping{Id: 1, Name: "Updated Shipping"}
	mockService.EXPECT().UpdateShipping(gomock.Any(), gomock.Any()).Return(expectedUpdatedShipping, nil)

	updatedShipping, err := mockService.UpdateShipping("shippingID", request.Shipping{Name: "New Method"})

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if updatedShipping.Id != expectedUpdatedShipping.Id {
		t.Errorf("Expected shipping ID %d, got %d", expectedUpdatedShipping.Id, updatedShipping.Id)
	}

	if updatedShipping.Name != expectedUpdatedShipping.Name {
		t.Errorf("Expected shipping method %s, got %s", expectedUpdatedShipping.Name, updatedShipping.Name)
	}
}

func TestUpdateShipping_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIShippingService(ctrl)

	expectedError := errors.New("some error")
	mockService.EXPECT().UpdateShipping(gomock.Any(), gomock.Any()).Return(response.Shipping{}, expectedError)

	updatedShipping, err := mockService.UpdateShipping("shippingID", request.Shipping{Name: "New Method"})

	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}

	if updatedShipping.Id != 0 {
		t.Errorf("Expected empty shipping ID, got %d", updatedShipping.Id)
	}

	if updatedShipping.Name != "" {
		t.Errorf("Expected empty shipping method, got %s", updatedShipping.Name)
	}
}