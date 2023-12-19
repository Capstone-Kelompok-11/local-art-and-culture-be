package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/response"
)

func ConvertFromSaveReqToModel(message, response string, userId uint) *models.SaveChatbot {
	return &models.SaveChatbot{
		Message:  message,
		Response: response,
		UserId:   userId,
	}
}

func ConvertFromModelToSaveRes(data models.SaveChatbot) *response.SaveChatbot {
	return &response.SaveChatbot{
		Id:       data.ID,
		Message:  data.Message,
		Response: data.Response,
	}
}

func ListConvertFromModelToSaveRes(data []models.SaveChatbot) []*response.SaveChatbot {
	list := []*response.SaveChatbot{}
	for _, v := range data {
		res := ConvertFromModelToSaveRes(v)
		list = append(list, res)
	}

	return list
}
