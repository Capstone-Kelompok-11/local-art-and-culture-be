package handler

import (
	"fmt"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"
	"lokasani/helpers/middleware"

	"github.com/labstack/echo/v4"
)

type CreatorHandler struct {
	creatorService services.ICreatorService
}

func NewCreatorHandler(iCreatorService services.ICreatorService) *CreatorHandler {
	return &CreatorHandler{creatorService: iCreatorService}
}

func (ch *CreatorHandler) CreateCreator(c echo.Context) error {
	userID, _, _, err := middleware.ExtractToken(c)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}

	var input request.Creator
	if err := c.Bind(&input); err != nil {
		return response.NewErrorResponse(c, err)
	}
	input.UserId = userID

	fmt.Println("UserID from token:", userID)


	//fmt.Println("UserID in struct:", input.Users.Id)

	res, err := ch.creatorService.CreateCreator(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ch *CreatorHandler) GetAllCreator(c echo.Context) error {
	var filter request.Creator
	//filter.Users.Username = c.QueryParam("name")
	res, err := ch.creatorService.GetAllCreator(filter)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ch *CreatorHandler) GetAllCreatorByRole(c echo.Context) error {
	var filter request.Creator
	filter.Role.Role = c.QueryParam("role")
	fmt.Println(filter.Role.Role)
	res, err := ch.creatorService.GetAllCreatorByRole(filter)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ch *CreatorHandler) GetCreator(c echo.Context) error {
	id := c.Param("id")
	res, err := ch.creatorService.GetCreator(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ch *CreatorHandler) UpdateCreator(c echo.Context) error {
	id := c.Param("id")
	var input request.Creator
	c.Bind(&input)

	res, err := ch.creatorService.UpdateCreator(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ch *CreatorHandler) DeleteCreator(c echo.Context) error {
	id := c.Param("id")

	res, err := ch.creatorService.DeleteCreator(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}