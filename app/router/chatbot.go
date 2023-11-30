package routes

import (
	"lokasani/features/handler"
	"lokasani/features/services"

	"github.com/labstack/echo"
	"github.com/sashabaranov/go-openai"
)

func ChatbotRoute(e *echo.Echo) {
	service := services.NewChatbotService(&openai.Client{})
	handler := handler.NewChatbotHandler(service)

	e.POST("/chatbot", handler.Chatbot)
}
