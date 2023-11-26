package request

type Wishlist struct {
	Id			uint	`json:"id"`
	Quantity	uint	`json:"quantity"`
	ProductId 	uint	`json:"productId"`
	TicketId 	uint	`json:"ticketId"`
}