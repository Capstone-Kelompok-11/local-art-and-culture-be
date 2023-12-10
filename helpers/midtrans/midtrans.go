package midtrans

import (
	"fmt"
	"lokasani/app/drivers/config"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	consts "lokasani/helpers/const"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type IMidtransService interface {
	Verification(orderId string) (error, response.Transaction)
}

type midtransService struct {
	transactionRepository repositories.ITransactionRepository
}

func NewMidtransService(transactionRepository repositories.ITransactionRepository) *midtransService {
	return &midtransService{transactionRepository: transactionRepository}
}

func Payment(input *response.Transaction) error {
	_, serverKey := config.MidtransCredential()
	var s snap.Client
	s.New(serverKey, midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  input.TransactionNumber,
			GrossAmt: int64(input.Total),
		},
	}

	snapResp, _ := s.CreateTransaction(req)
	input.SnapUrl = snapResp.RedirectURL
	return nil
}

func (m midtransService) Verification(orderId string) (error, response.Transaction) {
	var client coreapi.Client
	var transaction request.Transaction
	_, serverKey := config.MidtransCredential()
	client.New(serverKey, midtrans.Sandbox)

	transaction.Status = consts.OrderStatusPaid
	transaction.TransactionNumber = orderId
	// 4. Check transaction to Midtrans with param orderId
	transactionStatusResp, e := client.CheckTransaction(orderId)
	if e != nil {
		return e, response.Transaction{}
	} else {
		if transactionStatusResp != nil {
			// 5. Do set transaction status based on response from check transaction status
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					// TODO set transaction status on your database to 'challenge'
					// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
				} else if transactionStatusResp.FraudStatus == "accept" {
					err, res := m.transactionRepository.UpdateTransaction("", transaction)
					fmt.Println(1)
					if err != nil {
						return err, response.Transaction{}
					}
					return nil, res
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				err, res := m.transactionRepository.UpdateTransaction("", transaction)
				fmt.Println(1)
				if err != nil {
					return err, response.Transaction{}
				}
				return nil, res
			} else if transactionStatusResp.TransactionStatus == "deny" {
				// TODO you can ignore 'deny', because most of the time it allows payment retries
				// and later can become success
			} else if transactionStatusResp.TransactionStatus == "cancel" || transactionStatusResp.TransactionStatus == "expire" {
				// TODO set transaction status on your databaase to 'failure'
			} else if transactionStatusResp.TransactionStatus == "pending" {
				// TODO set transaction status on your databaase to 'pending' / waiting payment
			}
		}
	}
	return nil, response.Transaction{}
}
