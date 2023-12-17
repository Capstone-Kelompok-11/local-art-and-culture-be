package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/response"
)

func ConvertUserToAuthRes(user *models.Users) response.Auth {
	return response.Auth{
		Id: user.ID,
		Name: user.FirstName + user.LastName,
		Email: user.Email,
	}
}

func AuthResToAuthTokenRes(auth response.Auth, token string) *response.AuthToken {
    return &response.AuthToken{
        Name:  auth.Name,
        Email: auth.Email,
        Token: token,
    }
}