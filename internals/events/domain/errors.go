package domain

import (
	err "errors"
)

var (
	ErrSpotNameIsTooShort           = err.New("Spot name must be at least 2 characters long")
	ErrEventNotFound                = err.New("event not found")
	ErrSpotNotFound                 = err.New("spot not found")
	ErrSpotNameMustStartWithALetter = err.New("spot name must start with a letter")
	ErrSpotAlreadyReserved          = err.New("spot already reserved")
	ErrSpotAlreadyAvailable         = err.New("spot already available")
	ErrNameIsRequired               = err.New("name is required")
	ErrLocationIsRequired           = err.New("location is required")
	ErrDateIsRequired               = err.New("date is required")
	ErrRatingIsRequired             = err.New("rating is required")
	ErrOrganizerIsRequired          = err.New("organizer is required")
	ErrImageUrlIsRequired           = err.New("image url is required")
	ErrDescriptionIsRequired        = err.New("description is required")
	ErrDateIsInThePast              = err.New("date must be in the future")
	ErrCapacityIsZero               = err.New("capacity must be bigger then zero")
	ErrTicketPriceIsNegative        = err.New("ticket price must not be negative")
	ErrTicketsIsRequired            = err.New("tickets are required")
	ErrSpotsIsRequired              = err.New("spots are required")
)
