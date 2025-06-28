package main

import (
	"encoding/json"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type MessageHeader struct {
	Id         string `json:"id"`
	EventName  string `json:"event_name"`
	OccurredAt string `json:"occurred_at"`
}

func NewMessageHeader(id string, eventName string) MessageHeader {
	return MessageHeader{
		Id:         id,
		EventName:  eventName,
		OccurredAt: time.Now().Format(time.RFC3339),
	}
}

type ProductOutOfStock struct {
	Header    MessageHeader `json:"header"`
	ProductID string        `json:"product_id"`
}

type ProductBackInStock struct {
	Header    MessageHeader `json:"header"`
	ProductID string        `json:"product_id"`
	Quantity  int           `json:"quantity"`
}

type Publisher struct {
	pub message.Publisher
}

func NewPublisher(pub message.Publisher) Publisher {
	return Publisher{
		pub: pub,
	}
}

func (p Publisher) PublishProductOutOfStock(productID string) error {
	event := ProductOutOfStock{
		Header:    NewMessageHeader(watermill.NewUUID(), "ProductOutOfStock"),
		ProductID: productID,
	}

	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := message.NewMessage(watermill.NewUUID(), payload)

	return p.pub.Publish("product-updates", msg)
}

func (p Publisher) PublishProductBackInStock(productID string, quantity int) error {
	event := ProductBackInStock{
		Header:    NewMessageHeader(watermill.NewUUID(), "ProductBackInStock"),
		ProductID: productID,
		Quantity:  quantity,
	}

	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := message.NewMessage(watermill.NewUUID(), payload)

	return p.pub.Publish("product-updates", msg)
}
