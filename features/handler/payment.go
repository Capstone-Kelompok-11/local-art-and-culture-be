package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	paymentService services.IPaymentService
}

func NewPaymentHandler(iPaymentService services.IPaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: iPaymentService}
}

func (pa *PaymentHandler) CreatePayment(c echo.Context) error {
	var input request.Payment
	c.Bind(&input)

	res, err := pa.paymentService.CreatePayment(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (pa *PaymentHandler) GetAllPayment(c echo.Context) error {
	res, err := pa.paymentService.GetAllPayment()
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (pa *PaymentHandler) GetPayment(c echo.Context) error {
	id := c.Param("id")
	res, err := pa.paymentService.GetPayment(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (pa *PaymentHandler) UpdatePayment(c echo.Context) error {
	id := c.Param("id")
	var input request.Payment
	c.Bind(&input)
	res, err := pa.paymentService.UpdatePayment(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (pa *PaymentHandler) DeletePayment(c echo.Context) error {
	id := c.Param("id")
	res, err := pa.paymentService.DeletePayment(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}