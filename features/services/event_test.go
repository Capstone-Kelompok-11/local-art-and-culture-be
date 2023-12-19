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

func TestCreateEvent_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventRepository := mocks.NewMockIEventRepository(ctrl)

	mockRequestEvent := &request.Event{
		FromDate: time.Now(),
		ToDate: time.Now(),
		EventName: "test",
		EventDescription: "test",
		CreatorId: 1,
		CategoryId: 1,
		Creator: request.Creator{},
		Category: request.Category{},
		Guest: []request.Guest{},
	}

	expectedResponseEvent := response.Event{
		Id:          1,
		FromDate: time.Now(),
		ToDate: time.Now(),
		EventName: "test",
		EventDescription: "test",
		CreatorId: 1,
		CategoryId: 1,
		Creator: response.Creator{},
		Category: response.Category{},
		Guest: []response.Guest{},
	}
	mockEventRepository.EXPECT().CreateEvent(mockRequestEvent).Return(expectedResponseEvent, nil)

	EventService := NewEventService(mockEventRepository)
	result, err := EventService.CreateEvent(mockRequestEvent)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if result.Id != expectedResponseEvent.Id || result.EventName != expectedResponseEvent.EventName || result.EventDescription!= expectedResponseEvent.EventDescription || result.FromDate != expectedResponseEvent.FromDate || result.ToDate!= expectedResponseEvent.ToDate {
		t.Errorf("Expected response %v, but got %v", expectedResponseEvent, result)
	}
}

func TestCreateEvent_Failure_NameIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepository := mocks.NewMockIEventRepository(ctrl)

    mockRequestEvent := &request.Event{
		FromDate: time.Now(),
		ToDate: time.Now(),
		EventName: "",
		EventDescription: "test",
		CreatorId: 1,
		CategoryId: 1,
		Creator: request.Creator{},
		Category: request.Category{},
		Guest: []request.Guest{},
    }

    expectedError := errors.New("name is empty")

    EventService := NewEventService(mockEventRepository)

    result, err := EventService.CreateEvent(mockRequestEvent)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.EventName != "" || result.EventDescription != "" || result.CategoryId != 0 || result.CreatorId != 0 {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestCreateEvent_Failure_DescIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepository := mocks.NewMockIEventRepository(ctrl)

    mockRequestEvent := &request.Event{
        FromDate: time.Now(),
		ToDate: time.Now(),
		EventName: "test",
		EventDescription: "",
		CreatorId: 1,
		CategoryId: 1,
		Creator: request.Creator{},
		Category: request.Category{},
		Guest: []request.Guest{},
    }

    expectedError := errors.New("description is empty")

    EventService := NewEventService(mockEventRepository)

    result, err := EventService.CreateEvent(mockRequestEvent)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

	if result.Id != 0 || result.EventName != "" || result.EventDescription != "" || result.CategoryId != 0 || result.CreatorId != 0 {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestCreateEvent_Failure_DateIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepository := mocks.NewMockIEventRepository(ctrl)

    mockRequestEvent := &request.Event{
        FromDate: time.Time{},
		ToDate: time.Time{},
		EventName: "test",
		EventDescription: "test",
		CreatorId: 1,
		CategoryId: 1,
		Creator: request.Creator{},
		Category: request.Category{},
		Guest: []request.Guest{},
    }

    expectedError := errors.New("event date is empty")

    EventService := NewEventService(mockEventRepository)

    result, err := EventService.CreateEvent(mockRequestEvent)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

	if result.Id != 0 || result.EventName != "" || result.EventDescription != "" || result.CategoryId != 0 || result.CreatorId != 0 {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestCreateEvent_Failure_DatabaseError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepository := mocks.NewMockIEventRepository(ctrl)

    mockRequestEvent := &request.Event{
		FromDate: time.Now(),
		ToDate: time.Now(),
		EventName: "test",
		EventDescription: "test",
		CreatorId: 1,
		CategoryId: 1,
		Creator: request.Creator{},
		Category: request.Category{},
		Guest: []request.Guest{},
    }

    expectedError := errors.New("failed to create new Event")

    mockEventRepository.EXPECT().CreateEvent(mockRequestEvent).Return(response.Event{}, expectedError)

    EventService := NewEventService(mockEventRepository)

    result, _ := EventService.CreateEvent(mockRequestEvent)

    if result.Id != 0 || result.EventName != "" || result.EventDescription != "" || result.CategoryId != 0 || result.CreatorId != 0 {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestGetAllEvent_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepository := mocks.NewMockIEventRepository(ctrl)

    expectedEvents := []response.Event{
        {Id: 1, FromDate: time.Now(), ToDate: time.Now(), EventName: "test1", EventDescription: "test1", CreatorId: 1, CategoryId: 1},
        {Id: 2, FromDate: time.Now(), ToDate: time.Now(), EventName: "test2", EventDescription: "test2", CreatorId: 2, CategoryId: 2},
    }

    mockEventRepository.EXPECT().
        GetAllEvent("filter", "startDate", "endDate", gomock.Any(), gomock.Any()).
        Return(expectedEvents, 0, nil).
        Times(1)

    EventService := NewEventService(mockEventRepository)

    resultEvents, _, err := EventService.GetAllEvent("filter", "startDate", "endDate", 1, 10)

    assert.NoError(t, err)
    assert.Equal(t, expectedEvents, resultEvents)
}

func TestGetAllEvent_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventRepository := mocks.NewMockIEventRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")
	mockEventRepository.EXPECT().GetAllEvent("filter", "startDate", "endDate", gomock.Any(), gomock.Any()).Return(nil, 0, expectedError).Times(1)

	EventService := NewEventService(mockEventRepository)

	resultCategories, _, err := EventService.GetAllEvent("filter", "startDate", "endDate", 1, 10)

	assert.Error(t, err)
	assert.Nil(t, resultCategories)
	assert.EqualError(t, err, expectedError.Error())
}

func TestGetAllAvailableEvent_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepository := mocks.NewMockIEventRepository(ctrl)

    expectedEvents := []response.Event{
        {Id: 1, FromDate: time.Now(), ToDate: time.Now(), EventName: "test1", EventDescription: "test1", CreatorId: 1, CategoryId: 1},
        {Id: 2, FromDate: time.Now(), ToDate: time.Now(), EventName: "test2", EventDescription: "test2", CreatorId: 2, CategoryId: 2},
    }

    mockEventRepository.EXPECT().
        GetAllAvailableEvent("filter", "startDate", "endDate", gomock.Any(), gomock.Any()).
        Return(expectedEvents, 0, nil).
        Times(1)

    EventService := NewEventService(mockEventRepository)

    resultEvents, _, err := EventService.GetAllAvailableEvent("filter", "startDate", "endDate", 1, 10)

    assert.NoError(t, err)
    assert.Equal(t, expectedEvents, resultEvents)
}

func TestGetAllAvailableEvent_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventRepository := mocks.NewMockIEventRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")
	mockEventRepository.EXPECT().GetAllAvailableEvent("filter", "startDate", "endDate", gomock.Any(), gomock.Any()).Return(nil, 0, expectedError).Times(1)

	EventService := NewEventService(mockEventRepository)

	resultCategories, _, err := EventService.GetAllAvailableEvent("filter", "startDate", "endDate", 1, 10)

	assert.Error(t, err)
	assert.Nil(t, resultCategories)
	assert.EqualError(t, err, expectedError.Error())
}

