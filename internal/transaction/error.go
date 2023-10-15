package transaction

import "errors"

// Followings are the known errors returned from transaction.
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

	// errInvalidTicketJenisKelamin is returned when the request
	// is not valid format
	ErrInvalidTicketJenisKelamin = errors.New("invalid ticket jenis kelamin")
)
