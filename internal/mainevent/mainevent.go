package mainevent

import (
	"context"
	"time"
)

type Service interface {
	// ReplaceMainEventByEmail replace all mainevent
	// with the given email
	ReplaceMainEventByEmail(ctx context.Context, mainevent MainEvent) (int64, error)

	// GetMainEventByID returns a mainevent with the given
	// mainevent ID.
	GetMainEventByID(ctx context.Context, maineventID int64, nomorTiket string) (MainEvent, error)

	// GetAllMainEvents returns all mainevent and filter by status and tanggal.
	GetAllMainEvents(ctx context.Context, filter GetAllMainEventsFilter) ([]MainEvent, error)

	// UpdateCheckInStatus returns a ticket number and status with the giveb
	// mainevent ID and ticket number
	UpdateCheckInStatus(ctx context.Context, id int64, nomorTiket string) (string, error)

	// UpdatePaymentStatus update the payment status in DB
	// and send a email after completed payment
	UpdatePaymentStatus(ctx context.Context, reqMainEvent MainEvent) error
}

// MainEvent is a mainevent.

type MainEvent struct {
	ID                int64
	Nama              string
	Disabilitas       Disability
	NomorIdentitas    string
	AsalInstitusi     string
	Email             string
	NomorTelepon      string
	Type              Type
	JumlahTiket       int
	TotalHarga        int64
	Status            Status
	OrderID           string
	ImageURI          string
	NomorTiket        []string
	CheckInStatus     bool
	CheckInNomorTiket []string
	CreateTime        time.Time
	UpdateTime        time.Time
}

type GetAllMainEventsFilter struct {
	Type   Type
	Status Status
}

// Type denotes type of a schedule.
type Type int

// Followings are the known type.
const (
	TypeUnknown   Type = 0
	TypeEarlyBird Type = 1
	TypePresale1  Type = 2
	TypePresale2  Type = 3
)

var (
	// TypeList is a list of valid type.
	TypeList = map[Type]struct{}{
		TypeEarlyBird: {},
		TypePresale1:  {},
		TypePresale2:  {},
	}

	// typeName maps type to it's string representation.
	typeName = map[Type]string{
		TypeEarlyBird: "early-bird",
		TypePresale1:  "presale-1",
		TypePresale2:  "presale-2",
	}
)

// Value returns int value of a Type type.
func (s Type) Value() int {
	return int(s)
}

// String returns string representaion of a Type type.
func (s Type) String() string {
	return typeName[s]
}

// Status denotes status of a schedule.
type Status int

// Followings are the known status.
const (
	StatusUnknown    Status = 0
	StatusPending    Status = 1
	StatusUnpaid     Status = 2
	StatusSettlement Status = 3
)

var (
	// StatusList is a list of valid status.
	StatusList = map[Status]struct{}{
		StatusPending:    {},
		StatusUnpaid:     {},
		StatusSettlement: {},
	}

	// StatusName maps status to it's string representation.
	statusName = map[Status]string{
		StatusPending:    "pending",
		StatusUnpaid:     "unpaid",
		StatusSettlement: "settlement",
	}
)

// Value returns int value of a status type.
func (s Status) Value() int {
	return int(s)
}

// String returns string representaion of a status type.
func (s Status) String() string {
	return statusName[s]
}

type Disability int

const (
	DisabilityUnknown      Disability = 0
	SensoricDisability     Disability = 1
	PyshicDisability       Disability = 2
	IntellectualDisability Disability = 3
	MentalDisability       Disability = 4
)

var (
	DisabilityList = map[Disability]struct{}{
		SensoricDisability:     {},
		PyshicDisability:       {},
		IntellectualDisability: {},
		MentalDisability:       {},
	}

	disabilityName = map[Disability]string{
		SensoricDisability:     "Disabilitas Sensorik",
		PyshicDisability:       "Disabilitas Fisik",
		IntellectualDisability: "Disabilitas Intelektual",
		MentalDisability:       "Disabilitas Mental",
	}
)

func (d Disability) Value() int {
	return int(d)
}

func (d Disability) String() string {
	return disabilityName[d]
}
