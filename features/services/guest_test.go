package services

import (
	"errors"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/mocks"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateGuest_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGuestRepository := mocks.NewMockIGuestRepository(ctrl)

	mockRequestGuest := &request.Guest{
		Name:   "test",
		Role: "test",
		EventId:  1,
	}

	expectedResponseGuest := response.Guest{
		Id:      1,
		Name:   "test",
		Role: "test",
		EventId:  1,
	}

	mockGuestRepository.EXPECT().CreateGuest(mockRequestGuest).Return(expectedResponseGuest, nil)

	GuestService := NewGuestService(mockGuestRepository)
	result, err := GuestService.CreateGuest(mockRequestGuest)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if result.Id != expectedResponseGuest.Id || result.Name != expectedResponseGuest.Name || result.Role != expectedResponseGuest.Role || result.EventId != expectedResponseGuest.EventId {
		t.Errorf("Expected response %v, but got %v", expectedResponseGuest, result)
	}
}

func TestCreateGuest_Failure_NameIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockGuestRepository := mocks.NewMockIGuestRepository(ctrl)

    mockGuestRepository.EXPECT().CreateGuest(gomock.Any()).Times(0)

    mockRequestGuest := &request.Guest{
        Name:   "",
        Role:   "Participant",
        EventId: 1,
    }

    expectedError := errors.New("name is empty")

    guestService := NewGuestService(mockGuestRepository)

    resultGuest, err := guestService.CreateGuest(mockRequestGuest)

    assert.Error(t, err)
    assert.Equal(t, expectedError, err)
    assert.Equal(t, response.Guest{}, resultGuest)
}

func TestCreateGuest_Failure_RoleIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockGuestRepository := mocks.NewMockIGuestRepository(ctrl)

    mockGuestRepository.EXPECT().CreateGuest(gomock.Any()).Times(0)

    mockRequestGuest := &request.Guest{
        Name:   "test",
        Role:   "",
        EventId: 1,
    }

    expectedError := errors.New("role is empty")

    guestService := NewGuestService(mockGuestRepository)

    resultGuest, err := guestService.CreateGuest(mockRequestGuest)

    assert.Error(t, err)
    assert.Equal(t, expectedError, err)
    assert.Equal(t, response.Guest{}, resultGuest)
}

func TestCreateGuest_Failure_DatabaseError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockGuestRepository := mocks.NewMockIGuestRepository(ctrl)

    mockRequestGuest := &request.Guest{
		Name:   "test",
        Role:   "test",
        EventId: 1,
    }

    expectedError := errors.New("failed to create new Guest")

    mockGuestRepository.EXPECT().CreateGuest(mockRequestGuest).Return(response.Guest{}, expectedError)

    GuestService := NewGuestService(mockGuestRepository)

    result, _ := GuestService.CreateGuest(mockRequestGuest)

    if result.Id != 0 || result.Name != "" || result.Role != "" || result.EventId != 0 {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestGetAllGuest_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockGuestRepository := mocks.NewMockIGuestRepository(ctrl)

    mockGuestRepository.EXPECT().GetAllGuest("eventID").Return([]response.Event{{Id: 1, EventName: "test", EventDescription: "test"}}, nil).Times(1)

    guestService := NewGuestService(mockGuestRepository)

    resultGuests, err := guestService.GetAllGuest("eventID")

    assert.NoError(t, err)
    assert.Equal(t, []response.Event{{Id: 1, EventName: "test", EventDescription: "test"}}, resultGuests)
}

func TestGetAllGuest_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockGuestRepository := mocks.NewMockIGuestRepository(ctrl)

    mockGuestRepository.EXPECT().GetAllGuest("eventID").Return(nil, errors.New("failed to get guests")).Times(1)

    guestService := NewGuestService(mockGuestRepository)

    resultGuests, err := guestService.GetAllGuest("eventID")

    assert.Error(t, err)
    assert.Equal(t, []response.Event(nil), resultGuests)
}

