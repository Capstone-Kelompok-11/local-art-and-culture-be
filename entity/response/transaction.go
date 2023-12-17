package response

import (
	"time"
)

type Transaction struct {
	Id                uint                `json:"id"`
	TransactionDate   time.Time           `json:"transaction_date"`
	UserId            uint                `json:"user_id"`
	PaymentMethodId   uint                `json:"payment_method_id"`
	ShippingMethodId  *uint               `json:"shipping_method_id"`
	Status            string              `json:"status"`
	TransactionNumber string              `json:"transaction_number"`
	SnapUrl           string              `json:"snap_url"`
	User              User                `json:"user"`
	Shipping          Shipping            `json:"shipping_method"`
	Payment           Payment             `json:"payment_method"`
	Total             int64               `json:"total"`
	TransactionDetail []TransactionDetail `json:"transaction_detail"`
}

type TransactionReport struct {
	Id              uint      `json:"id"`
	TransactionDate time.Time `json:"transaction_date"`
	Status          string    `json:"status"`
	Qty             int32     `json:"qty"`
	Price           float64   `json:"price"`
	Nominal         float64   `json:"nominal"`
}
