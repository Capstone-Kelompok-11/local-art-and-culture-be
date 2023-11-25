package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo"
)

type CreatorHandler struct {
	creatorService services.ICreatorService
}

func NewCreatorHandler(iCreatorService services.ICreatorService) *CreatorHandler {
	return &CreatorHandler{creatorService: iCreatorService}
}

func (ch *CreatorHandler) CreateCreator(c echo.Context) error {
	var input request.Creator
	c.Bind(&input)

	res, err := ch.creatorService.CreateCreator(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ch *CreatorHandler) GetAllCreator(c echo.Context) error {
	nameFilter := c.QueryParam("name")
	res, err := ch.creatorService.GetAllCreator(nameFilter)
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