package response

import (
	"time"
)

type Ticket struct {
	Id          uint      `json:"id"`
	Type        string    `json:"type" form:"type"`
	Price       float64   `json:"price" form:"price"`
	Qty         int       `json:"qty" form:"qty"`
	Name        string    `json:"name" form:"name"`
	Description string    `json:"description" form:"description"`
	StartTime   time.Time `json:"start_time" form:"start_time"`
	EndTime     time.Time `json:"end_time" form:"end_time"`
	EventId     uint      `json:"event_id" form:"event_id"`
	Event       Events     `json:"event"`
}
