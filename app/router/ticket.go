package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func TicketRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewTicketRepository(db)
	service := services.NewTicketService(repository)
	handler := handler.NewTicketHandler(service)

	e.POST("/ticket", handler.CreateTicket)
	e.GET("/ticket", handler.GetAllTicket)
	e.GET("/ticket/:id", handler.GetTicket)
	e.PUT("/ticket/:id", handler.UpdateTicket)
	e.DELETE("/ticket/:id", handler.DeleteTicket)
}