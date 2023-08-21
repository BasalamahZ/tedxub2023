package ticket

import (
	"context"
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
	Status         bool
	NomorTiket     string
	CreateTime     time.Time
	UpdateTime     time.Time
}
