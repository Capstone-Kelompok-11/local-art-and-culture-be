package services

import (
	"lokasani/entity/domain"
	"lokasani/features/repositories"
	"lokasani/helpers/bcrypt"
)

type IUserService interface {
	Create(user domain.UsersDomain) (domain.UsersDomain, error)
	Login(email string, password string) (domain.UsersDomain, error)
}

type userService struct {
	repo repositories.IUserRepository
}

func NewUserService(repo repositories.IUserRepository) *userService {
	return &userService{repo}
}

func (u *userService) Create(user domain.UsersDomain) (domain.UsersDomain, error) {
	user.Password = helpers.Hash(user.Password)
}