package repositories

import (
	"log"
	"lokasani/entity/domain"
	"lokasani/entity/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(domain.UsersDomain) (domain.UsersDomain, error)
	Login(email string, password string) (domain.UsersDomain, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUsersRepository(DB *gorm.DB) *userRepository {
	return &userRepository{db: DB}
}

func (u *userRepository) Create(user domain.UsersDomain) (domain.UsersDomain, error) {
	insert := domain.FromDomainToUserModel(user)
	err := u.db.Create(&insert).Error
	if err != nil {
		return user, err
	}
	data := domain.FromModelToUser(insert)
	return data, nil
}

func (u *userRepository) Login(email string, password string) (domain.UsersDomain, error) {
	var user models.Users
	var data domain.UsersDomain

	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return data, err
	}

	data = domain.FromModelToUser(user)
	log.Println(data)
	return data, nil
}