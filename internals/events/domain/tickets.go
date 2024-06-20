package domain

type TicketType string

const (
	FullPrice TicketType = "full"
	HalfPrice TicketType = "half"
	Courtesy  TicketType = "courtesy"
)

type Ticket struct {
	ID      string
	EventID string
	Price   int
	Type    TicketType
	Spot    *Spot
}
