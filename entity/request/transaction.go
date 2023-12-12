package request

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
	User              User                `json:"user"`
	Shipping          Shipping            `json:"shipping_method"`
	Payment           Payment             `json:"payment_method"`
	TransactionDetail []TransactionDetail `json:"transaction_detail"`
}
