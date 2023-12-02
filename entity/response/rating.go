package response

type Rating struct {
	Id 			uint 	`json:"id"`
	Rating		uint16	`json:"rating"`
	Ulasan		string	`json:"ulasan"`
	UserId		uint	`json:"userId"`
	ProductId 	uint	`json:"productId"`
	Users Users			`json:"user"`
	Product Products	`json:"product"`
}