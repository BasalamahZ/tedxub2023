package service

import (
	"time"

	"github.com/midtrans/midtrans-go"
)

// New construts a new service.
type service struct {
	pgStore PGStore
	timeNow func() time.Time
}

// New returns a new service
func New(pgStore PGStore, midtransEnvironment midtrans.EnvironmentType, serverKey string) (*service, error) {
	// Set the Midtrans environment and server key
	midtrans.Environment = midtransEnvironment
	midtrans.ServerKey = serverKey

	return &service{
		pgStore: pgStore,
		timeNow: time.Now,
	}, nil
}
