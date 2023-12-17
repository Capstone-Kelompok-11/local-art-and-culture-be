package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo/v4"
)

type ShippingHandler struct {
	shippingService services.IShippingService
}

func NewShippingHandler(iShippingService services.IShippingService) *ShippingHandler {
	return &ShippingHandler{shippingService: iShippingService}
}

func (ah *ShippingHandler) CreateShipping(c echo.Context) error {
	var input request.Shipping
	c.Bind(&input)
	res, err := ah.shippingService.CreateShippingMethod(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *ShippingHandler) GetAllShipping(c echo.Context) error {
	res, err := ah.shippingService.GetAllShipping()
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *ShippingHandler) GetShipping(c echo.Context) error {
	id := c.Param("id")
	res, err := ah.shippingService.GetShipping(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *ShippingHandler) UpdateShipping(c echo.Context) error {
	id := c.Param("id")
	var input request.Shipping
	c.Bind(&input)
	res, err := ah.shippingService.UpdateShipping(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *ShippingHandler) DeleteShipping(c echo.Context) error {
	id := c.Param("id")
	res, err := ah.shippingService.DeleteShipping(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}