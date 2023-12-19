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

func TestDeleteAdmin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

	expectedAdmin := response.SuperAdmin{Id: 1, Name: "test", Email: "test@gmail.com", Token: "", PhoneNumber: "081342423", RoleId: 1, Role: response.Role{}}

	mockAdminRepository.EXPECT().DeleteAdmin("AdminID").Return(expectedAdmin, nil).Times(1)

	AdminService := NewAdminService(mockAdminRepository)

	resultAdmin, err := AdminService.DeleteAdmin("AdminID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultAdmin, expectedAdmin) {
		t.Errorf("Expected Admin %+v, but got %+v", expectedAdmin, resultAdmin)
	}
}

func TestDeleteAdmin_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	AdminService := NewAdminService(mockAdminRepository)

	_, err := AdminService.DeleteAdmin("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestDeleteAdmin_Failure_ErrDelete(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

    expectedError := errors.New("can't delete this admin")

    mockAdminRepository.EXPECT().DeleteAdmin(gomock.Any()).Return(response.SuperAdmin{}, errors.New("failed to delete data from database"))

    AdminService := NewAdminService(mockAdminRepository)

    result, err := AdminService.DeleteAdmin("AdminID")

    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error %v, but got %v", expectedError, err)
    }

    if result.Id != 0 || result.Name != "" || result.Email != "" || result.PhoneNumber != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestGetAdmin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

	expectedAdmin := response.SuperAdmin{Id: 1, Name: "test", Email: "test@gmail.com", Token: "", PhoneNumber: "081342423", RoleId: 1, Role: response.Role{}}

	mockAdminRepository.EXPECT().GetAdmin("AdminID").Return(expectedAdmin, nil).Times(1)

	AdminService := NewAdminService(mockAdminRepository)

	resultAdmin, err := AdminService.GetAdmin("AdminID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultAdmin, expectedAdmin) {
		t.Errorf("Expected Admin %+v, but got %+v", expectedAdmin, resultAdmin)
	}
}

func TestGetAdmin_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockAdminRepository.EXPECT().GetAdmin("AdminID").Return(response.SuperAdmin{}, expectedError).Times(1)

	AdminService := NewAdminService(mockAdminRepository)

	result, err := AdminService.GetAdmin("AdminID")

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if result.Id != 0 || result.Name != "" || result.Email != "" || result.PhoneNumber != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestGetAdmin_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	AdminService := NewAdminService(mockAdminRepository)

	_, err := AdminService.GetAdmin("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestUpdateAdmin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

	expectedUpdatedAdmin := response.SuperAdmin{Id: 1, Name: "test", Email: "test@gmail.com", Token: "", PhoneNumber: "081342423", RoleId: 1, Role: response.Role{}}

	var expectedError error

	mockAdminRepository.EXPECT().UpdateAdmin("AdminID", gomock.Any()).Return(expectedUpdatedAdmin, expectedError).Times(1)

	AdminService := NewAdminService(mockAdminRepository)

	requestAdmin := request.SuperAdmin{Name: "test", Email: "test@gmail.com", PhoneNumber: "081342423", RoleId: 1}

	resultAdmin, err := AdminService.UpdateAdmin("AdminID", requestAdmin)

	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if !reflect.DeepEqual(resultAdmin, expectedUpdatedAdmin) {
		t.Errorf("Expected %+v, but got %+v", expectedUpdatedAdmin, resultAdmin)
	}
}

func TestUpdateAdmin_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

	expectedError := errors.New("failed to update data")

	mockAdminRepository.EXPECT().UpdateAdmin("AdminID", gomock.Any()).Return(response.SuperAdmin{}, expectedError).Times(1)

	AdminService := NewAdminService(mockAdminRepository)

	requestAdmin := request.SuperAdmin{ Name: "test", Email: "test@gmail.com", PhoneNumber: "081342423", RoleId: 1}
	result, err := AdminService.UpdateAdmin("AdminID", requestAdmin)

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}
	if result.Id != 0 || result.Name != "" || result.Email != "" || result.PhoneNumber != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestUpdateAdmin_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	AdminService := NewAdminService(mockAdminRepository)

	_, err := AdminService.UpdateAdmin("", request.SuperAdmin{})

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestCalculatePaginationValues_Admin(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

    AdminService := NewAdminService(mockAdminRepository)

    page1, totalPages1 := AdminService.CalculatePaginationValues(0, 100, 8)
    assert.Equal(t, 1, page1)
    assert.Equal(t, 1, totalPages1)

    page2, totalPages2 := AdminService.CalculatePaginationValues(15, 100, 8)
    assert.Equal(t, 1, page2)
    assert.Equal(t, 1, totalPages2)

    page3, totalPages3 := AdminService.CalculatePaginationValues(2, 100, 8)
    assert.Equal(t, 1, page3)
    assert.Equal(t, 1, totalPages3)

    page4, totalPages4 := AdminService.CalculatePaginationValues(1, 95, 8)
    assert.Equal(t, 1, page4)
    assert.Equal(t, 1, totalPages4) 
}

