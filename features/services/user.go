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
	RegisterUser(data *request.User) (response.User, error)
	LoginUser(data *request.User) (response.User, error)
	GetAllUser(nameFilter string) ([]response.User, error)
	GetUser(id string) (response.User, error)
	UpdateUser(id string, input request.User) (response.User, error)
	DeleteUser(id string) (response.User, error)
}

type UserService struct {
	UserRepo repositories.IUserRepository
}

func NewUserService(repo repositories.IUserRepository) *UserService {
	return &UserService{UserRepo: repo}
}

func (u *UserService) RegisterUser(data *request.User) (response.User, error) {
	if data.FirstName == "" {
		return response.User{}, errors.ERR_NAME_IS_EMPTY
	}
	if data.LastName == "" {
		return response.User{}, errors.ERR_NAME_IS_EMPTY
	}
	if data.Email == "" {
		return response.User{}, errors.ERR_EMAIL_IS_EMPTY
	}
	if data.PhoneNumber == "" {
		return response.User{}, errors.ERR_PHONE_NUMBER_IS_EMPTY
	}

	hashPass, err := bcrypt.Hash(data.Password)
	if err != nil {
		return response.User{}, errors.ERR_BCRYPT_PASSWORD
	}

	data.Password = hashPass
	res, err := u.UserRepo.RegisterUser(data)
	if err != nil {
		return response.User{}, errors.ERR_REGISTER_USER_DATABASE
	}

	token, err := middleware.CreateToken(int(data.Id), data.Email)
	if err != nil {
		return response.User{}, errors.ERR_TOKEN
	}

	res.Token = token
	return res, nil
}

func (u *UserService) LoginUser(data *request.User) (response.User, error) {
	if data.Email == "" {
		return response.User{}, errors.ERR_EMAIL_IS_EMPTY
	} else if data.Password == "" {
		return response.User{}, errors.ERR_PASSWORD_IS_EMPTY
	}

	res, err := u.UserRepo.LoginUser(data)
	if err != nil {
		return response.User{}, err
	}

	token, err := middleware.CreateToken(int(data.Id), data.Email)
	if err != nil {
		return response.User{}, errors.ERR_TOKEN
	}

	res.Token = token
	return res, nil
}

func (u *UserService) GetAllUser(nameFilter string) ([]response.User, error) {
	err, res := u.UserRepo.GetAllUser(nameFilter)
	if err != nil {
		return err, nil
	}
	return nil, res
}

func (u *UserService) GetUser(id string) (response.User, error) {
	if id == "" {
		return response.User{}, errors.ERR_GET_USER_BAD_REQUEST_ID
	}

	res, err := u.UserRepo.GetUser(id)
	if err != nil {
		return response.User{}, err
	}
	return res, nil
}

func (u *UserService) UpdateUser(id string, data request.User) (response.User, error) {
	if id == "" {
		return response.User{}, errors.ERR_GET_USER_BAD_REQUEST_ID
	}

	res, err := u.UserRepo.UpdateUser(id, data)
	if err != nil {
		return response.User{}, errors.ERR_UPDATE_DATA
	}
	return res, nil
}

func (u *UserService) DeleteUser(id string) (response.User, error) {
	if id == "" {
		return response.User{}, errors.ERR_GET_USER_BAD_REQUEST_ID
	}

	res, err := u.UserRepo.DeleteUser(id)
	if err != nil {
		return response.User{}, errors.ERR_DELETE_USER
	}
	return res, nil
}