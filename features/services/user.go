package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/bcrypt"
	"lokasani/helpers/errors"
	"math"
)

type IUserService interface {
	RegisterUser(data *request.User) (response.User, error)
	LoginUser(data *request.User) (response.Creators, error)
	GetAllUser(nameFilter string, page, pageSize int) ([]response.User, map[string]int, error)
	GetUser(id string) (response.User, error)
	UpdateUser(id string, input request.User) (response.User, error)
	DeleteUser(id string) (response.User, error)
	CalculatePaginationValues(page, pageSize, allItmes int) (int, int)
	GetNextPage(currentPage, allPages int) int
	GetPrevPage(currentPage int) int
	CountUsersByRole(roleId uint) (int, error)
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

	hashPass, err := bcrypt.HashPassword(data.Password)
	if err != nil {
		return response.User{}, errors.ERR_BCRYPT_PASSWORD
	}

	data.Password = hashPass
	res, err := u.UserRepo.RegisterUser(data)
	if err != nil {
		return response.User{}, err
	}
	return res, nil
}

func (u *UserService) LoginUser(data *request.User) (response.Creators, error) {
	if data.Email == "" {
		return response.Creators{}, errors.ERR_EMAIL_IS_EMPTY
	}
	if data.Password == "" {
		return response.Creators{}, errors.ERR_PASSWORD_IS_EMPTY
	}

	res, err := u.UserRepo.LoginUser(data)
	if err != nil {
		return response.Creators{}, err
	}
	return res, nil
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

func (u *UserService) CountUsersByRole(roleId uint) (int, error) {
	count, err := u.UserRepo.CountUsersByRole(roleId)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (u *UserService) GetAllUser(nameFilter string, page, pageSize int) ([]response.User, map[string]int, error) {
	allUsers, _, err := u.UserRepo.GetAllUser(nameFilter, page, pageSize)
	if err != nil {
		return nil, nil, err
	}

	productCreators, err := u.CountUsersByRole(1)
	if err != nil {
		return nil, nil, err
	}

	regularUser, err := u.CountUsersByRole(0)
	if err != nil {
		return nil, nil, err
	}

	eventCreators, err := u.CountUsersByRole(3)
	if err != nil {
		return nil, nil, err
	}

	rolesCount := make(map[string]int)
	rolesCount["RegularUser"] = regularUser
	rolesCount["EventCreators"] = eventCreators
	rolesCount["ProductCreators"] = productCreators

	return allUsers, rolesCount, nil
}

func (u *UserService) CalculatePaginationValues(page, pageSize, allItmes int) (int, int) {
	pageInt := page
	if pageInt <= 0 {
		pageInt = 1
	}

	allPages := int(math.Ceil(float64(allItmes) / float64(pageSize)))

	if pageInt > allPages {
		pageInt = allPages
	}

	return pageInt, allPages
}

func (u *UserService) GetNextPage(currentPage, allPages int) int {
	if currentPage < allPages {
		return currentPage + 1
	}
	return allPages
}

func (u *UserService) GetPrevPage(currentPage int) int {
	if currentPage > 1 {
		return currentPage - 1
	}
	return 1
}
