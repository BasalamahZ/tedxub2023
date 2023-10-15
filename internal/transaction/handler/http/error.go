package http

import (
	"errors"

	"github.com/tedxub2023/internal/transaction"
)

// Followings are the known errors from Transaction HTTP handlers.
var (
	// errBadRequest is returned when the given request is
	// bad/invalid.
	errBadRequest = errors.New("BAD_REQUEST")

	// errInternalServer is returned when there is an
	// unexpected error encountered when processing a request.
	errInternalServer = errors.New("INTERNAL_SERVER_ERROR")

	// errInvalidTicketNama is returned when the given
	// ticket nama is invalid.
	errInvalidTicketNama = errors.New("INVALID_TICKET_NAMA")

	// errInvalidTicketNomorIdentitas is returned when the given
	// ticket nomor identitas is invalid.
	errInvalidTicketNomorIdentitas = errors.New("INVALID_TICKET_NOMOR_IDENTITAS")

	// errInvalidTicketAsalInstitusi is returned when the given
	// ticket asal institusi is invalid.
	errInvalidTicketAsalInstitusi = errors.New("INVALID_TICKET_ASAL_INSITITUSI")

	// errInvalidTicketDomisili is returned when the given
	// ticket domisili is invalid.
	errInvalidTicketDomisili = errors.New("INVALID_TICKET_DOMISILI")

	// errInvalidTicketEmail is returned when the given
	// ticket email is invalid.
	errInvalidTicketEmail = errors.New("INVALID_TICKET_EMAIL")

	// errInvalidTicketNomorTelepon is returned when the given
	// ticket nomor telepon is invalid.
	errInvalidTicketNomorTelepon = errors.New("INVALID_TICKET_NOMOR_TELEPON")

	// errInvalidTicketLineID is returned when the given
	// ticket line id is invalid.
	errInvalidTicketLineID = errors.New("INVALID_TICKET_LINE_ID")

	// errInvalidTicketInstagram is returned when the given
	// ticket instagram is invalid.
	errInvalidTicketInstagram = errors.New("INVALID_TICKET_INSTAGRAM")

	// errInvalidTicketJenisKelamin is returned when the request
	// is not valid format
	errInvalidTicketJenisKelamin = errors.New("INVALID_TICKET_JENIS_KELAMIN")

	// errMethodNotAllowed is returned when accessing not
	// allowed HTTP method.
	errMethodNotAllowed = errors.New("METHOD_NOT_ALLOWED")

	// errRequestTimeout is returned when processing time
	// has reached the timeout limit.
	errRequestTimeout = errors.New("REQUEST_TIMEOUT")
)

var (
	// mapHTTPError maps service error into HTTP error that
	// categorize as bad request error.
	//
	// Internal server error-related should not be mapped
	// here, and the handler should just return `errInternal`
	// as the error instead
	mapHTTPError = map[error]error{
		transaction.ErrInvalidTicketNama:           errInvalidTicketNama,
		transaction.ErrInvalidTicketNomorIdentitas: errInvalidTicketNomorIdentitas,
		transaction.ErrInvalidTicketAsalInstitusi:  errInvalidTicketAsalInstitusi,
		transaction.ErrInvalidTicketDomisili:       errInvalidTicketDomisili,
		transaction.ErrInvalidTicketEmail:          errInvalidTicketEmail,
		transaction.ErrInvalidTicketNomorTelepon:   errInvalidTicketNomorTelepon,
		transaction.ErrInvalidTicketLineID:         errInvalidTicketLineID,
		transaction.ErrInvalidTicketInstagram:      errInvalidTicketInstagram,
		transaction.ErrInvalidTicketJenisKelamin:   errInvalidTicketJenisKelamin,
	}
)
