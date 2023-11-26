package response

import "time"

type Ticket struct {
	Id          uint      `json:"id"`
	Name        string    `json:"name" form:"name"`
	Email       string    `json:"email" form:"email"`
	Type        string    `json:"type" form:"type"`
	Price       float64   `json:"price" form:"price"`
	PhoneNumber string    `json:"phoneNumber" form:"phoneNumber"`
	Gender      string    `json:"gender" form:"gender"`
	BirthDate   time.Time `json:"birthday" form:"birthday"`
	EventId     uint      `json:"eventId" form:"eventId"`
	Event Event
}