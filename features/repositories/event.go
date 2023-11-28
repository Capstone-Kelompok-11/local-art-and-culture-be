package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/errors"
	"time"

	"gorm.io/gorm"
)

type IEventRepository interface {
	CreateEvent(data *request.Event) (response.Event, error)
	GetAllEvent(nameFilter, startDate, endDate string, page, pageSize int) ([]response.Event, int, error)
	GetEvent(id string) (response.Event, error)
	UpdateEvent(id string, input request.Event) (response.Event, error)
	DeleteEvent(id string) (response.Event, error)
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *eventRepository {
	return &eventRepository{db}
}

func (er *eventRepository) CreateEvent(data *request.Event) (response.Event, error) {
	dataEvent := domain.ConvertFromEventReqToModel(*data)
	err := er.db.Create(&dataEvent).Error
	if err != nil {
		return response.Event{}, err
	}
	err = er.db.Preload("Category").Preload("Creator").First(&dataEvent, "id = ?", dataEvent.ID).Error
	return *domain.ConvertFromModelToEventRes(*dataEvent), nil
}

func (er *eventRepository) GetAllEvent(nameFilter, startDate, endDate string, page, pageSize int) ([]response.Event, int, error) {
    var allEvent []models.Event
    var resAllEvent []response.Event

    query := er.db.Preload("Category").Preload("Creator")
    if nameFilter != "" {
        query = query.Where("event_name LIKE ?", "%"+nameFilter+"%")
    }

    if startDate != "" && endDate != "" {
        startDateTime, err := time.Parse("2006-01-02", startDate)
        if err != nil {
            return nil, 0, err
        }

        endDateTime, err := time.Parse("2006-01-02", endDate)
        if err != nil {
            return nil, 0, err
        }

        query = query.Where("from_date BETWEEN ? AND ?", startDateTime, endDateTime)
    }

	offset := (page - 1) * pageSize

	query = query.Limit(pageSize).Offset(offset)

    err := query.Find(&allEvent).Error
    if err != nil {
        return nil, 0, err
    }

    for i := 0; i < len(allEvent); i++ {
        eventVm := domain.ConvertFromModelToEventRes(allEvent[i])
        resAllEvent = append(resAllEvent, *eventVm)
    }

	var allItems int64
	query.Count(&allItems)

    return resAllEvent, int(allItems), nil
}

func (er *eventRepository) GetEvent(id string) (response.Event, error) {
	var eventData models.Event
	err := er.db.Preload("Category").Preload("Creator").First(&eventData, "id = ?", id).Error

	if err != nil {
		return response.Event{}, err
	}

	return *domain.ConvertFromModelToEventRes(eventData), nil
}

func (er *eventRepository) UpdateEvent(id string, input request.Event) (response.Event, error) {
	eventData := models.Event{}
	err := er.db.Preload("Category").Preload("Creator").First(&eventData, "id = ?", id).Error

	if err != nil {
		return response.Event{}, errors.ERR_GET_EVENT_BAD_REQUEST_ID
	}
	if input.EventName != "" {
		eventData.EventName = input.EventName
	}
	if input.EventDescription != "" {
		eventData.EventDescription = input.EventDescription
	}
	if !input.FromDate.IsZero() {
		eventData.FromDate = input.FromDate
	}
	if !input.ToDate.IsZero() {
		eventData.ToDate = input.ToDate
	}

	if err = er.db.Save(&eventData).Error; err != nil {
		return response.Event{}, errors.ERR_UPDATE_DATA
	}
	return *domain.ConvertFromModelToEventRes(eventData), nil
}

func (er *eventRepository) DeleteEvent(id string) (response.Event, error) {
	eventData := models.Event{}
	res := response.Event{}
	find := er.db.Preload("Category").Preload("Creator").First(&eventData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToEventRes(eventData)
	}

	err := er.db.Delete(&eventData, "id = ?", id).Error
	if err != nil {
		return response.Event{}, err
	}
	return res, nil
}