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

func TestCreateCategory_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	mockRequestCategory := &request.Category{
		Category: "test",
		Type: "test",
	}

	expectedResponseCategory := response.Category{
		Id:      1,
		Category: "test",
		Type: "test",
	}

	mockCategoryRepository.EXPECT().CreateCategory(mockRequestCategory).Return(expectedResponseCategory, nil)

	CategoryService := NewCategoryService(mockCategoryRepository)
	result, err := CategoryService.CreateCategory(mockRequestCategory)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	if result.Id != expectedResponseCategory.Id || result.Category != expectedResponseCategory.Category || result.Type != expectedResponseCategory.Type {
		t.Errorf("Expected response %v, but got %v", expectedResponseCategory, result)
	}
}

func TestCreateCategory_Failure_CategoryIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

    mockRequestCategory := &request.Category{
		Category: "",
		Type: "test",
	}

    expectedError := errors.New("category is empty")

    CategoryService := NewCategoryService(mockCategoryRepository)

    result, err := CategoryService.CreateCategory(mockRequestCategory)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.Category != "" || result.Type != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestCreateCategory_Failure_TypeIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

    mockRequestCategory := &request.Category{
		Category: "test",
		Type: "",
	}

    expectedError := errors.New("type is empty")

    CategoryService := NewCategoryService(mockCategoryRepository)

    result, err := CategoryService.CreateCategory(mockRequestCategory)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.Category != "" || result.Type != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestCreateCategory_Failure_DatabaseError(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

    mockRequestCategory := &request.Category{
        Category: "test",
		Type: "test",
    }

    expectedError := errors.New("failed to create new Category")

    mockCategoryRepository.EXPECT().CreateCategory(mockRequestCategory).Return(response.Category{}, expectedError)

    CategoryService := NewCategoryService(mockCategoryRepository)

    result, _ := CategoryService.CreateCategory(mockRequestCategory)

	if result.Id != 0 || result.Category != "" || result.Type != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestGetAllCategory_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	expectedCategories := []response.Category{
		{Id: 1, Category: "Category 1", Type: "Type 1"},
		{Id: 2, Category: "Category 2", Type: "Type 2"},
	}

	mockCategoryRepository.EXPECT().GetAllCategory("filter").Return(expectedCategories, nil).Times(1)

	categoryService := NewCategoryService(mockCategoryRepository)

	resultCategories, err := categoryService.GetAllCategory("filter")

	assert.NoError(t, err)
	assert.Equal(t, expectedCategories, resultCategories)
}

func TestGetAllCategory_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockCategoryRepository.EXPECT().GetAllCategory("filter").Return([]response.Category{}, expectedError).Times(1)

	categoryService := NewCategoryService(mockCategoryRepository)

	resultCategories, err := categoryService.GetAllCategory("filter")

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Empty(t, resultCategories)
}

func TestGetCategory_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	expectedCategory := response.Category{Id: 1, Category: "Test Category", Type: "Test Type"}

	mockCategoryRepository.EXPECT().GetCategory("CategoryID").Return(expectedCategory, nil).Times(1)

	CategoryService := NewCategoryService(mockCategoryRepository)

	resultCategory, err := CategoryService.GetCategory("CategoryID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultCategory, expectedCategory) {
		t.Errorf("Expected Category %+v, but got %+v", expectedCategory, resultCategory)
	}
}

func TestGetCategory_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	expectedError := errors.New("failed to get Category")

	mockCategoryRepository.EXPECT().GetCategory("CategoryID").Return(response.Category{}, expectedError).Times(1)

	CategoryService := NewCategoryService(mockCategoryRepository)

	_, err := CategoryService.GetCategory("CategoryID")

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error containing '%v', but got '%v'", expectedError.Error(), err)
	}
}

func TestGetCategory_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	CategoryService := NewCategoryService(mockCategoryRepository)

	_, err := CategoryService.GetCategory("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestDeleteCategory_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

    expectedResponse := response.Category{
		Id:      1,
		Category: "test",
		Type: "test",
    }

    mockCategoryRepository.EXPECT().DeleteCategory(gomock.Any()).Return(expectedResponse, nil)

    CategoryService := NewCategoryService(mockCategoryRepository)

    _, err := CategoryService.DeleteCategory("CategoryID")

    if err != nil {
        t.Errorf("Expected no error, but got %v", err)
    }
}

func TestDeleteCategory_Failure_ErrDelete(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

    expectedError := errors.New("can't delete this category")

    mockCategoryRepository.EXPECT().DeleteCategory(gomock.Any()).Return(response.Category{}, errors.New("failed to delete data from database"))

    CategoryService := NewCategoryService(mockCategoryRepository)

    result, err := CategoryService.DeleteCategory("CategoryID")

    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error %v, but got %v", expectedError, err)
    }

    if result.Id != 0 || result.Category != "" || result.Type != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeleteCategory_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	CategoryService := NewCategoryService(mockCategoryRepository)

	_, err := CategoryService.DeleteCategory("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestUpdateCategory_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	expectedUpdatedCategory := response.Category{
		Id:      	1,
		Category: 	"test",
		Type: 		"test",
	}

	mockCategoryRepository.EXPECT().UpdateCategory("CategoryID", gomock.Any()).Return(expectedUpdatedCategory, nil).Times(1)

	CategoryService := NewCategoryService(mockCategoryRepository)

	requestCategory := request.Category{
		Category: "test",
		Type: 	  "test",
	}

	resultCategory, err := CategoryService.UpdateCategory("CategoryID", requestCategory)

	if err != nil {
    	t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultCategory, expectedUpdatedCategory) {
    	t.Errorf("Expected updated Category %+v, but got %+v", expectedUpdatedCategory, resultCategory)
	}
}

func TestUpdateCategory_Failure_BadRequestID(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

    expectedError := errors.New("can't find any category with this id")

    CategoryService := NewCategoryService(mockCategoryRepository)

    _, err := CategoryService.UpdateCategory("", request.Category{})

    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }
}

func TestUpdateCategory_Failure_UpdateData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	expectedError := errors.New("failed to update data")

	mockCategoryRepository.EXPECT().UpdateCategory("CategoryID", request.Category{}).Return(response.Category{}, expectedError).Times(1)

	CategoryService := NewCategoryService(mockCategoryRepository)

	_, err := CategoryService.UpdateCategory("CategoryID", request.Category{})

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error containing '%v', but got '%v'", expectedError.Error(), err)
	}
}