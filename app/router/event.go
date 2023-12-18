package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"
	"lokasani/helpers/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func EventRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewEventRepository(db)
	service := services.NewEventService(repository)
	handler := handler.NewEventHandler(service)

	eJwt := e.Group("")
	eJwt.Use(middleware.JWTMiddleware())

	//event creator can access
	eJwt.PUT("/event/:id", handler.UpdateEvent)
	eJwt.DELETE("/event/:id", handler.DeleteEvent)
	eJwt.POST("/event", handler.CreateEvent)

	//without token
	e.GET("/event", handler.GetAllEvent)
	e.GET("/event/home", handler.GetAllAvailableEvent)
	e.GET("/event/:id", handler.GetEvent)
}
