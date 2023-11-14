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

	err, res := ch.creatorService.CreateCreator(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ch *CreatorHandler) GetAllCreator(c echo.Context) error {
	err, res := ch.creatorService.GetAllCreator()
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ch *CreatorHandler) GetCreator(c echo.Context) error {
	id := c.Param("id")
	err, res := ch.creatorService.GetCreator(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}