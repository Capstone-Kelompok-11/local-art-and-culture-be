package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TicketRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewTicketRepository(db)
	service := services.NewTicketService(repository)
	handler := handler.NewTicketHandler(service)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))

	eJwt.POST("/ticket", handler.CreateTicket)
	eJwt.GET("/ticket", handler.GetAllTicket)
	eJwt.GET("/ticket/:id", handler.GetTicket)
	eJwt.PUT("/ticket/:id", handler.UpdateTicket)
	eJwt.DELETE("/ticket/:id", handler.DeleteTicket)
}