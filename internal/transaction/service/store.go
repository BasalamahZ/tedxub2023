package service

import (
	"context"
	"time"

	"github.com/tedxub2023/internal/transaction"
)

// PGStore is the PostgreSQL store for configuration service.
type PGStore interface {
	NewClient(useTx bool) (PGStoreClient, error)
}

type PGStoreClient interface {
	// Commit commits the transaction.
	Commit() error

	// Rollback aborts the transaction.
	Rollback() error

	// CreateTicket creates a new ticket and return the
	// created ticket ID.
	CreateTransaction(ctx context.Context, transaction transaction.Transaction) (int64, error)

	// GetTransactionByID returns a transaction with the given
	// transaction ID.
	GetTransactionByID(ctx context.Context, transactionID int64) (transaction.Transaction, error)

	// DeleteTransactionByEmail deletes all transaction
	// with the given email and tanggal.
	DeleteTransactionByEmail(ctx context.Context, email string, tanggal time.Time) error
}
