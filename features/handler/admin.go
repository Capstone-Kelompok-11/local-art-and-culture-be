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

	err, res := ah.adminService.RegisterAdmin(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ah *AdminHandler) LoginAdmin(c echo.Context) error {
	var input request.Admin
	c.Bind(&input)

	err, res := ah.adminService.LoginAdmin(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}
