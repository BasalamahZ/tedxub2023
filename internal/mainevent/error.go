package mainevent

import "errors"

// Followings are the known errors returned from mainevent.
var (
	// ErrDataNotFound is returned when the wanted data is
	// not found.
	ErrDataNotFound = errors.New("data not found")

	// ErrInvalidMainEventID is returned when the given main event
	// id is invalid.
	ErrInvalidMainEventID = errors.New("invalid main event id")

	// ErrInvalidMainEventNama is returned when the given main event
	// nama is invalid.
	ErrInvalidMainEventNama = errors.New("invalid main event nama")

	// ErrInvalidMainEventNomorIdentitas is returned when the given main event
	// nomor identitas is invalid.
	ErrInvalidMainEventNomorIdentitas = errors.New("invalid main event nomor identitas")

	// ErrInvalidMainEventAsalInstitusi is returned when the given main event
	// asal institusi is invalid.
	ErrInvalidMainEventAsalInstitusi = errors.New("invalid main event asal institusi")

	// ErrInvalidMainEventEmail is returned when the given main event
	// email is invalid.
	ErrInvalidMainEventEmail = errors.New("invalid main event email")

	// ErrInvalidMainEventNomorTelepon is returned when the given main event
	// nomor telepon is invalid.
	ErrInvalidMainEventNomorTelepon = errors.New("invalid main event nomor telepon")

	// ErrInvalidMainEventInstagram is returned when the given main event
	// instagram is invalid.
	ErrInvalidMainEventInstagram = errors.New("invalid main event instagram")

	// errInvalidMainEventJumlahTiket is returned when the given main event
	// jumlah tiket is invalid.
	ErrInvalidMainEventJumlahTiket = errors.New("invalid main event jumlah tiket")

	// errInvalidMainEventType is returned when the given main event
	// type is invalid.
	ErrInvalidMainEventType = errors.New("invalid main event type")

	// errInvalidMainEventStatus is returned when the given main event
	// status is invalid.
	ErrInvalidMainEventStatus = errors.New("invalid main event status")

	// errInvalidMainEventImageURI is returned when the given main event
	// image uri is invalid.
	ErrInvalidMainEventImageURI = errors.New("invalid main event image uri")

	// ErrTicketAlreadyCheckedIn is returned when the given
	// ticket_number is already checkedIn
	ErrTicketAlreadyCheckedIn = errors.New("ticket already checked in")

	// ErrTicketNotFound is returned when the given
	// ticket_number is not in the row with id given
	ErrTicketNotFound = errors.New("ticket not found")

	// ErrTicketNotYetPaid is returned when the given
	// ticket_number is not yet paid
	ErrTicketNotYetPaid = errors.New("ticket not yet paid")

	// ErrAllTicketAlreadyCheckedIn is returned when all ticket
	// already checked in
	ErrAllTicketAlreadyCheckedIn = errors.New("all ticket already checked in")

	// ErrPaymentNoSettlement is returned when the given payment
	// is not settlement.
	ErrPaymentNotSettlement = errors.New("payment not settlement")

	//ErrEarlyBirdTicketSoldOut is returned when the given early bird ticket
	//is sold out
	ErrEarlyBirdTicketSoldOut = errors.New("early bird ticket sold out or total ticket buy over the limit")
)
