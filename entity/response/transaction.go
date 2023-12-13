package response

import (
	"time"
)

type Transaction struct {
	Id                uint                `json:"id"`
	TransactionDate   time.Time           `json:"transaction_date"`
	UserId            uint                `json:"user_id"`
	PaymentMethodId   uint                `json:"payment_method_id"`
	TransactionNumber string              `json:"transaction_number"`
	ShippingMethodId  *uint               `json:"shipping_method_id"`
	Status            string              `json:"status"`
	SnapUrl           string              `json:"snap_url"`
	User              User                `json:"user"`
	Shipping          Shipping            `json:"shipping_method"`
	Payment           Payment             `json:"payment_method"`
	Total             float64             `json:"total"`
	TransactionDetail []TransactionDetail `json:"transaction_detail"`
}

type TransactionReport struct {
	Id              uint		`json:"id"`
	TransactionDate time.Time	`json:"transaction_date"`
	Status          string		`json:"status"`
	Qty             int32		`json:"qty"`
	Price           float64		`json:"price"`
	Nominal         float64		`json:"nominal"`
}