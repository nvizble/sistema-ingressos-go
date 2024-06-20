package domain

import (
	"time"
)

func (s *Spot) Validate() error {
	if s.Name == "" {
		return ErrNameIsRequired
	}

	if len(s.Name) < 2 {
		return ErrSpotNameIsTooShort
	}

	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotNameMustStartWithALetter

	}

	return nil
}

func (e *Event) Validate() error {
	if e.Name == "" {
		return ErrNameIsRequired
	}

	if e.Location == "" {
		return ErrLocationIsRequired
	}

	if e.Date.IsZero() {
		return ErrDateIsRequired
	}

	if e.Date.Before(time.Now()) {
		return ErrDateIsInThePast
	}

	if e.Rating == "" {
		return ErrRatingIsRequired
	}

	if e.Organizer == "" {
		return ErrOrganizerIsRequired
	}

	if e.ImageUrl == "" {
		return ErrImageUrlIsRequired
	}

	if e.Description == "" {
		return ErrDescriptionIsRequired
	}

	if e.Capacity <= 0 {
		return ErrCapacityIsZero
	}

	if e.TicketPrice < 0 {
		return ErrTicketPriceIsNegative
	}

	if len(e.Tickets) == 0 {
		return ErrTicketsIsRequired
	}

	if len(e.Spots) == 0 {
		return ErrSpotsIsRequired
	}

	return nil
}
