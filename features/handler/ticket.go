package handler

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo/v4"
)

type TicketHandler struct {
	ticketService services.ITicketService
}

func NewTicketHandler(iTicketService services.ITicketService) *TicketHandler {
	return &TicketHandler{ticketService: iTicketService}
}

func (ti *TicketHandler) CreateTicket(c echo.Context) error {
	var input request.Ticket
	c.Bind(&input)
	res, err := ti.ticketService.CreateTicket(&input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ti *TicketHandler) GetAllTicket(c echo.Context) error {
	nameFilter := c.QueryParam("name")
	res, err := ti.ticketService.GetAllTicket(nameFilter)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ti *TicketHandler) GetTicket(c echo.Context) error {
	id := c.Param("id")
	res, err := ti.ticketService.GetTicket(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ti *TicketHandler) UpdateTicket(c echo.Context) error {
	id := c.Param("id")
	var input request.Ticket
	c.Bind(&input)

	res, err := ti.ticketService.UpdateTicket(id, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}

func (ti *TicketHandler) DeleteTicket(c echo.Context) error {
	id := c.Param("id")
	res, err := ti.ticketService.DeleteTicket(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}