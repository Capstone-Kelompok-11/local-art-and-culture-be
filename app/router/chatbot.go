package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

func ChatbotRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewSaveRepository(db)
	serviceSave := services.NewSaveService(repository)
	service := services.NewChatbotService(&openai.Client{})
	handler := handler.NewChatbotHandler(service, serviceSave)

	eJwt := e.Group("")
	eJwt.Use(echojwt.JWT([]byte(os.Getenv("SECRET_JWT"))))
	
	eJwt.POST("/chatbot", handler.Chatbot)
	//eJwt.GET("/chatbot", handler.GetAllChatbot)
}