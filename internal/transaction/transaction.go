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
	GetTransactionByID(ctx context.Context, transactionID int64) (Transaction, error)
}

// Transaction is a transaction.
type Transaction struct {
	ID               int64
	Nama             string
	JenisKelamin     string
	NomorIdentitas   string
	AsalInstitusi    string
	Domisili         string
	Email            string
	NomorTelepon     string
	LineID           string
	Instagram        string
	JumlahTiket      int
	Harga            int64
	NomorTiket       []string
	ResponseMidtrans string
	CheckInStatus    bool
	CheckInCounter   int
	CreateTime       time.Time
	UpdateTime       time.Time
}
