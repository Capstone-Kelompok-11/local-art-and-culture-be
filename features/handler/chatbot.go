package handler

import (
	//"fmt"
	"lokasani/app/drivers/config"
	"lokasani/entity/domain"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/services"
	// "lokasani/helpers/middleware"
	// "strconv"

	// "lokasani/helpers/middleware"
	// "strconv"

	"github.com/labstack/echo/v4"
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

	// userID, _, _, _, err := middleware.ExtractToken(c)

	// if err != nil {
	// 	return response.NewErrorResponse(c, err)
	// }
	// input.UserId = userID

	res, err := ch.chatbotService.Chatbot(*client, input)
	if err != nil {
		return response.NewErrorResponse(c, err)
	}

	history := domain.ConvertFromSaveReqToModel(input.Message, res.Message, input.UserId)
	ch.save.SaveChatbot(*history)
	return response.NewSuccessResponse(c, res)
}

// func (ch *ChatbotHandler) GetAllChatbot(c echo.Context) error {
// 	// userID, _, _, err := middleware.ExtractToken(c)
// 	// if err != nil {
// 	// 	return response.NewErrorResponse(c, err)
// 	// }
// 	// fmt.Println(userID)
// 	//userIDString := strconv.FormatUint(uint64(userID), 10)

// 	// userIDUint, err := strconv.ParseUint(userIDString, 10, 64)
// 	// if err != nil {
// 	// 	return response.NewErrorResponse(c, err)
// 	// }

// 	// res, err := ch.save.GetAllChatbot(uint(userIDUint))
// 	// if err != nil {
// 	// 	return response.NewErrorResponse(c, err)
// 	// }
// 	// return response.NewSuccessResponse(c, res)
// }