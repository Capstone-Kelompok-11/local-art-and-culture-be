package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo/v4"
)

type GuestHandler struct {
	guestService services.IGuestService
}

func NewGuestHandler(iGuestService services.IGuestService) *GuestHandler {
	return &GuestHandler{guestService: iGuestService}
}

func (gu *GuestHandler) CreateGuest(c echo.Context) error {
	var input request.Guest
	c.Bind(&input)
	res, err := gu.guestService.CreateGuest(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (gu *GuestHandler) GetAllGuest(c echo.Context) error {
	nameFilter := c.QueryParam("name")

    res, err := gu.guestService.GetAllGuest(nameFilter)
    if err != nil {
        return response.NewErrorResponse(c, err)
    }
    return response.NewSuccessResponse(c, res)
}

func (gu *GuestHandler) GetGuest(c echo.Context) error {
	id := c.Param("id")
	res, err := gu.guestService.GetGuest(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (gu *GuestHandler) UpdateGuest(c echo.Context) error {
	id := c.Param("id")
	var input request.Guest
	c.Bind(&input)

	res, err := gu.guestService.UpdateGuest(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (gu *GuestHandler) DeleteGuest(c echo.Context) error {
	id := c.Param("id")
	res, err := gu.guestService.DeleteGuest(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}