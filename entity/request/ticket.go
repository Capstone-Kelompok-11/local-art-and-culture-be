package request

import "time"

type Ticket struct {
	Id          uint      `json:"id"`
	Type        string    `json:"type" form:"type"`
	Price       float64   `json:"price" form:"price"`
	Qty         int       `json:"qty" form:"qty"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	StartTime   time.Time `json:"startTime" form:"startTime"`
	EndTime     time.Time `json:"endTime" form:"endTime"`
	EventId     uint      `json:"eventId" form:"eventId"`
}
