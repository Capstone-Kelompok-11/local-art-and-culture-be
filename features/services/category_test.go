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

func TestCreateCategory_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	expectedCategory := response.Category{
		Id:   1,
		Category: "test",
		Type: "test",
	}

	mockCategoryRepository.EXPECT().CreateCategory(gomock.Any()).Return(expectedCategory, nil).Times(1)

	categoryService := NewCategoryService(mockCategoryRepository)

	requestCategory := &request.Category{
		Category: "test",
		Type: "test",
	}

	resultCategory, err := categoryService.CreateCategory(requestCategory)

	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, resultCategory)
}

func TestCreateCategory_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	expectedError := errors.New("failed to create new category")

	mockCategoryRepository.EXPECT().CreateCategory(gomock.Any()).Return(response.Category{}, expectedError).Times(1)

	categoryService := NewCategoryService(mockCategoryRepository)

	requestCategory := &request.Category{
		Category: "test",
		Type: 	  "test",
	}

	resultCategory, err := categoryService.CreateCategory(requestCategory)

	assert.EqualError(t, err, expectedError.Error())
	assert.Equal(t, response.Category{}, resultCategory)
}

func TestCreateCategory_Failure_Category(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

    mockRequestCategory := &request.Category{
		Category: "",
		Type: 	  "test",
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

func TestCreateCategory_Failure_Type(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

    mockRequestCategory := &request.Category{
		Category: "test",
		Type: 	  "",
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

func TestDeleteCategory_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	expectedCategory := response.Category{
		Id:   1,
		Category: "test",
		Type: "test",
	}

	mockCategoryRepository.EXPECT().DeleteCategory("categoryID").Return(expectedCategory, nil).Times(1)

	categoryService := NewCategoryService(mockCategoryRepository)

	resultCategory, err := categoryService.DeleteCategory("categoryID")

	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, resultCategory)
}

func TestDeleteCategory_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	expectedError := errors.New("can't delete this category")

	mockCategoryRepository.EXPECT().DeleteCategory("CategoryID").Return(response.Category{}, expectedError).Times(1)

	categoryService := NewCategoryService(mockCategoryRepository)

	resultCategory, err := categoryService.DeleteCategory("CategoryID")

	assert.EqualError(t, err, expectedError.Error())
	assert.Equal(t, response.Category{}, resultCategory)
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

func TestGetAllCategory_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	expectedCategories := []response.Category{
		{Id: 1, Category: "Category1", Type: "Type1"},
		{Id: 2, Category: "Category2", Type: "Type2"},
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
	mockCategoryRepository.EXPECT().GetAllCategory("filter").Return(nil, expectedError).Times(1)

	categoryService := NewCategoryService(mockCategoryRepository)

	resultCategories, err := categoryService.GetAllCategory("filter")

	assert.Error(t, err)
	assert.Nil(t, resultCategories)
	assert.EqualError(t, err, expectedError.Error())
}

func TestGetCategory_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	expectedCategory := response.Category{
		Id:   1,
		Category: "test",
		Type: "test",
	}

	mockCategoryRepository.EXPECT().GetCategory("CategoryID").Return(expectedCategory, nil).Times(1)

	categoryService := NewCategoryService(mockCategoryRepository)

	resultCategory, err := categoryService.GetCategory("CategoryID")

	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, resultCategory)
}

func TestGetCategory_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	expectedError := errors.New("category not found")

	mockCategoryRepository.EXPECT().GetCategory("NonExistentCategory").Return(response.Category{}, expectedError).Times(1)

	categoryService := NewCategoryService(mockCategoryRepository)

	resultCategory, err := categoryService.GetCategory("NonExistentCategory")

	assert.EqualError(t, err, expectedError.Error())
	assert.Equal(t, response.Category{}, resultCategory)
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

func TestGetTypeCategory_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	expectedCategory := response.Category{
		Id:   1,
		Category: "Technology",
		Type: "Article",
	}
	mockCategoryRepository.EXPECT().GetTypeCategory("Technology").Return(expectedCategory, nil).Times(1)

	categoryService := NewCategoryService(mockCategoryRepository)

	resultCategory, err := categoryService.GetTypeCategory("Technology")

	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, resultCategory)
}

func TestGetTypeCategory_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	expectedError := errors.New("type category not found")
	mockCategoryRepository.EXPECT().GetTypeCategory("NonExistentType").Return(response.Category{}, expectedError).Times(1)

	categoryService := NewCategoryService(mockCategoryRepository)

	resultCategory, err := categoryService.GetTypeCategory("NonExistentType")

	assert.EqualError(t, err, expectedError.Error())
	assert.Equal(t, response.Category{}, resultCategory)
}

func TestGeTypetCategory_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	CategoryService := NewCategoryService(mockCategoryRepository)

	_, err := CategoryService.GetTypeCategory("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestUpdateCategory_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	expectedCategory := response.Category{
		Id:   1,
		Category: "Update",
		Type: "Update",
	}
	mockCategoryRepository.EXPECT().UpdateCategory("categoryID", gomock.Any()).Return(expectedCategory, nil).Times(1)

	categoryService := NewCategoryService(mockCategoryRepository)
	resultCategory, err := categoryService.UpdateCategory("categoryID", request.Category{Category: "UpdatedCategory", Type: "UpdatedType"})

	assert.NoError(t, err)
	assert.Equal(t, expectedCategory, resultCategory)
}

func TestUpdateCategory_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	expectedError := errors.New("failed to update data")
	mockCategoryRepository.EXPECT().UpdateCategory("categoryID", gomock.Any()).Return(response.Category{}, expectedError).Times(1)

	categoryService := NewCategoryService(mockCategoryRepository)

	resultCategory, err := categoryService.UpdateCategory("categoryID", request.Category{Category: "UpdatedCategory", Type: "UpdatedType"})

	assert.EqualError(t, err, expectedError.Error())
	assert.Equal(t, response.Category{}, resultCategory)
}

func TestUpdateCategory_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	CategoryService := NewCategoryService(mockCategoryRepository)

	_, err := CategoryService.UpdateCategory("", request.Category{})

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}