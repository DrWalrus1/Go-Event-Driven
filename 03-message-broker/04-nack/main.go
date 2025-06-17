package main

import (
	"context"

	"github.com/ThreeDotsLabs/watermill/message"
)

type AlarmClient interface {
	StartAlarm() error
	StopAlarm() error
}

func ConsumeMessages(sub message.Subscriber, alarmClient AlarmClient) {
	messages, err := sub.Subscribe(context.Background(), "smoke_sensor")
	if err != nil {
		panic(err)
	}

	for msg := range messages {
		payload := string(msg.Payload)
		var err error
		if payload == "0" { // no smoke detected
			err = alarmClient.StopAlarm()
		} else if payload == "1" { // smoke detected
			err = alarmClient.StartAlarm()
		}
		if err != nil {
			msg.Nack()
		} else {
			msg.Ack()
		}
	}
}
