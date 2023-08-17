package service

import (
	"context"
	"math/rand"
	"net/mail"
	"strconv"
	"strings"
	"time"

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
		return err
	}

	tickets, err := pgStoreClient.GetAllTicket(ctx)

	if err != nil {
		pgStoreClient.Rollback()
		return err
	}

	var winnerTickets []ticket.Ticket

	if len(tickets) <= 25 {
		winnerTickets = tickets
	} else {
		rand.Seed(time.Now().UnixNano())

		for i := len(tickets) - 1; i > 0; i-- {
			j := rand.Intn(i + 1)
			tickets[i], tickets[j] = tickets[j], tickets[i]
		}

		winnerTickets = tickets[:25]
	}

	for i := 0; i < len(winnerTickets); i++ {
		ticketNumber := i + 1
		ticketKey := "TICKET/TEDXUB/" + strconv.Itoa(ticketNumber)

		err := pgStoreClient.UpdateTicket(ctx, ticketKey, int(tickets[i].ID), s.timeNow())

		if err != nil {
			pgStoreClient.Rollback()
			return err
		}
	}

	err = pgStoreClient.Commit()

	if err != nil {
		return err
	}

	return nil
}
