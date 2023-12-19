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

func TestCountUsersByRole_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mocks.NewMockIUserRepository(ctrl)

	expectedCount := 5
	expectedRoleID := uint(1)

	mockUserRepository.EXPECT().
		CountUsersByRole(expectedRoleID).
		Return(expectedCount, nil).
		Times(1)

	userService := NewUserService(mockUserRepository)

	count, err := userService.CountUsersByRole(expectedRoleID)

	assert.NoError(t, err)
	assert.Equal(t, expectedCount, count)
}

func TestCountUsersByRole_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mocks.NewMockIUserRepository(ctrl)

	expectedError := errors.New("some error")
	expectedRoleID := uint(1)

	mockUserRepository.EXPECT().
		CountUsersByRole(expectedRoleID).
		Return(0, expectedError).
		Times(1)

	userService := NewUserService(mockUserRepository)

	count, err := userService.CountUsersByRole(expectedRoleID)

	assert.Error(t, err)
	assert.Equal(t, 0, count) 
}

func TestDeleteUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mocks.NewMockIUserRepository(ctrl)

	expectedUser := response.User{Id:1, FirstName: "test", LastName: "test", Username: "test", Email: "test@gmail.com", Token: "", PhoneNumber: "081342423", NIK: "12345668", Gender: "pr", BirthDate: time.Time{}, Status: "test", DeletedAt: time.Time{}, CreatedAt: time.Time{}, Role: response.Role{}}

	mockUserRepository.EXPECT().DeleteUser("UserID").Return(expectedUser, nil).Times(1)

	UserService := NewUserService(mockUserRepository)

	resultUser, err := UserService.DeleteUser("UserID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultUser, expectedUser) {
		t.Errorf("Expected User %+v, but got %+v", expectedUser, resultUser)
	}
}

