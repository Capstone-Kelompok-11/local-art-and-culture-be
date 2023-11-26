package response

import "time"

type Event struct {
	Id               uint      `json:"id"`
	FromDate         time.Time `json:"fromDate"`
	ToDate           time.Time `json:"toDate"`
	EventName        string    `json:"eventName"`
	EventDescription string    `json:"eventDescription"`
	AddressId        *uint     `json:"addressId"`
	CreatorId        uint      `json:"creatorId"`
	CategoryId       uint      `json:"categoryId"`
	Creator          Creator   `json:"creator"`
	Category         Category  `json:"category"`
}

type Events struct {
	Id               uint      `json:"id"`
	FromDate         time.Time `json:"fromDate"`
	ToDate           time.Time `json:"toDate"`
	EventName        string    `json:"eventName"`
	EventDescription string    `json:"eventDescription"`
	AddressId        *uint     `json:"addressId"`
	CreatorId        uint      `json:"creatorId"`
	CategoryId       uint      `json:"categoryId"`
}
