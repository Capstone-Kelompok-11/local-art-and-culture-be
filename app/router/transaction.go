package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"
	"lokasani/helpers/middleware"
	"lokasani/helpers/midtrans"
	
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TransactionRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewTransactionRepository(db)
	tDRepository := repositories.NewTransactionDetailRepository(db)
	service := services.NewTransactionService(repository, tDRepository, midtrans.NewMidtransService(repository))
	handler := handler.NewTransactionHandler(service)

	eJwt := e.Group("")
	eJwt.Use(middleware.JWTMiddleware())

	
	e.POST("/transaction/payment-callback", handler.ConfirmPayment)

	//with token
	eJwt.POST("/transaction", handler.CreateTransaction)
	eJwt.GET("/transaction", handler.GetAllTransaction)
	eJwt.GET("/transaction/:id", handler.GetTransaction)
	eJwt.GET("/transaction/history", handler.GetHistoryTransaction)
	eJwt.GET("/transaction/report", handler.GetReportTransaction)
	eJwt.PUT("/transaction/:id", handler.UpdateTransaction)
	eJwt.DELETE("/transaction/:id", handler.DeleteTransaction)
}
