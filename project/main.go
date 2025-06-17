package main

import (
	"context"
	"os"

	"github.com/ThreeDotsLabs/go-event-driven/v2/common/clients"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/redis/go-redis/v9"

	"tickets/adapters"
	"tickets/service"
)

func main() {
	logger := watermill.NewSlogLogger(nil)
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	apiClients, err := clients.NewClients(os.Getenv("GATEWAY_ADDR"), nil)
	if err != nil {
		panic(err)
	}

	spreadsheetsAPI := adapters.NewSpreadsheetsAPIClient(apiClients)
	receiptsService := adapters.NewReceiptsServiceClient(apiClients)

	err = service.New(
		spreadsheetsAPI,
		receiptsService,
		logger,
		rdb,
	).Run(context.Background())
	if err != nil {
		panic(err)
	}
}
