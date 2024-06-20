package domain

import (
	err "errors"
	"time"
)

type Rating string

// EVENT
var (
	ErrEventNotFound         = err.New("event not found")
	ErrNameIsRequired        = err.New("name is required")
	ErrLocationIsRequired    = err.New("location is required")
	ErrDateIsRequired        = err.New("date is required")
	ErrRatingIsRequired      = err.New("rating is required")
	ErrOrganizerIsRequired   = err.New("organizer is required")
	ErrImageUrlIsRequired    = err.New("image url is required")
	ErrDescriptionIsRequired = err.New("description is required")
	ErrDateIsInThePast       = err.New("date must be in the future")
	ErrCapacityIsZero        = err.New("capacity must be bigger then zero")
	ErrTicketPriceIsNegative = err.New("ticket price must not be negative")
	ErrTicketsIsRequired     = err.New("tickets are required")
	ErrSpotsIsRequired       = err.New("spots are required")
)

const (
	RatingL  Rating = "L"
	Rating10 Rating = "L10"
	Rating12 Rating = "L12"
	Rating14 Rating = "L14"
	Rating16 Rating = "L16"
	Rating18 Rating = "L18"
)

type Event struct {
	ID          string
	Name        string
	Location    string
	Date        time.Time
	Rating      Rating
	Organizer   string
	ImageUrl    string
	Description string
	Capacity    int
	TicketPrice int
	Tickets     []Ticket
	Spots       []Spot
}

func (e *Event) AddSpots(name string) (*Spot, error) {
	Spot, err := NewSpot(e, name)
	if err != nil {
		return nil, err
	}
	e.Spots = append(e.Spots, *Spot)
	return Spot, nil
}
