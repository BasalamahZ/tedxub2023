package postgresql

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/tedxub2023/internal/transaction"
	"github.com/tedxub2023/internal/transaction/service"
)

var (
	errInvalidCommit   = errors.New("cannot do commit on non-transactional querier")
	errInvalidRollback = errors.New("cannot do rollback on non-transactional querier")
)

// store implements transaction/service.PGStore
type store struct {
	db *sqlx.DB
}

// storeClient implements transaction/service.PGStoreClient
type storeClient struct {
	q sqlx.Ext
}

// New creates a new store.
func New(db *sqlx.DB) (*store, error) {
	s := &store{
		db: db,
	}

	return s, nil
}

func (s *store) NewClient(useTx bool) (service.PGStoreClient, error) {
	var q sqlx.Ext

	// determine what object should be use as querier
	q = s.db
	if useTx {
		var err error
		q, err = s.db.Beginx()
		if err != nil {
			return nil, err
		}
	}

	return &storeClient{
		q: q,
	}, nil
}

func (sc *storeClient) Commit() error {
	if tx, ok := sc.q.(*sqlx.Tx); ok {
		return tx.Commit()
	}
	return errInvalidCommit
}

func (sc *storeClient) Rollback() error {
	if tx, ok := sc.q.(*sqlx.Tx); ok {
		return tx.Rollback()
	}
	return errInvalidRollback
}

type transactionDB struct {
	ID                int64          `db:"id"`
	Nama              string         `db:"nama"`
	JenisKelamin      string         `db:"jenis_kelamin"`
	NomorIdentitas    string         `db:"nomor_identitas"`
	AsalInstitusi     string         `db:"asal_institusi"`
	Domisili          string         `db:"domisili"`
	Email             string         `db:"email"`
	NomorTelepon      string         `db:"nomor_telepon"`
	LineID            string         `db:"line_id"`
	Instagram         string         `db:"instagram"`
	JumlahTiket       int            `db:"jumlah_tiket"`
	TotalHarga        int64          `db:"total_harga"`
	Tanggal           time.Time      `db:"tanggal"`
	OrderID           string         `db:"order_id"`
	StatusPayment     string         `db:"status_payment"`
	ResponseMidtrans  string         `db:"response_midtrans"`
	NomorTiket        pq.StringArray `db:"nomor_tiket"`
	CheckInStatus     *bool          `db:"checkin_status"`
	CheckInNomorTiket pq.StringArray `db:"checkin_nomor_tiket"`
	CreateTime        time.Time      `db:"create_time"`
	UpdateTime        *time.Time     `db:"update_time"`
}

// format formats database struct into domain struct.
func (tdb *transactionDB) format() transaction.Transaction {
	t := transaction.Transaction{
		ID:               tdb.ID,
		Nama:             tdb.Nama,
		JenisKelamin:     tdb.JenisKelamin,
		NomorIdentitas:   tdb.NomorIdentitas,
		AsalInstitusi:    tdb.AsalInstitusi,
		Domisili:         tdb.Domisili,
		Email:            tdb.Email,
		NomorTelepon:     tdb.NomorTelepon,
		LineID:           tdb.LineID,
		Instagram:        tdb.Instagram,
		JumlahTiket:      tdb.JumlahTiket,
		TotalHarga:       tdb.TotalHarga,
		Tanggal:          tdb.Tanggal,
		OrderID:          tdb.OrderID,
		StatusPayment:    tdb.StatusPayment,
		ResponseMidtrans: tdb.ResponseMidtrans,
		CreateTime:       tdb.CreateTime,
	}

	if len(tdb.NomorTiket) > 0 {
		ticketNumbers := make([]string, 0)
		for _, ticketNumber := range tdb.NomorTiket {
			ticketNumbers = append(ticketNumbers, ticketNumber)
		}
		t.NomorTiket = ticketNumbers
	}

	if tdb.CheckInStatus != nil {
		t.CheckInStatus = *tdb.CheckInStatus
	}

	if len(tdb.CheckInNomorTiket) > 0 {
		checkinTicketNumbers := make([]string, 0)
		for _, checkinTicketNumber := range tdb.CheckInNomorTiket {
			checkinTicketNumbers = append(checkinTicketNumbers, checkinTicketNumber)
		}
		t.CheckInNomorTiket = checkinTicketNumbers
	}

	if tdb.UpdateTime != nil {
		t.UpdateTime = *tdb.UpdateTime
	}

	return t
}
