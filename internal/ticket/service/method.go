package service

import (
	"context"
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
