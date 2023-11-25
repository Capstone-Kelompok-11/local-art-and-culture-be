package repositories

import (
	"log"
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/bcrypt"
	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type IUserRepository interface {
	RegisterUser(data *request.User) (response.User, error)
	LoginUser(data *request.User) (response.User, error)
	GetAllUser(nameFilter string) ([]response.User, error)
	GetUser(id string) (response.User, error)
	UpdateUser(id string, input request.User) (response.User, error)
	DeleteUser(id string) (response.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) RegisterUser(data *request.User) (response.User, error) {
	dataUser := domain.ConvertFromUserReqToModel(*data)
	err := u.db.Create(&dataUser).Error
	if err != nil {
		return response.User{}, err
	}
	return *domain.ConvertFromModelToUserRes(*dataUser), nil
}

func (u *userRepository) LoginUser(data *request.User) (response.User, error) {
	dataUser := domain.ConvertFromUserReqToModel(*data)
	err := u.db.Where("email = ? ", data.Email).First(&dataUser).Error
	if err != nil {
		return response.User{}, errors.ERR_EMAIL_NOT_FOUND
	}

	err = bcrypt.CheckPassword(data.Password, dataUser.Password)
	if err != nil {
		return response.User{}, errors.ERR_WRONG_PASSWORD
	}
	return *domain.ConvertFromModelToUserRes(*dataUser), nil
}

func (u *userRepository) GetAllUser(nameFilter string) ([]response.User, error) {
	var allUser []models.Users
	var resAllUser []response.User

	query := u.db.Model(&models.Users{})
	if nameFilter != "" {
    	query = query.Where("first_name LIKE ? OR last_name LIKE ?", "%"+nameFilter+"%", "%"+nameFilter+"%")
	}

	err := query.Find(&allUser).Error
	if err != nil {
		log.Printf("Error fetching data from database: %v", err)
		return nil, errors.ERR_GET_DATA
	}

	for i := 0; i < len(allUser); i++ {
		userVm := domain.ConvertFromModelToUserRes(allUser[i])
		resAllUser = append(resAllUser, *userVm)
	}

	return resAllUser, nil
}

func (u *userRepository) GetUser(id string) (response.User, error) {
	var userData models.Users
	err := u.db.First(&userData, "id = ?", id).Error
	if err != nil {
		return response.User{}, err
	}
	return *domain.ConvertFromModelToUserRes(userData), nil
}

func (u *userRepository) UpdateUser(id string, input request.User) (response.User, error) {
	userData := models.Users{}
	err := u.db.First(&userData, "id = ?", id).Error

	if err != nil {
		return response.User{}, err
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
	if input.PhoneNumber != "" {
		userData.PhoneNumber = input.PhoneNumber
	}
	if input.NIK != "" {
		userData.NIK = input.NIK
	}
	if input.Gender != "" {
		userData.Gender = input.Gender
	}

	if err = u.db.Save(&userData).Error; err != nil {
		return response.User{}, err
	}
	return *domain.ConvertFromModelToUserRes(userData), nil
}

func (u *userRepository) DeleteUser(id string) (response.User, error) {
	userData := models.Users{}
	res := response.User{}
	find := u.db.First(&userData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToUserRes(userData)
	}
	err := u.db.Delete(&userData, "id = ?", id).Error
	if err != nil {
		return response.User{}, err
	}
	return res, nil
}