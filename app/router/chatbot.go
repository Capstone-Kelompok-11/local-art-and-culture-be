package routes

import (
	"lokasani/features/handler"
	"lokasani/features/repositories"
	"lokasani/features/services"

	"github.com/labstack/echo"
	"github.com/sashabaranov/go-openai"
	"gorm.io/gorm"
)

func ChatbotRoute(e *echo.Echo, db *gorm.DB) {
	repository := repositories.NewSaveRepository(db)
	serviceSave := services.NewSaveService(repository)
	service := services.NewChatbotService(&openai.Client{})
	handler := handler.NewChatbotHandler(service, serviceSave)

	e.POST("/chatbot", handler.Chatbot)
	e.GET("/chatbot", handler.GetAllChatbot)
}
