package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteAdmin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepo := mocks.NewMockIAdminRepository(ctrl)

	adminService := NewAdminService(mockAdminRepo)

	mockAdminRepo.EXPECT().DeleteAdmin(gomock.Any()).Return(response.Admin{}, nil)

	result, err := adminService.DeleteAdmin("pudu")

	assert.NoError(t, err)
	assert.Equal(t, response.Admin{}, result)
}

func TestDeleteAdmin_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepo := mocks.NewMockIAdminRepository(ctrl)

	adminService := NewAdminService(mockAdminRepo)

	mockAdminRepo.EXPECT().DeleteAdmin(gomock.Any()).Return(response.Admin{}, someError)

	result, err := adminService.DeleteAdmin("pudu")

	assert.Error(t, err)
	assert.Equal(t, response.Admin{}, result)
}

func TestGetAdmin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepo := mocks.NewMockIAdminRepository(ctrl)

	adminService := NewAdminService(mockAdminRepo)

	mockAdminRepo.EXPECT().GetAdmin(gomock.Any()).Return(response.Admin{}, nil)

	result, err := adminService.GetAdmin("pudu")

	assert.NoError(t, err)
	assert.Equal(t, response.Admin{}, result)
}

func TestGetAdmin_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepo := mocks.NewMockIAdminRepository(ctrl)

	adminService := NewAdminService(mockAdminRepo)

	mockAdminRepo.EXPECT().GetAdmin(gomock.Any()).Return(response.Admin{}, someError)

	result, err := adminService.GetAdmin("pudu")

	assert.Error(t, err)
	assert.Equal(t, response.Admin{}, result)
}	

func TestGetAllAdmin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepo := mocks.NewMockIAdminRepository(ctrl)

	adminService := NewAdminService(mockAdminRepo)

	mockAdminRepo.EXPECT().GetAllAdmin().Return([]response.Admin{
		{
			Id: 1, 
			Name: "pudu",
			Email: "pudu@gmail.com",
			Token: "123",
			PhoneNumber: "08564327811",
		}}, nil)

	result, err := adminService.GetAllAdmin()

	assert.NotNil(t, err, result)
}

func TestGetAllAdmin_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepo := mocks.NewMockIAdminRepository(ctrl)

	adminService := NewAdminService(mockAdminRepo)

	mockAdminRepo.EXPECT().GetAllAdmin().Return(nil, someError)

	result, err := adminService.GetAllAdmin()

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestLoginAdmin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepo := mocks.NewMockIAdminRepository(ctrl)

	adminService := NewAdminService(mockAdminRepo)

	mockAdminRepo.EXPECT().LoginAdmin(gomock.Any()).Return(response.Admin{}, nil)

	result, err := adminService.LoginAdmin(&request.Admin{
		Id:          1,
		Name:   "pudu",
		Email:       "pudu@gmail.com",
		Password:    "123",
		PhoneNumber: "08564321789",
	})

	assert.NoError(t, err, result)
}

func TestLoginAdmin_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepo := mocks.NewMockIAdminRepository(ctrl)

	adminService := NewAdminService(mockAdminRepo)
	mockAdminRepo.EXPECT().LoginAdmin(gomock.Any()).Return(response.Admin{}, someError)

	result, err := adminService.LoginAdmin(&request.Admin{
		Id:          1,
		Name:   	"pudu",
		Email:       "pudu@gmail.com",
		Password:    "123",
		PhoneNumber: "08564321789",
	})

	assert.Error(t, err)
	assert.Equal(t, response.Admin{}, result)
}

func TestRegisterAdmin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepo := mocks.NewMockIAdminRepository(ctrl)

	adminService := NewAdminService(mockAdminRepo)

	mockAdminRepo.EXPECT().RegisterAdmin(gomock.Any()).Return(response.Admin{}, nil)

	result, err := adminService.RegisterAdmin(&request.Admin{
		Id:          1,
		Name:   "pudu",
		Email:       "pudu@gmail.com",
		Password:    "123",
		PhoneNumber: "08564321789",
	})

	assert.NoError(t, err, result)
}

func TestRegisterAdmin_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepo := mocks.NewMockIAdminRepository(ctrl)

	adminService := NewAdminService(mockAdminRepo)

	mockAdminRepo.EXPECT().RegisterAdmin(gomock.Any()).Return(response.Admin{}, someError)

	result, err := adminService.RegisterAdmin(&request.Admin{
		Id:          1,
		Name:   "pudu",
		Email:       "pudu@gmail.com",
		Password:    "123",
		PhoneNumber: "08564321789",
	})

	assert.Error(t, err)
	assert.Equal(t, response.Admin{}, result)
}

func TestUpdateAdmin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepo := mocks.NewMockIAdminRepository(ctrl)

	adminService := NewAdminService(mockAdminRepo)

	mockAdminRepo.EXPECT().UpdateAdmin(gomock.Any(), gomock.Any()).Return(response.Admin{}, nil)

	result, err := adminService.UpdateAdmin("pudu", request.Admin{})

	assert.NoError(t, err)
	assert.Equal(t, response.Admin{}, result)
}

func TestUpdateAdmin_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepo := mocks.NewMockIAdminRepository(ctrl)

	adminService := NewAdminService(mockAdminRepo)

	mockAdminRepo.EXPECT().UpdateAdmin(gomock.Any(), gomock.Any()).Return(response.Admin{}, someError)

	result, err := adminService.UpdateAdmin("pudu", request.Admin{})

	assert.Error(t, err)
	assert.Equal(t, response.Admin{}, result)
}