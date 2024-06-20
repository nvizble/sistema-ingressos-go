package domain

import (
	"time"
)

type Rating string

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
	return nil, nil
}
