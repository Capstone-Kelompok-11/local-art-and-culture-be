package services

// import (
// 	"errors"
// 	"lokasani/entity/request"
// 	"lokasani/entity/response"
// 	"lokasani/features/mocks"
// 	"testing"
// 	"time"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// )

// var someError = errors.New("some error")

// func TestDeleteUser_Success(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockUserRepo := mocks.NewMockIUserRepository(ctrl)

// 	userService := NewUserService(mockUserRepo)

// 	mockUserRepo.EXPECT().DeleteUser(gomock.Any()).Return(response.User{}, nil)

// 	result, err := userService.DeleteUser("user123")

// 	assert.NoError(t, err)
// 	assert.Equal(t, response.User{}, result)
// }

// func TestDeleteUser_Failure(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockUserRepo := mocks.NewMockIUserRepository(ctrl)

// 	userService := NewUserService(mockUserRepo)

// 	mockUserRepo.EXPECT().DeleteUser(gomock.Any()).Return(response.User{}, someError)

// 	result, err := userService.DeleteUser("user123")

// 	assert.Error(t, err)
// 	assert.Equal(t, response.User{}, result)
// }

// func TestGetAllUser_Success(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockUserRepo := mocks.NewMockIUserRepository(ctrl)

// 	userService := NewUserService(mockUserRepo)

// 	mockUserRepo.EXPECT().GetAllUser().Return([]response.User{
// 		{
// 			Id: 1, 
// 			FirstName: "pudu",
// 			LastName: "lee",
// 			Email: "pudu@gmail.com",
// 			PhoneNumber: "08564321789",
// 		}}, nil)

// 	users, err := userService.GetAllUser()

// 	assert.NoError(t, err)
// 	assert.Len(t, users, 1)
// 	assert.Equal(t, uint(1), users[0].Id)
// 	assert.Equal(t, "pudu", users[0].FirstName)
// 	assert.Equal(t, "lee", users[0].LastName)
// 	assert.Equal(t, "pudu@gmail.com", users[0].Email)
// 	assert.Equal(t, "08564321789", users[0].PhoneNumber)
// }

// func TestGetAllUser_Failure(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockUserRepo := mocks.NewMockIUserRepository(ctrl)

// 	userService := NewUserService(mockUserRepo)

// 	mockUserRepo.EXPECT().GetAllUser().Return(nil, someError)

// 	users, err := userService.GetAllUser()
// 	assert.Error(t, err)
// 	assert.Nil(t, users)
// }


// func TestGetUser_Success(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockUserRepo := mocks.NewMockIUserRepository(ctrl)

// 	userService := NewUserService(mockUserRepo)

// 	mockUserRepo.EXPECT().GetUser(gomock.Any()).Return(response.User{}, nil)

// 	result, err := userService.GetUser("pudu")

// 	assert.NoError(t, err)
// 	assert.Equal(t, response.User{}, result)
// }

// func TestGetUser_Failure(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockUserRepo := mocks.NewMockIUserRepository(ctrl)

// 	userService := NewUserService(mockUserRepo)

// 	mockUserRepo.EXPECT().GetUser(gomock.Any()).Return(response.User{}, someError)

// 	result, err := userService.GetUser("pudu")

// 	assert.Error(t, err)
// 	assert.Equal(t, response.User{}, result)
// }

// func TestLoginUser_Success(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockUserRepo := mocks.NewMockIUserRepository(ctrl)

// 	userService := NewUserService(mockUserRepo)

// 	birthDate, err := time.Parse(time.RFC3339, "2023-12-31T12:00:00Z")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	mockUserRepo.EXPECT().LoginUser(gomock.Any()).Return(response.User{}, nil)

// 	result, err := userService.LoginUser(&request.User{
// 		Id:          1,
// 		FirstName:   "pudu",
// 		LastName:    "lee",
// 		Email:       "pudu@gmail.com",
// 		Password:    "123",
// 		PhoneNumber: "08564321789",
// 		BirthDate:   birthDate, 
// 	})

// 	assert.NoError(t, err, result)
// }

// func TestLoginUser_Failure(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockUserRepo := mocks.NewMockIUserRepository(ctrl)

// 	userService := NewUserService(mockUserRepo)
// 	mockUserRepo.EXPECT().LoginUser(gomock.Any()).Return(response.User{}, someError)

// 	birthDate, err := time.Parse(time.RFC3339, "2023-12-31T12:00:00Z")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	result, err := userService.LoginUser(&request.User{
// 		Id:          1,
// 		FirstName:   "pudu",
// 		LastName:    "lee",
// 		Email:       "pudu@gmail.com",
// 		Password:    "123",
// 		PhoneNumber: "08564321789",
// 		BirthDate:   birthDate, 
// 	})

// 	assert.Error(t, err)
// 	assert.Equal(t, response.User{}, result)
// }

// func TestRegisterUser_Success(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockUserRepo := mocks.NewMockIUserRepository(ctrl)

// 	userService := NewUserService(mockUserRepo)

// 	mockUserRepo.EXPECT().RegisterUser(gomock.Any()).Return(response.User{}, nil)

// 	birthDate, err := time.Parse(time.RFC3339, "2023-12-31T12:00:00Z")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	result, err := userService.RegisterUser(&request.User{
// 		Id:          1,
// 		FirstName:   "pudu",
// 		LastName:    "lee",
// 		Email:       "pudu@gmail.com",
// 		Password:    "123",
// 		PhoneNumber: "08564321789",
// 		BirthDate:   birthDate, 
// 	})

// 	assert.NoError(t, err, result)
// }

// func TestRegisterUser_Failure(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockUserRepo := mocks.NewMockIUserRepository(ctrl)

// 	userService := NewUserService(mockUserRepo)

// 	mockUserRepo.EXPECT().RegisterUser(gomock.Any()).Return(response.User{}, someError)

// 	birthDate, err := time.Parse(time.RFC3339, "2023-12-31T12:00:00Z")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	result, err := userService.RegisterUser(&request.User{
// 		Id:          1,
// 		FirstName:   "pudu",
// 		LastName:    "lee",
// 		Email:       "pudu@gmail.com",
// 		Password:    "123",
// 		PhoneNumber: "08564321789",
// 		BirthDate:   birthDate, 
// 	})

// 	assert.Error(t, err)
// 	assert.Equal(t, response.User{}, result)
// }

// func TestUpdateUser_Success(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockUserRepo := mocks.NewMockIUserRepository(ctrl)

// 	userService := NewUserService(mockUserRepo)

// 	mockUserRepo.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(response.User{}, nil)

// 	result, err := userService.UpdateUser("pudu", request.User{})

// 	assert.NoError(t, err)
// 	assert.Equal(t, response.User{}, result)
// }

// func TestUpdateUser_Failure(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockUserRepo := mocks.NewMockIUserRepository(ctrl)

// 	userService := NewUserService(mockUserRepo)

// 	mockUserRepo.EXPECT().UpdateUser(gomock.Any(), gomock.Any()).Return(response.User{}, someError)

// 	result, err := userService.UpdateUser("pudu", request.User{})

// 	assert.Error(t, err)
// 	assert.Equal(t, response.User{}, result)
// }