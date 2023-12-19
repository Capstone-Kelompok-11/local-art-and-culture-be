package services

import (
	"errors"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/mocks"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateTicket_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockTicketRepository := mocks.NewMockITicketRepository(ctrl)

    mockRequestTicket := &request.Ticket{
        Type:        "test",
		Price:       10.000,
		Qty:         1,
		Name:        "test",
		Description: "test",
		StartTime:   time.Now(),
		EndTime:     time.Now(),
		EventId:     1,
    }

    expectedResponseTicket := response.Ticket{
		Id: 1,
		Type:        "test",
		Price:       10.000,
		Qty:         1,
		Name:        "test",
		Description: "test",
		StartTime:   time.Now(),
		EndTime:     time.Now(),
		EventId:     1,
		Event: response.Events{
			Id:              1,
			FromDate:        time.Now(),
			ToDate:          time.Now(),
			EventName:       "test",
			EventDescription: "test",
			CreatorId:       1, 
			CategoryId: 1,
		},
	}
    mockTicketRepository.EXPECT().CreateTicket(mockRequestTicket).Return(expectedResponseTicket, nil)

    TicketService := NewTicketService(mockTicketRepository)
    result, err := TicketService.CreateTicket(mockRequestTicket)

    if err != nil {
        t.Errorf("Expected no error, but got: %v", err)
    }

    if result.Id != expectedResponseTicket.Id || result.Type != expectedResponseTicket.Type || result.Price != expectedResponseTicket.Price || result.Qty != expectedResponseTicket.Qty || result.Name != expectedResponseTicket.Name{
        t.Errorf("Expected response %v, but got %v", expectedResponseTicket, result)
    }
}

