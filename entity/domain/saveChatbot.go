package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/response"
)

func ConvertFromSaveReqToModel(message, response string) *models.SaveChatbot{
	return &models.SaveChatbot{
		Message: 	message,
		Response: 	response,
		UserId: 	1,
	}
}

func ConvertFromModelToSaveRes(data models.SaveChatbot) *response.SaveChatbot{
	return &response.SaveChatbot{
		Id: 		data.ID,
		Message: 	data.Message,
		Response: 	data.Response,
		UserId: 	data.UserId,
	}
}