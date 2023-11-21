package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/mail"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/tedxub2023/global/helper"
	"github.com/tedxub2023/internal/mainevent"
)

func (s *service) ReplaceMainEventByEmail(ctx context.Context, reqMainEvent mainevent.MainEvent) (int64, error) {
	// validate field
	err := validateMainEvent(reqMainEvent)
	if err != nil {
		return 0, err
	}

	// get pg store client
	pgStoreClient, err := s.pgStore.NewClient(true)
	if err != nil {
		return 0, err
	}

	tickets, err := pgStoreClient.GetAllMainEvents(ctx, mainevent.GetAllMainEventsFilter{})
	if err != nil {
		return 0, err
	}

	totalPresaleTicket := checkPresaleTicket(tickets)
	if totalPresaleTicket+reqMainEvent.JumlahTiket > 35 {
		return 0, mainevent.ErrEarlyBirdTicketSoldOut
	}

	reqMainEvent.CreateTime = s.timeNow()
	reqMainEvent.TotalHarga = 49000 * int64(reqMainEvent.JumlahTiket)

	// rollback just before return if error
	defer func() {
		if err != nil {
			errTx := pgStoreClient.Rollback()
			if errTx != nil {
				// return err from rollback
				err = errTx
			}
		}
	}()

	// delete MainEvent specified with email in pgstore
	err = pgStoreClient.DeleteMainEventByEmail(ctx, reqMainEvent.Email)
	if err != nil {
		return 0, err
	}

	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1e10)
	reqMainEvent.OrderID = fmt.Sprintf("%010d", randomNum)

	ticketID, err := pgStoreClient.CreateMainEvent(ctx, reqMainEvent)
	if err != nil {
		return 0, err
	}

	go sendMainEventPendingMail(reqMainEvent)

	// commit changes
	err = pgStoreClient.Commit()
	if err != nil {
		return 0, err
	}
	s.AddCronJobs(ctx)

	return ticketID, nil
}

func checkPresaleTicket(tx []mainevent.MainEvent) int {
	counter := 0
	for _, t := range tx {
		if t.Type == mainevent.TypePresale {
			counter += t.JumlahTiket
		}
	}
	return counter
}

func (s *service) GetAllMainEvents(ctx context.Context, filter mainevent.GetAllMainEventsFilter) ([]mainevent.MainEvent, error) {
	// get pg store client without using MainEvent
	pgStoreClient, err := s.pgStore.NewClient(false)
	if err != nil {
		return nil, err
	}

	// get all MainEvents from postgre
	result, err := pgStoreClient.GetAllMainEvents(ctx, filter)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *service) GetMainEventByID(ctx context.Context, MainEventID int64, nomorTiket string) (mainevent.MainEvent, error) {
	// validate id
	if MainEventID <= 0 {
		return mainevent.MainEvent{}, mainevent.ErrInvalidMainEventID
	}

	// get pg store client without using MainEvent
	pgStoreClient, err := s.pgStore.NewClient(false)
	if err != nil {
		return mainevent.MainEvent{}, err
	}

	// get user from pgstore
	result, err := pgStoreClient.GetMainEventByID(ctx, MainEventID)
	if err != nil {
		return mainevent.MainEvent{}, err
	}

	if nomorTiket != "" {
		valid := false
		for _, nomorTikett := range result.NomorTiket {
			if nomorTikett == nomorTiket {
				valid = true
				break
			}
		}

		if !valid {
			return mainevent.MainEvent{}, mainevent.ErrDataNotFound
		}
	}

	return result, nil
}

func (s *service) UpdateCheckInStatus(ctx context.Context, id int64, ticketNumber string) (string, error) {
	if id <= 0 {
		return "", mainevent.ErrInvalidMainEventID
	}

	pgStoreClient, err := s.pgStore.NewClient(false)
	if err != nil {
		return "", err
	}

	tx, err := pgStoreClient.GetMainEventByID(ctx, id)
	if err != nil {
		return "", err
	}

	if tx.CheckInStatus {
		return "", mainevent.ErrAllTicketAlreadyCheckedIn
	} else if tx.Status != mainevent.StatusSettlement {
		return "", mainevent.ErrTicketNotYetPaid
	}

	var isTicketAvailable bool

	for _, ticNum := range tx.NomorTiket {
		if ticketNumber == ticNum {
			isTicketAvailable = true
			break
		}
	}

	if isTicketAvailable {
		for _, ticNum := range tx.CheckInNomorTiket {
			if ticketNumber == ticNum {
				return "", mainevent.ErrTicketAlreadyCheckedIn
			}
		}
	} else {
		return "", mainevent.ErrDataNotFound
	}

	tx.CheckInNomorTiket = append(tx.CheckInNomorTiket, ticketNumber)

	if len(tx.CheckInNomorTiket) == len(tx.NomorTiket) {
		tx.CheckInStatus = true
	}

	err = pgStoreClient.UpdateMainEventByID(ctx, tx)
	if err != nil {
		return "", err
	}

	return ticketNumber, nil
}

