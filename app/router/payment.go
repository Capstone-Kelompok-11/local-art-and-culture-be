package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func PaymentRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewPaymentRepository(db)
	service := services.NewPaymentService(repository)
	handler := handler.NewPaymentHandler(service)

	e.POST("/payment", handler.CreatePayment)
	e.GET("/payment", handler.GetAllPayment)
	e.GET("/payment/:id", handler.GetPayment)
	e.PUT("/payment/:id", handler.UpdatePayment)
	e.DELETE("/payment/:id", handler.DeletePayment)
}