package domain

import (
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"

	"github.com/jinzhu/gorm"
)

func ConvertFromWishlistReqToModel(data request.Wishlist) *models.Wishlist {
	return &models.Wishlist{
		Model: gorm.Model{
			ID: data.Id,
		},
		Quantity: 	data.Quantity,
		ProductId: 	data.ProductId,
		TicketId: 	data.TicketId,
	}
}

func ConvertFromModelToWishlistRes(data models.Wishlist) *response.Wishlist {
	return &response.Wishlist{
		Id: 		data.ID,
		Quantity: 	data.Quantity,
		ProductId: 	data.ProductId,
		TicketId: 	data.TicketId,
		Product: 	*ConvertFromModelToProductsRes(data.Product),
		Ticket: 	*ConvertFromModelToTicketRes(data.Ticket),
	}
}