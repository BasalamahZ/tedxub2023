package service

import (
	"context"

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

	// DeleteTransactionByEmail deletes all transaction
	// with the given email
	DeleteTransactionByEmail(ctx context.Context, email string) error

	// CreateTicket creates a new ticket and return the
	// created ticket ID.
	CreateTransaction(ctx context.Context, transaction transaction.Transaction) (int64, error)
}
