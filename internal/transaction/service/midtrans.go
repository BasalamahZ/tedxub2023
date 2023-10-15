package service

import (
	"encoding/json"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/tedxub2023/internal/transaction"
)

func payment(reqTransaction transaction.Transaction) (string, error) {
	midtrans.ServerKey = "SB-Mid-server-9TNsU_ksDobqWqY2ZHGsMWXA"
	midtrans.Environment = midtrans.Sandbox

	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeQris,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  reqTransaction.NomorTiket[0],
			GrossAmt: reqTransaction.Harga * int64(reqTransaction.JumlahTiket),
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    reqTransaction.NomorTiket[0],
				Name:  "Ticket Propaganda 2",
				Price: reqTransaction.Harga,
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
		return "", err
	}

	jsonData, _ := json.Marshal(coreApiRes)

	return string(jsonData), nil
}
