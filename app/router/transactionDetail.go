package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func TransactionDetailRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewTransactionDetailRepository(db)
	service := services.NewTransactionDetailService(repository)
	handler := handler.NewTransactionDetailHandler(service)

	e.POST("/transactionDetail", handler.CreateTransactionDetail)
	e.GET("/transactionDetail", handler.GetAllTransactionDetail)
	e.GET("/transactionDetail/:id", handler.GetTransactionDetail)
	e.PUT("/transactionDetail/:id", handler.UpdateTransactionDetail)
	e.DELETE("/transactionDetail/:id", handler.DeleteTransactionDetail)
}
