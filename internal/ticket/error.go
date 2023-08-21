package ticket

import "errors"

// Followings are the known errors returned from ticket.
var (
	// ErrInvalidTicketNama is returned when the given ticket
	// nama is invalid.
	ErrInvalidTicketNama = errors.New("invalid ticket nama")

	// ErrInvalidTicketNomorIdentitas is returned when the given ticket
	// nomor identitas is invalid.
	ErrInvalidTicketNomorIdentitas = errors.New("invalid ticket nomor identitas")

	// ErrInvalidTicketAsalInstitusi is returned when the given ticket
	// asal institusi is invalid.
	ErrInvalidTicketAsalInstitusi = errors.New("invalid ticket asal institusi")

	// ErrInvalidTicketDomisili is returned when the given ticket
	// domisili is invalid.
	ErrInvalidTicketDomisili = errors.New("invalid ticket domisili")

	// ErrInvalidTicketEmail is returned when the given ticket
	// email is invalid.
	ErrInvalidTicketEmail = errors.New("invalid ticket email")

	// ErrInvalidTicketNomorTelepon is returned when the given ticket
	// nomor telepon is invalid.
	ErrInvalidTicketNomorTelepon = errors.New("invalid ticket nomor telepon")

	// ErrInvalidTicketLineID is returned when the given ticket
	// line id is invalid.
	ErrInvalidTicketLineID = errors.New("invalid ticket line id")

	// ErrInvalidTicketInstagram is returned when the given ticket
	// instagram is invalid.
	ErrInvalidTicketInstagram = errors.New("invalid ticket instagram")
)

var (
	// errFailedTransaction is returned when the request
	// failed to make client with transaction
	ErrFailedTransaction = errors.New("FAILED_CREATE_TRANSACTION")

	// errSendEmail is returned when the request
	// failed to send email
	ErrSendEmail = errors.New("FAILED_SEND_EMAIL")

	// errParseBodyHTML is returned when the request
	// failed to parse body html
	ErrParseBodyHTML = errors.New("FAILED_PARSE_BODY_HTML")

	// errTicketNotFound is returned when the request
	// failed to find ticket
	ErrTicketNotFound = errors.New("TICKET_NOT_FOUND")
)
