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

func TestCreateComment_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCommentRepository := mocks.NewMockICommentRepository(ctrl)

	expectedComment := response.Comment{
		Id:      1,
		Comment: "Test Comment",
	}

	mockCommentRepository.EXPECT().CreateComment(gomock.Any()).Return(expectedComment, nil).Times(1)

	commentService := NewCommentService(mockCommentRepository)

	requestComment := &request.Comment{
		Comment: "Test Comment",
	}

	resultComment, err := commentService.CreateComment(requestComment)

	assert.NoError(t, err)
	assert.Equal(t, expectedComment, resultComment)
}

func TestCreateComment_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCommentRepository := mocks.NewMockICommentRepository(ctrl)

	expectedError := errors.New("failed to create new comment")

	mockCommentRepository.EXPECT().CreateComment(gomock.Any()).Return(response.Comment{}, expectedError).Times(1)

	commentService := NewCommentService(mockCommentRepository)

	requestComment := &request.Comment{
		Comment: "Test Comment",
	}

	resultComment, err := commentService.CreateComment(requestComment)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, response.Comment{}, resultComment)
}

func TestCreateComment_Failure_CommentIsEmpty(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockCommentRepository := mocks.NewMockICommentRepository(ctrl)

    mockRequestComment := &request.Comment{
		Comment: "",
    }

    expectedError := errors.New("comment is empty")

    CommentService := NewCommentService(mockCommentRepository)

    result, err := CommentService.CreateComment(mockRequestComment)

    if err.Error() != expectedError.Error() {
        t.Errorf("Expected error '%v', but got '%v'", expectedError, err)
    }

    if result.Id != 0 || result.Comment != ""  {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestGetAllComment_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCommentRepository := mocks.NewMockICommentRepository(ctrl)

	expectedComments := []response.Comment{
		{Id: 1, Comment: "test 1"},
		{Id: 2, Comment: "test 2"},
	}

	var expectedError error

	mockCommentRepository.EXPECT().GetAllComment().Return(expectedComments, expectedError).Times(1)

	CommentService := NewCommentService(mockCommentRepository)

	resultComments, err := CommentService.GetAllComment()

	if err != expectedError {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultComments, expectedComments) {
		t.Errorf("Expected Comments %+v, but got %+v", expectedComments, resultComments)
	}
}

func TestGetAllComment_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCommentRepository := mocks.NewMockICommentRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockCommentRepository.EXPECT().GetAllComment().Return(nil, expectedError).Times(1)

	CommentService := NewCommentService(mockCommentRepository)

	resultComments, err := CommentService.GetAllComment()

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if resultComments != nil {
		t.Errorf("Expected nil Comments, but got %+v", resultComments)
	}
}

func TestGetComment_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCommentRepository := mocks.NewMockICommentRepository(ctrl)

	expectedComment := response.Comment{
		Id:     1,
		Comment: "test",
	}

	mockCommentRepository.EXPECT().GetComment("CommentID").Return(expectedComment, nil).Times(1)

	CommentService := NewCommentService(mockCommentRepository)

	resultComment, err := CommentService.GetComment("CommentID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultComment, expectedComment) {
		t.Errorf("Expected Comment %+v, but got %+v", expectedComment, resultComment)
	}
}

func TestGetComment_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCommentRepository := mocks.NewMockICommentRepository(ctrl)

	expectedError := errors.New("database can't request any data right now")

	mockCommentRepository.EXPECT().GetComment("CommentID").Return(response.Comment{}, expectedError).Times(1)

	CommentService := NewCommentService(mockCommentRepository)

	resultComment, err := CommentService.GetComment("CommentID")

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if resultComment.Id != 0 || resultComment.Comment != "" {
		t.Errorf("Expected empty result, but got %+v", resultComment)
	}
}

func TestGetComment_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCommentRepository := mocks.NewMockICommentRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	CommentService := NewCommentService(mockCommentRepository)

	_, err := CommentService.GetComment("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestUpdateComment_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCommentRepository := mocks.NewMockICommentRepository(ctrl)

	expectedUpdatedComment := response.Comment{
		Id:          1,
		Comment:      "Update",
	}

	var expectedError error

	mockCommentRepository.EXPECT().UpdateComment("CommentID", gomock.Any()).Return(expectedUpdatedComment, expectedError).Times(1)

	CommentService := NewCommentService(mockCommentRepository)

	requestComment := request.Comment{
		Comment:      "Update",
	}

	resultComment, err := CommentService.UpdateComment("CommentID", requestComment)

	if err != expectedError {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if !reflect.DeepEqual(resultComment, expectedUpdatedComment) {
		t.Errorf("Expected %+v, but got %+v", expectedUpdatedComment, resultComment)
	}
}

func TestUpdateComment_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCommentRepository := mocks.NewMockICommentRepository(ctrl)

	expectedError := errors.New("failed to update data")

	mockCommentRepository.EXPECT().UpdateComment("CommentID", gomock.Any()).Return(response.Comment{}, expectedError).Times(1)

	CommentService := NewCommentService(mockCommentRepository)

	requestComment := request.Comment{
		Comment:      "Update",
	}

	resultComment, err := CommentService.UpdateComment("CommentID", requestComment)

	if err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v, but got %v", expectedError, err)
	}

	if resultComment.Id != 0 || resultComment.Comment != ""{
		t.Errorf("Expected empty result, but got %+v", resultComment)
	}
}

func TestUpdateComment_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCommentRepository := mocks.NewMockICommentRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	CommentService := NewCommentService(mockCommentRepository)

	_, err := CommentService.UpdateComment("", request.Comment{})

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}

func TestDeleteComment_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCommentRepository := mocks.NewMockICommentRepository(ctrl)

	expectedComment := response.Comment{
		Id:     1,
		Comment: "test",
	}

	mockCommentRepository.EXPECT().DeleteComment("CommentID").Return(expectedComment, nil).Times(1)

	CommentService := NewCommentService(mockCommentRepository)

	resultComment, err := CommentService.DeleteComment("CommentID")

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if !reflect.DeepEqual(resultComment, expectedComment) {
		t.Errorf("Expected Comment %+v, but got %+v", expectedComment, resultComment)
	}
}

func TestDeleteComment_Failure_ErrDelete(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockCommentRepository := mocks.NewMockICommentRepository(ctrl)

    expectedError := errors.New("failed to delete data from database")

    mockCommentRepository.EXPECT().DeleteComment(gomock.Any()).Return(response.Comment{}, errors.New("failed to delete data from database"))

    CommentService := NewCommentService(mockCommentRepository)

    result, err := CommentService.DeleteComment("CommentID")

    if err == nil || err.Error() != expectedError.Error() {
        t.Errorf("Expected error %v, but got %v", expectedError, err)
    }

    if result.Id != 0 || result.Comment != "" {
        t.Errorf("Expected empty response, but got: %v", result)
    }
}

func TestDeleteComment_Failure_BadRequestID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCommentRepository := mocks.NewMockICommentRepository(ctrl)

	_ = errors.New("can't find any data with this id")

	CommentService := NewCommentService(mockCommentRepository)

	_, err := CommentService.DeleteComment("")

	if err == nil  {
		t.Errorf("Expected error '%v', but got '%v'", "", err)
	}
}