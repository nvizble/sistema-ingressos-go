package domain

import (
	err "errors"
	"github.com/google/uuid"
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusReserved  SpotStatus = "reserved"
)

var (
	ErrSpotNameIsTooShort           = err.New("spot name must be at least 2 characters long")
	ErrSpotIdIsRequired             = err.New("spot id is required")
	ErrSpotEventNameIsRequired      = err.New("spot event name is required")
	ErrSpotStatusIsRequired         = err.New("spot status is required")
	ErrSpotTicketIdIsRequired       = err.New("spot ticket id is required")
	ErrSpotNameMustEndWithNumber    = err.New("spot name must end with a number")
	ErrSpotNotFound                 = err.New("spot not found")
	ErrSpotNameMustStartWithALetter = err.New("spot name must start with a letter")
	ErrSpotAlreadyReserved          = err.New("spot already reserved")
	ErrSpotNameIsRequired           = err.New("spot name is required")
)

type Spot struct {
	ID       string
	EventID  string
	Name     string
	Status   SpotStatus
	TicketiD string
}

func (s *Spot) Validate() error {
	if s.Name == "" {
		return ErrSpotNameIsRequired
	}

	if len(s.Name) < 2 {
		return ErrSpotNameIsTooShort
	}

	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotNameMustStartWithALetter
	}

	if s.Name[len(s.Name)-1] < '0' || s.Name[len(s.Name)-1] > '9' {
		return ErrSpotNameMustEndWithNumber
	}

	return nil
}

func NewSpot(event *Event, name string) (*Spot, error) {
	Spot := &Spot{
		ID:      uuid.New().String(),
		EventID: event.ID,
		Name:    name,
		Status:  SpotStatusAvailable,
	}

	if err := Spot.Validate(); err != nil {
		return nil, err
	}

	return Spot, nil
}

func (s *Spot) Reserve(ticketID string) error {
	if s.Status == SpotStatusReserved {
		return ErrSpotAlreadyReserved
	}

	s.Status = SpotStatusReserved
	s.TicketiD = ticketID

	return nil
}
