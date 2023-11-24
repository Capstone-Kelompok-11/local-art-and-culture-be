package services

// import (
// 	"lokasani/entity/request"
// 	"lokasani/entity/response"
// 	"lokasani/features/mocks"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCreateRole_Success(t *testing.T) {
//     ctrl := gomock.NewController(t)
//     defer ctrl.Finish()

//     mockRoleRepo := mocks.NewMockIRoleRepository(ctrl)

//     roleService := NewRoleService(mockRoleRepo)

//     mockRoleRepo.EXPECT().CreateRole(gomock.Any()).Return(response.Role{}, nil)

//     result, err := roleService.CreateRole(&request.Role{
//         Id: 1,
// 		Role: "user",
//     })

//     assert.NoError(t, err)
//     assert.Equal(t, response.Role{}, result)
// }

// func TestCreateRole_Failure(t *testing.T) {
// 	ctrl := gomock.NewController(t)
//     defer ctrl.Finish()

//     mockRoleRepo := mocks.NewMockIRoleRepository(ctrl)

//     roleService := NewRoleService(mockRoleRepo)

//     mockRoleRepo.EXPECT().CreateRole(gomock.Any()).Return(response.Role{}, someError)

//     result, err := roleService.CreateRole(&request.Role{
//         Id: 1,
// 		Role: "user",
//     })

//     assert.Error(t, err)
//     assert.Equal(t, response.Role{}, result)
// }

// func TestDeleteRole_Success(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRoleRepo := mocks.NewMockIRoleRepository(ctrl)

// 	roleService := NewRoleService(mockRoleRepo)

// 	mockRoleRepo.EXPECT().DeleteRole(gomock.Any()).Return(response.Role{}, nil)

// 	result, err := roleService.DeleteRole("pudu")

// 	assert.NoError(t, err)
// 	assert.Equal(t, response.Role{}, result)
// }

// func TestDeleteRole_Failure(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRoleRepo := mocks.NewMockIRoleRepository(ctrl)

// 	roleService := NewRoleService(mockRoleRepo)

// 	mockRoleRepo.EXPECT().DeleteRole(gomock.Any()).Return(response.Role{}, someError)

// 	result, err := roleService.DeleteRole("pudu")

// 	assert.Error(t, err)
// 	assert.Equal(t, response.Role{}, result)
// }

// func TestGetAllRole_Success(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRoleRepo := mocks.NewMockIRoleRepository(ctrl)

// 	roleService := NewRoleService(mockRoleRepo)

// 	expectedRoles := []response.Role{
// 		{Id: 1, Role: "Admin"},
// 		{Id: 2, Role: "User"},
// 	}
// 	mockRoleRepo.EXPECT().GetAllRole().Return(expectedRoles, nil)

// 	roles, err := roleService.GetAllRole()

// 	assert.NotNil(t, err, expectedRoles, roles)
// }

// func TestGetAllRole_Failure(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRoleRepo := mocks.NewMockIRoleRepository(ctrl)

// 	roleService := NewRoleService(mockRoleRepo)

// 	mockRoleRepo.EXPECT().GetAllRole().Return(nil, someError)

// 	roles, err := roleService.GetAllRole()

// 	assert.Error(t, err)
// 	assert.Nil(t, roles)
// }

// func TestGetRole_Success(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRoleRepo := mocks.NewMockIRoleRepository(ctrl)

// 	roleService := NewRoleService(mockRoleRepo)

// 	mockRoleRepo.EXPECT().GetRole(gomock.Any()).Return(response.Role{}, nil)

// 	result, err := roleService.GetRole("pudu")

// 	assert.NoError(t, err)
// 	assert.Equal(t, response.Role{}, result)
// }

// func TestGetRole_Failure(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRoleRepo := mocks.NewMockIRoleRepository(ctrl)

// 	roleService := NewRoleService(mockRoleRepo)

// 	mockRoleRepo.EXPECT().GetRole(gomock.Any()).Return(response.Role{}, someError)

// 	result, err := roleService.GetRole("pudu")

// 	assert.Error(t, err)
// 	assert.Equal(t, response.Role{}, result)
// }

// func TestUpdateRole_Success(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRoleRepo := mocks.NewMockIRoleRepository(ctrl)

// 	roleService := NewRoleService(mockRoleRepo)

// 	mockRoleRepo.EXPECT().UpdateRole(gomock.Any(), gomock.Any()).Return(response.Role{}, nil)

// 	result, err := roleService.UpdateRole("pudu", request.Role{})

// 	assert.NoError(t, err)
// 	assert.Equal(t, response.Role{}, result)
// }

// func TestUpdateRole_Failure(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRoleRepo := mocks.NewMockIRoleRepository(ctrl)

// 	roleService := NewRoleService(mockRoleRepo)

// 	mockRoleRepo.EXPECT().UpdateRole(gomock.Any(), gomock.Any()).Return(response.Role{}, someError)

// 	result, err := roleService.UpdateRole("pudu", request.Role{})

// 	assert.Error(t, err)
// 	assert.Equal(t, response.Role{}, result)
// }