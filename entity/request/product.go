package request

type Product struct {
	Id			uint	`json:"id"`
	Name		string	`json:"name"`
	Price		float64	`json:"price"`
	Description	string	`json:"description"`
	Status		string	`json:"status"`
	//AddressId	uint	`json:"addressId"``
	CategoryId	uint	`json:"categoryId"`
	CreatorId	uint	`json:"creatorId"`
}