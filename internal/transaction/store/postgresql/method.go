package postgresql

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/tedxub2023/internal/transaction"
)

func (sc *storeClient) DeleteTransactionByEmail(ctx context.Context, email string) error {
	// construct arguments filled with fields for the query
	argsKV := map[string]interface{}{
		"email": email,
	}

	// prepare query
	query, args, err := sqlx.Named(queryuDeleteTransactionByEmail, argsKV)
	if err != nil {
		return err
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return err
	}
	query = sc.q.Rebind(query)

	// execute query
	_, err = sc.q.Exec(query, args...)
	return err
}

func (sc *storeClient) CreateTransaction(ctx context.Context, reqTransaction transaction.Transaction) (int64, error) {
	ticketNumbers := make([]string, 0)
	ticketNumbers = append(ticketNumbers, reqTransaction.NomorTiket...)

	// construct arguments filled with fields for the query
	argsKV := map[string]interface{}{
		"nama":              reqTransaction.Nama,
		"jenis_kelamin":     reqTransaction.JenisKelamin,
		"nomor_identitas":   reqTransaction.NomorIdentitas,
		"asal_institusi":    reqTransaction.AsalInstitusi,
		"domisili":          reqTransaction.Domisili,
		"email":             reqTransaction.Email,
		"nomor_telepon":     reqTransaction.NomorTelepon,
		"line_id":           reqTransaction.LineID,
		"instagram":         reqTransaction.Instagram,
		"jumlah_tiket":      reqTransaction.JumlahTiket,
		"harga":             reqTransaction.Harga,
		"nomor_tiket":       pq.StringArray(ticketNumbers),
		"response_midtrans": reqTransaction.ResponseMidtrans,
		"create_time":       reqTransaction.CreateTime,
	}

	// prepare query
	query, args, err := sqlx.Named(queryCreateTransaction, argsKV)
	if err != nil {
		return 0, err
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return 0, err
	}
	query = sc.q.Rebind(query)

	// execute query
	var ticketID int64
	err = sc.q.QueryRowx(query, args...).Scan(&ticketID)
	if err != nil {
		return 0, err
	}

	return ticketID, nil
}
