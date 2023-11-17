package http

import (
	"errors"

	"github.com/tedxub2023/internal/mainevent"
)

// Followings are the known errors from Main Event HTTP handlers.
var (
	// errBadRequest is returned when the given request is
	// bad/invalid.
	errBadRequest = errors.New("BAD_REQUEST")

	// errInternalServer is returned when there is an
	// unexpected error encountered when processing a request.
	errInternalServer = errors.New("INTERNAL_SERVER_ERROR")

	// errDataNotFound is returned when the desired data is
	// not found.
	errDataNotFound = errors.New("DATA_NOT_FOUND")

	// errInvalidMainEventID is returned when the given
	// main event id is invalid.
	errInvalidMainEventID = errors.New("INVALID_MAIN_EVENT_ID")

	// errInvalidMainEventNama is returned when the given
	// main event nama is invalid.
	errInvalidMainEventNama = errors.New("INVALID_MAIN_EVENT_NAMA")

	// errInvalidMainEventNomorIdentitas is returned when the given
	// main event nomor identitas is invalid.
	errInvalidMainEventNomorIdentitas = errors.New("INVALID_MAIN_EVENT_NOMOR_IDENTITAS")

	// errInvalidMainEventAsalInstitusi is returned when the given
	// main event asal institusi is invalid.
	errInvalidMainEventAsalInstitusi = errors.New("INVALID_MAIN_EVENT_ASAL_INSITITUSI")

	// errInvalidMainEventEmail is returned when the given
	// main event email is invalid.
	errInvalidMainEventEmail = errors.New("INVALID_MAIN_EVENT_EMAIL")

	// errInvalidMainEventNomorTelepon is returned when the given
	// main event nomor telepon is invalid.
	errInvalidMainEventNomorTelepon = errors.New("INVALID_MAIN_EVENT_NOMOR_TELEPON")

	// errInvalidMainEventInstagram is returned when the given
	// main event instagram is invalid.
	errInvalidMainEventInstagram = errors.New("INVALID_MAIN_EVENT_INSTAGRAM")

	// errInvalidMainEventJumlahTiket is returned when the given
	// main event jumlah tiket is invalid.
	errInvalidMainEventJumlahTiket = errors.New("INVALID_MAIN_EVENT_JUMLAH_TIKET")

	// errInvalidMainEventType is returned when the given
	// main event type is invalid.
	errInvalidMainEventType = errors.New("INVALID_MAIN_EVENT_TYPE")

	// errInvalidMainEventStatus is returned when the given
	// main event status is invalid.
	errInvalidMainEventStatus = errors.New("INVALID_MAIN_EVENT_STATUS")

	// errInvalidMainEventImageURI is returned when the given
	// main event image uri is invalid.
	errInvalidMainEventImageURI = errors.New("INVALID_MAIN_EVENT_IMAGE_URI")

	// errInvalidMainEventFileURI is returned when the given
	// main event file uri is invalid.
	errInvalidMainEventFileURI = errors.New("INVALID_MAIN_EVENT_FILE_URI")

	// errMethodNotAllowed is returned when accessing not
	// allowed HTTP method.
	errMethodNotAllowed = errors.New("METHOD_NOT_ALLOWED")

	// errRequestTimeout is returned when processing time
	// has reached the timeout limit.
	errRequestTimeout = errors.New("REQUEST_TIMEOUT")

	// errInvalidTicketNumber is returned when the given
	// ticket number is invalid.
	errInvalidTicketNumber = errors.New("INVALID_TICKET_NUMBER")

	// errFailedJSONMarshal is returned when the JSONs
	// marshal process failed.
	errFailedJSONMarshal = errors.New("FAILED_JSON_MARSHAL")

	// errTicketAlreadyCheckedIn is returned when the given
	// ticket_number is already checkedIn
	errTicketAlreadyCheckedIn = errors.New("TICKET_ALREADY_CHECKED_IN")

	// errTicketNotFound is returned when the given
	// ticket_number is not in the row with id given
	errTicketNotFound = errors.New("TICKET_NOT_FOUND")

	// errTicketNotYetPaid is returned when the given
	// ticket_number is not yet paid
	errTicketNotYetPaid = errors.New("TIKET_NOT_YET_PAID")

	// errAllTicketAlreadyCheckedIn is returned when all ticket
	// already checked in
	errAllTicketAlreadyCheckedIn = errors.New("ALL_TICKET_ALREADY_CHECKED_IN")

	// errPaymentNoSettlement is returned when the given payment
	// is not settlement.
	errPaymentNotSettlement = errors.New("PAYMENT_NOT_SETTLEMENT")

	// errInvalidMainEventDisability is returned when the given
	// main event disability is invalid.
	errInvalidMainEventDisability = errors.New("INVALID_MAIN_EVENT_DISABILITY")

	// errEarlyBirdTicketSoldOut is returned when the early bird
	// ticket is sold out.
	errEarlyBirdTicketSoldOut = errors.New("EARLY_BIRD_TICKET_SOLD_OUT_OR_TOTAL_TICKET_BUY_OVER_THE_LIMIT")

	// errTicketNotAlreadyAccepted is returned when the given
	// ticketis not already accepted
	errTicketNotAlreadyAccepted = errors.New("TICKET_NOT_ALREADY_ACCEPTED")
)

var (
	// mapHTTPError maps service error into HTTP error that
	// categorize as bad request error.
	//
	// Internal server error-related should not be mapped
	// here, and the handler should just return `errInternal`
	// as the error instead
	mapHTTPError = map[error]error{
		mainevent.ErrDataNotFound:                   errDataNotFound,
		mainevent.ErrInvalidMainEventID:             errInvalidMainEventID,
		mainevent.ErrInvalidMainEventNama:           errInvalidMainEventNama,
		mainevent.ErrInvalidMainEventNomorIdentitas: errInvalidMainEventNomorIdentitas,
		mainevent.ErrInvalidMainEventAsalInstitusi:  errInvalidMainEventAsalInstitusi,
		mainevent.ErrInvalidMainEventEmail:          errInvalidMainEventEmail,
		mainevent.ErrInvalidMainEventNomorTelepon:   errInvalidMainEventNomorTelepon,
		mainevent.ErrInvalidMainEventInstagram:      errInvalidMainEventInstagram,
		mainevent.ErrInvalidMainEventType:           errInvalidMainEventType,
		mainevent.ErrInvalidMainEventStatus:         errInvalidMainEventStatus,
		mainevent.ErrInvalidMainEventJumlahTiket:    errInvalidMainEventJumlahTiket,
		mainevent.ErrInvalidMainEventImageURI:       errInvalidMainEventImageURI,
		mainevent.ErrTicketAlreadyCheckedIn:         errTicketAlreadyCheckedIn,
		mainevent.ErrTicketNotFound:                 errTicketNotFound,
		mainevent.ErrTicketNotYetPaid:               errTicketNotYetPaid,
		mainevent.ErrAllTicketAlreadyCheckedIn:      errAllTicketAlreadyCheckedIn,
		mainevent.ErrPaymentNotSettlement:           errPaymentNotSettlement,
		mainevent.ErrEarlyBirdTicketSoldOut:         errEarlyBirdTicketSoldOut,
		mainevent.ErrTicketNotAlreadyAccepted:       errTicketNotAlreadyAccepted,
	}
)
