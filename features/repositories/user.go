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
	RegisterUser(data *request.UserRequest) (error, response.UserResponse)
	LoginUser(data *request.UserRequest) (error, response.UserResponse)
	GetAllUser() (error, []response.UserResponse)
	GetUser(id string) (error, response.UserResponse)
	UpdateUser(id string, input request.UserRequest) (error, response.UserResponse)
	DeleteUser(id string) (error, response.UserResponse)
}

type userRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) RegisterUser(data *request.UserRequest) (error, response.UserResponse) {
	dataUser := domain.ConvertFromUserReqToModel(*data)
	err := u.db.Create(&dataUser).Error
	if err != nil {
		return err, response.UserResponse{}
	}
	return nil, *domain.ConvertFromModelToUserRes(*dataUser)
}

func (u *userRepository) LoginUser(data *request.UserRequest) (error, response.UserResponse) {
	dataUser := domain.ConvertFromUserReqToModel(*data)
	err := u.db.Where("email = ? ", data.Email).First(&dataUser).Error
	if err != nil {
		return errors.ERR_EMAIL_NOT_FOUND, response.UserResponse{}
	}

	err = bcrypt.CheckPassword(data.Password, dataUser.Password)
	if err != nil {
		return errors.ERR_WRONG_PASSWORD, response.UserResponse{}
	}
	return nil, *domain.ConvertFromModelToUserRes(*dataUser)
}

func (u *userRepository) GetAllUser() (error, []response.UserResponse) {
	var allUser []models.Users
	var resAllUser []response.UserResponse
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

func (u *userRepository) GetUser(id string) (error, response.UserResponse) {
	var userData models.Users
	err := u.db.First(&userData, "id = ?", id).Error

	if err != nil {
		return err, response.UserResponse{}
	}
	return nil, *domain.ConvertFromModelToUserRes(userData)
}

func (u *userRepository) UpdateUser(id string, input request.UserRequest) (error, response.UserResponse) {
	userData := models.Users{}
	err := u.db.First(&userData, "id = ?", id).Error

	if err != nil {
		return err, response.UserResponse{}
	}

	if input.FirstName != "" {
		userData.FirstName = input.FirstName
	}
	
	if input.LastName != "" {
		userData.LastName = input.LastName
	}

	if input.Email != "" {
		userData.Email = input.Email
	}

	if input.Password != "" {
		userData.Password = input.Password
	}

	// if input.AddressID != "" {
	// 	userData.AddressID = input.AddressID
	// }

	// if input.BirthDate != "" {
	// 	userData.BirthDate = input.BirthDate
	// }

	if input.PhoneNumber != "" {
		userData.PhoneNumber = input.PhoneNumber
	}

	if err = u.db.Save(&userData).Error; err != nil {
		return err, response.UserResponse{}
	}
	return nil, *domain.ConvertFromModelToUserRes(userData)
}

func (u *userRepository) DeleteUser(id string) (error, response.UserResponse) {
	userData := models.Users{}
	res := response.UserResponse{}
	find := u.db.First(&userData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToUserRes(userData)
	}
	err := u.db.Delete(&userData, "id = ?", id).Error
	if err != nil {
		return err, response.UserResponse{}
	}
	return nil, res
}