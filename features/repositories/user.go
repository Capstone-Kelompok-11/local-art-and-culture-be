package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(data *request.UserRequest) (response.UserResponse, error)
	//LoginUser(email string, password string) (*request.UserRequest, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) CreateUser(data *request.UserRequest) (response.UserResponse, error) {
	dataUser := domain.ConvertFromUserReqToModel(*data)
	err := u.db.Create(&dataUser).Error
	if err != nil {
		return response.UserResponse{}, err
	}
	return *domain.ConvertFromModelToAdminRes(*dataUser), nil
}

// func (u *userRepository) LoginUser(email string, password string) (*request.UserRequest, error) {
// 	var user models.Users
// 	var data domain.UsersDomain

// 	err := u.db.Where("email = ?", email).First(&user).Error
// 	if err != nil {
// 		return data, err
// 	}

// 	data = domain.FromModelToUser(user)
// 	log.Println(data)
// 	return data, nil
// }