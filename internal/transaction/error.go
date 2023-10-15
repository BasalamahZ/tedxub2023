package transaction

import "errors"

// Followings are the known errors returned from transaction.
var (
	// ErrDataNotFound is returned when the wanted data is
	// not found.
	ErrDataNotFound = errors.New("data not found")

	// ErrInvalidTransactionID is returned when the given transaction
	// id is invalid.
	ErrInvalidTransactionID = errors.New("invalid transaction id")

	// ErrInvalidTransactionNama is returned when the given transaction
	// nama is invalid.
	ErrInvalidTransactionNama = errors.New("invalid transaction nama")

	// ErrInvalidTransactionNomorIdentitas is returned when the given transaction
	// nomor identitas is invalid.
	ErrInvalidTransactionNomorIdentitas = errors.New("invalid transaction nomor identitas")

	// ErrInvalidTransactionAsalInstitusi is returned when the given transaction
	// asal institusi is invalid.
	ErrInvalidTransactionAsalInstitusi = errors.New("invalid transaction asal institusi")

	// ErrInvalidTransactionDomisili is returned when the given transaction
	// domisili is invalid.
	ErrInvalidTransactionDomisili = errors.New("invalid transaction domisili")

	// ErrInvalidTransactionEmail is returned when the given transaction
	// email is invalid.
	ErrInvalidTransactionEmail = errors.New("invalid transaction email")

	// ErrInvalidTransactionNomorTelepon is returned when the given transaction
	// nomor telepon is invalid.
	ErrInvalidTransactionNomorTelepon = errors.New("invalid transaction nomor telepon")

	// ErrInvalidTransactionLineID is returned when the given transaction
	// line id is invalid.
	ErrInvalidTransactionLineID = errors.New("invalid transaction line id")

	// ErrInvalidTransactionInstagram is returned when the given transaction
	// instagram is invalid.
	ErrInvalidTransactionInstagram = errors.New("invalid transaction instagram")

	// errInvalidTransactionJenisKelamin is returned when the given transaction
	// jenis kelamin is invalid.
	ErrInvalidTransactionJenisKelamin = errors.New("invalid transaction jenis kelamin")

	// errInvalidTransactionJumlahTiket is returned when the given transaction
	// jumlah tiket is invalid.
	ErrInvalidTransactionJumlahTiket = errors.New("invalid transaction jumlah tiket")

	// errInvalidTransactionTanggal is returned when the given transaction
	// tanggal is invalid.
	ErrInvalidTransactionTanggal = errors.New("invalid transaction tanggal")

	// ErrInvalidDateFormat is returned when the given date
	// format is invalid.
	ErrInvalidDateFormat = errors.New("invalid date format")
)
