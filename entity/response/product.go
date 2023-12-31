package response

import "time"

type Product struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name"`
	Stock 		bool	`json:"stock" form:"stock"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	TotalProduct int64	`json:"total_product" form:"total_product"`
	//AddressId	uint	`json:"addressId"``
	CategoryId uint      `json:"category_id"`
	CreatorId  uint      `json:"creator_id"`
	Creator    Creator   `json:"creator"`
	Category   Category  `json:"category"`
}

type Products struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name"`
	Stock 		bool	`json:"stock" form:"stock"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	TotalProduct int64	`json:"total_product" form:"total_product"`
	//AddressId	uint	`json:"addressId"``
	CategoryId uint     `json:"category_id"`
	CreatorId  uint     `json:"creator_id"`
	Category   Category `json:"category"`
	CreatedAt  time.Time `json:"posted_at"`
	TotalLike  uint     `json:"total_like"`
	Like       []Like   `json:"like"`
}