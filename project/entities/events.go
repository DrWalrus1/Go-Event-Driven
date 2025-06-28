package entities

import (
	"time"

	"github.com/ThreeDotsLabs/watermill"
)

type MessageHeader struct {
	ID          string    `json:"id"`
	PublishedAt time.Time `json:"published_at"`
}

func NewHeader() MessageHeader {
	return MessageHeader{
		ID:          watermill.NewUUID(),
		PublishedAt: time.Now(),
	}
}

type TicketBookingConfirmed struct {
	Header MessageHeader `json:"header"`

	TicketID      string `json:"ticket_id"`
	CustomerEmail string `json:"customer_email"`
	Price         Money  `json:"price"`
}

type TicketBookingCanceled struct {
	Header        MessageHeader `json:"header"`
	TicketID      string        `json:"ticket_id"`
	CustomerEmail string        `json:"customer_email"`
	Price         Money         `json:"price"`
}