func TestGetGuest_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGuestRepository := mocks.NewMockIGuestRepository(ctrl)

	expectedGuest := response.Event{Id: 1, EventName: "test", EventDescription: "test"}

	mockGuestRepository.EXPECT().GetGuest("GuestID").Return(expectedGuest, nil).Times(1)

	GuestService := NewGuestService(mockGuestRepository)

	resultGuest, err := GuestService.GetGuest("GuestID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultGuest, expectedGuest) {
		t.Errorf("Expected Guest %+v, but got %+v", expectedGuest, resultGuest)
	}
}

func TestGetGuest_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGuestRepository := mocks.NewMockIGuestRepository(ctrl)

	expectedError := errors.New("failed to get Guest")

	mockGuestRepository.EXPECT().GetGuest("GuestID").Return(response.Event{}, expectedError).Times(1)

	GuestService := NewGuestService(mockGuestRepository)

	_, err := GuestService.GetGuest("GuestID")

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error containing '%v', but got '%v'", expectedError.Error(), err)
	}
}

func TestGetGuest_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGuestRepository := mocks.NewMockIGuestRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	GuestService := NewGuestService(mockGuestRepository)

	_, err := GuestService.GetGuest("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestDeleteGuest_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockGuestRepository := mocks.NewMockIGuestRepository(ctrl)

    expectedResponse := response.Guest{
		Id:      1,
		Name: "test",
		Role: "test",
		EventId: 1,
    }

    mockGuestRepository.EXPECT().DeleteGuest(gomock.Any()).Return(expectedResponse, nil)

    GuestService := NewGuestService(mockGuestRepository)

    _, err := GuestService.DeleteGuest("GuestID")

    if err != nil {
        t.Errorf("Expected no error, but got %v", err)
    }
}

func TestDeleteGuest_Failure_ErrDelete(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockGuestRepository := mocks.NewMockIGuestRepository(ctrl)

    expectedError := errors.New("failed to delete data from database")

    mockGuestRepository.EXPECT().DeleteGuest(gomock.Any()).Return(response.Guest{}, errors.New("failed to delete data from database"))

    GuestService := NewGuestService(mockGuestRepository)

    result, err := GuestService.DeleteGuest("GuestID")

    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error %v, but got %v", expectedError, err)
    }

    if result.Id != 0 || result.Name != "" || result.Role != "" || result.EventId != 0 {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeleteGuest_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGuestRepository := mocks.NewMockIGuestRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	GuestService := NewGuestService(mockGuestRepository)

	_, err := GuestService.DeleteGuest("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestUpdateGuest_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGuestRepository := mocks.NewMockIGuestRepository(ctrl)

	expectedUpdatedGuest := response.Guest{
		Id:      1,
		Name: "test",
		Role: "test",
		EventId: 1,
	}

	mockGuestRepository.EXPECT().UpdateGuest("GuestID", gomock.Any()).Return(expectedUpdatedGuest, nil).Times(1)

	GuestService := NewGuestService(mockGuestRepository)

	requestGuest := request.Guest{
		Name: "test",
		Role: "test",
		EventId: 1,
	}

	resultGuest, err := GuestService.UpdateGuest("GuestID", requestGuest)

	if err != nil {
    	t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultGuest, expectedUpdatedGuest) {
    	t.Errorf("Expected updated Guest %+v, but got %+v", expectedUpdatedGuest, resultGuest)
	}
}

func TestUpdateGuest_Failure_BadRequestID(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockGuestRepository := mocks.NewMockIGuestRepository(ctrl)

    expectedError := errors.New("can't find any data with this id")

    GuestService := NewGuestService(mockGuestRepository)

    _, err := GuestService.UpdateGuest("", request.Guest{})

    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }
}

func TestUpdateGuest_Failure_UpdateData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGuestRepository := mocks.NewMockIGuestRepository(ctrl)

	expectedError := errors.New("failed to update data")

	mockGuestRepository.EXPECT().UpdateGuest("GuestID", request.Guest{}).Return(response.Guest{}, expectedError).Times(1)

	GuestService := NewGuestService(mockGuestRepository)

	_, err := GuestService.UpdateGuest("GuestID", request.Guest{})

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error containing '%v', but got '%v'", expectedError.Error(), err)
	}
}