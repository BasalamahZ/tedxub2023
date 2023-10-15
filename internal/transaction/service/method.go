package service

import (
	"context"
	"fmt"
	"hash/crc32"
	"net/mail"
	"strings"

	"github.com/tedxub2023/internal/transaction"
)

func (s *service) ReplaceTransactionByEmail(ctx context.Context, reqTransaction transaction.Transaction) (int64, error) {
	// validate field
	err := validateTransaction(reqTransaction)
	if err != nil {
		return 0, err
	}
	reqTransaction.CreateTime = s.timeNow()

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

	// delete transaction specified with email in pgstore
	err = pgStoreClient.DeleteTransactionByEmail(ctx, reqTransaction.Email)
	if err != nil {
		return 0, err
	}

	ticketNumbers, err := genereteTicketNumber(reqTransaction)
	if err != nil {
		return 0, err
	}
	reqTransaction.NomorTiket = ticketNumbers

	resPayment, err := payment(reqTransaction)
	if err != nil {
		return 0, err
	}
	reqTransaction.ResponseMidtrans = resPayment

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

func genereteTicketNumber(reqTransaction transaction.Transaction) ([]string, error) {
	// Calculate the CRC32 hash of the data
	hash := crc32.ChecksumIEEE([]byte(reqTransaction.Email))

	// Reduce the hash to a 3-digit number
	uniqueNumber := int(hash % 100)
	// uniqueNumberStr := strconv.Itoa(uniqueNumber)
	uniqueNumberStr := fmt.Sprintf("%02d", uniqueNumber)

	ticketNumbers := make([]string, 0)
	// combination := reqTransaction.JumlahTiket * 100
	for i := 0; i < reqTransaction.JumlahTiket; i++ {
		letter := string('A' + i)

		combination := letter + uniqueNumberStr

		ticketNumber := fmt.Sprintf("TICKET/TEDXUB/%s", combination)

		ticketNumbers = append(ticketNumbers, ticketNumber)
	}

	return ticketNumbers, nil
}

// validateTransaction validates fields of the given Ticket
// whether its comply the predetermined rules.
func validateTransaction(reqTransaction transaction.Transaction) error {
	if reqTransaction.Nama == "" {
		return transaction.ErrInvalidTicketNama
	}

	if reqTransaction.JenisKelamin == "" || (reqTransaction.JenisKelamin != "Pria" && reqTransaction.JenisKelamin != "Wanita") {
		return transaction.ErrInvalidTicketJenisKelamin
	}

	if reqTransaction.NomorIdentitas == "" || len(reqTransaction.NomorIdentitas) < 15 {
		return transaction.ErrInvalidTicketNomorIdentitas
	}

	if reqTransaction.AsalInstitusi == "" {
		return transaction.ErrInvalidTicketAsalInstitusi
	}

	if reqTransaction.Domisili == "" {
		return transaction.ErrInvalidTicketDomisili
	}

	_, err := mail.ParseAddress(reqTransaction.Email)
	if reqTransaction.Email == "" || err != nil {
		return transaction.ErrInvalidTicketEmail
	}

	if reqTransaction.NomorTelepon == "" || len(reqTransaction.NomorTelepon) < 10 || len(reqTransaction.NomorTelepon) > 13 {
		return transaction.ErrInvalidTicketNomorTelepon
	}

	if reqTransaction.LineID == "" || strings.HasPrefix(reqTransaction.LineID, "@") {
		return transaction.ErrInvalidTicketLineID
	}

	if reqTransaction.Instagram == "" || strings.HasPrefix(reqTransaction.LineID, "@") {
		return transaction.ErrInvalidTicketInstagram
	}

	return nil
}
