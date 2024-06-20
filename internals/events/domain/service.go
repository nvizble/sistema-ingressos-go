package domain

import (
	err "errors"
	"fmt"
)

type SpotService struct{}

var (
	ErrInvalidAmount = err.New("the amount of spots must be greater then zero")
)

func NewSpotService() *SpotService {
	return &SpotService{}
}

func (s *SpotService) GenerateSpots(e *Event, amount int) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	for i := range amount {
		SpotName := fmt.Sprintf("%c%d", 'A'+i/10, i%10+1)
		Spot, err := NewSpot(e, SpotName)
		if err != nil {
			return err
		}

		e.Spots = append(e.Spots, *Spot)
	}

	return nil
}
