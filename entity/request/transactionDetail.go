package request

type TransactionDetail struct {
	Id            uint    `json:"id"`
	TransactionId uint    `json:"transactionId"`
	ProductId     *uint   `json:"productId"`
	TicketId      *uint   `json:"ticketId"`
	Fullname      *string `json:"fullname"`
	Contact       *string `json:"contact"`
	Nik           *string `json:"nik"`
	Qty           int32   `json:"qty"`
	Product       Product `json:"product"`
	Ticket        Ticket  `json:"ticket"`
}
