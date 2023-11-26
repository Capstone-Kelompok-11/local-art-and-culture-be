package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type ITicketService interface {
	CreateTicket(data *request.Ticket) (response.Ticket, error)
	GetAllTicket(nameFilter string) ([]response.Ticket, error)
	GetTicket(id string) (response.Ticket, error)
	UpdateTicket(id string, input request.Ticket) (response.Ticket, error)
	DeleteTicket(id string) (response.Ticket, error)
}

type TicketService struct {
	ticketRepository repositories.ITicketRepository
}

func NewTicketService(repo repositories.ITicketRepository) *TicketService {
	return &TicketService{ticketRepository: repo}
}

func (ti *TicketService) CreateTicket(data *request.Ticket) (response.Ticket, error) {
	if data.Name != "" {
		return response.Ticket{}, errors.ERR_NAME_IS_EMPTY
	}
	if data.Email != "" {
		return response.Ticket{}, errors.ERR_EMAIL_IS_EMPTY
	}
	if data.Type != "" {
		return response.Ticket{}, errors.ERR_TYPE_IS_EMPTY
	}
	if data.Price != 0 {
		return response.Ticket{}, errors.ERR_PRICE_IS_EMPTY
	}
	
	res, err := ti.ticketRepository.CreateTicket(data)
	if err != nil {
		return response.Ticket{}, errors.ERR_CREATE_TICKET_DATABASE
	}
	return res, nil
}

func (ti *TicketService) GetAllTicket(nameFilter string) ([]response.Ticket, error) {
	res, err := ti.ticketRepository.GetAllTicket(nameFilter)
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return res, nil
}

func (ti *TicketService) GetTicket(id string) (response.Ticket, error) {
	if id == "" {
		return response.Ticket{}, errors.ERR_GET_TICKET_BAD_REQUEST_ID
	}
	res, err := ti.ticketRepository.GetTicket(id)
	if err != nil {
		return response.Ticket{}, err
	}
	return res, nil
}

func (ti *TicketService) UpdateTicket(id string, input request.Ticket) (response.Ticket, error) {
	if id == "" {
		return response.Ticket{}, errors.ERR_GET_TICKET_BAD_REQUEST_ID
	}
	res, err := ti.ticketRepository.UpdateTicket(id, input)
	if err != nil {
		return response.Ticket{}, err
	}
	return res, nil
}

func (ti *TicketService)DeleteTicket(id string) (response.Ticket, error) {
	if id == "" {
		return response.Ticket{}, errors.ERR_GET_TICKET_BAD_REQUEST_ID
	}
	res, err := ti.ticketRepository.DeleteTicket(id)
	if err != nil {
		return response.Ticket{}, err
	}
	return res, nil
}