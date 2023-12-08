package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"
	"lokasani/helpers/midtrans"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TransactionRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewTransactionRepository(db)
	tDRepository := repositories.NewTransactionDetailRepository(db)
	service := services.NewTransactionService(repository, tDRepository, midtrans.NewMidtransService(repository))
	handler := handler.NewTransactionHandler(service)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))

	eJwt.POST("/transaction", handler.CreateTransaction)
	e.POST("/transaction/payment-callback", handler.ConfirmPayment)
	eJwt.GET("/transaction", handler.GetAllTransaction)
	eJwt.GET("/transaction/:id", handler.GetTransaction)
	eJwt.GET("/report/transaction", handler.GetTransactionReport)
	eJwt.PUT("/transaction/:id", handler.UpdateTransaction)
	eJwt.DELETE("/transaction/:id", handler.DeleteTransaction)
}
