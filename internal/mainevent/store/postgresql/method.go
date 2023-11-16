package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/tedxub2023/internal/mainevent"
)

func (sc *storeClient) DeleteMainEventByEmail(ctx context.Context, email string) error {
	// construct arguments filled with fields for the query
	argsKV := map[string]interface{}{
		"email":  email,
		"status": mainevent.StatusUnpaid,
	}

	// prepare query
	query, args, err := sqlx.Named(queryuDeleteMainEventByEmail, argsKV)
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

func (sc *storeClient) CreateMainEvent(ctx context.Context, reqMainEvent mainevent.MainEvent) (int64, error) {
	// construct arguments filled with fields for the query
	argsKV := map[string]interface{}{
		"nama":            reqMainEvent.Nama,
		"disabilitas":     reqMainEvent.Disabilitas,
		"nomor_identitas": reqMainEvent.NomorIdentitas,
		"asal_institusi":  reqMainEvent.AsalInstitusi,
		"email":           reqMainEvent.Email,
		"nomor_telepon":   reqMainEvent.NomorTelepon,
		"jumlah_tiket":    reqMainEvent.JumlahTiket,
		"total_harga":     reqMainEvent.TotalHarga,
		"order_id":        reqMainEvent.OrderID,
		"status":          reqMainEvent.Status,
		"type":            reqMainEvent.Type,
		"image_uri":       reqMainEvent.ImageURI,
		"create_time":     reqMainEvent.CreateTime,
	}

	// prepare query
	query, args, err := sqlx.Named(queryCreateMainEvent, argsKV)
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

func (sc *storeClient) GetAllMainEvents(ctx context.Context, filter mainevent.GetAllMainEventsFilter) ([]mainevent.MainEvent, error) {
	// define variables to custom query
	argsKV := make(map[string]interface{})
	addConditions := make([]string, 0)

	if filter.Status != 0 {
		addConditions = append(addConditions, "m.status = :status")
		argsKV["status"] = filter.Status
	}
	if filter.Type != 0 {
		addConditions = append(addConditions, "m.type = :type")
		argsKV["type"] = filter.Type
	}
	// construct strings to custom query
	addCondition := strings.Join(addConditions, " AND ")

	// since the query does not contains "WHERE" yet, need
	// to add it if needed
	if len(addConditions) > 0 {
		addCondition = fmt.Sprintf("WHERE %s", addCondition)
	}
	query := fmt.Sprintf(queryGetMainEvent, addCondition)

	// prepare query
	query, args, err := sqlx.Named(query, argsKV)
	if err != nil {
		return nil, err
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return nil, err
	}
	query = sc.q.Rebind(query)

	// query to database
	rows, err := sc.q.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// read match
	match := make([]mainevent.MainEvent, 0)
	for rows.Next() {
		var row maineventDB
		err = rows.StructScan(&row)
		if err != nil {
			return nil, err
		}

		match = append(match, row.format())
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return match, nil
}

func (sc *storeClient) GetMainEventByID(ctx context.Context, maineventID int64) (mainevent.MainEvent, error) {
	query := fmt.Sprintf(queryGetMainEvent, "WHERE m.id = $1")

	// query single row
	var mdb maineventDB
	err := sc.q.QueryRowx(query, maineventID).StructScan(&mdb)
	if err != nil {
		if err == sql.ErrNoRows {
			return mainevent.MainEvent{}, mainevent.ErrDataNotFound
		}
		return mainevent.MainEvent{}, err
	}

	return mdb.format(), nil
}

func (sc *storeClient) UpdateMainEventByID(ctx context.Context, tx mainevent.MainEvent) error {
	argsKV := map[string]interface{}{
		"nama":            tx.Nama,
		"disabilitas":     tx.Disabilitas,
		"nomor_identitas": tx.NomorIdentitas,
		"asal_institusi":  tx.AsalInstitusi,
		"email":           tx.Email,
		"nomor_telepon":   tx.NomorTelepon,
		"jumlah_tiket":    tx.JumlahTiket,
		"total_harga":     tx.TotalHarga,
		"order_id":        tx.OrderID,
		"status":          tx.Status,
		"image_uri":       tx.ImageURI,
		"nomor_tiket":     tx.NomorTiket,
		"checkin_status":  tx.CheckInStatus,
		"update_time":     tx.UpdateTime,
		"id":              tx.ID,
	}
	query := fmt.Sprintf(queryUpdateMainEvent, "")

	if len(tx.CheckInNomorTiket) != 0 {
		argsKV["checkin_nomor_tiket"] = tx.CheckInNomorTiket
		query = fmt.Sprintf(queryUpdateMainEvent, ", checkin_nomor_tiket = ARRAY[:checkin_nomor_tiket]")
	}

	query, args, err := sqlx.Named(query, argsKV)
	if err != nil {
		return err
	}

	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return err
	}

	query = sc.q.Rebind(query)

	_, err = sc.q.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
