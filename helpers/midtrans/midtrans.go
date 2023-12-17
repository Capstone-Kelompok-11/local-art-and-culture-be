package midtrans

import (
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
	Verification(orderId, fraudStatus string) (error, response.Transaction)
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

func (m *midtransService) Verification(orderId, fraudStatus string) (error, response.Transaction) {
	var client coreapi.Client
	var transaction request.Transaction
	_, serverKey := config.MidtransCredential()
	client.New(serverKey, midtrans.Sandbox)

	transaction.Status = consts.OrderStatusPaid
	transaction.TransactionNumber = orderId
	// 4. Check transaction to Midtrans with param orderId

	if fraudStatus == "accept" {
		err, res := m.transactionRepository.UpdateTransaction("", transaction)
		if err != nil {
			return err, response.Transaction{}
		}
		return nil, res
	} else if fraudStatus == "settlement" {
		err, res := m.transactionRepository.UpdateTransaction("", transaction)
		if err != nil {
			return err, response.Transaction{}
		}
		return nil, res
	}

	return nil, response.Transaction{}
}
