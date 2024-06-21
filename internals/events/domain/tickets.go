package domain

import (
	err "errors"
)

type TicketType string

const (
	FullPrice TicketType = "full"
	HalfPrice TicketType = "half"
	Courtesy  TicketType = "courtesy"
)

var (
	ErrTicketIdIsRequired = err.New("ticket id is required")
)

type Ticket struct {
	ID      string
	EventID string
	Price   int
	Type    TicketType
	Spot    *Spot
}

func (t *Ticket) Validate() error {
	if t.ID == "" {
		return ErrTicketIdIsRequired
	}

	return nil
}

func isValidTicketTye(t TicketType) bool {
	return t == FullPrice || t == HalfPrice || t == Courtesy
}
