package handler

import (
	"lokasani/entity/request"
	"lokasani/features/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService services.IUserService
}

func NewUserHandler(userService services.IUserService) *userHandler {
	return &userHandler{userService}
}

func (u *userHandler) RegisterUsers(e echo.Context) error {
	user := new(request.UserRequest)
	if err := e.Bind(user); err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid request")
	}

	userRes, err := u.userService.CreateUser(user)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, "Failed to create user")
	}

	return e.JSON(http.StatusOK, userRes)
}