package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"lokasani/helpers/errors"
	"lokasani/helpers/middleware"

	"github.com/labstack/echo/v4"
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

	token, err := middleware.CreateToken(uint(res.Id), 0, uint(res.Id))
	if err != nil {
		return response.NewErrorResponse(e, errors.ERR_TOKEN)
	}
	res.Token = token

	middleware.SetTokenCookie(e, token)
	return response.NewSuccessResponse(e, res)
}

func (u *UserHandler) LoginUsers(e echo.Context) error {
    var input request.User
    e.Bind(&input)

    res, err := u.userService.LoginUser(&input)
    if err != nil {
        return response.NewErrorResponse(e, err)
    }

    token, err := middleware.CreateToken(uint(res.Users.Id), 0, uint(res.Id))
    if err != nil {
        return response.NewErrorResponse(e, errors.ERR_TOKEN)
    }
    res.Users.Token = token

    middleware.SetTokenCookie(e, token)
	return response.NewSuccessResponse(e, res)
}

func (u *UserHandler) GetAllUser(c echo.Context) error {
	nameFilter := c.QueryParam("name")
	page, pageSize := 1, 10

	res, allItems, err := u.userService.GetAllUser(nameFilter, page, pageSize)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}

	currentPage, allPages := u.userService.CalculatePaginationValues(page, pageSize, allItems)
	nextPage := u.userService.GetNextPage(currentPage, allPages)
	prevPage := u.userService.GetPrevPage(currentPage)

	responseData := map[string]interface{}{
		"allUsers": res,
		"pagination": map[string]int{
			"currentPage": currentPage,
			"nextPage":    nextPage,
			"prevPage":    prevPage,
			"allPages":    allPages,
		},
	}
	return response.NewSuccessResponse(c, responseData)
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
	_, roleId, _, err := middleware.ExtractToken(c)
    if err != nil {
        return response.NewErrorResponse(c, err)
    }
    if roleId != 4 {
		return response.NewErrorResponse(c, echo.ErrUnauthorized)
	}
	
	id := c.Param("id")
	res, err := u.userService.DeleteUser(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}
