package config

import (
	"os"

	"github.com/midtrans/midtrans-go"
)

func MidtransCredential() (ClientKey, ServerKey string) {
	midtrans.ClientKey = os.Getenv("MIDTRANS_CLIENT_KEY")
	midtrans.ServerKey = os.Getenv("MIDTRANS_SERVER_KEY")

	return midtrans.ClientKey, midtrans.ServerKey
}
