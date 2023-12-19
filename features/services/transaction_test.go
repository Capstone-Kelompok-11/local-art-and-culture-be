package services

import (
	"errors"
	"fmt"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/mocks"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func slicesEqual(a, b []response.TransactionDetail) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}

func TestCreateTransaction_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransactionRepo := mocks.NewMockITransactionRepository(ctrl)
    mockTransactionDetailRepo := mocks.NewMockITransactionDetailRepository(ctrl)
    mockMidtransService := mocks.NewMockIMidtransService(ctrl)

	testRequest := &request.Transaction{
		TransactionDate: time.Now(),
		UserId: 1,
		PaymentMethodId: 1,
		Status: "test",
		TransactionNumber: "12345",
		User: request.User{},
		Shipping: request.Shipping{},
		Payment: request.Payment{},
		TransactionDetail: []request.TransactionDetail{},
	}

	expectedResponse := response.Transaction{
		Id: 1,
		TransactionDate: time.Now(),
		UserId: 1,
		PaymentMethodId: 1,
		Status: "test",
		TransactionNumber: "12345",
		SnapUrl: "test",
		Total: 1000,
		User: response.User{},
		Shipping: response.Shipping{},
		Payment: response.Payment{},
		TransactionDetail: []response.TransactionDetail{},
	}
	mockTransactionRepo.EXPECT().CreateTransaction(testRequest).Return(expectedResponse, nil)

	service := NewTransactionService(mockTransactionRepo, mockTransactionDetailRepo, mockMidtransService)

	actualResponse, err := service.CreateTransaction(testRequest)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !slicesEqual(actualResponse.TransactionDetail, expectedResponse.TransactionDetail) {
        t.Errorf("Expected TransactionDetail %+v, but got %+v", expectedResponse.TransactionDetail, actualResponse.TransactionDetail)
    }
}

func TestCreateTransaction_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransactionRepo := mocks.NewMockITransactionRepository(ctrl)
	mockTransactionDetailRepo := mocks.NewMockITransactionDetailRepository(ctrl)
	mockMidtransService := mocks.NewMockIMidtransService(ctrl)

	testRequest := &request.Transaction{
		TransactionDate:   time.Now(),
		UserId:            1,
		PaymentMethodId:   1,
		Status:            "test",
		TransactionNumber: "12345",
		User:              request.User{},
		Shipping:          request.Shipping{},
		Payment:           request.Payment{},
		TransactionDetail: []request.TransactionDetail{},
	}

	expectedError := errors.New("transaction detail is empty")
	mockTransactionRepo.EXPECT().CreateTransaction(testRequest).Return(response.Transaction{}, expectedError).Do(func(arg interface{}) {
		fmt.Printf("Expected call to CreateTransaction with argument: %+v\n", arg)
	})

	service := NewTransactionService(mockTransactionRepo, mockTransactionDetailRepo, mockMidtransService)

	actualResponse, err := service.CreateTransaction(testRequest)
	fmt.Printf("Actual call to CreateTransaction with argument: %+v\n", testRequest)

	if err == nil {
		t.Error("Expected an error, but got nil")
	}

	if len(actualResponse.TransactionDetail) > 0 {
		t.Errorf("Expected empty TransactionDetail, but got non-empty: %+v", actualResponse.TransactionDetail)
	}
}

// func TestCreateTransaction_Failure_TDEmpty(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockTransactionRepo := mocks.NewMockITransactionRepository(ctrl)
// 	mockTransactionDetailRepo := mocks.NewMockITransactionDetailRepository(ctrl)
// 	mockMidtransService := mocks.NewMockIMidtransService(ctrl)

// 	testRequest := &request.Transaction{
// 		TransactionDate:   time.Now(),
// 		UserId:            1,
// 		PaymentMethodId:   1,
// 		Status:            "test",
// 		TransactionNumber: "12345",
// 		User:              request.User{},
// 		Shipping:          request.Shipping{},
// 		Payment:           request.Payment{},
// 		TransactionDetail: []request.TransactionDetail{},
// 	}

// 	expectedError := errors.New("can't find any item in transaction detail")

// 	mockTransactionRepo.EXPECT().
// 		CreateTransaction(gomock.Eq(testRequest)).
// 		Return(response.Transaction{}, expectedError).
// 		Times(1)

// 	service := NewTransactionService(mockTransactionRepo, mockTransactionDetailRepo, mockMidtransService)

// 	actualResponse, err := service.CreateTransaction(testRequest)
// 	fmt.Printf("Actual call to CreateTransaction with argument: %+v\n", testRequest)

// 	if err == nil {
// 		t.Error("Expected an error, but got nil")
// 	}

// 	if len(actualResponse.TransactionDetail) > 0 {
// 		t.Errorf("Expected empty TransactionDetail, but got non-empty: %+v", actualResponse.TransactionDetail)
// 	}
// }

