package response

type Wishlist struct {
	Id			uint	`json:"id"`
	Quantity	uint	`json:"quantity"`
	ProductId 	uint	`json:"product_id"`
	TicketId 	uint	`json:"ticket_id"`
	Product Products	`json:"product"`
	Ticket Ticket		`json:"ticket"`
}