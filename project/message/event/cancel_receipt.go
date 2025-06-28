package event

import (
	"context"
	"fmt"
	"log/slog"
	"tickets/entities"
)

func (h Handler) CancelReceipt(ctx context.Context, event entities.TicketBookingCanceled) error {
	slog.Info("Issuing receipt")

	err := h.spreadsheetsAPI.AppendRow(
		ctx,
		"tickets-to-refund",
		[]string{event.TicketID, event.CustomerEmail, event.Price.Amount, event.Price.Currency},
	)
	if err != nil {
		return fmt.Errorf("failed to issue receipt: %w", err)
	}

	return nil
}
