package response

import "time"

type Event struct {
	Id               uint      `json:"id"`
	FromDate         time.Time `json:"fromDate"`
	ToDate           time.Time `json:"toDate"`
	EventName        string    `json:"eventName"`
	StartTime        time.Time `json:"startTime"`
	EndTime          time.Time `json:"endTime"`
	EventDescription string    `json:"eventDescription"`
	TicketPrice      float64   `json:"ticketPrice"`
	AddressId        *uint     `json:"addressId"`
	CreatorId        uint      `json:"creatorId"`
	CategoryId       uint      `json:"categoryId"`
	Creator          Creator   `json:"creator"`
	Category         Category  `json:"category"`
}