package response

import "time"

type Event struct {
	Id               uint      `json:"id"`
	FromDate         time.Time `json:"from_date"`
	ToDate           time.Time `json:"to_date"`
	EventName        string    `json:"event_name"`
	EventDescription string    `json:"event_description"`
	AddressId        *uint     `json:"address_id"`
	CreatorId        uint      `json:"creator_id"`
	CategoryId       uint      `json:"category_id"`
	Creator          Creator   `json:"creator"`
	Category         Category  `json:"category"`
	Guest			 []Guest   `json:"guest"`
}

type Events struct {
	Id               uint      `json:"id"`
	FromDate         time.Time `json:"from_date"`
	ToDate           time.Time `json:"to_date"`
	EventName        string    `json:"event_name"`
	EventDescription string    `json:"event_description"`
	AddressId        *uint     `json:"address_id"`
	CreatorId        uint      `json:"creator_id"`
	CategoryId       uint      `json:"category_id"`
}
