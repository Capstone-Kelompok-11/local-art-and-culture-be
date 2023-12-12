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
	SnapUrl           string              `json:"snap_url"`
	User              User                `json:"user"`
	Shipping          Shipping            `json:"shipping_method"`
	Payment           Payment             `json:"payment_method"`
	Total             float64             `json:"total"`
	TransactionDetail []TransactionDetail `json:"transaction_detail"`
}
