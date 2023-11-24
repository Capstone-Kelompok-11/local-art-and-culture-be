package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func EventRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewEventRepository(db)
	service := services.NewEventService(repository)
	handler := handler.NewEventHandler(service)

	e.POST("/event", handler.CreateEvent)
	e.GET("/event", handler.GetAllEvent)
	e.GET("/event/:id", handler.GetEvent)
	e.PUT("/event/:id", handler.UpdateEvent)
	e.DELETE("/event/:id", handler.DeleteEvent)
}
