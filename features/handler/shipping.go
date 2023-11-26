package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo"
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
	err, res := ah.shippingService.GetAllShipping()
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *ShippingHandler) GetShipping(c echo.Context) error {
	id := c.Param("id")
	err, res := ah.shippingService.GetShipping(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *ShippingHandler) UpdateShipping(c echo.Context) error {
	id := c.Param("id")
	var input request.Shipping
	c.Bind(&input)
	err, res := ah.shippingService.UpdateShipping(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *ShippingHandler) DeleteShipping(c echo.Context) error {
	id := c.Param("id")
	err, res := ah.shippingService.DeleteShipping(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}