func TestDeleteUser_Failure_ErrDelete(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepository := mocks.NewMockIUserRepository(ctrl)

    expectedError := errors.New("can't delete this user")

    mockUserRepository.EXPECT().DeleteUser(gomock.Any()).Return(response.User{}, errors.New("failed to delete data from database"))

    UserService := NewUserService(mockUserRepository)

    result, err := UserService.DeleteUser("UserID")

    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error %v, but got %v", expectedError, err)
    }

    if result.Id != 0 || result.FirstName != "" || result.LastName != "" || result.Username != "" || result.Email != "" || result.PhoneNumber != "" || result.NIK != "" || result.Gender != ""{
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeleteUser_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mocks.NewMockIUserRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	UserService := NewUserService(mockUserRepository)

	_, err := UserService.DeleteUser("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestGetUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mocks.NewMockIUserRepository(ctrl)

	expectedUser := response.User{Id:1, FirstName: "test", LastName: "test", Username: "test", Email: "test@gmail.com", Token: "", PhoneNumber: "081342423", NIK: "12345668", Gender: "pr", BirthDate: time.Time{}, Status: "test", DeletedAt: time.Time{}, CreatedAt: time.Time{}, Role: response.Role{}}

	mockUserRepository.EXPECT().GetUser("UserID").Return(expectedUser, nil).Times(1)

	UserService := NewUserService(mockUserRepository)

	resultUser, err := UserService.GetUser("UserID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultUser, expectedUser) {
		t.Errorf("Expected User %+v, but got %+v", expectedUser, resultUser)
	}
}

func TestGetUser_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mocks.NewMockIUserRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockUserRepository.EXPECT().GetUser("UserID").Return(response.User{}, expectedError).Times(1)

	UserService := NewUserService(mockUserRepository)

	result, err := UserService.GetUser("UserID")

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if result.Id != 0 || result.FirstName != "" || result.LastName != "" || result.Username != "" || result.Email != "" || result.PhoneNumber != "" || result.NIK != "" || result.Gender != ""{
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestGetUser_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mocks.NewMockIUserRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	UserService := NewUserService(mockUserRepository)

	_, err := UserService.GetUser("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestUpdateUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mocks.NewMockIUserRepository(ctrl)

	expectedUpdatedUser := response.User{Id:1, FirstName: "test", LastName: "test", Username: "test", Email: "test@gmail.com", Token: "", PhoneNumber: "081342423", NIK: "12345668", Gender: "pr", BirthDate: time.Time{}, Status: "test", DeletedAt: time.Time{}, CreatedAt: time.Time{}, Role: response.Role{}}

	var expectedError error

	mockUserRepository.EXPECT().UpdateUser("UserID", gomock.Any()).Return(expectedUpdatedUser, expectedError).Times(1)

	UserService := NewUserService(mockUserRepository)

	requestUser := request.User{FirstName: "test", LastName: "test", Username: "test", Email: "test@gmail.com", PhoneNumber: "081342423", NIK: "12345668", Gender: "pr", BirthDate: time.Time{}}

	resultUser, err := UserService.UpdateUser("UserID", requestUser)

	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if !reflect.DeepEqual(resultUser, expectedUpdatedUser) {
		t.Errorf("Expected %+v, but got %+v", expectedUpdatedUser, resultUser)
	}
}

func TestUpdateUser_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mocks.NewMockIUserRepository(ctrl)

	expectedError := errors.New("failed to update data")

	mockUserRepository.EXPECT().UpdateUser("UserID", gomock.Any()).Return(response.User{}, expectedError).Times(1)

	UserService := NewUserService(mockUserRepository)

	requestUser := request.User{FirstName: "test", LastName: "test", Username: "test", Email: "test@gmail.com", PhoneNumber: "081342423", NIK: "12345668", Gender: "pr", BirthDate: time.Time{}}
	result, err := UserService.UpdateUser("UserID", requestUser)

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}
	if result.Id != 0 || result.FirstName != "" || result.LastName != "" || result.Username != "" || result.Email != "" || result.PhoneNumber != "" || result.NIK != "" || result.Gender != ""{
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestUpdateUser_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mocks.NewMockIUserRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	UserService := NewUserService(mockUserRepository)

	_, err := UserService.UpdateUser("", request.User{})

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestCalculatePaginationValues_User(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepository := mocks.NewMockIUserRepository(ctrl)

    UserService := NewUserService(mockUserRepository)

    page1, totalPages1 := UserService.CalculatePaginationValues(0, 100, 8)
    assert.Equal(t, 1, page1)
    assert.Equal(t, 1, totalPages1)

    page2, totalPages2 := UserService.CalculatePaginationValues(15, 100, 8)
    assert.Equal(t, 1, page2)
    assert.Equal(t, 1, totalPages2)

    page3, totalPages3 := UserService.CalculatePaginationValues(2, 100, 8)
    assert.Equal(t, 1, page3)
    assert.Equal(t, 1, totalPages3)

    page4, totalPages4 := UserService.CalculatePaginationValues(1, 95, 8)
    assert.Equal(t, 1, page4)
    assert.Equal(t, 1, totalPages4) 
}

func TestGetNextPage_User(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepository := mocks.NewMockIUserRepository(ctrl)

    UserService := NewUserService(mockUserRepository)

    result1 := UserService.GetNextPage(3, 10)
	assert.Equal(t, 4, result1)

	result2 := UserService.GetNextPage(10, 10)
	assert.Equal(t, 10, result2)
}

func TestGetPrevPage_User(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepository := mocks.NewMockIUserRepository(ctrl)

    UserService := NewUserService(mockUserRepository)

    result1 := UserService.GetPrevPage(5)
	assert.Equal(t, 4, result1)

	result2 := UserService.GetPrevPage(1)
	assert.Equal(t, 1, result2)
}

func TestGetAllUser_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepository := mocks.NewMockIUserRepository(ctrl)

	expectedUsers := []response.User{
        {Id: 1, FirstName: "test", LastName: "test", Username: "test", Email: "test@gmail.com", Token: "", PhoneNumber: "081342423", NIK: "12345668",
            Gender: "pr", BirthDate: time.Time{}, Status: "test", DeletedAt: time.Time{}, CreatedAt: time.Time{}, Role: response.Role{},},
        {Id: 2, FirstName: "test2", LastName: "test2", Username: "test2", Email: "test2@gmail.com", Token: "", PhoneNumber: "0813424232", NIK: "123456682",
            Gender: "pr", BirthDate: time.Time{}, Status: "test", DeletedAt: time.Time{}, CreatedAt: time.Time{}, Role: response.Role{},},
    }

    expectedUserCounts := map[string]int{
        "EventCreators": 15, 
		"ProductCreators": 15, 
		"RegularUser": 15,
    }

    mockUserRepository.EXPECT().
        GetAllUser("filter", gomock.Any(), gomock.Any()).
        DoAndReturn(func(_ string, _ int, _ int) ([]response.User, map[string]int, error) {
            return expectedUsers, expectedUserCounts, nil
        }).
        Times(1)

	mockUserRepository.EXPECT().
    CountUsersByRole(gomock.Any()).
    Return(15, nil).
    AnyTimes()

    userService := NewUserService(mockUserRepository)

    resultUsers, resultUserCounts, err := userService.GetAllUser("filter", 1, 10)

    assert.NoError(t, err)
    assert.Equal(t, expectedUsers, resultUsers)
    assert.Equal(t, expectedUserCounts, resultUserCounts)
}

func TestGetAllUser_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepository := mocks.NewMockIUserRepository(ctrl)

    expectedError := errors.New("mocked error")

    mockUserRepository.EXPECT().
        GetAllUser("filter", gomock.Any(), gomock.Any()).
        Return(nil, 0, expectedError).
        Times(1)

    mockUserRepository.EXPECT().
        CountUsersByRole(gomock.Any()).
        Return(0, expectedError).
        AnyTimes()

    userService := NewUserService(mockUserRepository)

    resultUsers, resultUserCounts, err := userService.GetAllUser("filter", 1, 10)

    assert.Error(t, err)
    assert.Nil(t, resultUsers)
    assert.Nil(t, resultUserCounts)
    assert.Equal(t, expectedError, err)
}

