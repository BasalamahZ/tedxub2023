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

	// errDataNotFound is returned when the desired data is
	// not found.
	errDataNotFound = errors.New("DATA_NOT_FOUND")

	// errInvalidTransactionID is returned when the given
	// transaction id is invalid.
	errInvalidTransactionID = errors.New("INVALID_TRANSACTION_ID")

	// errInvalidTransactionNama is returned when the given
	// transaction nama is invalid.
	errInvalidTransactionNama = errors.New("INVALID_TRANSACTION_NAMA")

	// errInvalidTransactionNomorIdentitas is returned when the given
	// transaction nomor identitas is invalid.
	errInvalidTransactionNomorIdentitas = errors.New("INVALID_TRANSACTION_NOMOR_IDENTITAS")

	// errInvalidTransactionAsalInstitusi is returned when the given
	// transaction asal institusi is invalid.
	errInvalidTransactionAsalInstitusi = errors.New("INVALID_TRANSACTION_ASAL_INSITITUSI")

	// errInvalidTransactionDomisili is returned when the given
	// transaction domisili is invalid.
	errInvalidTransactionDomisili = errors.New("INVALID_TRANSACTION_DOMISILI")

	// errInvalidTransactionEmail is returned when the given
	// transaction email is invalid.
	errInvalidTransactionEmail = errors.New("INVALID_TRANSACTION_EMAIL")

	// errInvalidTransactionNomorTelepon is returned when the given
	// transaction nomor telepon is invalid.
	errInvalidTransactionNomorTelepon = errors.New("INVALID_TRANSACTION_NOMOR_TELEPON")

	// errInvalidTransactionLineID is returned when the given
	// transaction line id is invalid.
	errInvalidTransactionLineID = errors.New("INVALID_TRANSACTION_LINE_ID")

	// errInvalidTransactionInstagram is returned when the given
	// transaction instagram is invalid.
	errInvalidTransactionInstagram = errors.New("INVALID_TRANSACTION_INSTAGRAM")

	// errInvalidTransactionJenisKelamin is returned when the request
	// is not valid format
	errInvalidTransactionJenisKelamin = errors.New("INVALID_TRANSACTION_JENIS_KELAMIN")

	// errInvalidTransactionJumlahTiket is returned when the given
	// transaction jumlah tiket is invalid.
	errInvalidTransactionJumlahTiket = errors.New("INVALID_TRANSACTION_JUMLAH_TIKET")

	// errInvalidTransactionTanggal is returned when the given
	// transaction tanggal is invalid.
	errInvalidTransactionTanggal = errors.New("INVALID_TRANSACTION_TANGGAL")

	// errInvalidTransactionImageURI is returned when the given
	// transaction image uri is invalid.
	errInvalidTransactionImageURI = errors.New("INVALID_TRANSACTION_IMAGE_URI")

	// errInvalidDateFormat is returned when the given date
	// string format is invalid.
	errInvalidDateFormat = errors.New("INVALID_DATE_FORMAT")

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
)

var (
	// mapHTTPError maps service error into HTTP error that
	// categorize as bad request error.
	//
	// Internal server error-related should not be mapped
	// here, and the handler should just return `errInternal`
	// as the error instead
	mapHTTPError = map[error]error{
		transaction.ErrDataNotFound:                     errDataNotFound,
		transaction.ErrInvalidTransactionID:             errInvalidTransactionID,
		transaction.ErrInvalidTransactionNama:           errInvalidTransactionNama,
		transaction.ErrInvalidTransactionNomorIdentitas: errInvalidTransactionNomorIdentitas,
		transaction.ErrInvalidTransactionAsalInstitusi:  errInvalidTransactionAsalInstitusi,
		transaction.ErrInvalidTransactionDomisili:       errInvalidTransactionDomisili,
		transaction.ErrInvalidTransactionEmail:          errInvalidTransactionEmail,
		transaction.ErrInvalidTransactionNomorTelepon:   errInvalidTransactionNomorTelepon,
		transaction.ErrInvalidTransactionLineID:         errInvalidTransactionLineID,
		transaction.ErrInvalidTransactionInstagram:      errInvalidTransactionInstagram,
		transaction.ErrInvalidTransactionJenisKelamin:   errInvalidTransactionJenisKelamin,
		transaction.ErrInvalidTransactionTanggal:        errInvalidTransactionTanggal,
		transaction.ErrInvalidTransactionJumlahTiket:    errInvalidTransactionJumlahTiket,
		transaction.ErrInvalidTransactionImageURI:       errInvalidTransactionImageURI,
		transaction.ErrInvalidDateFormat:                errInvalidDateFormat,
		transaction.ErrTicketAlreadyCheckedIn:           errTicketAlreadyCheckedIn,
		transaction.ErrTicketNotFound:                   errTicketNotFound,
		transaction.ErrTicketNotYetPaid:                 errTicketNotYetPaid,
		transaction.ErrAllTicketAlreadyCheckedIn:        errAllTicketAlreadyCheckedIn,
		transaction.ErrPaymentNotSettlement:             errPaymentNotSettlement,
	}
)
