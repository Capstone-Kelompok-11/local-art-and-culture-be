package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/bcrypt"
	"lokasani/helpers/errors"
	"lokasani/helpers/middleware"
)

type IUserService interface {
	RegisterUser(data *request.UserRequest) (error, response.UserResponse)
	LoginUser(data *request.UserRequest) (error, response.UserResponse)
}

type UserService struct {
	UserRepo repositories.IUserRepository
}

func NewUserService(repo repositories.IUserRepository) *UserService {
	return &UserService{UserRepo: repo}
}

func (u *UserService) RegisterUser(data *request.UserRequest) (error, response.UserResponse) {
    if data.FirstName == "" {
        return errors.ERR_NAME_IS_EMPTY, response.UserResponse{}
    } else if data.LastName == "" {
        return errors.ERR_NAME_IS_EMPTY, response.UserResponse{}
    } else if data.Email == "" {
        return errors.ERR_EMAIL_IS_EMPTY, response.UserResponse{}
    }

    hashPass, err := bcrypt.Hash(data.Password)
    if err != nil {
        return errors.ERR_BCRYPT_PASSWORD, response.UserResponse{}
    }

    data.Password = hashPass
    err, res := u.UserRepo.RegisterUser(data)
    if err != nil {
        return errors.ERR_REGISTER_USER_DATABASE, response.UserResponse{}
    }

    token, err := middleware.CreateToken(int(data.Id), data.Email)
    if err != nil {
        return errors.ERR_TOKEN, response.UserResponse{}
    }

    res.Token = token
    return nil, res
}

func (u *UserService) LoginUser(data *request.UserRequest) (error, response.UserResponse) {
    if data.FirstName == "" {
        return errors.ERR_NAME_IS_EMPTY, response.UserResponse{}
    } else if data.LastName == "" {
        return errors.ERR_NAME_IS_EMPTY, response.UserResponse{}
    } else if data.Email == "" {
        return errors.ERR_EMAIL_IS_EMPTY, response.UserResponse{}
    }

    hashPass, err := bcrypt.Hash(data.Password)
    if err != nil {
        return errors.ERR_BCRYPT_PASSWORD, response.UserResponse{}
    }

    data.Password = hashPass
    err, res := u.UserRepo.LoginUser(data)
    if err != nil {
        return errors.ERR_REGISTER_USER_DATABASE, response.UserResponse{}
    }

    token, err := middleware.CreateToken(int(data.Id), data.Email)
    if err != nil {
        return errors.ERR_TOKEN, response.UserResponse{}
    }

    res.Token = token
    return nil, res
}