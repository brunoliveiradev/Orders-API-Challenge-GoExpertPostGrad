package handler

import (
	"GoExpertPostGrad-Orders-Challenge/pkg/events"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
	"sync"
)

type OrderCreatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderCreatedHandler(rabbitMQChannel *amqp.Channel) *OrderCreatedHandler {
	return &OrderCreatedHandler{RabbitMQChannel: rabbitMQChannel}
}

func (h *OrderCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) error {
	defer wg.Done()
	log.Println("Event created:", event.GetPayload())

	jsonOutput, _ := json.Marshal(event.GetPayload())

	msg := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	err := h.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // routing key
		false,        // mandatory
		false,        // immediate
		msg,          // message to publish
	)
	if err != nil {
		log.Println("Error publishing event")
		return err
	}
	log.Println("Event published")
	return nil
}
