package postgresql

import (
	"context"

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
