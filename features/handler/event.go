package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo"
)

type EventHandler struct {
	eventService services.IEventService
}

func NewEventHandler(iEventService services.IEventService) *EventHandler {
	return &EventHandler{eventService: iEventService}
}

func (pr *EventHandler) CreateEvent(c echo.Context) error {
	var input request.Event
	c.Bind(&input)
	res, err := pr.eventService.CreateEvent(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (pr *EventHandler) GetAllEvent(c echo.Context) error {
	res, err := pr.eventService.GetAllEvent()
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (pr *EventHandler) GetEvent(c echo.Context) error {
	id := c.Param("id")
	res, err := pr.eventService.GetEvent(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (pr *EventHandler) UpdateEvent(c echo.Context) error {
	id := c.Param("id")
	var input request.Event
	c.Bind(&input)

	res, err := pr.eventService.UpdateEvent(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (pr *EventHandler) DeleteEvent(c echo.Context) error {
	id := c.Param("id")
	res, err := pr.eventService.DeleteEvent(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}
