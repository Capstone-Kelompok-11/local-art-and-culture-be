package services

import (
	"errors"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestCreatePayment_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIPaymentService(ctrl)

	expectedPayment := response.Payment{Id: 1, Name: "test"}
	mockService.EXPECT().CreatePayment(gomock.Any()).Return(expectedPayment, nil)

	createdPayment, err := mockService.CreatePayment(&request.Payment{Name: "test"})

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if createdPayment.Id != expectedPayment.Id {
		t.Errorf("Expected payment ID %d, got %d", expectedPayment.Id, createdPayment.Id)
	}

	if createdPayment.Name != expectedPayment.Name {
		t.Errorf("Expected payment Name %s, got %s", expectedPayment.Name, createdPayment.Name)
	}
}

func TestCreatePayment_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIPaymentService(ctrl)

	expectedError := errors.New("some error")
	mockService.EXPECT().CreatePayment(gomock.Any()).Return(response.Payment{}, expectedError)

	createdPayment, err := mockService.CreatePayment(&request.Payment{Name: "test"})

	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}

	if createdPayment.Id != 0 {
		t.Errorf("Expected empty payment ID, got %d", createdPayment.Id)
	}

	if createdPayment.Name != "" {
		t.Errorf("Expected empty payment Name, got %s", createdPayment.Name)
	}
}

