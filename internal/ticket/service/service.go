package service

import "time"

// New construts a new service.
type service struct {
	pgStore PGStore
	timeNow func() time.Time
}

// New returns a new service
func New(pgStore PGStore) (*service, error) {
	return &service{
		pgStore: pgStore,
		timeNow: time.Now,
	}, nil
}

type TicketUpdate struct {
	ID         int64
	Status     bool
	NomorTiket string
	UpdateTime time.Time
}
