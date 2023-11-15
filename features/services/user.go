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
	RegisterUser(data *request.User) (error, response.User)
	LoginUser(data *request.User) (error, response.User)
	GetAllUser() (error, []response.User)
	GetUser(id string) (error, response.User)
	UpdateUser(id string, input request.User) (error, response.User)
	DeleteUser(id string) (error, response.User)
}

type UserService struct {
	UserRepo repositories.IUserRepository
}

func NewUserService(repo repositories.IUserRepository) *UserService {
	return &UserService{UserRepo: repo}
}

func (u *UserService) RegisterUser(data *request.User) (error, response.User) {
	if data.FirstName == "" {
		return errors.ERR_NAME_IS_EMPTY, response.User{}
	} else if data.LastName == "" {
		return errors.ERR_NAME_IS_EMPTY, response.User{}
	} else if data.Email == "" {
		return errors.ERR_EMAIL_IS_EMPTY, response.User{}
	} else if data.PhoneNumber == "" {
		return errors.ERR_PHONE_NUMBER_IS_EMPTY, response.User{}
	}

	hashPass, err := bcrypt.Hash(data.Password)
	if err != nil {
		return errors.ERR_BCRYPT_PASSWORD, response.User{}
	}

	data.Password = hashPass
	err, res := u.UserRepo.RegisterUser(data)
	if err != nil {
		return errors.ERR_REGISTER_USER_DATABASE, response.User{}
	}

	token, err := middleware.CreateToken(int(data.Id), data.Email)
	if err != nil {
		return errors.ERR_TOKEN, response.User{}
	}

	res.Token = token
	return nil, res
}

func (u *UserService) LoginUser(data *request.User) (error, response.User) {
	if data.Email == "" {
		return errors.ERR_EMAIL_IS_EMPTY, response.User{}
	} else if data.Password == "" {
		return errors.ERR_PASSWORD_IS_EMPTY, response.User{}
	}

	err, res := u.UserRepo.LoginUser(data)
	if err != nil {
		return err, response.User{}
	}

	token, err := middleware.CreateToken(int(data.Id), data.Email)
	if err != nil {
		return errors.ERR_TOKEN, response.User{}
	}

	res.Token = token
	return nil, res
}

func (u *UserService) GetAllUser() (error, []response.User) {
	err, res := u.UserRepo.GetAllUser()
	if err != nil {
		return err, nil
	}
	return nil, res
}

func (u *UserService) GetUser(id string) (error, response.User) {
	if id == "" {
		return errors.ERR_GET_USER_BAD_REQUEST_ID, response.User{}
	}

	err, res := u.UserRepo.GetUser(id)
	if err != nil {
		return err, response.User{}
	}
	return nil, res
}

func (u *UserService) UpdateUser(id string, data request.User) (error, response.User) {
	if id == "" {
		return errors.ERR_GET_USER_BAD_REQUEST_ID, response.User{}
	}

	err, res := u.UserRepo.UpdateUser(id, data)
	if err != nil {
		return errors.ERR_UPDATE_DATA, response.User{}
	}
	return nil, res
}

func (u *UserService) DeleteUser(id string) (error, response.User) {
	if id == "" {
		return errors.ERR_GET_USER_BAD_REQUEST_ID, response.User{}
	}

	err, res := u.UserRepo.DeleteUser(id)
	if err != nil {
		return errors.ERR_DELETE_USER, response.User{}
	}
	return nil, res
}
