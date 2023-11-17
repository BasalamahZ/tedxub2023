package postgresql

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/tedxub2023/internal/mainevent"
	"github.com/tedxub2023/internal/mainevent/service"
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

type maineventDB struct {
	ID                int64                `db:"id"`
	Nama              string               `db:"nama"`
	Disabilitas       mainevent.Disability `db:"disabilitas"`
	NomorIdentitas    string               `db:"nomor_identitas"`
	AsalInstitusi     string               `db:"asal_institusi"`
	Email             string               `db:"email"`
	NomorTelepon      string               `db:"nomor_telepon"`
	JumlahTiket       int                  `db:"jumlah_tiket"`
	TotalHarga        int64                `db:"total_harga"`
	OrderID           string               `db:"order_id"`
	Type              mainevent.Type       `db:"type"`
	Status            mainevent.Status     `db:"status"`
	ImageURI          *string              `db:"image_uri"`
	NomorTiket        pq.StringArray       `db:"nomor_tiket"`
	CheckInStatus     *bool                `db:"checkin_status"`
	CheckInNomorTiket pq.StringArray       `db:"checkin_nomor_tiket"`
	CreateTime        time.Time            `db:"create_time"`
	UpdateTime        *time.Time           `db:"update_time"`
}

// format formats database struct into domain struct.
func (mdb *maineventDB) format() mainevent.MainEvent {
	t := mainevent.MainEvent{
		ID:             mdb.ID,
		Nama:           mdb.Nama,
		Disabilitas:    mdb.Disabilitas,
		NomorIdentitas: mdb.NomorIdentitas,
		AsalInstitusi:  mdb.AsalInstitusi,
		Email:          mdb.Email,
		NomorTelepon:   mdb.NomorTelepon,
		JumlahTiket:    mdb.JumlahTiket,
		TotalHarga:     mdb.TotalHarga,
		OrderID:        mdb.OrderID,
		Type:           mdb.Type,
		Status:         mdb.Status,
		CreateTime:     mdb.CreateTime,
	}

	if len(mdb.NomorTiket) > 0 {
		ticketNumbers := make([]string, 0)
		for _, ticketNumber := range mdb.NomorTiket {
			ticketNumbers = append(ticketNumbers, ticketNumber)
		}
		t.NomorTiket = ticketNumbers
	}

	if mdb.ImageURI != nil {
		t.ImageURI = *mdb.ImageURI
	}

	if mdb.CheckInStatus != nil {
		t.CheckInStatus = *mdb.CheckInStatus
	}

	if len(mdb.CheckInNomorTiket) > 0 {
		checkinTicketNumbers := make([]string, 0)
		for _, checkinTicketNumber := range mdb.CheckInNomorTiket {
			checkinTicketNumbers = append(checkinTicketNumbers, checkinTicketNumber)
		}
		t.CheckInNomorTiket = checkinTicketNumbers
	}

	if mdb.UpdateTime != nil {
		t.UpdateTime = *mdb.UpdateTime
	}

	return t
}
