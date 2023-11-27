package handler


import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo"
)

type TransactionDetailHandler struct {
	transactionDetailService services.ITransactionDetailService
}

func NewTransactionDetailHandler(ITransactionDetailService services.ITransactionDetailService) *TransactionDetailHandler {
	return &TransactionDetailHandler{transactionDetailService: ITransactionDetailService}
}

func (ah *TransactionDetailHandler) CreateTransactionDetail(c echo.Context) error {
	var input request.TransactionDetail
	c.Bind(&input)

	res, err := ah.transactionDetailService.CreateTransactionDetail(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *TransactionDetailHandler) GetAllTransactionDetail(c echo.Context) error {
	res, err := ah.transactionDetailService.GetAllTransactionDetail()
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *TransactionDetailHandler) GetTransactionDetail(c echo.Context) error {
	id := c.Param("id")
	res, err := ah.transactionDetailService.GetTransactionDetail(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *TransactionDetailHandler) UpdateTransactionDetail(c echo.Context) error {
	id := c.Param("id")
	var input request.TransactionDetail
	c.Bind(&input)
	res, err := ah.transactionDetailService.UpdateTransactionDetail(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *TransactionDetailHandler) DeleteTransactionDetail(c echo.Context) error {
	id := c.Param("id")
	res, err := ah.transactionDetailService.DeleteTransactionDetail(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}
