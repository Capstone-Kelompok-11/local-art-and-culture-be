package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/bcrypt"
	"lokasani/helpers/errors"
	"lokasani/helpers/middleware"
)

type IAdminService interface {
	RegisterAdmin(data *request.Admin) (error, response.Admin)
	LoginAdmin(data *request.Admin) (error, response.Admin)
	GetAllAdmin() (error, []response.Admin)
	GetAdmin(id string) (error, response.Admin)
	UpdateAdmin(id string, input request.Admin) (error, response.Admin)
	DeleteAdmin(id string) (error, response.Admin)
}

type AdminService struct {
	adminRepository repositories.IAdminRepository
}

func NewAdminService(repo repositories.IAdminRepository) *AdminService {
	return &AdminService{adminRepository: repo}
}

func (as *AdminService) RegisterAdmin(data *request.Admin) (error, response.Admin) {
	if data.Name == "" {
		return errors.ERR_NAME_IS_EMPTY, response.Admin{}
	}
	if data.Email == "" {
		return errors.ERR_EMAIL_IS_EMPTY, response.Admin{}
	}
	if data.Password == "" {
		return errors.ERR_PASSWORD_IS_EMPTY, response.Admin{}
	}
	if data.PhoneNumber == "" {
		return errors.ERR_PHONE_NUMBER_IS_EMPTY, response.Admin{}
	}
	hashPass, err := bcrypt.Hash(data.Password)
	if err != nil {
		return errors.ERR_BCRYPT_PASSWORD, response.Admin{}
	}

	data.Password = hashPass
	err, res := as.adminRepository.RegisterAdmin(data)
	if err != nil {
		return errors.ERR_REGISTER_USER_DATABASE, response.Admin{}
	}
	token, err := middleware.CreateToken(int(data.Id), data.Name)

	if err != nil {
		return errors.ERR_TOKEN, response.Admin{}
	}
	res.Token = token
	return nil, res
}

func (as *AdminService) LoginAdmin(data *request.Admin) (error, response.Admin) {
	if data.Email == "" {
		return errors.ERR_EMAIL_IS_EMPTY, response.Admin{}
	} else if data.Password == "" {
		return errors.ERR_PASSWORD_IS_EMPTY, response.Admin{}
	}

	err, res := as.adminRepository.LoginAdmin(data)
	if err != nil {
		return err, response.Admin{}
	}

	token, err := middleware.CreateToken(int(data.Id), data.Name)

	if err != nil {
		return errors.ERR_TOKEN, response.Admin{}
	}
	res.Token = token
	return nil, res
}

func (as *AdminService) GetAllAdmin() (error, []response.Admin) {
	err, res := as.adminRepository.GetAllAdmin()
	if err != nil {
		return errors.ERR_GET_DATA, nil
	}
	return nil, res
}

func (as *AdminService) GetAdmin(id string) (error, response.Admin) {
	if id == "" {
		return errors.ERR_GET_ADMIN_BAD_REQUEST_ID, response.Admin{}
	}
	err, res := as.adminRepository.GetAdmin(id)
	if err != nil {
		return err, response.Admin{}
	}
	return nil, res
}

func (as *AdminService) UpdateAdmin(id string, data request.Admin) (error, response.Admin) {
	if id == "" {
		return errors.ERR_GET_ADMIN_BAD_REQUEST_ID, response.Admin{}
	}
	err, res := as.adminRepository.UpdateAdmin(id, data)
	if err != nil {
		return errors.ERR_UPDATE_DATA, response.Admin{}
	}
	return nil, res
}

func (as *AdminService) DeleteAdmin(id string) (error, response.Admin) {
	if id == "" {
		return errors.ERR_GET_ADMIN_BAD_REQUEST_ID, response.Admin{}
	}
	err, res := as.adminRepository.DeleteAdmin(id)

	if err != nil {
		return errors.ERR_DELETE_ADMIN, response.Admin{}
	}

	return nil, res
}
