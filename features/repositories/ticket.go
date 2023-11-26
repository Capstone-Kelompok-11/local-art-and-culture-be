package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type ITicketRepository interface {
	CreateTicket(data *request.Ticket) (response.Ticket, error)
	GetAllTicket(nameFilter string) ([]response.Ticket, error)
	GetTicket(id string) (response.Ticket, error)
	UpdateTicket(id string, input request.Ticket) (response.Ticket, error)
	DeleteTicket(id string) (response.Ticket, error)
}

type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *ticketRepository {
	return &ticketRepository{db}
}

func (ti *ticketRepository) CreateTicket(data *request.Ticket) (response.Ticket, error) {
	dataTicket:= domain.ConvertFromTicketReqToModel(*data)
	err := ti.db.Create(&dataTicket).Error
	if err != nil {
		return response.Ticket{}, err
	}
	err = ti.db.Preload("Event").First(&dataTicket, "id = ?", dataTicket.ID).Error
	return *domain.ConvertFromModelToTicketRes(*dataTicket), nil
}

func (ti *ticketRepository) GetAllTicket(nameFilter string) ([]response.Ticket, error) {
	var allTicket []models.Ticket
	var resAllTicket []response.Ticket

	query := ti.db.Preload("Event")
	if nameFilter != "" {
		query = query.Where("name LIKE ?", "%"+nameFilter+"%")
	}

	err := query.Find(&allTicket).Error
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}

	for i := 0; i < len(allTicket); i++ {
		ticketVm := domain.ConvertFromModelToTicketRes(allTicket[i])
		resAllTicket = append(resAllTicket, *ticketVm)
	}
	return resAllTicket, nil
}

func (ti *ticketRepository) GetTicket(id string) (response.Ticket, error) {
	var dataTicket models.Ticket
	err := ti.db.Preload("Event").First(&dataTicket, "id = ?", dataTicket.ID).Error
	if err != nil {
		return response.Ticket{}, err
	}
	return *domain.ConvertFromModelToTicketRes(dataTicket), nil
}

func (ti *ticketRepository) UpdateTicket(id string, input request.Ticket) (response.Ticket, error) {
	ticketData := models.Ticket{}
	err := ti.db.First(&ticketData, "id = ?", id).Error
	if err != nil {
		return response.Ticket{}, err
	}

	if input.Type != "" {
		ticketData.Type = input.Type
	}
	if input.Price != 0 {
		ticketData.Price = input.Price
	}
	
	if err = ti.db.Save(&ticketData).Error; err != nil {
		return response.Ticket{}, err
	}
	return *domain.ConvertFromModelToTicketRes(ticketData),nil
}

func (ti *ticketRepository) DeleteTicket(id string) (response.Ticket, error) {
	ticketData := models.Ticket{}
	res := response.Ticket{}
	find := ti.db.Preload("Event").First(&ticketData).Error
	if find == nil {
		res = *domain.ConvertFromModelToTicketRes(ticketData)
	}

	err := ti.db.Delete(&ticketData, "id = ?", id).Error
	if err != nil {
		return response.Ticket{}, err
	}
	return res, nil
}