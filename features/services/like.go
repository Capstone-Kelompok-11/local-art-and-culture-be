package services

// import (
// 	"lokasani/entity/request"
// 	"lokasani/entity/response"
// 	"lokasani/features/repositories"
// 	"lokasani/helpers/errors"
// )

// type ILikeService interface {
// 	CreateLike(data *request.Like) (response.Like, error)
// 	GetAllLike() ([]response.Like, error)
// 	GetLike(id string) (response.Like, error)
// 	UpdateLike(id string, input request.Like) (response.Like, error)
// 	DeleteLike(id string) (response.Like, error)
// }

// type LikeService struct {
// 	likeRepository repositories.ILikeRepository
// }

// func NewLikeService(repo repositories.ILikeRepository) *LikeService {
// 	return &LikeService{likeRepository: repo}
// }

// func (co *LikeService) CreateLike(data *request.Like) (response.Like, error) {
// 	if data.Like == "" {
// 		return response.Like{}, errors.ERR_LIKE_IS_EMPTY
// 	}

// 	res, err := co.likeRepository.CreateLike(data)
// 	if err != nil {
// 		return response.Like{}, errors.ERR_CREATE_LIKE_DATABASE
// 	}
// 	return res, nil

// }

// func (co *LikeService) GetAllLike() ([]response.Like, error) {
// 	res, err := co.likeRepository.GetAllLike()
// 	if err != nil {
// 		return nil, errors.ERR_GET_DATA
// 	}
// 	return res, nil
// }

// func (co *LikeService) GetLiket(id string) (response.Like, error) {
// 	if id == "" {
// 		return response.Like{}, errors.ERR_GET_COMMENT_BAD_REQUEST_ID
// 	}
// 	res, err := co.likeRepository.GetLike(id)
// 	if err != nil {
// 		return response.Like{}, err
// 	}
// 	return res, nil
// }

// func (co *LikeService) UpdateLike(id string, input request.Like) (response.Like, error) {
// 	if id == "" {
// 		return response.Like{}, errors.ERR_GET_LIKE_BAD_REQUEST_ID
// 	}

// 	res, err := co.likeRepository.UpdateLike(id, input)
// 	if err != nil {
// 		return response.Like{}, err
// 	}
// 	return res, nil
// }

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
