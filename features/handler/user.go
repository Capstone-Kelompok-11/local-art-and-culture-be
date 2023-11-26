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
	var input request.User
	e.Bind(&input)

	res, err := u.userService.RegisterUser(&input)
	if err != nil {
		return response.NewErrorResponse(e, err)
	}
	return response.NewSuccessResponse(e, res)
}

func (u *UserHandler) LoginUsers(e echo.Context) error {
	var input request.User
	e.Bind(&input)

	res, err := u.userService.LoginUser(&input)
	if err != nil {
		return response.NewErrorResponse(e, err)
	}
	return response.NewSuccessResponse(e, res)
}

func (u *UserHandler) GetAllUser(c echo.Context) error {
	res, err := u.userService.GetAllUser()
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (u *UserHandler) GetUser(c echo.Context) error {
	id := c.Param("id")
	res, err := u.userService.GetUser(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (u *UserHandler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	var input request.User
	c.Bind(&input)
	res, err := u.userService.UpdateUser(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (u *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	res, err := u.userService.DeleteUser(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}