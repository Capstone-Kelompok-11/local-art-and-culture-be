package handler

import (
	"fmt"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"
	"lokasani/helpers/errors"
	"lokasani/helpers/middleware"
	"time"

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
	userID, _, _, _, _, err := middleware.ExtractToken(c)
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
	fmt.Println("call back")
	err := c.Bind(&paymentPayload)
	if err != nil {
		fmt.Println("tessst")
		return response.NewErrorResponse(c, err)
	}
	orderId, exists := paymentPayload["order_id"].(string)

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

func (pr *TransactionHandler) GetTransactionReport(c echo.Context) error {
	startDateStr := c.QueryParam("start_date")
	endDateStr := c.QueryParam("end_date")

	if startDateStr == "" || endDateStr == "" {
		return response.NewErrorResponse(c, echo.ErrBadGateway)
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}

	if endDate.Before(startDate) {
		return response.NewErrorResponse(c, err)
	}

	transactions, err := pr.transactionService.GetTransactionReport(startDate, endDate)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}

	return response.NewSuccessResponse(c, transactions)
}
