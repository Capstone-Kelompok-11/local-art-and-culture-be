package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateCreator_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreatorRepo := mocks.NewMockICreatorRepository(ctrl)

	creatorService := NewCreatorService(mockCreatorRepo)

	mockCreatorRepo.EXPECT().CreateCreator(gomock.Any()).Return(response.Creator{}, nil)

	result, err := creatorService.CreateCreator(&request.Creator{
		Id:          1,
		OutletName:  "example",
		Email:       "example@gmail.com",
		PhoneNumber: "087654321",
		UserId:      1,
		RoleId:      1,
	})

	assert.NoError(t, err)
	assert.Equal(t, response.Creator{}, result)
}

func TestCreateCreator_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreatorRepo := mocks.NewMockICreatorRepository(ctrl)

	creatorService := NewCreatorService(mockCreatorRepo)

	mockCreatorRepo.EXPECT().CreateCreator(gomock.Any()).Return(response.Creator{}, someError)

	result, err := creatorService.CreateCreator(&request.Creator{
		Id:          1,
		OutletName:  "example",
		Email:       "example@gmail.com",
		PhoneNumber: "087654321",
		UserId:      1,
		RoleId:      1,
	})

	assert.Error(t, err)
	assert.Equal(t, response.Creator{}, result)
}

func TestDeleteCreator_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreatorRepo := mocks.NewMockICreatorRepository(ctrl)

	creatorService := NewCreatorService(mockCreatorRepo)

	mockCreatorRepo.EXPECT().DeleteCreator(gomock.Any()).Return(response.Creator{}, nil)

	result, err := creatorService.DeleteCreator("example")

	assert.NoError(t, err)
	assert.Equal(t, response.Creator{}, result)
}

func TestDeleteCreator_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreatorRepo := mocks.NewMockICreatorRepository(ctrl)

	creatorService := NewCreatorService(mockCreatorRepo)

	mockCreatorRepo.EXPECT().DeleteCreator(gomock.Any()).Return(response.Creator{}, someError)

	result, err := creatorService.DeleteCreator("example")

	assert.Error(t, err)
	assert.Equal(t, response.Creator{}, result)
}

func TestGetAllCreator_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreatorRepo := mocks.NewMockICreatorRepository(ctrl)

	creatorService := NewCreatorService(mockCreatorRepo)

	mockCreatorRepo.EXPECT().GetAllCreator().Return([]response.Creator{
		{
			Id:          1,
			OutletName:  "pudu",
			RoleId:      1,
			UserId:      1,
			Email:       "pudu@gmail.com",
			PhoneNumber: "08564321789",
		}}, nil)

	result, err := creatorService.GetAllCreator()

	assert.NoError(t, err)
	assert.Equal(t, []response.Creator{
		{
			Id:          1,
			OutletName:  "pudu",
			RoleId:      1,
			UserId:      1,
			Email:       "pudu@gmail.com",
			PhoneNumber: "08564321789",
		}}, result)
}

func TestGetAllCreator_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreatorRepo := mocks.NewMockICreatorRepository(ctrl)

	creatorService := NewCreatorService(mockCreatorRepo)

	mockCreatorRepo.EXPECT().GetAllCreator().Return(nil, someError)

	result, err := creatorService.GetAllCreator()

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetCreator_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreatorRepo := mocks.NewMockICreatorRepository(ctrl)

	creatorService := NewCreatorService(mockCreatorRepo)

	mockCreatorRepo.EXPECT().GetCreator(gomock.Any()).Return(response.Creator{}, nil)

	result, err := creatorService.GetCreator("pudu")

	assert.NoError(t, err)
	assert.Equal(t, response.Creator{}, result)
}

func TestGetCreator_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreatorRepo := mocks.NewMockICreatorRepository(ctrl)

	creatorService := NewCreatorService(mockCreatorRepo)

	mockCreatorRepo.EXPECT().GetCreator(gomock.Any()).Return(response.Creator{}, someError)

	result, err := creatorService.GetCreator("pudu")

	assert.Error(t, err)
	assert.Equal(t, response.Creator{}, result)
}

func TestUpdateCreator_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreatorRepo := mocks.NewMockICreatorRepository(ctrl)

	creatorService := NewCreatorService(mockCreatorRepo)

	mockCreatorRepo.EXPECT().UpdateCreator(gomock.Any(), gomock.Any()).Return(response.Creator{}, nil)

	result, err := creatorService.UpdateCreator("pudu", request.Creator{})

	assert.NoError(t, err)
	assert.Equal(t, response.Creator{}, result)
}

func TestUpdateCreator_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCreatorRepo := mocks.NewMockICreatorRepository(ctrl)

	creatorService := NewCreatorService(mockCreatorRepo)

	mockCreatorRepo.EXPECT().UpdateCreator(gomock.Any(), gomock.Any()).Return(response.Creator{}, someError)

	result, err := creatorService.UpdateCreator("pudu", request.Creator{})

	assert.Error(t, err)
	assert.Equal(t, response.Creator{}, result)
}