func TestGetNextPage_Admin(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

    AdminService := NewAdminService(mockAdminRepository)

    result1 := AdminService.GetNextPage(3, 10)
	assert.Equal(t, 4, result1)

	result2 := AdminService.GetNextPage(10, 10)
	assert.Equal(t, 10, result2)
}

func TestGetPrevPage_Admin(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

    AdminService := NewAdminService(mockAdminRepository)

    result1 := AdminService.GetPrevPage(5)
	assert.Equal(t, 4, result1)

	result2 := AdminService.GetPrevPage(1)
	assert.Equal(t, 1, result2)
}

func TestGetAllAdmin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

	expectedAdmins := []response.SuperAdmin{
		{Id: 1, Name: "test", Email: "test@gmail.com", Token: "", PhoneNumber: "081342423", RoleId: 1, Role: response.Role{}},
		{Id: 1, Name: "test", Email: "test@gmail.com", Token: "", PhoneNumber: "081342423", RoleId: 1, Role: response.Role{}},
	}

	expectedTotal := len(expectedAdmins)
	var expectedError error

	mockAdminRepository.EXPECT().GetAllAdmin("category", 0, 10).Return(expectedAdmins, expectedTotal, expectedError).Times(1)
	AdminService := NewAdminService(mockAdminRepository)
	resultAdmins, _, err := AdminService.GetAllAdmin("category", 0, 10)

	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}
	if len(resultAdmins) != expectedTotal {
		t.Errorf("Expected %d Admins, but got %d", expectedTotal, len(resultAdmins))
	}
}

func TestGetAllAdmin_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockAdminRepository.EXPECT().GetAllAdmin("category", 0, 10).Return(nil, 0, expectedError).Times(1)

	AdminService := NewAdminService(mockAdminRepository)

	resultAdmins, resultTotal, err := AdminService.GetAllAdmin("category", 0, 10)

	if err.Error() != expectedError.Error(){
		t.Errorf("Expected error containing '%v', but got '%v'", expectedError.Error(), err)
	}

	if resultAdmins != nil || resultTotal != 0 {
		t.Errorf("Expected empty result, but got Admins: %v, total: %v", resultAdmins, resultTotal)
	}
}

func TestLoginAdmin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

    expectedAdmin := &request.SuperAdmin{Name: "test", Email: "test@gmail.com", PhoneNumber: "081342423", RoleId: 1, Password: "123"}
	expectedAdmins := response.SuperAdmin{Id: 1, Name: "test", Email: "test@gmail.com", Token: "", PhoneNumber: "081342423", RoleId: 1, Role: response.Role{}}

    mockAdminRepository.EXPECT().
        LoginAdmin(expectedAdmin).
        Return(expectedAdmins, nil).
        Times(1)

    AdminService := NewAdminService(mockAdminRepository)
    Admin, err := AdminService.LoginAdmin(expectedAdmin)

    assert.NoError(t, err)
    assert.Equal(t, expectedAdmins, Admin)
}

