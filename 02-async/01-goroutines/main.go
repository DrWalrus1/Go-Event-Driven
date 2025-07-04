package main

import (
	"log"
	"time"
)

type User struct {
	Email string
}

type UserRepository interface {
	CreateUserAccount(u User) error
}

type NotificationsClient interface {
	SendNotification(u User) error
}

type NewsletterClient interface {
	AddToNewsletter(u User) error
}

type Handler struct {
	repository          UserRepository
	newsletterClient    NewsletterClient
	notificationsClient NotificationsClient
}

func NewHandler(
	repository UserRepository,
	newsletterClient NewsletterClient,
	notificationsClient NotificationsClient,
) Handler {
	return Handler{
		repository:          repository,
		newsletterClient:    newsletterClient,
		notificationsClient: notificationsClient,
	}
}

func (h Handler) SignUp(u User) error {
	// This is a critical operation, so it should be done synchronously.
	if err := h.repository.CreateUserAccount(u); err != nil {
		return err
	}

	go func() {
		for {
			if err := h.newsletterClient.AddToNewsletter(u); err != nil {
				log.Printf("failed to add user to the newsletter: %v", err)
				time.Sleep(1 * time.Second)
				continue
			}
			break
		}
	}()

	go func() {
		for {
			if err := h.notificationsClient.SendNotification(u); err != nil {
				log.Printf("failed to add user to the newsletter: %v", err)
				time.Sleep(1 * time.Second)
				continue
			}
			break
		}
	}()

	return nil
}
