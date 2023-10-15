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
	ID               int64          `db:"id"`
	Nama             string         `db:"nama"`
	JenisKelamin     string         `db:"jenis_kelamin"`
	NomorIdentitas   string         `db:"nomor_identitas"`
	AsalInstitusi    string         `db:"asal_institusi"`
	Domisili         string         `db:"domisili"`
	Email            string         `db:"email"`
	NomorTelepon     string         `db:"nomor_telepon"`
	LineID           string         `db:"line_id"`
	Instagram        string         `db:"instagram"`
	JumlahTiket      int            `db:"jumlah_tiket"`
	Harga            int64          `db:"harga"`
	NomorTiket       pq.StringArray `db:"nomor_tiket"`
	ResponseMidtrans string         `db:"response_midtrans"`
	CheckInStatus    *bool          `db:"checkin_status"`
	CheckInCounter   *int           `db:"checkin_counter"`
	CreateTime       time.Time      `db:"create_time"`
	UpdateTime       *time.Time     `db:"update_time"`
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
		Harga:            tdb.Harga,
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

	if tdb.CheckInCounter != nil {
		t.CheckInCounter = *tdb.CheckInCounter
	}

	if tdb.UpdateTime != nil {
		t.UpdateTime = *tdb.UpdateTime
	}

	return t
}
