package ticket

import (
	"context"
	"database/sql"
	"time"
)

type Service interface {
	// CreateTicket creates a new ticket and return the
	// created ticket ID.
	CreateTicket(ctx context.Context, ticket Ticket) (string, error)

	// SendTicket sends a ticket to people who got the ticket
	// by randomize
	UpdateTicket(ctx context.Context) error
}

// Ticket is a ticket.
type Ticket struct {
	ID             int64
	Nama           string
	NomorIdentitas string
	AsalInstitusi  string
	Domisili       string
	Email          string
	NomorTelepon   string
	LineID         string
	Instagram      string
	Status         sql.NullBool
	NomorTiket     sql.NullString
	CreateTime     time.Time
	UpdateTime     sql.NullTime
}
