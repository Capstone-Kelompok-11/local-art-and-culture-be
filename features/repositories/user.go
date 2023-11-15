package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/bcrypt"
	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type IUserRepository interface {
	RegisterUser(data *request.User) (error, response.User)
	LoginUser(data *request.User) (error, response.User)
	GetAllUser() (error, []response.User)
	GetUser(id string) (error, response.User)
	UpdateUser(id string, input request.User) (error, response.User)
	DeleteUser(id string) (error, response.User)
}

type userRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) RegisterUser(data *request.User) (error, response.User) {
	dataUser := domain.ConvertFromUserReqToModel(*data)
	err := u.db.Create(&dataUser).Error
	if err != nil {
		return err, response.User{}
	}
	return nil, *domain.ConvertFromModelToUserRes(*dataUser)
}

func (u *userRepository) LoginUser(data *request.User) (error, response.User) {
	dataUser := domain.ConvertFromUserReqToModel(*data)
	err := u.db.Where("email = ? ", data.Email).First(&dataUser).Error
	if err != nil {
		return errors.ERR_EMAIL_NOT_FOUND, response.User{}
	}

	err = bcrypt.CheckPassword(data.Password, dataUser.Password)
	if err != nil {
		return errors.ERR_WRONG_PASSWORD, response.User{}
	}
	return nil, *domain.ConvertFromModelToUserRes(*dataUser)
}

func (u *userRepository) GetAllUser() (error, []response.User) {
	var allUser []models.Users
	var resAllUser []response.User
	err := u.db.Find(&allUser).Error
	if err != nil {
		return errors.ERR_GET_DATA, nil
	}

	for i := 0; i < len(allUser); i++ {
		userVm := domain.ConvertFromModelToUserRes(allUser[i])
		resAllUser = append(resAllUser, *userVm)
	}

	return nil, resAllUser
}

func (u *userRepository) GetUser(id string) (error, response.User) {
	var userData models.Users
	err := u.db.First(&userData, "id = ?", id).Error
	if err != nil {
		return err, response.User{}
	}
	return nil, *domain.ConvertFromModelToUserRes(userData)
}

func (u *userRepository) UpdateUser(id string, input request.User) (error, response.User) {
	userData := models.Users{}
	err := u.db.First(&userData, "id = ?", id).Error

	if err != nil {
		return err, response.User{}
	}
	
	if input.FirstName != "" {
		userData.FirstName = input.FirstName
	} else if input.LastName != "" {
		userData.LastName = input.LastName
	} else if input.Email != "" {
		userData.Email = input.Email
	} else if input.Password != "" {
		userData.Password = input.Password
	}else if input.PhoneNumber != "" {
		userData.PhoneNumber = input.PhoneNumber
	}

	if err = u.db.Save(&userData).Error; err != nil {
		return err, response.User{}
	}
	return nil, *domain.ConvertFromModelToUserRes(userData)
}

func (u *userRepository) DeleteUser(id string) (error, response.User) {
	userData := models.Users{}
	res := response.User{}
	find := u.db.First(&userData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToUserRes(userData)
	}
	err := u.db.Delete(&userData, "id = ?", id).Error
	if err != nil {
		return err, response.User{}
	}
	return nil, res
}