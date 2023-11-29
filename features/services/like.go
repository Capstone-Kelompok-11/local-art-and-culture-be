package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type ILikeService interface {
	GetAllLike(sourceId string) ([]response.Like, error)
	UpdateLike(input request.Like) (response.Like, error)
}

type LikeService struct {
	likeRepository repositories.ILikeRepository
}

func NewLikeService(repo repositories.ILikeRepository) *LikeService {
	return &LikeService{likeRepository: repo}
}

func (co *LikeService) GetAllLike(sourceId string) ([]response.Like, error) {
	res, err := co.likeRepository.GetAllLike(sourceId)
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return res, nil
}

func (co *LikeService) UpdateLike(input request.Like) (response.Like, error) {
	if input.SourceId == 0 {
		return response.Like{}, errors.ERR_SOURCE_IS_EMPTY
	}
	if input.UserId == 0 {
		return response.Like{}, errors.ERR_GET_USER_BAD_REQUEST_ID
	}
	if input.SourceStr == "" {
		return response.Like{}, errors.ERR_SOURCE_IS_EMPTY
	}

	res, err := co.likeRepository.UpdateLike(input)
	if err != nil {
		return response.Like{}, err
	}
	return res, nil
}
