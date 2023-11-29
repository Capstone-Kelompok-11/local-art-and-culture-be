package handler

import (
	"lokasani/app/drivers/config"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo"
)

type ChatbotHandler struct {
	chatbotService services.IChatbotService
}

// func NewChatbotHandler(iChatbotService services.ChatbotService) *ChatbotHandler {
// 	return &ChatbotHandler{chatbotService: &iChatbotService}
// }

func (ch *ChatbotHandler) CreateChatbot(c echo.Context) error {
	var input request.Chatbot
	c.Bind(&input)
	client := config.OpenAiClient()

	res, err := ch.chatbotService.Chatbot(client, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}
