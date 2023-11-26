package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type ILikeService interface {
	GetAllLike(articleId string) ([]response.Like, error)
	UpdateLike(input request.Like) (response.Like, error)
}

type LikeService struct {
	likeRepository repositories.ILikeRepository
}

func NewLikeService(repo repositories.ILikeRepository) *LikeService {
	return &LikeService{likeRepository: repo}
}


func (co *LikeService) GetAllLike(articleId string) ([]response.Like, error) {
	res, err := co.likeRepository.GetAllLike(articleId)
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return res, nil
}

// func (co *LikeService) GetLiket(articleId, userId string) (response.Like, error) {
// 	if id == "" {
// 		return response.Like{}, errors.ERR_GET_COMMENT_BAD_REQUEST_ID
// 	}
// 	res, err := co.likeRepository.GetLike(articleId, userId)
// 	if err != nil {
// 		return response.Like{}, err
// 	}
// 	return res, nil
// }

func (co *LikeService) UpdateLike(input request.Like) (response.Like, error) {
	if input.ArticleId == 0 {
		return response.Like{}, errors.ERR_GET_ARTICLE_BAD_REQUEST_ID
	}
	if input.UserId == 0 {
		return response.Like{}, errors.ERR_GET_USER_BAD_REQUEST_ID
	}

	res, err := co.likeRepository.UpdateLike(input)
	if err != nil {
		return response.Like{}, err
	}
	return res, nil
}

// func (co *LikeService) DeleteLike(id string) (response.Like, error) {
// 	if id == "" {
// 		return response.Like{}, errors.ERR_GET_LIKE_BAD_REQUEST_ID
// 	}

// 	res, err := co.likeRepository.DeleteLike(id)
// 	if err != nil {
// 		return response.Like{}, err
// 	}
// 	return res, nil
// }