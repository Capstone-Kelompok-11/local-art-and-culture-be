package request

type Product struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name" form:"name"`
	Stock 		bool	`json:"stock" form:"stock"`
	Price       float64 `json:"price" form:"price"`
	Description string  `json:"description" form:"description"`
	Status      string  `json:"status" form:"status"`
	TotalProduct int64	`json:"total_product" form:"total_product"`
	CategoryId  uint    `json:"category_id" form:"category_id"`
	CreatorId   uint    `json:"creator_id" form:"creator_id"`
	Category    string  `json:"category"`
	File        []byte  `json:"file" form:"file"`
}
