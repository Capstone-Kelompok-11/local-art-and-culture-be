package request

type Product struct {
	Id          uint    `json:"id"`
	Name        string  `json:"name" form:"name"`
	Price       float64 `json:"price" form:"price"`
	Description string  `json:"description" form:"description"`
	Status      string  `json:"status" form:"status"`
	CategoryId  uint    `json:"categoryId" form:"categoryId"`
	CreatorId   uint    `json:"creatorId" form:"creatorId"`
	Category    string
	File        []byte `json:"file" form:"file"`
}
