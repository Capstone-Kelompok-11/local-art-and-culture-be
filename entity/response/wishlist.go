package response

type Wishlist struct {
	Id			uint	`json:"id"`
	Quantity	uint	`json:"quantity"`
	ProductId 	uint	`json:"productId"`
	TicketId 	uint	`json:"ticketId"`
	Product Products	`json:"product"`
	Ticket Ticket		`json:"ticket"`
}