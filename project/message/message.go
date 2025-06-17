package message

import (
	"context"

	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
)

func RunReceiptsSubscription(redisStreamSubscriber *redisstream.Subscriber, topic string, Action func(context.Context, string) error) {
	msgs, err := redisStreamSubscriber.Subscribe(context.Background(), topic)
	if err != nil {
		panic(err)
	}

	for msg := range msgs {
		ticketId := string(msg.Payload)
		err := Action(msg.Context(), ticketId)
		if err != nil {
			msg.Nack()
		} else {
			msg.Ack()
		}
	}

}

func RunTicketsToPrintSubscription(redisStreamSubscriber *redisstream.Subscriber, topic string, Action func(context.Context, string, []string) error) {
	msgs, err := redisStreamSubscriber.Subscribe(context.Background(), topic)
	if err != nil {
		panic(err)
	}

	for msg := range msgs {
		ticketId := string(msg.Payload)
		err := Action(msg.Context(), "tickets-to-print", []string{ticketId})
		if err != nil {
			msg.Nack()
		} else {
			msg.Ack()
		}
	}

}
