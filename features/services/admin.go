package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/bcrypt"
	"lokasani/helpers/errors"
	"lokasani/helpers/middleware"
	"math"
)

type IAdminService interface {
	RegisterAdmin(data *request.SuperAdmin) (response.SuperAdmin, error)
	LoginAdmin(data *request.SuperAdmin) (response.SuperAdmin, error)
	GetAllAdmin(nameFilter string, page, pageSize int) ([]response.SuperAdmin, int, error)
	GetAdmin(id string) (response.SuperAdmin, error)
	UpdateAdmin(id string, input request.SuperAdmin) (response.SuperAdmin, error)
	DeleteAdmin(id string) (response.SuperAdmin, error)
	CalculatePaginationValues(page, pageSize, allItmes int) (int, int)
	GetNextPage(currentPage, allPages int) int
	GetPrevPage(currentPage int) int
}

type AdminService struct {
	adminRepository repositories.IAdminRepository
}

func NewAdminService(repo repositories.IAdminRepository) *AdminService {
	return &AdminService{adminRepository: repo}
}

func (as *AdminService) RegisterAdmin(data *request.SuperAdmin) (response.SuperAdmin, error) {
	if data.Name == "" {
		return response.SuperAdmin{}, errors.ERR_NAME_IS_EMPTY
	}
	if data.Email == "" {
		return response.SuperAdmin{}, errors.ERR_EMAIL_IS_EMPTY
	}
	if data.Password == "" {
		return response.SuperAdmin{}, errors.ERR_PASSWORD_IS_EMPTY
	}
	if data.PhoneNumber == "" {
		return response.SuperAdmin{}, errors.ERR_PHONE_NUMBER_IS_EMPTY
	}
	hashPass, err := bcrypt.HashPassword(data.Password)
	if err != nil {
		return response.SuperAdmin{}, errors.ERR_BCRYPT_PASSWORD
	}

	data.Password = hashPass
	res, err := as.adminRepository.RegisterAdmin(data)
	if err != nil {
		return response.SuperAdmin{}, errors.ERR_REGISTER_USER_DATABASE
	}
	token, err := middleware.CreateToken(uint(data.Id), data.Name, data.Role)

	if err != nil {
		return response.SuperAdmin{}, errors.ERR_TOKEN
	}
	res.Token = token
	return res, nil
}

func (as *AdminService) LoginAdmin(data *request.SuperAdmin) (response.SuperAdmin, error) {
	if data.Email == "" {
		return response.SuperAdmin{}, errors.ERR_EMAIL_IS_EMPTY
	} else if data.Password == "" {
		return response.SuperAdmin{}, errors.ERR_PASSWORD_IS_EMPTY
	}

	res, err := as.adminRepository.LoginAdmin(data)
	if err != nil {
		return response.SuperAdmin{}, err
	}

	data.Role = "superadmin"
	token, err := middleware.CreateToken(uint(data.Id), data.Name, data.Role)

	if err != nil {
		return response.SuperAdmin{}, errors.ERR_TOKEN
	}
	res.Token = token
	return res, nil
}

func (as *AdminService) GetAllAdmin(nameFilter string, page, pageSize int) ([]response.SuperAdmin, int, error) {
	res, allItems, err := as.adminRepository.GetAllAdmin(nameFilter, page, pageSize)
	if err != nil {
		return nil, 0, errors.ERR_GET_DATA
	}
	return res, allItems, nil
}

func (as *AdminService) GetAdmin(id string) (response.SuperAdmin, error) {
	if id == "" {
		return response.SuperAdmin{}, errors.ERR_GET_ADMIN_BAD_REQUEST_ID
	}
	res, err := as.adminRepository.GetAdmin(id)
	if err != nil {
		return response.SuperAdmin{}, err
	}
	return res, nil
}

func (as *AdminService) UpdateAdmin(id string, data request.SuperAdmin) (response.SuperAdmin, error) {
	if id == "" {
		return response.SuperAdmin{}, errors.ERR_GET_ADMIN_BAD_REQUEST_ID
	}
	res, err := as.adminRepository.UpdateAdmin(id, data)
	if err != nil {
		return response.SuperAdmin{}, errors.ERR_UPDATE_DATA
	}
	return res, nil
}

func (as *AdminService) DeleteAdmin(id string) (response.SuperAdmin, error) {
	if id == "" {
		return response.SuperAdmin{}, errors.ERR_GET_ADMIN_BAD_REQUEST_ID
	}
	res, err := as.adminRepository.DeleteAdmin(id)

	if err != nil {
		return response.SuperAdmin{}, errors.ERR_DELETE_ADMIN
	}

	return res, nil
}

func (pr *AdminService) CalculatePaginationValues(page, pageSize, allItmes int) (int, int) {
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

func (pr *AdminService) GetNextPage(currentPage, allPages int) int {
	if currentPage < allPages {
		return currentPage + 1
	}
	return allPages
}

func (pr *AdminService) GetPrevPage(currentPage int) int {
	if currentPage > 1 {
		return currentPage - 1
	}
	return 1
}