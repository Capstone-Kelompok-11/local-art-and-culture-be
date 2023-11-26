package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type IWishlistRepository interface {
	CreateWishlist(data *request.Wishlist) (response.Wishlist, error)
	GetAllWishlist(nameFilter string) ([]response.Wishlist, error)
	GetWishlist(id string) (response.Wishlist, error)
	UpdateWishlist(id string, input request.Wishlist) (response.Wishlist, error)
	DeleteWishlist(id string) (response.Wishlist, error)
}

type wishlistRepository struct {
	db *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) *wishlistRepository {
	return &wishlistRepository{db}
}

func (wi *wishlistRepository) CreateWishlist(data *request.Wishlist) (response.Wishlist, error) {
	dataWishlist := domain.ConvertFromWishlistReqToModel(*data)
	err := wi.db.Create(&dataWishlist).Error
	if err != nil {
		return response.Wishlist{}, err
	}
	err = wi.db.Preload("Product").Preload("Product.Category").Preload("Ticket").Preload("Ticket.Event").First(&dataWishlist, "id = ?", dataWishlist.ID).Error
	return *domain.ConvertFromModelToWishlistRes(*dataWishlist), nil
}

func (wi *wishlistRepository) GetAllWishlist(nameFilter string) ([]response.Wishlist, error) {
	var allWishlist []models.Wishlist
	var resAllWishlist []response.Wishlist

	query := wi.db.Preload("Product").Preload("Product.Category").Preload("Ticket").Preload("Ticket.Event")

	if nameFilter != "" {
		query = query.Where("quantity LIKE ?", "%"+nameFilter+"%")
	}

	err := query.Find(&allWishlist).Error
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}

	for i := 0; i < len(allWishlist); i++ {
		WishlistVm := domain.ConvertFromModelToWishlistRes(allWishlist[i])
		resAllWishlist = append(resAllWishlist, *WishlistVm)
	}

	return resAllWishlist, nil
}

func (wi *wishlistRepository) GetWishlist(id string) (response.Wishlist, error) {
	var wishlistData models.Wishlist
	err := wi.db.Preload("Product").Preload("Product.Category").Preload("Ticket").Preload("Ticket.Event").First(&wishlistData, "id = ?", id).Error

	if err != nil {
		return response.Wishlist{}, err
	}
	return *domain.ConvertFromModelToWishlistRes(wishlistData), nil
}

func (wi *wishlistRepository) UpdateWishlist(id string, input request.Wishlist) (response.Wishlist, error) {
	wishlistData := models.Wishlist{}
	err := wi.db.First(&wishlistData, "id = ?", id).Error
	if err != nil {
		return response.Wishlist{}, err
	}

	if input.Quantity != 0 {
		wishlistData.Quantity = input.Quantity
	}
	if input.ProductId != 0 {
		wishlistData.ProductId = input.ProductId
	}
	if input.TicketId != 0 {
		wishlistData.TicketId = input.TicketId
	}

	if err = wi.db.Save(&wishlistData).Error; err != nil {
		return response.Wishlist{}, err
	}
	return *domain.ConvertFromModelToWishlistRes(wishlistData), nil
}

func (wi *wishlistRepository) DeleteWishlist(id string) (response.Wishlist, error) {
	wishlistData := models.Wishlist{}
	res := response.Wishlist{}
	find := wi.db.Preload("Product").Preload("Product.Category").Preload("Ticket").Preload("Ticket.Event").First(&wishlistData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToWishlistRes(wishlistData)
	}

	err := wi.db.Delete(&wishlistData, "id = ?", id).Error
	if err != nil {
		return response.Wishlist{}, err
	}
	return res, nil
}