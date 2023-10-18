package transaction

import (
	"context"
	"time"
)

type Service interface {
	// ReplaceTransactionByEmail replace all transaction
	// with the given email
	ReplaceTransactionByEmail(ctx context.Context, transaction Transaction) (int64, error)

	// GetTransactionByID returns a transaction with the given
	// transaction ID.
	GetTransactionByID(ctx context.Context, transactionID int64, nomorTiket string) (Transaction, error)

	// UpdateCheckInStatus returns a ticket number and status with the giveb
	// transaction ID and ticket number
	UpdateCheckInStatus(ctx context.Context, id int64, nomorTiket string) (string, error)

	// UpdatePaymentStatus update the payment status in DB
	// and send a email after completed payment
	UpdatePaymentStatus(ctx context.Context, trasactionID int64) error
}

// Transaction is a transaction.

type Transaction struct {
	ID                int64
	Nama              string
	JenisKelamin      string
	NomorIdentitas    string
	AsalInstitusi     string
	Domisili          string
	Email             string
	NomorTelepon      string
	LineID            string
	Instagram         string
	JumlahTiket       int
	TotalHarga        int64
	Tanggal           time.Time
	StatusPayment     string
	OrderID           string
	ResponseMidtrans  string
	NomorTiket        []string
	CheckInStatus     bool
	CheckInNomorTiket []string
	CreateTime        time.Time
	UpdateTime        time.Time
}
