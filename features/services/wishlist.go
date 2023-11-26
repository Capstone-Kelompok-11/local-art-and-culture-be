package services

import (
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/errors"
)

type IWishlistService interface {
	CreateWishlist(data *request.Wishlist) (response.Wishlist, error)
	GetAllWishlist(nameFilter string) ([]response.Wishlist, error)
	GetWishlist(id string) (response.Wishlist, error)
	UpdateWishlist(id string, input request.Wishlist) (response.Wishlist, error)
	DeleteWishlist(id string) (response.Wishlist, error)
}

type WishlistService struct {
	wishlistRepository repositories.IWishlistRepository
}

func NewWishlistService(repo repositories.IWishlistRepository) *WishlistService {
	return &WishlistService{wishlistRepository: repo}
}

func (wi *WishlistService) CreateWishlist(data *request.Wishlist) (response.Wishlist, error) {
	if data.Quantity == 0 {
		return response.Wishlist{}, errors.ERR_QTY_IS_EMPTY
	}

	res, err := wi.wishlistRepository.CreateWishlist(data)
	if err != nil {
		return response.Wishlist{}, errors.ERR_CREATE_WISHLIST_DATABASE
	}
	return res, nil
}

func (wi *WishlistService) GetAllWishlist(nameFilter string) ([]response.Wishlist, error) {
	res, err := wi.wishlistRepository.GetAllWishlist(nameFilter)
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}
	return res, nil
}

func (wi *WishlistService) GetWishlist(id string) (response.Wishlist, error) {
	if id == "" {
		return response.Wishlist{}, errors.ERR_GET_BAD_REQUEST_ID
	}
	res, err := wi.wishlistRepository.GetWishlist(id)
	if err != nil {
		return response.Wishlist{}, err
	}
	return res, nil
}

func (wi *WishlistService) UpdateWishlist(id string, input request.Wishlist) (response.Wishlist, error) {
	if id == "" {
		return response.Wishlist{}, errors.ERR_GET_BAD_REQUEST_ID
	}
	res, err := wi.wishlistRepository.UpdateWishlist(id, input)
	if err != nil {
		return response.Wishlist{}, err
	}
	return res, nil
}

func (wi *WishlistService) DeleteWishlist(id string) (response.Wishlist, error) {
	if id == "" {
		return response.Wishlist{}, errors.ERR_GET_BAD_REQUEST_ID
	}
	res, err := wi.wishlistRepository.DeleteWishlist(id)
	if err != nil {
		return response.Wishlist{}, err
	}
	return res, nil
}