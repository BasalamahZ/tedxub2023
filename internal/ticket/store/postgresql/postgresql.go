package postgresql

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/tedxub2023/internal/ticket"
	"github.com/tedxub2023/internal/ticket/service"
)

var (
	errInvalidCommit   = errors.New("cannot do commit on non-transactional querier")
	errInvalidRollback = errors.New("cannot do rollback on non-transactional querier")
)

// store implements configuration/service.PGStore
type store struct {
	db *sqlx.DB
}

// storeClient implements configuration/service.PGStoreClient
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

type TicketDB struct {
	ID             int64      `db:"id"`
	Nama           string     `db:"nama"`
	NomorIdentitas string     `db:"nomor_identitas"`
	AsalInstitusi  string     `db:"asal_institusi"`
	Domisili       string     `db:"domisili"`
	Email          string     `db:"email"`
	NomorTelepon   string     `db:"nomor_telepon"`
	LineID         string     `db:"line_id"`
	Instagram      string     `db:"instagram"`
	Status         *bool      `db:"status"`
	NomorTiket     *string    `db:"nomor_tiket"`
	CreateTime     time.Time  `db:"create_time"`
	UpdateTime     *time.Time `db:"update_time"`
}

func (tdb *TicketDB) formatting() ticket.Ticket {
	t := ticket.Ticket{
		ID:             tdb.ID,
		Nama:           tdb.Nama,
		NomorIdentitas: tdb.NomorIdentitas,
		AsalInstitusi:  tdb.AsalInstitusi,
		Domisili:       tdb.Domisili,
		Email:          tdb.Email,
		NomorTelepon:   tdb.NomorTelepon,
		LineID:         tdb.LineID,
		Instagram:      tdb.Instagram,
		CreateTime:     tdb.CreateTime,
	}

	if tdb.Status != nil {
		t.Status = *tdb.Status
	}

	if tdb.NomorTiket != nil {
		t.NomorTiket = *tdb.NomorTiket
	}

	if tdb.UpdateTime != nil {
		t.UpdateTime = *tdb.UpdateTime
	}

	return t
}
