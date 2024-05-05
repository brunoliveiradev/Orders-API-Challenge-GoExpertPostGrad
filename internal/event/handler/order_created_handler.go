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

func (h *OrderCreatedHandler) HandleOrderCreatedEvent(event events.EventInterface, wg *sync.WaitGroup) error {
	defer wg.Done()
	log.Println("Order created:", event.GetPayload())

	jsonOutput, err := json.Marshal(event.GetPayload())
	if err != nil {
		log.Println("Error marshalling OrderCreated event payload")
		return err
	}

	msg := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	// Publish the event to the RabbitMQ
	err = h.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // routing key
		false,        // mandatory
		false,        // immediate
		msg,
	)
	if err != nil {
		log.Println("Error publishing OrderCreated event")
		return err
	}
	log.Println("OrderCreated event published")
	return nil
}
