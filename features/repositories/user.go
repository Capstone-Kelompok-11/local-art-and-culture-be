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
	RegisterUser(data *request.User) (response.User, error)
	LoginUser(data *request.User) (response.Creators, error)
	GetAllUser(nameFilter string, page, pageSize int) ([]response.User, int, error)
	CountUsersByRole(roleId uint) (int, error)
	GetUser(id string) (response.User, error)
	//getRoleName(roleID uint) string
	UpdateUser(id string, input request.User) (response.User, error)
	DeleteUser(id string) (response.User, error)
	FindByEmail(email string) (*models.Users, error)
	CreateUser(user *models.Users) (*models.Users, error)
	//SaveOTP(otp *models.OTP) (*models.OTP, error)
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
	err = u.db.Preload("Role").First(&dataUser, "id = ?", dataUser.ID).Error
	if err != nil {
		return response.User{}, errors.ERR_LOGIN
	}
	return *domain.ConvertFromModelToUserRes(*dataUser), nil
}

func (u *userRepository) LoginUser(data *request.User) (response.Creators, error) {
	dataUser := domain.ConvertFromUserReqToModel(*data)
	err := u.db.Where("email = ? ", data.Email).First(&dataUser).Error
	if err != nil {
		return response.Creators{}, errors.ERR_EMAIL_NOT_FOUND
	}

	err = bcrypt.CheckPassword(data.Password, dataUser.Password)
	if err != nil {
		return response.Creators{}, errors.ERR_WRONG_PASSWORD
	}
	var creator models.Creator
	err = u.db.Preload("Role").Preload("Users").First(&creator, "user_id = ?", dataUser.ID).Error
	if err != nil {
		creator.Users = *dataUser
		return *domain.ConvertFromModelToCreatorsRes(creator), nil
	}
	return *domain.ConvertFromModelToCreatorsRes(creator), nil
}

func (u *userRepository) GetAllUser(nameFilter string, page, pageSize int) ([]response.User, int, error) {
	var allUser []models.Users
	var resAllUser []response.User

	query := u.db.Preload("Role")
	if nameFilter != "" {
    	query = query.Where("first_name LIKE ? OR last_name LIKE ?", "%"+nameFilter+"%", "%"+nameFilter+"%")
	}

	offset := (page - 1) * pageSize

	query = query.Limit(pageSize).Offset(offset)

	err := query.Find(&allUser).Error
	if err != nil {
		return nil, 0, errors.ERR_GET_DATA
	}

	for i := 0; i < len(allUser); i++ {
		userVm := domain.ConvertFromModelToUserRes(allUser[i])
		resAllUser = append(resAllUser, *userVm)
	}

	var allItems int64
	query.Count(&allItems)

	return resAllUser, int(allItems), nil
}

func (u *userRepository) CountUsersByRole(roleId uint) (int, error) {
	var count int64
	var query = u.db.Model(&models.Users{})

	query.Where("role_id = ?", roleId)
	
	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (u *userRepository) GetUser(id string) (response.User, error) {
	var userData models.Users
	err := u.db.Preload("Role").First(&userData, "id = ?", id).Error
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
	if input.Username != "" {
		userData.Username = input.Username
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
	// if input.RoleId != 0 {
	// 	userData.RoleId = input.RoleId
	// }

	if err = u.db.Save(&userData).Error; err != nil {
		return response.User{}, err
	}
	return *domain.ConvertFromModelToUserRes(userData), nil
}

func (u *userRepository) DeleteUser(id string) (response.User, error) {
	userData := models.Users{}
	res := response.User{}

	find := u.db.Preload("Role").First(&userData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToUserRes(userData)
	}

	err := u.db.Delete(&models.Creator{}, "user_id = ?", id).Error
	if err != nil {
		return response.User{}, err
	}

	err = u.db.Delete(&userData, "id = ?", id).Error
	if err != nil {
		return response.User{}, err
	}

	return res, nil
}


func (u *userRepository) FindByEmail(email string) (*models.Users, error) {
	user := models.Users{}
	res := u.db.Where("email = ?", email).First(&user).Error
	if res != nil {
		return nil, res
	}
	return &user, nil
}

func (u *userRepository) CreateUser(user *models.Users) (*models.Users, error) {
	result := u.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
