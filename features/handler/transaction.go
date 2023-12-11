package handler

import (
	"fmt"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"
	consts "lokasani/helpers/const"
	"lokasani/helpers/errors"
	"lokasani/helpers/middleware"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	transactionService services.ITransactionService
}

func NewTransactionHandler(ITransactionService services.ITransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: ITransactionService}
}

func (ah *TransactionHandler) CreateTransaction(c echo.Context) error {
	var input request.Transaction
	c.Bind(&input)
	userID, _, _, err := middleware.ExtractToken(c)
	input.UserId = userID
	res, err := ah.transactionService.CreateTransaction(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *TransactionHandler) GetAllTransaction(c echo.Context) error {
	res, err := ah.transactionService.GetAllTransaction()
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *TransactionHandler) GetTransaction(c echo.Context) error {
	id := c.Param("id")
	res, err := ah.transactionService.GetTransaction(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *TransactionHandler) UpdateTransaction(c echo.Context) error {
	id := c.Param("id")
	var input request.Transaction
	c.Bind(&input)
	res, err := ah.transactionService.UpdateTransaction(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *TransactionHandler) DeleteTransaction(c echo.Context) error {
	id := c.Param("id")
	res, err := ah.transactionService.DeleteTransaction(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *TransactionHandler) ConfirmPayment(c echo.Context) error {
	var paymentPayload map[string]interface{}
	err := c.Bind(&paymentPayload)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	orderId, exists := paymentPayload["order_id"].(string)
	fmt.Println(orderId)
	if !exists {
		fmt.Println(paymentPayload)
		return response.NewErrorResponse(c, errors.ERR_INVALID_PAYLOAD)
	}

	res, err := ah.transactionService.ConfirmPayment(orderId)
	if err != nil {
		fmt.Println("cek midtrans")
		return response.NewErrorResponse(c, err)
	}

	return response.NewSuccessResponse(c, res)
}

func (ah *TransactionHandler) GetHistoryTransaction(c echo.Context) error {
    userId, _, _, err := middleware.ExtractToken(c)
    if err != nil {
        return response.NewErrorResponse(c, err)
    }	

	page, pageSize := 1, 10

	res, transactions, err := ah.transactionService.GetHistoryTransaction(userId, page, pageSize)
    if err != nil {
        return response.NewErrorResponse(c, err)
    }

	currentPage, allPages := ah.transactionService.CalculatePaginationValues(page, pageSize, transactions)
	nextPage := ah.transactionService.GetNextPage(currentPage, allPages)
	prevPage := ah.transactionService.GetPrevPage(currentPage)

	responseData := map[string]interface{}{
		"data": res,
		"pagination": map[string]int{
			"currentPage": currentPage,
			"nextPage":    nextPage,
			"prevPage":    prevPage,
			"allPages":    allPages,
		},
	}

    return response.NewSuccessResponse(c, responseData)
}

func (ah *TransactionHandler) GetReportTransaction(c echo.Context) error {
	_, role, creatorId, err := middleware.ExtractToken(c)
    if err != nil {
        return response.NewErrorResponse(c, err)
    }	
	fmt.Println(role)
	var constRole string
	constRole = consts.ProductCreator
	if role != constRole || role != constRole {
		fmt.Println("test")
		return response.NewErrorResponse(c, echo.ErrUnauthorized)
	} 

	res, err := ah.transactionService.GetReportTransaction(creatorId, role)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}