package service

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/tedxub2023/internal/transaction"
)

func (s *service) payment(reqTransaction transaction.Transaction) (*coreapi.ChargeResponse, error) {
	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeQris,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  reqTransaction.OrderID,
			GrossAmt: reqTransaction.TotalHarga,
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    "1",
				Name:  "Ticket Propaganda 2",
				Price: 25000,
				Qty:   int32(reqTransaction.JumlahTiket),
			},
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: reqTransaction.Nama,
			Email: reqTransaction.Email,
			Phone: reqTransaction.NomorTelepon,
		},
		CustomExpiry: &coreapi.CustomExpiry{
			ExpiryDuration: 30,
			Unit:           "day",
		},
	}

	coreApiRes, err := coreapi.ChargeTransaction(chargeReq)
	if err != nil {
		return nil, err
	}

	return coreApiRes, nil
}

func (s *service) checkStatusPayment(orderID string) (*coreapi.TransactionStatusResponse, error) {
	res, err := coreapi.CheckTransaction(orderID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
