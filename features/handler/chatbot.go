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

func NewChatbotHandler(iChatbotService services.IChatbotService) *ChatbotHandler {
	return &ChatbotHandler{chatbotService: iChatbotService}
}

func (ch *ChatbotHandler) Chatbot(c echo.Context) error {
	var input request.Chatbot
	c.Bind(&input)
	client := config.OpenAiClient()

	res, err := ch.chatbotService.Chatbot(*client, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}
