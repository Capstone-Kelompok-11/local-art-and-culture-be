package services

import (
	"errors"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateRole_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRoleRepository := mocks.NewMockIRoleRepository(ctrl)

    mockRequestRole := &request.Role{
        Role: "test",
    }

    expectedResponseRole := response.Role{
        Id:      1,
        Role: "test",
    }

    mockRoleRepository.EXPECT().CreateRole(mockRequestRole).Return(expectedResponseRole, nil)

    RoleService := NewRoleService(mockRoleRepository)
    result, err := RoleService.CreateRole(mockRequestRole)

    if err != nil {
        t.Errorf("Expected no error, but got: %v", err)
    }

    if result.Id != expectedResponseRole.Id || result.Role != expectedResponseRole.Role {
        t.Errorf("Expected response %v, but got %v", expectedResponseRole, result)
    }
}

func TestCreateRole_Failure_RoleIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRoleRepository := mocks.NewMockIRoleRepository(ctrl)

    mockRequestRole := &request.Role{
		Role: "",
    }

    expectedError := errors.New("role is empty")

    RoleService := NewRoleService(mockRoleRepository)

    result, err := RoleService.CreateRole(mockRequestRole)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.Role!= "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestCreateRole_Failure_DatabaseError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRoleRepository := mocks.NewMockIRoleRepository(ctrl)

    mockRequestRole := &request.Role{
        Role: "test",
    }

    expectedError := errors.New("failed to create new Role")

    mockRoleRepository.EXPECT().CreateRole(mockRequestRole).Return(response.Role{}, expectedError)

    RoleService := NewRoleService(mockRoleRepository)

    result, _ := RoleService.CreateRole(mockRequestRole)

    if result.Id != 0 || result.Role!= "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeleteRole_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRoleRepository := mocks.NewMockIRoleRepository(ctrl)

    expectedResponse := response.Role{
        Id:     1,
		Role: "test",
    }

    mockRoleRepository.EXPECT().DeleteRole(gomock.Any()).Return(expectedResponse, nil)

    RoleService := NewRoleService(mockRoleRepository)

    _, err := RoleService.DeleteRole("RoleID")

    if err != nil {
        t.Errorf("Expected no error, but got %v", err)
    }
}

func TestDeleteRole_Failure_ErrDelete(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRoleRepository := mocks.NewMockIRoleRepository(ctrl)

    expectedError := errors.New("can't delete this role")

    mockRoleRepository.EXPECT().DeleteRole(gomock.Any()).Return(response.Role{}, errors.New("failed to delete data from database"))

    RoleService := NewRoleService(mockRoleRepository)

    result, err := RoleService.DeleteRole("RoleID")

    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error %v, but got %v", expectedError, err)
    }

	if result.Id != 0 || result.Role!= "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeleteRole_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRoleRepository := mocks.NewMockIRoleRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	RoleService := NewRoleService(mockRoleRepository)

	_, err := RoleService.DeleteRole("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestMockIRoleService_GetAllRole_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRoleRepository := mocks.NewMockIRoleRepository(ctrl)

	expectedCategories := []response.Role{
		{Id: 1, Role: "Role1"},
		{Id: 2, Role: "Role2"},
	}
	mockRoleRepository.EXPECT().GetAllRole("filter").Return(expectedCategories, nil).Times(1)

	RoleService := NewRoleService(mockRoleRepository)

	resultCategories, err := RoleService.GetAllRole("filter")

	assert.NoError(t, err)
	assert.Equal(t, expectedCategories, resultCategories)
}

func TestGetAllRole_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRoleRepository := mocks.NewMockIRoleRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")
	mockRoleRepository.EXPECT().GetAllRole("filter").Return(nil, expectedError).Times(1)

	RoleService := NewRoleService(mockRoleRepository)

	resultCategories, err := RoleService.GetAllRole("filter")

	assert.Error(t, err)
	assert.Nil(t, resultCategories)
	assert.EqualError(t, err, expectedError.Error())
}

func TestGetRole_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRoleRepository := mocks.NewMockIRoleRepository(ctrl)

	expectedRole := response.Role{
		Id:   1,
		Role: "test",
	}

	mockRoleRepository.EXPECT().GetRole("RoleID").Return(expectedRole, nil).Times(1)

	RoleService := NewRoleService(mockRoleRepository)

	resultRole, err := RoleService.GetRole("RoleID")

	assert.NoError(t, err)
	assert.Equal(t, expectedRole, resultRole)
}

func TestGetRole_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRoleRepository := mocks.NewMockIRoleRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockRoleRepository.EXPECT().GetRole("NonExistentRole").Return(response.Role{}, expectedError).Times(1)

	RoleService := NewRoleService(mockRoleRepository)

	resultRole, err := RoleService.GetRole("NonExistentRole")

	assert.EqualError(t, err, expectedError.Error())
	assert.Equal(t, response.Role{}, resultRole)
}

func TestGetRole_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRoleRepository := mocks.NewMockIRoleRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	RoleService := NewRoleService(mockRoleRepository)

	_, err := RoleService.GetRole("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestUpdateRole_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRoleRepository := mocks.NewMockIRoleRepository(ctrl)

	expectedRole := response.Role{
		Id:   1,
		Role: "Update",
	}
	mockRoleRepository.EXPECT().UpdateRole("RoleID", gomock.Any()).Return(expectedRole, nil).Times(1)

	RoleService := NewRoleService(mockRoleRepository)
	resultRole, err := RoleService.UpdateRole("RoleID", request.Role{Role: "UpdatedRole"})

	assert.NoError(t, err)
	assert.Equal(t, expectedRole, resultRole)
}

func TestUpdateRole_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRoleRepository := mocks.NewMockIRoleRepository(ctrl)

	expectedError := errors.New("failed to update data")
	mockRoleRepository.EXPECT().UpdateRole("RoleID", gomock.Any()).Return(response.Role{}, expectedError).Times(1)

	RoleService := NewRoleService(mockRoleRepository)

	resultRole, err := RoleService.UpdateRole("RoleID", request.Role{Role: "UpdatedRole"})

	assert.EqualError(t, err, expectedError.Error())
	assert.Equal(t, response.Role{}, resultRole)
}

func TestUpdateRole_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRoleRepository := mocks.NewMockIRoleRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	RoleService := NewRoleService(mockRoleRepository)

	_, err := RoleService.UpdateRole("", request.Role{})

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}