package handler

import (
	"fmt"
	"log"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	//"lokasani/helpers/consts"
	consts "lokasani/helpers/const"
	"lokasani/helpers/middleware"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EventHandler struct {
	eventService services.IEventService
}

func NewEventHandler(iEventService services.IEventService) *EventHandler {
	return &EventHandler{eventService: iEventService}
}

func (pr *EventHandler) CreateEvent(c echo.Context) error {
	_, roleId, _, err := middleware.ExtractToken(c)
    if err != nil {
		log.Println("Error extracting token:", err)
        return response.NewErrorResponse(c, err)
    }
    if roleId != consts.EventCreator {
		return response.NewErrorResponse(c, echo.ErrUnauthorized)
	}

    var input request.Event
    c.Bind(&input)
    res, err := pr.eventService.CreateEvent(&input)
    if err != nil {
        fmt.Println(input)
        return response.NewErrorResponse(c, err)
    }
    return response.NewSuccessResponse(c, res)
}

func (pr *EventHandler) GetAllEvent(c echo.Context) error {
	nameFilter := c.QueryParam("name")
	startDate := c.QueryParam("startDate")
    endDate := c.QueryParam("endDate")
	page, pageSize := 1, 10

	res, allItems, err := pr.eventService.GetAllEvent(nameFilter, startDate, endDate, page, pageSize)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}

	currentPage, allPages := pr.eventService.CalculatePaginationValues(page, pageSize, allItems)
	nextPage := pr.eventService.GetNextPage(currentPage, allPages)
	prevPage := pr.eventService.GetPrevPage(currentPage)

	responseData := map[string]interface{}{
		"data": res,
		"pagination": map[string]int{
			"currentPage": currentPage,
			"nextPage":    nextPage,
			"prevPage":    prevPage,
			"allPages":    allPages,
		},
	}

	return response.NewSuccessResponse(c, responseData)
}

func (pr *EventHandler) GetAllAvailableEvent(c echo.Context) error {
	nameFilter := c.QueryParam("name")
	startDate := c.QueryParam("startDate")
	endDate := c.QueryParam("endDate")
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
	   page = 1
	}
	pageSize, err := strconv.Atoi(c.QueryParam("pageSize"))
	if err != nil {
	   pageSize = 10
	}

	res, allItems, err := pr.eventService.GetAllAvailableEvent(nameFilter, startDate, endDate, page, pageSize)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}

	currentPage, allPages := pr.eventService.CalculatePaginationValues(page, pageSize, allItems)
	nextPage := pr.eventService.GetNextPage(currentPage, allPages)
	prevPage := pr.eventService.GetPrevPage(currentPage)

	responseData := map[string]interface{}{
		"data": res,
		"pagination": map[string]int{
			"currentPage": currentPage,
			"nextPage":    nextPage,
			"prevPage":    prevPage,
			"allPages":    allPages,
		},
	}

	return response.NewSuccessResponse(c, responseData)
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
	_, roleId, _, err := middleware.ExtractToken(c)
    if err != nil {
		log.Println("Error extracting token:", err)
        return response.NewErrorResponse(c, err)
    }
    if roleId != consts.EventCreator {
		return response.NewErrorResponse(c, echo.ErrUnauthorized)
	}

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
	_, roleId, _, err := middleware.ExtractToken(c)
    if err != nil {
		log.Println("Error extracting token:", err)
        return response.NewErrorResponse(c, err)
    }
    if roleId != consts.EventCreator {
		return response.NewErrorResponse(c, echo.ErrUnauthorized)
	}
	
	id := c.Param("id")
	res, err := pr.eventService.DeleteEvent(id)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}