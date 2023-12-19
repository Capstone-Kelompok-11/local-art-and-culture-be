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

func TestCreateTransactionDetail_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransactionDetailRepository := mocks.NewMockITransactionDetailRepository(ctrl)

	expectedTransactionDetail := response.TransactionDetail{
		Id:   1,
		TransactionId: 1,
		Qty: 1,
		Subtotal: 100,
		Product: response.Product{},
		Ticket: response.Ticket{},
	}

	mockTransactionDetailRepository.EXPECT().CreateTransactionDetail(gomock.Any()).Return(expectedTransactionDetail, nil).Times(1)

	TransactionDetailService := NewTransactionDetailService(mockTransactionDetailRepository)

	requestTransactionDetail := &request.TransactionDetail{
		TransactionId: 1,
		Qty: 1,
		Product: request.Product{},
		Ticket: request.Ticket{},
	}

	resultTransactionDetail, err := TransactionDetailService.CreateTransactionDetail(requestTransactionDetail)

	assert.NoError(t, err)
	assert.Equal(t, expectedTransactionDetail, resultTransactionDetail)
}

func TestCreateTransactionDetail_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransactionDetailRepository := mocks.NewMockITransactionDetailRepository(ctrl)

	expectedError := errors.New("failed to create TransactionDetail")

	mockTransactionDetailRepository.EXPECT().CreateTransactionDetail(gomock.Any()).Return(response.TransactionDetail{}, expectedError).Times(1)

	TransactionDetailService := NewTransactionDetailService(mockTransactionDetailRepository)

	requestTransactionDetail := &request.TransactionDetail{
		TransactionId: 1,
		Qty: 1,
		Product: request.Product{},
		Ticket: request.Ticket{},
	}

	resultTransactionDetail, err := TransactionDetailService.CreateTransactionDetail(requestTransactionDetail)

	assert.Error(t, err)
	assert.Equal(t, response.TransactionDetail{}, resultTransactionDetail)
}

func TestDeleteTransactionDetail_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockITransactionDetailRepository(ctrl)

	expectedTransactionDetail := response.TransactionDetail{
		Id:             1,
        TransactionId: 1,
        Qty:            1,
        Subtotal:       100,
        Product:        response.Product{},
        Ticket:         response.Ticket{},
	}
	mockRepo.EXPECT().DeleteTransactionDetail("TransactionDetailID").Return(nil, expectedTransactionDetail)

	service := NewTransactionDetailService(mockRepo) 

	resultTransactionDetail, err := service.DeleteTransactionDetail("TransactionDetailID")

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if resultTransactionDetail != expectedTransactionDetail {
		t.Errorf("Expected transaction detail: %+v, but got: %+v", expectedTransactionDetail, resultTransactionDetail)
	}
}

func TestDeleteTransactionDetail_Failure_ErrDelete(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockTransactionDetailRepository := mocks.NewMockITransactionDetailRepository(ctrl)

    expectedError := errors.New("failed to delete data from database")

    mockTransactionDetailRepository.EXPECT().DeleteTransactionDetail(gomock.Any()).Return(errors.New("failed to delete data from database"), response.TransactionDetail{})

    TransactionDetailService := NewTransactionDetailService(mockTransactionDetailRepository)

    result, err := TransactionDetailService.DeleteTransactionDetail("TransactionDetailID")

    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error %v, but got %v", expectedError, err)
    }

    if result.Id != 0 || result.Qty != 0 || result.TransactionId != 0 || result.Subtotal != 0 {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeleteTransactionDetail_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransactionDetailRepository := mocks.NewMockITransactionDetailRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	TransactionDetailService := NewTransactionDetailService(mockTransactionDetailRepository)

	_, err := TransactionDetailService.DeleteTransactionDetail("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestGetAllTransactionDetail_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransactionDetailRepository := mocks.NewMockITransactionDetailRepository(ctrl)

	expectedTransactionDetails := []response.TransactionDetail{
		{Id: 1, TransactionId: 1, Qty: 1, Subtotal: 100, Product: response.Product{}, Ticket: response.Ticket{}},
		{Id: 2, TransactionId: 2, Qty: 2, Subtotal: 200, Product: response.Product{}, Ticket: response.Ticket{}},
	}

	var expectedError error

	mockTransactionDetailRepository.EXPECT().GetAllTransactionDetail().Return(expectedError, expectedTransactionDetails).Times(1)

	TransactionDetailService := NewTransactionDetailService(mockTransactionDetailRepository)

	resultTransactionDetails, err := TransactionDetailService.GetAllTransactionDetail()

	if err != expectedError {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultTransactionDetails, expectedTransactionDetails) {
		t.Errorf("Expected TransactionDetails %+v, but got %+v", expectedTransactionDetails, resultTransactionDetails)
	}
}

func TestGetAllTransactionDetail_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransactionDetailRepository := mocks.NewMockITransactionDetailRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockTransactionDetailRepository.EXPECT().GetAllTransactionDetail().Return(expectedError, nil).Times(1)

	TransactionDetailService := NewTransactionDetailService(mockTransactionDetailRepository)

	resultTransactionDetails, err := TransactionDetailService.GetAllTransactionDetail()

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if resultTransactionDetails != nil {
		t.Errorf("Expected nil TransactionDetails, but got %+v", resultTransactionDetails)
	}
}