func TestGetEvent_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventRepository := mocks.NewMockIEventRepository(ctrl)

	expectedEvent := response.Event{
		Id:     1,
		FromDate: time.Now(),
		ToDate: time.Now(),
		EventName: "test",
		EventDescription: "test",
		CreatorId: 1,
		CategoryId: 1,
		Creator: response.Creator{},
		Category: response.Category{},
		Guest: []response.Guest{},
	}

	mockEventRepository.EXPECT().GetEvent("EventID").Return(expectedEvent, nil).Times(1)

	EventService := NewEventService(mockEventRepository)

	resultEvent, err := EventService.GetEvent("EventID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultEvent, expectedEvent) {
		t.Errorf("Expected Event %+v, but got %+v", expectedEvent, resultEvent)
	}
}

func TestGetEvent_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventRepository := mocks.NewMockIEventRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockEventRepository.EXPECT().GetEvent("EventID").Return(response.Event{}, expectedError).Times(1)

	EventService := NewEventService(mockEventRepository)

	result, err := EventService.GetEvent("EventID")

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if result.Id != 0 || result.EventName != "" || result.EventDescription != "" || result.CategoryId != 0 || result.CreatorId != 0 {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestGetEvent_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventRepository := mocks.NewMockIEventRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	EventService := NewEventService(mockEventRepository)

	_, err := EventService.GetEvent("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestDeleteEvent_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventRepository := mocks.NewMockIEventRepository(ctrl)

	expectedEvent := response.Event{
		Id:     1,
		FromDate: time.Now(),
		ToDate: time.Now(),
		EventName: "test",
		EventDescription: "test",
		CreatorId: 1,
		CategoryId: 1,
		Creator: response.Creator{},
		Category: response.Category{},
		Guest: []response.Guest{},
	}

	mockEventRepository.EXPECT().DeleteEvent("EventID").Return(expectedEvent, nil).Times(1)

	EventService := NewEventService(mockEventRepository)

	resultEvent, err := EventService.DeleteEvent("EventID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultEvent, expectedEvent) {
		t.Errorf("Expected Event %+v, but got %+v", expectedEvent, resultEvent)
	}
}

func TestDeleteEvent_Failure_ErrDelete(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepository := mocks.NewMockIEventRepository(ctrl)

    expectedError := errors.New("failed to delete data from database")

    mockEventRepository.EXPECT().DeleteEvent(gomock.Any()).Return(response.Event{}, errors.New("failed to delete data from database"))

    EventService := NewEventService(mockEventRepository)

    result, err := EventService.DeleteEvent("EventID")

    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error %v, but got %v", expectedError, err)
    }

    if result.Id != 0 || result.EventName != "" || result.EventDescription != "" || result.CategoryId != 0 || result.CreatorId != 0 {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeleteEvent_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventRepository := mocks.NewMockIEventRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	EventService := NewEventService(mockEventRepository)

	_, err := EventService.DeleteEvent("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestUpdateEvent_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventRepository := mocks.NewMockIEventRepository(ctrl)

	expectedUpdatedEvent := response.Event{
		Id:          1,
		ToDate: time.Now(),
		EventName: "test",
		EventDescription: "test",
		CreatorId: 1,
		CategoryId: 1,
		Creator: response.Creator{},
		Category: response.Category{},
		Guest: []response.Guest{},
	}

	var expectedError error

	mockEventRepository.EXPECT().UpdateEvent("EventID", gomock.Any()).Return(expectedUpdatedEvent, expectedError).Times(1)

	EventService := NewEventService(mockEventRepository)

	requestEvent := request.Event{
		ToDate: time.Now(),
		EventName: "test",
		EventDescription: "test",
		CreatorId: 1,
		CategoryId: 1,
		Creator: request.Creator{},
		Category: request.Category{},
		Guest: []request.Guest{},
	}

	resultEvent, err := EventService.UpdateEvent("EventID", requestEvent)

	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if !reflect.DeepEqual(resultEvent, expectedUpdatedEvent) {
		t.Errorf("Expected %+v, but got %+v", expectedUpdatedEvent, resultEvent)
	}
}

func TestUpdateEvent_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventRepository := mocks.NewMockIEventRepository(ctrl)

	expectedError := errors.New("failed to update data")

	mockEventRepository.EXPECT().UpdateEvent("EventID", gomock.Any()).Return(response.Event{}, expectedError).Times(1)

	EventService := NewEventService(mockEventRepository)

	requestEvent := request.Event{
		FromDate: time.Now(),
		ToDate: time.Now(),
		EventName: "test",
		EventDescription: "test",
		CreatorId: 1,
		CategoryId: 1,
		Creator: request.Creator{},
		Category: request.Category{},
		Guest: []request.Guest{},
	}

	result, err := EventService.UpdateEvent("EventID", requestEvent)

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if result.Id != 0 || result.EventName != "" || result.EventDescription != "" || result.CategoryId != 0 || result.CreatorId != 0 {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestUpdateEvent_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventRepository := mocks.NewMockIEventRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	EventService := NewEventService(mockEventRepository)

	_, err := EventService.UpdateEvent("", request.Event{})

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestCalculatePaginationValues_Event(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepository := mocks.NewMockIEventRepository(ctrl)

    EventService := NewEventService(mockEventRepository)

    page1, totalPages1 := EventService.CalculatePaginationValues(0, 100, 8)
    assert.Equal(t, 1, page1)
    assert.Equal(t, 1, totalPages1)

    page2, totalPages2 := EventService.CalculatePaginationValues(15, 100, 8)
    assert.Equal(t, 1, page2)
    assert.Equal(t, 1, totalPages2)

    page3, totalPages3 := EventService.CalculatePaginationValues(2, 100, 8)
    assert.Equal(t, 1, page3)
    assert.Equal(t, 1, totalPages3)

    page4, totalPages4 := EventService.CalculatePaginationValues(1, 95, 8)
    assert.Equal(t, 1, page4)
    assert.Equal(t, 1, totalPages4) 
}

func TestGetNextPage_Event(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepository := mocks.NewMockIEventRepository(ctrl)

    EventService := NewEventService(mockEventRepository)

    result1 := EventService.GetNextPage(3, 10)
	assert.Equal(t, 4, result1)

	result2 := EventService.GetNextPage(10, 10)
	assert.Equal(t, 10, result2)
}

func TestGetPrevPage_Event(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepository := mocks.NewMockIEventRepository(ctrl)

    EventService := NewEventService(mockEventRepository)

    result1 := EventService.GetPrevPage(5)
	assert.Equal(t, 4, result1)

	result2 := EventService.GetPrevPage(1)
	assert.Equal(t, 1, result2)
}