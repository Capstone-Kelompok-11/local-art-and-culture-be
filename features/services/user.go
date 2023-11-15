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
	GetAllUser() (error, []response.UserResponse)
	GetUser(id string) (error, response.UserResponse)
	UpdateUser(id string, input request.UserRequest) (error, response.UserResponse)
	DeleteUser(id string) (error, response.UserResponse)
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
	} else if data.PhoneNumber == "" {
		return errors.ERR_PHONE_NUMBER_IS_EMPTY, response.UserResponse{}
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
	if data.Email == "" {
		return errors.ERR_EMAIL_IS_EMPTY, response.UserResponse{}
	} else if data.Password == "" {
		return errors.ERR_PASSWORD_IS_EMPTY, response.UserResponse{}
	}

	err, res := u.UserRepo.LoginUser(data)
	if err != nil {
		return err, response.UserResponse{}
	}

	token, err := middleware.CreateToken(int(data.Id), data.Email)
	if err != nil {
		return errors.ERR_TOKEN, response.UserResponse{}
	}

	res.Token = token
	return nil, res
}

func (u *UserService) GetAllUser() (error, []response.UserResponse) {
	err, res := u.UserRepo.GetAllUser()
	if err != nil {
		return err, nil
	}
	return nil, res
}

func (u *UserService) GetUser(id string) (error, response.UserResponse) {
	if id == "" {
		return errors.ERR_GET_USER_BAD_REQUEST_ID, response.UserResponse{}
	}

	err, res := u.UserRepo.GetUser(id)
	if err != nil {
		return err, response.UserResponse{}
	}
	return nil, res
}

func (u *UserService) UpdateUser(id string, data request.UserRequest) (error, response.UserResponse) {
	if id == "" {
		return errors.ERR_GET_USER_BAD_REQUEST_ID, response.UserResponse{}
	}

	err, res := u.UserRepo.UpdateUser(id, data)
	if err != nil {
		return errors.ERR_UPDATE_DATA, response.UserResponse{}
	}
	return nil, res
}

func (u *UserService) DeleteUser(id string) (error, response.UserResponse) {
	if id == "" {
		return errors.ERR_GET_USER_BAD_REQUEST_ID, response.UserResponse{}
	}

	err, res := u.UserRepo.DeleteUser(id)
	if err != nil {
		return errors.ERR_DELETE_USER, response.UserResponse{}
	}
	return nil, res
}
