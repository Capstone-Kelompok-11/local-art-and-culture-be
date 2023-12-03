package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo"
)

type AddressHandler struct {
	AddressService services.IAddressService
}

func NewAddressHandler(iAddressService services.IAddressService) *AddressHandler {
	return &AddressHandler{AddressService: iAddressService}
}

func (ad *AddressHandler) CreateAddress(c echo.Context) error {
	var input request.Address
	c.Bind(&input)

	res, err := ad.AddressService.CreateAddress(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ad *AddressHandler) GetAllAddress(c echo.Context) error {
	nameFilter := c.QueryParam("name")
	res, err := ad.AddressService.GetAllAddress(nameFilter)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ad *AddressHandler) GetAddress(c echo.Context) error {
	id := c.Param("id")
	res, err := ad.AddressService.GetAddress(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ad *AddressHandler) UpdateAddress(c echo.Context) error {
	id := c.Param("id")
	var input request.Address
	c.Bind(&input)
	res, err := ad.AddressService.UpdateAddress(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ad *AddressHandler) DeleteAddress(c echo.Context) error {
	id := c.Param("id")
	res, err := ad.AddressService.DeleteAddress(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}