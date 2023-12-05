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

func EventRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewEventRepository(db)
	service := services.NewEventService(repository)
	handler := handler.NewEventHandler(service)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))
	
	eJwt.GET("/event", handler.GetAllEvent)
	eJwt.GET("/event/home", handler.GetAllAvailableEvent)
	eJwt.GET("/event/:id", handler.GetEvent)

	e.PUT("/event/:id", handler.UpdateEvent)
	e.DELETE("/event/:id", handler.DeleteEvent)
	e.POST("/event", handler.CreateEvent)
}
