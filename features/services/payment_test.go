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

func TestCreatePayment_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentRepository := mocks.NewMockIPaymentRepository(ctrl)

	expectedPayment := response.Payment{
		Id:     1,
		Name: "test",
	}

	mockPaymentRepository.EXPECT().CreatePayment(gomock.Any()).Return(expectedPayment, nil).Times(1)

	paymentService := NewPaymentService(mockPaymentRepository)

	requestPayment := &request.Payment{
		Name: "test",
	}

	resultPayment, err := paymentService.CreatePayment(requestPayment)

	assert.NoError(t, err)
	assert.Equal(t, expectedPayment, resultPayment)
}

func TestCreatePayment_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentRepository := mocks.NewMockIPaymentRepository(ctrl)

	expectedError := errors.New("failed to create payment")

	mockPaymentRepository.EXPECT().CreatePayment(gomock.Any()).Return(response.Payment{}, expectedError).Times(1)

	paymentService := NewPaymentService(mockPaymentRepository)

	requestPayment := &request.Payment{
		Name: "test",
	}

	resultPayment, err := paymentService.CreatePayment(requestPayment)

	assert.Error(t, err)
	assert.Equal(t, response.Payment{}, resultPayment)
}

func TestCreatePayment_Failure_NameEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

	mockPaymentRepository := mocks.NewMockIPaymentRepository(ctrl)

    mockRequestPayment := &request.Payment{
		Name: "",
    }

    expectedError := errors.New("name is empty")

    paymentService := NewPaymentService(mockPaymentRepository)

    result, err := paymentService.CreatePayment(mockRequestPayment)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.Name != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeletePayment_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentRepository := mocks.NewMockIPaymentRepository(ctrl)

	expectedPayment := response.Payment{
		Id:     1,
		Name: "test",
	}

	mockPaymentRepository.EXPECT().DeletePayment("paymentID").Return(expectedPayment, nil).Times(1)

	paymentService := NewPaymentService(mockPaymentRepository)

	resultPayment, err := paymentService.DeletePayment("paymentID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultPayment, expectedPayment) {
		t.Errorf("Expected payment %+v, but got %+v", expectedPayment, resultPayment)
	}
}

func TestDeletePayment_Failure_ErrDelete(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockPaymentRepository := mocks.NewMockIPaymentRepository(ctrl)

    expectedError := errors.New("failed to delete data from database")

    mockPaymentRepository.EXPECT().DeletePayment(gomock.Any()).Return(response.Payment{}, errors.New("failed to delete data from database"))

    PaymentService := NewPaymentService(mockPaymentRepository)

    result, err := PaymentService.DeletePayment("PaymentID")

    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error %v, but got %v", expectedError, err)
    }

    if result.Id != 0 || result.Name != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeletePayment_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentRepository := mocks.NewMockIPaymentRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	PaymentService := NewPaymentService(mockPaymentRepository)

	_, err := PaymentService.DeletePayment("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestGetAllPayment_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentRepository := mocks.NewMockIPaymentRepository(ctrl)

	expectedPayments := []response.Payment{
		{Id: 1, Name: "test 1"},
		{Id: 2, Name: "test 2"},
	}

	var expectedError error

	mockPaymentRepository.EXPECT().GetAllPayment().Return(expectedPayments, expectedError).Times(1)

	paymentService := NewPaymentService(mockPaymentRepository)

	resultPayments, err := paymentService.GetAllPayment()

	if err != expectedError {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultPayments, expectedPayments) {
		t.Errorf("Expected payments %+v, but got %+v", expectedPayments, resultPayments)
	}
}

func TestGetAllPayment_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentRepository := mocks.NewMockIPaymentRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockPaymentRepository.EXPECT().GetAllPayment().Return(nil, expectedError).Times(1)

	paymentService := NewPaymentService(mockPaymentRepository)

	resultPayments, err := paymentService.GetAllPayment()

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if resultPayments != nil {
		t.Errorf("Expected nil payments, but got %+v", resultPayments)
	}
}

func TestGetPayment_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentRepository := mocks.NewMockIPaymentRepository(ctrl)

	expectedPayment := response.Payment{
		Id:     1,
		Name: "test",
	}

	mockPaymentRepository.EXPECT().GetPayment("paymentID").Return(expectedPayment, nil).Times(1)

	paymentService := NewPaymentService(mockPaymentRepository)

	resultPayment, err := paymentService.GetPayment("paymentID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultPayment, expectedPayment) {
		t.Errorf("Expected payment %+v, but got %+v", expectedPayment, resultPayment)
	}
}

func TestGetPayment_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentRepository := mocks.NewMockIPaymentRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockPaymentRepository.EXPECT().GetPayment("paymentID").Return(response.Payment{}, expectedError).Times(1)

	paymentService := NewPaymentService(mockPaymentRepository)

	resultPayment, err := paymentService.GetPayment("paymentID")

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if resultPayment.Id != 0 || resultPayment.Name != "" {
		t.Errorf("Expected empty result, but got %+v", resultPayment)
	}
}

func TestGetPayment_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentRepository := mocks.NewMockIPaymentRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	PaymentService := NewPaymentService(mockPaymentRepository)

	_, err := PaymentService.GetPayment("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestUpdatePayment_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentRepository := mocks.NewMockIPaymentRepository(ctrl)

	expectedUpdatedPayment := response.Payment{
		Id:          1,
		Name:      "Update",
	}

	var expectedError error

	mockPaymentRepository.EXPECT().UpdatePayment("paymentID", gomock.Any()).Return(expectedUpdatedPayment, expectedError).Times(1)

	paymentService := NewPaymentService(mockPaymentRepository)

	requestPayment := request.Payment{
		Name:      "Update",
	}

	resultPayment, err := paymentService.UpdatePayment("paymentID", requestPayment)

	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if !reflect.DeepEqual(resultPayment, expectedUpdatedPayment) {
		t.Errorf("Expected %+v, but got %+v", expectedUpdatedPayment, resultPayment)
	}
}

func TestUpdatePayment_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentRepository := mocks.NewMockIPaymentRepository(ctrl)

	expectedError := errors.New("failed to update data")

	mockPaymentRepository.EXPECT().UpdatePayment("paymentID", gomock.Any()).Return(response.Payment{}, expectedError).Times(1)

	paymentService := NewPaymentService(mockPaymentRepository)

	requestPayment := request.Payment{
		Name:      "Update",
	}

	resultPayment, err := paymentService.UpdatePayment("paymentID", requestPayment)

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if resultPayment.Id != 0 || resultPayment.Name != ""{
		t.Errorf("Expected empty result, but got %+v", resultPayment)
	}
}

func TestUpdatePayment_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockPaymentRepository := mocks.NewMockIPaymentRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	PaymentService := NewPaymentService(mockPaymentRepository)

	_, err := PaymentService.UpdatePayment("", request.Payment{})

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}