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

func TestCreateArticle_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)

    mockRequestArticle := &request.Article{
        Title:   "test",
        Content: "test",
        Status:  "true",
    }

    expectedResponseArticle := response.Article{
        Id:      1,
        Title:   "test",
        Content: "test",
        Status:  "true",
    }

    mockArticleRepository.EXPECT().CreateArticle(mockRequestArticle).Return(expectedResponseArticle, nil)

    articleService := NewArticleService(mockArticleRepository)
    result, err := articleService.CreateArticle(mockRequestArticle)

    if err != nil {
        t.Errorf("Expected no error, but got: %v", err)
    }

    if result.Id != expectedResponseArticle.Id || result.Title != expectedResponseArticle.Title || result.Content != expectedResponseArticle.Content || result.Status != expectedResponseArticle.Status {
        t.Errorf("Expected response %v, but got %v", expectedResponseArticle, result)
    }
}

func TestCreateArticle_Failure_NameIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)

    mockRequestArticle := &request.Article{
        Title:   "",        
        Content: "test",
        Status:  "true",
    }

    expectedError := errors.New("name is empty")

    articleService := NewArticleService(mockArticleRepository)

    result, err := articleService.CreateArticle(mockRequestArticle)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.Title != "" || result.Content != "" || result.Status != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestCreateArticle_Failure_StatusIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)

    mockRequestArticle := &request.Article{
        Title:   "test",        
        Content: "test",
        Status:  "",
    }

    expectedError := errors.New("status is empty")

    articleService := NewArticleService(mockArticleRepository)

    result, err := articleService.CreateArticle(mockRequestArticle)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.Title != "" || result.Content != "" || result.Status != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestCreateArticle_Failure_DatabaseError(t *testing.T) {
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

    if result.Id != 0 || result.Title != "" || result.Content != "" || result.Status != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeleteArticle_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)

    expectedResponse := response.Article{
        Id:     1,
        Title:  "Deleted Article",
        Content: "Content of the deleted article",
        Status: "deleted",
    }

    mockArticleRepository.EXPECT().DeleteArticle(gomock.Any()).Return(expectedResponse, nil)

    articleService := NewArticleService(mockArticleRepository)

    _, err := articleService.DeleteArticle("articleID")

    if err != nil {
        t.Errorf("Expected no error, but got %v", err)
    }
}


func TestDeleteArticle_Failure_ErrDelete(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)

    expectedError := errors.New("failed to delete data from database")

    mockArticleRepository.EXPECT().DeleteArticle(gomock.Any()).Return(response.Article{}, errors.New("failed to delete data from database"))

    articleService := NewArticleService(mockArticleRepository)

    result, err := articleService.DeleteArticle("articleID")

    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error %v, but got %v", expectedError, err)
    }

    if result.Id != 0 || result.Title != "" || result.Content != "" || result.Status != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeleteArticle_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	articleService := NewArticleService(mockArticleRepository)

	_, err := articleService.DeleteArticle("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestGetAllArticle_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)

	expectedArticles := []response.Article{
		{Id: 1, Title: "Article 1", Content: "Content 1"},
		{Id: 2, Title: "Article 2", Content: "Content 2"},
	}

	expectedTotal := len(expectedArticles)
	var expectedError error

	mockArticleRepository.EXPECT().GetAllArticle("category", 0, 10).Return(expectedArticles, expectedTotal, expectedError).Times(1)

	articleService := NewArticleService(mockArticleRepository)

	resultArticles, _, err := articleService.GetAllArticle("category", 0, 10)

	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if len(resultArticles) != expectedTotal {
		t.Errorf("Expected %d articles, but got %d", expectedTotal, len(resultArticles))
	}
}

func TestGetAllArticle_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockArticleRepository.EXPECT().GetAllArticle("category", 0, 10).Return(nil, 0, expectedError).Times(1)

	articleService := NewArticleService(mockArticleRepository)

	resultArticles, resultTotal, err := articleService.GetAllArticle("category", 0, 10)

	if err.Error() != expectedError.Error(){
		t.Errorf("Expected error containing '%v', but got '%v'", expectedError.Error(), err)
	}

	if resultArticles != nil || resultTotal != 0 {
		t.Errorf("Expected empty result, but got articles: %v, total: %v", resultArticles, resultTotal)
	}
}

func TestGetArticle_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)

	expectedArticle := response.Article{Id: 1, Title: "Test Article", Content: "Test Content"}

	mockArticleRepository.EXPECT().GetArticle("articleID").Return(expectedArticle, nil).Times(1)

	articleService := NewArticleService(mockArticleRepository)

	resultArticle, err := articleService.GetArticle("articleID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultArticle, expectedArticle) {
		t.Errorf("Expected article %+v, but got %+v", expectedArticle, resultArticle)
	}
}

