package response

type Product struct {
	Id			uint	`json:"id"`
	Name		string	`json:"name"`
	Price		string	`json:"price"`
	Description	string	`json:"description"`
	Status		string	`json:"status"`
	//AddressId	uint	`json:"addressId"``
	CategoryId	uint	`json:"categoryId"`
	CreatorId	uint	`json:"creatorId"`
}