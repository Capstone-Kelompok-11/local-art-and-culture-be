package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func TransactionRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewTransactionRepository(db)
	service := services.NewTransactionService(repository)
	handler := handler.NewTransactionHandler(service)

	e.POST("/transaction", handler.CreateTransaction)
	e.GET("/transaction", handler.GetAllTransaction)
	e.GET("/transaction/:id", handler.GetTransaction)
	e.PUT("/transaction/:id", handler.UpdateTransaction)
	e.DELETE("/transaction/:id", handler.DeleteTransaction)
}