func TestGetArticle_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)

	expectedError := errors.New("failed to get article")

	mockArticleRepository.EXPECT().GetArticle("articleID").Return(response.Article{}, expectedError).Times(1)

	articleService := NewArticleService(mockArticleRepository)

	_, err := articleService.GetArticle("articleID")

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error containing '%v', but got '%v'", expectedError.Error(), err)
	}
}

func TestGetArticle_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	articleService := NewArticleService(mockArticleRepository)

	_, err := articleService.GetArticle("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestGetTrendingArticle_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)

	expectedArticles := []response.Article{
		{Id: 1, Title: "Article 1", Content: "Content 1"},
		{Id: 2, Title: "Article 2", Content: "Content 2"},
	}

	expectedTotal := len(expectedArticles)
	var expectedError error

	mockArticleRepository.EXPECT().GetTrendingArticle("category", 0, 10).Return(expectedArticles, expectedTotal, expectedError).Times(1)

	articleService := NewArticleService(mockArticleRepository)

	resultArticles, _, err := articleService.GetTrendingArticle("category", 0, 10)

	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if len(resultArticles) != expectedTotal {
		t.Errorf("Expected %d articles, but got %d", expectedTotal, len(resultArticles))
	}
}

func TestGetTrendingArticle_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockArticleRepository.EXPECT().GetTrendingArticle("category", 0, 10).Return(nil, 0, expectedError).Times(1)

	articleService := NewArticleService(mockArticleRepository)

	resultArticles, resultTotal, err := articleService.GetTrendingArticle("category", 0, 10)

	if err.Error() != expectedError.Error(){
		t.Errorf("Expected error containing '%v', but got '%v'", expectedError.Error(), err)
	}

	if resultArticles != nil || resultTotal != 0 {
		t.Errorf("Expected empty result, but got articles: %v, total: %v", resultArticles, resultTotal)
	}
}

func TestUpdateArticle_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)

	expectedUpdatedArticle := response.Article{
   		Id:      1,
    	Title:   "Updated Test Article",
    	Content: "Updated Test Content",
	}

	mockArticleRepository.EXPECT().UpdateArticle("articleID", gomock.Any()).Return(expectedUpdatedArticle, nil).Times(1)

	articleService := NewArticleService(mockArticleRepository)

	requestArticle := request.Article{
    	Title:   "Updated Test Article",
    	Content: "Updated Test Content",
	}

	resultArticle, err := articleService.UpdateArticle("articleID", requestArticle)

	if err != nil {
    	t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultArticle, expectedUpdatedArticle) {
    	t.Errorf("Expected updated article %+v, but got %+v", expectedUpdatedArticle, resultArticle)
	}
}

func TestUpdateArticle_Failure_BadRequestID(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)

    expectedError := errors.New("can't find any data with this id")

    articleService := NewArticleService(mockArticleRepository)

    _, err := articleService.UpdateArticle("", request.Article{})

    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }
}

func TestUpdateArticle_Failure_UpdateData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)

	expectedError := errors.New("failed to update data")

	mockArticleRepository.EXPECT().UpdateArticle("articleID", request.Article{}).Return(response.Article{}, expectedError).Times(1)

	articleService := NewArticleService(mockArticleRepository)

	_, err := articleService.UpdateArticle("articleID", request.Article{})

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error containing '%v', but got '%v'", expectedError.Error(), err)
	}
}

func TestCalculatePaginationValues(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)

    articleService := NewArticleService(mockArticleRepository)

    page1, totalPages1 := articleService.CalculatePaginationValues(0, 100, 8)
    assert.Equal(t, 1, page1)
    assert.Equal(t, 1, totalPages1)

    page2, totalPages2 := articleService.CalculatePaginationValues(15, 100, 8)
    assert.Equal(t, 1, page2)
    assert.Equal(t, 1, totalPages2)

    page3, totalPages3 := articleService.CalculatePaginationValues(2, 100, 8)
    assert.Equal(t, 1, page3)
    assert.Equal(t, 1, totalPages3)

    page4, totalPages4 := articleService.CalculatePaginationValues(1, 95, 8)
    assert.Equal(t, 1, page4)
    assert.Equal(t, 1, totalPages4) 
}

func TestGetNextPage(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)

    articleService := NewArticleService(mockArticleRepository)

    result1 := articleService.GetNextPage(3, 10)
	assert.Equal(t, 4, result1)

	result2 := articleService.GetNextPage(10, 10)
	assert.Equal(t, 10, result2)
}

func TestGetPrevPage(t *testing.T) {
	ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockArticleRepository := mocks.NewMockIArticleRepository(ctrl)

    articleService := NewArticleService(mockArticleRepository)

    result1 := articleService.GetPrevPage(5)
	assert.Equal(t, 4, result1)

	result2 := articleService.GetPrevPage(1)
	assert.Equal(t, 1, result2)
}