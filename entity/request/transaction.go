package request

import (
	"time"
)

type Transaction struct {
	Id                uint                `json:"id"`
	TransactionDate   time.Time           `json:"transactionDate"`
	UserId            uint                `json:"userId"`
	PaymentMethodId   uint                `json:"paymentMethodId"`
	ShippingMethodId  *uint               `json:"shippingMethodId"`
	Status            string              `json:"status"`
	User              User                `json:"user"`
	Shipping          Shipping            `json:"shippingMethod"`
	Payment           Payment             `json:"paymentMethod"`
	TransactionDetail []TransactionDetail `json:"transactionDetail"`
}
