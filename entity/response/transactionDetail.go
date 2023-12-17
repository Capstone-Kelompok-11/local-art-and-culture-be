package response

type TransactionDetail struct {
	Id            uint    `json:"id"`
	TransactionId uint    `json:"transaction_id"`
	ProductId     *uint   `json:"product_id"`
	TicketId      *uint   `json:"ticket_id"`
	Fullname      *string `json:"fullname"`
	Contact       *string `json:"contact"`
	Nik           *string `json:"nik"`
	Email         *string `json:"email"`
	Qty           int32   `json:"qty"`
	Subtotal      int64   `json:"subtotal"`
	Product       Product `json:"product"`
	Ticket        Ticket  `json:"ticket"`
}
