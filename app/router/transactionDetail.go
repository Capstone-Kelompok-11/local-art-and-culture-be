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

func TransactionDetailRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewTransactionDetailRepository(db)
	service := services.NewTransactionDetailService(repository)
	handler := handler.NewTransactionDetailHandler(service)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))

	eJwt.POST("/transactionDetail", handler.CreateTransactionDetail)
	eJwt.GET("/transactionDetail", handler.GetAllTransactionDetail)
	eJwt.GET("/transactionDetail/:id", handler.GetTransactionDetail)
	eJwt.PUT("/transactionDetail/:id", handler.UpdateTransactionDetail)
	eJwt.DELETE("/transactionDetail/:id", handler.DeleteTransactionDetail)
}
