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

	total, err := pgStoreClient.CountEmail(ctx, reqTicket.Email)
	if err != nil {
		return "", err
	}

	if total > 0 {
		return "", ticket.ErrEmailAlreadyRegistered
	}

	ticketNama, err := pgStoreClient.CreateTicket(ctx, reqTicket)
	if err != nil {
		return "", err
	}

	go sendEmail(reqTicket)

	return ticketNama, nil
}

func sendEmail(reqTicket ticket.Ticket) error {
	mail := NewMailClient()
	mail.SetSender("tedxuniversitasbrawijaya@gmail.com")
	mail.SetReciever(reqTicket.Email)
	mail.SetSubject("Registrasi Panggung Swara Insan")
	if err := mail.SetBodyHTML(reqTicket.Nama); err != nil {
		return err
	}
	if err := mail.SendMail(); err != nil {
		return err
	}

	return nil
}

// validateTicket validates fields of the given Ticket
// whether its comply the predetermined rules.
func validateTicket(reqTicket ticket.Ticket) error {
	if reqTicket.Nama == "" {
		return ticket.ErrInvalidTicketNama
	}

	if reqTicket.JenisKelamin == "" || (reqTicket.JenisKelamin != "Pria" && reqTicket.JenisKelamin != "Wanita") {
		return ticket.ErrInvalidTicketJenisKelamin
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

	var womanTickets []ticket.Ticket
	var manTickets []ticket.Ticket

	for _, ticket := range tickets {
		if ticket.JenisKelamin == "Wanita" {
			womanTickets = append(womanTickets, ticket)
		} else {
			manTickets = append(manTickets, ticket)
		}
	}

	var womanCount int
	var manCount int
	if len(womanTickets) < 18 && len(womanTickets) < len(manTickets) {
		womanCount = len(womanTickets)
		manCount = 35 - len(womanTickets)
	} else if len(manTickets) < 17 && len(manTickets) < len(womanTickets) {
		manCount = len(manTickets)
		womanCount = 35 - len(manTickets)
	} else if len(womanTickets) < 18 && len(manTickets) < 17 {
		manCount = len(manTickets)
		womanCount = len(womanTickets)
	} else {
		manCount = 17
		womanCount = 18
	}

	for i := 0; i < len(womanTickets); i++ {
		status := i < womanCount
		ticketKey := fmt.Sprintf("TICKET/TEDXUB/%02d", i+1)

		req := ticket.Ticket{
			ID:         womanTickets[i].ID,
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

	for i := 0; i < len(manTickets); i++ {
		status := i < manCount
		ticketKey := fmt.Sprintf("TICKET/TEDXUB/%02d", (i + 1 + womanCount))

		req := ticket.Ticket{
			ID:         manTickets[i].ID,
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