func TestLoginAdmin_Failure_EmailIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

    mockRequestAdmin := &request.SuperAdmin{Name: "test", Email: "", PhoneNumber: "081342423", RoleId: 1, Password: "123"}
    expectedError := errors.New("email is empty")

    AdminService := NewAdminService(mockAdminRepository)
    result, err := AdminService.LoginAdmin(mockRequestAdmin)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }
    if result.Id != 0 || result.Name != "" || result.Email != "" || result.PhoneNumber != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestLoginAdmin_Failure_PassIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

    mockRequestAdmin := &request.SuperAdmin{Name: "test", Email: "test@gmail.com", PhoneNumber: "081342423", RoleId: 1, Password: ""}
    expectedError := errors.New("password is empty")

    AdminService := NewAdminService(mockAdminRepository)
    result, err := AdminService.LoginAdmin(mockRequestAdmin)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }
    if result.Id != 0 || result.Name != "" || result.Email != "" || result.PhoneNumber != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestLoginAdmin_Failure_DatabaseError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

    mockRequestAdmin := &request.SuperAdmin{Name: "test", Email: "test@gmail.com", PhoneNumber: "081342423", RoleId: 1, Password: "123"}
	expectedError := errors.New("failed to login Admin")

    mockAdminRepository.EXPECT().LoginAdmin(mockRequestAdmin).Return(response.SuperAdmin{}, expectedError)
    AdminService := NewAdminService(mockAdminRepository)
	result, err := AdminService.LoginAdmin(mockRequestAdmin)

    if err != expectedError {
        t.Errorf("Expected error: %v, but got: %v", expectedError, err)
    }
	if result.Id != 0 || result.Name != "" || result.Email != "" || result.PhoneNumber != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestRegisterAdmin_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

    mockAdminRepository.EXPECT().
        RegisterAdmin(gomock.Any()).
        Return(response.SuperAdmin{Id: 1, Name: "test", Email: "test@gmail.com", Token: "", PhoneNumber: "081342423", RoleId: 1, Role: response.Role{}}, nil)

    AdminService := NewAdminService(mockAdminRepository)
    AdminRequest := &request.SuperAdmin{Name: "test", Email: "test@gmail.com", PhoneNumber: "081342423", RoleId: 1, Password: "123"}
    resultAdmin, err := AdminService.RegisterAdmin(AdminRequest)

    assert.NoError(t, err)
    assert.NotNil(t, resultAdmin)
    assert.Equal(t, "test", resultAdmin.Name)
    assert.Equal(t, uint(1), resultAdmin.Id)
}

func TestRegistereAdmin_Failure_NameIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

    mockRequestAdmin := &request.SuperAdmin{Name: "", Email: "test@gmail.com", PhoneNumber: "081342423", RoleId: 1, Password: "123"}

    expectedError := errors.New("name is empty")

    AdminService := NewAdminService(mockAdminRepository)
    result, err := AdminService.RegisterAdmin(mockRequestAdmin)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }
	if result.Id != 0 || result.Name != "" || result.Email != "" || result.PhoneNumber != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestRegistereAdmin_Failure_PassIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

    mockRequestAdmin := &request.SuperAdmin{Name: "test", Email: "test@gmail.com", PhoneNumber: "081342423", RoleId: 1, Password: ""}

    expectedError := errors.New("password is empty")

    AdminService := NewAdminService(mockAdminRepository)
    result, err := AdminService.RegisterAdmin(mockRequestAdmin)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }
	if result.Id != 0 || result.Name != "" || result.Email != "" || result.PhoneNumber != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestRegistereAdmin_Failure_PhoneIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

    mockRequestAdmin := &request.SuperAdmin{Name: "test", Email: "test@gmail.com", PhoneNumber: "", RoleId: 1, Password: "123"}

    expectedError := errors.New("phone number is empty")

    AdminService := NewAdminService(mockAdminRepository)
    result, err := AdminService.RegisterAdmin(mockRequestAdmin)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }
	if result.Id != 0 || result.Name != "" || result.Email != "" || result.PhoneNumber != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestRegistereAdmin_Failure_EmailIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

    mockRequestAdmin := &request.SuperAdmin{Name: "test", Email: "", PhoneNumber: "081342423", RoleId: 1, Password: "123"}

    expectedError := errors.New("email is empty")

    AdminService := NewAdminService(mockAdminRepository)
    result, err := AdminService.RegisterAdmin(mockRequestAdmin)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }
	if result.Id != 0 || result.Name != "" || result.Email != "" || result.PhoneNumber != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestRegisterAdmin_Failure_DatabaseError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockAdminRepository := mocks.NewMockIAdminRepository(ctrl)

    mockRequestAdmin := &request.SuperAdmin{Name: "test", Email: "test@gmail.com", PhoneNumber: "081342423", RoleId: 1, Password: "123"}

    expectedError := errors.New("failed to create new Admin")

    mockAdminRepository.EXPECT().RegisterAdmin(mockRequestAdmin).Return(response.SuperAdmin{}, expectedError)

    AdminService := NewAdminService(mockAdminRepository)

    result, _ := AdminService.RegisterAdmin(mockRequestAdmin)

	if result.Id != 0 || result.Name != "" || result.Email != "" || result.PhoneNumber != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}