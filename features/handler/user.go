package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo"
)

type UserHandler struct {
	userService services.IUserService
}

func NewUserHandler(userService services.IUserService) *UserHandler {
	return &UserHandler{userService}
}

func (u *UserHandler) RegisterUsers(e echo.Context) error {
	var input request.UserRequest
	e.Bind(&input)

	err, res := u.userService.RegisterUser(&input)
	if err != nil {
		return response.NewErrorResponse(e, err)
	}
	return response.NewSuccessResponse(e, res)
}

func (u *UserHandler) LoginUsers(e echo.Context) error {
	var input request.UserRequest
	e.Bind(&input)

	err, res := u.userService.RegisterUser(&input)
	if err != nil {
		return response.NewErrorResponse(e, err)
	}
	return response.NewSuccessResponse(e, res)
}