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

func TestCreateFiles_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFilesRepository := mocks.NewMockIFilesRepository(ctrl)

	expectedFiles := response.Files{
		Id:   1,
		SourceId: 1,
		SourceStr: "test",
		Image: "test",
	}

	mockFilesRepository.EXPECT().CreateFiles(gomock.Any()).Return(expectedFiles, nil).Times(1)

	FilesService := NewFilesService(mockFilesRepository)

	requestFiles := &request.Files{
		SourceId: 1,
		SourceStr: "test",
		Image: "test",
	}

	resultFiles, err := FilesService.CreateFiles(requestFiles)

	assert.NoError(t, err)
	assert.Equal(t, expectedFiles, resultFiles)
}

func TestCreateFiles_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFilesRepository := mocks.NewMockIFilesRepository(ctrl)

	expectedError := errors.New("failed to create Files")

	mockFilesRepository.EXPECT().CreateFiles(gomock.Any()).Return(response.Files{}, expectedError).Times(1)

	FilesService := NewFilesService(mockFilesRepository)

	requestFiles := &request.Files{
		SourceId: 1,
		SourceStr: "test",
		Image: "test",
	}

	resultFiles, err := FilesService.CreateFiles(requestFiles)

	assert.Error(t, err)
	assert.Equal(t, response.Files{}, resultFiles)
}

func TestCreateFiles_Failure_ImageEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

	mockFilesRepository := mocks.NewMockIFilesRepository(ctrl)

    mockRequestFiles := &request.Files{
		SourceId: 1,
		SourceStr: "test",
		Image: "",
    }

    expectedError := errors.New("image is empty")

    FilesService := NewFilesService(mockFilesRepository)

    result, err := FilesService.CreateFiles(mockRequestFiles)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.Image != "" || result.SourceStr != "" || result.SourceId != 0{
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestCreateFiles_Failure_SrcEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

	mockFilesRepository := mocks.NewMockIFilesRepository(ctrl)

    mockRequestFiles := &request.Files{
		SourceId: 0,
		SourceStr: "",
		Image: "test",
    }

    expectedError := errors.New("source is empty`")

    FilesService := NewFilesService(mockFilesRepository)

    result, err := FilesService.CreateFiles(mockRequestFiles)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.Image != "" || result.SourceStr != "" || result.SourceId != 0{
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestGetAllFiles_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFilesRepository := mocks.NewMockIFilesRepository(ctrl)

	expectedFiless := []response.Files{
		{Id: 1, SourceId: 1, SourceStr: "test1", Image: "test1"},
		{Id: 2, SourceId: 1, SourceStr: "test1", Image: "test1"},
	}

	var expectedError error

	mockFilesRepository.EXPECT().GetAllFiles("filter").Return(expectedFiless, expectedError).Times(1)

	FilesService := NewFilesService(mockFilesRepository)

	resultFiless, err := FilesService.GetAllFiles("filter")

	if err != expectedError {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultFiless, expectedFiless) {
		t.Errorf("Expected Filess %+v, but got %+v", expectedFiless, resultFiless)
	}
}

func TestGetAllFiles_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFilesRepository := mocks.NewMockIFilesRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockFilesRepository.EXPECT().GetAllFiles("filter").Return(nil, expectedError).Times(1)

	FilesService := NewFilesService(mockFilesRepository)

	resultFiless, err := FilesService.GetAllFiles("filter")

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if resultFiless != nil {
		t.Errorf("Expected nil Filess, but got %+v", resultFiless)
	}
}

func TestGetFiles_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFilesRepository := mocks.NewMockIFilesRepository(ctrl)

	expectedFiles := response.Files{
		Id:     1,
		SourceId: 1,
		SourceStr: "test",
		Image: "test",
	}

	mockFilesRepository.EXPECT().GetFiles("FilesID").Return(expectedFiles, nil).Times(1)

	FilesService := NewFilesService(mockFilesRepository)

	resultFiles, err := FilesService.GetFiles("FilesID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultFiles, expectedFiles) {
		t.Errorf("Expected Files %+v, but got %+v", expectedFiles, resultFiles)
	}
}

