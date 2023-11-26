package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type IGuestRepository interface {
	CreateGuest(data *request.Guest) (response.Guest, error)
	GetAllGuest(nameFilter string) ([]response.Guest, error)
	GetGuest(id string) (response.Guest, error)
	UpdateGuest(id string, input request.Guest) (response.Guest, error)
	DeleteGuest(id string) (response.Guest, error)
}

type guestRepository struct {
	db *gorm.DB
}

func NewGuestRepository(db *gorm.DB) *guestRepository {
	return &guestRepository{db}
}

func (gu *guestRepository) CreateGuest(data *request.Guest) (response.Guest, error) {
	dataGuest := domain.ConvertFromGuestReqToModel(*data)
	err := gu.db.Create(&dataGuest).Error
	if err != nil {
		return response.Guest{}, err
	}
	err = gu.db.Preload("Event").First(&dataGuest, "id = ?", dataGuest.ID).Error
	return *domain.ConvertFromModelToGuestRes(*dataGuest), nil
}

func (gu *guestRepository) GetAllGuest(nameFilter string) ([]response.Guest, error) {
	var allGuest []models.Guest
	var resAllGuest []response.Guest

	query := gu.db.Preload("Event")

	if nameFilter != "" {
		query = query.Where("name LIKE ?", "%"+nameFilter+"%")
	}

	err := query.Find(&allGuest).Error
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}

	for i := 0; i < len(allGuest); i++ {
		GuestVm := domain.ConvertFromModelToGuestRes(allGuest[i])
		resAllGuest = append(resAllGuest, *GuestVm)
	}

	return resAllGuest, nil
}

func (gu *guestRepository) GetGuest(id string) (response.Guest, error) {
	var guestData models.Guest
	err := gu.db.Preload("Event").First(&guestData, "id = ?", id).Error

	if err != nil {
		return response.Guest{}, err
	}
	return *domain.ConvertFromModelToGuestRes(guestData), nil
}

func (gu *guestRepository) UpdateGuest(id string, input request.Guest) (response.Guest, error) {
	guestData := models.Guest{}
	err := gu.db.First(&guestData, "id = ?", id).Error
	if err != nil {
		return response.Guest{}, err
	}

	if input.Name != "" {
		guestData.Name = input.Name
	}
	if input.Role != "" {
		guestData.Role = input.Role
	}
	if input.EventId != 0 {
		guestData.EventId = input.EventId
	}

	if err = gu.db.Save(&guestData).Error; err != nil {
		return response.Guest{}, err
	}
	return *domain.ConvertFromModelToGuestRes(guestData), nil
}

func (gu *guestRepository) DeleteGuest(id string) (response.Guest, error) {
	guestData := models.Guest{}
	res := response.Guest{}
	find := gu.db.Preload("Event").First(&guestData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToGuestRes(guestData)
	}

	err := gu.db.Delete(&guestData, "id = ?", id).Error
	if err != nil {
		return response.Guest{}, err
	}
	return res, nil
}