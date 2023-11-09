package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	helpers "lokasani/helpers/bcrypt"
)

type IUserService interface {
	CreateUser(data *request.UserRequest) (response.UserResponse, error)
	//Login(email string, password string) (domain.UsersDomain, error)
}

type userService struct {
	repo repositories.IUserRepository
}

func NewUserService(repo repositories.IUserRepository) *userService {
	return &userService{repo}
}

func (u *userService) CreateUser(data *request.UserRequest) (response.UserResponse, error) {
    hashedPassword, err := helpers.Hash(data.Password)
    if err != nil {
        return response.UserResponse{}, err
    }

    data.Password = hashedPassword

    userResponse, err := u.repo.CreateUser(data)
    if err != nil {
        return response.UserResponse{}, err
    }

    return userResponse, nil
}