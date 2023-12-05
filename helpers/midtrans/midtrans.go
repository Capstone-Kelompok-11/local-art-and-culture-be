package midtrans

import (
	"lokasani/app/drivers/config"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/features/repositories"
	"lokasani/helpers/consts"
	"lokasani/helpers/errors"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type IMidtransService interface {
	Verification(data map[string]interface{}) (error, *string)
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
			OrderID:  "LKSNI/" + string(input.Id),
			GrossAmt: 100000,
		},
	}

	snapResp, _ := s.CreateTransaction(req)
	input.SnapUrl = snapResp.RedirectURL
	return nil
}

func (m midtransService) Verification(data map[string]interface{}) (error, *string) {
	var client coreapi.Client
	var transaction request.Transaction
	_, serverKey := config.MidtransCredential()
	client.New(serverKey, midtrans.Sandbox)

	orderId, exists := data["order_id"].(string)
	if !exists {
		// do something when key `order_id` not found
		return errors.ERR_INVALID_PAYLOAD, nil
	}

	// 4. Check transaction to Midtrans with param orderId
	transactionStatusResp, e := client.CheckTransaction(orderId)
	if e != nil {
		return e, nil
	} else {
		if transactionStatusResp != nil {
			// 5. Do set transaction status based on response from check transaction status
			if transactionStatusResp.TransactionStatus == "capture" {
				if transactionStatusResp.FraudStatus == "challenge" {
					// TODO set transaction status on your database to 'challenge'
					// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
				} else if transactionStatusResp.FraudStatus == "accept" {
					transaction.Status = consts.OrderStatusPaid
					m.transactionRepository.UpdateTransaction(orderId, transaction)
				}
			} else if transactionStatusResp.TransactionStatus == "settlement" {
				transaction.Status = consts.OrderStatusPaid
				m.transactionRepository.UpdateTransaction(orderId, transaction)
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
	return nil, nil
}