func TestGetAllUser_Failure_CountProduct(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepository := mocks.NewMockIUserRepository(ctrl)

    expectedError := errors.New("mocked error")

    mockUserRepository.EXPECT().
        GetAllUser("filter", gomock.Any(), gomock.Any()).
        Return(nil, 0, nil).
        Times(1)

    mockUserRepository.EXPECT().
        CountUsersByRole(uint(1)).
        Return(0, expectedError).
        AnyTimes()

    userService := NewUserService(mockUserRepository)

    _, resultUserCounts, err := userService.GetAllUser("filter", 1, 10)

    assert.Error(t, err)
    assert.Nil(t, resultUserCounts)
    assert.Equal(t, expectedError, err)
}

func TestLoginUser_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepository := mocks.NewMockIUserRepository(ctrl)

    expectedUser := &request.User{FirstName: "test", LastName: "test", Username: "test", Password: "123", Email: "test@gmail.com", PhoneNumber: "081342423", NIK: "12345668", Gender: "pr", BirthDate: time.Time{}}
	expectedCreators := response.Creators{Id: 1, OutletName: "test", Email: "test@gmail.com", PhoneNumber: "081342423", Address: "test", Roles: "test", Users: response.Users{}, Role: response.Role{}}

    mockUserRepository.EXPECT().
        LoginUser(expectedUser).
        Return(expectedCreators, nil).
        Times(1)

    userService := NewUserService(mockUserRepository)
    creators, err := userService.LoginUser(expectedUser)

    assert.NoError(t, err)
    assert.Equal(t, expectedCreators, creators)
}

func TestLoginUser_Failure_EmailIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepository := mocks.NewMockIUserRepository(ctrl)

    mockRequestUser := &request.User{FirstName: "test", LastName: "test", Username: "test", Password: "123", Email: "", PhoneNumber: "081342423", NIK: "12345668", Gender: "pr", BirthDate: time.Time{},}
    expectedError := errors.New("email is empty")

    UserService := NewUserService(mockUserRepository)
    result, err := UserService.LoginUser(mockRequestUser)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }
    if result.Id != 0 || result.OutletName != "" || result.Roles != "" || result.Address!= "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestLoginUser_Failure_PassIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepository := mocks.NewMockIUserRepository(ctrl)
    mockRequestUser := &request.User{FirstName: "test", LastName: "test", Username: "test", Password: "", Email: "test@gmail.com", PhoneNumber: "081342423", NIK: "12345668", Gender: "pr", BirthDate: time.Time{}}

    expectedError := errors.New("password is empty")

    UserService := NewUserService(mockUserRepository)
    result, err := UserService.LoginUser(mockRequestUser)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }
    if result.Id != 0 || result.OutletName != "" || result.Roles != "" || result.Address!= "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestLoginUser_Failure_DatabaseError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepository := mocks.NewMockIUserRepository(ctrl)

    mockRequestUser := &request.User{FirstName: "test", LastName: "test", Username: "test", Password: "123", Email: "test@gmail.com", PhoneNumber: "081342423", NIK: "12345668", Gender: "pr", BirthDate: time.Time{},}
	expectedError := errors.New("failed to login user")

    mockUserRepository.EXPECT().LoginUser(mockRequestUser).Return(response.Creators{}, expectedError)
    userService := NewUserService(mockUserRepository)
	result, err := userService.LoginUser(mockRequestUser)

    if err != expectedError {
        t.Errorf("Expected error: %v, but got: %v", expectedError, err)
    }

    if result.Id != 0 || result.OutletName != "" || result.Roles != "" || result.Address != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestRegisterUser_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepository := mocks.NewMockIUserRepository(ctrl)

    mockUserRepository.EXPECT().
        RegisterUser(gomock.Any()).
        Return(response.User{Id: 1, FirstName: "test", LastName: "test", Username: "test", Token: "", Email: "test@gmail.com", PhoneNumber: "081342423", NIK: "12345668", Gender: "pr", BirthDate: time.Time{}}, nil).Times(1)

    userService := NewUserService(mockUserRepository)
    userRequest := &request.User{FirstName: "test", LastName: "test", Username: "test", Password: "123", Email: "test@gmail.com", PhoneNumber: "081342423", NIK: "12345668", Gender: "pr", BirthDate: time.Time{}}
    resultUser, err := userService.RegisterUser(userRequest)

    assert.NoError(t, err)
    assert.NotNil(t, resultUser)
    assert.Equal(t, "test", resultUser.Username)
    assert.Equal(t, uint(1), resultUser.Id)
}

