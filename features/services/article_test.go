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

func TestCreateArticle_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)

    mockRequestArticle := &request.Article{
        Title:   "test",
        Content: "test",
		Status: "true",
    }

    expectedResponseArticle := response.Article{
        Title:   "test",
        Content: "test",
		Status: "true",
    }

    mockArticleRepository.EXPECT().CreateArticle(mockRequestArticle).Return(expectedResponseArticle, nil)

    articleService := NewArticleService(mockArticleRepository)
    result, err := articleService.CreateArticle(mockRequestArticle)

    if err != nil {
        t.Errorf("Expected no error, but got: %v", err)
    }

    if result.Title != expectedResponseArticle.Title || result.Content != expectedResponseArticle.Content {
    	t.Errorf("Expected response %v, but got %v", expectedResponseArticle, result)
	}
}

func TestCreateArticle_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)

    mockRequestArticle := &request.Article{
        Title:   "test",
        Content: "test",
        Status:  "true",
    }

    expectedError := errors.New("failed to create new article") 

    mockArticleRepository.EXPECT().CreateArticle(mockRequestArticle).Return(response.Article{}, expectedError)

    articleService := NewArticleService(mockArticleRepository)
    result, _ := articleService.CreateArticle(mockRequestArticle)

    if result.Title != "" || result.Content != "" || result.Status != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeleteArticleSuccess(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)
    mockArticleService := mocks.NewMockIArticleService(ctrl)

    expectedArticle := response.Article{
        Title:   "test",
        Content: "test",
        Status:  "true",
    }

    mockArticleRepository.EXPECT().DeleteArticle(gomock.Any()).Return(expectedArticle, nil)
    gomock.InOrder(
        mockArticleService.EXPECT().DeleteArticle(gomock.Any()).AnyTimes(),
    )

    article, err := mockArticleService.DeleteArticle("articleID")

    assert.Nil(t, err)
    assert.Equal(t, expectedArticle, article)
}