package request

type Rating struct {
	Id 			uint 	`json:"id"`
	Rating		uint16	`json:"rating"`
	Ulasan		string	`json:"ulasan"`
	UserId		uint	`json:"user_id"`
	ProductId 	uint	`json:"product_id"`
}