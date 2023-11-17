package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo"
)

type AdminHandler struct {
	adminService services.IAdminService
}

func NewAdminHandler(iAdminService services.IAdminService) *AdminHandler {
	return &AdminHandler{adminService: iAdminService}
}

func (ah *AdminHandler) RegisterAdmin(c echo.Context) error {
	var input request.Admin
	c.Bind(&input)

	res, err := ah.adminService.RegisterAdmin(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *AdminHandler) LoginAdmin(c echo.Context) error {
	var input request.Admin
	c.Bind(&input)

	res, err := ah.adminService.LoginAdmin(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *AdminHandler) GetAllAdmin(c echo.Context) error {
	res, err := ah.adminService.GetAllAdmin()
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *AdminHandler) GetAdmin(c echo.Context) error {
	id := c.Param("id")
	res, err := ah.adminService.GetAdmin(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *AdminHandler) UpdateAdmin(c echo.Context) error {
	id := c.Param("id")
	var input request.Admin
	c.Bind(&input)
	res, err := ah.adminService.UpdateAdmin(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *AdminHandler) DeleteAdmin(c echo.Context) error {
	id := c.Param("id")
	res, err := ah.adminService.DeleteAdmin(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}
