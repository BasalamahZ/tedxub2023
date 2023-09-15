package service

import (
	"context"

	"github.com/tedxub2023/internal/ticket"
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
	CreateTicket(ctx context.Context, ticket ticket.Ticket) (string, error)

	GetAllTicket(ctx context.Context) ([]ticket.Ticket, error)

	UpdateTicket(ctx context.Context, t ticket.Ticket) error

	CountEmail(ctx context.Context, email string) (int, error)
}
