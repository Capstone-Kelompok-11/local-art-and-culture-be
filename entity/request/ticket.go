package request

import "time"

type Ticket struct {
	Id        uint      `json:"id"`
	Type      string    `json:"type" form:"type"`
	Price     float64   `json:"price" form:"price"`
	Qty       int       `json:"qty" form:"qty"`
	BirthDate time.Time `json:"birthday" form:"birthday"`
	EventId   uint      `json:"eventId" form:"eventId"`
}
