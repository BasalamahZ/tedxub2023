package http

import (
	"errors"

	"github.com/tedxub2023/internal/ticket"
)

// Followings are the known errors from Ticket HTTP handlers.
var (
	// errBadRequest is returned when the given request is
	// bad/invalid.
	errBadRequest = errors.New("BAD_REQUEST")

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

	// errInternalServer is returned when there is an
	// unexpected error encountered when processing a request.
	errInternalServer = errors.New("INTERNAL_SERVER_ERROR")

	// errMethodNotAllowed is returned when accessing not
	// allowed HTTP method.
	errMethodNotAllowed = errors.New("METHOD_NOT_ALLOWED")

	// errRequestTimeout is returned when processing time
	// has reached the timeout limit.
	errRequestTimeout = errors.New("REQUEST_TIMEOUT")

	// errFailedTransaction is returned when the request
	// failed to make client with transaction
	errFailedTransaction = errors.New("FAILED_CREATE_TRANSACTION")

	// errSendEmail is returned when the request
	// failed to send email
	errSendEmail = errors.New("FAILED_SEND_EMAIL")

	// errParseBodyHTML is returned when the request
	// failed to parse body html
	errParseBodyHTML = errors.New("FAILED_PARSE_BODY_HTML")

	// errTicketNotFound is returned when the request
	// failed to find ticket
	errTicketNotFound = errors.New("TICKET_NOT_FOUND")

	// errInvalidTicketJenisKelamin is returned when the request
	// is not valid format
	errInvalidTicketJenisKelamin = errors.New("INVALID_TICKET_JENIS_KELAMIN")

	// errEmailAlreadyRegistered is returned when the request
	// email already in DB
	errEmailAlreadyRegistered = errors.New("EMAIL_ALREADY_REGISTERED")

	// errNumberIdentityAlreadyRegistered is returned when the request
	// number identity already in DB
	errNumberIdentityAlreadyRegistered = errors.New("NUMBER_IDENTITY_ALREADY_REGISTERED")
)

var (
	// mapHTTPError maps service error into HTTP error that
	// categorize as bad request error.
	//
	// Internal server error-related should not be mapped
	// here, and the handler should just return `errInternal`
	// as the error instead
	mapHTTPError = map[error]error{
		ticket.ErrInvalidTicketNama:               errInvalidTicketNama,
		ticket.ErrInvalidTicketNomorIdentitas:     errInvalidTicketNomorIdentitas,
		ticket.ErrInvalidTicketAsalInstitusi:      errInvalidTicketAsalInstitusi,
		ticket.ErrInvalidTicketDomisili:           errInvalidTicketDomisili,
		ticket.ErrInvalidTicketEmail:              errInvalidTicketEmail,
		ticket.ErrInvalidTicketNomorTelepon:       errInvalidTicketNomorTelepon,
		ticket.ErrInvalidTicketLineID:             errInvalidTicketLineID,
		ticket.ErrInvalidTicketInstagram:          errInvalidTicketInstagram,
		ticket.ErrFailedTransaction:               errFailedTransaction,
		ticket.ErrSendEmail:                       errSendEmail,
		ticket.ErrParseBodyHTML:                   errParseBodyHTML,
		ticket.ErrTicketNotFound:                  errTicketNotFound,
		ticket.ErrInvalidTicketJenisKelamin:       errInvalidTicketJenisKelamin,
		ticket.ErrEmailAlreadyRegistered:          errEmailAlreadyRegistered,
		ticket.ErrNumberIdentityAlreadyRegistered: errNumberIdentityAlreadyRegistered,
	}
)