func TestCreateTicket_Failure_TypeIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockTicketRepository := mocks.NewMockITicketRepository(ctrl)

    mockRequestTicket := &request.Ticket{
        Type:        "",
		Price:       10.000,
		Qty:         1,
		Name:        "test",
		Description: "test",
		StartTime:   time.Now(),
		EndTime:     time.Now(),
		EventId:     1,
    }

    expectedError := errors.New("type is empty")

    TicketService := NewTicketService(mockTicketRepository)

    result, err := TicketService.CreateTicket(mockRequestTicket)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.Type != "" || result.Price != 0 || result.Qty != 0 || result.Name != "" || result.Description != "" || result.EventId != 0{
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestCreateTicket_Failure_PriceIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockTicketRepository := mocks.NewMockITicketRepository(ctrl)

    mockRequestTicket := &request.Ticket{
        Type:        "test",
		Price:       0,
		Qty:         1,
		Name:        "test",
		Description: "test",
		StartTime:   time.Now(),
		EndTime:     time.Now(),
		EventId:     1,
    }

    expectedError := errors.New("price is empty")

    TicketService := NewTicketService(mockTicketRepository)

    result, err := TicketService.CreateTicket(mockRequestTicket)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.Type != "" || result.Price != 0 || result.Qty != 0 || result.Name != "" || result.Description != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestCreateTicket_Failure_DateIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockTicketRepository := mocks.NewMockITicketRepository(ctrl)

    mockRequestTicket := &request.Ticket{
        Type:        "test",
		Price:       10.000,
		Qty:         1,
		Name:        "test",
		Description: "test",
		StartTime:   time.Time{},
		EndTime:     time.Time{},
		EventId:     1,
    }

    expectedError := errors.New("event date is empty")

    TicketService := NewTicketService(mockTicketRepository)

    result, err := TicketService.CreateTicket(mockRequestTicket)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.Type != "" || result.Price != 0 || result.Qty != 0 || result.Name != "" || result.Description != "" || result.EventId != 0{
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestCreateTicket_Failure_QtyIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockTicketRepository := mocks.NewMockITicketRepository(ctrl)

    mockRequestTicket := &request.Ticket{
        Type:        "test",
		Price:       10.000,
		Qty:         0,
		Name:        "test",
		Description: "test",
		StartTime:   time.Now(),
		EndTime:     time.Now(),
		EventId:     1,
    }

    expectedError := errors.New("quantity is empty")

    TicketService := NewTicketService(mockTicketRepository)

    result, err := TicketService.CreateTicket(mockRequestTicket)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.Type != "" || result.Price != 0 || result.Qty != 0 || result.Name != "" || result.Description != "" || result.EventId != 0{
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestCreateTicket_Failure_NameIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockTicketRepository := mocks.NewMockITicketRepository(ctrl)

    mockRequestTicket := &request.Ticket{
        Type:        "test",
		Price:       10.000,
		Qty:         1,
		Name:        "",
		Description: "test",
		StartTime:   time.Now(),
		EndTime:     time.Now(),
		EventId:     1,
    }

    expectedError := errors.New("name is empty")

    TicketService := NewTicketService(mockTicketRepository)

    result, err := TicketService.CreateTicket(mockRequestTicket)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.Type != "" || result.Price != 0 || result.Qty != 0 || result.Name != "" || result.Description != "" || result.EventId != 0{
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestCreateTicket_Failure_DescIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockTicketRepository := mocks.NewMockITicketRepository(ctrl)

    mockRequestTicket := &request.Ticket{
        Type:        "test",
		Price:       10.000,
		Qty:         1,
		Name:        "test",
		Description: "",
		StartTime:   time.Now(),
		EndTime:     time.Now(),
		EventId:     1,
    }

    expectedError := errors.New("description is empty")

    TicketService := NewTicketService(mockTicketRepository)

    result, err := TicketService.CreateTicket(mockRequestTicket)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.Type != "" || result.Price != 0 || result.Qty != 0 || result.Name != "" || result.Description != "" || result.EventId != 0{
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestCreateTicket_Failure_DatabaseError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockTicketRepository := mocks.NewMockITicketRepository(ctrl)

    mockRequestTicket := &request.Ticket{
		Type:        "test",
		Price:       10.000,
		Qty:         1,
		Name:        "test",
		Description: "test",
		StartTime:   time.Now(),
		EndTime:     time.Now(),
		EventId:     1,
    }

    expectedError := errors.New("failed to create new Ticket")

    mockTicketRepository.EXPECT().CreateTicket(mockRequestTicket).Return(response.Ticket{}, expectedError)

    TicketService := NewTicketService(mockTicketRepository)

    result, _ := TicketService.CreateTicket(mockRequestTicket)

    if result.Id != 0 || result.Type != "" || result.Price != 0 || result.Qty != 0 || result.Name != "" || result.Description != "" || result.EventId != 0{
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestGetAllTicket_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTicketRepository := mocks.NewMockITicketRepository(ctrl)

	expectedCategories := []response.Ticket{
		{Id: 1, Type: "Type1", Price: 1000, Qty: 1, Name: "test1", Description: "test1", StartTime: time.Now(), EndTime: time.Now(), EventId: 1},
		{Id: 2, Type: "Type2", Price: 2000, Qty: 2, Name: "test2", Description: "test2", StartTime: time.Now(), EndTime: time.Now(), EventId: 2},
	}
	mockTicketRepository.EXPECT().GetAllTicket("filter").Return(expectedCategories, nil).Times(1)

	TicketService := NewTicketService(mockTicketRepository)

	resultCategories, err := TicketService.GetAllTicket("filter")

	assert.NoError(t, err)
	assert.Equal(t, expectedCategories, resultCategories)
}

func TestGetAllTicket_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTicketRepository := mocks.NewMockITicketRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")
	mockTicketRepository.EXPECT().GetAllTicket("filter").Return(nil, expectedError).Times(1)

	TicketService := NewTicketService(mockTicketRepository)

	resultCategories, err := TicketService.GetAllTicket("filter")

	assert.Error(t, err)
	assert.Nil(t, resultCategories)
	assert.EqualError(t, err, expectedError.Error())
}

func TestGetTicket_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTicketRepository := mocks.NewMockITicketRepository(ctrl)

	expectedTicket := response.Ticket{
		Type:        "test",
		Price:       10.000,
		Qty:         1,
		Name:        "test",
		Description: "test",
		StartTime:   time.Now(),
		EndTime:     time.Now(),
		EventId:     1,
	}

	mockTicketRepository.EXPECT().GetTicket("TicketID").Return(expectedTicket, nil).Times(1)

	TicketService := NewTicketService(mockTicketRepository)

	resultTicket, err := TicketService.GetTicket("TicketID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultTicket, expectedTicket) {
		t.Errorf("Expected Ticket %+v, but got %+v", expectedTicket, resultTicket)
	}
}

func TestGetTicket_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTicketRepository := mocks.NewMockITicketRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockTicketRepository.EXPECT().GetTicket("TicketID").Return(response.Ticket{}, expectedError).Times(1)

	TicketService := NewTicketService(mockTicketRepository)

	resultTicket, err := TicketService.GetTicket("TicketID")

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if resultTicket.Id != 0 || resultTicket.Name != "" {
		t.Errorf("Expected empty result, but got %+v", resultTicket)
	}
}

func TestGetTicket_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTicketRepository := mocks.NewMockITicketRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	TicketService := NewTicketService(mockTicketRepository)

	_, err := TicketService.GetTicket("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestUpdateTicket_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTicketRepository := mocks.NewMockITicketRepository(ctrl)

	expectedUpdatedTicket := response.Ticket{
		Id:          1,
		Type:        "test",
		Price:       10.000,
		Qty:         1,
		Name:        "test",
		Description: "test",
		StartTime:   time.Now(),
		EndTime:     time.Now(),
		EventId:     1,
	}

	var expectedError error

	mockTicketRepository.EXPECT().UpdateTicket("TicketID", gomock.Any()).Return(expectedUpdatedTicket, expectedError).Times(1)

	TicketService := NewTicketService(mockTicketRepository)

	requestTicket := request.Ticket{
		Type:        "test",
		Price:       10.000,
		Qty:         1,
		Name:        "test",
		Description: "test",
		StartTime:   time.Now(),
		EndTime:     time.Now(),
		EventId:     1,
	}

	resultTicket, err := TicketService.UpdateTicket("TicketID", requestTicket)

	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if !reflect.DeepEqual(resultTicket, expectedUpdatedTicket) {
		t.Errorf("Expected %+v, but got %+v", expectedUpdatedTicket, resultTicket)
	}
}

func TestUpdateTicket_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTicketRepository := mocks.NewMockITicketRepository(ctrl)

	expectedError := errors.New("failed to update data")

	mockTicketRepository.EXPECT().UpdateTicket("TicketID", gomock.Any()).Return(response.Ticket{}, expectedError).Times(1)

	TicketService := NewTicketService(mockTicketRepository)

	requestTicket := request.Ticket{
		Type:        "test",
		Price:       10.000,
		Qty:         1,
		Name:        "test",
		Description: "test",
		StartTime:   time.Now(),
		EndTime:     time.Now(),
		EventId:     1,
	}

	resultTicket, err := TicketService.UpdateTicket("TicketID", requestTicket)

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if resultTicket.Id != 0 || resultTicket.Name != ""{
		t.Errorf("Expected empty result, but got %+v", resultTicket)
	}
}

func TestUpdateTicket_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTicketRepository := mocks.NewMockITicketRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	TicketService := NewTicketService(mockTicketRepository)

	_, err := TicketService.UpdateTicket("", request.Ticket{})

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestDeleteTicket_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTicketRepository := mocks.NewMockITicketRepository(ctrl)

	expectedTicket := response.Ticket{
		Id:     1,
		Type:        "test",
		Price:       10.000,
		Qty:         1,
		Name:        "test",
		Description: "test",
		StartTime:   time.Now(),
		EndTime:     time.Now(),
		EventId:     1,
	}

	mockTicketRepository.EXPECT().DeleteTicket("TicketID").Return(expectedTicket, nil).Times(1)

	TicketService := NewTicketService(mockTicketRepository)

	resultTicket, err := TicketService.DeleteTicket("TicketID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultTicket, expectedTicket) {
		t.Errorf("Expected Ticket %+v, but got %+v", expectedTicket, resultTicket)
	}
}

func TestDeleteTicket_Failure_ErrDelete(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockTicketRepository := mocks.NewMockITicketRepository(ctrl)

    expectedError := errors.New("failed to delete data from database")

    mockTicketRepository.EXPECT().DeleteTicket(gomock.Any()).Return(response.Ticket{}, errors.New("failed to delete data from database"))

    TicketService := NewTicketService(mockTicketRepository)

    result, err := TicketService.DeleteTicket("TicketID")

    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error %v, but got %v", expectedError, err)
    }

    if result.Id != 0 || result.Name != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeleteTicket_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTicketRepository := mocks.NewMockITicketRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	TicketService := NewTicketService(mockTicketRepository)

	_, err := TicketService.DeleteTicket("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}