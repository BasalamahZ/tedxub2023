package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/tedxub2023/internal/transaction"
)

func (sc *storeClient) DeleteTransactionByEmail(ctx context.Context, email string, tanggal time.Time) error {
	// construct arguments filled with fields for the query
	argsKV := map[string]interface{}{
		"email":          email,
		"tanggal":        tanggal,
		"status_payment": "pending",
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
		"total_harga":       reqTransaction.TotalHarga,
		"tanggal":           reqTransaction.Tanggal,
		"order_id":          reqTransaction.OrderID,
		"status_payment":    reqTransaction.StatusPayment,
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

func (sc *storeClient) GetTransactionByID(ctx context.Context, transactionID int64) (transaction.Transaction, error) {
	query := fmt.Sprintf(queryGetTransaction, "WHERE t.id = $1")

	// query single row
	var tdb transactionDB
	err := sc.q.QueryRowx(query, transactionID).StructScan(&tdb)
	if err != nil {
		if err == sql.ErrNoRows {
			return transaction.Transaction{}, transaction.ErrDataNotFound
		}
		return transaction.Transaction{}, err
	}

	return tdb.format(), nil
}
