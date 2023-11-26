package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo"
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
