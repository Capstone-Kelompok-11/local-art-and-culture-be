package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type IRatingService interface {
	CreateRating(data *request.Rating) (response.Rating, error)
	GetAllRating(nameFilter string) ([]response.Rating, error)
	GetRating(id string) (response.Rating, error)
	UpdateRating(id string, input request.Rating) (response.Rating, error)
	DeleteRating(id string) (response.Rating, error)
}

type RatingService struct {
	ratingRepository repositories.IRatingRepository
}

func NewRatingService(repo repositories.IRatingRepository) *RatingService {
	return &RatingService{ratingRepository: repo}
}

func (ra *RatingService) CreateRating(data *request.Rating) (response.Rating, error) {
	if data.Rating == 0 {
		return response.Rating{}, errors.ERR_RATING_IS_EMPTY
	} 
	if data.Ulasan == "" {
		return response.Rating{}, errors.ERR_ULASAN_IS_EMPTY
	}

	res, err := ra.ratingRepository.CreateRating(data)
	if err != nil {
		return response.Rating{}, errors.ERR_CREATE_RATING_DATABASE
	}
	return res, nil
}

func (ra *RatingService) GetAllRating(nameFilter string) ([]response.Rating, error) {
	res, err := ra.ratingRepository.GetAllRating(nameFilter)
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return res, nil
}

func (ra *RatingService) GetRating(id string) (response.Rating, error) {
	if id == "" {
		return response.Rating{}, errors.ERR_GET_BAD_REQUEST_ID
	}
	res, err := ra.ratingRepository.GetRating(id)
	if err != nil {
		return response.Rating{}, err
	}
	return res, nil
}

func (ra *RatingService) UpdateRating(id string, input request.Rating) (response.Rating, error) {
	if id == "" {
		return response.Rating{}, errors.ERR_GET_BAD_REQUEST_ID
	}
	res, err := ra.ratingRepository.UpdateRating(id, input)
	if err != nil {
		return response.Rating{}, err
	}
	return res, nil
}

func (ra *RatingService) DeleteRating(id string) (response.Rating, error) {
	if id == "" {
		return response.Rating{}, errors.ERR_GET_BAD_REQUEST_ID
	}
	res, err := ra.ratingRepository.DeleteRating(id)
	if err != nil {
		return response.Rating{}, err
	}
	return res, nil
}