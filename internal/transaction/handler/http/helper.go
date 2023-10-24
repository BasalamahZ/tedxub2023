package http

import "github.com/tedxub2023/internal/transaction"

// formatTransaction formats the given transaction into the respective
// HTTP-format object.
func formatTransaction(t transaction.Transaction) (transactionHTTP, error) {
	tanggal := t.Tanggal.Format(dateFormat)

	return transactionHTTP{
		ID:                &t.ID,
		Nama:              &t.Nama,
		JenisKelamin:      &t.JenisKelamin,
		NomorIdentitas:    &t.NomorIdentitas,
		AsalInstitusi:     &t.AsalInstitusi,
		Domisili:          &t.Domisili,
		Email:             &t.Email,
		NomorTelepon:      &t.NomorTelepon,
		LineID:            &t.LineID,
		Instagram:         &t.Instagram,
		JumlahTiket:       &t.JumlahTiket,
		TotalHarga:        &t.TotalHarga,
		Tanggal:           &tanggal,
		OrderID:           &t.OrderID,
		StatusPayment:     &t.StatusPayment,
		ImageURI:          &t.ImageURI,
		NomorTiket:        &t.NomorTiket,
		CheckInStatus:     &t.CheckInStatus,
		CheckInNomorTiket: &t.CheckInNomorTiket,
	}, nil
}
