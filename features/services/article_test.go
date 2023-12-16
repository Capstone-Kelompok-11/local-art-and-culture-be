package services

import (
	"errors"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/mocks"
	"testing"

	"github.com/golang/mock/gomock"
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

    result, err := articleService.CreateArticle(mockRequestArticle)

    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != uint(0) || result.Title != "" || result.Content != "" || result.Status != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeleteArticle_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIArticleRepository(ctrl)

	mockRepo.EXPECT().DeleteArticle(gomock.Any()).Return(response.Article{}, nil)

	article, err := mockRepo.DeleteArticle("articleID")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if article.Id != 0 {
		t.Errorf("Expected empty article ID, got %d", article.Id)
	}
}

func TestDeleteArticle_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockIArticleRepository(ctrl)

	expectedError := errors.New("some error")
	mockRepo.EXPECT().DeleteArticle(gomock.Any()).Return(response.Article{}, expectedError)

	article, err := mockRepo.DeleteArticle("articleID")

	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}

	if article.Id != 0 {
		t.Errorf("Expected empty article ID, got %d", article.Id)
	}
}

func TestGetAllArticle_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIArticleService(ctrl)

	expectedArticles := []response.Article{{Id: 1, Title: "Article 1"}, {Id: 2, Title: "Article 2"}}
	expectedCount := 2
	mockService.EXPECT().GetAllArticle(gomock.Any(), gomock.Any(), gomock.Any()).Return(expectedArticles, expectedCount, nil)

	articles, count, err := mockService.GetAllArticle("category", 10, 1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if count != expectedCount {
		t.Errorf("Expected count %d, got %d", expectedCount, count)
	}

	if len(articles) != len(expectedArticles) {
		t.Errorf("Expected %d articles, got %d", len(expectedArticles), len(articles))
	}
}

func TestGetAllArticle_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIArticleService(ctrl)

	expectedError := errors.New("some error")
	mockService.EXPECT().GetAllArticle(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, 0, expectedError)

	articles, count, err := mockService.GetAllArticle("category", 10, 1)

	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}

	if count != 0 {
		t.Errorf("Expected count 0, got %d", count)
	}

	if articles != nil {
		t.Errorf("Expected nil articles, got %v", articles)
	}
}

func TestGetTrendingArticle_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIArticleService(ctrl)

	expectedArticles := []response.Article{{Id: 1, Title: "Article 1"}, {Id: 2, Title: "Article 2"}}
	expectedCount := 2
	mockService.EXPECT().GetTrendingArticle(gomock.Any(), gomock.Any(), gomock.Any()).Return(expectedArticles, expectedCount, nil)

	articles, count, err := mockService.GetTrendingArticle("category", 10, 1)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if count != expectedCount {
		t.Errorf("Expected count %d, got %d", expectedCount, count)
	}

	if len(articles) != len(expectedArticles) {
		t.Errorf("Expected %d articles, got %d", len(expectedArticles), len(articles))
	}
}

func TestGetTrendingArticle_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIArticleService(ctrl)

	expectedError := errors.New("some error")
	mockService.EXPECT().GetTrendingArticle(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, 0, expectedError)

	articles, count, err := mockService.GetTrendingArticle("category", 10, 1)

	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}

	if count != 0 {
		t.Errorf("Expected count 0, got %d", count)
	}

	if articles != nil {
		t.Errorf("Expected nil articles, got %v", articles)
	}
}

func TestUpdateArticle_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIArticleService(ctrl)

	expectedUpdatedArticle := response.Article{Id: 1, Title: "Updated Article"}
	mockService.EXPECT().UpdateArticle(gomock.Any(), gomock.Any()).Return(expectedUpdatedArticle, nil)

	updatedArticle, err := mockService.UpdateArticle("articleID", request.Article{Title: "New Title"})

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if updatedArticle.Id != expectedUpdatedArticle.Id {
		t.Errorf("Expected article ID %d, got %d", expectedUpdatedArticle.Id, updatedArticle.Id)
	}

	if updatedArticle.Title != expectedUpdatedArticle.Title {
		t.Errorf("Expected article title %s, got %s", expectedUpdatedArticle.Title, updatedArticle.Title)
	}
}

func TestUpdateArticle_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIArticleService(ctrl)

	expectedError := errors.New("some error")
	mockService.EXPECT().UpdateArticle(gomock.Any(), gomock.Any()).Return(response.Article{}, expectedError)

	updatedArticle, err := mockService.UpdateArticle("articleID", request.Article{Title: "New Title"})

	if err != expectedError {
		t.Errorf("Expected error %v, got %v", expectedError, err)
	}

	if updatedArticle.Id != 0 {
		t.Errorf("Expected empty article ID, got %d", updatedArticle.Id)
	}

	if updatedArticle.Title != "" {
		t.Errorf("Expected empty article title, got %s", updatedArticle.Title)
	}
}