func (s *service) UpdatePaymentStatus(ctx context.Context, reqMainEvent mainevent.MainEvent) error {
	// validate id
	if reqMainEvent.ID <= 0 {
		return mainevent.ErrInvalidMainEventID
	}

	reqMainEvent.UpdateTime = s.timeNow()

	pgStoreClient, err := s.pgStore.NewClient(false)
	if err != nil {
		return err
	}

	ticketNumbers := generateNumberTicket(reqMainEvent.ID, reqMainEvent.JumlahTiket)
	reqMainEvent.NomorTiket = ticketNumbers

	err = pgStoreClient.UpdateMainEventByID(ctx, reqMainEvent)
	if err != nil {
		return err
	}

	if reqMainEvent.Status == mainevent.StatusPending && reqMainEvent.ImageURI != "" {
		go sendMailInformAdmin(reqMainEvent)
	}

	if reqMainEvent.Status == mainevent.StatusSettlement {
		if err := generatePDF(reqMainEvent); err != nil {
			return err
		}

		go sendSuccessTransactionMail(reqMainEvent)
	}

	return nil
}

func generateNumberTicket(txID int64, totalTickets int) []string {
	var ticketNumbers []string

	asciiVal := 65
	for i := 0; i < totalTickets; i++ {
		ticketNumbers = append(ticketNumbers, fmt.Sprintf("MEMANTIKBASKARA-1/%c%d", rune(asciiVal), txID))
		asciiVal++
	}

	return ticketNumbers
}

func generatePDF(tx mainevent.MainEvent) error {
	err := helper.PDF(tx)
	if err != nil {
		return err
	}
	return nil
}

// validateMainEvent validates fields of the given MainEvent
// whether its comply the predetermined rules.
func validateMainEvent(reqMainEvent mainevent.MainEvent) error {
	if reqMainEvent.Nama == "" {
		return mainevent.ErrInvalidMainEventNama
	}

	if reqMainEvent.NomorIdentitas == "" || len(reqMainEvent.NomorIdentitas) < 15 {
		return mainevent.ErrInvalidMainEventNomorIdentitas
	}

	if reqMainEvent.AsalInstitusi == "" {
		return mainevent.ErrInvalidMainEventAsalInstitusi
	}

	_, err := mail.ParseAddress(reqMainEvent.Email)
	if reqMainEvent.Email == "" || err != nil {
		return mainevent.ErrInvalidMainEventEmail
	}

	if reqMainEvent.NomorTelepon == "" || len(reqMainEvent.NomorTelepon) < 10 || len(reqMainEvent.NomorTelepon) > 13 {
		return mainevent.ErrInvalidMainEventNomorTelepon
	}

	if reqMainEvent.JumlahTiket <= 0 {
		return mainevent.ErrInvalidMainEventJumlahTiket
	}

	return nil
}

func (s *service) AddCronJobs(ctx context.Context) error {
	pgStoreClient, err := s.pgStore.NewClient(false)
	if err != nil {
		return err
	}
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))
	go scheduler.AddFunc("*/1 * * * *", func() {
		results, err := s.GetAllMainEvents(ctx, mainevent.GetAllMainEventsFilter{
			Status: mainevent.StatusUnpaid,
			Type:   mainevent.TypePresale,
		})
		if err != nil {
			log.Println("err get all", err)
			return
		}

		for _, result := range results {
			timeDifference := time.Since(result.CreateTime)
			if timeDifference.Minutes() > 6 {
				err = pgStoreClient.DeleteMainEventByEmail(ctx, result.Email)
				if err != nil {
					log.Println("err delete", err, result.ID)
				}
				if err == nil {
					log.Println("deleted", result.ID)

					go sendTransactionDeclinedMail(result)
				}
			}
		}

	})
	scheduler.Start()

	return nil
}
