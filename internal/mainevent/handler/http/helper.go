package http

import "github.com/tedxub2023/internal/mainevent"

// formatMainEvent formats the given transaction into the respective
// HTTP-format objecm.
func formatMainEvent(m mainevent.MainEvent) (mainEventHTTP, error) {
	typeStr := m.Type.String()
	statusStr := m.Status.String()

	return mainEventHTTP{
		ID:                &m.ID,
		Nama:              &m.Nama,
		Disabilitas:       &m.Disabilitas,
		NomorIdentitas:    &m.NomorIdentitas,
		AsalInstitusi:     &m.AsalInstitusi,
		Email:             &m.Email,
		NomorTelepon:      &m.NomorTelepon,
		Instagram:         &m.Instagram,
		JumlahTiket:       &m.JumlahTiket,
		TotalHarga:        &m.TotalHarga,
		OrderID:           &m.OrderID,
		FileURI:           &m.FileURI,
		Status:            &statusStr,
		Type:              &typeStr,
		ImageURI:          &m.ImageURI,
		NomorTiket:        &m.NomorTiket,
		CheckInStatus:     &m.CheckInStatus,
		CheckInNomorTiket: &m.CheckInNomorTiket,
	}, nil
}

// parseType returns mainevent.Type
// from the given string.
func parseType(req string) (mainevent.Type, error) {
	switch req {
	case mainevent.TypeEarlyBird.String():
		return mainevent.TypeEarlyBird, nil
	case mainevent.TypePresale1.String():
		return mainevent.TypePresale1, nil
	case mainevent.TypePresale2.String():
		return mainevent.TypePresale2, nil
	}
	return mainevent.TypeUnknown, errInvalidMainEventType
}

// parseStatus returns mainevent.Status
// from the given string.
func parseStatus(req string) (mainevent.Status, error) {
	switch req {
	case mainevent.StatusPending.String():
		return mainevent.StatusPending, nil
	case mainevent.StatusUnpaid.String():
		return mainevent.StatusUnpaid, nil
	case mainevent.StatusSettlement.String():
		return mainevent.StatusSettlement, nil
	}
	return mainevent.StatusUnknown, errInvalidMainEventStatus
}
