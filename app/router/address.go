package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func AddressRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewAddressRepository(db)
	service := services.NewAddressService(repository)
	handler := handler.NewAddressHandler(service)

	e.POST("/address", handler.CreateAddress)
	e.GET("/address", handler.GetAllAddress)
	e.GET("/address/:id", handler.GetAddress)
	e.PUT("/address/:id", handler.UpdateAddress)
	e.DELETE("/address/:id", handler.DeleteAddress)
}