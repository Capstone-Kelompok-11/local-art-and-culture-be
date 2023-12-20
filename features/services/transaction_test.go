package services

import (
	"errors"
	"fmt"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/mocks"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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

func TestGetAllTransaction_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransactionRepo := mocks.NewMockITransactionRepository(ctrl)
	mockTransactionDetailRepo := mocks.NewMockITransactionDetailRepository(ctrl)
	mockMidtransService := mocks.NewMockIMidtransService(ctrl)

	expectedTransactions := []response.Transaction{
		{
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
		},
		{
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
		},
	}

	mockTransactionRepo.EXPECT().GetAllTransaction().Return(nil, expectedTransactions)

	service := NewTransactionService(mockTransactionRepo, mockTransactionDetailRepo, mockMidtransService)

	transactions, err := service.GetAllTransaction()

	assert.NoError(t, err)
	assert.Equal(t, expectedTransactions, transactions)
}

func TestGetAllTransaction_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransactionRepo := mocks.NewMockITransactionRepository(ctrl)
	mockTransactionDetailRepo := mocks.NewMockITransactionDetailRepository(ctrl)
	mockMidtransService := mocks.NewMockIMidtransService(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockTransactionRepo.EXPECT().GetAllTransaction().Return(expectedError, nil)

	service := NewTransactionService(mockTransactionRepo, mockTransactionDetailRepo, mockMidtransService)

	transactions, err := service.GetAllTransaction()

	assert.Error(t, err)
	assert.Nil(t, transactions)
	assert.Equal(t, expectedError, err)
}

func TestGetReportTransaction_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockTransactionRepo := mocks.NewMockITransactionRepository(ctrl)
    mockTransactionDetailRepo := mocks.NewMockITransactionDetailRepository(ctrl)
    mockMidtransService := mocks.NewMockIMidtransService(ctrl)

    expectedReport := response.TransactionReport{
        Id:              1,
        TransactionDate: time.Now(),
        Status:          "test",
        Qty:             100,
        Price:           1000,
        Nominal:         10000,
    }

    mockTransactionRepo.EXPECT().GetReportTransaction(gomock.Any(), gomock.Any()).Return([]response.TransactionReport{expectedReport}, nil)

    service := NewTransactionService(mockTransactionRepo, mockTransactionDetailRepo, mockMidtransService)

    reports, err := service.GetReportTransaction(123, "some_criteria")

    assert.NoError(t, err)
    assert.NotNil(t, reports)
    assert.Len(t, reports, 1) 
    assert.Equal(t, expectedReport, reports[0])
}

func TestGetReportTransaction_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

	mockTransactionRepo := mocks.NewMockITransactionRepository(ctrl)
    mockTransactionDetailRepo := mocks.NewMockITransactionDetailRepository(ctrl)
    mockMidtransService := mocks.NewMockIMidtransService(ctrl)

    expectedError := errors.New("database can't request any data right now")
    mockTransactionRepo.EXPECT().GetReportTransaction(gomock.Any(), gomock.Any()).Return(nil, expectedError)

	service := NewTransactionService(mockTransactionRepo, mockTransactionDetailRepo, mockMidtransService)

    reports, err := service.GetReportTransaction(123, "some_criteria")

    assert.Error(t, err)
    assert.Nil(t, reports)
    assert.EqualError(t, err, expectedError.Error())
}

func TestGetTransaction_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockTransactionRepo := mocks.NewMockITransactionRepository(ctrl)
    mockTransactionDetailRepo := mocks.NewMockITransactionDetailRepository(ctrl)
    mockMidtransService := mocks.NewMockIMidtransService(ctrl)
	service := NewTransactionService(mockTransactionRepo, mockTransactionDetailRepo, mockMidtransService)

    expectedTransaction := response.Transaction{ 
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
    mockTransactionRepo.EXPECT().GetTransaction(gomock.Any()).Return(nil, expectedTransaction)

    resultTransaction, err := service.GetTransaction("transactionID")

    if err != nil {
        t.Errorf("Expected no error, but got: %v", err)
    }

    if !reflect.DeepEqual(resultTransaction, expectedTransaction) {
        t.Errorf("Expected transaction %v, but got %v", expectedTransaction, resultTransaction)
    }
}