func TestGetTransactionDetail_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransactionDetailRepository := mocks.NewMockITransactionDetailRepository(ctrl)

	expectedTransactionDetail := response.TransactionDetail{Id: 1, TransactionId: 1, Qty: 1, Subtotal: 100, Product: response.Product{}, Ticket: response.Ticket{}}

	mockTransactionDetailRepository.EXPECT().GetTransactionDetail("TransactionDetailID").Return(nil, expectedTransactionDetail).Times(1)

	TransactionDetailService := NewTransactionDetailService(mockTransactionDetailRepository)

	resultTransactionDetail, err := TransactionDetailService.GetTransactionDetail("TransactionDetailID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultTransactionDetail, expectedTransactionDetail) {
		t.Errorf("Expected TransactionDetail %+v, but got %+v", expectedTransactionDetail, resultTransactionDetail)
	}
}

func TestGetTransactionDetail_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransactionDetailRepository := mocks.NewMockITransactionDetailRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockTransactionDetailRepository.EXPECT().GetTransactionDetail("TransactionDetailID").Return(expectedError, response.TransactionDetail{}).Times(1)

	TransactionDetailService := NewTransactionDetailService(mockTransactionDetailRepository)

	result, err := TransactionDetailService.GetTransactionDetail("TransactionDetailID")

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if result.Id != 0 || result.Qty != 0 || result.TransactionId != 0 || result.Subtotal != 0 {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestGetTransactionDetail_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransactionDetailRepository := mocks.NewMockITransactionDetailRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	TransactionDetailService := NewTransactionDetailService(mockTransactionDetailRepository)

	_, err := TransactionDetailService.GetTransactionDetail("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestUpdateTransactionDetail_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransactionDetailRepository := mocks.NewMockITransactionDetailRepository(ctrl)

	expectedUpdatedTransactionDetail := response.TransactionDetail{Id: 1, TransactionId: 1, Qty: 1, Subtotal: 100, Product: response.Product{}, Ticket: response.Ticket{}}

	var expectedError error
	mockTransactionDetailRepository.EXPECT().UpdateTransactionDetail("TransactionDetailID", gomock.Any()).Return(expectedError, expectedUpdatedTransactionDetail).Times(1)

	TransactionDetailService := NewTransactionDetailService(mockTransactionDetailRepository)

	requestTransactionDetail := request.TransactionDetail{
		TransactionId: 1,
		Qty: 1,
		Product: request.Product{},
		Ticket: request.Ticket{},
	}

	resultTransactionDetail, err := TransactionDetailService.UpdateTransactionDetail("TransactionDetailID", requestTransactionDetail)

	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if !reflect.DeepEqual(resultTransactionDetail, expectedUpdatedTransactionDetail) {
		t.Errorf("Expected %+v, but got %+v", expectedUpdatedTransactionDetail, resultTransactionDetail)
	}
}

func TestUpdateTransactionDetail_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransactionDetailRepository := mocks.NewMockITransactionDetailRepository(ctrl)

	expectedError := errors.New("failed to update data")

	mockTransactionDetailRepository.EXPECT().UpdateTransactionDetail("TransactionDetailID", gomock.Any()).Return(expectedError, response.TransactionDetail{}).Times(1)

	TransactionDetailService := NewTransactionDetailService(mockTransactionDetailRepository)

	requestTransactionDetail := request.TransactionDetail{
		TransactionId: 1,
		Qty: 1,
		Product: request.Product{},
		Ticket: request.Ticket{},
	}

	result, err := TransactionDetailService.UpdateTransactionDetail("TransactionDetailID", requestTransactionDetail)

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if result.Id != 0 || result.Qty != 0 || result.TransactionId != 0 || result.Subtotal != 0 {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestUpdateTransactionDetail_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransactionDetailRepository := mocks.NewMockITransactionDetailRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	TransactionDetailService := NewTransactionDetailService(mockTransactionDetailRepository)

	_, err := TransactionDetailService.UpdateTransactionDetail("", request.TransactionDetail{})

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}