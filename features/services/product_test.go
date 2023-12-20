package services

import (
	"lokasani/features/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCalculatePaginationValues_Product(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockProductRepository := mocks.NewMockIProductRepository(ctrl)
	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	ProductService := NewProductService(mockProductRepository, mockCategoryRepository)


    page1, totalPages1 := ProductService.CalculatePaginationValues(0, 100, 8)
    assert.Equal(t, 1, page1)
    assert.Equal(t, 1, totalPages1)

    page2, totalPages2 := ProductService.CalculatePaginationValues(15, 100, 8)
    assert.Equal(t, 1, page2)
    assert.Equal(t, 1, totalPages2)

    page3, totalPages3 := ProductService.CalculatePaginationValues(2, 100, 8)
    assert.Equal(t, 1, page3)
    assert.Equal(t, 1, totalPages3)

    page4, totalPages4 := ProductService.CalculatePaginationValues(1, 95, 8)
    assert.Equal(t, 1, page4)
    assert.Equal(t, 1, totalPages4)
}

func TestGetNextPage_Product(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepository := mocks.NewMockIProductRepository(ctrl)
	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	ProductService := NewProductService(mockProductRepository, mockCategoryRepository)

	result1 := ProductService.GetNextPage(3, 10)
	assert.Equal(t, 4, result1)

	result2 := ProductService.GetNextPage(10, 10)
	assert.Equal(t, 10, result2)
}

func TestGetPrevPage_Product(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProductRepository := mocks.NewMockIProductRepository(ctrl)
	mockCategoryRepository := mocks.NewMockICategoryRepository(ctrl)

	ProductService := NewProductService(mockProductRepository, mockCategoryRepository)

	result1 := ProductService.GetPrevPage(5)
	assert.Equal(t, 4, result1)

	result2 := ProductService.GetPrevPage(1)
	assert.Equal(t, 1, result2)
}