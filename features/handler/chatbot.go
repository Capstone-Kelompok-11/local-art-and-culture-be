package handler

import (
	"lokasani/app/drivers/config"
	"lokasani/entity/domain"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"

	"github.com/labstack/echo"
)

type ChatbotHandler struct {
	chatbotService services.IChatbotService
	save           services.ISaveService
}

func NewChatbotHandler(iChatbotService services.IChatbotService, save services.ISaveService) *ChatbotHandler {
	return &ChatbotHandler{
		chatbotService: iChatbotService,
		save: save,
	}
}

func (ch *ChatbotHandler) Chatbot(c echo.Context) error {
	var input request.Chatbot
	c.Bind(&input)
	client := config.OpenAiClient()

	res, err := ch.chatbotService.Chatbot(*client, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}

	history := domain.ConvertFromSaveReqToModel(input.Message, res.Message)
	ch.save.SaveChatbot(*history)
	return response.NewSuccessResponse(c, res)
}

func (ch *ChatbotHandler) GetAllChatbot(c echo.Context) error {
	res, err := ch.save.GetAllChatbot(1)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}
	return response.NewSuccessResponse(c, res)
}