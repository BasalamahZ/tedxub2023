package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/mail"
	"strings"
	"time"

	"github.com/tedxub2023/internal/transaction"
)

func (s *service) ReplaceTransactionByEmail(ctx context.Context, reqTransaction transaction.Transaction) (int64, error) {
	// validate field
	err := validateTransaction(reqTransaction)
	if err != nil {
		return 0, err
	}
	reqTransaction.CreateTime = s.timeNow()
	reqTransaction.TotalHarga = 30000 * int64(reqTransaction.JumlahTiket)

	// get pg store client with using transaction
	pgStoreClient, err := s.pgStore.NewClient(true)
	if err != nil {
		return 0, err
	}

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

	// delete transaction specified with email and tanggal in pgstore
	err = pgStoreClient.DeleteTransactionByEmail(ctx, reqTransaction.Email, reqTransaction.Tanggal)
	if err != nil {
		return 0, err
	}

	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(1e10)
	reqTransaction.OrderID = fmt.Sprintf("%010d", randomNum)

	resPayment, err := s.payment(reqTransaction)
	if err != nil {
		return 0, err
	}

	reqTransaction.StatusPayment = resPayment.TransactionStatus

	jsonData, err := json.Marshal(resPayment)
	if err != nil {
		return 0, err
	}
	reqTransaction.ResponseMidtrans = string(jsonData)

	ticketID, err := pgStoreClient.CreateTransaction(ctx, reqTransaction)
	if err != nil {
		return 0, err
	}

	// commit changes
	err = pgStoreClient.Commit()
	if err != nil {
		return 0, err
	}

	return ticketID, nil
}

func (s *service) GetTransactionByID(ctx context.Context, transactionID int64, nomorTiket string) (transaction.Transaction, error) {
	// validate id
	if transactionID <= 0 {
		return transaction.Transaction{}, transaction.ErrInvalidTransactionID
	}

	// get pg store client without using transaction
	pgStoreClient, err := s.pgStore.NewClient(false)
	if err != nil {
		return transaction.Transaction{}, err
	}

	// get user from pgstore
	result, err := pgStoreClient.GetTransactionByID(ctx, transactionID)
	if err != nil {
		return transaction.Transaction{}, err
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
			return transaction.Transaction{}, transaction.ErrDataNotFound
		}
	}

	return result, nil
}

// validateTransaction validates fields of the given transaction
// whether its comply the predetermined rules.
func validateTransaction(reqTransaction transaction.Transaction) error {
	if reqTransaction.Nama == "" {
		return transaction.ErrInvalidTransactionNama
	}

	if reqTransaction.JenisKelamin == "" || (reqTransaction.JenisKelamin != "Pria" && reqTransaction.JenisKelamin != "Wanita") {
		return transaction.ErrInvalidTransactionJenisKelamin
	}

	if reqTransaction.NomorIdentitas == "" || len(reqTransaction.NomorIdentitas) < 15 {
		return transaction.ErrInvalidTransactionNomorIdentitas
	}

	if reqTransaction.AsalInstitusi == "" {
		return transaction.ErrInvalidTransactionAsalInstitusi
	}

	if reqTransaction.Domisili == "" {
		return transaction.ErrInvalidTransactionDomisili
	}

	_, err := mail.ParseAddress(reqTransaction.Email)
	if reqTransaction.Email == "" || err != nil {
		return transaction.ErrInvalidTransactionEmail
	}

	if reqTransaction.NomorTelepon == "" || len(reqTransaction.NomorTelepon) < 10 || len(reqTransaction.NomorTelepon) > 13 {
		return transaction.ErrInvalidTransactionNomorTelepon
	}

	if reqTransaction.LineID == "" || strings.HasPrefix(reqTransaction.LineID, "@") {
		return transaction.ErrInvalidTransactionLineID
	}

	if reqTransaction.Instagram == "" || strings.HasPrefix(reqTransaction.LineID, "@") {
		return transaction.ErrInvalidTransactionInstagram
	}

	if reqTransaction.JumlahTiket <= 0 {
		return transaction.ErrInvalidTransactionJumlahTiket
	}

	if reqTransaction.Tanggal.IsZero() {
		return transaction.ErrInvalidTransactionTanggal
	}

	return nil
}

func (s *service) UpdateCheckInStatus(ctx context.Context, id int64, ticketNumber string) (string, error) {
	if id <= 0 {
		return "", transaction.ErrInvalidTransactionID
	}

	pgStoreClient, err := s.pgStore.NewClient(false)
	if err != nil {
		return "", err
	}

	tx, err := pgStoreClient.GetTransactionByID(ctx, id)
	if err != nil {
		return "", err
	}

	if tx.CheckInStatus {
		return "", transaction.ErrAllTicketAlreadyCheckedIn
	} else if tx.StatusPayment != "settlement" {
		return "", transaction.ErrTicketNotYetPaid
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
				return "", transaction.ErrTicketAlreadyCheckedIn
			}
		}
	} else {
		return "", transaction.ErrDataNotFound
	}

	tx.CheckInNomorTiket = append(tx.CheckInNomorTiket, ticketNumber)

	if len(tx.CheckInNomorTiket) == len(tx.NomorTiket) {
		tx.CheckInStatus = true
	}

	err = pgStoreClient.UpdateTransactionByID(ctx, tx, s.timeNow())
	if err != nil {
		log.Println(err)
		return "", err
	}

	return ticketNumber, nil
}