func TestRegistereUser_Failure_NameIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepository := mocks.NewMockIUserRepository(ctrl)

    mockRequestUser := &request.User{FirstName: "", LastName: "test", Username: "test", Password: "123", Email: "test@gmail.com", PhoneNumber: "081342423", NIK: "12345668", Gender: "pr", BirthDate: time.Time{},}

    expectedError := errors.New("name is empty")

    UserService := NewUserService(mockUserRepository)
    result, err := UserService.RegisterUser(mockRequestUser)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.FirstName != "" || result.LastName != "" || result.Username != "" || result.Email != "" || result.PhoneNumber != "" || result.NIK != "" || result.Gender != ""{
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestRegistereUser_Failure_Name2IsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepository := mocks.NewMockIUserRepository(ctrl)

    mockRequestUser := &request.User{FirstName: "test", LastName: "", Username: "test", Password: "123", Email: "test@gmail.com", PhoneNumber: "081342423", NIK: "12345668", Gender: "pr", BirthDate: time.Time{},}

    expectedError := errors.New("name is empty")
    UserService := NewUserService(mockUserRepository)

    result, err := UserService.RegisterUser(mockRequestUser)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.FirstName != "" || result.LastName != "" || result.Username != "" || result.Email != "" || result.PhoneNumber != "" || result.NIK != "" || result.Gender != ""{
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestRegistereUser_Failure_EmailIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepository := mocks.NewMockIUserRepository(ctrl)

    mockRequestUser := &request.User{FirstName: "test", LastName: "test", Username: "test", Password: "123", Email: "", PhoneNumber: "081342423", NIK: "12345668", Gender: "pr", BirthDate: time.Time{},}

    expectedError := errors.New("email is empty")
    UserService := NewUserService(mockUserRepository)

    result, err := UserService.RegisterUser(mockRequestUser)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.FirstName != "" || result.LastName != "" || result.Username != "" || result.Email != "" || result.PhoneNumber != "" || result.NIK != "" || result.Gender != ""{
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestRegistereUser_Failure_PhoneIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepository := mocks.NewMockIUserRepository(ctrl)

    mockRequestUser := &request.User{FirstName: "test", LastName: "test", Username: "test", Password: "123", Email: "test@gmail.com", PhoneNumber: "", NIK: "12345668", Gender: "pr", BirthDate: time.Time{},}

    expectedError := errors.New("phone number is empty")
    UserService := NewUserService(mockUserRepository)

    result, err := UserService.RegisterUser(mockRequestUser)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.FirstName != "" || result.LastName != "" || result.Username != "" || result.Email != "" || result.PhoneNumber != "" || result.NIK != "" || result.Gender != ""{
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestRegisterUser_Failure_DatabaseError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockUserRepository := mocks.NewMockIUserRepository(ctrl)

    mockRequestUser := &request.User{FirstName: "test", LastName: "test", Username: "test", Password: "123", Email: "test@gmail.com", PhoneNumber: "086754312", NIK: "12345668", Gender: "pr", BirthDate: time.Time{},}

    expectedError := errors.New("failed to create new User")

    mockUserRepository.EXPECT().RegisterUser(mockRequestUser).Return(response.User{}, expectedError)

    UserService := NewUserService(mockUserRepository)

    result, _ := UserService.RegisterUser(mockRequestUser)

	if result.Id != 0 || result.FirstName != "" || result.LastName != "" || result.Username != "" || result.Email != "" || result.PhoneNumber != "" || result.NIK != "" || result.Gender != ""{
        t.Errorf("Expected empty response, but got: %v", result)
    }
}