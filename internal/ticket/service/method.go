package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/mail"
	"strings"

	"github.com/tedxub2023/internal/ticket"
)

func (s *service) CreateTicket(ctx context.Context, reqTicket ticket.Ticket) (string, error) {
	// validate field
	err := validateTicket(reqTicket)
	if err != nil {
		return "", err
	}

	// these value should be same for all users
	var (
		createTime = s.timeNow()
	)

	reqTicket.CreateTime = createTime

	// get pg store client without using transaction
	pgStoreClient, err := s.pgStore.NewClient(false)
	if err != nil {
		return "", err
	}

	ticketNama, err := pgStoreClient.CreateTicket(ctx, reqTicket)
	if err != nil {
		return "", err
	}

	_ = NewMailClient()

	/*
		Implement here
		need email and name
	*/

	return ticketNama, nil
}

// validateTicket validates fields of the given Ticket
// whether its comply the predetermined rules.
func validateTicket(reqTicket ticket.Ticket) error {
	if reqTicket.Nama == "" {
		return ticket.ErrInvalidTicketNama
	}

	if reqTicket.NomorIdentitas == "" || len(reqTicket.NomorIdentitas) < 15 {
		return ticket.ErrInvalidTicketNomorIdentitas
	}

	if reqTicket.AsalInstitusi == "" {
		return ticket.ErrInvalidTicketAsalInstitusi
	}

	if reqTicket.Domisili == "" {
		return ticket.ErrInvalidTicketDomisili
	}

	_, err := mail.ParseAddress(reqTicket.Email)
	if reqTicket.Email == "" || err != nil {
		return ticket.ErrInvalidTicketEmail
	}

	if reqTicket.NomorTelepon == "" || len(reqTicket.NomorTelepon) < 10 || len(reqTicket.NomorTelepon) > 13 {
		return ticket.ErrInvalidTicketNomorTelepon
	}

	if reqTicket.LineID == "" || strings.HasPrefix(reqTicket.LineID, "@") {
		return ticket.ErrInvalidTicketLineID
	}

	if reqTicket.Instagram == "" || strings.HasPrefix(reqTicket.LineID, "@") {
		return ticket.ErrInvalidTicketInstagram
	}

	return nil
}

func (s *service) UpdateTicket(ctx context.Context) error {
	pgStoreClient, err := s.pgStore.NewClient(true)
	if err != nil {
		return ticket.ErrFailedTransaction
	}

	defer func() error {
		if err != nil {
			pgStoreClient.Rollback()
			log.Printf("[tedxub2023-api-service] Service got rollback, error occured: %s\n", err.Error())
			return err
		}
		return nil
	}()

	tickets, err := pgStoreClient.GetAllTicket(ctx)
	if err != nil {
		return err
	}

	if len(tickets) == 0 {
		return ticket.ErrTicketNotFound
	}

	rand.Shuffle(len(tickets), func(i, j int) {
		tickets[i], tickets[j] = tickets[j], tickets[i]
	})

	for i := 0; i < len(tickets); i++ {
		status := i < 25
		ticketKey := fmt.Sprintf("TICKET/TEDXUB/%02d", i+1)

		req := ticket.Ticket{
			ID:         tickets[i].ID,
			Status:     status,
			NomorTiket: ticketKey,
			UpdateTime: s.timeNow(),
		}

		if status {
			if err := pgStoreClient.UpdateTicket(ctx, req); err != nil {
				return err
			}
		} else {
			break
		}
	}

	if err = pgStoreClient.Commit(); err != nil {
		log.Printf("[tedxub2023-api-service] Failed to commit the transaction.: %s\n", err.Error())
		return err
	}
	return nil
}
