package http

import "github.com/tedxub2023/internal/mainevent"

// formatMainEvent formats the given transaction into the respective
// HTTP-format objecm.
func formatMainEvent(m mainevent.MainEvent) (mainEventHTTP, error) {
	typeStr := m.Type.String()
	statusStr := m.Status.String()
	disabilityStr := m.Disabilitas.String()

	return mainEventHTTP{
		ID:                &m.ID,
		Nama:              &m.Nama,
		Disabilitas:       &disabilityStr,
		NomorIdentitas:    &m.NomorIdentitas,
		AsalInstitusi:     &m.AsalInstitusi,
		Email:             &m.Email,
		NomorTelepon:      &m.NomorTelepon,
		JumlahTiket:       &m.JumlahTiket,
		TotalHarga:        &m.TotalHarga,
		OrderID:           &m.OrderID,
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
	case mainevent.TypePresale.String():
		return mainevent.TypePresale, nil
	case mainevent.TypeNormalSale.String():
		return mainevent.TypeNormalSale, nil
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

// parseDisability returns mainevent.Disability
// from the given string.
func parseDisability(req string) (mainevent.Disability, error) {
	switch req {
	case mainevent.SensoricDisability.String():
		return mainevent.SensoricDisability, nil
	case mainevent.PyshicDisability.String():
		return mainevent.PyshicDisability, nil
	case mainevent.IntellectualDisability.String():
		return mainevent.IntellectualDisability, nil
	case mainevent.MentalDisability.String():
		return mainevent.MentalDisability, nil
	case mainevent.NoneDisability.String():
		return mainevent.NoneDisability, nil
	}
	return mainevent.DisabilityUnknown, errInvalidMainEventDisability
}
