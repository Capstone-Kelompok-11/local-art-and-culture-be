package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/bcrypt"

	"gorm.io/gorm"
)

type IUserRepository interface {
	RegisterUser(data *request.UserRequest) (error, response.UserResponse)
	LoginUser(data *request.UserRequest) (error, response.UserResponse)
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
	return nil, *domain.ConvertFromModelToAdminRes(*dataUser)
}

func (u *userRepository) LoginUser(data *request.UserRequest) (error, response.UserResponse) {
	dataUser := domain.ConvertFromUserReqToModel(*data)
	err := u.db.Where("email = ? ", data.Email).First(&dataUser).Error
	if err != nil {
		return err, response.UserResponse{}
	}

	err = bcrypt.CheckPassword(data.Password, dataUser.Password)
	if err != nil {
		return err, response.UserResponse{}
	}
	return nil, *domain.ConvertFromModelToAdminRes(*dataUser)
}