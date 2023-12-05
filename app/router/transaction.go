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

func TransactionRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewTransactionRepository(db)
	service := services.NewTransactionService(repository)
	handler := handler.NewTransactionHandler(service)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))

	eJwt.POST("/transaction", handler.CreateTransaction)
	eJwt.GET("/transaction", handler.GetAllTransaction)
	eJwt.GET("/transaction/:id", handler.GetTransaction)
	eJwt.PUT("/transaction/:id", handler.UpdateTransaction)
	eJwt.DELETE("/transaction/:id", handler.DeleteTransaction)
}
