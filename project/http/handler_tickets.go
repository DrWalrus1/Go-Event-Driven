package http

import (
	"net/http"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/labstack/echo/v4"
)

type ticketsConfirmationRequest struct {
	Tickets []string `json:"tickets"`
}

func (h Handler) PostTicketsConfirmation(c echo.Context) error {
	var request ticketsConfirmationRequest
	err := c.Bind(&request)
	if err != nil {
		return err
	}

	for _, ticket := range request.Tickets {
		issueReceiptMessage := message.Message{
			Payload: message.Payload(ticket),
		}
		h.Publisher.Publish("issue-receipt", &issueReceiptMessage)
		ticketToPrintMessage := message.Message{
			Payload: message.Payload(ticket),
		}
		h.Publisher.Publish("append-to-tracker", &ticketToPrintMessage)
	}

	return c.NoContent(http.StatusOK)
}