func TestGetTransaction_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockTransactionRepo := mocks.NewMockITransactionRepository(ctrl)
    mockTransactionDetailRepo := mocks.NewMockITransactionDetailRepository(ctrl)
    mockMidtransService := mocks.NewMockIMidtransService(ctrl)
    service := NewTransactionService(mockTransactionRepo, mockTransactionDetailRepo, mockMidtransService)

    expectedError := errors.New("database can't request any data right now")
    mockTransactionRepo.EXPECT().GetTransaction(gomock.Any()).Return(expectedError, response.Transaction{})

    resultTransaction, err := service.GetTransaction("transactionID")

    if err == nil {
        t.Error("Expected an error, but got nil")
    }

    if resultTransaction.Id != 0 || resultTransaction.UserId != 0 {
        t.Errorf("Expected no transaction, but got %+v", resultTransaction)
    }
}

func TestGetTransaction_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransactionRepo := mocks.NewMockITransactionRepository(ctrl)
    mockTransactionDetailRepo := mocks.NewMockITransactionDetailRepository(ctrl)
    mockMidtransService := mocks.NewMockIMidtransService(ctrl)
    
	_ = errors.New("can't find any data with this id")

	service := NewTransactionService(mockTransactionRepo, mockTransactionDetailRepo, mockMidtransService)

	_, err := service.GetTransaction("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestDeleteTransaction_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransactionRepo := mocks.NewMockITransactionRepository(ctrl)
    mockTransactionDetailRepo := mocks.NewMockITransactionDetailRepository(ctrl)
    mockMidtransService := mocks.NewMockIMidtransService(ctrl)

	expectedTransaction := response.Transaction{
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

	mockTransactionRepo.EXPECT().DeleteTransaction("TransactionID").Return(nil, expectedTransaction).Times(1)

	service := NewTransactionService(mockTransactionRepo, mockTransactionDetailRepo, mockMidtransService)

	resultTransaction, err := service.DeleteTransaction("TransactionID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultTransaction, expectedTransaction) {
		t.Errorf("Expected Transaction %+v, but got %+v", expectedTransaction, resultTransaction)
	}
}

func TestDeleteTransaction_Failure_ErrDelete(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

	mockTransactionRepo := mocks.NewMockITransactionRepository(ctrl)
    mockTransactionDetailRepo := mocks.NewMockITransactionDetailRepository(ctrl)
    mockMidtransService := mocks.NewMockIMidtransService(ctrl)

    expectedError := errors.New("can't delete this transaction")

    mockTransactionRepo.EXPECT().DeleteTransaction(gomock.Any()).Return(errors.New("can't delete this transaction"), response.Transaction{})

	service := NewTransactionService(mockTransactionRepo, mockTransactionDetailRepo, mockMidtransService)

    result, err := service.DeleteTransaction("TransactionID")

    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error %v, but got %v", expectedError, err)
    }

    if result.Id != 0 || result.UserId != 0 || result.PaymentMethodId != 0 || result.Total != 0 {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeleteTransaction_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransactionRepo := mocks.NewMockITransactionRepository(ctrl)
    mockTransactionDetailRepo := mocks.NewMockITransactionDetailRepository(ctrl)
    mockMidtransService := mocks.NewMockIMidtransService(ctrl)
    
	_ = errors.New("can't find any data with this id")

	service := NewTransactionService(mockTransactionRepo, mockTransactionDetailRepo, mockMidtransService)

	_, err := service.DeleteTransaction("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestGetHistoryTransaction_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransactionRepo := mocks.NewMockITransactionRepository(ctrl)
    mockTransactionDetailRepo := mocks.NewMockITransactionDetailRepository(ctrl)
    mockMidtransService := mocks.NewMockIMidtransService(ctrl)

	userID := uint(1)
	page := 1
	limit := 10
	expectedTransactions := []*response.Transaction{
		{
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
	},
}
	expectedTotal := 100
	mockTransactionRepo.EXPECT().GetHistoryTransaction(userID, page, limit).Return(expectedTransactions, expectedTotal, nil)

	service := NewTransactionService(mockTransactionRepo, mockTransactionDetailRepo, mockMidtransService)
	transactions, total, err := service.GetHistoryTransaction(userID, page, limit)

	assert.NoError(t, err)
	assert.Equal(t, expectedTransactions, transactions)
	assert.Equal(t, expectedTotal, total)
}

func TestGetHistoryTransaction_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockTransactionRepo := mocks.NewMockITransactionRepository(ctrl)
    mockTransactionDetailRepo := mocks.NewMockITransactionDetailRepository(ctrl)
    mockMidtransService := mocks.NewMockIMidtransService(ctrl)

    userID := uint(1)
    page := 1
    limit := 10

    expectedError := errors.New("database can't request any data right now")

    mockTransactionRepo.EXPECT().GetHistoryTransaction(userID, page, limit).Return(nil, 0, expectedError)

    service := NewTransactionService(mockTransactionRepo, mockTransactionDetailRepo, mockMidtransService)

    _, _, err := service.GetHistoryTransaction(userID, page, limit)

    assert.Error(t, err)
    assert.EqualError(t, err, expectedError.Error())
}

func TestGetHistoryTransaction_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransactionRepo := mocks.NewMockITransactionRepository(ctrl)
	mockTransactionDetailRepo := mocks.NewMockITransactionDetailRepository(ctrl)
	mockMidtransService := mocks.NewMockIMidtransService(ctrl)

	_ = errors.New("can't find any data with this id")

	service := NewTransactionService(mockTransactionRepo, mockTransactionDetailRepo, mockMidtransService)

	_, _, err := service.GetHistoryTransaction(0, 0, 0)

	if err == nil {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

// func TestConfirmPayment_Success(t *testing.T) {
//     ctrl := gomock.NewController(t)
//     defer ctrl.Finish()

//     mockTransactionService := mocks.NewMockITransactionService(ctrl)

//     // Set up the expectation for ConfirmPayment
//     mockTransactionService.EXPECT().ConfirmPayment("arg0_value", "arg1_value").
//         Return(response.Transaction{}, nil)

//     // Call ConfirmPayment and check the result
//     result, err := mockTransactionService.ConfirmPayment("arg0_value", "arg1_value")

//     // Assert that there is no error and the result matches the expected response
//     assert.NoError(t, err)
//     assert.Equal(t, response.Transaction{}, result)
// }

func TestCalculatePaginationValues_Transaction(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockTransactionRepo := mocks.NewMockITransactionRepository(ctrl)
	mockTransactionDetailRepo := mocks.NewMockITransactionDetailRepository(ctrl)
	mockMidtransService := mocks.NewMockIMidtransService(ctrl)

    TransactionService := NewTransactionService(mockTransactionRepo, mockTransactionDetailRepo, mockMidtransService)

    page1, totalPages1 := TransactionService.CalculatePaginationValues(0, 100, 8)
    assert.Equal(t, 1, page1)
    assert.Equal(t, 1, totalPages1)

    page2, totalPages2 := TransactionService.CalculatePaginationValues(15, 100, 8)
    assert.Equal(t, 1, page2)
    assert.Equal(t, 1, totalPages2)

    page3, totalPages3 := TransactionService.CalculatePaginationValues(2, 100, 8)
    assert.Equal(t, 1, page3)
    assert.Equal(t, 1, totalPages3)

    page4, totalPages4 := TransactionService.CalculatePaginationValues(1, 95, 8)
    assert.Equal(t, 1, page4)
    assert.Equal(t, 1, totalPages4) 
}

func TestGetNextPage_Transaction(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

	mockTransactionRepo := mocks.NewMockITransactionRepository(ctrl)
	mockTransactionDetailRepo := mocks.NewMockITransactionDetailRepository(ctrl)
	mockMidtransService := mocks.NewMockIMidtransService(ctrl)

    TransactionService := NewTransactionService(mockTransactionRepo, mockTransactionDetailRepo, mockMidtransService)

    result1 := TransactionService.GetNextPage(3, 10)
	assert.Equal(t, 4, result1)

	result2 := TransactionService.GetNextPage(10, 10)
	assert.Equal(t, 10, result2)
}

func TestGetPrevPage_Transaction(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

	mockTransactionRepo := mocks.NewMockITransactionRepository(ctrl)
	mockTransactionDetailRepo := mocks.NewMockITransactionDetailRepository(ctrl)
	mockMidtransService := mocks.NewMockIMidtransService(ctrl)

    TransactionService := NewTransactionService(mockTransactionRepo, mockTransactionDetailRepo, mockMidtransService)

    result1 := TransactionService.GetPrevPage(5)
	assert.Equal(t, 4, result1)

	result2 := TransactionService.GetPrevPage(1)
	assert.Equal(t, 1, result2)
}