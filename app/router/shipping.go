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

func ShippingRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewShippingRepository(db)
	service := services.NewShippingService(repository)
	handler := handler.NewShippingHandler(service)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))

	eJwt.POST("/shippingMethod", handler.CreateShipping)
	eJwt.GET("/shippingMethod", handler.GetAllShipping)
	eJwt.GET("/shippingMethod/:id", handler.GetShipping)
	eJwt.PUT("/shippingMethod/:id", handler.UpdateShipping)
	eJwt.DELETE("/shippingMethod/:id", handler.DeleteShipping)
}