package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func ShippingRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewShippingRepository(db)
	service := services.NewShippingService(repository)
	handler := handler.NewShippingHandler(service)

	e.POST("/shippingMethod", handler.CreateShipping)
	e.GET("/shippingMethod", handler.GetAllShipping)
	e.GET("/shippingMethod/:id", handler.GetShipping)
	e.PUT("/shippingMethod/:id", handler.UpdateShipping)
	e.DELETE("/shippingMethod/:id", handler.DeleteShipping)
}