func TestGetFiles_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFilesRepository := mocks.NewMockIFilesRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockFilesRepository.EXPECT().GetFiles("FilesID").Return(response.Files{}, expectedError).Times(1)

	FilesService := NewFilesService(mockFilesRepository)

	resultFiles, err := FilesService.GetFiles("FilesID")

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if resultFiles.Id != 0 || resultFiles.SourceStr != "" || resultFiles.Image != "" || resultFiles.SourceId != 0 {
		t.Errorf("Expected empty result, but got %+v", resultFiles)
	}
}

func TestGetFiles_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFilesRepository := mocks.NewMockIFilesRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	FilesService := NewFilesService(mockFilesRepository)

	_, err := FilesService.GetFiles("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestUpdateFiles_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFilesRepository := mocks.NewMockIFilesRepository(ctrl)

	expectedUpdatedFiles := response.Files{
		Id:          1,
		SourceId: 1,
		SourceStr: "test",
		Image: "test",
	}

	var expectedError error

	mockFilesRepository.EXPECT().UpdateFiles("FilesID", gomock.Any()).Return(expectedUpdatedFiles, expectedError).Times(1)

	FilesService := NewFilesService(mockFilesRepository)

	requestFiles := request.Files{
		SourceId: 1,
		SourceStr: "test",
		Image: "test",
	}

	resultFiles, err := FilesService.UpdateFiles("FilesID", requestFiles)

	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if !reflect.DeepEqual(resultFiles, expectedUpdatedFiles) {
		t.Errorf("Expected %+v, but got %+v", expectedUpdatedFiles, resultFiles)
	}
}

func TestUpdateFiles_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFilesRepository := mocks.NewMockIFilesRepository(ctrl)

	expectedError := errors.New("failed to update data")

	mockFilesRepository.EXPECT().UpdateFiles("FilesID", gomock.Any()).Return(response.Files{}, expectedError).Times(1)

	FilesService := NewFilesService(mockFilesRepository)

	requestFiles := request.Files{
		SourceId: 1,
		SourceStr: "test",
		Image: "test",
	}

	resultFiles, err := FilesService.UpdateFiles("FilesID", requestFiles)

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if resultFiles.Id != 0 || resultFiles.SourceStr != "" || resultFiles.Image != "" || resultFiles.SourceId != 0 {
		t.Errorf("Expected empty result, but got %+v", resultFiles)
	}
}

func TestUpdateFiles_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFilesRepository := mocks.NewMockIFilesRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	FilesService := NewFilesService(mockFilesRepository)

	_, err := FilesService.UpdateFiles("", request.Files{})

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestDeleteFiles_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFilesRepository := mocks.NewMockIFilesRepository(ctrl)

	expectedFiles := response.Files{
		Id:     1,
		SourceId: 1,
		SourceStr: "test",
		Image: "test",
	}

	mockFilesRepository.EXPECT().DeleteFiles("FilesID").Return(expectedFiles, nil).Times(1)

	FilesService := NewFilesService(mockFilesRepository)

	resultFiles, err := FilesService.DeleteFiles("FilesID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultFiles, expectedFiles) {
		t.Errorf("Expected Files %+v, but got %+v", expectedFiles, resultFiles)
	}
}

func TestDeleteFiles_Failure_ErrDelete(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockFilesRepository := mocks.NewMockIFilesRepository(ctrl)

    expectedError := errors.New("failed to delete data from database")

    mockFilesRepository.EXPECT().DeleteFiles(gomock.Any()).Return(response.Files{}, errors.New("failed to delete data from database"))

    FilesService := NewFilesService(mockFilesRepository)

    result, err := FilesService.DeleteFiles("FilesID")

    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error %v, but got %v", expectedError, err)
    }

	if result.Id != 0 || result.Image != "" || result.SourceStr != "" || result.SourceId != 0{
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeleteFiles_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFilesRepository := mocks.NewMockIFilesRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	FilesService := NewFilesService(mockFilesRepository)

	_, err := FilesService.DeleteFiles("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}