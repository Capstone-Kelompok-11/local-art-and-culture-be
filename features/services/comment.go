package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type ICommentService interface {
	CreateComment(data *request.Comment) (response.Comment, error)
	GetAllComment() ([]response.Comment, error)
	GetComment(id string) (response.Comment, error)
	UpdateComment(id string, input request.Comment) (response.Comment, error)
	DeleteComment(id string) (response.Comment, error)
}

type CommentService struct {
	commentRepository repositories.ICommentRepository
}

func NewCommentService(repo repositories.ICommentRepository) *CommentService{
	return &CommentService{commentRepository: repo}
}

func (co *CommentService) CreateComment(data *request.Comment) (response.Comment, error) {
	if data.Comment == "" {
		return response.Comment{}, errors.ERR_COMMENT_IS_EMPTY
	} 

	res, err := co.commentRepository.CreateComment(data)
	if err != nil {
		return response.Comment{}, errors.ERR_CREATE_COMMENT_DATABASE
	}
	return res, nil

}

func (co *CommentService) GetAllComment() ([]response.Comment, error) {
	res, err := co.commentRepository.GetAllComment()
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return res, nil
}

func (co *CommentService) GetComment(id string) (response.Comment, error) {
	if id == "" {
		return response.Comment{}, errors.ERR_GET_COMMENT_BAD_REQUEST_ID
	}
	res, err := co.commentRepository.GetComment(id)
	if err != nil {
		return response.Comment{}, err
	}
	return res, nil
}

func (co *CommentService) UpdateComment(id string, input request.Comment) (response.Comment, error) {
	if id == "" {
		return response.Comment{}, errors.ERR_GET_COMMENT_BAD_REQUEST_ID
	}

	res, err := co.commentRepository.UpdateComment(id, input)
	if err != nil {
		return response.Comment{}, err
	}
	return res, nil
}

func (co *CommentService) DeleteComment(id string) (response.Comment, error) {
	if id == "" {
		return response.Comment{}, errors.ERR_GET_COMMENT_BAD_REQUEST_ID
	}

	res, err := co.commentRepository.DeleteComment(id)
	if err != nil {
		return response.Comment{}, err
	}
	return res, nil
}