package request

type Product struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name" form:"name"`
	Price       float64 `json:"price" form:"price"`
	Description string  `json:"description" form:"description"`
	Status      string  `json:"status" form:"status"`
	CategoryId  uint    `json:"category_id" form:"category_id"`
	CreatorId   uint    `json:"creator_id" form:"creator_id"`
	Category    string  `json:"category"`
	File        []byte  `json:"file" form:"file"`
}
