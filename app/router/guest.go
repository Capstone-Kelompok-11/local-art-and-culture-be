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

func GuestRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewGuestRepository(db)
	service := services.NewGuestService(repository)
	handler := handler.NewGuestHandler(service)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))
	
	e.POST("/guest", handler.CreateGuest)
	e.GET("/guest", handler.GetAllGuest)
	e.GET("/guest/:id", handler.GetGuest)
	e.PUT("/guest/:id", handler.UpdateGuest)
	e.DELETE("/guest/:id", handler.DeleteGuest)
}