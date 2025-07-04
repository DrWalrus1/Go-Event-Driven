package main

import (
	"context"
	"log"
	"os"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/redis/go-redis/v9"
)

func main() {
	logger := watermill.NewSlogLogger(nil)
	router := message.NewDefaultRouter(logger)

	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	sub, err := redisstream.NewSubscriber(redisstream.SubscriberConfig{
		Client: rdb,
	}, logger)
	if err != nil {
		panic(err)
	}

	router.AddNoPublisherHandler(
		"no_publisher_handler",
		"temperature-fahrenheit",
		sub,
		func(msg *message.Message) error {
			log.Printf("Temperature read: %s\n", string(msg.Payload))
			return nil
		},
	)

	if err := router.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
}
