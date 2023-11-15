package service

import (
	"context"

	"github.com/tedxub2023/internal/mainevent"
)

// PGStore is the PostgreSQL store for configuration service.
type PGStore interface {
	NewClient(useTx bool) (PGStoreClient, error)
}

type PGStoreClient interface {
	// Commit commits the mainevent.
	Commit() error

	// Rollback aborts the mainevent.
	Rollback() error

	// CreateTicket creates a new ticket and return the
	// created ticket ID.
	CreateMainEvent(ctx context.Context, mainevent mainevent.MainEvent) (int64, error)

	// GetTotalTicketByType(ctx context.Context, types int64) (int64, error)

	// GetAllMainEvents returns all mainevent.
	GetAllMainEvents(ctx context.Context, filter mainevent.GetAllMainEventsFilter) ([]mainevent.MainEvent, error)

	// GetMainEventByID returns a mainevent with the given
	// mainevent ID.
	GetMainEventByID(ctx context.Context, MainEventID int64) (mainevent.MainEvent, error)

	// UpdateMainEventByID updates a mainevent with the given
	// main event ID.
	UpdateMainEventByID(ctx context.Context, mainevent mainevent.MainEvent) error

	// DeleteMainEventByEmail deletes all mainevent
	// with the given email.
	DeleteMainEventByEmail(ctx context.Context, email string) error
}
