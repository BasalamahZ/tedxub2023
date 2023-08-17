package postgresql

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/tedxub2023/internal/ticket"
)

func (sc *storeClient) CreateTicket(ctx context.Context, reqTicket ticket.Ticket) (string, error) {
	// construct arguments filled with fields for the query
	argsKV := map[string]interface{}{
		"nama":            reqTicket.Nama,
		"nomor_identitas": reqTicket.NomorIdentitas,
		"asal_institusi":  reqTicket.AsalInstitusi,
		"domisili":        reqTicket.Domisili,
		"email":           reqTicket.Email,
		"nomor_telepon":   reqTicket.NomorTelepon,
		"line_id":         reqTicket.LineID,
		"instagram":       reqTicket.Instagram,
		"create_time":     reqTicket.CreateTime,
	}

	// prepare query
	query, args, err := sqlx.Named(queryCreateTicket, argsKV)
	if err != nil {
		return "", err
	}
	query, args, err = sqlx.In(query, args...)
	if err != nil {
		return "", err
	}
	query = sc.q.Rebind(query)

	// execute query
	var ticketNama string
	err = sc.q.QueryRowx(query, args...).Scan(&ticketNama)
	if err != nil {
		return "", err
	}

	return ticketNama, nil
}

func (sc *storeClient) GetAllTicket(ctx context.Context) ([]ticket.Ticket, error) {
	var tickets []ticket.Ticket

	rows, err := sc.q.Query(queryGetAllTicket)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var ticket ticket.Ticket
		err := rows.Scan(&ticket.ID, &ticket.Nama, &ticket.NomorIdentitas, &ticket.AsalInstitusi, &ticket.Domisili, &ticket.Email, &ticket.NomorTelepon, &ticket.LineID, &ticket.Instagram, &ticket.Status, &ticket.NomorTiket, &ticket.CreateTime, &ticket.UpdateTime)
		if err != nil {
			return nil, err
		}

		tickets = append(tickets, ticket)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tickets, nil
}

func (sc *storeClient) UpdateTicket(ctx context.Context, nomorTiket string, id int, updateTime time.Time) error {
	argsUpdate := map[string]interface{}{
		"status":      true,
		"nomor_tiket": nomorTiket,
		"update_time": updateTime,
		"id":          id,
	}

	queryUpdate, args, err := sqlx.Named(queryUpdateTicket, argsUpdate)

	if err != nil {
		return err
	}

	queryUpdate, args, err = sqlx.In(queryUpdate, args...)

	if err != nil {
		return err
	}

	queryUpdate = sc.q.Rebind(queryUpdate)

	_, err = sc.q.Exec(queryUpdate, args...)

	if err != nil {
		return err
	}

	return nil
}
