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

func PaymentRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewPaymentRepository(db)
	service := services.NewPaymentService(repository)
	handler := handler.NewPaymentHandler(service)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))

	eJwt.POST("/payment", handler.CreatePayment)
	eJwt.GET("/payment", handler.GetAllPayment)
	eJwt.GET("/payment/:id", handler.GetPayment)
	eJwt.PUT("/payment/:id", handler.UpdatePayment)
	eJwt.DELETE("/payment/:id", handler.DeletePayment)
}